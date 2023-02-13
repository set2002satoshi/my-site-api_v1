package infrastructure

import (
	"github.com/gin-gonic/gin"
	bc "github.com/set2002satoshi/my-site-api/interfaces/controllers/blog"
	uc "github.com/set2002satoshi/my-site-api/interfaces/controllers/user"
	"github.com/set2002satoshi/my-site-api/pkg/module/service/auth"
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
	blogsController := bc.NewBlogController(r.DB)

	userNotLoggedIn := r.Gin.Group("/api")
	{
		// user
		userNotLoggedIn.POST("/user/get", func(c *gin.Context) { usersController.FindById(c) })
		userNotLoggedIn.POST("/users", func(c *gin.Context) { usersController.FindAll(c) })
		userNotLoggedIn.POST("/users/create", func(c *gin.Context) { usersController.Create(c) })

		userNotLoggedIn.POST("/login", func(c *gin.Context) { usersController.Login(c) })
	}

	userLoggedIn := r.Gin.Group("/api")
	userLoggedIn.Use(auth.CheckLoggedIn())
	{
		// user
		userLoggedIn.POST("/users/update", func(c *gin.Context) { usersController.Update(c) })
		userLoggedIn.POST("/users/delete", func(c *gin.Context) { usersController.Delete(c) })

	}

	blogNotLoggedIn := r.Gin.Group("/api")
	{
		blogNotLoggedIn.POST("/blog", func(c *gin.Context) { blogsController.Find(c) })
		blogNotLoggedIn.POST("/blog/get", func(c *gin.Context) { blogsController.FindById(c) })
	}
	
	blogLoggedIn := r.Gin.Group("/api")
	blogLoggedIn.Use(auth.CheckLoggedIn())
	{
		// blog
		blogLoggedIn.POST("/blog/create", func(c *gin.Context) { blogsController.Create(c) })
		blogLoggedIn.POST("/blog/delete", func(c *gin.Context) { blogsController.Delete(c) })
	}

}

func (r *Routing) Run() {
	err := r.Gin.Run(r.Port)
	if err != nil {
		panic(err)
	}
}
