package resolver

import (
	"context"
	"gqlgen_test/model"
	"sync"

	"github.com/google/uuid"
)

func (r *mutationResolver) CreateEvent(ctx context.Context, name string, description string, price float64, date *string, imageUrl string, streetImages []*model.NewStreetImageInput) (*model.Event, error) {
	// Criar um novo evento
	event := &model.Event{
		ID:          "1", // Idealmente, gere um UUID aqui.
		Name:        name,
		Description: description,
		Price:       price,
		Date:        date,
		ImageURL:    imageUrl,
		LocationID:  "default-location", // Defina um valor padrão ou ajuste a lógica
	}

	// Criar imagens associadas ao evento
	var images []*model.StreetImage
	for _, img := range streetImages {
		images = append(images, &model.StreetImage{
			ID:      "img-id", // Idealmente, gere um ID único.
			URL:     img.URL,
			EventID: event.ID,
		})
	}
	event.StreetImages = images

	// Aqui você pode salvar o evento no banco de dados (se estiver usando GORM)
	// err := r.DB.Create(event).Error
	// if err != nil {
	// 	return nil, err
	// }
	r.SubscriptionResolver.PublishEvent(event)

	return event, nil
}

var mu sync.Mutex

// Mapa para armazenar assinantes da subscription
var subscribers = make(map[string]chan *model.Event)

func (r *subscriptionResolver) EventCreated(ctx context.Context) (<-chan *model.Event, error) {
	eventChan := make(chan *model.Event, 1) // Canal bufferizado para evitar bloqueios
	subscriberID := uuid.New().String()     // Criando um ID único para o assinante

	// Adicionar o assinante à lista de inscritos
	mu.Lock()
	subscribers[subscriberID] = eventChan
	mu.Unlock()

	// Remover o assinante quando a conexão for fechada
	go func() {
		<-ctx.Done()
		mu.Lock()
		delete(subscribers, subscriberID)
		close(eventChan)
		mu.Unlock()
	}()

	return eventChan, nil
}
