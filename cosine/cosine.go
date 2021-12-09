package cosine

import (
	"errors"
	"math"
)

func CalculateCosine(arr1 []int, arr2 []int) (float64, error) {

	// element wise do product
	var prod_slice []int

	// var norms
	norm_prod_arr1 := 0
	norm_prod_arr2 := 0
	var norm_arr1 []int
	var norm_arr2 []int

	//check if both the arrays are same
	if len(arr1) != len(arr2) {
		return 0, errors.New("for calculating cosine similarity, both the arrays need to be of same length")
	}

	// calculate dot products and norms
	for i := range arr1 {
		a1 := arr1[i]
		a2 := arr2[i]

		prod := a1 * a2
		prod_slice = append(prod_slice, prod)

		// norms
		norm_prod_arr1 = a1 * a1
		norm_prod_arr2 = a2 * a2
		norm_arr1 = append(norm_arr1, norm_prod_arr1)
		norm_arr2 = append(norm_arr2, norm_prod_arr2)

	}

	dot_product := 0
	norm_a := 0
	norm_b := 0

	for j := range prod_slice {
		dot_product += prod_slice[j]

		norm_a += norm_arr1[j]
		norm_b += norm_arr2[j]
	}

	pow_norm_a := math.Pow(float64(norm_a), 0.5)
	pow_norm_b := math.Pow(float64(norm_b), 0.5)
	norm_mul := pow_norm_a * pow_norm_b
	cosine := float64(dot_product) / norm_mul
	return cosine, nil
}
