package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"go.bug.st/serial.v1"
)

func sendSerialData(port serial.Port, data string) error {
	_, err := port.Write([]byte(data))
	return err
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Define the dimensions of the matrix
	const rows = 5
	const columns = 2

	// Create and populate the matrix
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, columns)
		for j := range matrix[i] {
			matrix[i][j] = 4294967295 //rand.Intn(100) // random numbers between 0 and 99
		}
	}

	// Print the matrix
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%d\t", val)
		}
		fmt.Println()
	}

	// Open the first serial port detected at 9600 baud
	mode := &serial.Mode{
		BaudRate: 9600,
	}
	port, err := serial.Open("/dev/ttyACM0", mode)
	if err != nil {
		panic(err)
	}
	defer port.Close()

	fmt.Printf("\n")

	// Send the number of rows first
	fmt.Printf("Envío de datos: \n")
	sendSerialData(port, strconv.Itoa(rows)+"\n")
	fmt.Printf("\n")
	fmt.Printf(strconv.Itoa(rows) + "\n")

	fmt.Printf("\n")

	// Send each row
	for _, row := range matrix {
		rowStr := " "
		for _, val := range row {
			rowStr += fmt.Sprintf("%d\t", val)
		}
		sendSerialData(port, rowStr+"\n")
		fmt.Printf(rowStr + "\n")
	}
}
