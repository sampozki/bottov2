package utils

import (
	"log"
	"strings"
)


func CheckNilErr(e error) {
 if e != nil {
  log.Fatal("Error message")
 }
}

func LogText(ID string, text string) {
	log.Println(ID + ": " + text)
}


// Returns true if some []words in content string, otherwise false
func ContainsAny(content string, words []string) bool {
    for _, w := range words {
        if strings.Contains(content, w) {
            return true
        }
    }
    return false
}