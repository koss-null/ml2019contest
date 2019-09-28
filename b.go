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

func (elem *confusionElem) F1() float64 {
	var pr, re float64
	if elem.tp + elem.fp != 0 {
		pr = float64(elem.tp) / float64(elem.tp+elem.fp)
	}
	if elem.tp + elem.tn!= 0 {
		re = float64(elem.tp) / float64(elem.tp+elem.tn)
	}
	if pr + re == 0 { return 0 }
	return 2 * pr * re / (pr + re)
}

func main() {
	var k int
	_, _ = fmt.Scanf("%d\n", &k)
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

	if  float64(elemsAmount) == 0 {
		fmt.Printf("0\n0\n")
		return
	}

	// macro
	f1 := float64(0)
	pr, re := float64(0), float64(0)
	for _, elem := range elems {
		weight := float64(elem.tp + elem.tn) / float64(elemsAmount)
		if elem.tp + elem.fp != 0 {
			pr += weight * float64(elem.tp) / float64(elem.tp+elem.fp)
		}
		if elem.tp + elem.tn != 0 {
			re += weight * float64(elem.tp) / float64(elem.tp+elem.tn)
		}
	}

	if pr + re != 0 {
		f1 = 2 * pr * re / (pr + re)
	}

	fmt.Println(f1)

	// micro
	f1 = 0.
	for _, elem := range elems {
		weight := float64(elem.tp + elem.tn) / float64(elemsAmount)
		f1 += weight * elem.F1()
	}

	fmt.Println(f1)
}

