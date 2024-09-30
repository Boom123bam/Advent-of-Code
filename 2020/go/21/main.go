package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type ingredient string
type allergen string
type food struct {
	ingredients []ingredient
	allergens   []allergen
}

func main() {
	file, err := os.Open("21/input.txt")
	// file, err := os.Open("21/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	foods := []food{}
	allIngredients := make(map[ingredient]bool)
	allAllergens := make(map[allergen]bool)

	for scanner.Scan() {
		stringIngredients := strings.Split(strings.Split(scanner.Text(), " (")[0], " ")
		allergensPart := strings.Split(scanner.Text(), " (contains ")[1]
		stringAllergens := strings.Split(allergensPart[:len(allergensPart)-1], ", ")

		food := food{}
		for _, s := range stringIngredients {
			food.ingredients = append(food.ingredients, ingredient(s))
			allIngredients[ingredient(s)] = true
		}
		for _, s := range stringAllergens {
			food.allergens = append(food.allergens, allergen(s))
			allAllergens[allergen(s)] = true
		}

		foods = append(foods, food)
	}

	ingredientsToPossibleAllergens := make(map[ingredient][]allergen)
	for ingredient := range allIngredients {
		tmp := []allergen{}
		for allergen := range allAllergens {
			tmp = append(tmp, allergen)
		}
		ingredientsToPossibleAllergens[ingredient] = tmp
	}

	for _, food := range foods {
		for ingredient := range allIngredients {
			if !ingredient.in(food.ingredients) {
				ingredientsToPossibleAllergens[ingredient] = subtract(ingredientsToPossibleAllergens[ingredient], food.allergens)
			}
		}
	}

	ingredientsWithNoAllergens := []ingredient{}
	for ingredient, allergens := range ingredientsToPossibleAllergens {
		if len(allergens) == 0 {
			ingredientsWithNoAllergens = append(ingredientsWithNoAllergens, ingredient)
		}
	}

	count := 0
	for _, food := range foods {
		for _, targetIng := range ingredientsWithNoAllergens {
			if targetIng.in(food.ingredients) {
				count++
			}
		}
	}
	fmt.Println(count)

	allAllergensArr := []allergen{}
	for allergen := range allAllergens {
		allAllergensArr = append(allAllergensArr, allergen)
	}
	unsolvedAllergens := make([]allergen, len(allAllergensArr))
	copy(unsolvedAllergens, allAllergensArr)

	for len(unsolvedAllergens) > 0 {
		targetIng := findIngWithSingleAllergen(unsolvedAllergens, ingredientsToPossibleAllergens)
		if targetIng == nil {
			panic("what")
		}
		targetAllergen := ingredientsToPossibleAllergens[*targetIng][0]
		eliminateAllergen(targetAllergen, ingredientsToPossibleAllergens)
		unsolvedAllergens = subtract(unsolvedAllergens, []allergen{targetAllergen})
	}

	type kv struct {
		Key   ingredient
		Value allergen
	}

	var ss []kv
	for k, v := range ingredientsToPossibleAllergens {
		if len(v) == 1 {
			ss = append(ss, kv{k, v[0]})
		}
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	result := ""
	for i, kv := range ss {
		if i != 0 {
			result += ","
		}
		result += string(kv.Key)
	}
	fmt.Println(result)
}

func findIngWithSingleAllergen(targets []allergen, ingredientToAllergens map[ingredient][]allergen) *ingredient {
	for ing, alls := range ingredientToAllergens {
		if len(alls) == 1 && alls[0].in(targets) {
			return &ing
		}
	}
	return nil
}

func eliminateAllergen(target allergen, ingredientToAllergens map[ingredient][]allergen) {
	for ing, alls := range ingredientToAllergens {
		if len(alls) == 1 {
			continue
		}
		ingredientToAllergens[ing] = subtract(alls, []allergen{target})
	}
}

func (target ingredient) in(ingredients []ingredient) bool {
	for _, item := range ingredients {
		if item == target {
			return true
		}
	}
	return false
}

func (target allergen) in(allergens []allergen) bool {
	for _, item := range allergens {
		if item == target {
			return true
		}
	}
	return false
}

func common(a, b []ingredient) []ingredient {
	result := []ingredient{}
	for _, i := range a {
		if i.in(b) {
			result = append(result, i)
		}
	}
	return result
}

func subtract(a, b []allergen) []allergen {
	for i := 0; i < len(a); {
		if a[i].in(b) {
			a = append(a[:i], a[i+1:]...)
		} else {
			i++
		}
	}
	return a
}

// func subtract(a *[]allergen, b []allergen) []allergen {
// 	for i := 0; i < len(*a); {
// 		if (*a)[i].in(b) {
// 			(*a) = append((*a)[:i], (*a)[i+1:]...)
// 		} else {
// 			i++
// 		}
// 	}
// 	return a
// }
