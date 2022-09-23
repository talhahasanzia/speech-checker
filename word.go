package main

import (
	"fmt"
	"strings"
)

func getWords() []Word {

	var words []Word

	data := ReadFile("pos.csv")

	for i, row := range data {

		if i == 0 {
			continue
		}
		word := Word{}
		for j, column := range row {

			if j == 0 {
				word.Name = strings.TrimSpace(column)
			}

			if j == 1 {
				word.P1 = strings.TrimSpace(column)
			}

			if j == 2 {
				word.P2 = strings.TrimSpace(column)
			}

			if j == 3 {
				word.P3 = strings.TrimSpace(column)
			}

		}
		fmt.Println(word.Name, word.P1)
		words = append(words, word)

		if i == 228_472 {
			break
		}

	}

	return words
}
