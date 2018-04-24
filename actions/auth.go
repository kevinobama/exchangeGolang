package actions

import (

	"github.com/gobuffalo/buffalo"
 
 
 
)

// AuthNew loads the signin page
func AuthNew(c buffalo.Context) error {
	//c.Set("user", models.User{})
	//return c.Render(200, r.HTML("auth/new.html"))
	return c.Render(200, r.JSON(map[string]string{"message": "Welcome to Buffalo!"}))
}
