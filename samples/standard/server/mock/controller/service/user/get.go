// Code generated by api_gen. DO NOT EDIT.
// generated version: devel
// nolint:dupl
package controller

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"time"

	types "github.com/go-generalize/api_gen/samples/standard/api/service/user"
	mock "github.com/go-generalize/api_gen/samples/standard/server/mock"
	"github.com/labstack/echo/v4"
)

// GetController ...
type GetController interface {
	Get(c echo.Context, req *types.GetRequest) (res *types.GetResponse, err error)
}

type getController struct {
}

// NewGetController ...
func NewGetController(cp interface{}) GetController {
	return &getController{}
}

// Get - GET
// @Accept json
// @Produce json
// @Success 200 {object} types.GetResponse
// @Failure 500 {object} apierror.APIError
// @Router /service/user/ [get]
func (ctrl *getController) Get(
	c echo.Context, req *types.GetRequest,
) (res *types.GetResponse, err error) {
	const jsonExt = ".json"

	option := &mock.HeaderOption{}
	ago := c.Request().Header.Get("Api-Gen-Option")
	if ago != "" {
		if err := json.Unmarshal([]byte(ago), option); err != nil {
			log.Printf("failed to JSON Unmarshal to `Api-Gen-Option` header(/service/user/): %+v", err)
			return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid Api-Gen-Option.",
			})
		}
	}

	if option.UseMatchRequest == nil {
		flag := true
		option.UseMatchRequest = &flag
	}

	if option.WaitMS > 0 {
		<-time.After(time.Duration(option.WaitMS) * time.Millisecond)
	}

	type Mock struct {
		Meta struct {
			Status       int               `json:"status"`
			MatchRequest *types.GetRequest `json:"match_request"`
		} `json:"meta"`
		Payload interface{} `json:"payload"`
	}

	jsons := make(map[string]*Mock)
	err = fs.WalkDir(mock.MockJSONFS, "json/service/user/get_", fs.WalkDirFunc(func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fp, err := mock.MockJSONFS.Open(path)
		if err != nil {
			log.Printf("SKIP: load mock json error in %s: %+v", path, err)
			return nil
		}
		defer fp.Close()

		mock := new(Mock)

		err = json.NewDecoder(fp).Decode(mock)
		if err != nil {
			log.Printf("SKIP: unmarshal mock json error in %s: %+v", path, err)
			return nil
		}

		jsons[info.Name()] = mock
		log.Printf("load %s", path)

		return nil
	}))

	if err != nil {
		m := fmt.Sprintf("jsons load error: %+v", err)
		log.Println(m)
		return nil, c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": m,
		})
	}

	var resMock *Mock = nil
	if option.TargetFile != "" {
		target := option.TargetFile
		if !strings.HasSuffix(target, jsonExt) {
			target += jsonExt
		}
		mock, ok := jsons[target]
		if ok {
			resMock = mock
		}
	} else if *option.UseMatchRequest {
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
			log.Printf("[%s] Return the %s because it matches a rule.", jsonName, jsonName)
			break
		}
	}
	if resMock == nil {
		mock, ok := jsons["default.json"]
		if !ok {
			m := fmt.Sprintf("not match request and not found default.json")
			log.Println(m)
			return nil, c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": m,
			})
		}

		if *option.UseMatchRequest {
			log.Println("[default.json] Return the default.json because it did not match rule.")
		} else {
			log.Println("[default.json] Return the default.json because use_match_request is disabled.")
		}

		resMock = mock
	}

	return nil, c.JSON(resMock.Meta.Status, resMock.Payload)
}

// AutoBind - use echo.Bind
func (ctrl *getController) AutoBind() bool {
	return true
}
