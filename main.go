package main 

import "fmt"

func main(){

	fmt.Print("Enter your age: ")
	var age int
	fmt.Scanf("%d" , &age)
	
		if age < 3{


		fmt.Println ("You are a baby")
		}else if age>2 && age<16 {
			fmt.Println("You are a children")
		}else if age>16 && age<=30 {
			fmt.Println("You are young adult") 
		
		

	}


