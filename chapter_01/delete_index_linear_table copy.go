// Package chapter01 contém exemplos de operações em estruturas de dados lineares.
package chapter01

import "fmt"

// remove é uma função interna que exclui um elemento de um índice específico,
// deslocando os elementos à direita para a esquerda para preencher o espaço.
func remove(array []int, index int) []int {
	// 1. Crie uma matriz temporária (tempArray) com o tamanho da original menos um
	var length = len(array)
	var tempArray = make([]int, length-1)

	// 2. Copie os valores da matriz original para o tempArray até o índice de exclusão
	// Usamos o slice [:index] para pegar todos os elementos antes do alvo
	copy(tempArray[:index], array[:index])

	// 3. Copie os valores restantes após o índice para o tempArray,
	// deslocando-os uma posição para a esquerda (pulando o índice excluído)
	// tempArray[index:] recebe os dados de array[index+1:]
	copy(tempArray[index:], array[index+1:])

	// 4. Finalmente, retorne a referência do tempArray (que será atribuída à matriz original)
	return tempArray
}

// RunDeleteIndexLinearTable demonstra a exclusão de um item em um índice específico.
func RunDeleteIndexLinearTable() {
	// Definição inicial da tabela linear
	var scores = []int{90, 70, 50, 80, 60, 85}

	// Executa a função interna para remover o elemento no índice 2 (valor 50)
	scores = remove(scores, 2)

	// Imprime os resultados usando o padrão moderno de range
	for _, score := range scores {
		fmt.Printf("%d ", score)
	}
	fmt.Println()
}
