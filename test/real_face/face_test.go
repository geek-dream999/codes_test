package real_face

import (
	"fmt"
	"testing"
)

func TestFace(t *testing.T) {
	// 替换为你的图片 URL 和模型路径
	imageURL := "https://image.wanniangqianxian.com/zjian/min_app/2024-11-19/tmp_80900b3bc1c3a006c6698609b5f280905e4a901617a25626.jpg"
	cascadePath := "haarcascade_frontalface_default.xml"

	// 从 URL 加载图片
	image, err := loadImageFromURL(imageURL)
	if err != nil {
		fmt.Printf("图片加载失败: %v\n", err)
		return
	}
	defer image.Close()

	// 检测图片中的人脸
	hasFace, err := detectFaces(image, cascadePath)
	if err != nil {
		fmt.Printf("人脸检测失败: %v\n", err)
		return
	}

	// 输出检测结果
	if hasFace {
		fmt.Println("图片中包含人脸。")
	} else {
		fmt.Println("图片中不包含人脸。")
	}
}
