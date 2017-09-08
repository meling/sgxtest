package main

import (
	"fmt"

	"github.com/klauspost/cpuid"
)

func main() {
	// Print basic CPU information:
	fmt.Println("Name:", cpuid.CPU.BrandName)

	// Test if we have a specific feature:
	if cpuid.CPU.SGX.Available {
		fmt.Println("Yeah, 🙂, you got SGX support, hopefully!")
	} else {
		fmt.Println("Nope, SGX not to be found here 😞")
	}
}
