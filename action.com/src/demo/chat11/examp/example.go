package examp

//user 用户
type user struct {
	Name  string
	Email string
}

//Admin 管理员
type Admin struct {
	user
	Rights int
}
