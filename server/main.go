package main

import (
	"github.com/niloysikdar/Sell-It/server/dbconnector"
	"github.com/niloysikdar/Sell-It/server/routes"
)

func main() {
	dbconnector.InitialMigration()
	routes.InitRouter()
}
