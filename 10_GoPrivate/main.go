package main

import (
	"fmt"

	"github.com/ElizCarvalho/fcutils/pkg/events"
)

func main() {
	fmt.Println("Hello, World!")
	ed := events.NewEventDispatcher()
	fmt.Println(ed)

	//existe uma variavel do go chamada GOPRIVATE
	//ela é usada para definir o caminho dos repositorios privados
	//exemplo: GOPRIVATE=github.com/ElizCarvalho/fcutils
	//ou
	//GOPRIVATE=github.com/ElizCarvalho/*
	//ou
	//GOPRIVATE=github.com/ElizCarvalho/fcutils,github.com/ElizCarvalho/fcutils/*

	//caso tenha que autenticar no repositorio privado, basta usar o comando:

	//SE QUISER USAR NETRC:
	// criar um arquivo .netrc na raiz do projeto
	// e colocar o seguinte conteudo:
	// machine github.com
	// login <seu_usuario>
	// password <seu_token_do_github> settings -> developer settings -> personal access tokens -> new token -> generate token -> full control of private repositories

	//SE QUISER USAR SSH:
	//via git config
	//vim .git/config
	//[url "ssh://git@github.com/"]
	//	insteadOf = https://github.com/

	//------------------------------------------------------------------------------------------------
	// SE QUISER TRABALHAR COM O BITBUCKET:
	// criar um arquivo .netrc na raiz do projeto
	// e colocar o seguinte conteudo:
	// machine api.bitbucket.org
	// login <seu_usuario>
	// password <seu_token_do_bitbucket>

}
