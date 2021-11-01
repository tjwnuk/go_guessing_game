package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Guessing game:")
	time := time.Now().Unix()

	rand.Seed(time)
	target := rand.Intn(100) + 1
	fmt.Println("Secret number is: ", target)

	try_count := 0

	reader := bufio.NewReader(os.Stdin)
	for {
		try_count++
		fmt.Println("Guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("an error occured")
		}
		input = strings.TrimSuffix(input, "\n")
		num, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal("Enter integers only")
		}
		if num == target {
			fmt.Println("Congratulations! You've guessed the number in the try number", try_count)
			break
		}
		if num > target {
			fmt.Println("Too HIGH")
			continue
		}
		if num < target {
			fmt.Println("Too LOW")
			continue
		}

	}
}
