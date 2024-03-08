package main

import "camarinb2096/form_example/internal/app/config/db"

func main() {
	config := db.NewConfig()
	_ = db.NewDb(config)

}
