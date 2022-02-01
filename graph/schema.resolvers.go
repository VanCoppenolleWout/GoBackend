package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/VanCoppenolleWout/GoBackend/graph/generated"
	"github.com/VanCoppenolleWout/GoBackend/graph/model"
	"github.com/VanCoppenolleWout/GoBackend/internal/auth"
	"github.com/VanCoppenolleWout/GoBackend/internal/movies"
	"github.com/VanCoppenolleWout/GoBackend/internal/pkg/jwt"
	"github.com/VanCoppenolleWout/GoBackend/internal/reviews"
	"github.com/VanCoppenolleWout/GoBackend/internal/users"
)

func (r *mutationResolver) CreateMovie(ctx context.Context, input model.MovieInput) (*model.Movie, error) {
	var movie movies.Movie
	movie.Title = input.Title
	movie.Genre = input.Genre
	movie.ImgUrl = input.ImgURL
	movie.Description = input.Description
	movie.ReleaseDate = input.ReleaseDate
	movie.Length = input.Length
	movie.Likes = input.Likes
	movie.Comments = input.Comments
	movieId := movie.Save()
	return &model.Movie{ID: strconv.FormatInt(movieId, 10), Title: movie.Title, Genre: movie.Genre, ImgURL: movie.ImgUrl, Description: movie.Description, ReleaseDate: movie.ReleaseDate, Length: movie.Length, Likes: movie.Likes, Comments: movie.Comments}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	user.Create()
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) CreateReview(ctx context.Context, input model.ReviewInput) (*model.Review, error) {
	var review reviews.Review
	review.Review = input.Review
	review.Date = input.Date
	review.Likes = input.Likes
	review.Comments = input.Comments

	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Review{}, fmt.Errorf("acces denied")
	}

	review.User = user
	reviewId := review.Save()
	graphqlUser := &model.User{
		ID:       user.ID,
		Username: user.Username,
	}
	return &model.Review{ID: strconv.FormatInt(reviewId, 10), Review: review.Review, Date: review.Date, Likes: review.Likes, Comments: review.Comments, User: graphqlUser}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	correct := user.Authenticate()
	if !correct {
		return "", &users.WrongUsernameOrPasswordError{}
	}
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	var resultMovies []*model.Movie
	var dbMovies []movies.Movie = movies.GetAll()
	for _, movie := range dbMovies {
		resultMovies = append(resultMovies, &model.Movie{ID: movie.ID, Title: movie.Title, Genre: movie.Genre, ImgURL: movie.ImgUrl, Description: movie.Description, ReleaseDate: movie.ReleaseDate, Length: movie.Length, Likes: movie.Likes, Comments: movie.Comments})
	}
	return resultMovies, nil
}

func (r *queryResolver) Reviews(ctx context.Context) ([]*model.Review, error) {
	var resultReviews []*model.Review
	var dbReviews []reviews.Review = reviews.GetAll()
	for _, review := range dbReviews {
		graphqlUser := &model.User{
			ID:       review.User.ID,
			Username: review.User.Username,
		}
		resultReviews = append(resultReviews, &model.Review{ID: review.ID, Review: review.Review, Date: review.Date, Likes: review.Likes, Comments: review.Comments, User: graphqlUser})
	}
	return resultReviews, nil
}

func (r *queryResolver) MovieByID(ctx context.Context, id *string) ([]*model.Movie, error) {
	var resultMovies []*model.Movie
	var dbMovies []movies.Movie = movies.GetMovieById(id)
	for _, movie := range dbMovies {
		resultMovies = append(resultMovies, &model.Movie{ID: movie.ID, Title: movie.Title, Genre: movie.Genre, ImgURL: movie.ImgUrl, Description: movie.Description, ReleaseDate: movie.ReleaseDate, Length: movie.Length, Likes: movie.Likes, Comments: movie.Comments})
	}
	return resultMovies, nil
}

func (r *queryResolver) MovieByGenre(ctx context.Context, genre *string) ([]*model.Movie, error) {
	var resultMovies []*model.Movie
	var dbMovies []movies.Movie = movies.GetMoviesByGenre(genre)
	for _, movie := range dbMovies {
		resultMovies = append(resultMovies, &model.Movie{ID: movie.ID, Title: movie.Title, Genre: movie.Genre, ImgURL: movie.ImgUrl, Description: movie.Description, ReleaseDate: movie.ReleaseDate, Length: movie.Length, Likes: movie.Likes, Comments: movie.Comments})
	}
	return resultMovies, nil
}

func (r *queryResolver) MovieByYear(ctx context.Context, releaseDate *int) ([]*model.Movie, error) {
	var resultMovies []*model.Movie
	var dbMovies []movies.Movie = movies.GetMoviesByYear(releaseDate)
	for _, movie := range dbMovies {
		resultMovies = append(resultMovies, &model.Movie{ID: movie.ID, Title: movie.Title, Genre: movie.Genre, ImgURL: movie.ImgUrl, Description: movie.Description, ReleaseDate: movie.ReleaseDate, Length: movie.Length, Likes: movie.Likes, Comments: movie.Comments})
	}
	return resultMovies, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
