package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// London;9.8
// Dubai;29.3
// ...
type TempStorage struct {
	station string
	val     float64
}

type Stats struct {
	min float64
	max float64
	avg float64
}

func calculateStats(values []float64) (min, max, avg float64) {
	min = values[0]
	max = values[0]
	sum := 0.0
	for _, val := range values {
		if val < min {
			min = val
		}

		if val > max {
			max = val
		}

		sum += val
	}

	avg = sum / float64(len(values))
	return
}

func main() {
	// reading the input file
	file, err := os.Open("measurements.txt")
	if err != nil {
		log.Fatalf("Error reading the file : %v", err)
	}

	defer file.Close()

	// the file size is : 16gb
	storageMap := make(map[string][]float64)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")
		if len(parts) != 2 {
			log.Fatalf("Invalid line : %s", line)
		}

		station := parts[0]
		val, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			log.Fatalf("Invalid value : %s", parts[1])
		}

		storageMap[station] = append(storageMap[station], val)
	}

	// calculating the min, max, avg for each station
	withStats := make(map[string]Stats)

	for station, values := range storageMap {
		min, max, avg := calculateStats(values)
		withStats[station] = Stats{min, max, avg}
	}

	// formating output
	var output []string
	for station := range withStats {
		output = append(output, station)
	}

	sort.Strings(output)

	fmt.Print("{")
	for i, station := range output {
		stats := withStats[station]
		fmt.Printf("%s=%.1f/%.1f/%.1f", station, stats.min, stats.max, stats.avg)
		if i < len(output)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("}")

}
