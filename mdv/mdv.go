package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/gholt/blackfridaytext"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Please specify *one* file name to render")
	}
	
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Could not read %s: %v\n", os.Args[1], err)
	}

	// Hard coded configuration
	opt := &blackfridaytext.Options{
		Color: true,
		HeaderPrefix: []byte("-["),
		HeaderSuffix: []byte("]-"),
	}

	metadata, output := blackfridaytext.MarkdownToText(data, opt)
	for _, item := range metadata {
		name, value := item[0], item[1]
		os.Stdout.WriteString(name)
		os.Stdout.WriteString(":\n    ")
		os.Stdout.WriteString(value)
		os.Stdout.WriteString("\n")
	}
	os.Stdout.WriteString("\n")
	os.Stdout.Write(output)
	os.Stdout.WriteString("\n")
}
