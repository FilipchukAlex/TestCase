// Without timeout 10 sek
package main

import ("fmt"
"encoding/json"
"strconv"
)

type User struct {
	Number int; SerialNumber int
}

func main() {
	index_number := 0
	correct_answer := 0
	incorrect_answer := 0
	var a, b= 0, 1
	for correct_answer <= 9 && incorrect_answer <= 2 {
		fmt.Println("Enter a Fibonacci number: ")
		var input string
		fmt.Scanf("%s", &input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Not an integer!")
			continue
		}
		if num == a {
			a, b = b, a + b
			index_number ++
			correct_answer ++
		} else {
			output := &User{Number: a, SerialNumber: index_number}
			c, _ := json.Marshal(output)
			fmt.Println(string(c))
			a, b = b, a + b
			index_number ++
			incorrect_answer ++
			correct_answer = 0
		}
		if correct_answer > 9 {
			fmt.Print(" You win!")
		}else if incorrect_answer > 2{
			fmt.Print("You lose!")
		}
	}

}
