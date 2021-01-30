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
	}
)

func TestCreateNormalUser(t *testing.T) {
	for _, c := range cases {
		if _,err:=CreateNormalUser(c.name,c.email,c.pass);err !=nil {
			t.Errorf("%s is expected to be %v", c.email, c.valid)
		}
	}
}