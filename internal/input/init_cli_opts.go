package input

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// コマンドオプション・フラグ指定によって
// フィールド値を更新する
type CliOptions struct {
	Platform string // 出力画像の投稿先
	Usecase  string // 出力画像の用途
}

// オプションの有効値を格納したマップ
// platform : usecase : [width, height]
var PatternMap = map[string]map[string][]int{
	"twitter": {
		"header": {1500, 500},
		"post":   {1024, 576},
	},
	"youtube": {
		"screen":    {1920, 1080},
		"thumbnail": {1280, 720},
	},
}

// platform や usecase の有効確認をする
// (true, true), (true, false), (false, false) の３通り
func mapKeysExist(patternMap map[string]map[string][]int, options *CliOptions) (bool, bool) {
	platform := options.Platform
	usecaseExists := false

	// platform が適値である場合に限り usecase の検証を行う
	_, platformExists := patternMap[platform]
	if platformExists {
		_, uExists := patternMap[platform][options.Usecase]
		usecaseExists = uExists
	}
	return platformExists, usecaseExists
}

// platform や usecase を決めるマップから key の一覧を文字列で取得する
// e.g. "twitter / youtube"
func getMapKeys(rawMap interface{}) (string, error) {
	refMap := reflect.ValueOf(rawMap)
	if refMap.Kind() != reflect.Map {
		return "", errors.New("Error: The argument is invalid because of the type or zero values")
	}

	var rawKeyArr []string
	refKeys := refMap.MapKeys()
	for _, key := range refKeys {
		rawKeyArr = append(rawKeyArr, key.String())
	}

	rawKeys := strings.Join(rawKeyArr, " / ")
	fmt.Println("rawKeys : " + rawKeys)
	return rawKeys, nil
}

// 対話型 CLI で platform や usecase の入力を求める
// 各マップの key 一覧を選択肢として提示する
func askMapKey(rawMap interface{}) error {
	refMap := reflect.ValueOf(rawMap)
	if refMap.Kind() != reflect.Map {
		return errors.New("Error: The argument is invalid because of the type or zero values")
	}

	mapKeys, err := getMapKeys(rawMap)
	if err != nil {
		return err
	}

	switch refMap.Interface().(type) {
	case map[string]map[string][]int:
		fmt.Println("\nEnter the platform where you will submit images. [" + mapKeys + "]")
	case map[string][]int:
		fmt.Println("\nEnter the usecase of output images. [" + mapKeys + "]")
	default:
		return errors.New("Error: The argument has invalid type of map")
	}

	return nil
}

// 対話型 CLI で Platform や Usecase の文字列を更新する
// ポインタ型を更新するため戻り値は error のみ
func updateCliOptions(rawMap interface{}, options *CliOptions) error {
	refMap := reflect.ValueOf(rawMap)
	if refMap.Kind() != reflect.Map {
		return errors.New("Error: The argument is invalid because of the type or zero values")
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return err
	}

	// 入力文字列が key となるマップが存在する場合
	inputText := strings.ToLower(scanner.Text())
	refKey := reflect.ValueOf(inputText)
	refVal := refMap.MapIndex(refKey)
	if exist := refVal.IsValid(); exist {
		switch refVal.Interface().(type) {
		case map[string][]int:
			options.Platform = inputText
			return nil
		case []int:
			options.Usecase = inputText
			return nil
		default:
			return errors.New("Error: The argument has invalid type of map")
		}
	}

	return errors.New("Error: \"" + inputText + "\" is not registered with this application.")
}

// ポインタ型の *cliOptions を適値で初期化する
// フラグ指定を受けた上で，不足分のオプションを対話型 CLI で補う
func InitCliOptions() (CliOptions, error) {
	// platform と usecase のフラグを用意する
	// デフォルトではフラグを指定しない
	var pFlag = flag.String("p", "", `The platform you want to post a image
	Assign "twitter" or "youtube".`)
	var uFlag = flag.String("u", "", `The usecase in your choosing platform
	twitter: "post" or "header"
	youtube: "screen" or "thumbnail"`)
	flag.Parse()
	cliOptions := &CliOptions{Platform: *pFlag, Usecase: *uFlag}
	usecaseMap := PatternMap[cliOptions.Platform]

	// フラグで適値を指定した場合は完了
	pExists, uExists := mapKeysExist(PatternMap, cliOptions)
	if pExists && uExists {
		return *cliOptions, nil
	}

	// フラグが不正・未指定の場合は対話型 CLI に切り替える
	// platform が適値かつ usecase は不正・未指定の場合
	if pExists && !uExists {
		// usecase の入力を求める
		if err := askMapKey(usecaseMap); err != nil {
			return CliOptions{}, err
		}

		// usecase が更新できたら完了
		if err := updateCliOptions(usecaseMap, cliOptions); err != nil {
			return CliOptions{}, err
		}

		return *cliOptions, nil
	}

	// platform が不正・未指定の場合
	// platform の入力を求める
	if err := askMapKey(PatternMap); err != nil {
		return CliOptions{}, err
	}

	// platform を更新する
	// 今後の処理では platform が適値であることが保証される
	if err := updateCliOptions(PatternMap, cliOptions); err != nil {
		return CliOptions{}, err
	}

	// フラグで予め指定した usecase が適値だった場合
	// platform と usecase が共に適値であるため完了
	if pExists, uExists := mapKeysExist(PatternMap, cliOptions); pExists && uExists {
		return *cliOptions, nil
	}

	// フラグで予め指定した usecase が不正・未指定だった場合
	// usecase の入力を求める
	if err := askMapKey(usecaseMap); err != nil {
		return CliOptions{}, err
	}

	// usecase が更新できたら完了
	if err := updateCliOptions(usecaseMap, cliOptions); err != nil {
		return CliOptions{}, err
	}

	return *cliOptions, nil
}
