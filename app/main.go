package main

import (
	"TheCollector/app/internal/archive"
	"TheCollector/app/internal/db"
	"TheCollector/app/internal/movie"
	"fmt"
)

func main() {
	database, err := db.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	movieRepo := movie.NewRepository(database)
	movieService := movie.NewService(movieRepo)

	archiveRepo := archive.NewRepository(database)
	archiveService := archive.NewService(archiveRepo)

	for {
		fmt.Println("1. Search for a movie\n2. Add a archive\n3. Update all archives\n4. Exit")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter movie title:")
			var title string
			fmt.Scanln(&title)
			movies, err := movieService.Search(title)
			if err != nil {
				fmt.Println(err)
				continue
			}
			for _, movie := range movies {
				fmt.Printf("Title: %s, URL: %s\n", movie.Title, movie.URL)
			}
		case 2:
			fmt.Println("Enter archive URL:")
			var url string
			fmt.Scanln(&url)
			fmt.Println("Enter depth:")
			var depth int
			fmt.Scanln(&depth)
			archive := &archive.Archive{URL: url, Depth: depth}
			err := archiveService.Create(archive)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case 3:
			err := archiveService.UpdateAll()
			if err != nil {
				fmt.Println(err)
				continue
			}
		case 4:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
