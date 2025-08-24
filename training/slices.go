package training

import "fmt"

// "fmt"
// "slices"
// "strings"

var a1 []int

func SlicesTrainig() {
	// slice := []string{"Alice", "Bob", "Vera"}
    // =================================================================
    // Backward - ревёртит срез для итерации к примеру
	// for i, v := range slices.Backward(slice) {
	// 	fmt.Println(i, ":", v)
	// }
    // =================================================================
	// BinarySearch - бинарный поиск 
    // n, isFound := slices.BinarySearch(slice, "Bob1")
	// fmt.Println(slice[n], isFound)
    // =================================================================
    // BinarySearch - бинарный поиск 
	// type Person struct {
	// 	Name string
	// 	Age  int
	// }
	// people := []Person{
	// 	{"Alice", 55},
	// 	{"Bob", 24},
	// 	{"Gopher", 13},
	// }
	// n, found := slices.BinarySearchFunc(people, Person{"Bob", 0}, func(a, b Person) int {
	// 	return strings.Compare(a.Name, b.Name)
	// })
	// =================================================================
	a2 := []int{1,2,3,4,5}
	// a3 := []int{2:1,4:1,7:1,9:2}
	a := a1
    show(a)
	test := &a2
	*test = append(*test,34)
	fmt.Println(test)
	// show(a)
	// a = a3
	// show(a)
}

func show(p []int) {
  fmt.Print("len=", len(p), " cap=", cap(p), " : ")
  for _, value := range p {
	fmt.Print(value, " ")
  }
  fmt.Println()
}