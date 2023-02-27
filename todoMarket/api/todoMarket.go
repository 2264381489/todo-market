package main

import (
	"flag"
	"fmt"
	"todo-market/todoMarket/api/internal/model"

	"todo-market/todoMarket/api/internal/config"
	"todo-market/todoMarket/api/internal/handler"
	"todo-market/todoMarket/api/internal/svc"

	"github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/todoMarket.yaml", "the config file")
var svcCtx *gorm.DB

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx = model.InitModel(c.DataSource)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
