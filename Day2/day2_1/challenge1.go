package RunChallenge1

import(
	"fmt"
	"log"
	"strconv"
)

func RunChallenge1(input [][]string){
	var horizontal int = 0
	var depth int = 0
	for i:=0; i < len(input); i++{
		steps, err := strconv.Atoi(input[i][1])
		if err != nil {
			log.Fatal(err)
		}
		switch(input[i][0]){
		case "forward":
			horizontal += steps
		case "up":
			depth -= steps
		case "down":
			depth += steps
		}
	}
	fmt.Println(horizontal)
	fmt.Println(depth)
	fmt.Println(horizontal*depth)
}
