package main

import (
	"fmt"
	"math"
)

func main() {

	for i := 10; i >= 1; i-- {
		fmt.Print(math.Pow10(i - 1))
		fmt.Println(" Louis Loop")
	}

}
