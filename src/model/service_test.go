package model

import (
	"context"
	"net/http"
	"testing"

	mockSvc "github.com/monitoring-service/src/communication/mock"
	"github.com/monitoring-service/src/entity"
	mockMetric "github.com/monitoring-service/src/web/metric/mock"

	"github.com/golang/mock/gomock"
)

func TestProcess(t *testing.T) {
	var (
		urls = []string{"test_url"}
		ctx  = context.Background()
	)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSvc := mockSvc.NewMockService(mockCtrl)
	mockPublisher := mockMetric.NewMockPublisher(mockCtrl)

	t.Run("Test_Success_200", func(t *testing.T) {
		mockSvc.EXPECT().
			MonitorService(ctx, urls[0]).
			Return(entity.ServiceResponse{StatusCode: http.StatusOK}, nil).
			Times(len(urls))

		mockPublisher.EXPECT().
			PublishResponseStatus(gomock.Any(), gomock.Any()).
			Return().
			Times(len(urls))

		mockPublisher.EXPECT().
			PublishResponseTime(gomock.Any(), gomock.Any()).
			Return().
			Times(len(urls))

		svc := NewServiceModel(urls, mockPublisher, mockSvc)
		svc.Process(ctx)
	})

	t.Run("Test_Success_503", func(t *testing.T) {
		mockSvc.EXPECT().
			MonitorService(ctx, urls[0]).
			Return(entity.ServiceResponse{StatusCode: http.StatusServiceUnavailable}, nil).
			Times(len(urls))

		mockPublisher.EXPECT().
			PublishResponseStatus(gomock.Any(), gomock.Any()).
			Return().
			Times(len(urls))

		mockPublisher.EXPECT().
			PublishResponseTime(gomock.Any(), gomock.Any()).
			Return().
			Times(len(urls))

		svc := NewServiceModel(urls, mockPublisher, mockSvc)
		svc.Process(ctx)
	})
}
