package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	type args struct {
		path       string
		HTTPMethod string
		reqBody    io.Reader
	}
	type want struct {
		code int
		body string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Test GET /movies",
			args: args{
				path:       "/movies",
				HTTPMethod: http.MethodGet,
				reqBody:    nil,
			},
			want: want{
				code: http.StatusOK,
				body: `{"data":null}`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			router = CreateRoutes(router)

			reqBody := tt.args.reqBody
			if tt.args.reqBody != nil {
				jsonBytes, _ := json.Marshal(tt.args.reqBody)
				reqBody = bytes.NewBuffer(jsonBytes)
			}

			req, _ := http.NewRequest(tt.args.HTTPMethod, tt.args.path, reqBody)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			jsonBytes, _ := io.ReadAll(resp.Body)

			assert.Equal(t, tt.want.code, resp.Code)
			assert.JSONEq(t, tt.want.body, string(jsonBytes))
		})
	}
}
