package export

type In struct {
	FromDate   string `json:"from_date"`
	ToDate     string `json:"to_date"`
	Filename   string `json:"filename"`
	TypeReport string `json:"type_report"`
}
