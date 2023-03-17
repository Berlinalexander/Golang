package main
import "fmt"
func main()
{
	var a int =42
	var b *int=&a
	fmt.Println(a,b)
	a=27
	fmt.Println(a,*b)
	*b=42
	fmt.Println(a,*b)
	a:=[5]int{1,2,3,4,5}
	b:=&a[1]
	c:=&a[2]
   // c:=&a[2]+4 it is not valid because the pointer will give the address of a[2]
	fmt.Printf( "%v %p %p\n",a,b,c)
    var ms * mystruct
	ms=new (mystruct)
	ms.foo=42 
	fmt.Print(ms.foo)

	ab:=[]int{1,2,3,4}
	c:=ab
	fmt.Println(ab,c)
	a[1]=32
	fmt.Println(ab,c)

	a:=map[string]string{"foo":"bex","baz":"baz"}
	b:=a
	fmt.Println(a,b)
	a["foo"]="qux"
	fmt.Println(a,b)
}
type mystruct struct{
	foo int
}