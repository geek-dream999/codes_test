package email

import "testing"

func TestSend(t *testing.T) {
	from := "l397608301@gmail.com"
	to := []string{"397608301@qq.com"}

	sendEmail(from, to)
}
