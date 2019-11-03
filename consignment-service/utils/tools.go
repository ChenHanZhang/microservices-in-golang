package utils

import (
	"fmt"
	"time"
)

func PrintWithTime(s string)  {
	o := time.Now().String() + ":" + s
	fmt.Println(o)
}
