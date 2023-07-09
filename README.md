# soat1-challenge1

Os diagramas estão linkados na wiki do projeto

## Collection do Postman

No diretório ./docs está a collection com as apis implementadas nesse projeto. Para utilizar, basta importar no postman.


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


Para executar o projeto, é necessário ter o docker destop instalado e subir as instancias usando o docker compose via IDE ou linha de comando.
' docker compose -f "docker-compose.yml" up -d --build '