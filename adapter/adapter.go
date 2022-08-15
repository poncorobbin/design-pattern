package main

import "fmt"

type Book struct {
	title  string
	author string
}

type Movie struct {
	title    string
	author   string
	duration int
}

// this is the adapter to convert Book and Movie to standard format
type CatalogAdapter interface {
	getTitle() string
}

type BookCatalogAdapter struct { // implement adapter for every class
	book Book
}

func (b BookCatalogAdapter) getTitle() string {
	return b.book.title + " by " + b.book.author
}

type MovieCatalogAdapter struct { // implement adapter for every class
	movie Movie
}

func (m MovieCatalogAdapter) getTitle() string {
	return fmt.Sprintf("%s by %s (%v)m",
		m.movie.title, m.movie.author, m.movie.duration)
}

func main() {
	listCatalog := []CatalogAdapter{}

	listCatalog = append(listCatalog, BookCatalogAdapter{book: Book{title: "Book 1", author: "Ponco"}})
	listCatalog = append(listCatalog, MovieCatalogAdapter{movie: Movie{title: "Movie 1", author: "Robbi", duration: 100}})

	for _, v := range listCatalog {
		fmt.Println(v.getTitle())
	}
}
