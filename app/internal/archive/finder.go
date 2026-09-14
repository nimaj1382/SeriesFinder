package archive

import (
	"TheCollector/app/internal/movie"
	"bytes"
	"io"
	"net/http"
)

func Crawl(url string, depth int, movieService *movie.Service) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	body = body[bytes.Index(body, []byte("<tbody")):]

	for start := bytes.Index(body, []byte("href=\"")); start != -1; start = bytes.Index(body, []byte("href=\"")) {
		trimed := body[start+6:]
		end := bytes.Index(trimed, []byte("\""))
		link := string(trimed[:end-1])
		if trimed[end-1] != '/' {
			body = trimed[end+1:]
			continue
		}
		body = trimed[end+1:]
		if link == ".." {
			continue
		}
		if depth > 0 {
			err := Crawl(url+link+"/", depth-1, movieService)
			if err != nil {
				return err
			}
		} else {
			newMovie := movie.Movie{
				Title: link,
				URL:   url + link,
			}
			err := movieService.Add(&newMovie)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
