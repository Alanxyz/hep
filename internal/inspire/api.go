package inspire

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "time"
	"strconv"
	"strings"
	"io"
)

type ITitle struct {
	Source string `json:"source"`
	Title string `json:"title"`
}

type IAbstract struct {
	Source string `json:"source"`
	Value string `json:"value"`
}

type IID struct {
	Value string `json:"value"`
	Schema string  `json:"schema"`
}

type IAuthor struct {
	FullName string `json:"full_name"`
	IDs []IID `json:"ids"`
}

type IArxiv struct {
	Value string `json:"value"`
}

type IInfo struct {
	Year int `json:"year"`
}

type IHit struct {
	ID string `json:"id"`
	Metadata struct {
		Titles []ITitle `json:"titles"`
		Abstracts []IAbstract `json:"abstracts"`
		Authors []IAuthor `json:"authors"`
		Arxiv []IArxiv `json:"arxiv_eprints"`
		Info []IInfo `json:"publication_info"`
	} `json:"metadata"`
}

type Response struct {
	Hits struct {
		Total int `json:"total"`
		Hits []IHit `json:"hits"`
	} `json:"hits"`
}

type Paper struct {
	ID int
	Title string
	Abstract string
	Authors []string
	AuthorIDs []string
	Date string
	Url string
}

// TODO validate id
func GetPaperByID(id int) Paper {
    url := fmt.Sprintf("https://inspirehep.net/api/literature/%d", id)

    client := &http.Client{
        Timeout: 100 * time.Second,
    }

    resp, err := client.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        panic(fmt.Sprintf("status inesperado: %s", resp.Status))
    }

    var dat IHit
    err = json.NewDecoder(resp.Body).Decode(&dat)
    if err != nil {
        panic(err)
    }

	paper := ProcessHit(dat)
	return paper
}

func ProcessHit(hit IHit) Paper {
	var paper Paper

	paper.ID, _ = strconv.Atoi(hit.ID)

	if len(hit.Metadata.Titles) > 0 {
		paper.Title = hit.Metadata.Titles[0].Title
	}
	if len(hit.Metadata.Abstracts) > 0 {
		paper.Abstract = hit.Metadata.Abstracts[0].Value
	}
	if len(hit.Metadata.Arxiv) > 0 {
		paper.Url = "https://arxiv.org/pdf/" + hit.Metadata.Arxiv[0].Value
	}
	if len(hit.Metadata.Info) > 0 {
		paper.Date = strconv.Itoa(hit.Metadata.Info[0].Year)
	}
	for _, author := range hit.Metadata.Authors {
		names := strings.Split(author.FullName, ", ")
		name := names[1] + " " + names[0]
		paper.Authors = append(paper.Authors, name)
	}
	for _, author := range hit.Metadata.Authors {
		for _, authorid := range author.IDs {
			if authorid.Schema == "INSPIRE BAI" {
				paper.AuthorIDs = append(paper.AuthorIDs, authorid.Value)
			}
		}
	}

	return paper
}

func Query(q string, size int, sort string) []Paper {

	if size < 1{
		size = 20
	}
	
	if sort == "" {
		sort = "mostrecient"
	}

    url := "https://inspirehep.net/api/literature?sort=" + sort + "&size=" + strconv.Itoa(size)+ "&q=" + url.QueryEscape(q) 

    client := &http.Client{
        Timeout: 100 * time.Second,
    }

    resp, err := client.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        panic(fmt.Sprintf("status inesperado: %s", resp.Status))
    }

    var dat Response
    err = json.NewDecoder(resp.Body).Decode(&dat)
    if err != nil {
        panic(err)
    }

	var papers []Paper
	for _, hit := range dat.Hits.Hits {
		paper := ProcessHit(hit)
		papers = append(papers, paper)
	}

	return papers
}

func FetchCitation(id int) string {
    url := "https://inspirehep.net/api/literature/" + strconv.Itoa(id) + "?format=bibtex"

    client := &http.Client{
        Timeout: 10 * time.Second,
    }

    resp, err := client.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        panic(fmt.Sprintf("status inesperado: %s", resp.Status))
    }
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(dat)
}
