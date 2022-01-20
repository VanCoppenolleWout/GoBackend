package movies

import (
	database "github.com/VanCoppenolleWout/GoBackend/internal/pkg/db/mysql"
	"log"
)

type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Genre       string `json:"genre"`
	ImgURL      string `json:"imgURL"`
	Description string `json:"description"`
	ReleaseDate int    `json:"releaseDate"`
	Length      string `json:"length"`
	Likes       int    `json:"likes"`
	Comments    int    `json:"comments"`
}

func (movie Movie) Save() int64 {
	statement, err := database.Db.Prepare("INSERT INTO Movies(Title, Genre, ImgURL, Description, ReleaseDate, Length, Likes, Comments) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := statement.Exec(movie.Title, movie.Genre, movie.ImgURL, movie.Description, movie.ReleaseDate, movie.Length, movie.Likes, movie.Comments)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
	log.Print("Row inserted")
	
	return id
}
