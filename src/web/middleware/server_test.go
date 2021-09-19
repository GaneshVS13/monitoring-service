package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"

	"github.com/monitoring-service/src/model/mock"
)

func TestServeHTTP(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockServiceModel := mock.NewMockServiceHandler(mockCtrl)

	t.Run("Test_Success_200", func(t *testing.T) {
		router := mux.NewRouter()

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/metrics", nil)

		mockServiceModel.EXPECT().Process(gomock.Any()).Return()

		middleware := NewMiddleware(router, mockServiceModel)
		middleware.ServeHTTP(rec, req)
	})
}
