package dto

type UserBlockchain struct {
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	Name      string `json:"name"`
}

type UserBlockchainCreate struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
