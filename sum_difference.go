package main 
import (
	"fmt"
    )


func calc(x int, y int)(int, int){
    sum:= x+y
	diff:= x-y
return sum, diff
}
func main(){
	
	result1, result2 := calc( 5, 5)
	fmt. Println( "The sum is", result1)
	fmt. Println( "The difference is", result2)

	
}
