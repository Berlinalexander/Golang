package main

import "fmt"

func main() {
	var car [5]int
	car[0] = 0
	car[1] = 1
	car[2] = 2
	car[3] = 3
	car[4] = 4
	fmt.Println(car)
	fmt.Println(len(car))
	bike := [3]string{"BMW", "duke", "eee"}
	fmt.Println(len(bike))
	bike2 := &bike
	fmt.Println(len(bike2))
	car1 := []int{1, 2, 3, 4}
	fmt.Println(len(car1))
	fmt.Println(cap(car1))
	car3 := car1[:]
	car4 :=
		car1[2:]
	car5 :=
		car1[:2]
	fmt.Println(car3)
	fmt.Println(car4)
	fmt.Println(car5)
	hello := []int{}
	hello = append(hello, 1)
	fmt.Println(len(hello))
	fmt.Println(cap(hello))
	hello = append(hello, 2, 3, 4, 5)
	fmt.Println(len(hello))
	fmt.Println(cap(hello))

}
