package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Declare as nil (zero value for map)
var mcmd map[string]cliCommand

//var mcmd = map[string]cliCommand{
//    "exit": {
//        name:        "exit",
//        description: "Exit the Pokedex",
//        callback:    commandExit,
//    },
//    "help": {
//        name:        "help",
//	description: "Displays a help message",
//        callback:    commandHelp,
//    },
//}


func commandExit() error {
    // print a goodbye message
    // then exit the program
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp() error {
	fmt.Println("Usage:")
        fmt.Println("")
        fmt.Println("")
	for key, val := range mcmd {
		fmt.Printf("%s: %s\n", key, val.description)
	}
	return nil
}


func myinit() error {
    // Initialize in init function
    mcmd = make(map[string]cliCommand)
    mcmd["exit"] = cliCommand{
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    }
    mcmd["help"] = cliCommand{
        name:        "help",
        description: "Displays a help message",
        callback:    commandHelp,
    }
    return nil
}

func main() {
    fmt.Println("Welcome to the Pokedex!")
    myinit()
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
	coiso, ok := mcmd[firstWord]
	if ok {
		coiso.callback()
	}else{
		fmt.Println("Unknown command")
	}

        
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
	//fmt.Printf("Your command was: %s\n", firstWord)
        
        // Optional: show if there were more words
//        if len(words) > 1 {
//            fmt.Printf("  (Ignored: %s)\n", strings.Join(words[1:], " "))
//        }
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Error:", err)
    }
}
