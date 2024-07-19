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

var tries = 0
var wordleOutputHTML = ""

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
		tries++

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
		wordleOutputHTML += sb.String()

		if strings.ToUpper(word) == input {
			sb.WriteString(fmt.Sprintf("<div class='character-container'>Korrekt, das Wort lautet %s</div>", word))
			reset()
		} else if tries >= 6 {
			sb.WriteString(fmt.Sprintf("<div class='character-container'>Das Wort lautet %s</div>", word))
			reset()
		} else {
			formHTML := `    <form hx-post="/submit-word">
        <div class="character-container">
            <input id="char1" name="char1" type="text" pattern="[a-zA-Z]{1}" maxlength="1"
                   class="character character-input" oninput="this.value = this.value.toUpperCase()">
            <input id="char2" name="char2" type="text" pattern="[a-zA-Z]{1}" maxlength="1"
                   class="character character-input" oninput="this.value = this.value.toUpperCase()">
            <input id="char3" name="char3" type="text" pattern="[a-zA-Z]{1}" maxlength="1"
                   class="character character-input" oninput="this.value = this.value.toUpperCase()">
            <input id="char4" name="char4" type="text" pattern="[a-zA-Z]{1}" maxlength="1"
                   class="character character-input" oninput="this.value = this.value.toUpperCase()">
            <input id="char5" name="char5" type="text" pattern="[a-zA-Z]{1}" maxlength="1"
                   class="character character-input" oninput="this.value = this.value.toUpperCase()">
        </div>

        <div class="character-container">
            <button type="submit">Prüfen</button>
        </div>
    </form>`
			sb.WriteString(formHTML)
		}

		html := sb.String()
		tmpl, _ := template.New("result").Parse(html)
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/submit-word", submitHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func reset() {
	tries = 0
	wordleOutputHTML = ""
}
