package main

func addition(a int, b int) (ac int, bc int) {
	ac = a + b
	bc = a + b
	return
}

type Person struct {
	name   string
	job    string
	age    int
	salary int
}

func main() {

	result1, result2 := addition(6, 7)
	println("first result: %v", result1)
	// result2 := addition(12, 789)
	println("result 2: %v", result2)

	var persone1 Person
	persone1.name = "Mohammad Mansoor Safi"
	persone1.age = 26
	persone1.job = "Senior Fullstack developer"
	persone1.salary = 60000

	var person2 Person
	person2.name = "Shafi"
	person2.age = 26
	person2.job = "Teacher"
	person2.salary = 6000

	println("person1 name: %v", persone1.name)
	println("person2 name: %v", person2.name)

}
