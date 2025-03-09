# GraphQL Event Subscription API

## 📌 Descrição
Este é um projeto de API GraphQL que permite a criação e assinatura de eventos em tempo real. Utiliza [gqlgen](https://github.com/99designs/gqlgen) para gerar o esquema GraphQL e inclui suporte a **Subscriptions** para notificar assinantes sempre que um novo evento é criado.

## 🚀 Tecnologias Utilizadas
- **Golang** (Linguagem de programação)
- **gqlgen** (Biblioteca para GraphQL)
- **UUID** (Gerador de identificadores únicos)
- **HTTP Server** para servir a API

## 📂 Estrutura do Projeto
```
├── gqlgen_test
│   ├── generated/           # Código gerado automaticamente pelo gqlgen
│   ├── model/               # Modelos do GraphQL
│   ├── resolver/            # Resolvers de Query, Mutation e Subscription
│   ├── server.go            # Inicialização do servidor
│   ├── gqlgen.yml           # Configuração do gqlgen
│   ├── schema.graphql       # Definição do esquema GraphQL
│   ├── go.mod               # Dependências do projeto
```

## 🛠 Configuração e Instalação
1. **Clone o repositório:**
   ```sh
   git clone https://github.com/iamrosada0/gqlgen-trainer.git
   cd gqlgen-trainer
   ```

2. **Instale as dependências:**
   ```sh
   go mod tidy
   ```

3. **Execute o servidor:**
   ```sh
   go run main.go
   ```

4. **Acesse o Playground GraphQL:**
   Abra o navegador e acesse:
   ```
   http://localhost:8080/
   ```

## 📌 Funcionalidades
### 🔍 Queries Disponíveis
```graphql
query {
  users {
    id
    name
  }
  posts {
    id
    title
    content
  }
}
```

### ✏️ Mutations Disponíveis
#### Criar um novo evento
```graphql
mutation {
  createEvent(
    name: "Evento Teste"
    description: "Descrição do evento"
    price: 99.99
    date: "2025-01-01"
    imageUrl: "https://example.com/image.jpg"
    streetImages: [
      { url: "https://example.com/street1.jpg" }
      { url: "https://example.com/street2.jpg" }
    ]
  ) {
    id
    name
    description
    date
    streetImages {
      id
      url
    }
  }
}
```

### 📡 Subscription para eventos criados
#### Receber eventos em tempo real
```graphql
subscription {
  eventCreated {
    id
    name
    description
    date
    imageUrl
  }
}
```

## 🏗 Estrutura dos Resolvers
### **MutationResolver** (Criação de Eventos)
- Cria um novo evento e suas imagens
- Notifica todos os assinantes conectados via Subscription

### **SubscriptionResolver** (Eventos em Tempo Real)
- Mantém uma lista de assinantes ativos
- Envia eventos criados para todos os assinantes conectados

## 🔥 Exemplo de Saída no Log
```
2025/03/09 18:53:48 🚀 Servidor rodando em http://localhost:8080/
2025/03/09 18:53:57 📡 Assinante conectado à EventCreated
2025/03/09 18:54:14 📥 Criando um novo evento...
2025/03/09 18:54:14 🎉 Evento criado com sucesso!
2025/03/09 18:54:14 📢 Evento enviado para assinantes
```

## 📌 Contribuição
Sinta-se à vontade para abrir **issues** e enviar **pull requests**! 😃

## 📜 Licença
Este projeto está sob a licença MIT. Veja [LICENSE](LICENSE) para mais detalhes.

