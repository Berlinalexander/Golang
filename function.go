package main

import (
	"fmt"
	"math/rand"
)

	func add(n1 int,n2 int) int
	{
      tatal:=n1+n2
	  return total
	}

	func mul(n1 int,n2 int )int{
		mul:=n1*n2
		return mul
	}

	func div(n1 int,n2 int)int{
		div:=n1/n2
		return div
	}
	func rem(n1 int,n2 int)int{
		rem:=n1%n2
		return rem
	}
	func random() int{
		random:=rand.Intn(10)
		return random
	}
	func sum(num ...int) int 
	{
		res:=0
		for _,n:=range num
	    {
			res+=num
		}
		return res
	}
	func main
	{
		sum:=add(10,100)
		fmt.Println(sum)
		mul:mul(10,10)
		fmt.println(mul)
		div:=div(100/10)
		fmt.Println(div)
		fmt.Println(rem(11,10))
		ans:=random()
		fmt.Println(ans)
		s1 := Sum(1, 2, 3) //varidic func with no of parameters
	s2 := Sum(1, 2, 3, 4)
	s3 := Sum(1, 2, 3, 4, 5)

	fmt.Println(s1, s2, s3)

	}
