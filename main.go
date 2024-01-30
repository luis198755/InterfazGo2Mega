package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jacobsa/go-serial/serial"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var fases, escenarios, ciclos int
	var input string

	/*
		// Pedir el número de fases
		fmt.Print("Ingrese el número de fases (máximo 10): ")
		fasesStr, _ := reader.ReadString('\n')
		fasesStr = strings.TrimSpace(fasesStr)
		fases, err := strconv.Atoi(fasesStr)
		if err != nil || fases < 1 || fases > 10 {
			fmt.Println("Número de fases inválido. Debe ser un número entre 1 y 10.")
			return
		}

		// Pedir el número de escenarios
		fmt.Print("Ingrese el número de escenarios: ")
		escenariosStr, _ := reader.ReadString('\n')
		escenariosStr = strings.TrimSpace(escenariosStr)
		escenarios, err := strconv.Atoi(escenariosStr)
		if err != nil || escenarios < 1 {
			fmt.Println("Número de escenarios inválido. Debe ser un número positivo.")
			return
		}
	*/
	// Pedir el número de fases
	for {
		fmt.Print("Ingrese el número de fases (entre 1 y 6): ")
		_, err := fmt.Scan(&input)

		// Elimina espacios en blanco al inicio y al final
		input = strings.TrimSpace(input)
		// Intenta convertir la entrada en un número entero
		fases, err = strconv.Atoi(input)

		if err != nil || fases <= 0 {
			fmt.Println("Por favor, ingrese un número entero positivo.")
			continue
		}

		if fases >= 1 && fases <= 6 {
			break
		} else {
			fmt.Println("El número debe estar entre 1 y 6. Inténtelo de nuevo.")
		}
	}
	fmt.Printf("Número de fases seleccionado: %d\n", fases)

	// Pedir el número de escenarios
	for {
		fmt.Print("Ingrese el número de escenarios (entre 1 y 8): ")
		_, err := fmt.Scan(&input)

		// Elimina espacios en blanco al inicio y al final
		input = strings.TrimSpace(input)
		// Intenta convertir la entrada en un número entero
		escenarios, err = strconv.Atoi(input)

		if err != nil || escenarios <= 0 {
			fmt.Println("Por favor, ingrese un número entero positivo.")
			continue
		}

		if fases >= 1 && escenarios <= 8 {
			break
		} else {
			fmt.Println("El número debe estar entre 1 y 8. Inténtelo de nuevo.")
		}
	}
	fmt.Printf("Número de escenarios seleccionado: %d\n", fases)

	// Pedir el número de Ciclos
	for {
		fmt.Print("Ingrese el número de Ciclos (entre 1 y 8): ")
		_, err := fmt.Scan(&input)

		// Elimina espacios en blanco al inicio y al final
		input = strings.TrimSpace(input)
		// Intenta convertir la entrada en un número entero
		ciclos, err = strconv.Atoi(input)

		if err != nil || ciclos <= 0 {
			fmt.Println("Por favor, ingrese un número entero positivo.")
			continue
		}

		if ciclos >= 1 && ciclos <= 6 {
			break
		} else {
			fmt.Println("El número debe estar entre 1 y 6. Inténtelo de nuevo.")
		}
	}
	fmt.Printf("Número de Ciclos seleccionado: %d\n", ciclos)

	// Crear una matriz de ceros
	columnas := 3*fases + 1 // 3 columnas por Fase + 1 columna para T-s
	matriz := make([][]int, escenarios)
	for i := range matriz {
		matriz[i] = make([]int, columnas)
	}

	// Nueva matriz para almacenar la conversión decimal y T-s
	matrizDecimal := make([][]int, escenarios)
	for i := range matrizDecimal {
		matrizDecimal[i] = make([]int, 2) // Una columna para el valor decimal, otra para T-s
	}

	// Función para convertir binario a decimal
	convertirBinarioADecimal := func(binario []int) int {
		decimal := 0
		exponente := len(binario) - 1
		for _, bit := range binario {
			if bit == 1 {
				decimal += 1 << exponente
			}
			exponente--
		}
		return decimal
	}

	// Función para imprimir la matriz original y la matriz decimal
	imprimirMatrices := func() {
		fmt.Println("Matriz Original:")
		fmt.Print("Fases     xx: |")
		for i := 1; i <= fases; i++ {
			fmt.Printf("- Fase  %d -|", i)
		}
		fmt.Println("T-s |")
		fmt.Print("Columnas  xx: |")
		for i := 1; i <= fases*3+1; i++ {

			if i >= 10 {
				fmt.Printf("%d |", i)
			} else {
				fmt.Printf(" %d |", i)
			}
		}
		fmt.Println("")

		for i, escenario := range matriz {
			fmt.Printf("Escenario %02d: |", i+1)
			for _, valor := range escenario {
				fmt.Printf(" %d |", valor)
			}
			fmt.Println()
		}

		fmt.Println("\nMatriz Decimal:")
		for i, escenario := range matrizDecimal {
			fmt.Printf("Escenario %02d: | %d | T-s: %d |\n", i+1, escenario[0], escenario[1])
		}
	}

	// Bucle para modificar la matriz y actualizar la matriz decimal
	for {
		imprimirMatrices()
		fmt.Print("Ingrese 'fila columna valor' para modificar (o 'salir' para terminar): ")
		entrada, _ := reader.ReadString('\n')
		entrada = strings.TrimSpace(entrada)

		if entrada == "salir" {
			break
		}

		partes := strings.Split(entrada, " ")
		if len(partes) != 3 {
			fmt.Println("Entrada inválida. Asegúrese de ingresar tres números.")
			continue
		}

		fila, err1 := strconv.Atoi(partes[0])
		columna, err2 := strconv.Atoi(partes[1])
		valor, err3 := strconv.Atoi(partes[2])

		if err1 != nil || err2 != nil || err3 != nil || fila < 1 || fila > escenarios || columna < 1 || columna > columnas {
			fmt.Println("Entrada inválida. Asegúrese de que fila y columna estén dentro de los límites de la matriz y que el valor sea un número.")
			continue
		}

		// Validación para las columnas de fases y T-s
		if columna < columnas {
			if valor != 0 && valor != 1 {
				fmt.Println("Valor inválido para las columnas de fases. Solo se aceptan 0 o 1.")
				continue
			}
		} else {
			if valor < 0 || valor > 100 {
				fmt.Println("Valor inválido para la columna T-s. Solo se aceptan números entre 0 y 100.")
				continue
			}
		}

		matriz[fila-1][columna-1] = valor

		// Actualizar la matriz decimal
		for i, filaMatriz := range matriz {
			matrizDecimal[i][0] = convertirBinarioADecimal(filaMatriz[:3*fases])
			matrizDecimal[i][1] = filaMatriz[columnas-1]
		}
	}

	//imprimirMatrices()
	if err := sendMatrix("/dev/ttyACM0", matrizDecimal); err != nil {
		log.Fatalf("Failed to send matrix: %v", err)
	}

	fmt.Println("Programa terminado.")
}

func sendMatrix(portName string, matrix [][]int) error {
	// Serial port options
	options := serial.OpenOptions{
		PortName:        portName,
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	// Open the port
	port, err := serial.Open(options)
	if err != nil {
		return err
	}
	defer port.Close()

	// Send number of rows first
	numRows := len(matrix)
	if _, err := io.WriteString(port, fmt.Sprintf("%d\n", numRows)); err != nil {
		return err
	}

	// Send matrix data row by row
	for _, row := range matrix {
		rowStr := ""
		for _, value := range row {
			rowStr += fmt.Sprintf("%d ", value)
		}
		rowStr += "\n" // End of row

		if _, err := io.WriteString(port, rowStr); err != nil {
			return err
		}
	}

	return nil
}
