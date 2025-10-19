# 🚀 Deploy Go App no Kubernetes

> Deploy de aplicação Go no Kubernetes local usando Kind

## 📌 Sobre

Este projeto demonstra como fazer deploy de uma aplicação Go no Kubernetes usando Kind (Kubernetes in Docker) com imagens otimizadas via multi-stage builds.

## 🚀 Quick Start

### 1. Instalar Kind

```bash
go install sigs.k8s.io/kind@latest
```

### 2. Criar Cluster

```bash
kind create cluster --name=goexpert
```

### 3. Build e Deploy

```bash
# Build da imagem
docker build -f Dockerfile.prod -t goapp:latest .

# Carregar imagem no Kind
kind load docker-image goapp:latest --name=goexpert

# Deploy no Kubernetes
kubectl apply -f k8s/

# Testar aplicação
kubectl port-forward service/goapp-service 8080:8080
curl http://localhost:8080
```

## 🐳 Dockerfile.prod

```dockerfile
FROM golang:1.23.5 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server main.go

FROM scratch
COPY --from=builder /app/server .
CMD ["./server"]
```

**Por que multi-stage?**

- **Builder**: Compila a aplicação com Go toolchain
- **Runtime**: Apenas o binário (imagem ~10MB vs ~800MB)

## 🎯 Kind - Kubernetes in Docker

**Kind** permite executar clusters Kubernetes locais usando containers Docker.

### Comandos Básicos

```bash
# Instalar
go install sigs.k8s.io/kind@latest

# Criar cluster
kind create cluster --name=goexpert

# Listar clusters
kind get clusters

# Deletar cluster
kind delete cluster --name=goexpert

# Carregar imagem
kind load docker-image goapp:latest --name=goexpert
```

## 📦 Kubernetes - Comandos Úteis

### Pods

```bash
# Listar pods
kubectl get pods

# Ver detalhes
kubectl describe pod <pod-name>

# Logs
kubectl logs -l app=goapp

# Deletar pods
kubectl delete pods -l app=goapp
```

### Deployments

```bash
# Listar deployments
kubectl get deployments

# Escalar
kubectl scale deployment goapp-deployment --replicas=5

# Atualizar imagem
kubectl set image deployment/goapp-deployment goapp=goapp:v2.0.0
```

### Services

```bash
# Listar services
kubectl get services

# Port forward
kubectl port-forward service/goapp-service 8080:8080

# Expor via NodePort
kubectl expose deployment goapp-deployment --type=NodePort --port=8080
```

### Aplicar/Deletar Recursos

```bash
# Aplicar todos os arquivos
kubectl apply -f k8s/

# Deletar tudo
kubectl delete -f k8s/

# Aplicar arquivo específico
kubectl apply -f k8s/deployment.yaml
```

## 🔧 Configuração dos Recursos

### Deployment (`k8s/deployment.yaml`)

- **3 réplicas** da aplicação
- **Recursos limitados**: 32Mi RAM, 100m CPU
- **imagePullPolicy: Never** (para Kind)
- **Porta 8080** exposta

### Service (`k8s/service.yaml`)

- **NodePort** na porta 30080
- **LoadBalancer** para produção
- **ClusterIP** para comunicação interna

## 🚀 Build e Deploy Completo

```bash
# 1. Build da imagem
docker build -f Dockerfile.prod -t goapp:latest .

# 2. Criar cluster Kind
kind create cluster --name=goexpert

# 3. Carregar imagem
kind load docker-image goapp:latest --name=goexpert

# 4. Deploy no K8s
kubectl apply -f k8s/

# 5. Verificar status
kubectl get pods
kubectl get services

# 6. Testar aplicação
kubectl port-forward service/goapp-service 8080:8080
curl http://localhost:8080
```

## 📊 Monitoramento

```bash
# Status geral
kubectl get all

# Logs em tempo real
kubectl logs -f deployment/goapp-deployment

# Métricas de recursos
kubectl top pods
kubectl top nodes

# Eventos
kubectl get events --sort-by=.metadata.creationTimestamp
```

## 🎯 Resumo

- ✅ **Kind**: Kubernetes local para desenvolvimento
- ✅ **Multi-stage build**: Imagem otimizada (~10MB)
- ✅ **Deployment**: 3 réplicas com recursos limitados
- ✅ **Service**: Exposição via NodePort/LoadBalancer
- ✅ **Port-forward**: Acesso local para testes
