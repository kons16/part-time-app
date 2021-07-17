package generate_message

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
)

// GenerateEnd はアルバイト終了時用の文章を生成する。
func GenerateEnd(fp *os.File, checkTime string, m map[string]string, exePath string) {
	m["{start_time}"] = readTime(fp)
	m["{end_time}"] = checkTime

	ft, err := os.OpenFile(exePath+"/templates/template_end.txt", os.O_RDWR|syscall.O_RDWR, 0777)
	if err != nil {
		log.Fatal("Error loading template_end.txt file")
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
			log.Fatal("Error loading template_end.txt file")
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

func readTime(fp *os.File) string {
	b := make([]byte, 216)
	var startTime string
	for {
		n, err := fp.Read(b)
		if n == 0 {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		startTime = string(b)
	}
	fp.WriteString("")
	fp.Close()

	return startTime
}
