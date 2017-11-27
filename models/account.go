package models

import (
	//	"database/sql"
	"time"
)

//Account account model
type Account struct {
	Id          int
	ImToken     string `xorm:" varchar(128) unique(imtoken) 'imtoken'"`
	CountryCode string `xorm:"varchar(20) null 'countrycode'",json:"AccountID"`
	PhoneNo     string `xorm:"varchar(20)  null unique(phoneno, openid, accountsrc )  'phoneno'"`
	OpenID      string `xorm:"varchar(64)  null unique(phoneno, openid, accountsrc)  'openid'"`
	//0：inner; 微信: 1; 微博:2 ; QQ :3
	AccountSrc   int8      `xorm:" tinyint   unique(phoneno, openid, accountsrc)  'accountsrc'"`
	Email        string    `xorm:" varchar(30)  null unique(email)  'email'"`
	PendingEmail string    `xorm:" varchar(30) null 'pendingemail'"` //新邮箱
	Password     string    `xorm:" varchar(50)  'password'"`
	CreateAt     time.Time `xorm:" created 'create_at'"`
	UpdateAt     time.Time `xorm:" updated 'update_at'"`
	DeletedAt    time.Time `xorm:" deleted 'deleted_at'"`
}

//TableName table name
func (m *Account) TableName() string {
	return "account"
}
