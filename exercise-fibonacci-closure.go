// https://tour.golang.org/moretypes/26 
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	nm1 := 1
	nm2 := 0
	n := 0
	count := 0
	
	return func() int {
		count++;
		if count == 1 {
			return 0
		}
		if count == 2 {
			return 1
		}
		
		n = nm1 + nm2
		nm2 = nm1
		nm1 = n
		return n
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
