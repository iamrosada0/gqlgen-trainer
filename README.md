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

  # 📌 Dica: Para projetos grandes, divida o schema em arquivos separados.
  # Isso melhora a organização e a manutenção do código. Exemplo:
  #
  # schema:
  # - graph/schema.graphqls
  # - graph/user.graphqls
  # - graph/event.graphqls
  # - graph/post.graphqls

exec:
  filename: generated/generated.go  # Caminho do arquivo gerado pelo gqlgen

model:
  filename: model/models_gen.go  # Arquivo onde os modelos gerados automaticamente serão salvos
  package: model  # Pacote onde os modelos estarão armazenados

  # 📌 Observação: Todos os modelos gerados automaticamente pelo gqlgen são baseados no schema GraphQL.
  # Isso significa que qualquer tipo definido com "type" ou "input" no schema será incluído nesse arquivo.
  # 
  # ❗ Importante: Se você tiver tabelas com relacionamentos complexos, é recomendável **não deixar o gqlgen gerar automaticamente**.
  # Isso porque os modelos gerados não podem ser modificados diretamente, o que pode causar problemas no seu projeto.
  # Para esses casos, crie seus próprios modelos manualmente.
  #
  # 🔎 Exemplo: Abra o arquivo `model/event.go` e observe as linhas 11, 17 e 18.
  # Elas representam um relacionamento que deveria ter sido tratado manualmente.

autobind:
  - "gqlgen_test/model"  # Pacote onde os modelos personalizados serão vinculados

  # 📌 Observação: O `autobind` permite que você vincule modelos que **não devem ser gerados automaticamente** pelo gqlgen.
  # Isso é útil para evitar que certos modelos sejam sobrescritos, especialmente aqueles que possuem lógica específica.
  #
  # Se você precisar criar vários modelos manuais, basta adicioná-los nesta seção.

resolver:
  type: Resolver  # Tipo principal do resolver, responsável por mapear as queries, mutations e subscriptions
  # 📌 Observação: 
  #resolver:
  # type: Resolver  # Tipo principal do resolver
  # layout: follow-schema  # Usa a estrutura baseada no schema
  # dir: graph/resolvers   # Define a pasta onde os resolvers estarão
  # package: resolvers     # Define o pacote dos resolvers

  # OBS: O gqlgen gera um único arquivo resolver.go, mas para projetos grandes é recomendável separar os resolvers por funcionalidades. 
  # Você pode ter resolvers específicos para User, Event, Subscription, Post, etc., organizando-os em arquivos separados dentro do diretório de resolvers.
  #
  # ATENÇÃO: Mesmo definindo essa estrutura, ao regenerar os resolvers, o gqlgen pode sobrescrever e excluir arquivos personalizados fora do padrão gerado.
  # Por isso, é mais seguro utilizar apenas "type: Resolver" e gerenciar manualmente os resolvers, evitando perdas de código ao regenerar os arquivos.

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

