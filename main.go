package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Pedir el número de canales
	fmt.Print("Ingrese el número de canales (máximo 10): ")
	canalesStr, _ := reader.ReadString('\n')
	canalesStr = strings.TrimSpace(canalesStr)
	canales, err := strconv.Atoi(canalesStr)
	if err != nil || canales < 1 || canales > 10 {
		fmt.Println("Número de canales inválido. Debe ser un número entre 1 y 6.")
		return
	}

	// Pedir el número de filas
	fmt.Print("Ingrese el número de filas: ")
	filasStr, _ := reader.ReadString('\n')
	filasStr = strings.TrimSpace(filasStr)
	filas, err := strconv.Atoi(filasStr)
	if err != nil || filas < 1 {
		fmt.Println("Número de filas inválido. Debe ser un número positivo.")
		return
	}

	// Crear una matriz de ceros
	columnas := 3*canales + 1 // 3 columnas por canal + 1 columna para T-s
	matriz := make([][]int, filas)
	for i := range matriz {
		matriz[i] = make([]int, columnas)
	}

	// Nueva matriz para almacenar la conversión decimal y T-s
	matrizDecimal := make([][]int, filas)
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
		fmt.Print("Canales   xx: |")
		for i := 1; i <= canales; i++ {
			fmt.Printf("- Canal %d -|", i)
		}
		fmt.Println("T-s |")
		fmt.Print("Columnas  xx: |")
		for i := 1; i <= canales*3+1; i++ {

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

		if err1 != nil || err2 != nil || err3 != nil || fila < 1 || fila > filas || columna < 1 || columna > columnas {
			fmt.Println("Entrada inválida. Asegúrese de que fila y columna estén dentro de los límites de la matriz y que el valor sea un número.")
			continue
		}

		// Validación para las columnas de canales y T-s
		if columna < columnas {
			if valor != 0 && valor != 1 {
				fmt.Println("Valor inválido para las columnas de canales. Solo se aceptan 0 o 1.")
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
			matrizDecimal[i][0] = convertirBinarioADecimal(filaMatriz[:3*canales])
			matrizDecimal[i][1] = filaMatriz[columnas-1]
		}
	}

	imprimirMatrices()
	fmt.Println("Programa terminado.")
}
