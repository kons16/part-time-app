package generate_message

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
)

// GenerateStart はアルバイト開始時の文章を生成する。
func GenerateStart(fp *os.File, checkTime string, m map[string]string, exePath string) {
	startTime := checkTime
	fp.WriteString(startTime)
	fp.Close()

	m["{start_time}"] = checkTime

	ft, err := os.OpenFile(exePath+"/templates/template_start.txt", os.O_RDWR|syscall.O_RDWR, 0777)
	if err != nil {
		log.Fatal("Error loading template_start.txt file")
		return
	}
	defer ft.Close()

	buf := make([]byte, 1024)
	for {
		n, err := ft.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			log.Fatal("Error loading template_start.txt file")
			return
		}

		line := string(buf)
		for key, value := range m {
			if strings.Contains(line, key) {
				line = strings.Replace(line, key, value, -1)
			}
		}
		fmt.Println(line)
	}
}
