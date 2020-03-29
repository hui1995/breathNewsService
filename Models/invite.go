package Models

import (
	"breathNewsService/Databases/Mysql"
	"github.com/jinzhu/gorm"
)

/**
 * @Author: hui
 * @Email: breathcoder@gmail.com
 * @Description:
 * @WebSite : https://www.breathcoder.cn
 * @Version: 1.0.0
 * @Date: 2020/3/21 10:16 PM
 */
type Invite struct {
	gorm.Model
	Inviter      int
	Invitee      int
	State        int
	OptionUserId int
	Desction     string
}

func init() {
	table := Mysql.DB.HasTable(Invite{})
	if !table {
		Mysql.DB.CreateTable(Invite{})
	}
}
func (this *Invite) InsertInvite(userId, inviteeId int) Invite {
	var invite Invite
	Mysql.DB.Where(Invite{Inviter: userId, Invitee: inviteeId}).FirstOrCreate(&invite)
	return invite

}

func (this *Invite) FindeByInvitee(userId int) bool {

	var count int
	Mysql.DB.Model(&Invite{}).Where("invitee = ? ", userId).Count(&count)
	if count == 0 {

		return true

	}

	return false

}

func (this *Invite) FindCountByInviter(userId int) int {

	var count int
	Mysql.DB.Model(&Invite{}).Where("inviter = ? ", userId).Count(&count)

	return count
}

func (this *Invite) FIndInviteeByInviter(userId int) []Invite {
	Db := Mysql.DB
	var invitelst []Invite
	Db = Db.Where("inviter = ?", userId).Find(&invitelst)
	return invitelst

}
