package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Kawan Lama")
	var s string = "USOMAANAPAIUMASYDNIP"
	var sArr []string
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "S" || string(s[i]) == "I" || string(s[i]) == "A" || string(s[i]) == "P" {
			sArr = append(sArr, string(s[i]))
		}
	}
	var maybeCount int
	newS := strings.Join(sArr, " ")
	countS := strings.Count(newS, "S")
	countI := strings.Count(newS, "I")
	countA := strings.Count(newS, "A")
	countP := strings.Count(newS, "P")
	if countS <= countI || countS <= countP {
		maybeCount = countS
	} else if countI <= countS || countI <= countP {
		maybeCount = countI
	} else if countP <= countS || countP <= countI {
		maybeCount = countP
	}
	if maybeCount <= countA/maybeCount {
		fmt.Println("SIAPA:", maybeCount)
	} else {
		fmt.Println("SIAPA:", countA/maybeCount)
	}
}
