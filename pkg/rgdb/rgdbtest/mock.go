package rgdbtest

import (
	"errors"
	"github.com/juicyluv/rgdb/pkg/rgdb"
	"github.com/juicyluv/rgutils/pkg/logger"
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
	RowsError       error
	WaitResponseFor time.Duration
}

func PrepareMock(t *testing.T, v ExpectedBehaviour) *rgdb.Client {
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
		WillReturnError(v.RowsError).
		WillDelayFor(v.WaitResponseFor).
		RowsWillBeClosed()

	client := &rgdb.Client{
		Logger: logger.New(&logger.Config{LogToConsole: true}),
		Driver: mock,
	}

	return client
}

func CheckResult(t *testing.T, actual, expected interface{}, actualError, expectedError error) {
	if !errors.Is(actualError, expectedError) && !reflect.DeepEqual(actualError, expectedError) {
		t.Fatalf("Mismatch error\n\nGot:%+v\nExpected:%+v", actualError, expectedError)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Mismatch response\n\nGot:%+v\nExpected:%+v", actual, expected)
	}
}
