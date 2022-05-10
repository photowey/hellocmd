package user

import (
	"github.com/photowey/hellocmd/app/helloapp/internal/app/model/db"
)

type Repository interface {
	Get(id int) (user db.User, err error)
}
