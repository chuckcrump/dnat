package main

import "fmt"

func joeprint(mrna string) {

	codons := theExtraStuff(mrna)
	for _, codon := range codons {
		if codon == "AUG" {
			fmt.Println("Methionine")
		} else if codon == "UAA" || codon == "UAG" || codon == "UGA" {
			fmt.Println("Stop!")
		} else if codon == "UUU" || codon == "UUC" {
			fmt.Println("Phenylalanine")
		} else if codon == "UUA" || codon == "UUG" || codon == "CUU" || codon == "CUC" || codon == "CUA" || codon == "CUG" {
			fmt.Println("Leucine")
		} else if codon == "AUU" || codon == "AUC" || codon == "AUA" {
			fmt.Println("Isoleucine")
		} else if codon == "GUU" || codon == "GUC" || codon == "GUG" || codon == "GUA" {
			fmt.Println("Valine")
		} else if codon == "UCU" || codon == "UCC" || codon == "UCA" || codon == "UCG" {
			fmt.Println("Serine")
		} else if codon == "CCU" || codon == "CCC" || codon == "CCA" || codon == "CCG" {
			fmt.Println("Proline")
		} else if codon == "ACU" || codon == "ACC" || codon == "ACA" || codon == "ACG" {
			fmt.Println("Threonine")
		} else if codon == "GCU" || codon == "GCC" || codon == "GCA" || codon == "GCG" {
			fmt.Println("Alanine")
		} else if codon == "UAU" || codon == "UAC" {
			fmt.Println("Tyrosine")
		} else if codon == "CAU" || codon == "CAC" {
			fmt.Println("Histidine")
		} else if codon == "AAU" || codon == "AAC" {
			fmt.Println("Asparagine")
		} else if codon == "AAA" || codon == "AAG" {
			fmt.Println("Lysine")
		} else if codon == "GAU" || codon == "GAC" {
			fmt.Println("Aspartic acid")
		} else if codon == "GAG" || codon == "GAA" {
			fmt.Println("Glutamic acid")
		} else if codon == "UGU" || codon == "UGC" {
			fmt.Println("Cysteine")
		} else if codon == "UGG" {
			fmt.Println("Tryptophan")
		} else if codon == "CGU" || codon == "CGC" || codon == "CGA" || codon == "CGG" || codon == "AGG" || codon == "AGA" {
			fmt.Println("Arginine")
		} else if codon == "AGU" || codon == "AGC" {
			fmt.Println("Serine")
		} else if codon == "GGU" || codon == "GGC" || codon == "GGA" || codon == "GGG" {
			fmt.Println("Glycine")
		}
	}
}
