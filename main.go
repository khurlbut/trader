package main

import(
"fmt"
"github.com/khurlbut/trader/hello"
"github.com/khurlbut/trader/prototype"
)

func main() {
	fmt.Println(hello.BuildHello())
	fmt.Println(hello.BuildHi())
	fmt.Println("hey", prototype.PricingLoop())
}
