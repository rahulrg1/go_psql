// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateJobListingInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Comapny     string `json:"comapny"`
	URL         string `json:"url"`
}

type DeleteJobResponse struct {
	DeleteJobID string `json:"deleteJobId"`
}

type JobListing struct {
	ID          string `json:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Comapny     string `json:"comapny"`
	URL         string `json:"url"`
}

type UpdateJobListingInput struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	URL         *string `json:"url,omitempty"`
}
