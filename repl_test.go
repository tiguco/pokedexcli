package main

import (
    "reflect"
    "testing"
)


func TestCleanInput(t *testing.T) {
    tests := map[string]struct {
        input string
        want  []string
    }{ 
        "simple":       {input: "  hello  world  ", want: []string{"hello", "world"}}, 
        "verysimple":       {input: "helloworld", want: []string{"helloworld"}}, 
        "simple2":       {input: "hello world", want: []string{"hello", "world"}}, 
        "simple3":       {input: "hello    world", want: []string{"hello", "world"}}, 
        "simple4":       {input: "hello    cruel          world", want: []string{"hello", "cruel", "world"}}, 
        "simple5":       {input: "hello    cruel          world    ", want: []string{"hello", "cruel", "world"}}, 
        "simple6":       {input: "     hello    cruel          world    ", want: []string{"hello", "cruel", "world"}}, 
        "simple7":       {input: "   hello cruel world ", want: []string{"hello", "cruel", "world"}}, 
        "med1":       {input: "Charmander Bulbasaur PIKACHU", want: []string{"charmander", "bulbasaur", "pikachu"}}, 
    }

    for name, tc := range tests {
        got := cleanInput(tc.input)
        if !reflect.DeepEqual(tc.want, got) {
            t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
        }
    }
}
