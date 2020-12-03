package main

import "fmt"

//-----------------------------------------------------
// func main() {
// 	var s string
// 	s = "complete"
// 	ConnectedConfirm(&s)
// 	fmt.Println(s)
// }

// func ConnectedConfirm(status *string) {
// 	if *status == "complete" {
// 		*status = "connected"
// 	}
// }
//-----------------------------------------------------

type foo struct {
	id   int
	name string
}

type bigFoo struct {
	Foo foo
}
type bigStarFoo struct {
	Foo *foo
}

type too struct {
	id *int
}

func main() {
	// var f1 bigStarFoo
	// fmt.Println(f1)
	// f1.Foo = &foo{}
	// fmt.Println(f1)
	// fmt.Printf("f1.Foo's id: %d, name: %s \n", f1.Foo.id, f1.Foo.name)
	// f1.Foo.id = 1
	// f1.Foo.name = "f1"
	// fmt.Println(f1)
	// fmt.Println(f1.Foo)
	// fmt.Println(*f1.Foo)
	// fmt.Println(&f1.Foo)
	// fmt.Printf("f1.Foo's id: %d, name: %s \n", f1.Foo.id, f1.Foo.name)

	// // pointer asign to object
	// var f2 bigFoo
	// fmt.Print("f2: ")
	// fmt.Println(f2)
	// f2.Foo = *f1.Foo
	// fmt.Print("f2: ")
	// fmt.Println(f2)

	// // object asign to pointer
	// var f3 bigStarFoo
	// fmt.Print("f3: ")
	// fmt.Println(f3)
	// f3.Foo = &f2.Foo
	// fmt.Print("f3: ")
	// fmt.Println(f3)
	// fmt.Printf("f3.Foo's id: %d, name: %s \n", f3.Foo.id, f3.Foo.name)

	// var Too too
	// id := 1
	// Too.id = &id
	// fmt.Println(*Too.id)
	var Foo foo
	fmt.Println(&Foo)
	fmt.Println(&Foo.id)
	fmt.Println(&Foo.name)
	var Too foo
	Too = Foo
	fmt.Println(&Too.id)
	fmt.Println(&Too.name)
}
