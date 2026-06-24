package inspire

import (
	"fmt"
	"strings"
	"github.com/logrusorgru/aurora/v4"
)


func QueryResults(q string, omitAbstract bool) {
	papers := Query(q, 0, "")
	for _, paper := range papers {
		fmt.Println("[", aurora.Green(paper.ID), "]", aurora.Blue(paper.Title))
		authors := strings.Join(paper.Authors, ", ")
		fmt.Println(aurora.Cyan(authors))
		fmt.Println(paper.Date)
		if !omitAbstract {
			fmt.Println(paper.Abstract)
		}
		if paper.Url != "" {
			fmt.Println(aurora.Yellow("Preprint avalible"))
		}
		fmt.Println()
	}
}

func FeedMenu(OnlyFollowers, omitAbstract  bool) {
	authors := []string{
		"O.Loaiza.Brito.1",
		"N.G.Cabo.Bizet.1",
		"Victor.Lopez.1",
		"Alan.Josue.Sierra.Torres.1",
		"G.PerezCuellar.1",
		"C.A.Vaquera.Araujo.1",
	}

	var q string
	if OnlyFollowers {
		q = "author: " + strings.Join(authors, " OR ")
	} else {
		q = ""
	}

	QueryResults(q, omitAbstract)
}

func ShowPaper(id int) {
	paper := GetPaperByID(id)
	fmt.Println(aurora.Blue(paper.Title))
	fmt.Println(paper.Abstract)
}

