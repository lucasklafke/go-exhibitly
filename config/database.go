package config

import (
	"exhibitly/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB é a conexão com o banco de dados
var DB *gorm.DB

// ConnectToDatabase realiza a conexão com o banco de dados
func ConnectToDatabase() {
	// Configuração do banco de dados (leitura de variáveis de ambiente)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),     // Ex: "localhost"
		os.Getenv("DB_USER"),     // Ex: "postgres"
		os.Getenv("DB_PASSWORD"), // Ex: "senha"
		os.Getenv("DB_NAME"),     // Ex: "meubanco"
		os.Getenv("DB_PORT"),     // Ex: "5432"
	)

	var err error
	// Inicializa a conexão com o banco de dados
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	log.Println("Banco de dados conectado com sucesso!")

	// Realiza a migração dos modelos, apenas após a conexão ser estabelecida
	err = DB.AutoMigrate(&models.Product{}) // Adicione outros modelos conforme necessário
	if err != nil {
		log.Fatalf("Falha ao migrar os modelos: %v", err)
	}

	log.Println("Migração de modelos realizada com sucesso!")
}
