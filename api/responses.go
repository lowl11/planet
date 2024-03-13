package api

type justOK struct {
	Status string `json:"status"`
}

type createdWithID struct {
	ID any `json:"id"`
}
