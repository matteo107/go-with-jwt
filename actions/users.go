package actions

import "github.com/gobuffalo/buffalo"

// UsersLogin default implementation.
func UsersLogin(c buffalo.Context) error {
	return c.Render(200, r.HTML("users/login.html"))
}
