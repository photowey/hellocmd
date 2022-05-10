package user

import (
	"github.com/photowey/hellocmd/app/helloapps/internal/app/model/db"
)

type Repository interface {
	Get(id int) (user db.User, err error)
}
