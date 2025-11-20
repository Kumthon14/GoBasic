package Models

type Fileupload struct {
	Id         uint   `gorm:"column:id"`
	Filename   string `gorm:"column:filename"`
	Filepath   string `gorm:"column:filepath"`
	CreateTime string `gorm:"column:createtime"`
}

type ResponseUploadFile struct {
	Message  string `json:"msg"`
	Code     string `json:"code"`
	Filepath string `json:"filepath"`
	Fileid   int    `json:"fileid"`
}

func (b *Fileupload) TableName() string {
	return "fileupload"
}
