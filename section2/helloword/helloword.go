package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello word now")

	nb, err := fmt.Println("Hello word 2")
	_, _ = fmt.Println(nb, err)
}
