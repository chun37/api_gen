// Code generated by server_generator. DO NOT EDIT.
// generated version: unknown

package user

import (
	"net/http"

	"github.com/labstack/echo/v4"

	props "github.com/go-generalize/api_gen/server_generator/sample/props"
)

// Routes ...
type Routes struct {
	router *echo.Group
}

// NewRoutes ...
func NewRoutes(p *props.ControllerProps, router *echo.Group) *Routes {
	r := &Routes{
		router: router,
	}
	router.POST("update_user_name", r.PostUpdateUserName(p))
	router.POST("update_user_password", r.PostUpdateUserPassword(p))
	return r
}

// PostUpdateUserName ...
func (r *Routes) PostUpdateUserName(p *props.ControllerProps) echo.HandlerFunc {
	i := NewPostUpdateUserNameController(p)
	return func(c echo.Context) error {
		req := new(PostUpdateUserNameRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}
		res, err := i.PostUpdateUserName(c, req)
		if err != nil {
			return err
		}
		if res == nil {
			return nil
		}

		return c.JSON(http.StatusOK, res)
	}
}

// PostUpdateUserPassword ...
func (r *Routes) PostUpdateUserPassword(p *props.ControllerProps) echo.HandlerFunc {
	i := NewPostUpdateUserPasswordController(p)
	return func(c echo.Context) error {
		req := new(PostUpdateUserPasswordRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}
		res, err := i.PostUpdateUserPassword(c, req)
		if err != nil {
			return err
		}
		if res == nil {
			return nil
		}

		return c.JSON(http.StatusOK, res)
	}
}

// IPostUpdateUserNameController ...
type IPostUpdateUserNameController interface {
	PostUpdateUserName(c echo.Context, req *PostUpdateUserNameRequest) (res *PostUpdateUserNameResponse, err error)
}

// IPostUpdateUserPasswordController ...
type IPostUpdateUserPasswordController interface {
	PostUpdateUserPassword(c echo.Context, req *PostUpdateUserPasswordRequest) (res *PostUpdateUserPasswordResponse, err error)
}
