package dto

type CardCate struct {
	ID       int    `json:"id"`
	CateName string `json:"cate_name"`
}

type CardPrefix struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Currency struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type CardAttrs struct {
	CateList    []CardCate   `json:"cate_list"`
	Currencies  []Currency   `json:"currencies"`
	CardPrefixs []CardPrefix `json:"card_prefixs"`
}
