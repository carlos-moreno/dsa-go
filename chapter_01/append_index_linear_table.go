// Package chapter01 contém exemplos de operações em estruturas de dados lineares.
package chapter01

import "fmt"

// insert é uma função interna que insere um valor em um índice específico,
// deslocando os elementos à direita para abrir espaço na tabela linear.
func insert(array []int, value int, insertIndex int) []int {
	// 1. Primeiro, crie uma matriz temporária (tempArray) maior que o comprimento da matriz original
	var length = len(array)
	var tempArray = make([]int, length+1)

	// 2. Copie os valores da matriz original para o tempArray até o índice de inserção
	// Usamos o slice [:insertIndex] para copiar apenas a parte inicial
	copy(tempArray[:insertIndex], array[:insertIndex])

	// 3. Atribua o novo valor ao índice especificado no tempArray
	tempArray[insertIndex] = value

	// 4. Copie os valores restantes da matriz original para o tempArray,
	// deslocando-os uma posição à direita (daqui em diante: index + 1)
	copy(tempArray[insertIndex+1:], array[insertIndex:])

	// 5. Finalmente, retorne a referência do tempArray (que será atribuída à matriz original)
	return tempArray
}

// RunInsertLinearTable demonstra a inserção de um item em um índice específico de uma tabela linear.
func RunInsertLinearTable() {
	// Definição inicial da tabela linear
	var scores = []int{90, 70, 50, 80, 60, 85}

	// Executa a função interna para inserir o valor 75 no índice 2
	scores = insert(scores, 75, 2)

	// Imprime os resultados usando o padrão moderno de range
	for _, score := range scores {
		fmt.Printf("%d ", score)
	}
	fmt.Println()
}
