package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	set, _ := reader.ReadString('\n')

	set = strings.TrimSpace(set)
	setSlice := strings.Split(set, " ")
	var num float64
	var err error
	var floatSlice []float64

	for _, s := range setSlice {

		num, err = strconv.ParseFloat(s, 64)
		if err != nil {
			panic(err)
		}
		floatSlice = append(floatSlice, num)
	}

	var numb float64
	var index int
	for i := 0; i < len(floatSlice)-1; i++ {
		index = i
		numb = floatSlice[i]
		for j := i; j < len(floatSlice); j++ {

			if floatSlice[j] < numb {
				numb = floatSlice[j]
				index = j
			}
		}
		if index != i {
			floatSlice[index] = floatSlice[i]
			floatSlice[i] = numb
		}
	}
	fmt.Println(floatSlice)

}
