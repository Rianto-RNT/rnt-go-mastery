package main

import "fmt"

type car struct {
	brand string
	price int
}

func changeCar(c car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func (c *car) changeCar2(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice
}

// type distance *int

// func (d *distance) m1() {
// 	fmt.Println("example rnt message")
// }

func main() {
	myCar := car{brand: "McLaren", price: 70000}
	changeCar(myCar, "Ferrari", 100000)
	fmt.Println(myCar)

	myCar.changeCar1("Bugati", 21999)
	fmt.Println(myCar)

	myCar.changeCar2("Dokar", 40000)
	fmt.Println(myCar)

	var yourCar *car
	yourCar = &myCar
	fmt.Println(*yourCar)

	// valid way
	yourCar.changeCar2("Toyota", 20000)
	fmt.Println(*yourCar)

	(*yourCar).changeCar2("More Toyota", 20000)
	fmt.Println(*yourCar)

	fmt.Println(myCar)
}