package database_test

import (
	"context"
	"fmt"
	"reflect"
	"sekareco_srv/usecase/database"
	"testing"
)

var tr database.SqlTransaction

func TestTx_Do(t *testing.T) {
	ctx := context.Background()
	t.Run("Do test", func(t *testing.T) {
		got, err := tr.Do(ctx, func(ctx context.Context) (interface{}, error) {
			return 1, nil
		})
		if (err != nil) != false {
			t.Errorf("tx.Do() error = %v, wantErr %v", err, false)
			return
		}
		if !reflect.DeepEqual(got, 1) {
			t.Errorf("tx.Do() = %v, want %v", got, 1)
		}
	})

	t.Run("Do test failed", func(t *testing.T) {
		got, err := tr.Do(ctx, func(ctx context.Context) (interface{}, error) {
			return nil, fmt.Errorf("")
		})
		if (err != nil) != true {
			t.Errorf("tx.Do() error = %v, wantErr %v", err, false)
			return
		}
		if !reflect.DeepEqual(got, nil) {
			t.Errorf("tx.Do() = %v, want %v", got, 1)
		}
	})
}
