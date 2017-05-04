package main

func incrementPointer(input []byte) []byte {
	return append(input, []byte("p++; ")...)
}

func decrementPointer(input []byte) []byte {
	return append(input, []byte("p--; ")...)
}

func incrementValueAtPointer(input []byte) []byte {
	return append(input, []byte("(*p)++; ")...)
}

func decrementValueAtPointer(input []byte) []byte {
	return append(input, []byte("(*p)--; ")...)
}

func printValueAtPointer(input []byte) []byte {
	return append(input, []byte("putchar(*p); ")...)
}

func getAndStoreInputAtPointer(input []byte) []byte {
	return append(input, []byte("*p=getchar(); ")...)
}

func jumpForwardIfValueAtPointerIsZero(input []byte) []byte {
	return append(input, []byte("while (*p) { ")...)
}

func jumpBackwards(input []byte) []byte {
	return append(input, []byte(" } ")...)
}
