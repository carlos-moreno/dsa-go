// Package chapter01 contém exemplos de operações em estruturas de dados lineares.
package chapter01

import "fmt"

// Run executa a demonstração de uma tabela linear simples,
// imprimindo os elementos de uma fatia (slice) de inteiros.
func RunLinearTable() {
	// 1. Definir uma matriz unidimensional (slice) de pontuações.
	// Em Go, slices são usados para representar sequências dinâmicas.
	var scores = []int{90, 70, 50, 80, 60, 85}

	// 2. Iterar sobre a tabela linear usando 'range'.
	// O range nos dá o índice (que ignoramos com '_') e o valor de cada elemento.
	// Isso é mais seguro e limpo do que controlar um contador manualmente.
	for _, score := range scores {
		// 3. Imprimir cada elemento (score) formatado com um espaço
		fmt.Printf("%d ", score)
	}

	// 4. Imprimir uma nova linha ao final para organizar a saída
	fmt.Println()
}
