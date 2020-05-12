package main

import (
	"fmt"
	"github.com/edreg/awesome/algorithms/module01"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	//Loop:
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		gcd := module01.GCD(a, b)
		fmt.Println(gcd)
		//if 3 == 4+2 {
		//	break Loop
		//}

	}
}
