package pojo

type PojoPagination struct {
	ListData interface{} `json:"data"`
	Count    int         `json:"count"`
}
