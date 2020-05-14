package convert

import (
	"image"
	"image-combinator/internal/input"
	"image/draw"
)

func Combine(imgs input.Images, screenWidth, screenHeight, paddingX, paddingY, densityCol, densityRow int) *image.RGBA {
	// 背景画像の作成
	screen := image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight))

	// 背景画像への書き込み
	posX := paddingX
	posY := paddingY
	for i, img := range imgs {
		rect := image.Rect(posX, posY, posX+img.Width, posY+img.Height)
		draw.Draw(screen, rect, img.Src, image.Point{0, 0}, draw.Over)

		// １行分の連結を続ける場合
		if densityCol-(i%densityCol) > 1 {
			posX += paddingX + img.Width
			continue
		}

		// １行分の連結を終える場合
		// ２列目以降も連結を行う場合
		if densityRow > 2 {
			posX = paddingX
			posY += paddingY + img.Height
			continue
		}

		// １列目・２列目で連結が完了する場合
		posX = paddingX
		posY += 2*paddingY + img.Height
	}

	return screen
}
