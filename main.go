package main 

import "fmt"

func main(){

	fmt.Print("Enter your age: ")
	var age int
	fmt.Scanf("%d" , &age)
	
	/*	if age < 3{


		fmt.Println ("You are a baby")
		}else if age>2 && age<16 {
			fmt.Println("You are a children")
		}else if age>16 && age<=30 {
			fmt.Println("You are young adult") 
		}else if age >30 && age<=45 {
			fmt.Println("You are middle aged adult")
		}else {fmt.Println("You are old adult")
		} 
	*/


	switch age {
	case 2:
		fmt.Println ("You are a baby")

	case 3,4,5,6,7,8,9,10,11,12,13,14,15:
		fmt.Println ("You are a children")

	case 16,17,18,19,20,21,22,23,24,25,26,27,28,29,30:
		fmt.Println ("You are a young adult")

	case 31,32,33,34,35,36,37,38,39,40,41,42,43,44,45:
		fmt.Println ("You are middle aged adult")

	default:
		fmt.Println ("You are old adult")

			}













		

	}



