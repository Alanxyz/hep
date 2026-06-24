package pdf

import (
	"os"
	"net/http"
	"io"
	"fmt"
	"hep/internal/inspire"
)

func downloadFile(filepath string, url string) (err error) {
  out, err := os.Create(filepath)
  if err != nil  {
    return err
  }
  defer out.Close()

  resp, err := http.Get(url)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    return fmt.Errorf("bad status: %s", resp.Status)
  }

  _, err = io.Copy(out, resp.Body)
  if err != nil  {
    return err
  }

  return nil
}

func OpenPaper(id int) {
	paper := inspire.GetPaperByID(id)

	if paper.Url == "" {
		fmt.Println("PDF file not avalible")
		return
	}

	paperPath := fmt.Sprintf("/tmp/%d.pdf", id)
	err := downloadFile(paperPath, paper.Url)

	if err != nil {
		panic(err)
	}

	OpenFile(paperPath)
}
