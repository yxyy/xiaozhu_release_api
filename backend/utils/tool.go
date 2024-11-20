package utils

import (
	"crypto/md5"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"time"
)

func Uuid() string {

	return time.Now().Format("20060102150405") + strconv.Itoa(Range())

}

func Range() int {

	rand.Seed(time.Now().UnixNano())

	return rand.Intn(8999) + 1000
}

func Salt() string {

	return Random(10)

}

func Random(length int) string {

	if length <= 0 {
		length = 8
	}

	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	arr := make([]byte, length)
	for i := 0; i < length; i++ {
		arr[i] = str[rand.Intn(len(str))]
	}

	return string(arr)
}

func Md5(b string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(b)))
}

func Md5SaltAndPassword(salt, password string) string {
	if salt == "" || password == "" {
		return ""
	}
	return Md5(salt + password)
}

func RandString(n int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890*/.&%$#@!~"
	rand.Seed(time.Now().UnixNano())

	var b []byte
	for i := 0; i < n; i++ {
		b = append(b, str[rand.Intn(len(str))])
	}

	return string(b)
}

func ParseUrl(urls string) error {

	uri, err := url.ParseRequestURI(urls)
	if err != nil {
		return err
	}

	if uri.Scheme != "http" && uri.Scheme != "https" {
		return errors.New("协议错误")
	}

	if uri.Host == "" {
		return errors.New("域名错误")
	}

	return nil
}
