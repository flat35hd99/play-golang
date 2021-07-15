package greetings

import "fmt"

// 名前付きで挨拶を返す
func Hello(name string) string {
	msg := fmt.Sprintf("Hi, %v. Welcome!", name)
	return msg
}