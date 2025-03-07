package resolver

import (
	"context"
	"gqlgen_test/model"
	"sync"

	"github.com/google/uuid"
)

// Estrutura para armazenar os canais das subscriptions
type SubscriptionResolver struct {
	mu          sync.Mutex
	subscribers map[string]chan *model.Event
}

// Criamos um novo SubscriptionResolver
func NewSubscriptionResolver() *SubscriptionResolver {
	return &SubscriptionResolver{
		subscribers: make(map[string]chan *model.Event),
	}
}

// Método para a Subscription
func (r *SubscriptionResolver) EventCreated(ctx context.Context) (<-chan *model.Event, error) {
	// Criar um canal para enviar eventos
	eventChan := make(chan *model.Event, 1)

	// Gerar um ID único para o assinante
	subscriberID := uuid.New().String()

	// Adicionar assinante ao mapa
	r.mu.Lock()
	r.subscribers[subscriberID] = eventChan
	r.mu.Unlock()

	// Remover assinante quando a conexão for fechada
	go func() {
		<-ctx.Done()
		r.mu.Lock()
		delete(r.subscribers, subscriberID)
		close(eventChan)
		r.mu.Unlock()
	}()

	return eventChan, nil
}

// Método para disparar eventos para todos os assinantes
func (r *SubscriptionResolver) PublishEvent(event *model.Event) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, ch := range r.subscribers {
		ch <- event
	}
}
