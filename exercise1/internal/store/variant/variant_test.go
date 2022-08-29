package variant

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"trainig/exercise1/internal/models"

	"developer.zopsmart.com/go/gofr/pkg/datastore"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*gofr.Context, sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	ctx := gofr.NewContext(nil, nil, &gofr.Gofr{DataStore: datastore.DataStore{ORM: db}})
	ctx.Context = context.Background()

	return ctx, mock, db
}

func TestGetVariant(t *testing.T) {
	ctx, mock, db := NewMock()
	defer db.Close()
	v := NewVariantRepo()

	input := models.Variant{ID: "1", ProductID: "1", VariantName: "parle", VariantDetails: "biscuit"}
	rows := sqlmock.NewRows([]string{"id", "product_id", "variantName", "variantDetails"}).
		AddRow(input.ID, input.ProductID, input.VariantName, input.VariantDetails)

	mock.ExpectQuery("select id, product_id, variantName, variantDetails from variant where product_id = ? and id = ?").
		WithArgs(input.ID, input.ProductID).WillReturnRows(rows)

	tests := []struct {
		desc string
		id   string
		out  *models.Variant
		err  error
	}{
		{
			desc: "success",
			id:   "1",
			out:  &models.Variant{ID: "1", ProductID: "1", VariantName: "parle", VariantDetails: "biscuit"},
			err:  nil,
		},
	}
	for _, tc := range tests {
		res := v.GetVariant(ctx, input.ProductID, tc.id)
		assert.Equal(t, tc.out, res, tc.desc)
	}
}
