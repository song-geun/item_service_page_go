package dto

import (
	"database/sql"
	"encoding/json"
	_ "io"
	_ "io/ioutil"
	"item_service_page/src/datastruct"
	"item_service_page/src/entity"
	"item_service_page/src/repository5"
	"log"
	_ "net/http"
	"strconv"
	"strings"
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

func ProductInsert(str string, st int, en int) {
	datastruct.New()
	var str2 = str
	var d entity.T_product
	for i := st; i <= en; i++ {
		if strings.Compare(str[i:i+1], "{") == 0 {
			datastruct.Stack.Push_back(i)
		}
		if strings.Compare(str[i:i+1], "}") == 0 {
			str3 := str2[datastruct.Stack.Top() : i+1]
			err := json.Unmarshal([]byte(str3), &d)
			if err != nil {
				log.Fatal(err)
			}
			repository5.RepositoryInstance.USP_ProductManage("I1", d.P_id, d.P_name, d.Value, d.Quantity)
			datastruct.Stack.Del()
		}
	}
}

func ProductDto(sq *sql.Rows) []entity.T_product {
	//var sq = repository5.RepositoryInstance.USP_ProductManage("S1", 0, " ", 0, 0)
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

// []byte는 *void
// any는 *void
