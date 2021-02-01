package main

import (
	"111/model"
	"fmt"
)


func main(){
	c,_ := model.GetCommentByID(1)
	fmt.Println(c)
}