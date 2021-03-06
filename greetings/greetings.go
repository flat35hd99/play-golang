package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// 名前付きで挨拶を返す
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	msg := fmt.Sprintf(randomFormat(), name)
	return msg, nil
}

func Hellos(names []string) (map[string]string, error) {
	msgs := make(map[string]string)

	for _, name := {
		msg, err := range names {
			msg, err := Hello(name)
			if err != nil {
				return nil, err
			}
			msgs[name] = msg
		}
	}
}

// プログラムの初期化処理の一環で呼ばれる。Global varのinit後に実行される。
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// `slice`というらしい。
	// []の中になにも入れないとdynamicallyだそう。
	formats := []string{
		"Hi, %v. Welcome",
		"Great to see you %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}