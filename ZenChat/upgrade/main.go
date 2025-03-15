// ZenChat db upgrade

package main

import (
	"flag"
	"fmt"
	"github.com/mazezen/zenchat/common/sdk"
	"github.com/mazezen/zenchat/internel/models"
	"os"
	"xorm.io/xorm"
)

var c string

var db *xorm.Engine

func init() {
	flag.StringVar(&c, "c", "./config.yaml", "config file")
}

func main() {
	fmt.Println("*******************************")
	fmt.Println("***********开始同步表结构*********")
	sdk.ParseConfig(c)
	connect()

	if err := db.Sync2(
		new(models.CUser),
		new(models.CUserRelation),
		new(models.CCommunity),
		new(models.CCommunityRelation),
		new(models.CMessage),
	); err != nil {
		fmt.Printf("同步表结构失败err: %v\n", err)
		os.Exit(-1)
	}
	if err := db.CreateIndexes(new(models.CUser)); err != nil {
		fmt.Printf("创建c_user表索引失败err: %v\n", err)
		os.Exit(-1)
	}

	fmt.Println("***********同步表结构结束*********")
	fmt.Printf("已同步的表:%s\n", user())
	fmt.Printf("已同步的表:%s\n", UserRelation())
	fmt.Printf("已同步的表:%s\n", community())
	fmt.Printf("已同步的表:%s\n", communityRelation())
	fmt.Printf("已同步的表:%s\n", message())
	fmt.Println("*******************************")
}

func user() string {
	u := new(models.CUser)
	return u.TableName()
}

func UserRelation() string      { return new(models.CUserRelation).TableName() }
func community() string         { return new(models.CCommunity).TableName() }
func communityRelation() string { return new(models.CCommunityRelation).TableName() }
func message() string           { return new(models.CMessage).TableName() }

func connect() {
	engine, err := xorm.NewEngine("mysql", compact())
	if err != nil {
		fmt.Printf("connet mysql error: %v\n", err)
		os.Exit(-1)
	}

	db = engine
	if err = db.Ping(); err != nil {
		fmt.Printf("connet mysql error: %v\n", err)
		os.Exit(-1)
	}
}

func compact() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		sdk.GetConf().Db.Username, sdk.GetConf().Db.Password, sdk.GetConf().Db.Host, sdk.GetConf().Db.Port, sdk.GetConf().Db.Database, sdk.GetConf().Db.Charset)
}
