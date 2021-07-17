package check_time

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
)

// startWrite はアルバイト開始時用の文章を生成する。
func StartWrite(fp *os.File, checkTime string, m map[string]string) {
	startTime := checkTime
	fp.WriteString(startTime)
	fp.Close()

	m["{start_time}"] = checkTime

	ft, err := os.OpenFile("./templates/template_start.txt", os.O_RDWR|syscall.O_RDWR, 0777)
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
