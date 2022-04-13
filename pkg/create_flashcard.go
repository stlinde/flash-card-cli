// pkg/create_flashcard
package pkg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"github.com/rs/xid"
)

func CreateFlashCard() {
    // We need to locate where the file will be stored

    // Opening database file
    // os.CREATE ensures that if there is no file, the file will be created.
    jsonFile, err := os.OpenFile("flashcard.json", os.O_RDWR|os.O_CREATE, 0755)
    if err != nil {
        log.Fatal(err)
    }

    // Reading from database file
    jsonData, err := ioutil.ReadAll(jsonFile)
    if err != nil {
        fmt.Println(err)
    }

    // flashcards is used for updating the database
    // flashcard is the new flashcard to be added. 
    var flashcards []FlashCard
    var flashcard FlashCard

    // Unmarshalling the flashcard json file
    if err := json.Unmarshal(jsonData, &flashcards); err != nil {
        log.Println(err)
    }

    // Creating input reader
    // The input reader is needed to be able to read strings with spaces.
    inputReader := bufio.NewReader(os.Stdin)

    // Enter data
    fmt.Printf("Enter question: ")
    question, _ := inputReader.ReadString('\n') 
    flashcard.Question = strings.Trim(question, "\n")

    fmt.Printf("Enter answer: ")
    answer, _ := inputReader.ReadString('\n') 
    flashcard.Answer = strings.Trim(answer, "\n")

    fmt.Printf("Enter tag: ")
    tag, _ := inputReader.ReadString('\n') 
    flashcard.Tag = strings.Trim(tag, "\n")

    flashcard.Completed = 0
    flashcard.Goal = 5
    flashcard.Tries = 0
    flashcard.SuccessRate = 0.0
    flashcard.ID = xid.New().String()


    // Adding the new flashcard to the array of flashcards read from the db. 
    flashcards = append(flashcards, flashcard)

    newJsonData, err := json.MarshalIndent(flashcards, "", " ")
    if err != nil {
        log.Println(err)
    }

    err = ioutil.WriteFile("flashcard.json", newJsonData, 0644)

}

