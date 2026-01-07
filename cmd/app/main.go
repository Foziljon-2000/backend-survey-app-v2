package main

import (
	"backend-survey-app-v2/internal/storage"
	"backend-survey-app-v2/internal/transport/http/router"
	"backend-survey-app-v2/pkg/databases"
	"fmt"
	"net/http"
)

func main() {
	db := databases.DB()
	defer db.Close()

	storage.GetDB(db)

	fmt.Println("listening localhost:2025")
	http.ListenAndServe("localhost:2025", router.NewRouterCompl())
}
