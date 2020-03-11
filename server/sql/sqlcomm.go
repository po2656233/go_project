/*
# -*- coding: utf-8 -*-
# @Time : 2020/3/10 23:13
# @Author : Pitter
# @File : sqlcomm.go
# @Software: GoLand
*/
package sql

import "database/sql"

// 数据库的常规接口
type ISqlOP interface {
	Connect(user, psw, addr, port, db string) (*sql.DB, error) //连接数据库
	Check(statement interface{}) interface{}
	Add(statement interface{}) interface{}
	Delete(statement interface{}) interface{}
	Modify(statement interface{}) interface{}
	CloseMysql()
}
