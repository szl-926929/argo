package main

import "os"

func main() {
	switch os.Args[1] {
	case "docgen":
		generateDocs()
	case "readmegen":
		readmeGen()
	default:
		panic(os.Args[1])
	}
}
