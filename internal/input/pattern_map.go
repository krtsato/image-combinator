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
type densityMapType map[int]measureMapType
type aspectMapType map[string]densityMapType

var AspectMap = aspectMapType{
	"3:1": {
		3: {
			"column": 3,
			"row":    1,
		},
		75: {
			"column": 15,
			"row":    5,
		},
	},
	"16:9": {
		10: {
			"column": 5,
			"row":    2,
		},
		144: {
			"column": 16,
			"row":    9,
		},
	},
}
