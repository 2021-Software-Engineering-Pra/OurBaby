package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type OREDER struct {
	ORID     string `json:"oid" db:"orid"`
	Uid      int    `json:"uid" db:"uid"`
	Count    string `json:"count" db:"count"`
	Pid      string `json:"pid" db:"pid"`
	Aprice   int32  `json:"aprice" db:"aprice"`
	Status   string `json:"status" db:"status"`
	Pname    string `json:"pname" db:"pname"`
	Peizhi   string `json:"peizhi" db:"peizhi"`
	Uname    string `json:"uname" db:"uname"`
	Linkid   int32  `json:"linkid" db:"linkid" `
	Recvname string `json:"recvname" db:"recvname" form:"recvname"`
	Phone    string `json:"phone" db:"phone" form:"phone"`
	Detailed string `json:"detailed" db:"detailed" form:"detailed"`
	Creatime string `json:"creatime" db:"creatime"`
	Updatime string `json:"updatime" db:"updatime"`
}

func QueryOrderColumn(column, orid string) (res string, err error) {
	err = Db.Get(&res, "select `"+column+"` from orders where orid=?", orid)
	return
}
func UpdateOrderColumn(column, newvalue, orid string) (err error) {
	_, err = Db.Exec("update orders set `"+column+"`=? where orid=?", newvalue, orid)
	return
}
func AddOrder(c *gin.Context) {
	ck, istrue := CheckCookie(c)
	if istrue {
		var temporder struct {
			Pid   []interface{} `json:"pid"`
			Count []int32       `json:"count"`
			Aid   string        `json:"aid"`
		}
		c.ShouldBind(&temporder)

		if len(temporder.Pid) == 0 || len(temporder.Pid) != len(temporder.Count) {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "请求错误，请重试",
			})
			return
		}
		pid := fmt.Sprint(temporder.Pid)
		pid = pid[1 : len(pid)-1]
		pids := strings.Split(pid, " ")
		temproduct, err := getProductByPid(pids)
		if err != nil {
			fmt.Println(err)
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "请求错误，请重试",
			})
		}
		count := fmt.Sprint(temporder.Count)
		count = count[1 : len(count)-1]
		for k, v := range temproduct {
			if v.Kucun < temporder.Count[k] {
				c.JSON(202, &ErrJson{
					Status: 0,
					Msg:    "库存不足，下单失败",
				})
			}
			cutnum := strconv.Itoa((int)(temporder.Count[k]))
			rwLock.Lock()
			UpdateProductColumn("kucun", "kucun-"+cutnum, v.PID)
			rwLock.Unlock()
		}
		addresinfo, err := QueryAdress(temporder.Aid)
		if err != nil {
			fmt.Println(err)
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "地址id有误",
			})
			return
		}
		ntime := time.Now().Format("20060102150405")
		orid := ntime + ck.Uid
		_, err = Db.Exec("insert into orders(orid,uid,pid,uname,pname,peizhi,count,aprice,linkid,recvname,phone,detailed) values(?,?,?,?,?,?,?,?,?,?,?,?)", orid, ck.Uid, pid, ck.Uname, temproduct[0].Pname, temproduct[0].Peizhi, count, temproduct[0].Price, temproduct[0].Linkid, addresinfo.Recvname, addresinfo.Phone, addresinfo.Detailed)
		if err != nil {
			fmt.Println(err)
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "系统繁忙，请稍后再试",
			})
		}

		c.JSON(200, &ErrJson{
			Status: 1,
			Msg:    orid,
		})
		return
	} else {
		c.JSON(202, &ErrJson{
			Status: 0,
			Msg:    "请先登录",
		})
		return
	}
}
func CancelOrder(c *gin.Context) {
	ck, istrue := CheckCookie(c)
	if istrue {
		order := &OREDER{}
		orid := c.Query("orid")
		if orid == "" {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "请求错误，请重试",
			})
			return
		}
		err := Db.Get(order, "select * from orders where orid=? and uid=?", orid, ck.Uid)
		if err != nil {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "请求错误，请重试",
			})
			return
		}
		uid, err := strconv.Atoi(ck.Uid)
		if err != nil || order.Uid != uid {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "请求错误，请重试",
			})
		}
		statu, _ := QueryOrderColumn("status", orid)
		if statu == "已完成" {
			return
		}
		rwLock.Lock()
		err = UpdateOrderColumn("status", "已取消", orid)
		rwLock.Unlock()
		if err != nil {
			fmt.Println(err)
			c.JSON(500, &ErrJson{
				Status: 0,
				Msg:    "系统繁忙，请稍后再试",
			})
			return
		}
		pid := strings.Split(order.Pid, " ")
		count := strings.Split(order.Count, " ")
		for k, v := range pid {
			UpdateProductColumn("kucun", "kucun+"+count[k], v)
		}
		if statu == "已支付" {
			var intcounts []int
			for _, v := range count {
				intcount, _ := strconv.Atoi(v)
				intcounts = append(intcounts, intcount)
			}
			getp, _ := getProductByPid(pid)
			yue, _ := QueryUserColumn("yue", ck.Uid)
			intyue, _ := strconv.Atoi(yue)
			total := 0
			for k, v := range getp {
				total += (int)(v.Price) * intcounts[k]
			}
			rest := intyue + total
			rests := strconv.Itoa(rest)
			rwLock.Lock()
			err = UpdateUserColumn("yue", rests, ck.Uid)
			rwLock.Unlock()
		}

		if err != nil {
			fmt.Println(err)
			c.JSON(500, &ErrJson{
				Status: 0,
				Msg:    "系统繁忙，请稍后再试",
			})
		} else {
			c.JSON(200, &ErrJson{
				Status: 1,
				Msg:    "您已成功取消该订单",
			})
			return
		}
		// UpdateProductColumn("kucun", "kucun+"+count, strconv.Itoa(order.Pid))

	} else {
		c.JSON(http.StatusNetworkAuthenticationRequired, &ErrJson{
			Status: 0,
			Msg:    "请先登录",
		})
	}
}
func QueryOrderByOrid(orid, uid string) (res OREDER, ishave bool) {

	err := Db.Get(&res, "SELECT * FROM `orders` WHERE orid=? and uid=?", orid, uid)
	if err != nil {
		fmt.Println(err)
		ishave = false
		return
	}

	ishave = true
	return
}
func QueryOrderByUid(uid string) (res []OREDER, err error) {
	err = Db.Select(&res, "SELECT * from orders WHERE uid= ?", uid)
	return
}
func PayOrder(c *gin.Context) {
	ck, istrue := CheckCookie(c)
	if istrue {
		order := &OREDER{}
		orid := c.Query("orid")
		paypass := c.Query("paypass")
		if len(orid) == 0 {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "请求错误，请重试",
			})
			return
		}
		err := Db.Get(order, "select * from orders where orid=? and uid=?", orid, ck.Uid)
		if err != nil {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "请求错误，请重试",
			})
			return
		}
		if order.Status == "已支付" {
			c.JSON(200, &ErrJson{
				Status: 1,
				Msg:    "您已成功支付订单",
			})
			return
		}
		uid, err := strconv.Atoi(ck.Uid)
		if err != nil || order.Uid != uid {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "请求错误，请重试",
			})
		}
		res, err := QueryUserColumn("paywd", ck.Uid)
		if strings.Compare(res, paypass) == 0 {
			pids := strings.Split(order.Pid, " ")
			counts := strings.Split(order.Count, " ")
			var intcounts []int
			for _, v := range counts {
				intcount, _ := strconv.Atoi(v)
				intcounts = append(intcounts, intcount)
			}
			getp, _ := getProductByPid(pids)
			yue, _ := QueryUserColumn("yue", ck.Uid)
			intyue, _ := strconv.Atoi(yue)
			total := 0
			for k, v := range getp {
				total += (int)(v.Price) * intcounts[k]
			}
			if intyue < total {
				c.JSON(202, &ErrJson{
					Status: 0,
					Msg:    "余额不足",
				})
				return
			}
			rest := intyue - total
			rests := strconv.Itoa(rest)
			rwLock.Lock()
			err = UpdateUserColumn("yue", rests, ck.Uid)
			err = UpdateOrderColumn("status", "已支付", orid)
			rwLock.Unlock()
			if err != nil {
				fmt.Println(err)
				c.JSON(202, &ErrJson{
					Status: 0,
					Msg:    "系统繁忙，请稍后再试",
				})
			} else {
				c.JSON(200, &ErrJson{
					Status: 1,
					Msg:    "您已成功支付订单",
				})
				return
			}
		} else {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "支付密码错误",
			})
		}

	} else {
		c.JSON(http.StatusNetworkAuthenticationRequired, &ErrJson{
			Status: 0,
			Msg:    "请先登录",
		})
	}
}
func GetOrder(c *gin.Context) {
	ck, is := CheckCookie(c)
	if is == false {
		c.JSON(200, &ErrJson{
			Status: 0,
			Msg:    "请先登录",
		})
		return
	}
	par := c.Query("orid")
	if par != "" {
		u, is := QueryOrderByOrid(par, ck.Uid)
		if is {
			r := &ResJson{}
			r.fuchuzhi()
			r.Data = append(r.Data, u)
			c.JSON(200, r)
		} else {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "无此订单，请确认后重试",
			})
		}
		return
	}
	u, err := QueryOrderByUid(ck.Uid)
	if err == nil {
		r := &ResJson{}
		r.fuchuzhi()
		r.Data = append(r.Data, u)
		c.JSON(200, r)
	} else {
		c.JSON(202, &ErrJson{
			Status: 0,
			Msg:    "服务器出错啦",
		})
	}
	return
}
func DelOrder(c *gin.Context) {
	ck, is := CheckCookie(c)
	if is {
		orid := c.Query("orid")
		if orid == "" {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "请求错误",
			})
			return
		}
		_, err := Db.Exec("DELETE FROM orders WHERE orid=? and uid=?", orid, ck.Uid)
		if err != nil {
			fmt.Println(err)
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "服务器错误，请稍后再试",
			})
			return
		}
		c.JSON(200, &ErrJson{
			Status: 1,
			Msg:    "删除成功",
		})
		return
	}
	c.JSON(202, &ErrJson{
		Status: 0,
		Msg:    "请先登录",
	})
	return
}
func ReceiveProduct(c *gin.Context) {
	_, is := CheckCookie(c)
	if is {
		orid := c.Query("orid")
		if orid == "" {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "请求参数不足",
			})
			return
		}
		err := UpdateOrderColumn("status", "已完成", orid)
		if err != nil {
			fmt.Println(err)
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "服务器错误，请稍后再试",
			})
			return
		}
		c.JSON(200, &ErrJson{
			Status: 1,
			Msg:    "收货成功",
		})
		return
	}
	c.JSON(202, &ErrJson{
		Status: 0,
		Msg:    "请先登录",
	})
	return
}

// func main() {
// 	initDb()
// 	res, err := QueryOrderByOrid("201912262014561", "1")
// 	fmt.Println(res, err)
// }
