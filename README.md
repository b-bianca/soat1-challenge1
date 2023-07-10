# Software Architecture - FASE 1 - Tech Challenge

Veja a Wiki do projeto em: [https://github.com/b-bianca/soat1-challenge1/wiki](https://github.com/b-bianca/soat1-challenge1/wiki)

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

## Collection do Postman

No diretório ./docs está a collection com as apis implementadas nesse projeto. Para utilizar, basta importar no Postman.<br>
Também é possível acessar o Postam Online:<br>
[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/16227218-ad366006-d6e5-41a8-8b14-0e5b79002ac0?action=collection%2Ffork&collection-url=entityId%3D16227218-ad366006-d6e5-41a8-8b14-0e5b79002ac0%26entityType%3Dcollection%26workspaceId%3De76668fb-982b-4d15-ab75-26131dab7174#?env%5BDEV%5D=W3sia2V5IjoiYmFzZV91cmwucmVzdGF1cmFudCIsInZhbHVlIjoibG9jYWxob3N0OjgwODAvYXBpL3YxIiwiZW5hYmxlZCI6dHJ1ZSwidHlwZSI6ImRlZmF1bHQifV0=)

## Demonstração Rodando Docker Compose e Consumindo API
https://github.com/b-bianca/soat1-challenge1/assets/83218983/865c92df-56b0-4cc7-b328-3921039d1f9c



