package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	words := strings.Fields(input)
	for i := len(words) - 1; i >= 1; i-- {
		fmt.Printf("%s ", words[i])
	}
	fmt.Printf(words[0])
}
