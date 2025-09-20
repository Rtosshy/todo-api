package router

import (
	"fmt"
	"net/http"
	"os"
	"todo-api/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	feURL := os.Getenv("FE_URL")
	fmt.Printf("▶︎ Loaded FE_URL: %q\n", feURL)

	apiDomain := os.Getenv("API_DOMAIN")
	fmt.Printf("▶︎ Loaded API_DOMAIN: %q\n", apiDomain)

	env := os.Getenv("ENV")
	isProd := env == "prod"
	sameSite := http.SameSiteLaxMode
	if !isProd {
		sameSite = http.SameSiteNoneMode
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", feURL},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   apiDomain,
		CookieHTTPOnly: true,
		CookieSameSite: sameSite,
		CookieMaxAge:   60 * 60 * 24,
	}))
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)
	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllTasks)
	t.GET("/:task_id", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:task_id", tc.UpdateTask)
	t.DELETE("/:task_id", tc.DeleteTask)
	return e
}
