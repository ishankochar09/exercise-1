package product

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

func TestAddProduct(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ProductMockInterface := service.NewMockProductStore(mockCtrl)
	ProductService := NewProductService((ProductMockInterface))
	testCases := []struct {
		desc          string
		mock          interface{}
		expOut        int
		expectedError error
		expProduct    *models.Product
	}{
		{
			desc:          "success",
			mock:          ProductMockInterface.EXPECT().AddProduct(gomock.Any(), gomock.Any()).Return(1, nil),
			expOut:        1,
			expectedError: nil,
			expProduct: &models.Product{
				Name: "ishan",
			},
		},
		{
			desc:          "empty name",
			expOut:        0,
			expectedError: errors.MissingParam{Param: []string{"name"}},
			expProduct:    &models.Product{},
		},
		{
			desc:          "add product failure",
			mock:          ProductMockInterface.EXPECT().AddProduct(gomock.Any(), gomock.Any()).Return(0, errors.Error("error")),
			expOut:        0,
			expectedError: errors.Error("error"),
			expProduct: &models.Product{
				Name: "ishan",
			},
		},
	}
	for _, tc := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/product/", nil)
		r := request.NewHTTPRequest(req)
		k := gofr.NewContext(nil, r, gofr.New())

		t.Run("testing create service", func(t *testing.T) {
			output, err := ProductService.AddProduct(k, tc.expProduct)

			if !reflect.DeepEqual(err, tc.expectedError) {
				t.Errorf(" Expected Error: %v , Found : %v", tc.expectedError, err)
			}

			if !reflect.DeepEqual(output, tc.expOut) {
				t.Errorf(" Expected Output: %v , Found : %v", tc.expOut, output)
			}
		})

	}
}

func TestGetProduct(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ProductMockInterface := service.NewMockProductStore(mockCtrl)
	ProductService := NewProductService((ProductMockInterface))
	testCases := []struct {
		desc          string
		mock          interface{}
		expOut        models.Products
		expectedError error
		productID     int
	}{
		{
			desc: "success case",
			mock: []interface{}{ProductMockInterface.EXPECT().GetProduct(gomock.Any(), gomock.Any(), gomock.Any()).Return(models.Product{}),
				ProductMockInterface.EXPECT().GetProductVariants(gomock.Any(), gomock.Any()).Return([]models.Variant{}),
			},
			expOut: models.Products{
				Product: models.Product{},
				Variant: []models.Variant{},
			},
			expectedError: nil,
			productID:     1,
		},
	}
	for _, tc := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/product/"+strconv.Itoa(tc.productID), nil)

		r := request.NewHTTPRequest(req)
		k := gofr.NewContext(nil, r, gofr.New())

		t.Run("testing create service", func(t *testing.T) {
			output, err := ProductService.GetProduct(k, &models.Product{}, tc.productID)

			if !reflect.DeepEqual(err, tc.expectedError) {
				t.Errorf(" Expected Error: %v , Found : %v", tc.expectedError, err)
			}

			if !reflect.DeepEqual(tc.expOut, output) {
				t.Errorf(" Expected Output: %v , Found : %v", tc.expOut, output)
			}
		})

	}
}
