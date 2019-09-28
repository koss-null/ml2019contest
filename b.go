package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func sum(arr []float64) (s float64) {
	for i := range arr {
		s += arr[i]
	}
	return
}

func sumRow(arr [][]float64, rowInd int) (s float64) {
	for i := range arr {
		s += arr[i][rowInd]
	}
	return
}

type confusionElem struct {
	tp float64
	tn float64
	fp float64
}

func normalize(m [][]float64) [][]float64 {
	for row := range m {
		divider := sumRow(m, row)
		for i := range m[row] {
			m[row][i] /= divider
		}
	}
	return m
}

func main() {
	var k int
	fmt.Scanf("%d\n", &k)
	reader := bufio.NewReader(os.Stdin)

	confusion := make([][]float64, 0, k)

	for i := 0; i < k; i++ {
		confusion = append(confusion, make([]float64, 0, k))
		text, _ := reader.ReadString('\n')
		for _, numStr := range strings.Fields(text) {
			num, _ := strconv.Atoi(numStr)
			confusion[i] = append(confusion[i], float64(num))
		}
	}

	confusion = normalize(confusion)

	elemsAmount := 0.
	elems := make([]confusionElem, k)
	for i := range elems {
		elems[i] = confusionElem{
			confusion[i][i],
			sum(confusion[i]) - confusion[i][i],
			sumRow(confusion, i) - confusion[i][i],
		}
		elemsAmount += elems[i].tp + elems[i].tn
	}

	// macro
	f1 := float64(0)
	pr, re := float64(0), float64(0)
	for _, elem := range elems {
		weight := 1. // float64(elem.tp + elem.tn) / float64(elemsAmount)
		if elem.tp + elem.tn != 0 {
			pr += float64(elem.tp) / float64(elem.tp+elem.fp)
		}
		if elem.tp + elem.fp != 0 {
			re += float64(elem.tp) / float64(elem.tp+elem.tn)
		}
		if pr + re != 0 {
			f1 = weight * 2 * pr * re / (pr + re)
		}
	}

	fmt.Println(f1)

	// micro
	f1 = float64(0)
	tp, fp := 0., 0.
	for _, elem := range elems {
		weight := 1. // float64(elem.tp + elem.tn) / float64(elemsAmount)
		tp += float64(elem.tp) * weight
		fp += float64(elem.fp) * weight
	}

	f1 = float64(tp) / float64(tp+fp)

	fmt.Println(f1)
}

