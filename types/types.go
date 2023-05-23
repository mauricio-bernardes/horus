package types

type Service struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Subscribed struct {
	Url   string   `json:"url"`
	Names []string `json:"names"`
}

type Services struct {
	Services []string `json:"services"`
}
