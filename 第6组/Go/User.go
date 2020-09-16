package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wenzhenxi/gorsa"
)

type USER struct {
	UID      uint   `json:"uid" db:"uid"`
	Uname    string `json:"uname" db:"uname" form:"uname" bind:"require"`
	Upasswd  string `json:"passwd" db:"upasswd" form:"passwd" bind:"require"`
	Paywd    string `json:"-" db:"paywd" form:"paywd"`
	Email    string `json:"email" db:"email" form:"email"`
	Phone    string `json:"phone" db:"phone" form:"phone"`
	Yue      int32  `json:"yue" db:"yue"`
	Status   uint8  `json:"status" db:"status"`
	Creatime string `json:"creatime" db:"creatime"`
	Updatime string `json:"updatime" db:"updatime"`
}

// Mycookie cookie 结构
type Mycookie struct {
	Uid   string
	Uname string
	Pnkey string
	Ptkey string
}

func ChangePasswd(c *gin.Context) {
	ck, istrue := CheckCookie(c)
	if istrue {
		pa := c.Query("opasswd")
		npa := c.Query("npasswd")
		var a string
		err := Db.Get(&a, "select upasswd from user where uid=?", ck.Uid)
		if err != nil {
			fmt.Println(err)
			c.JSON(202, ErrJson{
				Status: 0,
				Msg:    "服务器错误，请稍后再试",
			})
			return
		}
		if strings.Compare(a, pa) == 0 {
			err = UpdateUserColumn("upasswd", npa, ck.Uid)
			if err != nil {
				fmt.Println(err)
				c.JSON(202, ErrJson{
					Status: 0,
					Msg:    "服务器错误，请稍后再试",
				})
				return
			}
			c.JSON(200, ErrJson{
				Status: 1,
				Msg:    "更改成功",
			})
			return
		}
		c.JSON(202, ErrJson{
			Status: 0,
			Msg:    "旧密码错误，请稍后再试",
		})
		return
	}
	c.JSON(202, ErrJson{
		Status: 0,
		Msg:    "请重新登录",
	})
	return
}

func IsName(c *gin.Context) {

	name := c.Query("uname")
	if name == "" {
		c.JSON(202, &ErrJson{
			Status: 0,
			Msg:    "请求错误",
		})
		return
	}
	result := ""
	err := Db.Get(&result, "select uname from user where uname=?", name)
	if err != nil {
		c.JSON(200, &ErrJson{
			Status: 1,
			Msg:    "ok",
		})
		return
	}
	c.JSON(202, &ErrJson{
		Status: 0,
		Msg:    "该用户名已经使用",
	})
}

func QueryUserColumn(column, uid string) (res string, err error) {
	err = Db.Get(&res, "select `"+column+"` from user where uid=?", uid)
	return
}
func UpdateUserColumn(column, newvalue, uid string) (err error) {
	_, err = Db.Exec("update user set `"+column+"`='"+newvalue+"' where uid=?", uid)
	return
}

func ChangeEmail(c *gin.Context) {
	ck, is := CheckCookie(c)
	if is {
		email := c.Query("email")
		if email == "" {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "请求错误",
			})
			return
		}
		err := UpdateUserColumn("email", email, ck.Uid)
		if err != nil {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "服务器错误",
			})
			return
		}
		c.JSON(200, &ErrJson{
			Status: 1,
			Msg:    "更改成功",
		})
	}

}
func ChangePhone(c *gin.Context) {
	ck, is := CheckCookie(c)
	if is {
		phone := c.Query("phone")
		if phone == "" {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "请求错误",
			})
			return
		}
		err := UpdateUserColumn("phone", phone, ck.Uid)
		if err != nil {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "服务器错误",
			})
			return
		}
		c.JSON(200, &ErrJson{
			Status: 1,
			Msg:    "更改成功",
		})
	}

}
func QueryUser(uname string) (*USER, error) {
	result := USER{}
	err := Db.Get(&result, "select * from user where uname=?", uname)
	if err != nil {
		log.Println(err)
		result.Upasswd = ""
		return &result, err
	}
	result.Upasswd = ""
	return &result, err
}
func CreatCookie(u *USER) (res Mycookie, err error) {

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
func UpdatePaypass(c *gin.Context) {
	ck, is := CheckCookie(c)
	if is {
		opaypass := c.Query("opaypass")
		paypass := c.Query("npaypass")
		if paypass == "" {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "参数不足，请确认后重试",
			})
			return
		}
		res, err := QueryUserColumn("paywd", ck.Uid)
		if err != nil {
			fmt.Println(err)
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "服务器错误，请稍后重试",
			})
		}
		if strings.Compare(res, opaypass) == 0 {
			err = UpdateUserColumn("paywd", paypass, ck.Uid)
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
				Msg:    "更改成功",
			})
			return
		}
		c.JSON(202, &ErrJson{
			Status: 0,
			Msg:    "旧密码错误，请确认后重试",
		})
		return

	}

}
func NothavePaypass(c *gin.Context) {
	ck, is := CheckCookie(c)
	if is {
		res, err := QueryUserColumn("paywd", ck.Uid)
		if err != nil {
			fmt.Println(err)
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "服务器错误，请稍后重试",
			})
			return
		}
		if res == "" {
			c.JSON(200, &ErrJson{
				Status: 1,
				Msg:    "你还没有设置支付密码",
			})
			return
		}
		c.JSON(202, &ErrJson{
			Status: 0,
			Msg:    "ok",
		})
		return
	}
	c.JSON(202, &ErrJson{
		Status: 0,
		Msg:    "请先登录",
	})
}
func CheckCookie(c *gin.Context) (ck Mycookie, res bool) {
	res = false
	un, err := c.Cookie("uname")
	if err != nil {
		return
	}
	ui, err := c.Cookie("uid")
	if err != nil {
		return
	}
	unkey, err := c.Cookie("token")
	if err != nil {
		return
	}
	ct, err := gorsa.PriKeyEncrypt(ui, string(privateKey))
	if err != nil {
		return
	}
	if strings.HasPrefix(ct, unkey) {
		unkey, err := c.Cookie("ptkey")
		if err != nil {
			return
		}
		dptkey, err := gorsa.PublicDecrypt(unkey, string(publicKey))
		if err != nil {

			return
		}
		et, err := time.Parse("2006 01 02 15 04", dptkey)
		if err != nil {

			return
		}
		if et.After(time.Now()) {
			ck.Uname = un
			ck.Uid = ui
			res = true
			return
		}
		return
	}
	return
}
func GetUserinfo(c *gin.Context) {
	ck, is := CheckCookie(c)
	if is {
		res, err := QueryUser(ck.Uname)
		if err != nil {
			fmt.Println(err)
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "服务器错误，请稍后重试",
			})
			return
		}
		re := &ResJson{}
		re.fuchuzhi()
		re.Data = append(re.Data, res)
		c.JSON(200, re)
		return
	}
}

// func main() {
// 	u := USER{
// 		Uname:   "test2",
// 		Upasswd: "test2",
// 		Paywd:   "123456",
// 		Email:   "ouc@stu.ouc.edu.cn",
// 		Phone:   "13123456789",
// 	}
// }
