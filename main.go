package main

import (
	"fmt"

	"github.com/bradMDev/stoics"
)

func main() {
	ism := stoics.SayAurelius()
	ism2 := stoics.SaySeneca()

	fmt.Println(ism)
	fmt.Println(ism2)
}
