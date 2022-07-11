package dto

type Files struct {
	CreatedAt string   `json:"createdAt"`
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Owner     string   `json:"owner"`
	Customers []string `json:"customers"`
	Url       string   `json:"url"`
	Size      int      `json:"size"`
	Type      string   `json:"type"`
}

type FilesCreateDto struct {
	Name      string   `json:"name"`
	Owner     string   `json:"owner"`
	Customers []string `json:"customers"`
	Url       string   `json:"url"`
	Size      int      `json:"size"`
	Type      string   `json:"type"`
}

type FilesUpdateDto struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Size int    `json:"size"`
	Type string `json:"type"`
}

type FileTransferDto struct {
	UserId string `json:"userId"`
}

type FileGetAllQuery struct {
	Owner string `json:"owner"`
}
