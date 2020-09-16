package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func queryallOrder() []OREDER {
	res := make([]OREDER, 0)
	err := Db.Select(&res, "select * from orders")
	if err != nil {
		fmt.Println(err)
	}
	return res
}
func getallorder(c *gin.Context) {
	orders := queryallOrder()
	res := ResJson{}
	for _, v := range orders {
		res.Data = append(res.Data, v)
	}
	res.fuchuzhi()
	c.JSON(200, res)
}

//Update NULL
func updateorder(c *gin.Context) {
	res := make(map[string]interface{})
	orid, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(orid, &res)
	if res["orid"] == nil {
		c.JSON(200, &ResJson{
			Status: 0,
			Msg:    "id为空，无法操作！",
		})
	} else {
		_, err := Db.Exec("UPDATE `order` SET status = ?, recvname = ?, phone = ?, detailed = ? WHERE orid = ?", res["status"], res["recvname"], res["phone"], res["detailed"], res["orid"])
		if err == nil {
			c.JSON(200, &ResJson{
				Status: 1,
				Msg:    "更新成功！",
			})
		} else {
			c.JSON(200, &ResJson{
				Status: 0,
				Msg:    "更新失败，请重试！",
			})
		}
	}
}
func main() {
	initDb()
	r := gin.Default()
	r.GET("/api/getallorder", getallorder)
	r.Run(":9001")
}
