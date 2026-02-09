package main

type API struct {
	config config
}	

type config struct {
	addr string
	db dbConfig
}

type dbConfig struct {
	dbUrl string
}