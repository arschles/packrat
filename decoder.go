package main

import (
	"context"

	"github.com/arschles/receipt-classifier/file"
	"github.com/arschles/receipt-classifier/matchers"
)

func walkDir(ctx context.Context, dir string) ([]decoded, error) {
	// diskFS, err := fs.NewFS(dir)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil

}

type decoded struct {
	Regions []region `json:"regions"`
}

func (d decoded) allWords() file.Words {
	ret := []file.Word{}
	for _, region := range d.Regions {
		for _, line := range region.Lines {
			for _, word := range line.Words {
				ret = append(ret, file.Word(word.Text))
			}
		}
	}
	return ret
}

type region struct {
	Lines []line `json:"lines"`
}

type line struct {
	Words []word `json:"words"`
}

type word struct {
	Text string `json:"text"`
}

// TODO: read this from a file
func companies() file.Words {
	return []file.Word{
		"marriott",
		"microsoft",
		"github",
		"privacy.com",
		"netlify",
		"statuscode",
		"papercall",
		"gophercon",
		"netflix",
		"software engineering daily",
	}
}

func runMatches(d decoded) (*file.Matches, error) {
	words := d.allWords()
	return &file.Matches{
		Companies: matchers.Companies(words, companies()),
	}, nil
}
