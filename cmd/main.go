package main

import "os"


func main() {
	config := config{
		addr: ":8080",
		db: dbConfig{
			dbUrl: "postgres://	user:password@localhost:5432/mydb",
		},
	}

	api := API{
		config: config,
	}

	if err := api.run(api.mount()); err != nil {
		println("Server has failed to start", err.Error())
		os.Exit(1)
	}
	
}
