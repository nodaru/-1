package model

import "errors"



//UserPrivilege 代表着用户的权限
type UserPrivilege int16
/*
baned 指代被禁止访问的用户 其无法注销，无法浏览，也无法发言
muted 指代被禁言的用户 其权限与匿名用户相同
admin 指代最高级的管理员 拥有最高权限，可以恢复被删除的内容
manager 指代一般的管理员 可以处理举报，删除发言，有限次数的封禁用户/每三天，解封被自身封禁的用户
normalUser 指代正常用户 可以正常的使用功能
anonymous 指代未登录的用户 与banned用户不同是其未登录
*/
const (
	banedUserPrivilege     = -100
	mutedUserPrivilege     = -1
	adminPrivilege         = 0
	managerPrivilege       = 1
	normalUserPrivilege    = 10
	anonymousUserPrivilege = 100
)

type state int
const (
	state_block   state = -10
	state_nice    state = 1
	state_normal  state = 0
	state_deleted state = -2
	state_edited  state = -1
	state_pined   state = 2
)


//SysNoticeType 是系统通知的类型
var SysNoticeType map[int]string = map[int]string{
	1: "新增回复",
	2: "系统通知",
}

var (
	ERR_EMAIL_HAVE_BEEN_REGISTED = errors.New("email have been regisisted")
	ERR_POST_DOES_NOT_EXISTED = errors.New("post does not existed")
	ERR_USER_DOES_NOT_EXISTED = errors.New("user does not existed")
	ERR_COMMENT_DOES_NOT_EXISTED = errors.New("comment does not existed")
)

//ReportReasonType 代表举报相关的原因
type ReportReasonType map[int]string

//ReportReasonKeyType 代表举报相关的原因的key
type ReportReasonKeyType int

//ReportReason 是举报的类别
var ReportReason ReportReasonType = map[int]string{
	0: "政治有害",
	1: "引战",
	2: "令人不适",
}

//HandleReason 代表着 存档原因/分类
var HandleReason map[int]string = map[int]string{
	0: "Admin Delete",
	1: "Admin Ban",
	2: "Admin unDelete",
	3: "Admin unBan",
	4: "User Delete",
	5: "User Edit",
}