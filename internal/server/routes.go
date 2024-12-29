package server

import (
	"encoding/gob"
	"log"
	"net/http"
	"smp15-oidc/authenticator"
	"smp15-oidc/cmd/api/callback"
	"smp15-oidc/cmd/api/login"
	"smp15-oidc/cmd/api/logout"
	"smp15-oidc/cmd/api/user"
	"smp15-oidc/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	// r.GET("/", s.HelloWorldHandler)

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("auth-session", store))

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	r.Static("/public", "cmd/static")
	r.LoadHTMLGlob("cmd/template/*")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})
	r.GET("/login", login.Handler(auth))
	r.GET("/callback", callback.Handler(auth))
	r.GET("/user", middleware.IsAuthenticated, user.Handler)
	r.GET("/logout", logout.Handler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World!!"

	c.JSON(http.StatusOK, resp)
}
