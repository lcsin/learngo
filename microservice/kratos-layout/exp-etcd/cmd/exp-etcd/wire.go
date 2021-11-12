// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"exp-etcd/internal/biz"
	"exp-etcd/internal/conf"
	"exp-etcd/internal/data"
	"exp-etcd/internal/server"
	"exp-etcd/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger, *conf.Registry) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
