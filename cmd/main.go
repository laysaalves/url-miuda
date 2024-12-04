package main

import (
	"fmt"
	"github.com/laysaalves/url-miuda/db"
	"github.com/laysaalves/url-miuda/internal/handlers"
)

func main() {
	db.ConnectDB()
	url := handlers.GerarUrlAleatoria()
	fmt.Println("URL gerada:", url)
}
