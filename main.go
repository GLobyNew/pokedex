package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

func main() {
	sc := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for {
		fmt.Print("Pokedex > ")
		sc.Scan()
		curCom := strings.Fields(strings.ToLower(sc.Text()))
		fmt.Printf("Your command was: %v\n", curCom[0])
	}
}
