package rgdbtest

import (
	"errors"
	"github.com/juicyluv/rgdb/pkg/rgdb"
	"github.com/pashagolub/pgxmock"
	"reflect"
	"testing"
	"time"
)

type ExpectedBehaviour struct {
	Query           string
	Args            []interface{}
	Columns         []string
	Rows            [][]interface{}
	ExpectedError   error
	WaitResponseFor time.Duration
}

func PrepareMock(t *testing.T, v ExpectedBehaviour) rgdb.Interface {
	dataSet := pgxmock.NewRows(v.Columns)

	for _, r := range v.Rows {
		dataSet.AddRow(r...)
	}

	mock, err := pgxmock.NewPool()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectQuery(v.Query).
		WithArgs(v.Args...).
		WillReturnRows(dataSet).
		WillReturnError(v.ExpectedError).
		WillDelayFor(v.WaitResponseFor).
		RowsWillBeClosed()

	return mock
}

func CheckResult(t *testing.T, actual, expected interface{}, actualError, expectedError error) {
	if !errors.Is(actualError, expectedError) && !reflect.DeepEqual(actualError, expectedError) {
		t.Fatalf("Mismatch error\n\nGot:%+v\nWant:%+v", actualError, expectedError)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Mismatch response\n\nGot:%+v\nWant:%+v", actual, expected)
	}
}
