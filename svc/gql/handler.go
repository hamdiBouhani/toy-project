package gql

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"toy-project/common"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/sirupsen/logrus"
)

// Handler a graphql Handle responds to an HTTP request.
type Handler struct {
	Schema               *graphql.Schema
	Log                  logrus.FieldLogger
	EnableDashboardCache bool
}

// Serve implementation function of graphql handler
func (h *Handler) Serve(c *gin.Context) {
	r := c.Request
	w := c.Writer

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.Log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	schema := string(body)
	res, err := h.Query(c, schema)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func (h *Handler) Query(c *gin.Context, schema string) (*graphql.Response, error) {
	startTime := time.Now().Unix()
	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}
	if err := json.NewDecoder(strings.NewReader(schema)).Decode(&params); err != nil {
		return nil, err
	}

	ctx := c.Request.Context()
	res := h.Schema.Exec(context.WithValue(ctx, common.RequestHeaderCtx, c.Request.Header), params.Query, params.OperationName, params.Variables)

	h.Log.Infof("schema:%v, time cost:%v", schema, time.Now().Unix()-startTime)

	return res, nil
}
