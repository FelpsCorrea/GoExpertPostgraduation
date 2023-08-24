package main

func main() {

	a := 1
	b := 2
	c := 3

	if a > b && c > a {
		println("a > b && c > a")
	}

	if b > a || c > a {
		println("b > a || c > a")
	}

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("default")
	}

}
