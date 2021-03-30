package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/yeka/zip"
)

const (
	zipPath             = "./assets/lerolero_protected.zip"
	caminhoArquivoSenha = "./assets/rockyou.txt"
	threadNumber        = 1
	linesPerThread      = 15000
)

func main() {
	abrirArquivoZip()
}

func bruteForce(zipPath string, listaDeSenhas []string, canal chan<- string) {
	//output := color.New(outputColor)
	arquivo, err := zip.OpenReader(zipPath)

	if err != nil {
		panic("Error !")
	}

	defer arquivo.Close()

	zipFile := arquivo.File[0]

	for _, value := range listaDeSenhas {
		//	output.Printf("Trying to crack the file with password: %v \n", string(value))

		zipFile.SetPassword(string(value))
		_, err := zipFile.Open()

		if err == nil {
			fmt.Printf("Aeeeee encontrei a senha\n")

			zipReader, err := zipFile.Open()

			if err != nil {
				panic("Erro")
			}

			buf, err := ioutil.ReadAll(zipReader)
			if err != nil {
				panic("Erro")
			}

			defer zipReader.Close()

			fmt.Printf("Size of %v: %v byte(s)\n", zipFile.Name, len(buf))

			canal <- string(value)
			break
		}
	}
}

func abrirArquivoZip() {
	arquivos, err := zip.OpenReader(zipPath)

	if err != nil {
		panic(fmt.Sprintf("Um erro aconteceu: %v", err))
	}

	defer arquivos.Close()

	zipFile := arquivos.File[0] //Pega o primeiro objeto do array
	if zipFile.IsEncrypted() {
		listaSenhas := obterListaDeSenhas(caminhoArquivoSenha)
		fmt.Printf("O arquivo esta protegido por senha. Bora quebrar issae... \n")

		canal := make(chan string, 1)

		start := time.Now()

		linhaInicial := 0
		for i := 0; i < threadNumber; i++ {
			linhaFinal := linesPerThread * (i + 1)

			//outputColor := RandomOutputColor(i)
			//output := color.New(outputColor)
			fmt.Printf("Iniciando a thread %d lendo da linha %d ate %d\n", i+1, linhaInicial, linhaFinal)
			go bruteForce(zipPath, listaSenhas[linhaInicial:linhaFinal], canal)

			linhaInicial = linhaFinal + 1
		}

		fmt.Println("----------------------------------------------------")

		fmt.Printf("Quebrando a senha...\n")
		fmt.Println()

		select {
		case senha := <-canal:
			fmt.Printf("\nA senha eh:\"%v\"\n", senha)
			fmt.Printf("A quebra da senha demorou: %v \n", time.Since(start))
		case <-time.After(time.Duration(15) * time.Second):
			//fmt.Printf("Timeout after: %d seconds \n", timeout)
			fmt.Printf("Senha nao encontrada :( \n")
		}

	}
}

func obterListaDeSenhas(caminhoArquivoSenha string) []string {
	arquivo, erro := os.Open(caminhoArquivoSenha)

	if erro != nil {
		panic(fmt.Sprintf("Um erro aconteceu: %v", erro))
	}

	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)

	scanner.Split(bufio.ScanLines)
	var senhas []string

	for scanner.Scan() {
		senhas = append(senhas, scanner.Text())
	}

	arquivoStatus, erro := os.Stat(caminhoArquivoSenha)

	if erro != nil {
		panic(fmt.Sprintf("Um erro aconteceu: %v", erro))
	}

	fmt.Printf("Status do arquivo => Nome: %v, Tamanho: %v KB \n", arquivoStatus.Name(), arquivoStatus.Size()/(1024))
	fmt.Printf("Quantidade total de senhas no arquivo: %d \n", len(senhas))

	return senhas
}
