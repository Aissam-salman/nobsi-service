package data

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnexionDB() (*gorm.DB, error) {
	dsn := "host=localhost user=salman password=secret dbname=nobsi port=5432 sslmode=disable TimeZone=Europe/Paris"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestConnexionDB() {
	db, err := ConnexionDB()
	if err != nil {
		fmt.Println("Erreur de connexion à la base de données :", err)
		return
	}
	fmt.Println("Connexion à la base de données réussie")
	print(db)
}