package main

import (
	"fmt"
	"os"
	"strings"
)

func translatemrna(dna string) string {
	var mrna strings.Builder
	for _, nuctide := range dna {
		switch nuctide {
		case 'A':
			mrna.WriteRune('U')
		case 'T':
			mrna.WriteRune('A')
		case 'G':
			mrna.WriteRune('C')
		case 'C':
			mrna.WriteRune('G')
		default:
			fmt.Println("Invalid nucleotide")
		}
	}
	return mrna.String()
}

func translatetrna(dna string) string {
	var trna strings.Builder
	for _, nuctide := range dna {
		switch nuctide {
		case 'A':
			trna.WriteRune('U')
		case 'U':
			trna.WriteRune('A')
		case 'G':
			trna.WriteRune('C')
		case 'C':
			trna.WriteRune('G')
		default:
			fmt.Println("Invalid nucleotide")
		}
	}
	return trna.String()
}

func theExtraStuff(mrna string) []string {
	fmt.Println("Condons")
	var condons []string
	for i := 0; i < len(mrna); i += 3 {
		end := i + 3
		if end > len(mrna) {
			end = len(mrna)
		}
		condons = append(condons, mrna[i:end])
	}
	return condons
}

func main() {

	for {
		fmt.Println("What do you want to do? \n1. Translate \n2. Exit")
		var choice int
		fmt.Scanln(&choice)

		if choice == 2 {
			fmt.Println("Adiós")
			break
		} else if choice == 1 {
			fmt.Println("Input dna sequence")
			var dnamain string
			fmt.Scanln(&dnamain)

			mrna := translatemrna(dnamain)
			trna := translatetrna(mrna)
			fmt.Println("mRNA:", mrna, "\ntRNA:", trna)
			mRNA := translatemrna(dnamain)
			joeprint(mRNA)

		}
	}
	os.Exit(0)
}
