package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func TidyDirectory(filepath string) (err error) {
	_, err = os.Stat(filepath)
	if !os.IsExist(err) {
		err = os.MkdirAll(filepath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func NormalizePath(path string) string {
	return filepath.ToSlash(path)
}

func GetFileExt(filepath string) string {
	index := strings.LastIndex(filepath, ".")
	if index == -1 {
		return ""
	}

	return filepath[index:]
}

func GetFileMd5(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hash := md5.New()
	if _, err = io.Copy(hash, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
