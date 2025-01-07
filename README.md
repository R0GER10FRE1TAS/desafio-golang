
# API de Gerenciamento de Livros (com PostgreSQL e GORM)

Este projeto é uma API REST para gerenciar livros, desenvolvida em *GoLang, utilizando **Gin Framework* e *PostgreSQL*. A API permite criar, listar, atualizar e deletar livros.

## Funcionalidades

- *Criar Livro:* Adiciona um novo livro.
- *Listar Livros:* Retorna todos os livros cadastrados.
- *Atualizar Livro:* Atualiza informações de um livro existente.
- *Deletar Livro:* Remove um livro do sistema.

Cada livro possui os seguintes campos:
- *Título* (string): Nome do livro.
- *Autor* (string): Autor do livro.
- *Categoria* (string): Gênero ou categoria literária.
- *Sinopse* (string): Breve descrição ou resumo do livro.

## Configuração

### Pré-requisitos
- *GoLang* (1.18 ou superior)
- *PostgreSQL* (13 ou superior)

### Passos para Configurar

1. Instale o PostgreSQL:
   - Baixe e instale do site oficial: [postgresql.org](https://www.postgresql.org/download/).

2. Crie o banco de dados:
   bash
   psql -U postgres
   CREATE DATABASE desafio_golang;
   

3. Configure a aplicação:
   - No código, ajuste o DSN na função connectDB com suas credenciais do PostgreSQL.

4. *Migração Automática de Tabelas*:
   - Não é necessário criar manualmente as tabelas no banco de dados. O método AutoMigrate do GORM é usado para criar e configurar as tabelas automaticamente na primeira execução da aplicação. Isso simplifica o processo de configuração inicial e evita erros manuais na estrutura do banco de dados.
   - *Exemplo no código:*
     go
     db.AutoMigrate(&Book{})
     

## Como Rodar

1. Instale as dependências:
   bash
   go mod tidy
   

2. Inicie a aplicação:
   bash
   go run main.go
   

A API estará disponível em: http://localhost:8080

## Documentação das Rotas

### POST /books
*Descrição:* Adiciona um novo livro.
- *Corpo da Requisição:*
  json
  {
      "title": "O Senhor dos Anéis",
      "author": "J.R.R. Tolkien",
      "category": "Fantasia",
      "synopsis": "Uma história fantástica na Terra Média."
  }
  
- *Status de Sucesso:* 201 (Created)

### GET /books
*Descrição:* Retorna a lista de todos os livros cadastrados.

### PUT /books/:id
*Descrição:* Atualiza um livro específico.
- *Corpo da Requisição:*
  Apenas os campos que precisam ser alterados.
  json
  {
      "title": "O Hobbit"
  }
  
- *Status de Sucesso:* 200 (OK)

### DELETE /books/:id
*Descrição:* Remove um livro específico do sistema.
- *Status de Sucesso:* 200 (OK)
-

## Observações e Considerações Pessoais

  

Eu não tinha experiência prévia com as tecnologias que utilizei nesse projeto, então precisei pesquisar bastante.  
A escolha pelo PostgreSQL veio por alguns motivos. Durante o meu curso de FullStack, tive um contato breve com o MySQL, o que me deixou um pouco mais familiarizado com a estrutura SQL, mesmo que de forma bem superficial.  
 Ao pesquisar mais sobre as linguagens e comparar as opções, percebi que o PostgreSQL teria uma integração melhor com o GORM, o que fez a escolha se tornar mais viável. O GORM, por sua vez, facilitou bastante o trabalho, simplificando muitas etapas.
Embora eu já tenha utilizado outras tecnologias para criar APIs, como FastAPI no Python e Node.js, nenhuma delas me pareceu tão tranquila quanto o Go, especialmente com o uso do GORM.
Além disso, no desenvolvimento da estrutura básica do CRUD, optei por criar um POST mais flexível, evitando que o usuário precisasse digitar todos os dados novamente. E, graças à combinação do GORM com o SQL, implementei o AutoMigrate, que gerencia as tabelas automaticamente, sem a necessidade de escrever scripts SQL manualmente.
