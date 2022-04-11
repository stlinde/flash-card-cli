// pkg/flashcard.go
package pkg

type FlashCard struct {
    Question        string      `json:"question"`
    Answer          string      `json:"answer"`
    Tag             string      `json:"tag"`
    Goal            int         `json:"goal"`     
    Completed       int         `json:"completed"` 
    Tries           int         `json:"tries"` 
    SuccessRate     float64     `json:"successrate"` 
}

type FlashCardDB struct {
    FlashCards []FlashCard      `json:"flashcards"`
}

