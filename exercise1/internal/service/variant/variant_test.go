package variant

import (
	"net/http"
	"net/http/httptest"
	reflect "reflect"
	"strconv"
	"testing"
	"trainig/exercise1/internal/models"
	"trainig/exercise1/internal/service"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	gomock "github.com/golang/mock/gomock"
)

func TestAddVariant(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	VariantMockInterface := service.NewMockVariantStore(mockCtrl)
	VariantService := NewVariantService((VariantMockInterface))
	testCases := []struct {
		desc          string
		mock          interface{}
		expOut        int
		productID     string
		expectedError error
		expVariant    *models.Variant
	}{
		{
			desc:          "success",
			mock:          VariantMockInterface.EXPECT().AddVariant(gomock.Any(), gomock.Any()).Return(1, nil),
			expOut:        1,
			productID:     "1",
			expectedError: nil,
			expVariant: &models.Variant{
				ProductID: "1",
			},
		},
		{
			desc:          "empty ID",
			expOut:        0,
			expectedError: errors.MissingParam{Param: []string{"product_id"}},
			expVariant:    &models.Variant{},
		},
	}
	for _, tc := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/product/"+tc.productID+"/variant", nil)
		r := request.NewHTTPRequest(req)
		k := gofr.NewContext(nil, r, gofr.New())

		t.Run("testing create service", func(t *testing.T) {
			output, err := VariantService.AddVariant(k, tc.expVariant)

			if !reflect.DeepEqual(err, tc.expectedError) {
				t.Errorf(" Expected Error: %v , Found : %v", tc.expectedError, err)
			}

			if !reflect.DeepEqual(output, tc.expOut) {
				t.Errorf(" Expected Output: %v , Found : %v", tc.expOut, output)
			}
		})

	}
}

func TestGetVariant(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	VariantMockInterface := service.NewMockVariantStore(mockCtrl)
	VariantService := NewVariantService((VariantMockInterface))
	testCases := []struct {
		desc          string
		mock          interface{}
		expOut        *models.Variant
		expectedError error
		productID     int
		varID         int
	}{
		{
			desc:          "success case",
			mock:          VariantMockInterface.EXPECT().GetVariant(gomock.Any(), gomock.Any(), gomock.Any()).Return(&models.Variant{}),
			expOut:        &models.Variant{},
			expectedError: nil,
			productID:     1,
		},
	}
	for _, tc := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/product/"+strconv.Itoa(tc.productID)+"/variant/"+strconv.Itoa(tc.varID), nil)

		r := request.NewHTTPRequest(req)
		k := gofr.NewContext(nil, r, gofr.New())

		t.Run("testing create service", func(t *testing.T) {
			output, err := VariantService.GetVariant(k, strconv.Itoa(tc.productID), strconv.Itoa(tc.varID))

			if !reflect.DeepEqual(err, tc.expectedError) {
				t.Errorf(" Expected Error: %v , Found : %v", tc.expectedError, err)
			}

			if !reflect.DeepEqual(tc.expOut, output) {
				t.Errorf(" Expected Output: %v , Found : %v", tc.expOut, output)
			}
		})

	}
}
