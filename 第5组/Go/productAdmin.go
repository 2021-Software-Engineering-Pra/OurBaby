package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

func addproduct(c *gin.Context) {
	tdata := PRODUCT{}
	tdata.Status = 1
	err := c.ShouldBind(&tdata)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(tdata)
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
		err = os.Mkdir("D:/code/web/shop/web/src/static/"+pid, 066)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		_, err := Db.Exec("UPDATE product SET pname=? ,pdesp=?,peizhi=?,price=?,kucun=?,type=?,kind=?,`status`=? WHERE pid=?", tdata.Pname, tdata.Pdesp, tdata.Peizhi, tdata.Price, tdata.Kucun, tdata.Type, tdata.Kind, tdata.Status, tdata.PID)
		if err != nil {
			fmt.Println(err)
		}
		var t int32
		err = Db.Get(&t, "SELECT linkid FROM `product` WHERE pid=?", pid)
		if err != nil {
			fmt.Println(err)
		} else {
			pid = fmt.Sprint(t)
		}
	}
	// pid = "5"
	// err = os.Mkdir("D:/code/web/shop/web/src/static/"+pid, 066)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	files, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range files.File["intro"] {
		st := ""
		if k == 0 {
			st = "D:/code/web/shop/web/src/static/" + pid + "/intro" + ".jpg"
		} else {
			st = "D:/code/web/shop/web/src/static/" + pid + "/intro" + fmt.Sprint(k) + ".jpg"
		}
		c.SaveUploadedFile(v, st)
	}
	for _, v := range files.File["main"] {
		st := "D:/code/web/shop/web/src/static/" + pid + "/main" + ".jpg"
		c.SaveUploadedFile(v, st)
	}
	for _, v := range files.File["big"] {
		st := "D:/code/web/shop/web/src/static/" + pid + "/big" + ".jpg"
		c.SaveUploadedFile(v, st)
	}
	c.JSON(200, &ErrJson{
		Status: 1,
		Msg:    "ok",
	})
}
func delProduct(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	t := make(map[string]interface{})
	json.Unmarshal(data, &t)
	pid := fmt.Sprint(t["pid"])
	if pid != "" {
		_, err := Db.Exec("UPDATE product SET `status`=0 WHERE pid=?", pid)
		if err != nil {
			c.JSON(202, &ErrJson{
				Status: 0,
				Msg:    "服务器错误，请稍后重试",
			})
		} else {
			c.JSON(200, &ErrJson{
				Status: 1,
				Msg:    "ok",
			})
		}
	}
}
func getAllProduct(c *gin.Context) {

	data := make([]PRODUCT, 0)
	err := Db.Select(&data, "SELECT * FROM product LIMIT 1000")
	if err != nil {
		fmt.Println(err)
		c.JSON(202, &ErrJson{
			Status: 0,
			Msg:    "服务器错误，请稍后重试",
		})
	} else {
		res := ResJson{}
		res.fuchuzhi()
		for _, v := range data {
			res.Data = append(res.Data, v)
		}
		c.JSON(200, &res)
	}
}

// func main() {
// 	initDb()
// 	r := gin.Default()
// 	r.POST("/api/addproduct", addproduct)
// 	r.POST("/api/updateproduct", addproduct)
// 	r.POST("/api/delproduct", delProduct)
// 	r.GET("/api/getallproduct", getAllProduct)
// 	r.Run(":9001")
// }
