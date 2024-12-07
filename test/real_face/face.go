package real_face

import (
	"bytes"
	"fmt"
	"gocv.io/x/gocv"
	"io"
	"net/http"
)

func loadImageFromURL(url string) (gocv.Mat, error) {
	// 下载图片数据
	resp, err := http.Get(url)
	if err != nil {
		return gocv.NewMat(), fmt.Errorf("无法下载图片: %v", err)
	}
	defer resp.Body.Close()

	// 将图片数据加载为内存字节
	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return gocv.NewMat(), fmt.Errorf("无法读取图片数据: %v", err)
	}

	// 解码图片数据为 OpenCV 的 Mat
	img, err := gocv.IMDecode(buf.Bytes(), gocv.IMReadColor)
	if err != nil {
		return gocv.NewMat(), fmt.Errorf("无法解码图片: %v", err)
	}
	return img, nil
}

func detectFaces(image gocv.Mat, cascadePath string) (bool, error) {
	// 加载 Haar 分类器
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(cascadePath) {
		return false, fmt.Errorf("无法加载人脸检测模型: %s", cascadePath)
	}

	// 转为灰度图（人脸检测需要灰度图）
	grayImage := gocv.NewMat()
	defer grayImage.Close()
	gocv.CvtColor(image, &grayImage, gocv.ColorBGRToGray)

	// 检测人脸
	faces := classifier.DetectMultiScale(grayImage)
	return len(faces) > 0, nil
}
