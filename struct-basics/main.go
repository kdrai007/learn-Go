package main

type car struct {
	brand     string
	backWheel wheel
}

type wheel struct {
	radius int
}

type wheelType struct {
	wheel
	name string
}

// struct method

type ract struct {
	width  int
	height int
}

func (r ract) area() ract {
	r.width = r.width * r.height
	r.height = r.width / r.height
	return r
}

func isClear(hisCar car) bool {
	if hisCar.brand == "" || hisCar.backWheel.radius <= 0 {
		return false
	}
	return true
}

func main() {
	//myCar := car{}
	//myCar.brand = "benz"
	//myCar.backWheel.radius = 0

	//myWheel := wheelType{}
	//myWheel.radius = 0
	//myWheel.name = "something in it"
	//
	//fmt.Println(myWheel)

	// var foo = ract{
	// 	width:  25,
	// 	height: 10,
	// }

	// bar := foo.area()
	// fmt.Println(bar)

	//Exercises

	// exercise()
	// exercise_two()
	exercise_three()
}
