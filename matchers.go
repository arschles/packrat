package main

import (
	"encoding/json"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

// MatchCompany returns all of the companies that match s over a given threshold
func matchCompanies(s string) []string {
	companies := []string{
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
	lower := strings.ToLower(s)
	ret := []string{}
	for _, c := range companies {
		if fuzzy.Match(lower, c) && (len(lower) == len(c)) {
			ret = append(ret, c)
		}
	}
	return ret
}

type matches struct {
	json.Marshaler
	companies []string `json:"companies"`
}

func (m *matches) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"companies": m.companies,
	})
}

func runMatches(d decoded) (*matches, error) {
	ret := &matches{}
	txt := d.allText()
	addCompany, getCompanies := companySet()
	for _, s := range txt {
		matchedCompanies := matchCompanies(s)
		addCompany(matchedCompanies...)
	}
	ret.companies = getCompanies()
	return ret, nil
}

func companySet() (func(...string), func() []string) {
	set := map[string]struct{}{}
	setter := func(companies ...string) {
		for _, c := range companies {
			set[c] = struct{}{}
		}
	}
	getter := func() []string {
		ret := []string{}
		for c := range set {
			ret = append(ret, c)
		}
		return ret
	}
	return setter, getter
}
