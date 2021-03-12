package main

import "fmt"

var (
	coins = 5000
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distrbution = make(map[string]int, len(users))
)

func dispatchCoin() (left int) {
	for _, name := range users {
		for _, c := range name {
			switch c {
			case 'e', 'E':
				distrbution[name]++
				coins--
			case 'i', 'I':
				distrbution[name] += 2
				coins -= 2
			case 'o', 'O':
				distrbution[name] += 3
				coins -= 3
			case 'u', 'U':
				distrbution[name] += 4
				coins -= 4
			}
		}
	}
	left = coins
	return
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下： ", left)
	for k, v := range distrbution {
		fmt.Println(k, ":", v)
	}
}
