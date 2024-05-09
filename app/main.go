package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

var Ctx = context.Background()

func main() {
    // Initialisez un client Redis
    client := redis.NewClient(&redis.Options{
        Addr:     "redis-10596.c300.eu-central-1-1.ec2.redns.redis-cloud.com:10596", // Adresse du serveur Redis
        Password: "YDadsk5SrAfuFeYPGFisSQX4YvsQs2HV",               // Mot de passe, si nécessaire
        DB:       0,                // Numéro de la base de données, par défaut 0
    })

    // Test de la connexion
    pong, err := client.Ping().Result()
    if err != nil {
        panic(fmt.Errorf("erreur lors de la connexion a redis : %v", err))
    }
    fmt.Println("Connexion réussie à Redis :", pong)


    TestConnexionDB
}