// Script que retorna a quantidade necessária de notas para atingir um valor decimal entrado pelo usuário.
package main

// Importação dos pacotes.
import (
	"fmt"
	"strings"
)

// Objeto que representa a quantidade de notas acumuladas de um valor.
type Notas struct {
	// Quantidade de notas de R$200,00.
	Notas200 int

	// Quantidade de notas de R$100,00.
	Notas100 int

	// Quantidade de notas de R$50,00.
	Notas50 int

	// Quantidade de notas de R$20,00.
	Notas20 int

	// Quantidade de notas de R$10,00.
	Notas10 int

	// Quantidade de notas de R$5,00.
	Notas5 int

	// Quantidade de notas de R$2,00.
	Notas2 int

	// Valor total da sobra que impossível de ser retornada utilizando apenas as notas.
	Sobra float64
}

// Função que formata a mensagem de impressão para quantidade de notas acumuladas.
func (notas Notas) Mensagem() string {
	// Mensagem para impressão.
	var mensagem string

	// Verificação se existem notas de R$200,00 para formatar na mensagem de impressão.
	if notas.Notas200 > 0 {
		mensagem = fmt.Sprintf("Notas de R$200,00: %d\n", notas.Notas200)
	}

	// Verificação se existem notas de R$100,00 para formatar na mensagem de impressão.
	if notas.Notas100 > 0 {
		mensagem = mensagem + fmt.Sprintf("Notas de R$100,00: %d\n", notas.Notas100)
	}

	// Verificação se existem notas de R$50,00 para formatar na mensagem de impressão.
	if notas.Notas50 > 0 {
		mensagem = mensagem + fmt.Sprintf("Notas de R$50,00: %d\n", notas.Notas50)
	}

	// Verificação se existem notas de R$20,00 para formatar na mensagem de impressão.
	if notas.Notas20 > 0 {
		mensagem = mensagem + fmt.Sprintf("Notas de R$20,00: %d\n", notas.Notas20)
	}

	// Verificação se existem notas de R$10,00 para formatar na mensagem de impressão.
	if notas.Notas10 > 0 {
		mensagem = mensagem + fmt.Sprintf("Notas de R$10,00: %d\n", notas.Notas10)
	}

	// Verificação se existem notas de R$5,00 para formatar na mensagem de impressão.
	if notas.Notas5 > 0 {
		mensagem = mensagem + fmt.Sprintf("Notas de R$5,00: %d\n", notas.Notas5)
	}

	// Verificação se existem notas de R$2,00 para formatar na mensagem de impressão.
	if notas.Notas2 > 0 {
		mensagem = mensagem + fmt.Sprintf("Notas de R$2,00: %d\n", notas.Notas2)
	}

	// Verificação se existem valor de sobra para formatar na mensagem de impressão.
	if notas.Sobra > 0 {
		mensagem = mensagem + fmt.Sprintf("Total do valor impossível de obter utilizando apenas notas: R$%f\n", notas.Sobra)
	}

	// Validação se a mensagem de impressão foi preenchida.
	if mensagem == "" {
		mensagem = "O Valor está incorreto."
	} else {
		// Remoção do último caractere especial para não existir uma linha vazia.
		mensagem = strings.TrimSuffix(mensagem, "\n")
	}

	/* Função de substituição dos caracteres especiais para implantação de uma nova linha.
	E retorno da mensagem para a impressão já formatada.*/
	return strings.Replace(mensagem, `\n`, "\n", -1)
}

// Calcula a quantidade de notas necessárias para atingir um valor decimal.
func calcular_notas(input float64) Notas {
	// Variável utilizada na função que contabiliza a quantidade de notas de R$200.
	var notas200 int
	// Variável utilizada na função que contabiliza a quantidade de notas de R$100.
	var notas100 int
	// Variável utilizada na função que contabiliza a quantidade de notas de R$50.
	var notas50 int
	// Variável utilizada na função que contabiliza a quantidade de notas de R$20.
	var notas20 int
	// Variável utilizada na função que contabiliza a quantidade de notas de R$10.
	var notas10 int
	// Variável utilizada na função que contabiliza a quantidade de notas de R$5.
	var notas5 int
	// Variável utilizada na função que contabiliza a quantidade de notas de R$2.
	var notas2 int
	// Variável utilizada na função que contabiliza o total da sobra impossível de retornar utilizando apenas notas.
	var sobra float64

	// Variável para o controle do loop.
	breakLoop := false

	// Loop para calcular a quantidade de notas do valor entrado.
	for {
		// Atribuição do valor atual da entrada na variável do switch-case.
		switch valor := input; {

		/* Valida se o valor atual é maior ou igual a 200.
		Se sim, remove 200 do valor da entrada e acresenta a variável de quantidade de notas de R$200,00.*/
		case valor >= 200.00:
			input -= 200
			notas200++

		/* Valida se o valor atual é maior ou igual a 100.
		Se sim, remove 100 do valor da entrada e acresenta a variável de quantidade de notas de R$100,00.*/
		case valor >= 100.00:
			input -= 100
			notas100++

		/* Valida se o valor atual é maior ou igual a 50.
		Se sim, remove 50 do valor da entrada e acresenta a variável de quantidade de notas de R$50,00.*/
		case valor >= 50.00:
			input -= 50
			notas50++

		/* Valida se o valor atual é maior ou igual a 20.
		Se sim, remove 200 do valor da entrada e acresenta a variável de quantidade de notas de R$20,00.*/
		case valor >= 20.00:
			input -= 20
			notas20++

		/* Valida se o valor atual é maior ou igual a 10.
		Se sim, remove 10 do valor da entrada e acresenta a variável de quantidade de notas de R$10,00.*/
		case valor >= 10.00:
			input -= 10
			notas10++

		/* Valida se o valor atual é maior ou igual a 5.
		Se sim, remove 5 do valor da entrada e acresenta a variável de quantidade de notas de R$5,00.*/
		case valor >= 5.00:
			input -= 5
			notas5++

		/* Valida se o valor atual é maior ou igual a 2.
		Se sim, remove 2 do valor da entrada e acresenta a variável de quantidade de notas de R$2,00.*/
		case valor >= 2.00:
			input -= 2
			notas2++

		/* Acresenta o restante do valor de entrada atual na váriavel de sobra
		E altera o estado da variável de controle do loop.*/
		default:
			sobra = input
			breakLoop = true
		}

		/* Valida se o estado da variável de controle do loop foi alterado.
		Se sim, é realizdo a quebra do loop.*/
		if breakLoop {
			break
		}
	}

	// Criação e retorno do objeto da quantidade de notas.
	return Notas{
		Notas200: notas200,
		Notas100: notas100,
		Notas50:  notas50,
		Notas20:  notas20,
		Notas10:  notas10,
		Notas5:   notas5,
		Notas2:   notas2,
		Sobra:    sobra,
	}
}

// Função principal do script.
func main() {
	// Variável de armazenamento do valor decimal da entrada.
	var input float64

	// Recupera o valor decimal da entrada.
	fmt.Print("Entre com o valor decimal: R$")
	fmt.Scan(&input)

	// Impressão da mensagem que reporta a quantidade de notas necessárias para atingir o valor de entrada.
	fmt.Print(calcular_notas(input).Mensagem())
}
