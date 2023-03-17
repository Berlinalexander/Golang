package main
import {
	"fmt"
	"reflect"
}
type netflix struct{
	 Number int
	 Name string 
	 Naration [] string 
}
type animal struct{
	name string
	location string
}
type bird struct{
	animal
	speed int
	canfly bool
}
type animal1 struct{
	name string 'required max="100"'
	origin string
}
func main()
{
	// netflix1:=netflix
	// {
	// 	1,
	// 	"moneyheist",
	// 	[] string {
	// 		" hero",
	// 		" heroin",
	// 	},
	// }
	netflix1:=netflix
	{
		Number:1,
		Name:"moneyheist",
		Narration:[] string {
			" hero",
			" heroin",
		},
	}
	fmt.Println(netflix1)
	fmt.Println(netflix1.Naration)
	aname:=struct{name string}{name:berlin}//anonymous struct
	anotheraname:=aname
	anotheraname.name=rookie
	fmt.Println(anotheraname)
	fmt.Println(aname)
	anotheraname:=&aname //pointer  to the aname 
	anotheraname.name=rookie
	fmt.Println(anotheraname) 
	fmt.Println(aname)//it change to the aname

	b:=bird{}
	b.name="emu"
	b.location="australia"
	b.speed="125"
	b.canfly=false
	fmt.Print(b.name)
    t:=reflect.TypeOf(animal1{})
	field,_:=t.FieldByname{"name"}
	fmt.Println(field.Tag)
   

     
}