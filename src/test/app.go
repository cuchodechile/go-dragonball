package main

import (
    "fmt"
	"log"
	"cuchodechile.cl/reto-amaris/internal/character"
	"github.com/redis/go-redis/v9"
	"context"

   // "github.com/gin-gonic/gin"
)






// internal/drangonball/modelo.go

//test var c Character
//if err := json.Unmarshal(data, &c); err != nil {
 //   log.Fatal(err)
//}



func main() {

	fmt.Println("Inicio")

	rdb    := redis.NewClient(&redis.Options{Addr: "192.168.16.30:6379"})
	repo   := character.NewRedisRepository(rdb)
	client := character.NewDragonBallHTTPClient()           // <-- implementa ExternalClient
	svc    := character.NewCharacterService(repo, client)
	fmt.Println("buscar")

	ch, err := svc.FindOrCreate(context.Background(), "Goku")
	if err != nil { log.Fatal(err) }
	fmt.Printf("%+v\n", ch)

/*
		// 1. Crea repo y servicio
	repo    := character.NewCharacterRepository()
		service   := character.NewCharacterService(repo)

	// 2. Registra un personaje de prueba
	goku := character.Character{
		Name: "mono", Ki: "60.000.000", Race: "Saiyan", Gender: "Male",
	}
	if err := service.RegisterCharacter(goku); err != nil {
		log.Fatal(err)
	}

	// 3. Búscalo por nombre
	found, err := service.FindByName("Goku")
	if err != nil {
		log.Fatal(err)
	}

	// 4. Muestra el resultado
	fmt.Printf("Personaje encontrado: %+v\n", *found)
*/
}