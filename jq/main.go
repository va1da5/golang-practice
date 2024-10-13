package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

var jsonData interface{}

func main() {
	reader := bufio.NewReader(os.Stdin)

	content := make([]byte, 0)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error reading from pipe: %s", err)
		}

		content = append(content, buffer[:n]...)
	}

	content = bytes.ReplaceAll(content, []byte{'\x00'}, []byte{})

	err := json.Unmarshal(content, &jsonData)

	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %s", err)
	}

	beautifiedJSON, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		log.Fatalf("Failed to beautify JSON: %v", err)
	}

	fmt.Printf("%s", string(beautifiedJSON))
}
