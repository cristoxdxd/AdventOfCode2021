package main

import (
	"fmt"

	d1run1 "github.com/cristoxdxd/AdventOfCode2021/Day1/day1_1"
	d1run2 "github.com/cristoxdxd/AdventOfCode2021/Day1/day1_2"
	d2run1 "github.com/cristoxdxd/AdventOfCode2021/Day2/day2_1"
	d2run2 "github.com/cristoxdxd/AdventOfCode2021/Day2/day2_2"
)

func main(){
	// Day 1:
	fmt.Println("\tDay 1: Sonar Sweep")
	d1run1.RunChallenge1();
	d1run2.RunChallenge2();
	// Day 2:
	fmt.Println("\n\n\tDay 2: Dive!")
	d2run1.RunChallenge1()
	d2run2.RunChallenge2()
}
