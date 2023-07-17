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
- [X] \- Criar usecase para listar as orders;
- [X] \- Criar endpoint http para listar orders;
- [ ] \- Criar Service GRPC para listar orders;
- [ ] \- Criar Query GraphQL para listar orders;
- [ ] \- Criar Containerfile para rodar a aplicação;
- [ ] \- Configurar arquivos a serem ignorados pelo container;
- [ ] \- Documentar processo de build da aplicação;
- [ ] \- Documentar processo para rodar a aplicação;

### Commands
```shell
migrate create -ext=sql -dir=sql/migrations -seq init
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up
```