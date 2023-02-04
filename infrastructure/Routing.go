package infrastructure

import (
	"github.com/gin-gonic/gin"
	uc "github.com/set2002satoshi/my-site-api/interfaces/controllers/user"
)

type Routing struct {
	DB   *DB
	Gin  *gin.Engine
	Port string
}

func NewRouting(db *DB) *Routing {
	r := &Routing{
		DB:   db,
		Gin:  gin.Default(),
		Port: ":8080",
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	usersController := uc.NewUserController(r.DB)
	userNotLoggedIn := r.Gin.Group("/api")
	{
		userNotLoggedIn.POST("/users", func(c *gin.Context) { usersController.FindAll(c) })
		userNotLoggedIn.POST("/user/id", func(c *gin.Context) { usersController.FindById(c) })
		userNotLoggedIn.POST("/users/create", func(c *gin.Context) { usersController.Create(c) })
		userNotLoggedIn.POST("/users/update", func(c *gin.Context) { usersController.Update(c) })
		userNotLoggedIn.POST("/users/delete", func(c *gin.Context) { usersController.Delete(c) })
	}
}

func (r *Routing) Run() {
	err := r.Gin.Run(r.Port)
	if err != nil {
		panic(err)
	}
}