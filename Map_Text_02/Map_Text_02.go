package main

import (
	"fmt"
	"strings"
)

//Store the words of a text on a map

func main() {
	text := `Bicycle, bicycle, bicycle
I want to ride my bicycle, bicycle, bicycle
I want to ride my bicycle
I want to ride my bike
I want to ride my bicycle
I want to ride it where I like
You say black, I say white
You say bark, I say bite
You say shark, I say hey man
Jaws was never my scene
And I don't like Star Wars`

	words := strings.Split(text, " ")
	dict := make(map[string]int)
	for _, word := range words {
		dict[word]++
	}

	for word, quantity := range dict {
		fmt.Println(word, " ->", quantity)
	}

}
