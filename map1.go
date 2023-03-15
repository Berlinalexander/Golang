package map1

import "fmt"

func map1()
{
 m:make(map[string] int)
 m["abc"]=7
 m["bca"]=11
 fmt.Println("map:",m)
 value:=m["abc"]
 fmt.Println(value)

}
