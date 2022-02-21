package main

import(
	"fmt"
	"github.com/khurlbut/trader/drive"
	"github.com/khurlbut/trader/prototype"
)

func main() {
	d := drive.NewDrive()
	fmt.Println(prototype.PricingLoop(d))
}
