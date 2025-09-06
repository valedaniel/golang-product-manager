# Product Manager API

> **⚠️ PROJETO APENAS PARA ESTUDOS**
>
> Esta aplicação foi desenvolvida exclusivamente para fins educacionais e aprendizado de Go (Golang).

## Sobre o Projeto

API REST simples para gerenciamento de produtos, desenvolvida em Go com PostgreSQL. O objetivo é demonstrar conceitos básicos de:

- Estruturação de projetos Go
- APIs REST com handlers HTTP
- Integração com banco de dados PostgreSQL
- Padrões de arquitetura (Repository Pattern)
- Context e tratamento de erros

## Tecnologias Utilizadas

- **Go 1.25+**
- **PostgreSQL**
- **database/sql** (driver nativo)
- **net/http** (servidor HTTP nativo)
- **github.com/lib/pq** (driver PostgreSQL)
- **github.com/joho/godotenv** (carregamento de variáveis de ambiente)

## Estrutura do Projeto

```
golang-product-manager/
├── cmd/
│   └── api/
│       └── main.go          # Ponto de entrada da aplicação
├── db/
│   └── migrations/          # Scripts de migração do banco
├── internal/
│   ├── models/              # Modelos de dados
│   ├── storage/             # Camada de persistência
│   └── handler/             # Handlers HTTP
├── .env                     # Variáveis de ambiente
├── go.mod                   # Dependências do módulo
└── README.md
```

## Funcionalidades

- ✅ Criar produto
- ✅ Listar produtos
- ✅ Buscar produto por ID
- ✅ Atualizar produto
- ✅ Deletar produto

## Como Executar

### Pré-requisitos

- Go 1.25 ou superior
- PostgreSQL rodando localmente

### Configuração do Banco

1. Crie o banco de dados:

```sql
CREATE DATABASE products_db;
```

2. Configure as variáveis de ambiente no arquivo `.env`:

```env
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=products_db
```

3. Execute a migração (ou crie a tabela manualmente):

```sql
CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  price NUMERIC(10, 2) NOT NULL,
  createdAt TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updatedAt TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

### Executando a Aplicação

```bash
# Clone o repositório
git clone <url-do-repositorio>
cd golang-product-manager

# Instale as dependências
go mod tidy

# Execute a aplicação
go run cmd/api/main.go
```

A API estará disponível em `http://localhost:3000` (ou na porta configurada na variável `DB_ADDR`)

## Endpoints da API

| Método | Endpoint         | Descrição                |
| ------ | ---------------- | ------------------------ |
| POST   | `/products`      | Criar produto            |
| GET    | `/products`      | Listar todos os produtos |
| GET    | `/products/{id}` | Buscar produto por ID    |
| PUT    | `/products/{id}` | Atualizar produto        |
| DELETE | `/products/{id}` | Deletar produto          |

## Exemplo de Uso

```bash
# Criar produto
curl -X POST http://localhost:3000/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Notebook","price":2500.00}'

# Listar produtos
curl http://localhost:3000/products

# Buscar produto por ID
curl http://localhost:3000/products/1
```

## Objetivos de Aprendizado

Este projeto foi criado para praticar:

- [x] Organização de código Go
- [x] Manipulação de banco de dados
- [x] Criação de APIs REST
- [x] Tratamento de erros
- [x] Uso de interfaces
- [x] Context em Go

## Limitações

- Sem autenticação/autorização
- Sem validações robustas
- Sem testes unitários
