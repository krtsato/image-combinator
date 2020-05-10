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

// 出力画像のサイズを切り替える
type CliOptions struct {
	Platform string // 画像の投稿先
	Usecase  string // 画像の用途
}

// 対話型 CLI を使う場合に参照するマップ
// platform : usecase : [width, height]
var patternMap = map[string]map[string][]int{
	"twitter": {
		"header": {1500, 500},
		"post":   {1024, 576},
	},
	"youtube": {
		"screen":    {1920, 1080},
		"thumbnail": {1280, 720},
	},
}

// platform や usecase の登録を確認する
func mapKeysExist(options *CliOptions) (bool, bool) {
	platform := options.Platform
	usecaseExists := false

	_, platformExists := patternMap[platform]
	if platformExists {
		_, uExists := patternMap[platform][options.Usecase]
		usecaseExists = uExists
	}
	return platformExists, usecaseExists
}

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
		fmt.Println("\nEnter the platform of output images. [" + mapKeys + "]")
	case map[string][]int:
		fmt.Println("\nEnter the usecase of output images. [" + mapKeys + "]")
	default:
		return errors.New("Error: The argument has invalid type of map")
	}

	return nil
}

// 対話型 CLI で platform 文字列を更新する
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

func GetCliOptions() (CliOptions, error) {
	// CLI フラグで直接指定する場合
	// デフォルトはフラグ未指定
	var pFlag = flag.String("p", "", `The platform you want to post a image
	Assign "twitter" or "youtube".`)
	var uFlag = flag.String("u", "", `The usecase in your choosing platform
	twitter: "post" or "header"
	youtube: "screen" or "thumbnail"`)
	flag.Parse()
	cliOptions := &CliOptions{Platform: *pFlag, Usecase: *uFlag}

	// フラグが不正・未指定の場合
	// 対話型 CLI に切り替える
	pExists, uExists := mapKeysExist(cliOptions)
	if pExists && uExists {
		return *cliOptions, nil
	}

	// platform は登録済みかつ，usecase が不正・未指定の場合
	if pExists && !uExists {
		// usecase の入力を求める
		if err := askMapKey(patternMap[cliOptions.Platform]); err != nil {
			return CliOptions{}, err
		}

		// update
		if err := updateCliOptions(patternMap[cliOptions.Platform], cliOptions); err != nil {
			return CliOptions{}, err
		}

		return *cliOptions, nil
	}

	// platform が不正・未指定の場合
	// ask
	if err := askMapKey(patternMap); err != nil {
		return CliOptions{}, err
	}

	// update
	if err := updateCliOptions(patternMap, cliOptions); err != nil {
		return CliOptions{}, err
	}

	// usecase が登録されていた場合
	if pExists, uExists := mapKeysExist(cliOptions); pExists && uExists {
		return *cliOptions, nil
	}

	// usecase が登録されていない場合
	if err := askMapKey(patternMap[cliOptions.Platform]); err != nil {
		return CliOptions{}, err
	}

	// update
	if err := updateCliOptions(patternMap[cliOptions.Platform], cliOptions); err != nil {
		return CliOptions{}, err
	}

	// 最終的に適値で更新された場合
	return *cliOptions, nil
}
