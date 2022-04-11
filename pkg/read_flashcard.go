// pkg/read_flashcard.go
package pkg

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

func ReadFlashCard() {
    file, _ := ioutil.ReadFile("$HOME/Documents/Repos/flash-card-cli/flashcard.json")

    flashcards := FlashCardDB{}

    _ = json.Unmarshal([]byte(file), &flashcards)

    for i := 0; i < len(flashcards.FlashCards); i++ {
        fmt.Println("Question: ", flashcards.FlashCards[i].Question)
        fmt.Println("Answer: ", flashcards.FlashCards[i].Answer)
    }
}
