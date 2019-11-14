package main

import (
	// "fmt"
	// "math"
	// "strings"
)

type person struct {
	name string
	// firstName string
	// lastName string
	age int
	// tangan tangan
}

// type tangan struct{
// 	warna string
// 	panjang int
// }

// func (P person) getName ()string {
// 	return P.name + " amazing!"
// }

// func (P *person) increaseAge(){
// 	P.age = P.age + 1
// }

// type Rect struct {
// 	width float64
// 	height float64
// 	area float64
// }

// type Circle struct {
// 	radius float64
// }

// func (r Rect) Area () float64 {
// 	return r.width * r.height
// }

// func (r *Rect) CalculateArea () {
// 	r.area = r.width * r.height
// }

// func (c Circle) Area () float64{
// 	return math.Pi * c.radius * c.radius
// }

type square struct{
	side int
}

type calculate interface {
	large () int
	doubleSide () int
	nSide(n int) int
}


func day6()  {
	//DEMO MATERI
	// ========= contoh interface =========
	// var secret interface {}

	// secret = 2
	// var number = secret.(int) * 10
	// fmt.Println(secret, "multiplied by 10 is :", number)

	// secret = []string{"apple", "mango", "banana"}
	// var gruits = strings.Join(secret.([]string), ", ")
	// fmt.Println(gruits, "is my favorite fruits")

	// secret = 45
	// explain(secret)
	// fmt.Println()

	// var dimResult calculate
	// dimResult = square{10}
	// fmt.Println("large square:", dimResult.large())
	// fmt.Println("large square:", dimResult.doubleSide())

	// func (s square) large () int {
	// 	return s.side * s.side
	// }

	// func (s square) doubleSide () int {
	// 	return s.side * 2
	// }

	// func (s square) nSide(n int) int {
	// 	return s.side * n
	// }

	// ========= contoh method ==========
	// rect := Rect{
	// 	width: 5.0,
	// 	height: 4.0,
	// }
	// // cir := Circle{5.0}
	// fmt.Printf ("Area of retangle rect = %0.2f\n", rect.Area())
	// // fmt.Printf ("Area of circle cir = %0.2f\n", cir.Area())

	// fmt.Printf ("Retangle before call pointer method = %v \n", rect)
	// rect.CalculateArea()
	// fmt.Printf ("Retangle after call pointer method = %v\n", rect)

	// =========== contoh struct ===========
	// personA := person{"john", 50}
	// fmt.Printf ("%v \n", personA)
	// fmt.Println(personA.getName())

	// personA.increaseAge()
	// fmt.Println (personA.age)

	// var Person person
	// Person.firstName = "sephia"
	// Person.lastName = "aldiariza"
	// Person.age = 23
	// fmt.Println (Person.firstName, Person.lastName, Person.age)

	// var Person1 = person {"rizky","Kurniawan",23,  tangan {"putih",40}}
	// fmt.Println(Person1)

	// var Person2 = person {
	// 	firstName : "aldia",
	// 	lastName : "hernandha",
	// 	age : 23,
	// }
	// fmt.Println (Person2)

	// Person3 := person{"pranadya","bagus", 23}
	// fmt.Println(Person3)

	// Person4 := new(person)
	// Person4.firstName = "muhammad"
	// Person4.lastName = "ismail"
	// Person4.age = 30
	// fmt.Println (*Person4)

	// var Person5 = person{
	// 	firstName : "sesa",
	// 	lastName : "sesi",
	// 	tangan: tangan{
	// 		warna : "kuning",
	// 		panjang : 30,
	// 	},
	// }
	// fmt.Println(Person5)
}