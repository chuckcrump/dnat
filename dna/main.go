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

func main() {

	for {
		fmt.Println("What do you want to translate? \n1. DNA to mRNA \n2. mRNA to tRNA \n3. Exit")
		var choice int
		fmt.Scanln(&choice)

		if choice == 3 {
			fmt.Println("Adiós")
			break
		} else {
			fmt.Println("Invalid choice")
		}

		fmt.Println("Input dna sequence")
		var dnamain string
		fmt.Scanln(&dnamain)

		if choice == 1 {
			mrna := translatemrna(dnamain)
			fmt.Println("mRNA =", mrna)
		} else if choice == 2 {
			trna := translatetrna(dnamain)
			fmt.Println("tRNA =", trna)
		}
	}
	os.Exit(0)
}
