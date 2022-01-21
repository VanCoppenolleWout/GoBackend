package reviews

import (
	database "github.com/VanCoppenolleWout/GoBackend/internal/pkg/db/mysql"
	"github.com/VanCoppenolleWout/GoBackend/internal/users"
	"log"
)

type Review struct {
	ID       string `json:"id"`
	Review   string `json:"review"`
	Date     string `json:"date"`
	Likes    string `json:"likes"`
	Comments string `json:"comments"`
	User     *users.User  `json:"user"`
}

func (review Review) Save() int64 {
	statement, err := database.Db.Prepare("INSERT INTO Reviews(Review, Date, Likes, Comments) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := statement.Exec(review.Review, review.Date, review.Likes, review.Comments)
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