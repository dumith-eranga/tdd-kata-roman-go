package main

import (
	"fmt"
	"strings"
)

func ArabicToRoman(number int) string {
	symbols := []struct {
		arabic int
		roman  string
	}{
		{1, "I"},
		{5, "V"},
		{10, "X"},
	}

	type transformationDefinition struct {
		from string
		to   string
	}

	transformations := []transformationDefinition{}

	if 0 >= number {
		return ""
	}

	roman := ""

	for i := len(symbols); i > 0; i-- {
		symbol := symbols[i-1]
		if number >= symbol.arabic {
			factor := number / symbol.arabic
			number = number % symbol.arabic

			roman = roman + strings.Repeat(symbol.roman, factor)
		}

		if i > 1 {
			lesserSymbol := symbols[i-2]
			mapFrom := strings.Repeat(lesserSymbol.roman, 4)
			mapTo := lesserSymbol.roman + symbol.roman

			transformation1 := transformationDefinition{from: mapFrom, to: mapTo}
			transformations = append(transformations, transformation1)

			// if i < (len(symbols) - 1) {
			// 	higherSymbol := symbols[i]
			// 	mapFrom2 := symbol.roman + lesserSymbol.roman + symbol.roman
			// 	mapTo2 := lesserSymbol.roman + higherSymbol.roman

			// 	transformation2 := transformationDefinition{from: mapFrom2, to: mapTo2}
			// 	transformations = append(transformations, transformation2)
			// }

		}
	}

	fmt.Println("2d: ", transformations)

	for _, transformation := range transformations {
		roman = strings.ReplaceAll(roman, transformation.from, transformation.to)
	}

	return roman
}

func main() {

}
