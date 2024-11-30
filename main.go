package main

func sum(x, y int) int {
	return x + y
}

func div(x, y int) (float32, error) {
	return div(x, y), nil
}

func main() {
	x := sum(1, 2)
	println("Hello", x, " World!")
}
