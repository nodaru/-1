package model

import (
	"testing"
)

var (
	case1 = []struct {
		name  string
		email string
		pass string
		valid bool
	}{
		{name:"aow",email: "admin@mails.tsinghua.edu.cn",pass:"124141241", valid: true},
		{name:"rux",email: "pp@mails.tsinghua.edu.cn",pass:"124adaf1", valid: true},
		{name:"oho",email: "admin@mails.tsinghua.edu.cn",pass:"124141241", valid: true},
	}
	case2 = []struct {
		uid int
		path string
	}{
		{uid: 1, path: "halo"},
		{uid: 2, path: "ha"},
	}
	uid =[]int{1}
	case3 = []int{1,2}
)


func TestCreateNormalUser(t *testing.T) {
	for _, c := range case1 {
		if _,err:=CreateNormalUser(c.name,c.email,c.pass);err !=nil {
			t.Errorf(" %v", err)
		}
	}
}

func TestCreatPost(t *testing.T) {
	for _, c := range case2 {
		u,err := GetUserByID(c.uid)
		if err !=nil {
			t.Errorf(" %v",err)
		}
		_,err = u.CreatePost(c.path)
		if err!=nil {
			t.Errorf("%v", err)
		}
	}
}

func TestGetUserByID(t *testing.T) {
	for _, c := range uid {
		if _,err:=GetUserByID(c);err !=nil {
			t.Errorf("%v",err)
		}
	}
}

func TestGetPostById(t *testing.T){
	for _, c := range case3 {
		if _,err:=GetPostByID(c);err !=nil {
			t.Errorf("%v",err)
		}
	}
}