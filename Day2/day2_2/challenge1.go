package RunChallenge2

import(
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	command := make([][]string, 0)
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
	return command
}

func RunChallenge2(){
	input := GetCommands()
	var horizontal int = 0
	var depth int = 0
	var aim int = 0
	for i:=0; i < len(input); i++{
		steps, err := strconv.Atoi(input[i][1])
		if err != nil {
			log.Fatal(err)
		}
		switch(input[i][0]){
		case "forward":
			horizontal += steps
			depth += aim*steps
		case "up":
			aim -= steps
		case "down":
			aim += steps
		}
	}
	//fmt.Println(horizontal)
	//fmt.Println(depth)
	fmt.Printf("\nChallenge 2 answer: %d", horizontal*depth)
}
