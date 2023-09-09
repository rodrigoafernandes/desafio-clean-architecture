# desafio-clean-architecture
Pra este desafio, precisaremos criar a listagem das orders.

Esta listagem precisa ser feita com:

- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL

## TODOS
- [X] \- Configurar Rabbitmq através de variáveis de ambiente;
- [X] \- Criar arquivo para subir infra através do podman;
- [X] \- Configurar Database migration;
- [X] \- Configurar rotas de acordo com o método http;
- [X] \- Criar query para buscar as orders;
- [X] \- Criar teste unitário para a busca de orders;
- [X] \- Criar usecase para listar as orders;
- [X] \- Criar endpoint http para listar orders;
- [X] \- Criar Service GRPC para listar orders;
- [X] \- Criar Query GraphQL para listar orders;
- [X] \- Criar Containerfile para rodar a aplicação;
- [X] \- Configurar arquivos a serem ignorados pelo container;
- [X] \- Documentar processo de build da aplicação;
- [X] \- Documentar processo para rodar a aplicação;

## Build da App
A aplicação pode ser construída de 2 formas:
1. Binário Golang
2. Imagem de Container
### Build do Binário
Para construir a aplicação para rodar como um binário, basta executar o comando:
```shell
go build -o desafio-clean-architecture github.com/rodrigoafernandes/desafio-clean-architecture/cmd/ordersystem
```
Um binário com o nome "desafio-clean-architecture" será gerado no diretório atual
### Build de Imagem de Container
Para construir uma Imagem de Container, desenvolvi um "Containerfile" para facilitar a criação. Existem diversas ferramentas para facilitar a criação da imagem, listarei abaixo tres opções: Buildah, Podman e Docker
#### Build da Imagem com Buildah
Para construir a imagem do container com a ferramenta Buildah, utilize o seguinte comando:
```shell
buildah bud -f Containerfile -t rodrigoafernandes/golang-expert/desafio-clean-architecture:1.0.0 .
```
#### Build da Imagem com Podman
Para construir a imagem do container com a ferramenta Podman, utilize o seguinte comando:
```shell
podman build -f Containerfile -t rodrigoafernandes/golang-expert/desafio-clean-architecture:1.0.0 .
```
#### Build da Imagem com Docker
Para construir a imagem do container com a ferramenta Docker, utilize o seguinte comando:
```shell
docker build -f Containerfile -t rodrigoafernandes/golang-expert/desafio-clean-architecture:1.0.0 .
```

## Rodando a App
Assim como para construir, rodar a aplicação pode ser feito de duas formas: executar o binário ou executar a imagem do container construída.<br/>
Caso os componentes de infraestrutura não rodem na mesma rede(localhost), é necessário configura-los através de variáveis de ambiente. As variáveis que a aplicação espera receber para configurar são:<br/>
| Nome                 | Finalidade                        |
|----------------------|-----------------------------------|
| DB_DRIVER            | Nome do driver do Database        |
| DB_HOST              | Host do Database                  |
| DB_PORT              | Porta do Database                 |
| DB_USER              | Usuário do Database               |
| DB_PASSWORD          | Senha do Usuário do Database      |
| DB_NAME              | Nome do Schema no Database        |
| WEB_SERVER_PORT      | Porta para rodar o WebServer      |
| GRPC_SERVER_PORT     | Porta para rodar o gRPC Server    |
| GRAPHQL_SERVER_PORT  | Porta para rodar o GraphQL Server |
| RABBITMQ_USER        | Usuário do Rabbimq                |
| RABBITMQ_PWD         | Senha do Usuário do RabbiMQ       |
| RABBITMQ_HOST        | Host do Rabbitmq                  |
| RABBITMQ_PORT        | Porta do Rabbitmq                 |
### Rodando o Binário
./desafio-clean-architecture
### Rodando a Imagem de Container com Podman
Para rodar a imagem do container com a ferramenta Podman, utilize o seguinte comando:
```shell
podman container run -d --rm --name desafio-clean-architecture -p 8000:8000 -p 8080:8080 -p 50051:50051 -e DB_DRIVER=mysql -e DB_HOST=172.16.10.1 -e DB_PORT=3306 -e DB_PASSWORD=root -e DB_USER=root -e DB_NAME=orders -e WEB_SERVER_PORT=":8000" -e GRPC_SERVER_PORT=50051 -e GRAPHQL_SERVER_PORT=8080 -e RABBITMQ_USER=guest -e RABBITMQ_PWD=guest -e RABBITMQ_HOST=172.16.10.2 -e RABBITMQ_PORT=5672 rodrigoafernandes/golang-expert/desafio-clean-architecture:1.0.0
```
Caso queira rodar toda a infraestrutura(mysql e rabbitmq) junto com a imagem da aplicação, existem dois arquivos que facilitam a execução dos containers:
1. clean-arch-infra.yaml
2. clean-arch-infra-with-mysql-volume.yaml
O primeiro arquivo, cria a infraestrutura de forma efêmera, ou seja, quando finalizar os containers os dados serão perdidos. Para rodar desta forma, utilize o seguinte comando:
```shell
podman play kube clean-arch-infra.yaml
```
O segundo arquivo, utiliza um volume para armazenar os dados do Mysql, podendo parar os containers e rodar novamente, mantendo os registros criados.<br /> Altere no arquivo "clean-arch-infra-with-mysql-volume.yaml" o valor "<CAMINHO-PARA-VOLUME-MYSQL>" para um diretório que deseja armazenar os dados do container Mysql.<br /> Para rodar utilizando esta forma, utilize o comando:
```shell
podman play kube clean-arch-infra-with-mysql-volume.yaml
```
### Rodando a Imagem de Container com Docker
Para rodar a imagem do container com a ferramenta Docker, utilize o seguinte comando:
```shell
docker container run -d --rm --name desafio-clean-architecture -p 8000:8000 -p 8080:8080 -p 50051:50051 -e DB_DRIVER=mysql -e DB_HOST=172.16.10.1 -e DB_PORT=3306 -e DB_PASSWORD=root -e DB_USER=root -e DB_NAME=orders -e WEB_SERVER_PORT=":8000" -e GRPC_SERVER_PORT=50051 -e GRAPHQL_SERVER_PORT=8080 -e RABBITMQ_USER=guest -e RABBITMQ_PWD=guest -e RABBITMQ_HOST=172.16.10.2 -e RABBITMQ_PORT=5672 rodrigoafernandes/golang-expert/desafio-clean-architecture:1.0.0
```
### Comandos para executar ao criar a infraestrutura
Após subir os componentes de infraestrutura(MySql e RabbitMQ), é necessário rodar os seguintes comandos: 
```shell
migrate create -ext=sql -dir=sql/migrations -seq init
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up
```

### Referências
- [Install Golang](https://go.dev/doc/install)
- [Install GVM(Golang Version Manager)](https://github.com/moovweb/gvm#installing)
- [Install Golang Migrate DB Tool](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)
- [Install GRPC Tools](https://grpc.io/docs/languages/go/quickstart/)
- [Install Evans(gRPC command line cli tool)](https://github.com/ktr0731/evans#installation)
- [Install Buildah](https://github.com/containers/buildah/blob/main/install.md)
- [Install Podman](https://podman.io/docs/installation)
- [Install Docker](https://docs.docker.com/get-docker/)
- [Build Golang Binary](https://go.dev/doc/tutorial/compile-install)
- [Build Container Image With Buildah](https://manpages.ubuntu.com/manpages/impish/man1/buildah-bud.1.html)
- [Build Container Image With Podman](https://docs.podman.io/en/latest/markdown/podman-build.1.html)
- [Build Container Image With Docker](https://docs.docker.com/engine/reference/commandline/build/)