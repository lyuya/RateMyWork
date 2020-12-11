package model

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	UserType string `json:"userType"`
}

type File struct {
	Id       string `json:"id"`
	FileName string `json:"fileName"`
	FileType string `json:"fileType"`
}
