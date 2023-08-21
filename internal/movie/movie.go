package movie

type Movie struct {
	ID         string
	Title      string
	Rating     string
	Cast       []string
	Director   string
	Poster     *string
	UserRating *int32
}

type MovieInput struct {
	Title      string
	Rating     string
	Cast       []string
	Director   string
	Poster     *string
	UserRating *int32
	Genres     []string
}
