package main

import (
	"database/sql"
	"flag"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
	"wordle/console"
	"wordle/database"
	"wordle/web"
)

var modus string

type WordleGame interface {
	Play(word string)
}

func main() {
	flag.StringVar(&modus, "modus", "web", "game modus, either web oder console")
	flag.Usage()
	flag.Parse()

	db, err := sql.Open("sqlite3", "words.db")

	word := strings.ToUpper(database.GetRandomWord(db))

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var game WordleGame
	if strings.EqualFold(modus, "console") {
		game = &console.ConsoleGame{}
	} else {
		game = &web.WebGame{}
	}

	game.Play(word)

}
