package main

import (
	"context"
	"log"

	"cuchodechile.cl/reto-amaris/internal/character"
	ginhandler "cuchodechile.cl/reto-amaris/internal/character/handler"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	/* ---------- infraestructura ---------- */
	rdb := redis.NewClient(&redis.Options{Addr: "op-redis:6379"})
	repo := character.NewRedisRepository(rdb)
	api  := character.NewDragonBallHTTPClient( )

	/* ---------- dominio ---------- */
	svc := character.NewCharacterService(repo, api)

	/* ---------- transporte Gin ---------- */
	router := gin.Default()
	ginhandler.New(svc).Register(router)

	log.Println("API Gin corriendo en :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

	// opcional: cerrar Redis al terminar
	<-ctx.Done()
	_ = rdb.Close()
}



	