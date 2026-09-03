package main

// "reflect"
// "strconv"

func addNumbers(a int, b int) int {
	return a + b
}

func operation(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

// func operation(a int, b int) (int, int) {
// 	sum := a + b
// 	diff := a - b
// 	return sum, diff
// }

func sumNumbers(numbers ...int) int {
	result := 0
	for _, element := range numbers {
		result += element
	}
	return result
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

type shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perimeter() float64 {
	return 4 * s.side
}

type Student struct {
	name   string
	rollNo int
	marks  []int
	grades map[string]int
}

type Circle struct {
	x      int
	y      int
	radius float64
	area   float64
}

func calcArea(c *Circle) {
	const PI float64 = 3.14
	area := (PI * c.radius * c.radius)
	(*c).area = area
}

func (c *Circle) calArea() {
	c.area = 3.14 * c.radius * c.radius
}

func main() {
	// var first_name, last_name string = "Abhradip", "Paul"
	// var (
	// 	age       int     = 24
	// 	net_worth float32 = 2.66
	// )
	// is_adult := true
	// fmt.Print("Hello", "\n")
	// fmt.Println("World")
	// fmt.Printf("The name of the user is %v %s \n", first_name, last_name)
	// fmt.Printf("Age of the user is %d \n", age)
	// fmt.Println("Is the user adult", is_adult)
	// fmt.Println("The networth of the user is", net_worth)
	// var (
	// 	random_string string
	// 	random_number int
	// )
	// fmt.Print("Enter random string and random number: ")
	// count, err := fmt.Scanf("%s %d", &random_string, &random_number)
	// fmt.Println("Count:", count)
	// fmt.Println("Err:", err)
	// fmt.Printf("Here is the random string: %s %d \n", random_string, random_number)
	// fmt.Printf("Data type of the variable %s is %T \n", random_string, random_string)
	// fmt.Printf("Type: %v \n", reflect.TypeOf(8.8))
	// fmt.Println(int(net_worth), reflect.TypeOf(strconv.Itoa(age)))

	// Array
	// marks := [3]int{10, 20, 30}
	// fmt.Println(marks)
	// var grades [5]int = [5]int{10, 20, 30, 40, 50}
	// fmt.Println(grades)
	// fruits := [...]string{"apples", "oranges"}
	// for index, element := range fruits {
	// 	fmt.Println(index, "=>", element)
	// }
	// fmt.Println(len(fruits))
	// fmt.Println("Sum of two number is", addNumbers(1, 1))
	// sum, diff := operation(7, 3)
	// fmt.Println(sum, diff)
	// fmt.Println("Sum of numbers:", sumNumbers(1, 2, 3, 4))
	// fmt.Println("Factorial of 5 is", factorial(5))
	// ana := func(a int, b int) int {
	// 	return a * b
	// }
	// fmt.Printf("%T \n", ana)
	// fmt.Println(ana(20, 30))
	// ana := func(a int, b int) int {
	// 	return a * b
	// }(20, 30)
	// fmt.Printf("%T \n", ana)
	// fmt.Println(ana)
	// i := 10
	// str := "Hello"
	// var ptr_i *int = &i
	// var s = &str
	// ptr_s := &str
	// fmt.Println(ptr_i)
	// fmt.Println(s)
	// fmt.Println(ptr_s)
	// 	fmt.Println("Hello")
	// 	st := Student{
	// 		name:   "Abhradip Paul",
	// 		rollNo: 21,
	// 	}
	// 	fmt.Printf("%+v \n", st)
	// 	c := Circle{x: 5, y: 5, radius: 5, area: 0}
	// 	fmt.Printf("%+v", c)
	// 	calcArea(&c)
	// 	fmt.Printf("%+v", c)
}
