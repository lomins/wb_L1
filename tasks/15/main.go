package main

// полная строка остается в памяти, и не происходит ее освобождение

func someFunc() string {
	v := createHugeString(1 << 10)
	return v[:100]
}

func main() {
	justString := someFunc()
}
