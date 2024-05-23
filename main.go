package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/nbd-wtf/go-nostr"
)

func main() {
	enc := json.NewEncoder(os.Stdout)
	dec := json.NewDecoder(os.Stdin)

	envelope := [3]any{"EVENT", "s", nil}
	for {
		var ev nostr.Event
		err := dec.Decode(&ev)
		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}
		envelope[2] = ev
		err = enc.Encode(envelope)
		if err != nil {
			os.Exit(2)
			break
		}
	}
}
