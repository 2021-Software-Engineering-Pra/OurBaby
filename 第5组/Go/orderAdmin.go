package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func updateOrder(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	t := make(map[string]interface{}, 0)
	e := json.Unmarshal(data, &t)
	if e != nil {
		fmt.Println(e)
		c.JSON(202, &ErrJson{
			Status: 0,
			Msg:    "请求错误,请稍后重试",
		})
		return
	}
	orid := t["orid"]
	if orid == nil {
		c.JSON(202, &ErrJson{
			Status: 0,
			Msg:    "缺少订单号",
		})
	} else {
		recv := t["recvname"]
		phone := t["phone"]
		detail := t["detail"]
		status := t["status"]
		if recv != nil && phone != nil && detail != nil {
			_, err := Db.Exec("UPDATE orders SET recvname=?, phone=?,detailed=? WHERE orid=?", recv, phone, detail, orid)
			if err != nil {
				fmt.Println(err)
				c.JSON(202, &ErrJson{
					Status: 0,
					Msg:    "服务器错误,请稍后重试",
				})
			}
		}
		if status != nil {
			_, err := Db.Exec("UPDATE orders SET status=? WHERE orid=?", status, orid)
			if err != nil {
				fmt.Println(err)
				c.JSON(202, &ErrJson{
					Status: 0,
					Msg:    "服务器错误,请稍后重试",
				})
			}
		}

		c.JSON(200, &ErrJson{
			Status: 1,
			Msg:    "ok",
		})

	}
}
func getAllOrder(c *gin.Context) {
	data := make([]OREDER, 0)
	err := Db.Select(&data, "SELECT * FROM orders LIMIT 1000")
	if err != nil {
		fmt.Println(err)
		c.JSON(202, &ErrJson{
			Status: 0,
			Msg:    "服务器错误，请稍后重试",
		})
	} else {
		res := ResJson{}
		res.fuchuzhi()
		for _, v := range data {
			res.Data = append(res.Data, v)
		}
		c.JSON(200, &res)
	}
}

// func main() {
// 	initDb()
// 	r := gin.Default()
// 	r.POST("/api/updateorder", updateOrder)
// 	r.GET("/api/getallorder", getAllOrder)
// 	r.Run(":9001")
// }
