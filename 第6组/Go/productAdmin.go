package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

func queryallProduct() []PRODUCT {
	res := make([]PRODUCT, 0)
	err := Db.Select(&res, "select * from order")
	if err != nil {
		fmt.Println(err)
	}
	return res
}
func getallproduct(c *gin.Context) {
	products := queryallProduct()
	res := ResJson{}
	for _, v := range products {
		res.Data = append(res.Data, v)
	}
	res.fuchuzhi()
	c.JSON(200, res)
}

func addproduct(c *gin.Context) {
	tdata := PRODUCT{}
	tdata.Status = 1
	err := c.ShouldBind(&tdata)
	if err != nil {
		fmt.Println(err)
	}
	pid := tdata.PID
	if pid == "" {
		sqlr, err := Db.Query("SELECT linkid FROM `product` WHERE type=?", tdata.Type)
		if err != nil {
			fmt.Println(err)
		} else {
			var tlinkid int32
			for sqlr.Next() {
				sqlr.Scan(&tlinkid)
			}
			if tlinkid > 0 {
				tdata.Linkid = tlinkid
			} else {
				err = Db.Get(&tlinkid, "SELECT linkid FROM product ORDER BY linkid DESC LIMIT 1")
				if err != nil {
					fmt.Println(err)
				} else {
					tdata.Linkid = tlinkid + 1
				}
			}
			sqlres, err := Db.Exec("INSERT INTO product(pname,pdesp,peizhi,price,kucun,type,kind,linkid,`status`) VALUES(?,?,?,?,?,?,?,?,?)", tdata.Pname, tdata.Pdesp, tdata.Peizhi, tdata.Price, tdata.Kucun, tdata.Type, tdata.Kind, tdata.Linkid, tdata.Status)
			if err != nil {
				fmt.Println(err)
			} else {
				t, _ := sqlres.RowsAffected()
				if t == 1 {
					pid = fmt.Sprint(tdata.Linkid)
				}
			}
		}
		err = os.Mkdir("E:/code/src/static/"+pid, 066)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		var t int32
		err := Db.Get(&t, "SELECT linkid FROM `product` WHERE pid=?", pid)
		if err != nil {
			fmt.Println(err)
		} else {
			pid = fmt.Sprint(t)
		}
	}
	files, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range files.File["intro"] {
		st := "E:/code/src/static/" + pid + "/intro" + fmt.Sprint(k) + ".jpg"
		c.SaveUploadedFile(v, st)
	}
	for _, v := range files.File["main"] {
		st := "E:/code/src/static/" + pid + "/main" + ".jpg"
		c.SaveUploadedFile(v, st)
	}
	for _, v := range files.File["big"] {
		st := "E:/code/src/static/" + pid + "/big" + ".jpg"
		c.SaveUploadedFile(v, st)
	}
	c.JSON(200, &ErrJson{
		Status: 1,
		Msg:    "ok",
	})
}

func delproduct(c *gin.Context) {
	res := make(map[string]interface{})
	pid, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(pid, &res)
	err := UpdateProductColumn("status", "0", fmt.Sprint(res["pid"]))
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

// func main() {
// 	initDb()
// 	r := gin.Default()
// 	r.POST("/api/addproduct", addproduct)
// 	r.Run(":9001")
// }
