// Code generated by server_generator. DO NOT EDIT.
// generated version: 0.4.0

package user2

import (
	"net/http"

	"github.com/labstack/echo/v4"

	props "github.com/go-generalize/api_gen/server_generator/sample/props"
)

type Routes struct {
	router *echo.Group
}

func NewRoutes(p *props.ControllerProps, router *echo.Group) *Routes {
	r := &Routes{
		router: router,
	}

	router.GET(":userID", r.GetUser(p))
	router.POST("update_user_name", r.PostUpdateUserName(p))
	router.POST("update_user_password", r.PostUpdateUserPassword(p))

	return r
}

func (r *Routes) GetUser(p *props.ControllerProps) echo.HandlerFunc {
	i := NewGetUserController(p)
	return func(c echo.Context) error {
		req := new(GetUserRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}
		res, err := i.GetUser(c, req)
		if err != nil {
			return err
		}
		if res == nil {
			return nil
		}

		return c.JSON(http.StatusOK, res)
	}
}

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

type IGetUserController interface {
	GetUser(c echo.Context, req *GetUserRequest) (res *GetUserResponse, err error)
}

type IPostUpdateUserNameController interface {
	PostUpdateUserName(c echo.Context, req *PostUpdateUserNameRequest) (res *PostUpdateUserNameResponse, err error)
}

type IPostUpdateUserPasswordController interface {
	PostUpdateUserPassword(c echo.Context, req *PostUpdateUserPasswordRequest) (res *PostUpdateUserPasswordResponse, err error)
}
