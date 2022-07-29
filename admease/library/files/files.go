package files

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

//执行的目录
func GetCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func MkdirIfNotExist(path string) error {
	exist, err := PathExists(path)
	if err != nil {
		return err
	}

	if !exist {
		err := os.Mkdir(path, os.ModePerm)
		return err
	}
	return nil
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func ParentDir(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func ProjectDir() string {
	dirctory, _ := os.Getwd()
	projectName := "admease"
	arr := strings.Split(dirctory, projectName)
	if len(arr) > 0 {
		return arr[0] + projectName
	}
	return dirctory
}
