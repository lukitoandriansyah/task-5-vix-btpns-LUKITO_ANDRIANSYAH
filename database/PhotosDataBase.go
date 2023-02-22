package database

type PhotosUpdateData struct {
	ID       uint   `json:"id" form:"id"`
	Title    string `json:"title" form:"title"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photoUrl" form:"photoUrl"`
	UserId   uint   `json:"userId,omitempty" form:"userId, omitempty"`
}

type PhotoCreateData struct {
	Title    string `json:"title" form:"title"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photoUrl" form:"photoUrl"`
	UserId   uint   `json:"userId,omitempty" form:"userId, omitempty"`
}
