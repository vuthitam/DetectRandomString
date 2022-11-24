package corpus

import (
	"bufio"
	"log"
	"os"
)

func HashTable() map[string]int {
	corpus := make(map[string]int)

	// Buld corpus hash table from a book
	wordlist, err := os.Open("./resource/english_world.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer wordlist.Close()
	scanner := bufio.NewScanner(wordlist)
	for scanner.Scan() {
		word := scanner.Text()

		if _, ok := corpus[word]; ok {
			//exist in corpus
			corpus[word]++
		} else {
			//not exist in corpus
			corpus[word] = 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return corpus
}
