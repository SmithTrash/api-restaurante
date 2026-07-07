# 🍽️ API Restaurante Online

![Go](https://img.shields.io/badge/Go-1.24-00ADD8?style=for-the-badge\&logo=go)
![REST API](https://img.shields.io/badge/API-REST-green?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-v1.0-blue?style=for-the-badge)

## 📖 Sobre o projeto

A API Restaurante Online é uma aplicação desenvolvida em Go (Golang) como parte da disciplina de Projeto Integrador da UNINTER.

O objetivo do projeto é simular o funcionamento do back-end de um sistema de gerenciamento de pedidos de um restaurante online, implementando um fluxo completo de criação de pedidos, processamento de pagamentos e atualização automática do status do pedido.

Além da implementação da API, o projeto foi estruturado utilizando arquitetura em camadas, permitindo a evolução para futuras versões com banco de dados, autenticação e deploy em produção.

---

## 🎯 Objetivos

* Desenvolver uma API REST utilizando Go.
* Implementar um fluxo crítico completo.
* Aplicar boas práticas de organização de código.
* Criar documentação técnica.
* Validar os endpoints utilizando o Postman.
* Preparar o projeto para futuras evoluções.

---

## 🚀 Tecnologias utilizadas

* Go (Golang)
* REST API
* JSON
* HTTP
* Postman
* Git
* GitHub

---

## 📂 Estrutura do Projeto

```text
api-restaurante/
│
├── handler/
│   ├── pedido_handler.go
│   └── pagamento_handler.go
│
├── models/
│   ├── pedido.go
│   └── pagamento.go
│
├── repository/
│   └── repository.go
│
├── routes/
│   └── routes.go
│
├── main.go
│
├── go.mod
│
└── README.md
```

---

## 📡 Endpoints

### GET /pedidos

Lista todos os pedidos cadastrados.

---

### POST /pedidos

Cria um novo pedido.

Exemplo:

```json
{
    "cliente": "Carlos"
}
```

---

### POST /pagamentos

Processa o pagamento de um pedido.

Exemplo:

```json
{
    "pedido_id": 1,
    "status": "PAGO"
}
```

---

## 🔄 Fluxo da Aplicação

```text
Cliente

↓

POST /pedidos

↓

Pedido Criado

↓

POST /pagamentos

↓

Status atualizado para PAGO

↓

GET /pedidos
```

---

## 🏗️ Arquitetura

O projeto foi organizado em camadas:

* Models
* Repository
* Handlers
* Routes

Essa organização facilita a manutenção, testes e futuras evoluções.

---

## 🧪 Testes

Todos os endpoints foram testados utilizando o Postman.

Fluxos validados:

* Criação de pedidos
* Consulta de pedidos
* Processamento de pagamento
* Atualização de status
* Validação de erros

---

## 🔒 Segurança

A API realiza validações básicas dos dados recebidos e utiliza códigos HTTP apropriados para cada situação.

Em versões futuras serão implementados:

* JWT
* Controle de acesso
* PostgreSQL
* HTTPS
* Logs
* Docker

---

## 🗺️ Roadmap

### ✅ Versão 1.0

* API REST
* Fluxo Pedido → Pagamento
* Arquitetura em camadas
* Documentação
* Testes

### 🚀 Próximas versões

* PostgreSQL
* Docker
* Swagger
* JWT
* Deploy em Cloud
* Flutter
* Painel Administrativo

---

## 👨‍💻 Autor

**Diego Moraes**

Projeto desenvolvido para a disciplina de Projeto da UNINTER.

Este projeto continuará evoluindo como projeto pessoal, recebendo novas funcionalidades e melhorias em versões futuras.
