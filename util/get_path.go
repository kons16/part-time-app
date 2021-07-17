package util

import (
	"log"
	"os"
	"strings"
)

func GetPath() string {
	exeFullPath, err := os.Executable()
	if err != nil {
		log.Fatal("Error exec loading path.")
		return ""
	}
	exePathSlice := strings.Split(exeFullPath, "/")
	exePath := strings.Join(exePathSlice[:len(exePathSlice)-1], "/")
	return exePath
}
