package input

// CLI オプションの有効値を格納したマップ

type screenMapType map[string]int
type usecaseMapType map[string]screenMapType
type platformMapType map[string]usecaseMapType

var PlatformMap = platformMapType{
	"twitter": {
		"header": {
			"width":  1500,
			"height": 500,
		},
		"post": {
			"width":  1024,
			"height": 576,
		},
	},
	"youtube": {
		"screen": {
			"width":  1920,
			"height": 1080,
		},
		"thumbnail": {
			"width":  1280,
			"height": 720,
		},
	},
}

var densityMap = usecaseMapType{
	"3:1": {
		"max": 75,
		"min": 3,
	},
	"16:1": {
		"max": 144,
		"min": 10,
	},
}
