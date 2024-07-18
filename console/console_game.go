package console

import (
	"fmt"
	"strings"
	"wordle/game"
)

var Reset = "\033[0m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Grey = "\033[37m"

type ConsoleGame struct {
}

func (h *ConsoleGame) Play(word string) {
	PlayOnConsole(word)
}

func PlayOnConsole(word string) {
	var colours = map[int]string{
		game.IsInCorrectPosition: Green,
		game.OccursInWord:        Yellow,
		game.DoesNotOccurInWord:  Grey,
	}
	doConsoleGameLoop(word, colours, 0)
}

func doConsoleGameLoop(word string, colours map[int]string, counter int) {
	if counter <= 6 {
		input := readInputFromConsole()
		checkInput(word, colours, counter, input)

		if strings.EqualFold(word, input) {
			fmt.Println(Green + input + Reset)
			fmt.Printf("Korrekt, das gesuchte Wort lautet: %s", word)
			return
		} else {
			result := game.CompareInputAndWord(strings.ToUpper(input), word)

			for i := 0; i < 5; i++ {
				colour := colours[result[i]]
				fmt.Print(colour + string([]rune(input)[i]) + Reset)
			}
			fmt.Println()
			updatedCounter := counter + 1
			doConsoleGameLoop(word, colours, updatedCounter)
		}
	} else {
		fmt.Println("Gescheitert, weil zu viele Versuche.")
		fmt.Printf("Das Wort lautet: %s", word)
	}
}

func checkInput(word string, colours map[int]string, counter int, input string) {
	if len(input) != 5 {
		fmt.Printf("Die Eingabe %s muss genau aus fünf Buchstaben bestehen\n", input)
		fmt.Println("Gib ein weiteres Wort mit fünf Buchstaben ein:")
		doConsoleGameLoop(word, colours, counter)
	}
}
