package disney

import "fmt"

func FindPort(c []Cruise) {
	for i, _ := range c {
		fmt.Println(i)
	}
}
