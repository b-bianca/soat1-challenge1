# Software Architecture - Tech Challenge

Veja a Wiki do projeto em: [https://github.com/b-bianca/soat1-challenge1/wiki](https://github.com/b-bianca/soat1-challenge1/wiki)

<details>

<summary>Entrega FASE 3 - Devops | AWS e Terraform</summary>

## Requisitos

|Recurso|Versão|Obrigatório|Nota|
|-|-|-|-|
|Terraform| 1.6.2|Não|Necessário apenas no caso de rodar localmente|

## Como executar o projeto localmente
Garanta que os requisitos obrigatórios estejam instalados. Após isso siga:

### Etapa 1: Inicialize o Terraform
Inicie o terraform com o seguinte comando:
~~~bash
terraform init
~~~

### Etapa 2: Executar as ações no manifesto do Terraform
Na pasta onde se encontra o arquivo iniciador, execute:
~~~bash
terraform apply
~~~

>Nota: caso já exista uma infraestrutura montada, é possível rodar o comando `terraform plan` para verificar a infra remota já existente e o que as mudanças propostas irá alterar

---
---
</details>

<details>

<summary>Entrega FASE 2 - Kubernetes & Clean Architecture</summary>

# Software Architecture - FASE 2 - Tech Challenge

## Requisitos

|Recurso|Versão|Obrigatório|Nota|
|-|-|-|-|
|kubectl|1.27 ou atual|Sim|Necessário para interagir com o cluster Kubernetes que você executa localmente usando o Minikube ou o Docker Desktop Kubernetes|
|Docker Desktop| 4.22 ou atual|Sim*|Necessário para executar ambiente Kubernetes localmente. Deve ser habilitada a opção para utilizar kubernetes|
|minikube|1.31 ou atual|Sim*| Necessário para executar ambiente Kubernetes localmente.|
|Golang| 1.20|Não|Necessário apenas no caso de rodar localmente sem container|
|Kubernetes (extensão VSCode)|atual|Não|Ajuda visualizar em árvore os recursos e objetos kubernetes, possibilitando inspecionar e comandos com menu por mouse|

(*) Ao menos uma das duas ferramentas (minikube ou Docker Desktop) para executar localmente ambiente Kubernetes.

## Como executar o projeto
Garanta que os requisitos obrigatórios estejam instalados. Após isso siga:

### Etapa 1: Ative o ambiente para Kubernetes
###### Opção 1 | Utilizando Docker-Desktop
Ative o recurso Kubernetes nas configurações, conforme instrução [aqui](https://docs.docker.com/desktop/kubernetes/).

###### Opção 2 | Utilizando Minikube
Inicie o minikube com o seguinte comando:
~~~bash
minikube start
~~~


### Etapa 2: Apliques os arquivos de configuração
Na pasta raiz do projeto, execute:
~~~bash
kubectl apply -f k8s
~~~

>Nota: o comando aplicará todos os arquivos `.yaml` de configuração contidos no diretório `k8s`. Caso deseje executar os arquivos individualmente, siga nessa ordem: 1º `*-secrets.ymal`, 2º `*-configmap.ymal`, 3º `*-psc.ymal`, 4º `*-deployament.ymal`, 5º `*-service.ymal`


Com isso a aplicação estará pronta para consumo.

### Etapa 3: Consuma a API
Os endpoints disponíveis constam na documentação [aqui](#collection-do-postman)

O loadbalancer disponibilizará API para ser consumida na seguinte endereço:
>localhost:8080

>exemplo:
http://localhost:8080/api/v1/categories

Se por ventura não estiver disponível, tente conectar a porta no service kubernetes com a porta local:
~~~bash
kubectl port-forward service/restaurant-api-service 8080:8080 -n default
~~~


## Demonstração rodando Kubernetes localmente com Docker-Desktop
https://github.com/b-bianca/soat1-challenge1/assets/83218983/8cba5dfb-dee9-48d1-b0ed-57098d742101


---
---
</details>

<details>
<summary>Entrega FASE 1 - Docker & DDD</summary>

# Software Architecture - FASE 1 - Tech Challenge

## Requisitos Mínimos
* Docker Desktop | última versão
   

## Como executar o projeto

### Arquivo de Variáveis de Ambiente
Para executar o projeto, crie um arquivo '.env' no diretório raiz do projeto adicionando os valores conforme desejado. Abaixo um exemplo de um arquivo .env com dados de acesso localmente
~~~
API_CONTAINER_PORT='8080'
API_HOST_PORT='8080'

MYSQL_ROOT_PASSWORD='root'
MYSQL_USR='user'
MYSQL_PASS='user'
MYSQL_HOST='database-mysql'
MYSQL_DBNAME='restaurant'
MYSQL_CONTAINER_PORT='3306'
MYSQL_HOST_PORT='3306'
~~~


Para executar o projeto, é necessário ter o `Docker Desktop` instalado e subir as instancias usando o docker compose via IDE ou linha de comando conforme a seguir:
~~~bash
docker compose -f "docker-compose.yml" up -d --build
~~~

## Demonstração Rodando Docker Compose e Consumindo API
https://github.com/b-bianca/soat1-challenge1/assets/83218983/865c92df-56b0-4cc7-b328-3921039d1f9c

</details>

## Collection do Postman

O diretório `./docs` contém a collection da API implementada nesse projeto. Para utilizar, basta importar no Postman.<br>
Também é possível acessar o Postam Online:<br>
[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/16227218-ad366006-d6e5-41a8-8b14-0e5b79002ac0?action=collection%2Ffork&collection-url=entityId%3D16227218-ad366006-d6e5-41a8-8b14-0e5b79002ac0%26entityType%3Dcollection%26workspaceId%3De76668fb-982b-4d15-ab75-26131dab7174#?env%5BDEV%5D=W3sia2V5IjoiYmFzZV91cmwucmVzdGF1cmFudCIsInZhbHVlIjoibG9jYWxob3N0OjgwODAvYXBpL3YxIiwiZW5hYmxlZCI6dHJ1ZSwidHlwZSI6ImRlZmF1bHQifV0=)
