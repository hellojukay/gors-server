package database

type Record struct {
	Filename string `json:"filename"`
	FilePath string `json:"filepath"`
}

func SaveRecord(filename string, filePath string) error {
	db.Save(Record{
		FilePath: filePath,
		Filename: filename,
	})
	return db.Error
}

func AllRecords() ([]Record, error) {
	var records = []Record{}
	db.Find(&records)
	return records, db.Error
}
