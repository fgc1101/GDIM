package main

import (
	"GDim/lib"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// gd im start
	fmt.Println("运行gin框架")
	r := gin.Default()

	r.GET("/index", func(context *gin.Context) {
		context.JSON(200, gin.H{"code": 1, "msg": "success"})
	})

	r.GET("/user_list", GetUserList)

	r.Run()

}

func GetUserList(ctx *gin.Context) {

	// id, _ := strconv.Atoi(ctx.DefaultQuery("id", "1"))
	// fmt.Printf("值是%d\n", id)

	// username := ctx.PostForm("name")
	// fmt.Printf("值是%v\n", username)
	var userRequest lib.UserRequest

	if err := ctx.BindJSON(&userRequest); err != nil {
		return
	}
	fmt.Printf("值为：%v\n", userRequest.Name)
	fmt.Printf("值为：%v\n", userRequest.Age)

	var userList []string
	userList = append(userList, "姜媛", "运河")
	fmt.Println(userList[1])
	data := make(map[string][]string)
	data["user_list"] = userList

	ctx.JSON(200, gin.H{"code": 1, "msg": "success", "data": data})
}
