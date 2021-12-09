package main

import (
	"fmt"
	"log"

	"cosine.com/cosine"
)

func main() {

	log.SetPrefix("Cosine: ")
	log.SetFlags(0)
	//define two arrays
	vec_a := []int{1, 2, 3, 4, 5}
	vec_b := []int{1, 3, 5, 7, 9}

	cosine, err := cosine.CalculateCosine(vec_a, vec_b)
	if err != nil {
		log.Fatal("Error occured")
	}
	fmt.Println("The cosine similarity between two vectors is", cosine)
}
