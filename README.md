# 🗓️ Agendamento API (Golang)

Backend de uma plataforma de **agendamentos inteligentes**, desenvolvido em **Golang**, com autenticação **JWT**, regras reais de negócio e pronto para integração com **frontend Vue 3**.

---

## Pré-requisitos

- Go 1.19 ou superior

## Instalação

1. Clone o repositório:
   ```
   git clone https://github.com/lazaaro01/agendamento-api.git
   cd agendamento-api
   ```

2. Instale as dependências:
   ```
   go mod tidy
   ```

3. Execute o servidor:
   ```
   go run main.go  # Supondo que haja um main.go na raiz
   ```

A API estará rodando em `http://localhost:8080`.


## 🚀 Tecnologias Utilizadas

- Golang  
- Gin (framework HTTP)  
- GORM (ORM)  
- PostgreSQL  
- JWT (JSON Web Token)  
- Docker & Docker Compose  
- bcrypt (hash de senha)  

---

## 🧠 Funcionalidades

### 🔐 Autenticação
- Registro de usuário  
- Login  
- JWT  
- Middleware de proteção de rotas  

### 👤 Usuários
- Perfil **CLIENT** (padrão)  

### 🛠️ Serviços
- Criar serviço  
- Listar serviços  

### 🗓️ Agendamentos
- Criar agendamento  
- Listar agendamentos do usuário  
- Cancelar agendamento  
- Validação de conflito de horário  

---

## 📁 Arquitetura de Pastas

```text
agendamento-api/
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── config/
│   │   └── env.go
│   ├── database/
│   │   └── postgres.go
│   ├── models/
│   │   ├── user.go
│   │   ├── service.go
│   │   └── appointment.go
│   ├── handlers/
│   │   ├── auth_handler.go
│   │   ├── service_handler.go
│   │   └── appointment_handler.go
│   ├── services/
│   │   ├── jwt.go
│   │   └── password.go
│   ├── middlewares/
│   │   └── auth.go
│   └── routes/
│       └── routes.go
│
├── docker-compose.yml
├── .env
├── go.mod
└── go.sum