package main

import (
	"fmt"
	"practise/Product"
)


func main() {
	var a,b int
	fmt.Println("Enter cost os the product ")
	fmt.Scanf("%d",&a)
	fmt.Println("enter teh number of product ")
	fmt.Scanf("%d",&b)
	k := Product.Product(a,b)
	println("total price ",k)
}
