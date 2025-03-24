package dto

import (
	_ "io"
	_ "io/ioutil"
	"item_service_page/src/entity"
	"item_service_page/src/repository5"
	_ "net/http"
	"strconv"
)

var ans string = ""

func solve(list []entity.T_product, num int) {
	if len(list) <= num {
		return
	}
	ans += "{"
	ans += "\"p_id\" : "
	ans += strconv.Itoa(list[num].P_id)
	ans += ", "
	ans += "\"p_name\" : "
	ans += list[num].P_name
	ans += ", "
	ans += "\"value\" : "
	ans += strconv.Itoa(list[num].Value)
	ans += ", "
	ans += "\"quantity\" : "
	ans += strconv.Itoa(list[num].Quantity)
	ans += "}"
	if len(list) > num+1 {
		ans += ","
	}
	solve(list, num+1)
}

func FindAllD() []entity.T_product {
	var sq = repository5.RepositoryInstance.USP_ProductManage("S1", 0, " ", 0, 0)
	mylist := []entity.T_product{}

	for sq.Next() {
		p_id := 0
		p_name := " "
		p_value := 0
		p_quantity := 0
		sq.Scan(&p_id, &p_name, &p_value, &p_quantity)
		mylist = append(mylist, entity.T_product{p_id, p_name, p_value, p_quantity})
	}
	if len(mylist) < 0 {
		return nil
	}
	ans = "{"
	solve(mylist, 0)
	ans += "}"
	return mylist
}
