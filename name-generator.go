package main

import "fmt"
import "math/rand"

func main() {
	first_names := []string{"John", "Sam", "Lily", "Mary", "Roger"}
	last_names := []string{"Smith", "Blackwood", "Berger", "Haley", "Foster"}

	rand1 := rand.Intn(len(first_names))
	rand2 := rand.Intn(len(last_names))

	fmt.Println(first_names[rand1], last_names[rand2])
}