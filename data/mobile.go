package data

type Mobile struct {
	Id        string `json:"id"`
	Brand     string `json:"brand"`
	Model     string `json:"model"`
	Processor string `json:"processor"`
	Ram       string `json:"ram"`
	Storage   string `json:"storage"`
}