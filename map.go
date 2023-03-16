package main
import "fmt"
func main()
{
	localdata:=map[string] int
	{
		"berlin" :001,
		 "london":002,
		 "losAngels":003,
		 "newdelhi":004,
	}
	fmt.Println(localdata)
	localdata[bangalore]="005"
	fmt.Println(localdata[bangalore])
	delete(localdata,"bangalore")
	pop, ok:=localdata["bank"]
	fmt.println(pop,ok)
	fmt.println(len(localdata))
}