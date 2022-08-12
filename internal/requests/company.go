package requests

type CreateCompanyRequest struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}
