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

- **Go 1.21+**
- **PostgreSQL**
- **database/sql** (driver nativo)
- **net/http** (servidor HTTP nativo)

## Estrutura do Projeto

```
golang-product-manager/
├── cmd/
│   └── main.go              # Ponto de entrada da aplicação
├── internal/
│   ├── models/              # Modelos de dados
│   ├── storage/             # Camada de persistência
│   └── handlers/            # Handlers HTTP
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

- Go 1.21 ou superior
- PostgreSQL rodando localmente

### Configuração do Banco

```sql
CREATE DATABASE product_manager;

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Executando a Aplicação

```bash
# Clone o repositório
git clone <url-do-repositorio>
cd golang-product-manager

# Execute a aplicação
go run cmd/main.go
```

A API estará disponível em `http://localhost:8080`

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
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Notebook","price":2500.00}'

# Listar produtos
curl http://localhost:8080/products

# Buscar produto por ID
curl http://localhost:8080/products/1
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
- Sem logs estruturados
