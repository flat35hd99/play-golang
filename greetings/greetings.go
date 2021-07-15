package greetings

import "fmt"

import "errors"

// 名前付きで挨拶を返す
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	msg := fmt.Sprintf("Hi, %v. Welcome!", name)
	return msg, nil
}