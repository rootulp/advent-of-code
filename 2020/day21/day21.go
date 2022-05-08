package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Printf("Starting day 21...\n")

	partOne := PartOne("input.txt")
	fmt.Printf("Part One: %v\n", partOne)

	partTwo := PartTwo("input.txt")
	fmt.Printf("Part Two: %v\n", partTwo)
}

type Set struct {
	set map[string]bool
}

func NewSet(list []string) Set {
	set := map[string]bool{}
	for _, item := range list {
		set[item] = true
	}

	return Set{set: set}
}

func Intersection(a Set, b Set) Set {
	set := map[string]bool{}

	for aItem := range a.set {
		for bItem := range b.set {
			if aItem == bItem {
				set[aItem] = true
			}
		}
	}
	return Set{set: set}
}

func GetOnlyKey(s Set) string {
	if len(s.set) != 1 {
		log.Fatalf("can not get only key for set with len %v", len(s.set))
	}
	keys := []string{}
	for k := range s.set {
		keys = append(keys, k)
	}
	return keys[0]
}

// PartOne returns the number of times a "safe ingredient" appears in the list of recipes in filename
// A "safe ingredient" is one that can not possibly contain any of the allergens
func PartOne(filename string) int {
	lines := readLines(filename)
	allIngredients, allAllergens := parseLines(lines)
	allergenIngredients := getAllergenIngredients(allAllergens)
	safeIngredients := getSafeIngredients(allIngredients, NewSet(allergenIngredients))

	// fmt.Printf("allIngredients %v\n", allIngredients)
	// fmt.Printf("allAllergens %v\n", allAllergens)
	// fmt.Printf("allergenFoods %v\n", allergenIngredients)
	// fmt.Printf("safeIngredients %v\n", safeIngredients)
	return len(safeIngredients)
}

// PartTwo returns a canonical dangerous list. The canonical dangerous list is a
// list of ingredients that contain an allergen. The list is sorted
// alphabetically and contains no spaces.
func PartTwo(filename string) (canonicalDangerous string) {
	lines := readLines(filename)
	_, allAllergens := parseLines(lines)

	dangerousList := getDangerousList(allAllergens)
	return strings.Join(dangerousList, ",")
}

func getDangerousList(allAllergens map[string]Set) (dangerous []string) {
	allergenToIngredient := map[string]string{}

	for len(allAllergens) > 0 {
		for k, v := range allAllergens {
			if len(v.set) == 1 {
				ingredient := GetOnlyKey(v)
				allergenToIngredient[k] = ingredient
				delete(allAllergens, k)
				for _, setToRemoveIngredient := range allAllergens {
					delete(setToRemoveIngredient.set, ingredient)
				}
			}
		}
	}

	allergens := []string{}
	for allergen := range allergenToIngredient {
		allergens = append(allergens, allergen)
	}
	sort.Strings(allergens)

	for _, allergen := range allergens {
		dangerous = append(dangerous, allergenToIngredient[allergen])
	}

	// fmt.Printf("allergenToIngredient %v\n", allergenToIngredient)
	// fmt.Printf("allergens %v\n", allergens)
	// fmt.Printf("dangerous %v\n", dangerous)
	return dangerous
}

func getSafeIngredients(allIngredients []string, allergenIngredients Set) (safeIngredients []string) {
	for _, ingredient := range allIngredients {
		if !allergenIngredients.set[ingredient] {
			safeIngredients = append(safeIngredients, ingredient)
		}
	}
	return safeIngredients
}

func getAllergenIngredients(allAllergens map[string]Set) (allergenIngredients []string) {
	for _, ingredients := range allAllergens {
		for ingredient := range ingredients.set {
			allergenIngredients = append(allergenIngredients, ingredient)
		}
	}
	return allergenIngredients
}

func parseLines(lines []string) (allIngredients []string, allAllergens map[string]Set) {
	allAllergens = map[string]Set{}

	for _, line := range lines {
		ingredients, allergens := parseLine(line)
		allIngredients = append(allIngredients, ingredients...)
		for _, allergen := range allergens {
			allergenSet := NewSet(ingredients)
			if len(allAllergens[allergen].set) == 0 {
				allAllergens[allergen] = allergenSet
			} else {
				allAllergens[allergen] = Intersection(allAllergens[allergen], allergenSet)
			}
		}
	}
	return allIngredients, allAllergens
}

func parseLine(line string) (ingredients []string, allergens []string) {
	parts := strings.Split(line, " (contains ")
	if len(parts) != 2 {
		log.Fatalf("unexpected number of parts %v for line %v", len(parts), line)
	}
	partOne, partTwo := parts[0], parts[1]
	ingredients = strings.Split(partOne, " ")
	allergens = strings.Split(strings.Trim(partTwo, ")"), ", ")
	return ingredients, allergens
}

func readLines(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
