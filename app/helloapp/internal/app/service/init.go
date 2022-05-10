package service

import (
	"github.com/photowey/hellocmd/app/helloapp/internal/app/model"
	"github.com/photowey/hellocmd/app/helloapp/internal/app/service/user"
	"github.com/photowey/hellocmd/app/helloapp/internal/app/service/user/impl"
)

var (
	UserRepository user.Repository
)

// Init instantiate the service
func Init() {
	UserRepository = impl.NewMysqlImpl(model.MysqlHandler)
}
