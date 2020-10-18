// +build mock
// Code generated by server_generator. DO NOT EDIT.
// generated version: unknown

package user2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

	props "github.com/go-generalize/api_gen/server_generator/sample/props"
)

// MockRoutes ...
type MockRoutes struct {
	router *echo.Group
}

// apiGenMockOption api-gen mock options
type apiGenMockOption struct {
	WaitMS     int64  `json:"wait_ms"`
	TargetFile string `json:"target_file"`
}

// NewRoutes ...
func NewMockRoutes(p *props.ControllerProps, router *echo.Group, jsonDir string) *MockRoutes {
	r := &MockRoutes{
		router: router,
	}
	{
		jd := filepath.Join(jsonDir, "/get_user/")
		router.GET(":userID", r.GetUser(p, jd))
	}
	{
		jd := filepath.Join(jsonDir, "/post_update_user_name/")
		router.POST("update_user_name", r.PostUpdateUserName(p, jd))
	}
	{
		jd := filepath.Join(jsonDir, "/post_update_user_password/")
		router.POST("update_user_password", r.PostUpdateUserPassword(p, jd))
	}
	return r
}

// GetUser ...
func (r *MockRoutes) GetUser(p *props.ControllerProps, jsonDir string) echo.HandlerFunc {
	type Mock struct {
		Meta struct {
			Status       int             `json:"status"`
			MatchRequest *GetUserRequest `json:"match_request"`
		} `json:"meta"`
		Payload interface{} `json:"payload"`
	}
	return func(c echo.Context) error {
		req := new(GetUserRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}

		option := &apiGenMockOption{}
		xago := c.Request().Header.Get("X-Api-Gen-Option")
		if xago != "" {
			if err := json.Unmarshal([]byte(xago), option); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid X-Api-Gen-Option.",
				})
			}
		}

		if option.WaitMS > 0 {
			ticker := time.NewTicker(time.Duration(option.WaitMS) * time.Millisecond)
			<-ticker.C
			ticker.Stop()
		}

		jsons := make(map[string]*Mock)
		err := filepath.Walk(jsonDir, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			mock := new(Mock)

			j, err := ioutil.ReadFile(path)
			if err != nil {
				log.Printf("SKIP: load mock json error in %s: %+v", path, err)
				return nil
			}

			err = json.Unmarshal(j, mock)
			if err != nil {
				log.Printf("SKIP: unmarshal mock json error in %s: %+v", path, err)
				return nil
			}

			jsons[info.Name()] = mock
			log.Printf("load %s", path)

			return nil
		})
		if err != nil {
			m := fmt.Sprintf("jsons load error: %+v", err)
			log.Println(m)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": m,
			})
		}

		var resMock *Mock = nil
		if option.TargetFile != "" {
			target := option.TargetFile
			if !strings.HasSuffix(target, ".json") {
				target += ".json"
			}
			mock, ok := jsons[target]
			if ok {
				resMock = mock
			}
		} else {
			jsonNameList := make([]string, 0, len(jsons))
			for key := range jsons {
				jsonNameList = append(jsonNameList, key)
			}
			sort.Strings(jsonNameList)

			for _, jsonName := range jsonNameList {
				m := jsons[jsonName]
				if !reflect.DeepEqual(m.Meta.MatchRequest, req) {
					continue
				}
				resMock = jsons[jsonName]
				log.Printf("[%s] Return the %s because it match rule.", jsonName, jsonName)
				break
			}
		}
		if resMock == nil {
			mock, ok := jsons["default.json"]
			if !ok {
				m := fmt.Sprintf("not match request and not found default.json")
				log.Println(m)
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    http.StatusInternalServerError,
					"message": m,
				})
			}
			log.Println("[default.json] Return the default.json because it did not match rule.")
			resMock = mock
		}

		return c.JSON(resMock.Meta.Status, resMock.Payload)
	}
}

// PostUpdateUserName ...
func (r *MockRoutes) PostUpdateUserName(p *props.ControllerProps, jsonDir string) echo.HandlerFunc {
	type Mock struct {
		Meta struct {
			Status       int                        `json:"status"`
			MatchRequest *PostUpdateUserNameRequest `json:"match_request"`
		} `json:"meta"`
		Payload interface{} `json:"payload"`
	}
	return func(c echo.Context) error {
		req := new(PostUpdateUserNameRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}

		option := &apiGenMockOption{}
		xago := c.Request().Header.Get("X-Api-Gen-Option")
		if xago != "" {
			if err := json.Unmarshal([]byte(xago), option); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid X-Api-Gen-Option.",
				})
			}
		}

		if option.WaitMS > 0 {
			ticker := time.NewTicker(time.Duration(option.WaitMS) * time.Millisecond)
			<-ticker.C
			ticker.Stop()
		}

		jsons := make(map[string]*Mock)
		err := filepath.Walk(jsonDir, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			mock := new(Mock)

			j, err := ioutil.ReadFile(path)
			if err != nil {
				log.Printf("SKIP: load mock json error in %s: %+v", path, err)
				return nil
			}

			err = json.Unmarshal(j, mock)
			if err != nil {
				log.Printf("SKIP: unmarshal mock json error in %s: %+v", path, err)
				return nil
			}

			jsons[info.Name()] = mock
			log.Printf("load %s", path)

			return nil
		})
		if err != nil {
			m := fmt.Sprintf("jsons load error: %+v", err)
			log.Println(m)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": m,
			})
		}

		var resMock *Mock = nil
		if option.TargetFile != "" {
			target := option.TargetFile
			if !strings.HasSuffix(target, ".json") {
				target += ".json"
			}
			mock, ok := jsons[target]
			if ok {
				resMock = mock
			}
		} else {
			jsonNameList := make([]string, 0, len(jsons))
			for key := range jsons {
				jsonNameList = append(jsonNameList, key)
			}
			sort.Strings(jsonNameList)

			for _, jsonName := range jsonNameList {
				m := jsons[jsonName]
				if !reflect.DeepEqual(m.Meta.MatchRequest, req) {
					continue
				}
				resMock = jsons[jsonName]
				log.Printf("[%s] Return the %s because it match rule.", jsonName, jsonName)
				break
			}
		}
		if resMock == nil {
			mock, ok := jsons["default.json"]
			if !ok {
				m := fmt.Sprintf("not match request and not found default.json")
				log.Println(m)
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    http.StatusInternalServerError,
					"message": m,
				})
			}
			log.Println("[default.json] Return the default.json because it did not match rule.")
			resMock = mock
		}

		return c.JSON(resMock.Meta.Status, resMock.Payload)
	}
}

// PostUpdateUserPassword ...
func (r *MockRoutes) PostUpdateUserPassword(p *props.ControllerProps, jsonDir string) echo.HandlerFunc {
	type Mock struct {
		Meta struct {
			Status       int                            `json:"status"`
			MatchRequest *PostUpdateUserPasswordRequest `json:"match_request"`
		} `json:"meta"`
		Payload interface{} `json:"payload"`
	}
	return func(c echo.Context) error {
		req := new(PostUpdateUserPasswordRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}

		option := &apiGenMockOption{}
		xago := c.Request().Header.Get("X-Api-Gen-Option")
		if xago != "" {
			if err := json.Unmarshal([]byte(xago), option); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid X-Api-Gen-Option.",
				})
			}
		}

		if option.WaitMS > 0 {
			ticker := time.NewTicker(time.Duration(option.WaitMS) * time.Millisecond)
			<-ticker.C
			ticker.Stop()
		}

		jsons := make(map[string]*Mock)
		err := filepath.Walk(jsonDir, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			mock := new(Mock)

			j, err := ioutil.ReadFile(path)
			if err != nil {
				log.Printf("SKIP: load mock json error in %s: %+v", path, err)
				return nil
			}

			err = json.Unmarshal(j, mock)
			if err != nil {
				log.Printf("SKIP: unmarshal mock json error in %s: %+v", path, err)
				return nil
			}

			jsons[info.Name()] = mock
			log.Printf("load %s", path)

			return nil
		})
		if err != nil {
			m := fmt.Sprintf("jsons load error: %+v", err)
			log.Println(m)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": m,
			})
		}

		var resMock *Mock = nil
		if option.TargetFile != "" {
			target := option.TargetFile
			if !strings.HasSuffix(target, ".json") {
				target += ".json"
			}
			mock, ok := jsons[target]
			if ok {
				resMock = mock
			}
		} else {
			jsonNameList := make([]string, 0, len(jsons))
			for key := range jsons {
				jsonNameList = append(jsonNameList, key)
			}
			sort.Strings(jsonNameList)

			for _, jsonName := range jsonNameList {
				m := jsons[jsonName]
				if !reflect.DeepEqual(m.Meta.MatchRequest, req) {
					continue
				}
				resMock = jsons[jsonName]
				log.Printf("[%s] Return the %s because it match rule.", jsonName, jsonName)
				break
			}
		}
		if resMock == nil {
			mock, ok := jsons["default.json"]
			if !ok {
				m := fmt.Sprintf("not match request and not found default.json")
				log.Println(m)
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    http.StatusInternalServerError,
					"message": m,
				})
			}
			log.Println("[default.json] Return the default.json because it did not match rule.")
			resMock = mock
		}

		return c.JSON(resMock.Meta.Status, resMock.Payload)
	}
}
