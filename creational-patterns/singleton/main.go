package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 30; i++ {
		go func(index int){
			logger := getInstance()
			logger.log(strconv.Itoa(i))
		}(i)
	}
	fmt.Scanln()
}
