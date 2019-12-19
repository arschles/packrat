package main

type decoded struct {
	Regions []region `json:"regions"`
}

func (d decoded) allText() []string {
	ret := []string{}
	for _, region := range d.Regions {
		for _, line := range region.Lines {
			for _, word := range line.Words {
				ret = append(ret, word.Text)
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
