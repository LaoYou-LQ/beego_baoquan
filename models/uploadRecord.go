package models

import (
	"DataCertProject/db"
)

/*
	上传文件的结构体
*/
type UploadRecord struct {
	Id        int
	FileName  string
	FileSize  int64
	FileCert  string
	FileTitle string
	CertTime  int64
	Phone     string
}

func (u UploadRecord) SeveRecord() (int64, error) {
	r, err := db.Db.Exec("insert into upload_record (file_name,file_size,file_cert,file_title,cert_time,phone)"+
		"values(?,?,?,?,?,?)",
		u.FileName, u.FileSize, u.FileCert, u.FileTitle, u.CertTime, u.Phone)
	//fmt.Println("111",err)
	if err != nil {
		return -1, err
	}
	id, err := r.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}

/*
读取数据库中phone用户对于的所有认证数据
*/
func QueryRecordbyPhone(phone string) ([]UploadRecord, error) {
	ros, err := db.Db.Query("select id, file_name ,file_size ,file_cert , file_title ,cert_time, phone from upload_record where phone =?", phone)
	if err != nil {
		return nil, err
	}
	records := make([]UploadRecord, 0)
	for ros.Next() {
		var record UploadRecord
		err := ros.Scan(&record.Id, &record.FileName, &record.FileSize, &record.FileCert, &record.FileTitle, &record.CertTime, &record.Phone)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records,nil
}
