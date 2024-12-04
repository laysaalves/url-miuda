package main

import (
	"fmt"
	"github.com/laysaalves/url-miuda/db"
	"github.com/laysaalves/url-miuda/internal/handlers"
)

func main() {
	dbInstance, err := db.ConnectDB()
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return
	}

	url := handlers.GerarUrlAleatoria()
	fmt.Println("URL gerada:", url)

	handlers.EncurtarUrl(dbInstance, url)
}
