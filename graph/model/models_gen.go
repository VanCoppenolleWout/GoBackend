// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Genre       string `json:"genre"`
	ImgURL      string `json:"imgUrl"`
	Description string `json:"description"`
	ReleaseDate int    `json:"releaseDate"`
	Length      string `json:"length"`
	Likes       int    `json:"likes"`
	Comments    int    `json:"comments"`
}

type MovieInput struct {
	Title       string `json:"title"`
	Genre       string `json:"genre"`
	ImgURL      string `json:"imgUrl"`
	Description string `json:"description"`
	ReleaseDate int    `json:"releaseDate"`
	Length      string `json:"length"`
	Likes       int    `json:"likes"`
	Comments    int    `json:"comments"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type Review struct {
	ID       string `json:"id"`
	Review   string `json:"review"`
	Date     string `json:"date"`
	Likes    int    `json:"likes"`
	Comments int    `json:"comments"`
	User     *User  `json:"user"`
}

type ReviewInput struct {
	Review   string `json:"review"`
	Date     string `json:"date"`
	Likes    int    `json:"likes"`
	Comments int    `json:"comments"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
