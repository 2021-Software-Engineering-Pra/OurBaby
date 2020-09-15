package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wenzhenxi/gorsa"
)

type ADMIN struct {
	UID      uint   `json:"uid" db:"uid"`
	Uname    string `json:"uname" db:"uname" form:"uname" bind:"require"`
	Upasswd  string `json:"passwd" db:"upasswd" form:"passwd" bind:"require"`
	Status   uint8  `json:"status" db:"status"`
	Creatime string `json:"creatime" db:"creatime"`
	Updatime string `json:"updatime" db:"updatime"`
}

func CreatAdminCookie(u *ADMIN) (res Mycookie, err error) {

	res.Uid = strconv.Itoa(int(u.UID))
	res.Uname = u.Uname
	pek, err := gorsa.PriKeyEncrypt(res.Uid, string(privateKey))
	if err != nil {
		return
	}
	etime := time.Now().Format("2006 01 02 15 04")
	ptk, err := gorsa.PriKeyEncrypt(etime, string(privateKey))
	if err != nil {
		return
	}
	res.Ptkey = ptk
	res.Pnkey = pek[0:18]
	return

}
func adminLogin(c *gin.Context) {

	uname := c.Query("uname")
	// uname := c.PostForm("uname")
	if uname == "" {
		c.JSON(202, &ErrJson{
			Status: 0,
			Msg:    "请求错误",
		})
		return
	}
	upass := c.Query("upasswd")
	// upass := c.PostForm("passwd")
	if upass == "" {
		c.JSON(202, &ErrJson{
			Status: 0,
			Msg:    "请求错误",
		})
		return
	}
	res := ADMIN{}
	err := Db.Get(&res, "select * from adminer where uname=?", uname)
	if err != nil {
		c.JSON(202, &ErrJson{
			Status: 0,
			Msg:    "用户名错误",
		})
		return
	}
	if res.Status == 0 {
		c.JSON(202, &ErrJson{
			Status: 0,
			Msg:    "管理员已禁用",
		})
		return
	}
	if strings.Compare(upass, res.Upasswd) == 0 {
		cook, err := CreatAdminCookie(&res)
		if err != nil {
			c.JSON(500, &ErrJson{
				Status: 0,
				Msg:    "服务器繁忙，请稍后再试",
			})
			return
		}
		c.SetCookie("uname", cook.Uname, 0, "/", dominname, false, false) //本次会话有效
		c.SetCookie("token", cook.Pnkey, 0, "/", dominname, false, true)
		c.SetCookie("uid", cook.Uid, 0, "/", dominname, false, true)
		c.SetCookie("ptkey", cook.Ptkey, 0, "/", dominname, false, true)
		c.JSON(200, &ErrJson{
			Status: 1,
			Msg:    "登录成功",
		})
	} else {
		c.JSON(202, &ErrJson{
			Status: 0,
			Msg:    "密码错误",
		})
	}
	return
}
func isAdmin(uname string) (res bool) {
	res = false
	t := ADMIN{}
	err := Db.Get(&t, "select * from adminer where uname=?", uname)
	if err != nil {
		return
	} else {
		res = true
		return
	}
}

// func main() {
// 	initDb()
// 	res := ADMIN{}
// 	uname := "testq"
// 	err := Db.Get(&res, "select * from adminer where uname=?", uname)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	if res.Status == 1 {
// 		fmt.Println(res)
// 	}

// }
