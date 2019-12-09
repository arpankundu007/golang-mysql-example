package models

type Mobile struct {
	Id        string `json:"id"`
	Brand     string `json:"brand"`
	Model     string `json:"model"`
	Processor string `json:"processor"`
	Ram       string `json:"ram"`
	Storage   string `json:"storage"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}
