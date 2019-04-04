package fizzbuzz

import (
	"fmt"
	"strconv"
)

// FizzBuzz はフィズバズです。code buildのテストのために使用します。
// ユニットテストのため、あえて気持ち悪いコードにします。
func FizzBuzz() {
	for i := 1; i < 101; i++ {
		fmt.Println(getFizzBuzz(i))
	}
}

func getFizzBuzz(num int) string {
	switch {
	case num%15 == 0:
		return ("FIZZ BUZZ!")
	case num%3 == 0:
		return ("FIZZ!")
	case num%5 == 0:
		return ("BUZZ!")
	default:
		return strconv.FormatInt(int64(num), 10)
	}
}
