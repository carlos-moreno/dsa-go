// Package chapter01 contém exemplos de operações em estruturas de dados lineares.
package chapter01

import "fmt"

// add é uma função interna que adiciona um valor ao final de um slice.
// Nota: Usamos 'add' para evitar conflito (shadowing) com a função nativa 'append'.
func add(array []int, value int) []int {
	// 1. Primeiro, crie uma matriz temporária (tempArray) maior que o comprimento da matriz original
	var length = len(array)
	var tempArray = make([]int, length+1)

	// 2. Copie cada valor da matriz para tempArray
	// A função copy(dest, src) é a forma mais eficiente e idiomática no Go.
	copy(tempArray, array)

	// 3. Atribua o valor ao último índice de tempArray
	tempArray[length] = value

	// 4. Finalmente, retorne a referência do tempArray (que será atribuída à matriz original)
	return tempArray
}

// RunAppendLinearTable demonstra o processo manual de adição de um item em uma tabela linear.
func RunAppendLinearTable() {
	// Definição inicial da tabela linear
	var scores = []int{90, 70, 50, 80, 60, 85}

	// Executa a função interna para adicionar o valor 75
	scores = add(scores, 75)

	// Imprime os resultados usando o padrão moderno de range
	for _, score := range scores {
		fmt.Printf("%d ", score)
	}
	fmt.Println()
}
