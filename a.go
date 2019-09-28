package main
 
import (
	"fmt"
	"bufio"
	"os"
)
 
type Round struct {
	i int
	roundLength int
}
 
func (r *Round) next() {
	r.i++
	if r.i == r.roundLength {
		r.i = 0
	}
}
 
func NewRound(k int) Round {
	return Round {0, k}
}
 
// fast input
func getArray(line []byte) ([]int) {
	a := make([]int, 0, 10000)
 
	val := 0
	for i := 0; i < len(line); i++ {
		char := line[i]
		if char == ' ' {
			a = append(a, val)
			val = 0
			continue
		}
		val = val*10 + int(char) - '0'
	}
 
	a = append(a, val)
	return a
}
 
func main() {
	var objN, classN, partsN int
	fmt.Scanf("%d %d %d\n", &objN, &classN, &partsN)
 
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, objN * 8), objN * 10)
	scanner.Scan()
 
	objects := make([][]int, classN)
 
	objsRaw := getArray(scanner.Bytes())
	for i, obj := range objsRaw {
		obj--
		objects[obj] = append(objects[obj], i+1)
	}
 
	splitted := make([][] int, partsN)
	rnd := NewRound(partsN)
	for _, object := range objects {
		for i := range object {
			splitted[rnd.i] = append(splitted[rnd.i], object[i])
			rnd.next()
		}
	}
 
	for _, class := range splitted {
		fmt.Printf("%d ", len(class))
		for _, item := range class {
			fmt.Printf("%d ", item)
		}
		fmt.Println()
	}
}

