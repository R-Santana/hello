package main 

import ("fmt"
		"github.com/r-santana/stringutil"
		
		)

func main(){
	fmt.Printf(stringutil.Reverse("!oG ,olleH")+",rev func in stringutil package\n")
	fmt.Printf(stringutil.Richie("!boj ,doog")+",rich func in stringutil package\n")
	rsantan()
	
}
func rsantan(){
	fmt.Printf("rsantan func in main (local) package\n") 
}	
//asdfasd