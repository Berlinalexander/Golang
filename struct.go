package main
import "fmt"
type netflix struct{
	 Number int
	 Name string 
	 Naration [] string 
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
}