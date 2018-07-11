// Package dbmodel provides ...
package dbmodel

import (
	"fmt"

	"github.com/duminghui/go-pos-service/db"
)

type dstUser struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	CreateAt    int    `db:"create_at"`
	CreateAtStr string `db:"create_at_str"`
}

func FindUserByID(id string) *dstUser {
	const sqlStr = "select * a from dst_user where id=?"
	user := &dstUser{}
	err := db.DB.Get(user, sqlStr, id)
	if checkQueryErr(err, sqlStr, fmt.Sprintf("No User:%s[%s]", id, sqlStr)) {
		return nil
	}
	return user
}

type dstUserAddr struct {
	Addr     string `db:"addr"`
	UserID   string `db:"userid"`
	UserName string `db:"username"`
}

func FindUserByAddr(addr string) *dstUserAddr {
	const sqlStr = "select a.addr,a.userid,b.name as username from dst_user_addr a,dst_user b where a.userid=b.id and a.addr=?"
	userAddr := &dstUserAddr{}
	err := db.DB.Get(userAddr, sqlStr, addr)
	if checkQueryErr(err, sqlStr, fmt.Sprintf("No User By Addr:%s[%s]", addr, sqlStr)) {
		return nil
	}
	return userAddr
}
