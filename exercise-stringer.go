// https://tour.golang.org/methods/18
package main

import  ( 
	"fmt"
	"strings"
	"strconv"
)

type IPAddr [4]byte

func (p IPAddr) String() string {	
	var s []string
	for _, x := range p {
		s = append(s, strconv.Itoa(int(x)))
	}
	return fmt.Sprintf("%v", strings.Join(s, "."))
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

