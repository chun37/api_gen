//go:build !mock
// +build !mock

// Code generated by api_gen. DO NOT EDIT.
// generated version: (devel)
package controller

import (
	"net/http"

	types_api_e3c6fe7f "github.com/go-generalize/api_gen/v2/samples/standard/api"
	types_service_004a0f6b "github.com/go-generalize/api_gen/v2/samples/standard/api/service"
	types_groups_3384feb9 "github.com/go-generalize/api_gen/v2/samples/standard/api/service/groups"
	types_static_page_88819fa5 "github.com/go-generalize/api_gen/v2/samples/standard/api/service/static_page"
	types_user_b2fb2440 "github.com/go-generalize/api_gen/v2/samples/standard/api/service/user"
	types_user2_fc0fa272 "github.com/go-generalize/api_gen/v2/samples/standard/api/service/user2"
	types__userID_341bfdbf "github.com/go-generalize/api_gen/v2/samples/standard/api/service/user2/_userID"
	types__JobID_550441cf "github.com/go-generalize/api_gen/v2/samples/standard/api/service/user2/_userID/_JobID"
	ctrl_controller_dcb773a3 "github.com/go-generalize/api_gen/v2/samples/standard/server/controller"
	ctrl_service_348fab3e "github.com/go-generalize/api_gen/v2/samples/standard/server/controller/service"
	ctrl_groups_82ababd5 "github.com/go-generalize/api_gen/v2/samples/standard/server/controller/service/groups"
	ctrl_static_page_918c2446 "github.com/go-generalize/api_gen/v2/samples/standard/server/controller/service/static_page"
	ctrl_user_471e7c77 "github.com/go-generalize/api_gen/v2/samples/standard/server/controller/service/user"
	ctrl_user2_390d1774 "github.com/go-generalize/api_gen/v2/samples/standard/server/controller/service/user2"
	ctrl__userID_880730e5 "github.com/go-generalize/api_gen/v2/samples/standard/server/controller/service/user2/_userID"
	ctrl__JobID_7ce20d9f "github.com/go-generalize/api_gen/v2/samples/standard/server/controller/service/user2/_userID/_JobID"
	apierror "github.com/go-generalize/api_gen/v2/samples/standard/server/pkg/apierror"
	props "github.com/go-generalize/api_gen/v2/samples/standard/server/props"
	echo "github.com/labstack/echo/v4"
	xerrors "golang.org/x/xerrors"
)

func addRoutes(e *echo.Echo, p *props.ControllerProps, opt *options) {
	add := func(method, path string, handler func(c echo.Context) (interface{}, error)) {
		e.Add(method, path, func(c echo.Context) error {
			var werr *apierror.APIError

			res, err := handler(c)

			if err != nil {
				if !opt.disableErrorHandling && xerrors.As(err, &werr) {
					c.Logger().Errorf("%+v", werr)
					return c.JSON(werr.Status, werr.Body)
				}
				return xerrors.Errorf("%s %s returned an error: %w", method, path, err)
			}
			if res == nil {
				return nil
			}

			return c.JSON(http.StatusOK, res)
		})
	}

	{
		ctrl := ctrl_controller_dcb773a3.NewGetController(p)

		add("GET", "/", func(c echo.Context) (interface{}, error) {
			req := new(types_api_e3c6fe7f.GetRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				return nil, xerrors.Errorf("the validator returned an error: %w", err)
			}

			res, err := ctrl.Get(c, req)

			if err != nil {
				return nil, xerrors.Errorf("the handler(Get) returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_controller_dcb773a3.NewPostCreateTableController(p)

		add("POST", "/create_table", func(c echo.Context) (interface{}, error) {
			req := new(types_api_e3c6fe7f.PostCreateTableRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/create_table): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				return nil, xerrors.Errorf("the validator returned an error: %w", err)
			}

			res, err := ctrl.PostCreateTable(c, req)

			if err != nil {
				return nil, xerrors.Errorf("the handler(PostCreateTable) returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_controller_dcb773a3.NewPostCreateUserController(p)

		add("POST", "/create_user", func(c echo.Context) (interface{}, error) {
			req := new(types_api_e3c6fe7f.PostCreateUserRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/create_user): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				return nil, xerrors.Errorf("the validator returned an error: %w", err)
			}

			res, err := ctrl.PostCreateUser(c, req)

			if err != nil {
				return nil, xerrors.Errorf("the handler(PostCreateUser) returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_service_348fab3e.NewGetArticleController(p)

		add("GET", "/service/article", func(c echo.Context) (interface{}, error) {
			req := new(types_service_004a0f6b.GetArticleRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/article): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				return nil, xerrors.Errorf("the validator returned an error: %w", err)
			}

			res, err := ctrl.GetArticle(c, req)

			if err != nil {
				return nil, xerrors.Errorf("the handler(GetArticle) returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_groups_82ababd5.NewGetGroupsController(p)

		add("GET", "/service/groups/groups", func(c echo.Context) (interface{}, error) {
			req := new(types_groups_3384feb9.GetGroupsRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/groups/groups): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				return nil, xerrors.Errorf("the validator returned an error: %w", err)
			}

			res, err := ctrl.GetGroups(c, req)

			if err != nil {
				return nil, xerrors.Errorf("the handler(GetGroups) returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_static_page_918c2446.NewGetStaticPageController(p)

		add("GET", "/service/static_page/static_page", func(c echo.Context) (interface{}, error) {
			req := new(types_static_page_88819fa5.GetStaticPageRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/static_page/static_page): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				return nil, xerrors.Errorf("the validator returned an error: %w", err)
			}

			res, err := ctrl.GetStaticPage(c, req)

			if err != nil {
				return nil, xerrors.Errorf("the handler(GetStaticPage) returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_user_471e7c77.NewGetController(p)

		add("GET", "/service/user/", func(c echo.Context) (interface{}, error) {
			req := new(types_user_b2fb2440.GetRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/user/): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				return nil, xerrors.Errorf("the validator returned an error: %w", err)
			}

			res, err := ctrl.Get(c, req)

			if err != nil {
				return nil, xerrors.Errorf("the handler(Get) returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_user_471e7c77.NewPostUpdateUserNameController(p)

		add("POST", "/service/user/update_user_name", func(c echo.Context) (interface{}, error) {
			req := new(types_user_b2fb2440.PostUpdateUserNameRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/user/update_user_name): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				return nil, xerrors.Errorf("the validator returned an error: %w", err)
			}

			res, err := ctrl.PostUpdateUserName(c, req)

			if err != nil {
				return nil, xerrors.Errorf("the handler(PostUpdateUserName) returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_user_471e7c77.NewPostUpdateUserPasswordController(p)

		add("POST", "/service/user/update_user_password", func(c echo.Context) (interface{}, error) {
			req := new(types_user_b2fb2440.PostUpdateUserPasswordRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/user/update_user_password): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				return nil, xerrors.Errorf("the validator returned an error: %w", err)
			}

			res, err := ctrl.PostUpdateUserPassword(c, req)

			if err != nil {
				return nil, xerrors.Errorf("the handler(PostUpdateUserPassword) returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_user2_390d1774.NewDeleteUserController(p)

		add("DELETE", "/service/user2/:user_id", func(c echo.Context) (interface{}, error) {
			req := new(types_user2_fc0fa272.DeleteUserRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/user2/:user_id): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				return nil, xerrors.Errorf("the validator returned an error: %w", err)
			}

			res, err := ctrl.DeleteUser(c, req)

			if err != nil {
				return nil, xerrors.Errorf("the handler(DeleteUser) returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_user2_390d1774.NewGetUserController(p)

		add("GET", "/service/user2/:user_id", func(c echo.Context) (interface{}, error) {
			req := new(types_user2_fc0fa272.GetUserRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/user2/:user_id): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				return nil, xerrors.Errorf("the validator returned an error: %w", err)
			}

			res, err := ctrl.GetUser(c, req)

			if err != nil {
				return nil, xerrors.Errorf("the handler(GetUser) returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_user2_390d1774.NewPostUpdateUserNameController(p)

		add("POST", "/service/user2/update_user_name", func(c echo.Context) (interface{}, error) {
			req := new(types_user2_fc0fa272.PostUpdateUserNameRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/user2/update_user_name): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				return nil, xerrors.Errorf("the validator returned an error: %w", err)
			}

			res, err := ctrl.PostUpdateUserName(c, req)

			if err != nil {
				return nil, xerrors.Errorf("the handler(PostUpdateUserName) returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_user2_390d1774.NewPostUpdateUserPasswordController(p)

		add("POST", "/service/user2/update_user_password", func(c echo.Context) (interface{}, error) {
			req := new(types_user2_fc0fa272.PostUpdateUserPasswordRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/user2/update_user_password): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				return nil, xerrors.Errorf("the validator returned an error: %w", err)
			}

			res, err := ctrl.PostUpdateUserPassword(c, req)

			if err != nil {
				return nil, xerrors.Errorf("the handler(PostUpdateUserPassword) returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl__userID_880730e5.NewGetUserJobGetController(p)

		add("GET", "/service/user2/:userID/user_job_get", func(c echo.Context) (interface{}, error) {
			req := new(types__userID_341bfdbf.GetUserJobGetRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/user2/:userID/user_job_get): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				return nil, xerrors.Errorf("the validator returned an error: %w", err)
			}

			res, err := ctrl.GetUserJobGet(c, req)

			if err != nil {
				return nil, xerrors.Errorf("the handler(GetUserJobGet) returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl__JobID_7ce20d9f.NewPutJobController(p)

		add("PUT", "/service/user2/:userID/:JobID/job", func(c echo.Context) (interface{}, error) {
			req := new(types__JobID_550441cf.PutJobRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/user2/:userID/:JobID/job): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				return nil, xerrors.Errorf("the validator returned an error: %w", err)
			}

			res, err := ctrl.PutJob(c, req)

			if err != nil {
				return nil, xerrors.Errorf("the handler(PutJob) returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

}
