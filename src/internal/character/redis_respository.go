package character

import (
	"context"
	"strconv"
	"strings"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// redisRepo cumple CharacterRepository usando Hashes.
type redisRepo struct {
	rdb *redis.Client
}

// NewRedisRepository crea el adaptador; injéctale un *redis.Client.
func NewRedisRepository(rdb *redis.Client) CharacterRepository {
	return &redisRepo{rdb: rdb}
}

// ---------- Save ----------

func (r *redisRepo) Save(ctx context.Context, c Character) error {
	key := keyFor(c.Name)

	values := map[string]interface{}{
		"id":          c.ID,
		"name":        c.Name,
		"ki":          c.Ki,
		"maxKi":       c.MaxKi,
		"race":        c.Race,
		"gender":      c.Gender,
		"description": c.Description,
		"image":       c.Image,
		"affiliation": c.Affiliation,
	}
	return r.rdb.HSet(ctx, key, values).Err()
}

// ---------- FindByName ----------

func (r *redisRepo) FindByName(ctx context.Context, name string) (*Character, error) {
	key := keyFor(name)
		fmt.Printf("buscar en redis [%s] %x", key,ctx)

	m, err := r.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		fmt.Println("no encontrado")
		return nil, err
	}
	if len(m) == 0 {
		return nil, ErrNotFound
	}

	id, _ := strconv.Atoi(m["id"]) // ignoramos error: ""→0
	return &Character{
		ID:          id,
		Name:        m["name"],
		Ki:          m["ki"],
		MaxKi:       m["maxKi"],
		Race:        m["race"],
		Gender:      m["gender"],
		Description: m["description"],
		Image:       m["image"],
		Affiliation: m["affiliation"],
	}, nil
}

// Helpers

func keyFor(name string) string {
	return "character:" + strings.ToLower(name)
}
