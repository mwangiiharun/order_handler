package server

import (
	"context"
	"encoding/json"
	"io/fs"

	"github.com/qri-io/jsonschema"
)

type ValidationFactory struct {
	method string
	fs     fs.FS
	schema *jsonschema.Schema
}

func (v *ValidationFactory) GetSchema(method string) (*jsonschema.Schema, error) {
	if v.schema != nil && v.method == method {
		return v.schema, nil
	}

	fileInfo, err := fs.Stat(v.fs, method+".json")
	if err != nil {
		return nil, err
	}

	schemaData, err := fs.ReadFile(v.fs, fileInfo.Name())
	if err != nil {
		return nil, err
	}

	rs := &jsonschema.Schema{}
	if err := json.Unmarshal(schemaData, rs); err != nil {
		return nil, err
	}

	v.schema = rs
	return rs, nil
}

func (v *ValidationFactory) Validate(ctx context.Context, data []byte) []error {
	if v.schema == nil {
		_, err := v.GetSchema(v.method)
		if err != nil {
			return []error{err}
		}
	}

	keyErrors, err := v.schema.ValidateBytes(ctx, data)
	if err != nil {
		return []error{err}
	}

	if len(keyErrors) > 0 {
		errs := make([]error, len(keyErrors))
		for i, ke := range keyErrors {
			errs[i] = ke
		}
		return errs
	}

	return nil
}

func NewValidationFactory(ctx context.Context, method string, embed fs.FS) (*ValidationFactory, error) {
	factory := &ValidationFactory{
		method: method,
		fs:     embed,
	}
	_, err := factory.GetSchema(method)
	if err != nil {
		return nil, err
	}
	return factory, nil
}
