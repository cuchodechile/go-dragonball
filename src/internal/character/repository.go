package character

import (
	"fmt"
	"context"
)

// Character es tu DTO / entidad simple.
type Character struct {
	ID          int
	Name        string
	Ki          string
	MaxKi       string
	Race        string
	Gender      string
	Description string
	Image       string
	Affiliation string
}

// ErrNotFound indica que el registro no existe en la BD.
var ErrNotFound = fmt.Errorf("character not found")

// CharacterRepository es el **puerto** de persistencia.
// El servicio solo dependerá de esta abstracción.
type CharacterRepository interface {
	Save(ctx context.Context, c Character) error
	FindByName(ctx context.Context, name string) (*Character, error)
}
