package input

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type CliOption struct {
	Platform string // 画像の投稿先
	Usecase  string // 画像の用途
}

// 対話型 CLI を使う場合に参照するマップ
// platform: usecase
var patternMap = map[string][]string{
	"twitter": {"post", "header"},
	"youtube": {"screen", "thumbnail"},
}

// 対話型 CLI で platform 文字列を取得する
func getPlatform() (string, error) {
	var platform string
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputText := strings.ToLower(scanner.Text())

	// platform の登録を確認する
	if _, isPlatform := patternMap[inputText]; isPlatform {
		platform = inputText
	} else {
		err = errors.New("Error: \"" + inputText + "\" is not register with this application.")
	}

	// スキャンエラーに対応する
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return platform, err
}

// 対話型 CLI で usecase 文字列を取得する
func getUsecase(platform string) (string, error) {
	var (
		usecase string
		err     error
	)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputText := strings.ToLower(scanner.Text())

	// usecase の登録を確認する
	usecaseArr := patternMap[platform]
	for i, v := range usecaseArr {
		if inputText == v {
			usecase = inputText
			break
		}

		// 未登録の場合はエラーが発生する
		if i == len(usecaseArr) {
			err = errors.New("Error: \"" + inputText + "\" is not register with this application.")
		}
	}

	// スキャンエラーに対応する
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return usecase, err
}

// 対話型 CLI で platform の入力を求める
func askPlatform() {
	var platformArr []string
	for k := range patternMap {
		platformArr = append(platformArr, k)
	}
	platform := strings.Join(platformArr, " / ")
	fmt.Println("画像の投稿先を入力して下さい [" + platform + "]")
}

// 対話型 CLI で usecase の入力を求める
func askUsecase(platform string) {
	usecaseArr := patternMap[platform]
	usecase := strings.Join(usecaseArr, " / ")
	fmt.Println("画像の用途を入力して下さい [" + usecase + "]")
}

func GetCliOptions() (CliOption, error) {
	// CLI フラグで直接指定する
	// デフォルトはフラグ未指定
	var Platform = flag.String("p", "", `The platform you want to post a image
	Assign "twitter" or "youtube".`)
	var Usecase = flag.String("u", "", `The usecase in your choosing platform
	twitter: "post" or "header"
	youtube: "screen" or "thumbnail"`)
	flag.Parse()

	// フラグ未指定の場合は対話型 CLI に切り替える
	// 標準入力から platform を取得する
	if *Platform == "" {
		askPlatform()
		platform, err := getPlatform()
		if err != nil {
			return CliOption{}, err
		}
		*Platform = platform
	}

	// 標準入力から usecase を取得する
	// platform によって usecase が替わる
	if *Usecase == "" {
		askUsecase(*Platform)
		usecase, err := getUsecase(*Platform)
		if err != nil {
			return CliOption{}, err
		}
		*Usecase = usecase
	}

	return CliOption{*Platform, *Usecase}, nil
}
