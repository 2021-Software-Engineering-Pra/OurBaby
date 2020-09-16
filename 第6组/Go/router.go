package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

//ResJson 返回的json数据结构
type ResJson struct {
	Status uint8         `json:"status"`
	Msg    string        `json:"msg"`
	Data   []interface{} `json:"data"`
}
type ErrJson struct {
	Status uint8  `json:"status"`
	Msg    string `json:"msg"`
}

func (c *ResJson) fuchuzhi() {
	c.Status = 1
	c.Msg = "ok"
	return
}
func login(c *gin.Context) {

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
	res := USER{}
	err := Db.Get(&res, "select * from user where uname=?", uname)
	if err != nil {
		c.JSON(202, &ErrJson{
			Status: 0,
			Msg:    "用户名错误",
		})
		return
	}

	if strings.Compare(upass, res.Upasswd) == 0 {
		cook, err := CreatCookie(&res)
		if err != nil {
			c.JSON(500, &ErrJson{
				Status: 0,
				Msg:    "服务器繁忙，请稍后再试",
			})
			return
		}
		c.SetCookie("uname", cook.Uname, 10800, "/", dominname, false, false)
		c.SetCookie("token", cook.Pnkey, 10800, "/", dominname, false, true)
		c.SetCookie("uid", cook.Uid, 10800, "/", dominname, false, true)
		c.SetCookie("ptkey", cook.Ptkey, 10800, "/", dominname, false, true)
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

func ShowCookie(c *gin.Context) {
	_, is := CheckCookie(c)
	if is {
		c.JSON(200, gin.H{
			"msg": "good",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "bad",
	})
	return
}

func Zhuce(c *gin.Context) {
	u := USER{}
	err := c.Bind(&u)
	if err != nil || u.Uname == "" {
		fmt.Println(err)
		c.JSON(202, &ErrJson{
			Status: 0,
			Msg:    "错误请求，请确认后重试",
		})
		return
	}
	_, err = Db.Exec("INSERT INTO `user`(uname,upasswd,email,phone) VALUES(?,?,?,?)", u.Uname, u.Upasswd, u.Email, u.Phone)
	if err != nil {
		fmt.Println(err)
		c.JSON(202, &ErrJson{
			Status: 0,
			Msg:    "服务器错误，请稍后重试",
		})
		return
	}
	c.JSON(200, &ErrJson{
		Status: 1,
		Msg:    "注册成功",
	})

}
func Logout(c *gin.Context) {
	cook, is := CheckCookie(c)
	if is {
		c.SetCookie("uname", cook.Uname, -1, "/", dominname, false, false)
		c.SetCookie("token", cook.Pnkey, -1, "/", dominname, false, true)
		c.SetCookie("uid", cook.Uid, -1, "/", dominname, false, true)
		c.SetCookie("ptkey", cook.Ptkey, -1, "/", dominname, false, true)
		c.JSON(200, &ErrJson{
			Status: 1,
			Msg:    "注销成功",
		})
		return
	}
}

// func main() {
// 	gin.SetMode(gin.ReleaseMode)
// 	initDb()
// 	dominname = "localhost"
// 	r := gin.Default()
// 	r.GET("/api/login", login)
// 	r.GET("/", ShowCookie)
// 	r.POST("/api/cancelorder", CancelOrder)
// 	r.POST("/api/addorder", AddOrder)
// 	r.GET("/api/getorder", GetOrder)
// 	r.GET("/api/cancelorder", CancelOrder)
// 	r.GET("/api/receive", ReceiveProduct)
// 	r.GET("/api/delorder", DelOrder)
// 	r.GET("/api/search", SearchProduct)
// 	r.GET("/api/getnew", FindNewProduct)
// 	r.GET("/api/gethot", FindBestSellProduct)
// 	r.GET("/api/getproduct", QueryProductByType)
// 	r.POST("/api/products", Products)
// 	r.GET("/api/getkind", QueryProductByKind)
// 	r.POST("/api/zhuce", Zhuce)
// 	r.GET("/api/addcart", AddCart)
// 	r.GET("/api/delcart", RemovCart)
// 	r.POST("/api/decart", DeCartNum)
// 	r.GET("/api/setcart", SetCartNum)
// 	r.GET("/api/getcart", GetCart)
// 	r.POST("/api/addaddress", AddAdress)
// 	r.GET("/api/deladdress", DelAdress)
// 	r.GET("/api/getaddress", GetAddress)
// 	r.GET("/api/userinfo", GetUserinfo)
// 	r.GET("/api/isuname", IsName)
// 	r.POST("/api/nothavepaypass", NothavePaypass)
// 	r.GET("/api/updateemail", ChangeEmail)
// 	r.GET("/api/updatephone", ChangePhone)
// 	r.GET("/api/updatepasswd", ChangePasswd)
// 	r.GET("/api/updatepaypasswd", UpdatePaypass)
// 	r.GET("/api/nothavepaypass", NothavePaypass)
// 	r.GET("/api/topay", PayOrder)
// 	r.POST("/api/updateaddress", UpdateAddress)
// 	r.POST("/api/logout", Logout)

// 	r.Run(":9001")
// }
