//The sum of the squares of the first ten natural numbers is,
//The square of the sum of the first ten natural numbers is,
//Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 ? 385 = 2640.
//Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

package main

import (
	"fmt"
	"time"
)


func main() {
	var sumSq,sum,dif int
        t := time.Now();
	for j:= 1; j<=1000;j++{
        	for i:= 1; i<j;i++ {
			sumSq += i*i
			sum += i  
		}
		dif = (sum*sum) - sumSq
		fmt.Println("dif = ",dif)
		sumSq = 0
		sum = 0  
	}
	fmt.Println(" time: ", time.Since(t))
}