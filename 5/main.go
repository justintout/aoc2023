package main

import (
	"cmp"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type category struct {
	maps [][]int
}

func (c *category) addMap(m []int) {
	c.maps = append(c.maps, m)
}

func (c *category) sort() {
	slices.SortFunc(c.maps, func(a, b []int) int {
		return cmp.Compare(a[1], b[1])
	})
}

func (c category) apply(n int) int {
	var i int
	for i = 0; i < len(c.maps); i++ {
		if n < c.maps[i][1] {
			i = -1
			break
		}
		if n > c.maps[i][1]+c.maps[i][2]-1 {
			continue
		}
		break
	}
	if i == -1 || i >= len(c.maps) || (n < c.maps[i][1] || n > c.maps[i][1]+c.maps[i][2]) {
		return n
	}
	return c.maps[i][0] + (n - c.maps[i][1])
}

var digitRegex = regexp.MustCompile(`\d+`)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	all := string(b)
	sections := strings.Split(all, "\n\n")

	var (
		seeds                 []int
		seedToSoil            category
		soilToFertilizer      category
		fertilizerToWater     category
		waterToLight          category
		lightToTemperature    category
		temperatureToHumidity category
		humidityToLocation    category
	)

	for i, section := range sections {
		if i == 0 {
			nn := digitRegex.FindAllString(section, -1)
			for _, n := range nn {
				sn, err := strconv.Atoi(n)
				if err != nil {
					panic(err)
				}
				seeds = append(seeds, sn)
			}
			continue
		}
		var c category
		s := strings.Split(section, "\n")
		for _, ms := range s[1:] {
			var m []int
			nn := digitRegex.FindAllString(ms, -1)
			for _, n := range nn {
				x, err := strconv.Atoi(n)
				if err != nil {
					panic(err)
				}
				m = append(m, x)
			}
			if len(m) < 3 {
				panic(fmt.Sprintf("%s: map not length 3: %v", s[0], m))
			}
			c.addMap(m)
		}
		c.sort()
		switch i {
		case 1:
			seedToSoil = c
		case 2:
			soilToFertilizer = c
		case 3:
			fertilizerToWater = c
		case 4:
			waterToLight = c
		case 5:
			lightToTemperature = c
		case 6:
			temperatureToHumidity = c
		case 7:
			humidityToLocation = c
		default:
			panic("category out of range: " + strconv.Itoa(i))
		}
	}

	solve := func(n int) int {
		return humidityToLocation.apply(temperatureToHumidity.apply(lightToTemperature.apply(waterToLight.apply(fertilizerToWater.apply(soilToFertilizer.apply(seedToSoil.apply(n)))))))
	}

	var locations []int
	for _, s := range seeds {
		locations = append(locations, solve(s))
	}
	slices.Sort(locations)
	fmt.Printf("part 1: %d\n", locations[0])

	// fmt.Println(seedToSoil.apply(82))
	// fmt.Println(soilToFertilizer.apply(seedToSoil.apply(82)))
	// fmt.Println(fertilizerToWater.apply(soilToFertilizer.apply(seedToSoil.apply(82))))
	// fmt.Println(waterToLight.apply(fertilizerToWater.apply(soilToFertilizer.apply(seedToSoil.apply(82)))))
	// fmt.Println(lightToTemperature.apply(waterToLight.apply(fertilizerToWater.apply(soilToFertilizer.apply(seedToSoil.apply(82))))))
	// fmt.Println(temperatureToHumidity.apply(lightToTemperature.apply(waterToLight.apply(fertilizerToWater.apply(soilToFertilizer.apply(seedToSoil.apply(82)))))))
	// fmt.Println(humidityToLocation.apply(temperatureToHumidity.apply(lightToTemperature.apply(waterToLight.apply(fertilizerToWater.apply(soilToFertilizer.apply(seedToSoil.apply(82))))))))

	// there's probably some smart way to analyze breakpoints to solve this quickly...

	lowest := math.MaxInt
	for i := 1; i < len(seeds); i += 2 {
		for j := seeds[i-1]; j < seeds[i-1]+seeds[i]; j++ {
			l := solve(j)
			if l < lowest {
				lowest = l
			}
		}
	}

	fmt.Printf("part 2: %d\n", lowest)
}
