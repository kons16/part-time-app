package main

import (
	"flag"
	"github.com/joho/godotenv"
	"github.com/kons16/part-time-app/check_time"
	"log"
	"os"
	"strconv"
	"syscall"
	"time"
)

func main() {
	flag.Parse()
	fileType := flag.Args()[0]
	checkTime := flag.Args()[1]

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	t := time.Now()
	wdays := []string{"日", "月", "火", "水", "木", "金", "土"}
	dayOfWeek := wdays[t.Weekday()]

	// .envに書かれている初期設定と曜日をmapに登録
	m := map[string]string{
		"{major}":     os.Getenv("MAJOR"),
		"{lab}":       os.Getenv("LAB"),
		"{name}":      os.Getenv("NAME"),
		"{date}":      strconv.Itoa(int(t.Month())) + "月" + strconv.Itoa(int(t.Day())) + "日",
		"{dayOfWeek}": dayOfWeek,
	}

	fp, err := os.OpenFile("start_time.txt", os.O_RDWR|syscall.O_RDWR, 0777)
	if err != nil {
		log.Fatal("Error loading start_time.txt file")
		return
	}
	defer fp.Close()

	if fileType == "s" {
		check_time.StartWrite(fp, checkTime, m)
	} else if fileType == "e" {
		check_time.EndWrite(fp, checkTime, m)
	}
}
