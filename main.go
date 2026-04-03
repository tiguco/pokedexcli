package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    
//    fmt.Println("=== Word REPL ===")
//    fmt.Println("Commands: 'exit', 'help'")
//    fmt.Println()
    
    for {
        fmt.Print("Pokedex > ")
        
        if !scanner.Scan() {
            break
        }
        
        input := strings.TrimSpace(scanner.Text())
	input = strings.ToLower(input)
        
        // Handle empty input
        if input == "" {
            continue
        }
        
        // Get first word
        words := strings.Fields(input)
        firstWord := words[0]
        
        // Command handling
//        switch firstWord {
//        case "exit", "quit":
//            fmt.Println("Exiting...")
//            return
//        case "help":
//            fmt.Println("  Type anything, I'll show you the first word")
//            fmt.Println("  Commands: exit, quit, help")
//            continue
//        }
        
        // Display the first word
	fmt.Printf("Your command was: %s\n", firstWord)
        
        // Optional: show if there were more words
//        if len(words) > 1 {
//            fmt.Printf("  (Ignored: %s)\n", strings.Join(words[1:], " "))
//        }
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Error:", err)
    }
}
