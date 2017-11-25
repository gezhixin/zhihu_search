package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"zhihu_search/service/core/db"
	"zhihu_search/spider"
)

func Router(r *gin.Engine) {
	grounp := r.Group("/user")
	{
		grounp.GET("/s", search)
	}
}

func search(c *gin.Context) {

	fmt.Println("user index")

	loction, _ := c.GetQuery("location")
	limitStr, _ := c.GetQuery("count")
	offsetStr, _ := c.GetQuery("offset")

	limit := int(200)
	offset := int(0)

	if len(limitStr) > 0 {
		limit_tmp, ok := strconv.Atoi(limitStr)
		if ok == nil {
			limit = limit_tmp
		}
	}

	if len(offsetStr) > 0 {
		offset_tmp, ok := strconv.Atoi(offsetStr)
		if ok == nil {
			offset = offset_tmp
		}
	}

	dbSesion, db, err := db.MongoDB()
	defer dbSesion.Close()

	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"code": 2001, "msg": "db connection error", "data": err})
		return
	}

	var userList = make([]spider.ZUser, 0, 10)
	err = db.C("User").Find(bson.M{"location": bson.M{"$regex": loction}, "gender": 0}).Sort("-_id").Limit(limit).Skip(offset).All(&userList)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"code": 2002, "msg": "db qurry error", "data": err})
		return
	}

	c.JSON(200, gin.H{"code": 0, "msg": "sucicess", "data": userList})
}
