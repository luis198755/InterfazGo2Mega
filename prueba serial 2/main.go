package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tarm/serial"
)

func main() {
	matrix := [][]uint32{
		{4294967295, 10}, // Row 0
		{4294967294, 20}, // Row 1
		{4294967293, 30}, // Row 2
		{4294967292, 40}, // Row 3
	}

	// Step 1: Generate a random 5x2 matrix
	//rand.Seed(time.Now().UnixNano())
	//matrix := generateRandomMatrix(4, 2)

	// Step 2: Convert matrix to string format
	matrixStr := matrixToString(matrix)

	// Print the matrix to console
	printMatrix(matrixStr)

	// Step 3: Send through serial port
	sendThroughSerial(matrixStr)
}

// Generates a random 5x2 matrix with uint32 values
/*
func generateRandomMatrix(rows, cols int) [][]uint32 {
	matrix := make([][]uint32, rows)
	for i := range matrix {
		matrix[i] = make([]uint32, cols)
		for j := range matrix[i] {
			matrix[i][j] = rand.Uint32()
		}
	}
	return matrix
}*/

// Converts the matrix to the specified string format
func matrixToString(matrix [][]uint32) string {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(len(matrix)) + "\n")
	for _, row := range matrix {
		for j, val := range row {
			sb.WriteString(strconv.FormatUint(uint64(val), 10))
			if j < len(row)-1 {
				sb.WriteString(",")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

// Sends the matrix string through a serial port
func sendThroughSerial(matrixStr string) {
	// Configure the serial port
	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600} // Replace "COM1" with your port
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println("Error opening serial port:", err)
		return
	}
	defer s.Close()

	// Write the matrix string to the serial port
	_, err = s.Write([]byte(matrixStr))
	if err != nil {
		fmt.Println("Error writing to serial port:", err)
		return
	}
	fmt.Println("Matrix sent successfully.")
}

// Prints the matrix string to the console
func printMatrix(matrixStr string) {
	fmt.Println("Matrix to be sent:")
	fmt.Println(matrixStr)
}
