package resolver

import (
	"context"
	"log"
	"sync"

	"gqlgen_test/model"

	"github.com/google/uuid"
)

// Canal global para eventos
var eventChannel = make(chan *model.Event, 1)

// Lista de assinantes
var subscribers = make(map[chan *model.Event]bool)
var mu sync.Mutex

// EventCreated agora repassa eventos para todos os assinantes
func (r *subscriptionResolver) EventCreated(ctx context.Context) (<-chan *model.Event, error) {
	log.Println("📡 Novo assinante conectado à EventCreated")

	eventStream := make(chan *model.Event, 1)

	mu.Lock()
	subscribers[eventStream] = true
	mu.Unlock()

	// Remover o assinante quando desconectar
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

// CreateEvent agora notifica todos os assinantes
func (r *mutationResolver) CreateEvent(ctx context.Context, name string, description string, price float64, date *string, imageUrl string, streetImages []*model.NewStreetImageInput) (*model.Event, error) {
	log.Println("📥 Criando um novo evento...")

	event := &model.Event{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		Date:        date,
		ImageURL:    imageUrl,
		LocationID:  "default-location",
	}

	// Criar imagens associadas ao evento
	var images []*model.StreetImage
	for _, img := range streetImages {
		imageID := uuid.New().String()
		images = append(images, &model.StreetImage{
			ID:      imageID,
			URL:     img.URL,
			EventID: event.ID,
		})
		log.Printf("🖼️ Imagem associada ao evento: ID=%s, URL=%s\n", imageID, img.URL)
	}
	event.StreetImages = images

	// Publicar evento para todos os assinantes
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
