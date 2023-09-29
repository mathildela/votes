package main

import (
	"fmt"
	"ia04/comsoc"
)

func main() {
	prefs := [][]comsoc.Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}
	/*
		count, err := comsoc.MajoritySWF(prefs)
		fmt.Println(count)
		fmt.Println(err)

		alt, err := comsoc.MajoritySCF(prefs)
		fmt.Println(alt)
		fmt.Println(err)*/

	count, err := comsoc.BordaSWF(prefs)
	fmt.Println(count)
	fmt.Println(err)

	alt, err := comsoc.BordaSCF(prefs)
	fmt.Println(alt)
	fmt.Println(err)
}
