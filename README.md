# Titulo

Descriçao.

# Requisitos

Algumas ferramentas são necessárias para utilizar a aplicação:

- Docker

# Execução

O ambiente de execução é inserido em um Docker Compose, onde possui tanto os serviços de banco de dados Postgres quanto a própria API Golang. Para executar, clone o repositório e execute os comandos:

```
$ docker compose up --build -d
```

# Tecnologias e Bibliotecas

- Golang 1.22 - Linguagem de programação
- PostgreSQL - Banco de dados
- SQLX - Executar operações no banco de dados
- Golang Migrate - Executar migrações no banco de dados
- Testify - Realizar testes mais claros

# API

O servidor HTTP foi todo feito com ferramentas nativas do Golang.

## Autenticação

A autenticação foi feita no formato JWT, armazenando o `id` do usuário no campo `sub` do token.

Para acessar uma rota que requer autenticação, é necessário enviar um `token de acesso`, que é recebido quando se autentica, e enviar no cabeçalho da requisição o header `Authorization` e o valor `Bearer seu_token`.

## Rotas

| Método | URL                                   | Privada | Funcionalidade                                                                            |
| :----- | ------------------------------------- | :-----: | ----------------------------------------------------------------------------------------- |
| POST   | /api/v1/login                         |    -    | Autentica utilizando CPF e Secret da conta, retornando JWT para acesso de rotas privadas. |
| GET    | /api/v1/accounts                      |    -    | Retorna todas contas cadastradas.                                                         |
| POST   | /api/v1/accounts                      |    -    | Cria uma nova conta.                                                                      |
| GET    | /api/v1/accounts/{account_id}/balance |    -    | Retorna o saldo da conta.                                                                 |
| GET    | /api/v1/transfers                     |   Sim   | Retorna a lista de transferências em que a conta serviu como origem da operação.          |
| POST   | /api/v1/transfers                     |   Sim   | Faz uma transferência da conta autenticada para uma conta informada.                      |

## Dados para execução

Modelo padrão de resposta

```
{
  "message": "mensagem do contexto",
  "status_code": 200,
  "data": {
    ...conteúdo
  }
}
```

### `[POST]` /api/v1/accounts

Caso de Sucesso: Rota para criar uma conta, retornando os dados da conta e status code 201.

Header: `Authorization: Bearer seu_token`

Body da requisição:

```
{
    "name": "dummyy doe",
	"cpf": "24628728097",
	"secret": "password123"
}
```

Body da resposta:

```
{
  "message": "Account created",
  "status_code": 201
}
```
