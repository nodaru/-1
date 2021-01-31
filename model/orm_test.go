package model

import (
	"testing"
)

var (
	case1 = []struct {
		name     string
		email    string
		pass     string
		expected error
	}{
		{name: "aow", email: "admin@mails.tsinghua.edu.cn", pass: "124141241", expected: nil},
		{name: "rux", email: "pp@mails.tsinghua.edu.cn", pass: "124adaf1", expected: nil},
		{name: "rux", email: "pp@mails.tsinghua.edu.cn", pass: "124adaf1", expected: ERR_EMAIL_HAVE_BEEN_REGISTED},
	}
	case2 = []struct {
		uid      int
		path     string
		expected error
	}{
		{uid: 1, path: "halo", expected: nil},
		{uid: 2, path: "halooooa", expected: nil},
	}
	case3 = []struct {
		pid      int
		expected error
	}{
		{pid: 1, expected: nil},
		{pid: 2, expected: nil},
		{pid: 0, expected: ERR_POST_DOES_NOT_EXISTED}}
	case4 = []struct {
		pid      int
		uid      int
		refCid   int
		path     string
		expected error
	}{
		{pid: 1, uid: 1, refCid: 1, path: "aoao", expected: nil},
		{pid: 1, uid: 2, refCid: 2, path: "jiao", expected: nil},
	}
	case5 = []struct {
		uid      int
		expected error
	}{
		{uid: 1, expected: nil},
		{uid: 2, expected: nil},
		{uid: 0, expected: ERR_USER_DOES_NOT_EXISTED}}
)

func TestCreateNormalUser(t *testing.T) {
	for _, c := range case1 {
		if _, err := CreateNormalUser(c.name, c.email, c.pass); err != c.expected {
			t.Errorf(" %v", err)
		}
	}
}

func TestCreatPost(t *testing.T) {
	for _, c := range case2 {
		u, err := GetUserByID(c.uid)
		_, err = u.CreatePost(c.path)
		if err != c.expected {
			t.Errorf("%v", err)
		}
	}
}

func TestGetPostById(t *testing.T) {
	for _, c := range case3 {
		if _, err := GetPostByID(c.pid); err != c.expected {
			t.Errorf("%v", err)
		}
	}
}

func TestCreateComment(t *testing.T) {
	for _, c := range case4 {
		user, err := GetUserByID(c.uid)
		post, err := GetPostByID(c.pid)
		refcomment, err := GetCommentByID(c.refCid)
		comment, err := user.CreateComment(c.path, post, refcomment)
		if err != c.expected {
			t.Errorf("%v %v", err, comment)
		}
	}
}


func TestGetUserByID(t *testing.T) {
	for _, c := range case5 {
		if _, err := GetUserByID(c.uid); err != c.expected {
			t.Errorf("%v", err)
		}
	}
}