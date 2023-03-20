package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	number := random_number()

	fmt.Println("Please enter 4 digits number.")

	for {
		user := readGuess()
		konum, sayi := checkGuess(user, number)

		if konum == 4 {
			fmt.Println("You won!")
			break
		}

		fmt.Printf("Not correct exactly. %d truth digit, %d truth number.\n", konum, sayi)
	}
}

func random_number() string {
	rand.Seed(time.Now().UnixNano())
	number := ""
	for i := 0; i < 4; i++ {
		number += strconv.Itoa(rand.Intn(10))
	}
	return number
}

func readGuess() string {
	var user string
	for {
		fmt.Scanln(&user)
		if len(user) != 4 {
			fmt.Println("Not valid.Please enter 4 digits.")
			continue
		}
		break
	}
	return user
}

func checkGuess(user, number string) (konum, sayi int) {
	for i := 0; i < 4; i++ {
		if user[i] == number[i] {
			konum++
		} else if strings.Contains(number, string(user[i])) {
			sayi++
		}
	}
	return konum, sayi
}
