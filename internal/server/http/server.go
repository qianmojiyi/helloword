package http

import (
	// "errors"
	"fmt"
	"helloword/internal/model"
	"net/http"

	"go-common/library/conf/paladin.v2"
	"go-common/library/log"
	bm "go-common/library/net/http/blademaster"
	pb "helloword/api"

	"helloword/internal/service"
)

var svc *service.Service

// New new a bm server.
func New(s *service.Service) (engine *bm.Engine, err error) {
	var (
		cfg bm.ServerConfig
		ct  paladin.TOML
	)
	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	svc = s
	engine = bm.DefaultServer(&cfg)
	pb.RegisterDemoBMServer(engine, s)
	initRouter(engine)
	err = engine.Start()
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/helloword")
	{
		g.GET("/hello", howToStart1)
		g.GET("/adduser", AddUser)
		g.GET("/searchmap", SearchUser)
		g.GET("/searchrow", SearchStructUser)
		g.GET("/search", SearchStruct)
		g.GET("/update", UpdateUser)
		g.GET("/delete", DeleteUser)
	}
}

func ping(ctx *bm.Context) {
	if _, err := svc.Ping(ctx, nil); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
func howToStart1(c *bm.Context) {
	k := &model.Kratos{
		Hello: "Hello Kratos!",
	}
	fmt.Println("howToStart1:", k)
	c.JSON(k, nil)
}

// example for http request handler.
func AddUser(c *bm.Context) {
	id, err := svc.InsertUser()
	if err != nil {
		c.JSON(0, err)
		return
	}
	c.JSON(id, nil)
}

func SearchUser(c *bm.Context) {
	users, err := svc.SearchUser()
	if err != nil {
		c.JSON("error:", err)
		return
	}
	c.JSON(users, nil)
}

//SearchStructUser row单条查询
func SearchStructUser(c *bm.Context) {
	users, err := svc.SearchStructUser()
	if err != nil {
		c.JSON("error:", err)
		return
	}
	c.JSON(users, nil)
}

//SearchStruct 多条查询
func SearchStruct(c *bm.Context) {
	users, err := svc.SearchStruct()
	if err != nil {
		c.JSON("error:", err)
		return
	}
	fmt.Println("howt5:", users)
	c.JSON(users, nil)
}

func UpdateUser(c *bm.Context) {
	fmt.Println("111")
	id, err := svc.UpdateUser()
	if err != nil {
		c.JSON(0, err)
		return
	}
	c.JSON(id, nil)
}

func DeleteUser(c *bm.Context) {
	fmt.Println("111")
	id, err := svc.DeleteUser()
	if err != nil {
		c.JSON(0, err)
		return
	}
	c.JSON(id, nil)
}
