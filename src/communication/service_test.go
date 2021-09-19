package communication_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/monitoring-service/src/communication"
	"github.com/monitoring-service/src/communication/rest/mock"
	"github.com/monitoring-service/src/entity"
)

func TestMonitorService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRestSvc := mock.NewMockRestService(mockCtrl)

	t.Run("Test_Success", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(entity.ContentType, entity.ApplicationJSON)
			w.WriteHeader(http.StatusOK)
		}))
		defer ts.Close()

		ctx := context.Background()

		mockRestSvc.EXPECT().Do(ctx, http.MethodGet, ts.URL, nil, nil).
			Return(nil, http.StatusOK, nil)

		svc := communication.NewService(mockRestSvc)
		_, err := svc.MonitorService(ctx, ts.URL)
		assert.Nil(t, err)
	})

	t.Run("Test_Response_Error", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer ts.Close()

		ctx := context.Background()

		mockRestSvc.EXPECT().Do(ctx, http.MethodGet, ts.URL, nil, nil).
			Return(nil, http.StatusBadRequest, errors.New("test error"))

		svc := communication.NewService(mockRestSvc)
		_, err := svc.MonitorService(ctx, ts.URL)
		assert.NotNil(t, err)
	})
}
