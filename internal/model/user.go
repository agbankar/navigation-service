package model

type EmptyStruct struct {
}

var VoidStruct EmptyStruct

type User struct {
	UserId string `json:"userId"`
	Url    string `json:"url"`
}

type PageDetails struct {
	UserIds map[string]EmptyStruct
	Counter int
}
type ApiResponse struct {
	Page           string `json:"page,omitempty"`
	ErrorMessage   string `json:"errorMessage,omitempty"`
	SuccessMessage string `json:"successMessage,omitempty"`
	UniqueVisits   *int   `json:"uniqueVisits,omitempty"`
}
