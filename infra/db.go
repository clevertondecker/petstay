package infra

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Função para inicializar a conexão com o banco de dados
func InitDB() {
	// Obtendo as variáveis de ambiente
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Construindo a string de conexão com as variáveis de ambiente
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Conectando ao banco de dados
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	// Verificando se a conexão foi bem-sucedida
	if err := db.Ping(); err != nil {
		log.Fatal("Erro ao testar a conexão:", err)
	}

	// Create table owners if not exists.
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS owners (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(100),
			email VARCHAR(100),
			phone VARCHAR(20),
			address VARCHAR(200)
		);
	`)
	if err != nil {
	log.Fatal("Erro ao criar a tabela owners:", err)
	}

	log.Println("Conectado com sucesso ao MySQL!")
}

// Função para retornar a conexão com o banco
func GetDB() *sql.DB {
	return db
}

