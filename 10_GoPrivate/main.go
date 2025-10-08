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

	//------------------------------------------------------------------------------------------------
	// se quiser que o go mantenha os pacotes no projeto para nao correr o risco de
	// perder as dependencias, basta usar o comando:
	// go mod tidy -> remove as dependencias que nao sao mais usadas
	// go mod vendor -> copia os pacotes para o diretorio vendor
	// go mod download -> baixa os pacotes para o diretorio vendor
	// go mod verify -> verifica se os pacotes sao validos
	// go mod why -> mostra porque o pacote foi instalado

}
