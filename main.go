package main

import (
	"bufio"
	"encoding/csv"
	"encoding/hex"
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"github.com/capitalone/fpe/ff1"
)

var fname = flag.String("fname", "", "input filename")
var col = flag.Int("column", 0, "column to tokenise")
var keyString = flag.String("key", "FF4359D8D580AA4F7F036D6F04FC6A94", "key for the FF1 algorithm")
var tweakString = flag.String("tweak", "D8E7920AFA330A73", "tweak for the FF1 algorithm")

// panic(err) is just used for example purposes.
func main() {
	flag.Parse()
	key, err := hex.DecodeString(*keyString)
	if err != nil {
		panic(err)
	}
	tweak, err := hex.DecodeString(*tweakString)
	if err != nil {
		panic(err)
	}

	// Create a new FF1 cipher "object"
	// 62 is the radix/base, and 8 is the tweak length.
	// 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
	FF1, err := ff1.NewCipher(62, 8, key, tweak)
	if err != nil {
		panic(err)
	}

	f, _ := os.Open(*fname)
	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	r.TrimLeadingSpace = true

	// Write to stdout
	w := csv.NewWriter(os.Stdout)

	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		original := record[*col]
		originalSplit := strings.Split(original, " ")
		log.Println(originalSplit)
		tokenised := make([]string, len(originalSplit))

		for i, word := range originalSplit {
			// Call the encryption function on an example SSN
			tokenised[i], err = FF1.Encrypt(word)
			if err != nil {
				log.Println("failed to encrypt")
				log.Println(word)
				panic(err)
			}
		}

		// replace the col with the ciphertext
		record[*col] = strings.Join(tokenised, " ")
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}

	}
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

}
