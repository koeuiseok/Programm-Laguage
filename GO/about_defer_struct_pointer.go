package main

import "fmt"

type Vertex struct {
	X int
	Y int
	Z bool
}

func main() {
	v := Vertex{1, 2, false}
	w := &v
	defer fmt.Println("line 14 : ", v.X)
	fmt.Println("line 15", v.X)
	v.X += 12
	fmt.Println("line 17", v.X)
	fmt.Printf("w : %v, *w : %v\n", w, *w)
	fmt.Printf("*w.X : %v\n", (*w).X)
	fmt.Printf("w.X : %v\n", w.X)
}


/*
output :
=================
line 15 1
line 17 13
w : &{13 2 false}, *w : {13 2 false}
*w.X : 13
w.X : 13
line 14 :  1
=================
go는 포인터, defer(나중에 실행), struct 기능을 지원
포인터는 값이 아닌 그 자체로 타입을 갖는다.
구조체 포인터는 (*(&v)).X와 (&v).X와 v.X가 모두 같은 것을 확인할 수 있다
*/
