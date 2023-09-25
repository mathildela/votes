package main

import (
	"fmt"
	"ia04/comsoc"
)

func main() {
	var alt comsoc.Alternative = 3
	tab := [...]comsoc.Alternative{1, 2, 3, 4, 5}
	prefs := tab[:]

	fmt.Println(alt, prefs)
}
