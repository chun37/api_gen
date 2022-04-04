// Code generated by api_gen. DO NOT EDIT.
// Package controller binds handlers to echo
// generated version: devel
package main

import (
	"net/http"
	"runtime/debug"
	"strings"

	apierror "github.com/go-generalize/api_gen/v2/pkg/server/tmp/output/pkg/apierror"
	props "github.com/go-generalize/api_gen/v2/pkg/server/tmp/output/props"
	"github.com/go-utils/echo-multipart-binder/mjbinder"
	"github.com/go-utils/echo-multipart-binder/mpbinder"
	"github.com/labstack/echo/v4"
)

type middleware struct {
	path       string
	middleware echo.MiddlewareFunc
}

type options struct {
	disableErrorHandling bool
}

// Controllers binds handlers to echo
type Controllers struct {
	middlewares []middleware
	options
}

// NewControllers returns a new Controllers
func NewControllers(
	props *props.ControllerProps, e *echo.Echo,
) *Controllers {
	ctrl := &Controllers{}

	e.Binder = mjbinder.NewMultipartJSONBinder(mpbinder.NewMultipartFileBinder(e.Binder))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			for _, m := range ctrl.middlewares {
				if strings.HasPrefix(c.Request().URL.Path, m.path) {
					next = m.middleware(next)
				}
			}

			return next(c)
		}
	})
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			defer func() {
				recoverErr := recover()
				if recoverErr == nil {
					return
				}

				debug.PrintStack()

				if httpErr, ok := recoverErr.(*echo.HTTPError); ok {
					err = c.JSON(httpErr.Code, httpErr.Message)

					return
				}

				err = apierror.NewAPIError(http.StatusInternalServerError, "internal server error.")
			}()

			return next(c)
		}
	})

	addRoutes(e, props, &ctrl.options)

	return ctrl
}

// AddMiddleware adds 'm' to the paths starting with the 'path'.
func (c *Controllers) AddMiddleware(path string, m echo.MiddlewareFunc) {
	c.middlewares = append(c.middlewares, middleware{
		path:       path,
		middleware: m,
	})
}

// DisableErrorHandling disables
func (c *Controllers) DisableErrorHandling() {
	c.disableErrorHandling = true
}
