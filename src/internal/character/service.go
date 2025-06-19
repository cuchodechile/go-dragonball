// Aqui se implementa la logica de negocio que es buscar en local o ir a buscar a wwww
package character

import (
	"context"
	"errors"
	"fmt"
)

// ExternalClient describe al adaptador HTTP (puerto secundario).
type ExternalClient interface {
	FetchByName(ctx context.Context, name string) (*Character, error)
}

// CharacterService contiene la lógica de orquestación.
type CharacterService struct {
	repo   CharacterRepository
	client ExternalClient
}

// Constructor
func NewCharacterService(repo CharacterRepository, client ExternalClient) *CharacterService {
	return &CharacterService{repo: repo, client: client}
}

// FindOrCreate busca en la BD y, si no existe, llama al API externa y persiste.
func (s *CharacterService) FindOrCreate(ctx context.Context, name string) (*Character, error) {
	// 1) Lookup en la base local
	ch, err := s.repo.FindByName(ctx, name)
	if err == nil {
		return ch, nil // cache hit
	}
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err // fallo real
	}

	// 2) No existe → API externa
	ext, err := s.client.FetchByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("external api: %w", err)
	}

	// 3) Persistir y devolver
	if err := s.repo.Save(ctx, *ext); err != nil {
		return nil, err
	}
	return ext, nil
}
