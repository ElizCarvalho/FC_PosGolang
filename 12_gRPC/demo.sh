#!/bin/bash

# Cores
BLUE='\033[0;34m'
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[0;33m'
NC='\033[0m'

echo -e "${BLUE}🧮 Demonstração do Serviço gRPC Calculadora${NC}"
echo "=============================================="
echo ""

# Verificar se o protoc está instalado
if ! command -v protoc &> /dev/null; then
    echo -e "${RED}❌ Protocol Buffers compiler (protoc) não encontrado${NC}"
    echo "Instale com: brew install protobuf"
    exit 1
fi

# Verificar se os plugins Go estão instalados
if ! command -v protoc-gen-go &> /dev/null; then
    echo -e "${RED}❌ protoc-gen-go não encontrado${NC}"
    echo "Instale com: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"
    exit 1
fi

if ! command -v protoc-gen-go-grpc &> /dev/null; then
    echo -e "${RED}❌ protoc-gen-go-grpc não encontrado${NC}"
    echo "Instale com: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"
    exit 1
fi

echo -e "${GREEN}✅ Todas as dependências estão instaladas${NC}"
echo ""

# Configurar o ambiente
echo -e "${BLUE}🔧 Configurando ambiente...${NC}"
make setup
echo ""

# Gerar código
echo -e "${BLUE}📝 Gerando código Go...${NC}"
make proto
echo ""

# Verificar código
echo -e "${BLUE}🔍 Verificando código...${NC}"
make lint
echo ""

echo -e "${GREEN}🎉 Ambiente configurado com sucesso!${NC}"
echo ""
echo -e "${YELLOW}Para testar o serviço:${NC}"
echo "1. Em um terminal: ${BLUE}make run-server${NC}"
echo "2. Em outro terminal: ${BLUE}make run-client${NC}"
echo ""
echo -e "${YELLOW}Comandos disponíveis:${NC}"
make help
