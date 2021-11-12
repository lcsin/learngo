// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"exp-yoka/internal/biz"
	"exp-yoka/internal/conf"
	"exp-yoka/internal/data"
	"exp-yoka/internal/server"
	"exp-yoka/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger, *conf.Registry) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
