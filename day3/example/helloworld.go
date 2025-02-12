package main

import "fmt"

type student struct {
	Name string
	Rgno float64
	Dept string

} 

func main() {	
	var q string = "hello world"
	fmt.Println(q)
	st :=student{Name: "student1", Rgno: 12.2 , Dept: "cs" }
    fmt.Println(st)
	fmt.Println("name: ",st.Name,"\nrgno: ",st.Rgno,"\ndept: ",st.Dept)
	
}
func add() {	
	var x,y int =1,2
	var z int =x+y
    fmt.Println(z)
}


func ifelsedemo() {
	var a,b int
	fmt.Scanln(&a,&b)
	if (a>b) {
		fmt.Println("a is larger than b")
	} else if (a==b) {
		fmt.Println("a is equal to b")
	}else {
		fmt.Println("a is smaller than b")
	}
}


func forthree() {
	s :=0
	for i :=0 ; i<5 ; i++ {
		s=s+i
	}
	fmt.Println(s)
}


func forcond() {
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n)
}

func forpythonstyle() {
	strings := []string{"hello","world","golang","nie"}
	for i,s := range strings {
		fmt.Println(i,s)
	}
}