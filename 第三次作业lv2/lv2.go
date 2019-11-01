package main

import "fmt"

func Primenumber(num1 int) (int,bool) {
	if num1==1 {
		return num1,true
	}
	for num2:=2; num2<num1; num2++ {
		if num1%num2== 0 {
			return num1,false
		}
	}
	return num1,true
}

func main(){
	for num3:=1; num3<=10000; num3++ {
		go Primenumber(num3)
		fmt.Println(Primenumber(num3))
	}
}