package main

import (
	"github.com/eternaleight/go-backend/config"
	"github.com/eternaleight/go-backend/di"
	"github.com/eternaleight/go-backend/interfaces"
)

func main() {
	dsn := config.LoadConfig() // dsn Data Source Name
	db := config.InitializeDatabase(dsn)

	// 依存性を初期化
	authHandler, productHandler, purchaseHandler, userHandler, postHandler := di.InitializeHandlers(db)

	// ルーターを設定
	r := interfaces.SetupRouter(authHandler, productHandler, purchaseHandler, userHandler, postHandler)

	r.Run(":8001")
}
