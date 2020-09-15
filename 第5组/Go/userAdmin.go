package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func forbidUser(c *gin.Context) {
	// ck, e := CheckCookie(c)
	// if e {
	// if isAdmin(ck.Uname) {
	tem := make(map[string]interface{})
	uid, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(uid, &tem)
	_, err := Db.Exec("UPDATE `user` SET `status`=0 WHERE uid=?", tem["uid"])
	if err != nil {
		c.JSON(202, ErrJson{
			Status: 0,
			Msg:    "服务器错误",
		})

	} else {
		c.JSON(200, ErrJson{
			Status: 1,
			Msg:    "ok",
		})
		// 	}
		// }
	}
	return
}
func QueryAllUser() []USER {
	res := make([]USER, 0)
	err := Db.Select(&res, "select * from user order by status DESC")
	if err != nil {
		fmt.Println(err)
	}
	return res
}
func getAllUser(c *gin.Context) {
	res := QueryAllUser()
	r := &ResJson{
		Status: 1,
		Msg:    "ok",
	}
	for _, v := range res {
		r.Data = append(r.Data, v)
	}
	c.JSON(200, r)
}
func resetUser(c *gin.Context) {
	t := make(map[string]interface{})
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &t)
	if t["uid"] != nil {
		_, err := Db.Exec("UPDATE `user` SET upasswd=\"123456\",paywd=\"\" WHERE uid=?", t["uid"])
		if err != nil {
			fmt.Println(err)
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "服务器错误，请稍后再试",
			})
		} else {
			c.JSON(200, &ErrJson{
				Status: 1,
				Msg:    "重置成功，新的登录密码为123456",
			})
		}

	} else {
		c.JSON(202, &ErrJson{
			Status: 0,
			Msg:    "请求不规范，请包含uid",
		})
	}
}
func updateuser(c *gin.Context) {
	t := make(map[string]interface{})
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &t)
	if t["uid"] != nil {
		if t["email"] != nil && t["phone"] != nil && t["yue"] != nil {
			_, err := Db.Exec("UPDATE `user` SET email=?,phone=?,yue=? where uid=?", t["email"], t["phone"], t["yue"], t["uid"])
			if err != nil {
				fmt.Println(err)
				c.JSON(202, &ErrJson{
					Status: 0,
					Msg:    "服务器错误，请稍后再试",
				})
			} else {
				c.JSON(200, &ErrJson{
					Status: 1,
					Msg:    "ok",
				})
			}
		}
	} else {
		c.JSON(202, &ErrJson{
			Status: 0,
			Msg:    "请求不规范，请包含uid",
		})
	}
}

// func main() {
// 	initDb()
// 	r := gin.Default()
// 	r.POST("/api/deluser", forbidUser)
// 	r.GET("/api/getalluser", getAllUser)
// 	r.POST("/api/resetuser", resetUser)
// 	r.POST("/api/updateuser", updateuser)
// 	r.Run(":9001")
// }
