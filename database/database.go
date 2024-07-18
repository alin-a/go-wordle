package database

import (
	"database/sql"
	"log"
	"math/rand"
)

func GetRandomWord(db *sql.DB) string {
	count := countWords(db)
	v := rand.Intn(count-1) + 1
	row := db.QueryRow("SELECT word FROM words WHERE id = ? ", v)

	var word string
	if err := row.Scan(&word); err != nil {
		log.Fatal(err)
	}
	return word
}

func countWords(db *sql.DB) int {
	rows, err := db.Query("SELECT COUNT(*) FROM words")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}
	return count
}
