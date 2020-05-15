package input

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"image-combinator/internal/calc"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// フラグ指定・対話型 CLIによってフィールド値を更新する
type CliOptions struct {
	AspectRatio string // 出力画像のアスペクト比 : 再計算コスト削減のため保持する
	Density     int    // 出力画像１枚あたりの入力画像枚数
	Platform    string // 出力画像の投稿先
	Usecase     string // 出力画像の用途
}

// platform や usecase の有効確認をする
// (true, true), (true, false), (false, false) の３通り
func mapKeysExist(platformMap platformMapType, options *CliOptions) (bool, bool) {
	platform := options.Platform
	usecaseExists := false

	// platform が適値である場合に限り usecase の検証を行う
	_, platformExists := platformMap[platform]
	if platformExists {
		_, uExists := platformMap[platform][options.Usecase]
		usecaseExists = uExists
	}
	return platformExists, usecaseExists
}

// platform や usecase を決めるマップから key の一覧を文字列で取得する
// e.g. "twitter / youtube"
func getMapKeys(rawMap interface{}) (string, error) {
	refMap := reflect.ValueOf(rawMap)
	if refMap.Kind() != reflect.Map {
		return "", errors.New("Error: The argument is invalid because of the type or zero values.")
	}

	var rawKeyArr []string
	refKeys := refMap.MapKeys()

	for _, key := range refKeys {
		var keyStr string
		if key.Kind() == reflect.Int {
			keyStr = strconv.Itoa(key.Interface().(int))
		} else {
			keyStr = key.String()
		}
		rawKeyArr = append(rawKeyArr, keyStr)
	}

	rawKeys := strings.Join(rawKeyArr, " / ")
	return rawKeys, nil
}

// 対話型 CLI で platform や usecase の入力を求める
// 各マップの key 一覧を選択肢として提示する
func askMapKey(rawMap interface{}) error {
	refMap := reflect.ValueOf(rawMap)
	if refMap.Kind() != reflect.Map {
		return errors.New("Error: The argument is invalid because of the type or zero values.")
	}

	mapKeys, err := getMapKeys(rawMap)
	if err != nil {
		return err
	}

	switch refMap.Interface().(type) {
	case platformMapType:
		fmt.Printf("\nEnter the platform where you will submit images. [%s]\n", mapKeys)
	case usecaseMapType:
		fmt.Printf("\nEnter the usecase of output images. [%s]\n", mapKeys)
	case densityMapType:
		fmt.Printf("\nEnter the quantity of materials per output image. [%s]\n", mapKeys)
	default:
		return errors.New("Error: The argument has invalid type of map.")
	}

	return nil
}

// 対話型 CLI で Platform や Usecase の文字列を更新する
// ポインタ型を更新するため戻り値は error のみ
func updateCliOptions(rawMap interface{}, options *CliOptions) error {
	refMap := reflect.ValueOf(rawMap)
	if refMap.Kind() != reflect.Map {
		return errors.New("Error: The argument is invalid because of the type or zero values.")
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return err
	}

	// 入力文字列が key となるマップが存在する場合
	var inputParam interface{}
	inputText := strings.ToLower(scanner.Text())
	if _, exist := refMap.Interface().(densityMapType); exist {
		inputParam, _ = strconv.Atoi(inputText)
	} else {
		inputParam = inputText
	}

	refKey := reflect.ValueOf(inputParam)
	refVal := refMap.MapIndex(refKey)
	if exist := refVal.IsValid(); exist {
		switch refVal.Interface().(type) {
		case usecaseMapType:
			options.Platform = inputParam.(string)
			return nil
		case screenMapType:
			options.Usecase = inputParam.(string)
			return nil
		case measureMapType:
			options.Density = inputParam.(int)
			return nil
		default:
			fmt.Println(refVal.Type())
			return errors.New("Error: The argument has invalid type.")
		}
	}

	return fmt.Errorf("Error: \"%s\" is not registered with this application.", inputText)
}

// usecase を入力・更新する
func updateUsecase(options *CliOptions) error {
	// usecase の入力を求める
	usecaseMap := PlatformMap[options.Platform]
	if err := askMapKey(usecaseMap); err != nil {
		return err
	}

	// usecase が更新できたら完了
	if err := updateCliOptions(usecaseMap, options); err != nil {
		return err
	}

	return nil
}

// aspectRatio・density を入力・更新する
func updateAspectDensity(options *CliOptions) error {
	// aspectRatio を更新する
	usecaseMap := PlatformMap[options.Platform][options.Usecase]
	aspectRatio := calc.AspectRatio(usecaseMap["width"], usecaseMap["height"])
	options.AspectRatio = aspectRatio

	// フラグで予め指定した density が適値ならば完了
	densityMap := AspectMap[aspectRatio]

	if _, dExists := densityMap[options.Density]; dExists {
		return nil
	}

	// density が不正・未指定の場合
	// density の入力を求める
	if err := askMapKey(densityMap); err != nil {
		return err
	}

	// density が更新できたら完了
	if err := updateCliOptions(densityMap, options); err != nil {
		return err
	}

	return nil
}

// ポインタ型の *cliOptions を適値で初期化する
// フラグ指定を受けた上で，不足分のオプションを対話型 CLI で補う
func InitCliOptions() (*CliOptions, error) {
	// platform と usecase のフラグを用意する
	// デフォルトではフラグを指定しない
	var pFlag = flag.String("p", "", `The platform you want to post a image
	Assign "twitter" or "youtube".`)
	var uFlag = flag.String("u", "", `The usecase in your choosing platform
	twitter: Assign "post" or "header".
	youtube: Assign "screen" or "thumbnail".`)
	var dFlag = flag.Int("d", 0, "The density of materials per output image for your choosing usecase.")
	flag.Parse()
	cliOptions := &CliOptions{AspectRatio: "", Density: *dFlag, Platform: *pFlag, Usecase: *uFlag}

	// platform と usecase を検証する
	// フラグで適値を指定した場合
	pExists, uExists := mapKeysExist(PlatformMap, cliOptions)
	if pExists && uExists {
		// density の入力・更新に成功したら完了
		if err := updateAspectDensity(cliOptions); err != nil {
			return &CliOptions{}, err
		}

		return cliOptions, nil
	}

	// フラグが不正・未指定の場合は対話型 CLI に切り替える
	// platform が適値かつ usecase は不正・未指定の場合
	if pExists && !uExists {
		// usecase の入力・更新する
		// platform と usecase  が適値であることが保証される
		if err := updateUsecase(cliOptions); err != nil {
			return &CliOptions{}, err
		}

		// density の入力・更新に成功したら完了
		if err := updateAspectDensity(cliOptions); err != nil {
			return &CliOptions{}, err
		}

		return cliOptions, nil
	}

	// platform が不正・未指定の場合
	// platform の入力を求める
	if err := askMapKey(PlatformMap); err != nil {
		return &CliOptions{}, err
	}

	// platform を更新する
	// platform が適値であることが保証される
	if err := updateCliOptions(PlatformMap, cliOptions); err != nil {
		return &CliOptions{}, err
	}

	// フラグで予め指定した usecase が適値だった場合
	// platform と usecase  が適値であることが保証される
	if pExists, uExists := mapKeysExist(PlatformMap, cliOptions); pExists && uExists {
		// density の入力・更新に成功したら完了
		if err := updateAspectDensity(cliOptions); err != nil {
			return &CliOptions{}, err
		}

		return cliOptions, nil
	}

	// フラグで予め指定した usecase が不正・未指定だった場合
	// usecase の入力・更新する
	// platform と usecase  が適値であることが保証される
	if err := updateUsecase(cliOptions); err != nil {
		return &CliOptions{}, err
	}

	// density の入力・更新に成功したら完了
	if err := updateAspectDensity(cliOptions); err != nil {
		return &CliOptions{}, err
	}

	return cliOptions, nil
}
