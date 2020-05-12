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

// フラグ指定・対話型 CLIによってフィールド値を更新する
type CliOptions struct {
	Density  string // 出力画像１枚あたりの入力画像枚数
	Platform string // 出力画像の投稿先
	Usecase  string // 出力画像の用途
}

// オプションの有効値を格納したマップ
// platform : usecase : size (density) : int
type formatDependency map[string]map[string]int
type usecaseDependency map[string]formatDependency
type platformDependency map[string]usecaseDependency

var PatternMap = platformDependency{
	"twitter": {
		"header": {
			"size": {
				"width":  1500,
				"height": 500,
			},
			"density": {
				"min": 3,
				"max": 75,
			},
		},
		"post": {
			"size": {
				"width":  1024,
				"height": 576,
			},
			"density": {
				"min": 10,
				"max": 144,
			},
		},
	},
	"youtube": {
		"screen": {
			"size": {
				"width":  1920,
				"height": 1080,
			},
			"density": {
				"min": 10,
				"max": 144,
			},
		},
		"thumbnail": {
			"size": {
				"width":  1280,
				"height": 720,
			},
			"density": {
				"min": 10,
				"max": 144,
			},
		},
	},
}

// platform や usecase の有効確認をする
// (true, true), (true, false), (false, false) の３通り
func mapKeysExist(patternMap platformDependency, options *CliOptions) (bool, bool) {
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
		return "", errors.New("Error: The argument is invalid because of the type or zero values.")
	}

	var rawKeyArr []string
	refKeys := refMap.MapKeys()

	for _, key := range refKeys {
		rawKeyArr = append(rawKeyArr, key.String())
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
	case platformDependency:
		fmt.Printf("\nEnter the platform where you will submit images. [%s]\n", mapKeys)
	case usecaseDependency:
		fmt.Printf("\nEnter the usecase of output images. [%s]\n", mapKeys)
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
	inputText := strings.ToLower(scanner.Text())
	refKey := reflect.ValueOf(inputText)
	refVal := refMap.MapIndex(refKey)
	if exist := refVal.IsValid(); exist {
		switch refVal.Interface().(type) {
		case usecaseDependency:
			options.Platform = inputText
			return nil
		case formatDependency:
			options.Usecase = inputText
			return nil
		default:
			fmt.Println(refVal.Type())
			return errors.New("Error: The argument has invalid type of map.!!!!")
		}
	}

	return fmt.Errorf("Error: \"%s\" is not registered with this application.", inputText)
}

// ポインタ型の *cliOptions を適値で初期化する
// フラグ指定を受けた上で，不足分のオプションを対話型 CLI で補う
func InitCliOptions() (CliOptions, error) {
	// platform と usecase のフラグを用意する
	// デフォルトではフラグを指定しない
	var pFlag = flag.String("p", "", `The platform you want to post a image
	Assign "twitter" or "youtube".`)
	var uFlag = flag.String("u", "", `The usecase in your choosing platform
	twitter: Assign "post" or "header".
	youtube: Assign "screen" or "thumbnail".`)
	var dFlag = flag.String("d", "min", `The density of materials per output image for your choosing usecase
	Assign "min" or "max".`)
	flag.Parse()
	cliOptions := &CliOptions{Platform: *pFlag, Usecase: *uFlag, Density: *dFlag}

	// density を検証する
	// フラグがデフォルトを更新する不正値の場合
	if !(*dFlag == "min" || *dFlag == "max") {
		return CliOptions{}, errors.New("Error: The argument of \"-d\" must be \"min\" or \"max\".")
	}

	// platform と usecase を検証する
	// フラグで適値を指定した場合は完了
	pExists, uExists := mapKeysExist(PatternMap, cliOptions)
	if pExists && uExists {
		return *cliOptions, nil
	}

	// フラグが不正・未指定の場合は対話型 CLI に切り替える
	// platform が適値かつ usecase は不正・未指定の場合
	if pExists && !uExists {
		// usecase の入力を求める
		if err := askMapKey(PatternMap[cliOptions.Platform]); err != nil {
			return CliOptions{}, err
		}

		// usecase が更新できたら完了
		if err := updateCliOptions(PatternMap[cliOptions.Platform], cliOptions); err != nil {
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
	if err := askMapKey(PatternMap[cliOptions.Platform]); err != nil {
		return CliOptions{}, err
	}

	// usecase が更新できたら完了
	if err := updateCliOptions(PatternMap[cliOptions.Platform], cliOptions); err != nil {
		return CliOptions{}, err
	}

	return *cliOptions, nil
}
