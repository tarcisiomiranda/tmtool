// cmd/root.go
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var number bool

var rootCmd = &cobra.Command{
	Use:   "leitor",
	Short: "Leitor de arquivos que ignora comentários e linhas em branco",
	Long: `Leitor de arquivos em Go que utiliza Cobra para gerenciar argumentos de linha de comando.
    Ele lê um arquivo especificado, ignorando linhas que começam com '#' e linhas em branco.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]

		// Tenta abrir o arquivo
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("Erro ao abrir o arquivo: %v", err)
		}
		defer file.Close()

		// Cria um scanner para ler o arquivo linha por linha
		scanner := bufio.NewScanner(file)
		lineNumber := 1
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())

			// Ignora linhas em branco
			if line == "" {
				continue
			}

			// Ignora linhas que começam com #
			if strings.HasPrefix(line, "#") {
				continue
			}

			// Processa a linha (neste exemplo, apenas imprime)
			if number {
				fmt.Printf("%d: %s\n", lineNumber, line)
			} else {
				fmt.Println(line)
			}

			lineNumber++
		}

		// Verifica se ocorreu algum erro durante a leitura
		if err := scanner.Err(); err != nil {
			log.Fatalf("Erro ao ler o arquivo: %v", err)
		}
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&number, "number", "n", false, "Exibir números das linhas")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
