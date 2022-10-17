package main

import (
	"fmt"
	"os"
)

//func readInput() (text string, err error) {
//	reader := bufio.NewReader(os.Stdin)
//	text, err = reader.ReadString('\n')
//	return text, err
//}

func findWord(text string) (flag bool, position int) {
	bytes := []byte(text)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] != 32 && bytes[i] != 10 {
			return true, i
		}
	}
	return false, -1
}

func countWords(text string) (count int) {
	count = 0
	flag, position := findWord(text)
	if flag {
		count++
		for i := position; i < len(text)-1; i++ {
			if string(text[i]) == " " && text[i+1] != text[len(text)-1] {
				if string(text[i+1]) != " " {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	text := os.Args[1]
	count := countWords(text)
	fmt.Println(count)
}
