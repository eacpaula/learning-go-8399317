package main

import (
	"fmt"
	"sort"
)

/*
 * Reference: Chapter 3 > Lesson 3
 *
 * Store ordered values in array
 */

func main() {
	states := make(map[string]string)
	fmt.Println(states)

	states["NY"] = "New York"
	states["NJ"] = "New Jersey"
	states["CA"] = "California"

	fmt.Println(states)

	newYork := states["NY"]
	fmt.Println(newYork)

	delete(states, "NJ")

	fmt.Println(states)

	states["WA"] = "Washington"
	fmt.Println(states)

	for key, value := range states {
		fmt.Println(key, value)
	}

	keys := make([]string, 0, len(states))
	for key := range states {
		keys = append(keys, key)
	}
	fmt.Println(keys)

	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, states[key])
	}
}
