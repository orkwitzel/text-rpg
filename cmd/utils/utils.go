package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ClearScreen() {
	if os.Getenv("OS") == "windows" {
		fmt.Print("\033[H\033[2J")
	} else {
		fmt.Print("\033[H\033[2J")
	}
}

func InputString() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}
	return strings.TrimSpace(input)
}

func InputInt() int {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}
	n, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing input: %v\n", err)
		os.Exit(1)
	}
	return n
}

func InputFloat() float64 {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}
	f, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing input: %v\n", err)
		os.Exit(1)
	}
	return f
}

func InputBool() bool {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}
	b, err := strconv.ParseBool(strings.TrimSpace(input))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing input: %v\n", err)
		os.Exit(1)
	}
	return b
}

var linkingWords = []string{"the", "a", "an", "of", "in", "on", "to", "from", "with", "as"}

func RemoveLinkingWordsFromArgs(args []string) []string {
	newArgs := []string{}
	for i, arg := range linkingWords {
		formmatedArg := strings.TrimSpace(strings.ToLower(args[i]))
		if formmatedArg != arg {
			newArgs = append(newArgs, arg)
		}
	}
	return newArgs
}

func SectionTitlePrint(title string) {
	fmt.Println("--------------------------------")
	fmt.Println(title)
	fmt.Println("--------------------------------")
}
