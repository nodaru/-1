package model

import (
	"testing"
)

var (
	cases = []struct {
		name  string
		email string
		pass string
		valid bool
	}{
		{name:"aow",email: "admin@mails.tsinghua.edu.cn",pass:"124141241", valid: true},
		{name:"rux",email: "pp@mails.tsinghua.edu.cn",pass:"124adaf1", valid: true},
		{name:"oho",email: "admin@mails.tsinghua.edu.cn",pass:"124141241", valid: true},
	}
)


var (
	case1 = []struct {
		uid int
		path string
	}{
		{uid: 1, path: "halo"},
		{uid: 2, path: "ha"},
	}
)

var (uid =[]int{0,1,2,3,4,5,6})

func TestCreateNormalUser(t *testing.T) {
	for _, c := range cases {
		if _,err:=CreateNormalUser(c.name,c.email,c.pass);err !=nil {
			t.Errorf(" %v", err)
		}
	}
}

func TestCreatPost(t *testing.T) {
	for _, c := range case1 {
		if _,err:=CreatPost(c.uid, c.path);err !=nil {
			t.Errorf("%v", err)
		}
	}
}


func TestGetUserByID(t *testing.T) {
	for _, c := range uid {
		if _,err:=GetUserByID(c);err !=nil {
			t.Errorf(" %v",err)
		}
	}
}