package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Adress struct {
	Aid      string `json:"aid" db:"aid" form:"aid"`
	Uid      string `json:"-" db:"uid"`
	Recvname string `json:"recvname" db:"recvname" form:"recvname"`
	Phone    string `json:"phone" db:"phone" form:"phone"`
	Detailed string `json:"detailed" db:"detailed" form:"detailed"`
	Creatime string `json:"creatime" db:"creatime"`
	Updatime string `json:"updatime" db:"updatime"`
}

func AddAdress(c *gin.Context) {
	ck, istrue := CheckCookie(c)
	if istrue {
		a := &Adress{}
		err := c.Bind(a)
		if err != nil {
			fmt.Println(err)
			c.JSON(202, ErrJson{
				Status: 0,
				Msg:    "请求错误，请稍后再试",
			})
			return
		}
		err = addadre(ck.Uid, a.Recvname, a.Phone, a.Detailed)
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
			Msg:    "添加成功",
		})
		return

	}
	c.JSON(202, ErrJson{
		Status: 0,
		Msg:    "请先登录",
	})
	return
}
func QueryAdress(aid string) (res Adress, err error) {
	err = Db.Get(&res, "SELECT * FROM `address` WHERE aid=?", aid)
	return
}
func DelAdress(c *gin.Context) {
	ck, istrue := CheckCookie(c)
	if istrue {

		aid := c.Query("aid")
		if aid == "" {
			c.JSON(202, ErrJson{
				Status: 0,
				Msg:    "请求错误，请稍后再试",
			})
			return
		}
		err := deleteAdre(ck.Uid, aid)
		if err != nil {
			fmt.Println(err)
			c.JSON(202, ErrJson{
				Status: 0,
				Msg:    "服务器错误，请稍后再试",
			})
			return
		}
		fmt.Println(err)
		c.JSON(200, ErrJson{
			Status: 1,
			Msg:    "删除成功",
		})
		return

	}
	c.JSON(202, ErrJson{
		Status: 0,
		Msg:    "请先登录",
	})
	return
}
func deleteAdre(uid, aid string) (err error) {
	_, err = Db.Exec("delete from address where uid=? and aid=?", uid, aid)
	return
}
func addadre(uid, recvname, phone, detailed string) (err error) {
	_, err = Db.Exec("insert into address(uid,recvname,phone,detailed) values(?,?,?,?)", uid, recvname, phone, detailed)
	return
}
func UpdateAdressColumn(column, newvalue, aid, uid string) (err error) {
	_, err = Db.Exec("update adress set `"+column+"`="+newvalue+" where uid=? and aid=?", uid, aid)
	return
}
func queryAddress(uid string) (res []Adress, err error) {
	tem, err := Db.Query("select * from address where uid=?", uid)
	if err != nil {
		return
	}
	a := Adress{}
	for tem.Next() {
		err = tem.Scan(&a.Aid, &a.Uid, &a.Recvname, &a.Phone, &a.Detailed, &a.Creatime, &a.Updatime)
		if err != nil {
			fmt.Println(err)
			return
		}
		res = append(res, a)
	}
	return
}
func GetAddress(c *gin.Context) {
	ck, is := CheckCookie(c)
	if is {
		res, err := queryAddress(ck.Uid)
		if err != nil {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "请求错误",
			})
			return
		}
		c.JSON(200, gin.H{
			"status": 1,
			"msg":    "ok",
			"data":   res,
		})
		return
	}
}
func UpdateAddress(c *gin.Context) {
	ck, ishava := CheckCookie(c)
	if ishava {
		a := &Adress{}
		err := c.Bind(a)
		if err != nil {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "请求错误",
			})
			return
		}
		_, err = Db.Exec("UPDATE address SET recvname=?,phone=?,detailed=? where aid=? and uid=?", a.Recvname, a.Phone, a.Detailed, a.Aid, ck.Uid)
		if err != nil {
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
		Msg:    "请重新登录",
	})
	return
}
