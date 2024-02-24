package main

import (
	"fmt"

	stoicisms "github.com/bradMDev/stoics"
)

func main() {
	ism := stoicisms.SayAurelius()
	ism2 := stoicisms.SaySeneca()

	fmt.Println(ism)
	fmt.Println(ism2)
}
