package variant


import (
	"bytes"
	"net/http"
	"net/http/httptest"
	reflect "reflect"
	"testing"
	"trainig/exercise1/internal/models"
	"trainig/exercise1/internal/handler"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	gomock "github.com/golang/mock/gomock"
)

func TestAddVariant(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	MockInterface := handler.NewMockVariantService(mockCtrl)
	handler := NewVariantHandler(MockInterface)
	testCases := []struct {
		mock          *gomock.Call
		varID         string
		expOut        interface{}
		expectedError error
		body          []byte
	}{
		{
			mock: MockInterface.EXPECT().AddVariant(gomock.Any(), gomock.Any()).
				Return(0, nil),
			varID:         "1",
			expOut:        0,
			expectedError: nil,
			body: []byte(`{
				"variantName":"kurkure",
				"variantDetails":"red color"
			}`),
		},
		{
			expOut:        nil,
			expectedError: errors.InvalidParam{Param: []string{"id"}},
			body:          nil,
		},
	}

	for _, tc := range testCases {
		r := httptest.NewRequest("POST", "/variant", bytes.NewReader(tc.body))
		w := httptest.NewRecorder()

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, nil)
		ctx.SetPathParams(map[string]string{"pid": tc.varID})
		output, err := handler.AddVariant(ctx)

		if !reflect.DeepEqual(tc.expectedError, err) {
			t.Errorf("Expected error %v, got:%v", tc.expectedError, err)
		}

		if !reflect.DeepEqual(tc.expOut, output) {
			t.Errorf("Expected error %v, got:%v", tc.expOut, output)
		}
	}
}

func TestGetVariant(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	MockInterface := handler.NewMockVariantService(mockCtrl)
	handler := NewVariantHandler(MockInterface)
	tests := []struct {
		desc      string
		productid string
		varid     string
		mock      *gomock.Call
		expOut    interface{}
		Exerr     error
	}{
		{
			desc:      "Success",
			productid: "1",
			varid:     "2",
			mock:      MockInterface.EXPECT().GetVariant(gomock.Any(), gomock.Any(), gomock.Any()).Return(&models.Variant{}, nil),
			expOut:    &models.Variant{},
			Exerr:     nil,
		},
		{
			desc:      "Failure",
			productid: "1",
			varid:     "2",
			mock:      MockInterface.EXPECT().GetVariant(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil,errors.Error("error")),
			expOut:    nil,
			Exerr:     errors.Error("error"),
		},
	}

	for _, tc := range tests {
		link := "/product" + "/" + tc.productid + "/variant" + "/" + tc.varid
		r := httptest.NewRequest(http.MethodGet, link, nil)
		w := httptest.NewRecorder()

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, nil)
		ctx.SetPathParams(map[string]string{"pid": tc.productid})
		ctx.SetPathParams(map[string]string{"pid": tc.varid})

		output, err := handler.GetVariant(ctx)

		if !reflect.DeepEqual(tc.Exerr, err) {
			t.Errorf("Expected error %v, got:%v", tc.Exerr, err)
		}

		if !reflect.DeepEqual(tc.expOut, output) {
			t.Errorf("Expected output %v, got:%v", tc.expOut, output)
		}
	}
}
