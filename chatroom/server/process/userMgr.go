package process2

import (
	"fmt"
)

//userMgr实例只有一个，在服务器很多地方要使用到，
//定义为全局变量
var (
	usrMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成对userMgr初始化工作
func init()  {
	usrMgr = &UserMgr{
		onlineUsers : make(map[int]*UserProcess, 1024),
	}
}

//完成对onlineUsers添加
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//删除
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

//返回当前所有在线用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

//根据id返回对应值
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d不存在", userId)
		return
	}
	return
}