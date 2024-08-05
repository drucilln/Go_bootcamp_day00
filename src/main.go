package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	meanF := flag.Bool("mean", false, "Print mean")
	medianF := flag.Bool("median", false, "Print median")
	modeF := flag.Bool("mode", false, "Print mode")
	sdF := flag.Bool("sd", false, "Print standard deviation")
	all := flag.Bool("all", true, "Print all metrics")

	flag.Parse()

	if *meanF || *medianF || *modeF || *sdF {
		*all = false
	}

	arr := []int{}
	var sum float64
	arr, sum = fillArray(arr, sum)
	cnt := len(arr)
	if cnt != 0 {
		meanCalc(sum, cnt, *meanF, *all)
		medianCalc(arr, cnt, *medianF, *all)
		modeCalc(arr, *modeF, *all)
		sdCalc(arr, sum, cnt, *sdF, *all)
	}
}

func fillArray(arr []int, sum float64) ([]int, float64) {
	var tmp, cnt int
	scanner := bufio.NewScanner(os.Stdin)
	var err error
	fmt.Println("Введите количество элементов последовательности:")
	scanner.Scan()
	input := scanner.Text()
	cnt, err = strconv.Atoi(strings.TrimSpace(input))
	if err == nil {
		fmt.Println("Введите последовательность:")
		for i := 0; i < cnt; i++ {
			scanner.Scan()
			input := scanner.Text()
			tmp, err = strconv.Atoi(strings.TrimSpace(input))
			if err != nil || tmp < -100000 || tmp > 100000 {
				fmt.Println("Ошибка: невалидный элемент")
				arr = arr[:0]
				cnt = 0
			} else {
				arr = append(arr, tmp)
				sum += float64(tmp)
			}
		}
	} else {
		fmt.Println("Ошибка: невалидное количество элементов")
	}
	return arr, sum
}

func meanCalc(sum float64, cnt int, meanF bool, all bool) {
	if meanF || all {
		fmt.Printf("Mean: %.2f\n", sum/float64(cnt))
	}
}

func medianCalc(arr []int, cnt int, medianF bool, all bool) {
	if medianF || all {
		sort.Ints(arr)
		var median float64
		if cnt%2 == 0 {
			median = float64(arr[cnt/2-1]+arr[cnt/2]) / 2
		} else {
			median = float64(arr[cnt/2])
		}
		fmt.Printf("Median: %.2f\n", median)
	}
}

func modeCalc(arr []int, modeF bool, all bool) {
	if modeF || all {
		counts := make(map[int]int)
		for _, num := range arr {
			counts[num]++
		}
		var maxCount, mode int
		for num, count := range counts {
			if count > maxCount || (count == maxCount && num < mode) {
				mode = num
				maxCount = count
			}
		}
		fmt.Println("Mode: ", mode)
	}
}

func sdCalc(arr []int, sum float64, cnt int, sdF bool, all bool) {
	if sdF || all {
		sdArr := []float64{}
		var sd float64
		mean := sum / float64(cnt)
		sum = 0
		for idx := range arr {
			sdArr = append(sdArr, math.Pow(float64(arr[idx])-mean, 2))
			sum += sdArr[idx]
		}
		sd = sum / float64(cnt)
		sd = math.Sqrt(sd)
		fmt.Printf("Sd: %.2f\n", sd)
	}
}
