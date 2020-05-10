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

// platform や usecase の登録を確認する
func _mapKeysExist(options *CliOptions) (bool, bool) {
	platform := options.Platform
	usecaseExists := false

	_, platformExists := patternMap[platform]
	if platformExists {
		_, uExists := patternMap[platform][options.Usecase]
		usecaseExists = uExists
	}
	return platformExists, usecaseExists
}

func _getMapKeys(rawMap interface{}) (string, error) {
	refMap := reflect.ValueOf(rawMap)
	if refMap.Kind() != reflect.Map {
		return "", errors.New("Error: cannot get map keys because the argment type")
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
func _askMapKey(rawMap interface{}) error {
	refMap := reflect.ValueOf(rawMap)
	if refMap.Kind() != reflect.Map {
		return errors.New("Error: cannot get map keys because the argment type")
	}

	mapKeys, err := _getMapKeys(rawMap)
	if err != nil {
		return err
	}
	fmt.Println("\nEnter the platform of output images. [" + mapKeys + "]")
	return nil
}

// 対話型 CLI で platform 文字列を更新する
func _updateCliOptions(rawMap interface{}, options *CliOptions) error {
	refMap := reflect.ValueOf(rawMap)
	if refMap.Kind() != reflect.Map {
		return errors.New("Error: cannot get map keys because the argment type")
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
			return errors.New("Error: The argment has invalid type of map")
		}
	}

	return errors.New("Error: \"" + inputText + "\" is not registered with this application.")
}

func _GetCliOptions() (CliOptions, error) {
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
	pExists, uExists := _mapKeysExist(cliOptions)
	if pExists && uExists {
		return *cliOptions, nil
	}

	// platform は登録済みかつ，usecase が不正・未指定の場合
	if pExists && !uExists {
		// usecase の入力を求める
		if err := _askMapKey(patternMap[cliOptions.Platform]); err != nil {
			return CliOptions{}, err
		}

		// update
		if err := _updateCliOptions(patternMap[cliOptions.Platform], cliOptions); err != nil {
			return CliOptions{}, err
		}
	}

	// platform が不正・未指定の場合
	// ask
	if err := _askMapKey(patternMap); err != nil {
		return CliOptions{}, err
	}

	// update
	if err := _updateCliOptions(patternMap, cliOptions); err != nil {
		return CliOptions{}, err
	}

	// usecase が登録されていた場合
	if pExists, uExists := _mapKeysExist(cliOptions); pExists && uExists {
		return *cliOptions, nil
	}

	// usecase が登録されていない場合
	if err := _askMapKey(patternMap[cliOptions.Platform]); err != nil {
		return CliOptions{}, err
	}

	// update
	if err := _updateCliOptions(patternMap[cliOptions.Platform], cliOptions); err != nil {
		return CliOptions{}, err
	}

	// 最終的に適値で更新された場合
	return *cliOptions, nil
}
