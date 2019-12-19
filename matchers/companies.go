package matchers

import (
	"strings"

	"github.com/arschles/receipt-classifier/file"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

// Companies returns all of the companies that match any of the words in w
func Companies(words file.Words, companies file.Words) file.Words {
	addCompany, getCompanies := file.WordSet()
	for _, word := range words {
		for _, company := range companies {
			lowerWord := strings.ToLower(string(word))
			lowerCompany := strings.ToLower(string(company))
			if fuzzy.Match(lowerWord, lowerCompany) && len(lowerCompany) == len(lowerWord) {
				addCompany(file.Word(lowerCompany))
			}
		}
	}
	return getCompanies()
}
