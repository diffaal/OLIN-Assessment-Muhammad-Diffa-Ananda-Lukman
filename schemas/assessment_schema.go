package schemas

import "errors"

type CodeOneRequest struct {
	Nums   []int `json:"nums"`
	Target int   `json:"target"`
}

type CodeTwoRequest struct {
	Nums []int `json:"nums"`
}

type ResponseBase struct {
	Result interface{} `json:"result"`
	Status string      `json:"status"`
	Detail string      `json:"detail"`
}

func (cor *CodeOneRequest) ValidatePayload() error {
	if len(cor.Nums) < 2 {
		return errors.New("Nums array length cannot be less than 2")
	}
	return nil
}
func (ctr *CodeTwoRequest) ValidatePayload() error {
	if len(ctr.Nums) < 3 {
		return errors.New("Nums array length cannot be less than 3")
	}
	return nil
}
