package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type PRODUCT struct {
	PID       string `json:"pid" db:"pid"`
	Pname     string `json:"pname" db:"pname" form:"name"`
	Pdesp     string `json:"pdesp" db:"pdesp" form:"pdesp"`
	Peizhi    string `json:"peizhi" db:"peizhi" form:"peizhi"`
	Kucun     int32  `json:"kucun" db:"kucun" form:"kucun"`
	Price     int32  `json:"price" db:"price" form:"price"`
	Type      string `json:"type" db:"type" form:"type"`
	Kind      string `json:"kind" db:"kind" form:"kind"`
	Sellcount int32  `json:"sellcount" db:"sellcount"`
	Status    uint8  `json:"status" db:"status"`
	Linkid    int32  `json:"linkid" db:"linkid"`
	Creatime  string `json:"creatime" db:"creatime"`
	Updatime  string `json:"updatime" db:"updatime"`
}

//QueryProduct 查询商品
func QueryProductByType(c *gin.Context) {
	key := c.Query("type")
	initDb()
	res, err := Db.Query("SELECT * FROM product WHERE type=?", key)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, &ErrJson{
			Status: 0,
			Msg:    "服务器错误，请重试",
		})
		return
	}
	muban := PRODUCT{}
	resp := ResJson{
		Status: 1,
		Msg:    "ok",
	}
	for res.Next() {
		err = res.Scan(&muban.PID, &muban.Pname, &muban.Pdesp, &muban.Peizhi, &muban.Kucun, &muban.Price, &muban.Type, &muban.Kind, &muban.Sellcount, &muban.Status, &muban.Linkid, &muban.Creatime, &muban.Updatime)
		if err != nil {
			fmt.Println(err)
		}
		resp.Data = append(resp.Data, muban)
	}
	c.JSON(200, resp)
}
func QueryProductByKind(c *gin.Context) {
	key := c.Query("kind")
	initDb()
	res, err := Db.Query("SELECT * FROM product WHERE kind=?", key)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, &ErrJson{
			Status: 0,
			Msg:    "服务器错误，请重试",
		})
		return
	}
	muban := PRODUCT{}
	resp := ResJson{
		Status: 1,
		Msg:    "ok",
	}
	for res.Next() {
		err = res.Scan(&muban.PID, &muban.Pname, &muban.Pdesp, &muban.Peizhi, &muban.Kucun, &muban.Price, &muban.Type, &muban.Kind, &muban.Sellcount, &muban.Status, &muban.Linkid, &muban.Creatime, &muban.Updatime)
		if err != nil {
			fmt.Println(err)
		}
		resp.Data = append(resp.Data, muban)
	}
	c.JSON(200, resp)
}
func QueryProductOrderBy(Column string) (rest []PRODUCT, err error) {

	res, err := Db.Query("SELECT * FROM product order by ? DESC limit 4", Column)
	if err != nil {
		return
	}
	muban := PRODUCT{}
	for res.Next() {
		err = res.Scan(&muban.PID, &muban.Pname, &muban.Pdesp, &muban.Peizhi, &muban.Kucun, &muban.Price, &muban.Type, &muban.Kind, &muban.Sellcount, &muban.Status, &muban.Linkid, &muban.Creatime, &muban.Updatime)
		if err != nil {
			fmt.Println(err)
		}
		rest = append(rest, muban)
	}
	return
}
func FindBestSellProduct(c *gin.Context) {
	res, err := QueryProductOrderBy("sellcount")
	if err != nil {
		c.JSON(500, &ErrJson{
			Status: 0,
			Msg:    "服务器错误，请稍后重试",
		})
		return
	}
	ret := &ResJson{}
	ret.fuchuzhi()
	ret.Data = append(ret.Data, res)
	c.JSON(200, ret)
	return
}
func FindNewProduct(c *gin.Context) {
	res, err := QueryProductOrderBy("pid")
	if err != nil {
		c.JSON(500, &ErrJson{
			Status: 0,
			Msg:    "服务器错误，请稍后重试",
		})
		return
	}
	ret := &ResJson{}
	ret.fuchuzhi()
	ret.Data = append(ret.Data, res)
	c.JSON(200, ret)
	return
}
func SearchProduct(c *gin.Context) {
	key := c.Query("key")
	res, err := Db.Query("SELECT * FROM product WHERE pname LIKE '%" + key + "%' order by pid DESC limit 100")
	if err != nil {
		fmt.Println(err)
		c.JSON(500, &ErrJson{
			Status: 0,
			Msg:    "服务器错误，请重试",
		})
		return
	}
	muban := PRODUCT{}
	resp := ResJson{
		Status: 1,
		Msg:    "ok",
	}
	for res.Next() {
		err = res.Scan(&muban.PID, &muban.Pname, &muban.Pdesp, &muban.Peizhi, &muban.Kucun, &muban.Price, &muban.Type, &muban.Kind, &muban.Sellcount, &muban.Status, &muban.Linkid, &muban.Creatime, &muban.Updatime)
		if err != nil {
			fmt.Println(err)
		}
		resp.Data = append(resp.Data, muban)
	}
	c.JSON(200, resp)
}

func QueryProductColumn(column, pid string) (res string, err error) {
	err = Db.Get(&res, "select `"+column+"` from product where pid=?", pid)
	return
}
func UpdateProductColumn(column, newvalue, pid string) (err error) {
	_, err = Db.Exec("update product set `"+column+"`="+newvalue+" where pid=?", pid)
	return
}
func getProductByPid(pid []string) (res []PRODUCT, err error) {
	stmt, err := Db.Prepare("select * from product where pid=?")
	if err != nil {
		return
	}
	for _, v := range pid {
		gets, er := stmt.Query(v)
		if er != nil {
			return
		}
		temp := PRODUCT{}
		for gets.Next() {
			er = gets.Scan(&temp.PID, &temp.Pname, &temp.Pdesp, &temp.Peizhi, &temp.Kucun, &temp.Price, &temp.Type, &temp.Kind, &temp.Sellcount, &temp.Status, &temp.Linkid, &temp.Creatime, &temp.Updatime)
			if er != nil {
				return
			}
			res = append(res, temp)
		}
	}
	return
}
func Products(c *gin.Context) {
	// pids := c.QueryArray("pid")
	var temps struct {
		Pid []interface{} `json:"pid"`
	}
	c.ShouldBind(&temps)
	var last []string
	re := fmt.Sprint(temps.Pid)
	re = re[1 : len(re)-1]
	last = strings.Split(re, " ")
	if len(last) > 0 {
		res, err := getProductByPid(last)
		if err != nil {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "服务器错误",
			})
			return
		}

		rj := &ResJson{}
		rj.fuchuzhi()
		rj.Data = append(rj.Data, res)
		c.JSON(200, rj)
		return
	}
	c.JSON(202, &ErrJson{
		Status: 0,
		Msg:    "请求错误",
	})
}

// func main() {
// 	initDb()
// 	res, err := getProductByPid([]string{"1", "2", "3"})
// 	fmt.Println(res, err)
// }
