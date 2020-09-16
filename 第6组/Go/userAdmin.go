package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func deluser(c *gin.Context) {
	res := make(map[string]interface{})
	uid, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(uid, &res)
	err := UpdateUserColumn("status", "0", fmt.Sprint(res["uid"]))
	if err == nil {
		c.JSON(200, &ResJson{
			Status: 1,
			Msg:    "删除成功！",
		})
	} else {
		c.JSON(200, &ResJson{
			Status: 0,
			Msg:    "删除失败，请重试！",
		})
	}
}
func queryallUser() []USER {
	res := make([]USER, 0)
	err := Db.Select(&res, "select * from user")
	if err != nil {
		fmt.Println(err)
	}
	return res
}
func selectalluser(c *gin.Context) {
	users := queryallUser()
	res := ResJson{}
	for _, v := range users {
		res.Data = append(res.Data, v)
	}
	res.fuchuzhi()
	c.JSON(200, res)
}

//Update NULL
func updateuser(c *gin.Context) {
	res := make(map[string]interface{})
	uid, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(uid, &res)
	if res["uid"] == nil {
		c.JSON(200, &ResJson{
			Status: 0,
			Msg:    "id为空，无法操作！",
		})
	} else {
		_, err := Db.Exec("UPDATE `user` SET email = ?, phone = ?, yue = ? WHERE uid = ?", res["email"], res["phone"], res["yue"], res["uid"])
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

func resetuser(c *gin.Context) {
	res := make(map[string]interface{})
	uid, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(uid, &res)
	if res["uid"] == nil {
		c.JSON(200, &ResJson{
			Status: 0,
			Msg:    "id为空，无法操作！",
		})
	} else {
		err := UpdateUserColumn("upasswd", "123456", fmt.Sprint(res["uid"]))
		if err == nil {
			c.JSON(200, &ResJson{
				Status: 1,
				Msg:    "重置成功，重置密码为：123456！",
			})
		} else {
			c.JSON(200, &ResJson{
				Status: 0,
				Msg:    "重置失败，请重试！",
			})
		}
	}

}

// func main() {
// 	initDb()
// 	t := queryallUser()
// 	fmt.Println(t)
// }
