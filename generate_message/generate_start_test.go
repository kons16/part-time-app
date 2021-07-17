package generate_message

import (
	"bufio"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
	"syscall"
	"testing"
)

func Test_GenerateStart(t *testing.T) {
	dirPath, err := os.Getwd()
	if err != nil {
		log.Fatal("Error loading path")
		return
	}
	exePath := getBeforePath(dirPath)
	err = godotenv.Load(exePath + "/.env.test")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	fp, err := os.OpenFile(exePath+"/save_time/start_time_test.txt", os.O_RDWR|syscall.O_RDWR, 0777)
	if err != nil {
		t.Fatalf("Error %s", err)
		return
	}
	defer fp.Close()

	m := map[string]string{
		"{name}": os.Getenv("NAME"),
	}

	templateName := "template_start_test.txt"
	checkTime := "10:00"
	out := extractStdout(t, fp, checkTime, m, exePath, templateName)
	want := "氏名: 田中太郎"
	outRemoveZero := removeZeroBytes(out)

	if outRemoveZero != want {
		t.Fatalf("Error %s", err)
	}
}

// getBeforePath は受け取った dirPath の一つ前までの絶対パスを返す
func getBeforePath(dirPath string) string {
	exePathSlice := strings.Split(dirPath, "/")
	exePath := strings.Join(exePathSlice[:len(exePathSlice)-1], "/")
	return exePath
}

// extractStdout は Stdout に書き込まれた文字列を抽出する
func extractStdout(t *testing.T, fp *os.File, checkTime string, m map[string]string, exePath string, templateName string) string {
	t.Helper()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Error %s", err)
	}

	stdout := os.Stdout
	os.Stdout = w
	GenerateStart(fp, checkTime, m, exePath, templateName)
	w.Close()
	os.Stdout = stdout

	fr := bufio.NewScanner(r)

	outStr := ""
	for fr.Scan() {
		outStr += fr.Text()
	}
	r.Close()

	return outStr
}

// removeZeroBytes は out を[]byteにしたとき、0で padding されているものを取り除く。
func removeZeroBytes(out string) string {
	var aSlice []string
	for _, k := range out {
		if k != 0 {
			aSlice = append(aSlice, string(k))
		}
	}
	a := strings.Join(aSlice, "")
	return a
}
