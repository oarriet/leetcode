package utils

import "fmt"

func PrintMatrixOfBytes(byteMatrix [][]byte) {
	println("----- Printing Matrix of Byte")
	for i := 0; i < len(byteMatrix); i++ {
		for j := 0; j < len(byteMatrix[i]); j++ {
			print(byteMatrix[i][j])
		}
		println("")
	}
	println("-----")
}

func PrintMatrixOfInts(intMatrix [][]int) {
	println("----- Printing Matrix of Int")
	for i := 0; i < len(intMatrix); i++ {
		for j := 0; j < len(intMatrix[i]); j++ {
			fmt.Printf(" %d ", intMatrix[i][j])
		}
		println("")
	}
	println("-----")
}
