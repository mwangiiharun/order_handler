package server_test

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"testing"

	"github.com/order_handler/pkg/server"
)

//go:embed internal/validationtest/testdata/*.json
var testdata embed.FS

func Schemas() (fs.FS, error) {
	return fs.Sub(testdata, "internal/validationtest/testdata")
}

func TestListAllEmbededValidationFiles(t *testing.T) {
	schemas, e := Schemas()
	if e != nil {
		t.Errorf("Schema() error = %v", e)
		return
	}
	err := fs.WalkDir(schemas, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		t.Errorf("fs.WalkDir() error = %v", err)
	}
}

func TestValidationFactory_GetSchema_Sucess(t *testing.T) {
	schemas, e := Schemas()
	if e != nil {
		t.Errorf("Schema() error = %v", e)
		return
	}
	ctx := context.Background()
	vf, err := server.NewValidationFactory(ctx, "orders.order.add", schemas)
	if err != nil {
		t.Errorf("NewValidationFactory() error = %v", err)
		return
	}
	if vf == nil {
		t.Errorf("GetSchema() schema is nil")
	}

}

func TestValidationFactory_GetSchema_Fail(t *testing.T) {
	schemas, e := Schemas()
	if e != nil {
		t.Errorf("Schema() error = %v", e)
		return
	}
	ctx := context.Background()
	vf, err := server.NewValidationFactory(ctx, "orders.order.add", schemas)
	if err != nil {
		t.Errorf("NewValidationFactory() error = %v", err)
		return
	}
	if vf == nil {
		t.Errorf("GetSchema() schema is nil")
	}

	_, err = vf.GetSchema("orders.orders.add")
	if err == nil {
		t.Errorf("GetSchema() expected not found error")
	}
}

func TestValidationFactory_Validate_Sucess(t *testing.T) {
	schemas, e := Schemas()
	if e != nil {
		t.Errorf("Schema() error = %v", e)
		return
	}
	ctx := context.Background()
	vf, err := server.NewValidationFactory(ctx, "orders.order.add", schemas)
	if err != nil {
		t.Errorf("NewValidationFactory() error = %v", err)
		return
	}
	if vf == nil {
		t.Errorf("GetSchema() schema is nil")
	}

	data := []byte(`{"items": [
    {
      "itemId": "item001",
      "quantity": 2
    },
    {
      "itemId": "item002",
      "quantity": 1
    }
  ]}`)
	errs := vf.Validate(ctx, data)
	if len(errs) > 0 {
		t.Errorf("Validate() error = %v", errs[0])
	}
}


func TestValidationFactory_Validate_Fail(t *testing.T) {
	schemas, e := Schemas()
	if e != nil {
		t.Errorf("Schema() error = %v", e)
		return
	}
	ctx := context.Background()
	vf, err := server.NewValidationFactory(ctx, "orders.order.add", schemas)
	if err != nil {
		t.Errorf("NewValidationFactory() error = %v", err)
		return
	}
	if vf == nil {
		t.Errorf("GetSchema() schema is nil")
	}

	data := []byte(`{"items": [
    {
      "itemId": "item001",
      "quantit": 2
    },
    {
      "iteId": "item002",
      "quantity": 1
    }
  ]}`)
	errs := vf.Validate(ctx, data)
	if len(errs) != 2 {
		t.Errorf("Validate() error = %v", errs[0])
	}
}