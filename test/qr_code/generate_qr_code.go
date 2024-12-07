package qr_code

import (
	"bytes"
	"fmt"
	"github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"image/draw"
	"log"
)

// GenerateQrCode 生成二维码
func GenerateQrCode() {
	content := "https://yahuihui.cn/" // 二维码内容
	err := qrcode.WriteFile(content, qrcode.Medium, 256, "img/qrcode.png")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("二维码已生成，保存在当前目录的 qrcode.png 文件中")

}

// 生成二维码图像
func generateQRCode(content string, size int) (image.Image, error) {
	qrData, err := qrcode.Encode(content, qrcode.Medium, size)
	if err != nil {
		return nil, err
	}

	// 解码二维码字节为图像
	qrImage, _, err := image.Decode(bytes.NewReader(qrData))
	if err != nil {
		return nil, err
	}

	return qrImage, nil
}

// 将二维码嵌入到背景图像中，并调整透明度和颜色
func embedQRCodeInImage(background image.Image, qrCode image.Image, position image.Point, transparency uint8) image.Image {
	// 创建一个新的图像，用于绘制带二维码的背景
	rgba := image.NewRGBA(background.Bounds())
	draw.Draw(rgba, background.Bounds(), background, image.Point{}, draw.Src)

	// 调整二维码颜色和透明度
	qrWidth, qrHeight := qrCode.Bounds().Dx(), qrCode.Bounds().Dy()
	for x := 0; x < qrWidth; x++ {
		for y := 0; y < qrHeight; y++ {
			qrColor := qrCode.At(x, y)
			r, g, b, a := qrColor.RGBA()

			if a > 0 { // 非透明像素
				r = r * uint32(transparency) / 255
				g = g * uint32(transparency) / 255
				b = b * uint32(transparency) / 255
				rgba.Set(position.X+x, position.Y+y, color.RGBA{
					R: uint8(r >> 8),
					G: uint8(g >> 8),
					B: uint8(b >> 8),
					A: uint8(transparency),
				})
			}
		}
	}
	return rgba
}
