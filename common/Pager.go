package common

import "fmt"

//type StringString = CodeValue[string, string]

var Asc = StringString{
	Code:        "asc",
	ValueString: "asc",
}

var Desc = StringString{
	Code:        "desc",
	ValueString: "desc",
}

type Order struct {
	Field string `json:"field" form:"field"`
	Sort  string `json:"sort" form:"sort"`
}

type Pager struct {
	Page       int    `json:"page" form:"page"`
	Size       int    `json:"size" form:"size"`
	OrderField string `json:"order_field" form:"order_field"`
	Order      string `json:"order" form:"order"`
}

func (p *Pager) GetOrderField() string {

	if p.OrderField == "" {
		return "id"
	}

	// 默认降序
	return p.OrderField
}

func (p *Pager) GetOrder() string {

	// 默认降序
	var order = Desc
	if p.Order == Asc.Code {
		order = Asc
	}

	return fmt.Sprintf("%s %s", p.GetOrderField(), order.Code)
}

func (p *Pager) GetSafePage() int {

	if p.Page <= 0 {
		return 1
	}

	return p.Page
}

func (p *Pager) GetSafeSize() int {

	if p.Size <= 0 {
		return 10
	}

	if p.Size > 1000 {
		return 1000
	}

	return p.Size
}

func (p *Pager) GetSafeOffset() int {

	page := p.GetSafePage()
	size := p.GetSafeSize()

	return (page - 1) * size
}
