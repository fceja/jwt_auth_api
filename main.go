package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, firstName, lastName, pw string) *Account {
	acnt, err := newAccount(firstName, lastName, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.createAccount(acnt); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account => ", acnt.Number)

	return acnt
}

func seedAcounts(s Storage) {
	seedAccount(s, "fseed", "cseed", "seeeeed")
}

func main() {
	seed := flag.Bool("seed", false, "seed db")
	flag.Parse()

	store, err := newPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("seeding the db")
		seedAcounts(store)
	}

	server := newAPIServer(":3000", store)
	server.run()
}
