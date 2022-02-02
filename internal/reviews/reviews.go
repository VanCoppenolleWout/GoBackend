package reviews

import (
	"log"

	database "github.com/VanCoppenolleWout/GoBackend/internal/pkg/db/mysql"
	"github.com/VanCoppenolleWout/GoBackend/internal/users"
)

type Review struct {
	ID       string      `json:"id"`
	Review   string      `json:"review"`
	Date     string      `json:"date"`
	Likes    int      `json:"likes"`
	Comments int      `json:"comments"`
	User     *users.User `json:"user"`
}

func (review Review) Save() int64 {
	statement, err := database.Db.Prepare("INSERT INTO Reviews(Review, Date, Likes, Comments, UserID) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := statement.Exec(review.Review, review.Date, review.Likes, review.Comments, review.User.ID)
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

func GetAll() []Review {
	statement, err := database.Db.Prepare("SELECT L.id, L.review, L.date, L.likes, L.comments, L.UserId, U.Username from Reviews L inner join Users U on L.UserID = U.ID")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	rows, err := statement.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var reviews []Review
	var username string
	var id string
	for rows.Next() {
		var review Review
		err := rows.Scan(&review.ID, &review.Review, &review.Date, &review.Likes, &review.Comments, &id, &username)
		if err != nil {
			log.Fatal(err)
		}
		review.User = &users.User{
			ID:       id,
            Username: username,
		}
		reviews = append(reviews, review)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return reviews
}


func UpdateReviewLike(id *string) []Review {
	statement, err := database.Db.Prepare("UPDATE Reviews SET Likes = Likes + 1 WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	rows, err := statement.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var reviews []Review
	for rows.Next() {
		var review Review
		err := rows.Scan(&review.ID, &review.Review, &review.Date, &review.Likes, &review.Comments, &review.User.ID)
		if err != nil {
			log.Fatal(err)
		}
		reviews = append(reviews, review)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return reviews
 }
