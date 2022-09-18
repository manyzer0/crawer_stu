package model

type User struct {
	NickName  string `json:"nickname"`
	Age       int    `json:"age"`
	Sarary    int    `json:"sarary"`
	Status    string `json:"status"`
	Height    int    `json:"height"`
	Education string `json:"education"`
}
