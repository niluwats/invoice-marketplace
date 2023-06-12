package main

import (
	"fmt"

	"github.com/niluwats/invoice-marketplace/pkg/db"
)

func main() {
	dbClient := db.SetupDBConn()
	fmt.Println(dbClient)
}
