package model

import (
	"time"
)

//数据库表结构
type Users struct {
	Uid   int64     `json:"uid"`
	Name  string    `json:"name"`
	Age   int64     `json:"age"`
	Ctime time.Time `json:"ctime"`
	Mtime time.Time `json:"mtime"`
}
type Users1 struct {
	Uid   int64
	Name  string
	Age   int64
	Ctime time.Time
	Mtime time.Time
}
