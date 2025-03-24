package entity

import "math/big"

type T_product_data struct {
	Prod_data_id big.Int `json:"prod_data_id"` //primary_key
	P_id         big.Int `json:"p_id"`
	P_name       string  `json:"p_name"`
	Value        big.Int `json:"value"`
	P_quantity   big.Int `json:"p_quantity"`
	DATE         string  `json:"date"`
} //go 엔티티는 구조체와 같다.
