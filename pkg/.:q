// pkg/create_flashcard
package pkg

import (
	"encoding/json"
	"fmt"
	"os"
    "log" 
)

func CreateFlashCard() {
    // Creating empty instance of the FlashCard struct.
    var flashcard FlashCard

    // Enter data
    fmt.Printf("Enter question: ")
    fmt.Scanf("%s\n", &flashcard.Question) 

    fmt.Printf("Enter answer: ")
    fmt.Scanf("%s\n", &flashcard.Answer) 

    fmt.Printf("Enter tag: ")
    fmt.Scanf("%s\n", &flashcard.Tag)

    flashcard.Completed = 0
    flashcard.Goal = 5
    flashcard.Tries = 0
    flashcard.SuccessRate = 0.0


    // We need to open the file in append mode.
    f, err := os.OpenFile("flashcard.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }

    // The file variable is used to write a struct to json format.
    file, _ := json.MarshalIndent(flashcard, "", " ")

    // Appending newly created flashcard to our json file
    if _, err := f.Write(file); err != nil {
        log.Fatal(err)
    }
    if err := f.Close(); err != nil {
        log.Fatal(err) 
    }

}
