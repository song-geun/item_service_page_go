package entity

type T_product struct {
	/*@Column
	private String p_name;
	@Column
	private Long value;
	@Column
	private Long quantity;*/

	P_id     int    `json:"p_id"`
	P_name   string `json:"p_name"`
	Value    int    `json:"value"`
	Quantity int    `json:"quantity"`
}

var T_ProducT T_product
