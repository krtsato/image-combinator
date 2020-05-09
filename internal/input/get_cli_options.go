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
type CliOption struct {
	Platform string // 画像の投稿先
	Usecase  string // 画像の用途
}

// 関数 keyExists のデフォルト引数を空文字にする
type keyExistsArgs struct {
	platform string
	usecase  string
}

// 対話型 CLI を使う場合に参照するマップ
// platform : usecase : {width, height}
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
func keyExists(args keyExistsArgs) bool {
	platform := args.platform
	if platform == "" {
		return false
	}

	if args.usecase == "" {
		_, exists := patternMap[platform]
		return exists
	}

	_, exists := patternMap[platform][args.usecase]
	return exists
}

// 対話型 CLI で platform 文字列を取得する
func getPlatform() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	// platform の登録を確認する
	inputText := strings.ToLower(scanner.Text())
	if keyExists(keyExistsArgs{platform: inputText}) {
		return inputText, nil
	}

	return "", errors.New("Error: \"" + inputText + "\" is not registered with this application.")
}

// 対話型 CLI で usecase 文字列を取得する
func getUsecase(platform string) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	// usecase の登録を確認する
	inputText := strings.ToLower(scanner.Text())
	if keyExists(keyExistsArgs{platform: platform, usecase: inputText}) {
		return inputText, nil
	}

	return "", errors.New("Error: \"" + inputText + "\" is not registered with this application.")
}

// マップの key 一覧を文字列で返す
func getMapKeys(rawMap interface{}) (string, error) {
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
	return rawKeys, nil
}

// 対話型 CLI で platform の入力を求める
func askPlatform() error {
	mapKeys, err := getMapKeys(patternMap)
	if err != nil {
		return err
	}
	fmt.Println("\nEnter the platform where you will submit images. [" + mapKeys + "]")
	return nil
}

// 対話型 CLI で usecase の入力を求める
func askUsecase(platform string) error {
	usecaseMap := patternMap[platform]
	mapKeys, err := getMapKeys(usecaseMap)
	if err != nil {
		return err
	}
	fmt.Println("\nEnter the usecase of output images. [" + mapKeys + "]")
	return nil
}

func GetCliOptions() (CliOption, error) {
	// CLI フラグで直接指定する場合
	// デフォルトはフラグ未指定
	var pFlag = flag.String("p", "", `The platform you want to post a image
	Assign "twitter" or "youtube".`)
	var uFlag = flag.String("u", "", `The usecase in your choosing platform
	twitter: "post" or "header"
	youtube: "screen" or "thumbnail"`)
	flag.Parse()

	// フラグが不正・未指定の場合
	// 対話型 CLI に切り替える
	// 標準入力から platform を取得する
	if !keyExists(keyExistsArgs{platform: *pFlag}) {
		if err := askPlatform(); err != nil {
			return CliOption{}, err
		}

		platform, err := getPlatform()
		if err != nil {
			return CliOption{}, err
		}
		*pFlag = platform
	}

	// フラグが不正・未指定の場合
	// 標準入力から usecase を取得する
	// platform に応じて usecase が替わる
	if keyExists(keyExistsArgs{platform: *pFlag, usecase: *uFlag}) {
		if err := askUsecase(*pFlag); err != nil {
			return CliOption{}, err
		}

		usecase, err := getUsecase(*pFlag)
		if err != nil {
			return CliOption{}, err
		}
		*uFlag = usecase
	}

	return CliOption{*pFlag, *uFlag}, nil
}
