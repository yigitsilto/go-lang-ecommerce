package dto

type UserMeModel struct {
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	Phone          string  `json:"phone"`
	Point          float64 `json:"point"`
	VergiDairesi   string  `json:"vergi_dairesi"`
	VergiNo        string  `json:"vergi_no"`
	Email          string  `json:"email"`
	CompanyGroupId float64 `json:"company_group_id"`
}
