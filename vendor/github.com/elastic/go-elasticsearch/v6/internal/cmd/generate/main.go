package main

import (
	"github.com/elastic/go-elasticsearch/v6/internal/cmd/generate/commands"
	_ "github.com/elastic/go-elasticsearch/v6/internal/cmd/generate/commands/gensource"
	_ "github.com/elastic/go-elasticsearch/v6/internal/cmd/generate/commands/gentests"
)

func main() {
	commands.Execute()
}
