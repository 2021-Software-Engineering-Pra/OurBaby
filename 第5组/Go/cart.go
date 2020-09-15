package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CART struct {
	CID      int    `json:"cid" db:"cid"`
	UID      int    `json:"uid" db:"uid"`
	Pid      int    `json:"pid" db:"pid"`
	Pname    string `json:"pname" db:"pname"`
	Peizhi   string `json:"peizhi" db:"peizhi"`
	Linkid   int32  `json:"linkid" db:"linkid"`
	Type     string `json:"type" db:"type"`
	Count    int32  `json:"count" db:"count"`
	Aprice   int32  `json:"aprice" db:"aprice"`
	Creatime string `json:"creatime" db:"creatime"`
}

func AddCart(c *gin.Context) {
	ck, ishave := CheckCookie(c)
	if ishave {
		temproduct := &PRODUCT{}
		pid := c.Query("pid")
		count := c.Query("count")
		if pid == "" || count == "" {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "请求错误，请稍后重试",
			})
			return
		}
		cart := CART{}

		err := Db.Get(&cart, "SELECT * FROM `cart` WHERE uid=?&&pid=?", ck.Uid, pid)

		if err == nil {
			icid := strconv.Itoa(cart.CID)
			err = UpdateCartColumn("count", "count+"+count, icid)
			if err != nil {
				fmt.Println(err)
				c.JSON(500, &ErrJson{
					Status: 0,
					Msg:    "服务器错误，请稍后重试",
				})
				return
			}
			c.JSON(200, &ErrJson{
				Status: 1,
				Msg:    "添加成功",
			})
			return
		}
		err = Db.Get(temproduct, "select * from product where pid=?", pid)
		if err != nil {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "服务器错误，请稍后重试",
			})
			return
		}
		_, err = Db.Exec("insert into cart(uid,pid,pname,peizhi,type,linkid,count,aprice) values(?,?,?,?,?,?,?,?)", ck.Uid, pid, temproduct.Pname, temproduct.Peizhi, temproduct.Type, temproduct.Linkid, count, temproduct.Price)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, &ErrJson{
				Status: 0,
				Msg:    "服务器错误，请稍后重试",
			})
			return
		}
		c.JSON(200, &ErrJson{
			Status: 1,
			Msg:    "添加成功",
		})
		return
	}
	c.JSON(202, &ErrJson{
		Status: 0,
		Msg:    "请先登录",
	})

}
func RemovCart(c *gin.Context) {
	ck, ishave := CheckCookie(c)
	if ishave {
		pid := c.Query("pid")
		if pid == "" {
			c.JSON(400, ErrJson{
				Msg: "请求错误",
			})
			return
		}
		_, err := Db.Exec("DELETE FROM cart WHERE uid=?&&pid=?", ck.Uid, pid)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, ErrJson{
				Msg: "服务器错误，请稍后再试",
			})
			return
		}
		c.JSON(200, ErrJson{
			Status: 1,
			Msg:    "成功移除",
		})
		return
	}
	c.JSON(400, ErrJson{
		Msg: "请先登录",
	})
}
func QueryCartColumn(column, cid string) (res string, err error) {
	err = Db.Get(&res, "select `"+column+"` from cart where cid=?", cid)
	return
}
func UpdateCartColumn(column, newvalue, cid string) (err error) {
	_, err = Db.Exec("update cart set `"+column+"`="+newvalue+" where cid=?", cid)
	return
}
func GetCart(c *gin.Context) {
	ck, ishave := CheckCookie(c)
	if ishave {
		res, err := Db.Query("select * from cart where uid=? order by cid desc", ck.Uid)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, &ErrJson{
				Status: 0,
				Msg:    "服务器错误，请稍后重试",
			})
			return
		}
		temp := CART{}
		resp := &ResJson{}
		resp.fuchuzhi()
		for res.Next() {
			err = res.Scan(&temp.CID, &temp.UID, &temp.Pid, &temp.Pname, &temp.Peizhi, &temp.Type, &temp.Linkid, &temp.Count, &temp.Aprice, &temp.Creatime)
			if err != nil {
				fmt.Println(err)
				c.JSON(500, &ErrJson{
					Status: 0,
					Msg:    "服务器错误，请稍后重试",
				})
				return
			}
			resp.Data = append(resp.Data, temp)
		}
		c.JSON(200, resp)
		return
	}
	c.JSON(500, &ErrJson{
		Status: 0,
		Msg:    "请先登录",
	})
	return
}
func DeCartNum(c *gin.Context) {
	ck, ishave := CheckCookie(c)
	if ishave {
		pid := c.Query("pid")
		count := c.Query("count")
		if pid == "" {
			c.JSON(400, &ErrJson{
				Msg: "请求错误，请重试",
			})
			return
		}
		cart := CART{}
		err := Db.Get(&cart, "SELECT * FROM `cart` WHERE uid=?&&pid=?", ck.Uid, pid)
		if err == nil {
			cid := strconv.Itoa(cart.CID)
			err = UpdateCartColumn("count", "count-"+count, cid)
			if err != nil {
				c.JSON(500, &ErrJson{
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
		c.JSON(400, &ErrJson{
			Status: 0,
			Msg:    "错误的请求",
		})
	}
	c.JSON(500, &ErrJson{
		Status: 0,
		Msg:    "请先登录",
	})
	return
}
func SetCartNum(c *gin.Context) {
	ck, ishave := CheckCookie(c)
	if ishave {
		pid := c.Query("pid")
		count := c.Query("count")
		if pid == "" {
			c.JSON(400, &ErrJson{
				Msg: "请求错误，请重试",
			})
			return
		}
		cart := CART{}
		err := Db.Get(&cart, "SELECT * FROM `cart` WHERE uid=?&&pid=?", ck.Uid, pid)
		if err == nil {
			cid := strconv.Itoa(cart.CID)
			err = UpdateCartColumn("count", count, cid)
			if err != nil {
				c.JSON(500, &ErrJson{
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
		c.JSON(400, &ErrJson{
			Status: 0,
			Msg:    "错误的请求",
		})
	}
	c.JSON(500, &ErrJson{
		Status: 0,
		Msg:    "请先登录",
	})
	return
}
