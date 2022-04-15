package tool

import (
	"math/rand"
	"time"
)

func RandString(num int) string {
	str := "qwertyuiopasdfghjklzxcvbnm"
	newStr := ""
	rand.Seed(time.Now().Unix())
	var i = 0
	for i < num {
		index := rand.Intn(25)
		newStr = newStr + str[index:index+1]
		i = i + 1
	}
	return newStr
}
