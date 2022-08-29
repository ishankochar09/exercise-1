package product

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

func TestAddProduct(t *testing.T) {
	ctx, mock, db := NewMock()
	defer db.Close()
	p:=NewProductRepo()

	input := models.Product{ID: "1",  Name:"Parle Frooti 200ml",BrandName:"Parle", Details:"really tasty",ImageURL:"www.googlefruity.com"}

	mock.ExpectExec("insert into product(name, brandName,details,imageURL) values (?,?,?,?);").
		WithArgs(input.Name, input.BrandName, input.Details, input.ImageURL).
		WillReturnResult(sqlmock.NewResult(1, 1))

	tests := []struct {
		desc string
		err  error
	}{
		{
			desc: "Success",
			err:  nil,
		},
	}
	for _, tc := range tests {
		_, err := p.AddProduct(ctx, &input)
		assert.Equal(t, tc.err, err, tc.desc)
	}
}

func TestGetProduct(t *testing.T){
	ctx, mock, db := NewMock()
	defer db.Close()
	p:=NewProductRepo()

	input := models.Product{ID: "1",  Name:"Parle Frooti 200ml",BrandName:"Parle", Details:"really tasty",ImageURL:"www.googlefruity.com"}
	rows := sqlmock.NewRows([]string{"id", "name", "brandName", "details", "imageURL"}).
		AddRow(input.ID, input.Name, input.BrandName, input.Details, input.ImageURL)

		mock.ExpectQuery("select id, name, brandName, details, imageURL from product where id = ?").
		WithArgs(input.ID).WillReturnRows(rows)

		tests := []struct {
			desc string
			id   int
			out  models.Product
			err  error
		}{
			{
				desc: "success",
				id:   1,
				out:  models.Product{},
				err:  nil,
			},
		}
		for _, tc := range tests {
			res := p.GetProduct(ctx,&input,tc.id )
			assert.Equal(t, tc.out, res, tc.desc)
		}
}