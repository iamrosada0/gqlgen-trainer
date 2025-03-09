# 🚀 Projeto GraphQL com gqlgen

Este projeto é uma API GraphQL construída com **gqlgen** para gerenciar eventos, imagens de rua e usuários. Ele inclui suporte a **subscriptions**, permitindo que clientes recebam notificações em tempo real sobre eventos criados.

## 📂 Estrutura do Projeto

```
/gqlgen_test
│── /generated      # Código gerado automaticamente pelo gqlgen
│── /model          # Modelos usados no GraphQL
│── /resolver       # Resolvers das queries, mutations e subscriptions
│── main.go         # Ponto de entrada do servidor
│── gqlgen.yml      # Configuração do gqlgen
│── graph/schema.graphqls  # Definição do schema GraphQL
```

## 📜 Schema GraphQL

### 🧑 Usuários e Posts
```graphql
type User {
  id: ID!
  name: String!
  post: [Post!]!
}

type Post {
  id: ID!
  title: String!
  content: String!
  author: User!
}
```

### 📍 Eventos e Imagens
```graphql
type StreetImage {
  id: ID!
  url: String!
  event: Event!
}

type Event {
  id: ID!
  name: String!
  description: String!
  date: Date
  imageUrl: String!
  streetImages: [StreetImage!]!
}
```

### 🔍 Queries
```graphql
type Query {
  users: [User!]!
  posts: [Post!]!
}
```

### ✍️ Mutations
```graphql
type Mutation {
  createEvent(
    name: String!,
    description: String!,
    price: Float!,
    date: Date,
    imageUrl: String!,
    streetImages: [NewStreetImageInput!]!
  ): Event!
  createPost(title: String!, content: String!): Post!
}
```

### 📡 Subscription (Eventos em tempo real)
```graphql
type Subscription {
  eventCreated: Event!
}
```

## ⚙️ Configuração do gqlgen (gqlgen.yml)

O arquivo `gqlgen.yml` define como os arquivos do GraphQL são organizados e gerados:

```yaml
schema:
  - graph/schema.graphqls  # Caminho do schema GraphQL

exec:
  filename: generated/generated.go  # Arquivo onde o código gerado será salvo

model:
  filename: model/models_gen.go  # Local dos modelos gerados automaticamente
  package: model  # Pacote onde os modelos estão

autobind:
  - "gqlgen_test/model"  # Pacote para vincular automaticamente os modelos

resolver:
  type: Resolver  # Tipo principal do resolver
```

## 🚀 Rodando o Servidor

1. Instale as dependências:
   ```sh
   go mod tidy
   ```
2. Gere os resolvers e código GraphQL:
   ```sh
   go run github.com/99designs/gqlgen generate
   ```
3. Inicie o servidor:
   ```sh
   go run main.go
   ```
4. Acesse o playground do GraphQL em:
   [http://localhost:8080/](http://localhost:8080/)

## 🛠️ Implementação do Subscription

A implementação do `eventCreated` permite que clientes recebam notificações sempre que um evento for criado.

### Resolver do Subscription
```go
func (r *subscriptionResolver) EventCreated(ctx context.Context) (<-chan *model.Event, error) {
    log.Println("📡 Novo assinante conectado à EventCreated")
    eventStream := make(chan *model.Event, 1)
    
    mu.Lock()
    subscribers[eventStream] = true
    mu.Unlock()
    
    go func() {
        <-ctx.Done()
        mu.Lock()
        delete(subscribers, eventStream)
        close(eventStream)
        mu.Unlock()
        log.Println("❌ Assinante desconectado de EventCreated")
    }()
    
    return eventStream, nil
}
```

### Criando um Evento e Notificando os Assinantes
```go
func (r *mutationResolver) CreateEvent(ctx context.Context, name string, description string, price float64, date *string, imageUrl string, streetImages []*model.NewStreetImageInput) (*model.Event, error) {
    log.Println("📥 Criando um novo evento...")
    event := &model.Event{
        ID:          uuid.New().String(),
        Name:        name,
        Description: description,
        Price:       price,
        Date:        date,
        ImageURL:    imageUrl,
    }

    go func() {
        mu.Lock()
        for subscriber := range subscribers {
            subscriber <- event
        }
        mu.Unlock()
        log.Println("📢 Evento enviado para assinantes")
    }()
    
    log.Println("🎉 Evento criado com sucesso!")
    return event, nil
}
```

## 🏗️ Contribuindo

1. **Fork** o repositório
2. **Clone** o projeto
   ```sh
   git clone https://github.com/seu-usuario/gqlgen_test.git
   ```
3. **Crie uma branch** para sua funcionalidade
   ```sh
   git checkout -b minha-feature
   ```
4. **Implemente e envie suas mudanças**
   ```sh
   git commit -m "Adicionando nova funcionalidade"
   git push origin minha-feature
   ```
5. **Abra um Pull Request** 🚀

## 📜 Licença

Este projeto está licenciado sob a **MIT License**.

