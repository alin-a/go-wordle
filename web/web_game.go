package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"wordle/game"
)

type WebGame struct {
}

func (h *WebGame) Play(word string) {
	playInWeb(word)
}

func playInWeb(word string) {
	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	}

	submitHandler := func(w http.ResponseWriter, r *http.Request) {
		char1 := r.PostFormValue("char1")
		char2 := r.PostFormValue("char2")
		char3 := r.PostFormValue("char3")
		char4 := r.PostFormValue("char4")
		char5 := r.PostFormValue("char5")

		input := strings.ToUpper(fmt.Sprintf("%s%s%s%s%s", char1, char2, char3, char4, char5))
		result := game.CompareInputAndWord(input, word)

		var sb strings.Builder
		sb.WriteString("<div class='character-container'>")

		for index, elem := range result {
			letter := string(input[index])
			if game.DoesNotOccurInWord == elem {
				sb.WriteString(fmt.Sprintf("<div class='character character-grey'>%s</div>", letter))
			} else if game.OccursInWord == elem {
				sb.WriteString(fmt.Sprintf("<div class='character character-yellow'>%s</div>", letter))
			} else {
				sb.WriteString(fmt.Sprintf("<div class='character character-green'>%s</div>", letter))
			}
		}
		sb.WriteString("</div>")
		html := sb.String()

		tmpl, _ := template.New("result").Parse(html)
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/submit-word", submitHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
