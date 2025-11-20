package Models

import (
	"GoPractice/Config"
	"mime/multipart"
)

func UploadFile(file *multipart.FileHeader, filepath string) (fileId int, err error) {
	tsql := "INSERT INTO [dbo].[fileupload] (filename,filepath,createtime) OUTPUT INSERTED.id VALUES ('" + file.Filename + "','" + filepath + "',GETDATE())"
	result := Config.DB.Raw(tsql)
	result.Scan(&fileId)
	return fileId, result.Error
}

func GetUploadLists(fileupload *[]Fileupload) error {
	if err := Config.DB.Find(fileupload).Error; err != nil {
		return err
	}
	return nil
}

func SearchById(FileId string) (res Fileupload, err error) {
	err = Config.DB.Where("id = ?", FileId).First(&res).Error
	return res, err
}
