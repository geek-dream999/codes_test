package qr_code

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"testing"
)

func TestQrCode1(t *testing.T) {
	// 加载背景图片和二维码图片
	backgroundPath := "./img/background1.png"
	qrCodePath := "./img/qrcode.png"

	backgroundImg, err := imaging.Open(backgroundPath)
	if err != nil {
		log.Fatalf("无法打开背景图片: %v", err)
	}

	qrCodeImg, err := imaging.Open(qrCodePath)
	if err != nil {
		log.Fatalf("无法打开二维码图片: %v", err)
	}

	// 调整二维码大小
	qrCodeImg = imaging.Resize(qrCodeImg, 800, 800, imaging.Lanczos)

	// 设置透明度
	alpha := 0.3 // 透明度30%
	qrCodeImg = adjustAlpha(qrCodeImg, alpha)

	// 调整亮度和对比度
	qrCodeImg = imaging.AdjustBrightness(qrCodeImg, -20) // 降低亮度
	qrCodeImg = imaging.AdjustContrast(qrCodeImg, -30)   // 降低对比度

	// 将二维码叠加到背景图片指定位置
	offset := image.Pt(150, 300) // 放置二维码的位置
	result := imaging.Overlay(backgroundImg, qrCodeImg, offset, alpha)

	// 保存合成图片
	outputPath := "./img/final_output.png"
	err = imaging.Save(result, outputPath)
	if err != nil {
		log.Fatalf("无法保存最终图片: %v", err)
	}

	fmt.Printf("最终合成效果图已保存到: %s\n", outputPath)
}

// adjustAlpha 用于调整图像的透明度
func adjustAlpha(img image.Image, alpha float64) image.Image {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			r, g, b, a := originalColor.RGBA()
			a = uint32(float64(a) * alpha) // 调整透明度
			newColor := color.RGBA64{uint16(r), uint16(g), uint16(b), uint16(a)}
			newImg.Set(x, y, newColor)
		}
	}
	return newImg
}

func TestQrCode(t *testing.T) {
	// 生成二维码图像
	qrContent := "https://yahuihui.cn/"
	qrCode, err := generateQRCode(qrContent, 256) // 生成 256x256 的二维码
	if err != nil {
		log.Fatal("生成二维码失败:", err)
	}

	// 加载背景图像（目标图片）
	backgroundFile, err := os.Open("background.png")
	if err != nil {
		log.Fatal("加载背景图片失败:", err)
	}
	defer backgroundFile.Close()

	background, _, err := image.Decode(backgroundFile)
	if err != nil {
		log.Fatal("解码背景图片失败:", err)
	}

	// 调整二维码透明度为 20%（默认是完全不透明）
	transparency := uint8(51) // 范围 0 - 255，255 表示完全不透明，0 表示完全透明

	// 将二维码嵌入到目标图片中，位置可以自己调整
	newImage := embedQRCodeInImage(background, qrCode, image.Point{X: 100, Y: 100}, transparency)

	// 保存合成后的图像
	outputFile, err := os.Create("output.png")
	if err != nil {
		log.Fatal("保存合成图像失败:", err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, newImage)
	if err != nil {
		log.Fatal("编码合成图像失败:", err)
	}

	fmt.Println("合成图像已保存为 output.png")
}

func TestQrCode2(t *testing.T) {
	// 加载背景图片
	backgroundPath := "./img/background1.png" // 替换为背景图片路径
	bgFile, err := os.Open(backgroundPath)
	if err != nil {
		fmt.Println("无法打开背景图片:", err)
		return
	}
	defer bgFile.Close()

	bgImg, _, err := image.Decode(bgFile)
	if err != nil {
		fmt.Println("无法解码背景图片:", err)
		return
	}

	// 加载二维码图片
	qrCodePath := "./img/qrcode.png" // 替换为二维码图片路径
	qrFile, err := os.Open(qrCodePath)
	if err != nil {
		fmt.Println("无法打开二维码图片:", err)
		return
	}
	defer qrFile.Close()

	qrImg, _, err := image.Decode(qrFile)
	if err != nil {
		fmt.Println("无法解码二维码图片:", err)
		return
	}

	// 调整二维码的大小（如有需要，可更改尺寸）
	const qrWidth, qrHeight = 800, 800
	qrResized := resizeImage(qrImg, qrWidth, qrHeight)

	// 创建一个新的图像，并将背景图和二维码合成
	output := image.NewRGBA(bgImg.Bounds())
	draw.Draw(output, bgImg.Bounds(), bgImg, image.Point{0, 0}, draw.Src)

	// 定义二维码的位置 (xOffset, yOffset)
	xOffset, yOffset := 150, 300
	for y := 0; y < qrResized.Bounds().Dy(); y++ {
		for x := 0; x < qrResized.Bounds().Dx(); x++ {
			qrPixel := qrResized.At(x, y)
			r, g, b, a := qrPixel.RGBA()

			// 设置透明度：alpha 值控制透明度，范围0-65535（最大）
			alpha := 0.3
			a = uint32(float64(a) * alpha)

			// 组合二维码与背景图片的像素
			bgPixel := output.At(x+xOffset, y+yOffset)
			br, bg, bb, ba := bgPixel.RGBA()

			// 混合颜色，使用透明度融合
			r = (r*a + br*(65535-a)) / 65535
			g = (g*a + bg*(65535-a)) / 65535
			b = (b*a + bb*(65535-a)) / 65535
			a = (a + ba) / 2

			output.Set(x+xOffset, y+yOffset, color.RGBA64{uint16(r), uint16(g), uint16(b), uint16(a)})
		}
	}

	// 保存结果图像
	outputPath := "merged_output.png"
	outFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("无法创建输出文件:", err)
		return
	}
	defer outFile.Close()

	err = png.Encode(outFile, output)
	if err != nil {
		fmt.Println("无法保存输出图像:", err)
		return
	}

	fmt.Println("合成图像已保存到:", outputPath)
}

// resizeImage 函数，用于调整二维码图片大小
func resizeImage(img image.Image, width, height int) image.Image {
	dc := gg.NewContext(width, height)
	dc.DrawImageAnchored(img, width/2, height/2, 0.5, 0.5)
	return dc.Image()
}
