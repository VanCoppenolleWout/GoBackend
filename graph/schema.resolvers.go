package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/VanCoppenolleWout/GoBackend/graph/generated"
	"github.com/VanCoppenolleWout/GoBackend/graph/model"
	"github.com/VanCoppenolleWout/GoBackend/internal/movies"
)

func (r *mutationResolver) CreateMovie(ctx context.Context, input model.MovieInput) (*model.Movie, error) {
	var movie movies.Movie
	movie.Title = input.Title
	movie.Genre = input.Genre
	movie.ImgURL = input.ImgURL
	movie.Description = input.Description
	movie.ReleaseDate = input.ReleaseDate
	movie.Length = input.Length
	movie.Likes = input.Likes
	movie.Comments = input.Comments
	movieId := movie.Save()
	return &model.Movie{ID: strconv.FormatInt(movieId, 10), Title: movie.Title, Genre: movie.Genre, ImgURL: movie.ImgURL, Description: movie.Description, ReleaseDate: movie.ReleaseDate, Length: movie.Length, Likes: movie.Likes, Comments: movie.Comments}, nil
}

func (r *mutationResolver) Createuser(ctx context.Context, input model.UserInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	var movies []*model.Movie
	movies = append(movies, &model.Movie{
		Title:       "The conjuring",
		Genre:       "horror",
		ImgURL: "",
		Description: "super duper movie",
		ReleaseDate: 2020,
		Length:      "String",
		Likes:       50,
		Comments:    30,
	})
	return movies, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
