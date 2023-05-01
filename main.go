package main

// Lesson 2
// import "fmt"

// func main(){

// 	fmt.Println("Hello, ningas!")
// }

// Lesson 3

// import "fmt"

// func main(){
// var nameOne string = "mario"
// var nameTwo = "luigi"
// var nameThree string

// fmt.Println(nameOne, nameTwo, nameThree)

// nameOne = "peach"
// nameThree = "bowser"

// fmt.Println(nameOne, nameTwo, nameThree)

// // Shorthand for declaring a variable
// nameFour := "yoshi"
// fmt.Println(nameFour)

//Ints
// var ageOne int = 20
// var ageTwo =30
// ageThree := 40

// fmt.Println(ageOne, ageTwo, ageThree)

// bits and memory
// var numOne int8 = 25
// var numTwo int8 = -128
// var numThree uint8 = 25 // uint means it cannot be a negative number

// floats
// var scoreOne float32 = 25.5
// var scoreTwo float64 = 330900244769855.32
// }

// Lesson 4

// import "fmt"

// func main(){

// age := 35
// name := "Shaun"
// Print
// fmt.Print("Hello, ")
// fmt.Print("world \n")
// fmt.Print("new line \n")

//Println
// fmt.Println("hello, ninjas")
// fmt.Println("goodbye ningas")
// fmt.Println("my age is", age, "and my name is", name)

// Printf(formatting strings) %_ = formatting specifier
// fmt.Printf("my age is %v and my name is %v \n", age, name)
// fmt.Printf("my age is %q and my name is %q \n", age, name)
// fmt.Printf("my age is %T \n", age)
// fmt.Printf("You have scored %f points \n", 255.3443)
// fmt.Printf("You have scored %0.1f points \n", 255.3443)

//Sprintf(save formatted string)
// var str = fmt.Sprintf("my age is %v and my name is %v", age, name)
// fmt.Println("The saved string is:", str)

// }

// Lesson 5
// Arrays

// import "fmt"

// func main(){
// var ages [3]int = [3]int{20, 30, 40}
// var ages = [3]int{20, 30, 40}

// names := [4]string{"yoshi", "mario", "peach", "bowser"}

// fmt.Println(ages, len(ages))
// fmt.Println(names, len(names))

// Slice
// var scores = []int{100, 50, 60}
// scores[1] = 78
// scores = append(scores, 85)
// fmt.Println(scores, len(scores))

// Slice ranges
// 	rangeOne := scores[1:3]
// 	rangeTwo := scores[2:]
// 	rangeThree := scores[:3]

// 	fmt.Println(rangeOne, rangeTwo, rangeThree)
// }

// Lesson 6
// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// greeting := "hello there friends!"

// fmt.Println(strings.Contains(greeting, "hello"))
// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
// fmt.Println(strings.ToUpper(greeting))
// fmt.Println(strings.ToLower(greeting))
// fmt.Println(strings.Index(greeting, "ll"))
// fmt.Println(strings.Split(greeting, "e"))

// ages := []int{23, 39, 54, 65, 76, 70, 21, 45}
// sort.Ints(ages)
// // fmt.Println(ages)

// index := sort.SearchInts(ages, 76)
// fmt.Println(index)

// names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
// sort.Strings(names)
// fmt.Println(names)
// fmt.Println(sort.SearchStrings(names, "bowser"))
// }

// Lesson 7
// For loops

// import "fmt"

// func main(){
// x := 0
// while loop
// for x < 5 {
// 	fmt.Println("The value of x is:", x)
// 	x++
// }

// for loop
// for i := 0; i < 5; i++ {
// 	fmt.Println("The value of i is:", i)
// }

// names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

// for i := 0; i < len(names); i++ {
// 	fmt.Println(names[i])
// }

// for in loop
// for index, value := range names {
// 	fmt.Printf("The value at index %v is %v \n", index, value)
// }

// 	for _, value := range names {
// 		fmt.Printf("The value is %v \n", value)
// 	}
// }

// Lesson 8

// import "fmt"

// func main(){
// age := 45

// fmt.Println(age <= 50)
// fmt.Println(age >= 50)
// fmt.Println(age == 45)
// fmt.Println(age != 50)

// if age < 30 {
// 	fmt.Println("age is less than 30")
// } else if age < 40 {
// 	fmt.Println("age is less than 40")
// } else {
// 	fmt.Println("age is not less than 45")
// }

// 	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

// 	for index, value := range names {
// 		if index == 1 {
// 			fmt.Println("continuing at pos", index)
// 			continue
// 		}

// 		if index > 2 {
// 			fmt.Println("breaking at pos", index)
// 			break
// 		}
// 		fmt.Printf("The value at pos %v is %v \n", index, value)
// 	}
// }

// Lesson 9
// Function

// import (
// 	"fmt"
// 	"math"
// )

// func sayGreeting(n string) {
// 	fmt.Printf("Good morning %v \n", n)
// }

// func sayBye(n string) {
// 	fmt.Printf("Goodbye %v \n", n)
// }

// func cycleName (n []string, f func(string)){
// 	for _, v := range n {
// 		f(v)
// 	}
// }

// func circleArea(r float64) float64{
// 	return math.Pi * r * r
// }

// func main( ) {
// sayGreeting("mario")
// sayGreeting("luigi")
// sayBye("mario")

// 	cycleName([]string{"cloud", "tifa", "barret"}, sayGreeting)
// 	cycleName([]string{"cloud", "tifa", "barret"}, sayBye)

// 	a1 := circleArea(10.5)
// 	a2 := circleArea(15)

// 	fmt.Println(a1, a2)
// 	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f\n", a1, a2)
// }

// import (
// 	"fmt"
// 	"strings"
// )

// func getInitials(n string)(string, string){
// 	s := strings.ToUpper(n)
// 	names := strings.Split(s, " ")

// 	var initials []string
// 	for _, v := range names {
// 		initials = append(initials, v[:1])
// 	}

// 	if len(initials) > 1 {
// 		return initials[0], initials[1]
// 	}

// 	return initials[0], "_"
// }

// func main(){
// 	fn, sn := getInitials("tifa lockhart")
// 	fmt.Println(fn, sn)

// 	fn2, sn2 := getInitials("barret")
// 	fmt.Println(fn2, sn2)
// }

// import "fmt"

// var score = 99.5

// func main (){
// 	sayHello("Mario")

// 	for _, v := range points {
// 		fmt.Println(v)
// 	}

// 	showScore()
// }

// import "fmt"

// func main(){
// 	menu := map[string]float64{
// 		"soup": 4.99,
// 		"pie": 7.99,
// 		"salad": 6.99,
// 		"toffee pudding": 3.55,
// 	}
// 	fmt.Println(menu)
// 	fmt.Println(menu["pie"])

// Looping maps
// for k, v := range menu{
// 	fmt.Println(k, "-", v)
// }

// int as key type

// 	phonebook := map[int]string{
// 		234234: "mario",
// 		234423: "peach",
// 		342342: "luigi",
// 	}

// 	fmt.Println(phonebook)
// 	fmt.Println(phonebook[234234])

// 	phonebook[234423] = "bowser"
// 	fmt.Println(phonebook)
// }

// import "fmt"

// func updateName(x string) string{
// 	x = "wedge"
// 	return x
// }

// func updateMenu(y map[string]float64){
// 	y["coffee"] = 2.99
// }

// func main(){
// 	name := "tifa"
// types: strings, int, float, bools, array, struct

// name = updateName(name)

// fmt.Println(name)

// types: slices, maps, functions

// 	menu := map[string]float64{
// 		"pie": 5.95,
// 		"ice cream" : 3.99,
// 	}

// 	updateMenu(menu)
// 	fmt.Println(menu)
// }

// import "fmt"

// func updateName(x *string){
// 	*x = "wedge"
// }

// func main(){
// 	name := "tifa"

// updateName(name)

// fmt.Println("The memory location for name is:", &name)

// m := &name
// fmt.Println("memory location is:", m)
// fmt.Println("Value at memory location is:", *m)
// 	fmt.Println(name)
// 	updateName(m)
// 	fmt.Println(name)
// }
// import "fmt"

// func main(){
// 	mybill := newBill("mario's bill")

// 	mybill.addItem("onion soup", 4.50)
// 	mybill.addItem("veg pie", 5.50)
// 	mybill.updateTip(10)

// 	fmt.Println(mybill.format())
// }

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func getInput(prompt string, r *bufio.Reader) (string, error) {
// 	fmt.Print(prompt)
// 	input, err := r.ReadString('\n')
// 	return strings.TrimSpace(input), err
// }

// func createBill() bill{
// 	reader := bufio.NewReader(os.Stdin)

// 	name, _ := getInput("Create a new bill name: ", reader)

// 	b := newBill(name)
// 	fmt.Println("Created the bill - ", b.name)
// 	return b
// }

// func promptOptions(b bill){
// 	reader := bufio.NewReader(os.Stdin)

// 	opt, _ := getInput("Choose options (a - add item, s - save bill, t - add tip): ", reader)

// 	switch opt {
// 	case "a":
// 		name, _ := getInput("Item name: ", reader)
// 		price, _ := getInput("Price: ", reader)

// 		p, err := strconv.ParseFloat(price, 64)
// 		if err != nil {
// 			fmt.Println("The price must be a number")
// 			promptOptions(b)
// 		}
// 		b.addItem(name, p)

// 		fmt.Println("Item added - ", name, price)
// 		promptOptions(b)
// 	case "t":
// 		tip, _ := getInput("Enter tip amount ($): ", reader)

// 		t, err := strconv.ParseFloat(tip, 64)
// 		if err != nil {
// 			fmt.Println("The tip must be a number")
// 			promptOptions(b)
// 		}
// 		b.updateTip(t)
// 		fmt.Println(tip)
// 		promptOptions(b)
// 	case "s":
// 		b.save()
// 		fmt.Println("You saved the file - :", b.name)

// 	default:
// 		fmt.Println("That was not a valid option...")
// 	promptOptions(b)
// 	}
// }

// func main(){
// 	mybill := createBill()
// 	promptOptions(mybill)
// }

import (
	"fmt"
	"math"
)

// shape interface
type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}
type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}

// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.circumf())
}

func main() {
	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}
}