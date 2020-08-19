# fct7-devops-kubernetes-hpa-desafio
Code Education | Full Cycle Turma 7 | DevOps | Kubernetes &amp; HPA - Projeto Prático

Repositório Docker Hub:  
<https://hub.docker.com/r/kaissi/go-hpa>

Fazer o build da imagem Docker:  
`docker build ./src -t kaissi/go-hpa:latest`

## 1. Executar standalone

Testar localmente subindo o servidor:  
`docker run --rm --name go-hpa --publish 8001:8001 kaissi/go-hpa:latest`  

Executar um dos dois comandos (fora do container):  
`wget -qO- http://localhost:8001`  
ou  
`curl -sf http://localhost:8001`  

Resultado:
> Code.education Rocks

## 2. Executar no Kubernetes

Criar um Pod loader e entrar no mesmo:  
`kubectl run --rm -it loader --image=busybox --generator=run-pod/v1 /bin/sh`  

Executar um dos seguintes comandos (dentro do container loader):  
`wget -qO- http://go-hpa:8001`  
ou  
`wget -qO- http://go-hpa.demo.svc.cluster.local:8001`  

Resultado:  
> Code.education Rocks

Para gerar a carga, executar em looping. Isso irá sobrecarregar o servidor, fazendo com que o HPA funcione:  
`while true; do wget -qO- http://go-hpa:8001; done`

