package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var collection []float64
	var command []string

	Reader := bufio.NewReader(os.Stdin)
	commadnRead, err := Reader.ReadString('\n')

	if err != nil {

		panic(err)
	}

	commadnRead = strings.TrimSpace(commadnRead)
	command = strings.Split(commadnRead, " ")

	var val float64
	var numf float64
	var errf error

	for command[0] != "exit" {

		if len(command) >= 3 {

			numf, errf = strconv.ParseFloat(command[2], 64)

			if errf != nil {
				panic(errf)
			}
		}
		if len(command) > 1 {
			val, errf = strconv.ParseFloat(command[1], 64)
			if errf != nil {
				panic(err)
			}
		}
		switch command[0] {

		case "show":

			fmt.Println(collection)

		case "add":
			collection = append(collection, val)

		case "inc":

			collection[int(val)] += 1

		case "sub":

			collection[int(val)] -= numf

		case "mul":

			collection[int(val)] *= numf

		case "div":

			collection[int(val)] /= numf

		}

		commadnRead, err = Reader.ReadString('\n')
		if err != nil {

			panic(err)
		}
		commadnRead = strings.TrimSpace(commadnRead)
		command = strings.Split(commadnRead, " ")

	}

}
