package main

import (
    "strings"
)

func cleanInput(text string) []string {
   // Trim leading and trailing whitespace
    trimmed := strings.TrimSpace(text)
    
    // Split by whitespace (handles multiple spaces, tabs, newlines)
    words := strings.Fields(trimmed)
    
    // Lowercase each word
    for i, word := range words {
        words[i] = strings.ToLower(word)
    }
    
    return words


}
