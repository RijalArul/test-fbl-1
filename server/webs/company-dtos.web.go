package webs

type CompanyDTO struct {
	CompanyName string `json:"company_name" form:"company_name"`
	CompanyCode string `json:"company_code" form:"company_code"`
}

type RepsonseCompanyBody struct {
	Name string `json:"company_name" form:"company_name"`
	Code string `json:"company_code" form:"company_code"`
}
