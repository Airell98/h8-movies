package dto

type NewMovieRequest struct {
	Title    string `json:"title" valid:"required~title cannot be empty" example:"Jelangkung"`
	ImageUrl string `json:"imageUrl" valid:"required~image url cannot be empty" example:"http://imageurl.com"`
	Price    int    `json:"price" valid:"required~price cannot be empty" example:"20000"`
}

type NewMovieResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
