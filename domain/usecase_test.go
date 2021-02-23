package domain_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/pkg/errors"

	. "github.com/zitryss/perfmon/domain"
	"github.com/zitryss/perfmon/infrastructure/database"
)

var (
	dbMem = database.NewMem()
)

func TestGetProduct(t *testing.T) {
	use := NewUsecase(dbMem)
	products, err := use.Products()
	if err != nil {
		t.Errorf("error occurred: %+v", err)
	}
	if len(products) == 0 {
		t.Error("no product")
	} else {
		fmt.Println(fmt.Sprintf("Product size: %v", len(products)))
	}
}

func TestGetVersion(t *testing.T) {
	testCases := []struct {
		product string
		success bool
	}{
		{
			product: "p1",
			success: true,
		},
		{
			product: "p2",
			success: true,
		},
		{
			product: "p3",
			success: true,
		},
		{
			product: "p4",
			success: true,
		},
		{
			product: "p5",
			success: true,
		},
		{
			product: "p6",
			success: true,
		},
		{
			product: "p7",
			success: true,
		},
	}
	use := NewUsecase(dbMem)
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			versions, err := use.Versions(tc.product)
			if tc.success {
				if err != nil {
					t.Errorf("error occurred: %+v", err)
				}
			} else {
				if err == nil {
					t.Errorf("error occurred: %+v", err)
				}
			}
			fmt.Println(versions)
		})
	}
}

func TestUsecaseJobsMem(t *testing.T) {
	use := NewUsecase(dbMem)
	t.Run("", func(t *testing.T) {
		product := "p6"
		version := "v6"
		attributes := []string{"a61", "a62"}
		name := "n6"
		measurement := "m6"
		lbound := time.Date(2017, 8, 19, 8, 40, 23, 0, time.UTC)
		rbound := time.Date(2018, 7, 2, 6, 34, 54, 0, time.UTC)
		jobs, err := use.Jobs(product, version, attributes, name, measurement, lbound, rbound)
		if err != nil {
			t.Errorf("error occurred: %+v", err)
		}
		got := len(jobs)
		want := 3
		if got != want {
			t.Errorf("len(use.Jobs(%v, %v, %v, %v, %v, %v, %v)) = %v; want %v", product, version, attributes, name, measurement, lbound, rbound, got, want)
		}
	})
	t.Run("", func(t *testing.T) {
		product := "p7"
		version := "v7"
		attributes := []string{"a71", "a72"}
		name := "n7"
		measurement := "m7"
		lbound := time.Date(2018, 2, 8, 0, 51, 55, 0, time.UTC)
		rbound := time.Date(2018, 2, 16, 2, 16, 51, 0, time.UTC)
		_, gotErr := use.Jobs(product, version, attributes, name, measurement, lbound, rbound)
		wantErr := ErrNotEnoughAttributes
		if errors.Cause(gotErr) != wantErr {
			t.Errorf("use.Jobs(%v, %v, %v, %v, %v, %v, %v) = %v; want %v", product, version, attributes, name, measurement, lbound, rbound, gotErr, wantErr)
		}
	})
}
