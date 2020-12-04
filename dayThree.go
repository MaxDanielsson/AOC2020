package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var multRes = 1
	var listD = [5]int{1, 1, 1, 1, 2}
	var listR = [5]int{1, 3, 5, 7, 1}

	for i := 0; i < len(listR); i++ {
		var val = threeOne("input.txt", listD[i], listR[i])
		multRes *= val
		fmt.Printf("val for i=%d is = %d\n", i, val)
	}
	fmt.Printf("Answer = %d\n", multRes)
	//fmt.Printf("val = %d\n", threeOne("input.txt", 2, 1))
}

func threeOne(s string, down, right int) int {
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	var position int = 0
	var row int = 0
	var amount int = 0

	for sc.Scan() {
		if row%down == 0 {
			var treeRow = sc.Text()

			if string(treeRow[position]) == "#" {
				amount++
			}
			position += right

			if position >= len(treeRow) {
				position -= len(treeRow)
			}
		}
		row++
	}
	//fmt.Printf("The number of encountered trees are: %d\n", amount)
	return amount
}
