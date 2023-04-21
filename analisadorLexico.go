package main

import (
	"bufio"
	"fmt"
	"os"
)

// Estrutura para representar um token
type Token struct {
	TokenType  string // tipo do token
	Lexeme     string // lexema do token
	LineNumber int    // número da linha
	ColNumber  int    // número da coluna
}

func main() {

	// Abre o arquivo de entrada
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Cria um scanner para ler o arquivo
	scanner := bufio.NewScanner(file)

	// Inicializa a tabela de símbolos
	symbolsTable := make(map[string]Token)

	// Loop através de cada linha do arquivo
	var lineNumber int
	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		// Loop através de cada caractere na linha
		var colNumber int
		for i := 0; i < len(line); i++ {
			colNumber++
			char := line[i]

			// Verifica se o caractere é uma letra maiúscula ou minúscula
			if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
				// Se for uma letra, inicia um novo token
				token := Token{TokenType: "Identificador", Lexeme: "", LineNumber: lineNumber, ColNumber: colNumber}
				token.Lexeme += string(char)

				// Loop através dos próximos caracteres até que não seja mais uma letra ou um número
				for j := i + 1; j < len(line); j++ {
					nextChar := line[j]
					if (nextChar >= 'a' && nextChar <= 'z') || (nextChar >= 'A' && nextChar <= 'Z') || (nextChar >= '0' && nextChar <= '9') {
						token.Lexeme += string(nextChar)
					} else {
						break
					}
				}

				// Adiciona o token à tabela de símbolos, se ainda não estiver lá
				if _, ok := symbolsTable[token.Lexeme]; !ok {
					symbolsTable[token.Lexeme] = token
				}

				// Imprime o token
				fmt.Println(token)

				// Move o índice i para o final do token para que o loop externo possa continuar a partir do próximo caractere
				i += len(token.Lexeme) - 1
			}

			// Verifica se o caractere é um número
			if char >= '0' && char <= '9' {
				// Se for um número, inicia um novo token
				token := Token{TokenType: "Número", Lexeme: "", LineNumber: lineNumber, ColNumber: colNumber}
				token.Lexeme += string(char)

				// Loop através dos próximos caracteres até que não seja mais um número
				for j := i + 1; j < len(line); j++ {
					nextChar := line[j]
					if nextChar >= '0' && nextChar <= '9' {
						token.Lexeme += string(nextChar)
					} else {
						break
					}
				}

				// Imprime o token
				fmt.Println(token)

				// Move o índice i para o final do token para que o loop externo possa continuar a partir do próximo caractere
				i += len(token.Lexeme) - 1
			}

			// Verifica se o caractere é um espaço em branco
			if char == ' ' {
				continue
			}

			// Verifica se o caractere é um símbolo
			if char == '+' || char == '-' || char == '*' || char == '/' {
				// Se for um símbolo, inicia um novo token
				token := Token{TokenType: "Símbolo", Lexeme: string(char), LineNumber: lineNumber, ColNumber: colNumber}

				// Adiciona o token à tabela de símbolos, se ainda não estiver lá
				if _, ok := symbolsTable[token.Lexeme]; !ok {
					symbolsTable[token.Lexeme] = token
				}

				// Imprime o token
				fmt.Println(token)
			}

			// Verifica se o caractere é um parêntese de abertura
			if char == '(' {
				// Se for um parêntese de abertura, inicia um novo token
				token := Token{TokenType: "Parêntese", Lexeme: string(char), LineNumber: lineNumber, ColNumber: colNumber}

				// Adiciona o token à tabela de símbolos, se ainda não estiver lá
				if _, ok := symbolsTable[token.Lexeme]; !ok {
					symbolsTable[token.Lexeme] = token
				}

				// Imprime o token
				fmt.Println(token)
			}

			// Verifica se o caractere é um parêntese de fechamento
			if char == ')' {
				// Se for um parêntese de fechamento, inicia um novo token
				token := Token{TokenType: "Parêntese", Lexeme: string(char), LineNumber: lineNumber, ColNumber: colNumber}

				// Adiciona o token à tabela de símbolos, se ainda não estiver lá
				if _, ok := symbolsTable[token.Lexeme]; !ok {
					symbolsTable[token.Lexeme] = token
				}

				// Imprime o token
				fmt.Println(token)
			}

			// Verifica se o caractere é um ponto e vírgula
			if char == ';' {
				// Se for um ponto e vírgula, inicia um novo token
				token := Token{TokenType: "Ponto e vírgula", Lexeme: string(char), LineNumber: lineNumber, ColNumber: colNumber}

				// Adiciona o token à tabela de símbolos, se ainda não estiver lá
				if _, ok := symbolsTable[token.Lexeme]; !ok {
					symbolsTable[token.Lexeme] = token
				}

				// Imprime o token
				fmt.Println(token, colNumber, lineNumber)
			}
		}
	}

	fmt.Println("Tabela de símbolos:")
	fmt.Printf("%-20s %-20s %-20s %-20s %-20s\n", "Lexema", "Tipo de token", "Valor do token", "Linha", "Coluna")
	fmt.Println("----------------------------------------------------------------------------------------------")

	for lexeme, token := range symbolsTable {
		fmt.Printf("%-20s %-20s %-20s %-20d %-20d\n", lexeme, token.TokenType, token.Lexeme, token.LineNumber, token.ColNumber)
	}
}
