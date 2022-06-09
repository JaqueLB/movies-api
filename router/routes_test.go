package router

import (
	"bytes"
	"fmt"
	"io"
	"moviesapi/controller"
	"moviesapi/entity"
	"moviesapi/test/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	// dates setup
	location, err := time.LoadLocation("America/Sao_Paulo")
	if err == nil {
		time.Local = location
	}
	dateMock := time.Date(2022, time.June, 6, 12, 30, 0, 0, time.Local)
	dateStr := dateMock.Format("02/01/2006 15:04")

	// storage mocks
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	localDBMock := mock.NewMockIStorage(ctrl)

	localDBMock.EXPECT().List().Return([]entity.Movie{
		{
			ID:   1,
			Name: "Movie 1",
			Sessions: []entity.MovieSession{
				{
					DateTime: dateMock,
					Room:     1,
				},
			},
		},
		{
			ID:   2,
			Name: "Movie 2",
			Sessions: []entity.MovieSession{
				{
					DateTime: dateMock,
					Room:     2,
				},
			},
		},
		{
			ID:   3,
			Name: "Movie 3",
			Sessions: []entity.MovieSession{
				{
					DateTime: dateMock,
					Room:     3,
				},
			},
		},
	}).Times(1)
	localDBMock.EXPECT().GetByID(1).Return(&entity.Movie{
		ID:   1,
		Name: "Movie 1",
		Sessions: []entity.MovieSession{
			{
				DateTime: dateMock,
				Room:     1,
			},
		},
	}).Times(1)
	localDBMock.EXPECT().UpdateByID(2, gomock.Any()).Return(true).Times(1)
	localDBMock.EXPECT().DeleteByID(3).Return(true).Times(1)
	localDBMock.EXPECT().Create(gomock.Any()).Return(&entity.Movie{
		ID:   4,
		Name: "Movie Name",
		Sessions: []entity.MovieSession{
			{
				DateTime: dateMock,
				Room:     1,
			},
		},
	}).Times(1)

	// dependency injection and router creation
	controller := &controller.MovieController{
		Storage: localDBMock,
	}
	engine := gin.New()
	engine = CreateRoutes(engine, controller)

	// test data
	type args struct {
		path       string
		ID         int
		HTTPMethod string
		reqBody    string
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
			name: "Test get all movies",
			args: args{
				path:       "/movies",
				HTTPMethod: http.MethodGet,
				reqBody:    "",
			},
			want: want{
				code: http.StatusOK,
				body: `{"data":[{
                    "ID": 1,
                    "name": "Movie 1",
                    "sessions": [{
                        "datetime": "2022-06-06T12:30:00-03:00",
                        "room": 1
                    }]
                },
                {
                    "ID": 2,
                    "name": "Movie 2",
                    "sessions": [{
                        "datetime": "2022-06-06T12:30:00-03:00",
                        "room": 2
                    }]
                },
                {
                    "ID": 3,
                    "name": "Movie 3",
                    "sessions": [{
                        "datetime": "2022-06-06T12:30:00-03:00",
                        "room": 3
                    }]
                }]}`,
			},
		},
		{
			name: "Test create a movie",
			args: args{
				path:       "/movies",
				HTTPMethod: http.MethodPut,
				reqBody: fmt.Sprintf(`{
					"name": "Movie Name",
					"sessions": [{
                        "datetime": "%s",
                        "room": 1
                    }]
				}`, dateStr),
			},
			want: want{
				code: http.StatusOK,
				body: `{"data":{
                    "ID":4,
                    "name":"Movie Name",
                    "sessions":[{"datetime":"2022-06-06T12:30:00-03:00", "room":1}]}}`,
			},
		},
		{
			name: "Test get a movie",
			args: args{
				path:       fmt.Sprintf("/movies/%d", 1),
				ID:         1,
				HTTPMethod: http.MethodGet,
				reqBody:    "",
			},
			want: want{
				code: http.StatusOK,
				body: `{"data":{
                    "ID":1,
                    "name":"Movie 1",
                    "sessions":[{"datetime":"2022-06-06T12:30:00-03:00", "room":1}]}}`,
			},
		},
		{
			name: "Test update a movie",
			args: args{
				path:       fmt.Sprintf("/movies/%d", 2),
				ID:         2,
				HTTPMethod: http.MethodPost,
				reqBody: fmt.Sprintf(`{
					"name": "Updated movie Name",
					"sessions": [{
                        "datetime": "%s",
                        "room": 1
                    }]}`, dateStr),
			},
			want: want{
				code: http.StatusOK,
				body: `{"data":"OK"}`,
			},
		},
		{
			name: "Test delete a movie",
			args: args{
				path:       fmt.Sprintf("/movies/%d", 3),
				ID:         3,
				HTTPMethod: http.MethodDelete,
				reqBody:    "",
			},
			want: want{
				code: http.StatusOK,
				body: `{"data":"OK"}`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.args.HTTPMethod, tt.args.path, bytes.NewBuffer([]byte(tt.args.reqBody)))
			resp := httptest.NewRecorder()
			engine.ServeHTTP(resp, req)

			jsonBytes, _ := io.ReadAll(resp.Body)

			assert.Equal(t, tt.want.code, resp.Code)
			assert.JSONEq(t, tt.want.body, string(jsonBytes))
		})
	}
}
