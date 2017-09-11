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

// panic(err) is just used for example purposes.
func main() {
	flag.Parse()
	// Key and tweak should be byte arrays. Put your key and tweak here.
	// To make it easier for demo purposes, decode from a hex string here.
	key, err := hex.DecodeString("FF4359D8D580AA4F7F036D6F04FC6A94")
	if err != nil {
		panic(err)
	}
	tweak, err := hex.DecodeString("D8E7920AFA330A73")
	if err != nil {
		panic(err)
	}

	// Create a new FF1 cipher "object"
	// 10 is the radix/base, and 8 is the tweak length.
	FF1, err := ff1.NewCipher(36, 8, key, tweak)
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
		original := strings.Split(record[*col], " ")
		tokenised := make([]string, len(original))

		for i, word := range original {
			// Call the encryption function on an example SSN
			tokenised[i], err = FF1.Encrypt(word)
			if err != nil {
				log.Println("failed to encrypt")
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
