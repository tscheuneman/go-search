package utils

type Status struct {
	Status  int    `json:"status"`  // low-level runtime error
	Message string `json:"message"` // http response status code
}
