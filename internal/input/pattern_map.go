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

type measureMapType map[string]int
type densityMapType map[string]measureMapType
type aspectMapType map[string]densityMapType

var aspectMap = aspectMapType{
	"3:1": {
		"3": {
			"width":  3,
			"height": 1,
		},
		"75": {
			"width":  15,
			"height": 5,
		},
	},
	"16:9": {
		"10": {
			"width":  5,
			"height": 2,
		},
		"144": {
			"width":  16,
			"height": 9,
		},
	},
}
