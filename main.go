package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	Black   = "\033[1;30m"
	Red     = "\033[1;31m"
	Green   = "\033[1;32m"
	Yellow  = "\033[1;33m"
	Blue    = "\033[1;34m"
	Magenta = "\033[1;35m"
	Cyan    = "\033[1;36m"
	White   = "\033[1;37m"
	Reset   = "\033[0m"
)

func animateText(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(delay)
	}
	fmt.Println()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(101)
	var guess int
	tries := 1
	maxTries := 5

	animateText(Magenta+"Annyeonghaseyo"+Reset, 100*time.Millisecond)
	time.Sleep(1 * time.Second)

	fmt.Println(Cyan + "I'm thinking of a number between 1 and 100. Try to guess it:" + Reset)

	for tries <= maxTries {
		fmt.Scan(&guess)
		if guess == n {
			fmt.Println(Green + "Congratulations, you got it right!" + Reset)
			fmt.Println("/ \\ / \\")
	        fmt.Println("\\     / see you soon")
	        fmt.Println(" \\   /    .--.")
			fmt.Println("  \\ /")
			break
		} else if guess < n {
			fmt.Println(Red + "Your guess is too low." + Reset)
		} else {
			fmt.Println(Red + "Your guess is too high." + Reset)
		}
		tries++

		if tries == 4 {
			var hint string
			one:
			fmt.Println(Blue + "You've had 3 tries. Do you want a hint? (yes/no)" + Reset)
			fmt.Scan(&hint)
			if hint == "yes" {
				fmt.Println("The number is between:", n-2, "and", n+2)
			} else if hint == "no" {
				fmt.Println("ok")
			} else {
				fmt.Println("What? That's not an answer.")
				goto one
			}
		}
	}

	if tries > maxTries {
		fmt.Println("Sorry, you've used up all your tries. The number was:", n)
		fmt.Println("Do you want a second chance? (sure/no)")
		var answer string
		fmt.Scan(&answer)
		if answer == "sure" {
			main()
		} else if answer == "no" {
			os.Exit(0)
		} else {
			fmt.Println(answer + ", that's not an option.")
			os.Exit(1)
		} 
	}
}
