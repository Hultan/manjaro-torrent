package manjaro

import (
	"bufio"
	"github.com/PuerkitoBio/goquery"
	"os"
	"strings"
)

const nonBreakingSpace = string(160)

type Manjaro struct {
	Distributions map[string]Distribution
}

type Distribution struct {
	Name    string
	Type    string
	Version string
}

func NewDistribution(name, distributionType, version string) Distribution {
	d := Distribution{name, distributionType, version}
	return d
}

func New() *Manjaro {
	compare := new(Manjaro)
	return compare
}

func (m *Manjaro) ParseHtml(path string) (err error) {
	buf, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func() {
		err = buf.Close()
	}()

	r := bufio.NewReader(buf)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return err
	}

	m.internalGetDistributions(doc)

	return nil
}

func (m *Manjaro) internalGetDistributions(doc *goquery.Document) {
	m.Distributions = make(map[string]Distribution, 1)

	// Find the official distributions
	doc.Find("ul.dropdown-menu ul.dropdown-menu[aria-labelledby=\"Official\"] a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the distribution
		dist := m.getDistribution(s.Text(), "Official")
		m.Distributions[dist.Name] = dist
	})
	// Find the official distributions
	doc.Find("ul.dropdown-menu ul.dropdown-menu[aria-labelledby=\"Community\"] a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the distribution
		dist := m.getDistribution(s.Text(), "Community")
		m.Distributions[dist.Name] = dist
	})
}

func (m *Manjaro) getDistribution(text, distributionType string) Distribution {
	index := strings.LastIndex(text, nonBreakingSpace)
	name := strings.Trim(text[0:index], nonBreakingSpace)
	version := strings.Trim(text[index:], nonBreakingSpace)
	return NewDistribution(name, distributionType, version)
}
