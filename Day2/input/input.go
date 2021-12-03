package GetCommands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetCommands()[][]string {
	file, err := os.Open("./input.txt")
	if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
 
    scanner := bufio.NewScanner(file)
	
	var input [1000]string
	var command [1000][]string
	var i int = 0
    for scanner.Scan() {
		input[i] = scanner.Text()
		command[i] = strings.SplitAfter(input[i], " ")
		i++
    }
	fmt.Println(command)
     if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
	return command[:]
}
