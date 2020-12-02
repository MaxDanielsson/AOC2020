package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	partOne("input.txt")
	partTwo("input.txt")

}

func partOne(s string) {
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	var valid = 0

	for sc.Scan() {
		//fmt.Printf("we got: %s \n", sc.Text())
		var splits = strings.Split(sc.Text(), " ")

		var span = strings.Split(splits[0], "-")
		var let = strings.Split(splits[1], ":")
		var psswrd = splits[2]
		var count = 0

		for i := 0; i < len(psswrd); i++ {
			if string(psswrd[i]) == let[0] {
				count++
			}
		}
		min, _ := strconv.Atoi(span[0])
		max, _ := strconv.Atoi(span[1])
		if count >= min && count <= max {
			valid++
		}
	}
	fmt.Printf("valid psswrds one: %d\n", valid)
}
func partTwo(s string) {
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	var valid = 0

	for sc.Scan() {
		var splits = strings.Split(sc.Text(), " ")
		var span = strings.Split(splits[0], "-")
		var let = strings.Split(splits[1], ":")
		var psswrd = splits[2]
		var a = ""
		var b = ""

		min, _ := strconv.Atoi(span[0])
		max, _ := strconv.Atoi(span[1])

		if len(psswrd) >= min {
			a = string(psswrd[min-1])
		}
		if len(psswrd) >= max {
			b = string(psswrd[max-1])
		}

		//(X || Y) && !(X && Y)
		if (a == let[0] || b == let[0]) && !(a == let[0] && b == let[0]) {
			valid++
		}

	}
	fmt.Printf("valid psswrds two: %d\n", valid)
}
