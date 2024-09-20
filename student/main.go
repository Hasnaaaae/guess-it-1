package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 1 {
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	Data_Manipulation(scanner)
}

func Data_Manipulation(scanner *bufio.Scanner){
	var tab []float64
	var average, sd, max, min float64
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println(err)
		} else if num > 1000 {
			fmt.Printf("%v %v", 1200, 1500)
			fmt.Println()
		} else {
			tab = append(tab, num)
			average = Average(tab)
			sd = Standard_Deviation(Variance(tab, average))
			if len(tab) == 1 {
				max = 950
				min = 1
			} else {
				max = average + sd
				min = math.Max(0, average-sd)
			}
			fmt.Println(int(min), int(max))
		}
	}
}

// Median function
func Average(tab []float64) float64 {
	var average, somme float64
	for i := 0; i < len(tab); i++ {
		somme = somme + tab[i]
	}
	average = somme / float64(len(tab))
	return average
}

// Variance function
func Variance(tab []float64, average float64) float64 {
	var variance float64
	for i := 0; i < len(tab); i++ {
		variance += math.Pow(tab[i]-average, 2)
	}
	variance = variance / float64(len(tab))
	return variance
}

// Standard_Deviation function
func Standard_Deviation(variance float64) float64 {
	sd := math.Sqrt(variance)
	return sd
}
