package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * Ini untuk mendapatkan jumlah inputan lebih dari 1 kata.
 */
var reader = bufio.NewReader(os.Stdin)

func ReadLine(prompt string) string {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Input cannot be empty, please try again.")
			continue
		}

		return input
	}
}

func ReadInt(prompt string) (int, error) {
	input := ReadLine(prompt)
	return strconv.Atoi(input)
}
