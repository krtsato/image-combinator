package input

import (
	"flag"
	"fmt"
)

var (
	Platform = flag.String("p", "twitter", "The platform you want to post a image")
	Usecase  = flag.String("u", "post", "The usecase in your choosing platform")
)

func CliFlagParse() {
	flag.Parse()

	fmt.Println(*Usecase)
	fmt.Println(*Platform)
}
