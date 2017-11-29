package model

type ReqUpdateUser struct {
	Password string `json:"password"`
}

type ReqLoginBody struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ReqNewRoomBody struct {
	Room string `json:"room"`
	User string `json:"user"`
	Max  int    `json:"max"`
}

type ReqNewRoomTokenBody struct {
	Room    string `json:"room"`
	User    string `json:"user"`
	Version string `json:"version"`
}
