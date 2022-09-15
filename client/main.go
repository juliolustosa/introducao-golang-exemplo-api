package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatalf("Erro ao tentar acessar api server: %s", err)
	}

	// defer: comando executado apos fim da funcao
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	// retorna conteudo da resposta da api
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	// retorna erro caso tenha falhado a leitura do conteudo
	if err := scanner.Err(); err != nil {
		log.Fatalf("Erro ao ler conteudo da resposta: %s", err)
	}
}
