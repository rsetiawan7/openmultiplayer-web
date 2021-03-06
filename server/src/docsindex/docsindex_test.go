package docsindex

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/openmultiplayer/web/server/src/config"
)

func TestIndex_Build(t *testing.T) {
	idx, err := New(config.Config{
		DocsSourcesPath: "../../../docs/scripting",
		DocsIndexPath:   "docsearch.bleve",
	})
	if err != nil {
		panic(err)
	}

	res1, err := idx.Search("position")
	if err != nil {
		panic(err)
	}

	b, _ := json.MarshalIndent(res1, "", "  ")

	fmt.Printf("result: %v\n", string(b))
}
