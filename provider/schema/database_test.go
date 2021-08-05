// Code generated by mockery v0.0.0-dev. DO NOT EDIT.
package schema

import (
	"context"
	"testing"

	pgx "github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

type jsonTestType struct {
	Name        string `json:"name"`
	Description string `json:"decription"`
	Version     int    `json:"version"`
}

var (
	stringJson    = "{\"test\":true}"
	jsonTestTable = Table{
		Name: "test_table_validator",
		Columns: []Column{
			{
				Name: "test",
				Type: TypeJSON,
			},
		},
	}
	resources = []Resource{
		{
			data: map[string]interface{}{
				"test": stringJson,
				"meta": make(map[string]string),
			},
			table: &jsonTestTable,
		},
		{
			data: map[string]interface{}{
				"test": &stringJson,
			},
			table: &jsonTestTable,
		},
		{
			data: map[string]interface{}{
				"test": map[string]interface{}{
					"test": 1,
					"test1": map[string]interface{}{
						"test": 1,
					},
				},
			},
			table: &jsonTestTable,
		},
		{
			data: map[string]interface{}{
				"test": []interface{}{
					map[string]interface{}{
						"test":  1,
						"test1": true,
					},
					map[string]interface{}{
						"test":  1,
						"test1": true,
					},
				},
			},
			table: &jsonTestTable,
		},
		{
			data: map[string]interface{}{
				"test": nil,
			},
			table: &jsonTestTable,
		},
		{
			data: map[string]interface{}{
				"test": []interface{}{
					nil,
				},
			},
			table: &jsonTestTable,
		},
		{
			data: map[string]interface{}{
				"test": jsonTestType{
					Name:        "test",
					Description: "test1",
					Version:     10,
				},
			},
			table: &jsonTestTable,
		},
	}
)

// databaseMock is an autogenerated mock type for the databaseMock type
type databaseMock struct {
	mock.Mock
}

// CopyFrom provides a mock function with given fields: ctx, resources, shouldCascade, CascadeDeleteFilters
func (_m *databaseMock) CopyFrom(ctx context.Context, resources Resources, shouldCascade bool, CascadeDeleteFilters map[string]interface{}) error {
	ret := _m.Called(ctx, resources, shouldCascade, CascadeDeleteFilters)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, Resources, bool, map[string]interface{}) error); ok {
		r0 = rf(ctx, resources, shouldCascade, CascadeDeleteFilters)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, t, args
func (_m *databaseMock) Delete(ctx context.Context, t *Table, args []interface{}) error {
	ret := _m.Called(ctx, t, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *Table, []interface{}) error); ok {
		r0 = rf(ctx, t, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exec provides a mock function with given fields: ctx, query, args
func (_m *databaseMock) Exec(ctx context.Context, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) error); ok {
		r0 = rf(ctx, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Insert provides a mock function with given fields: ctx, t, instance
func (_m *databaseMock) Insert(ctx context.Context, t *Table, instance Resources) error {
	ret := _m.Called(ctx, t, instance)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *Table, Resources) error); ok {
		r0 = rf(ctx, t, instance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Query provides a mock function with given fields: ctx, query, args
func (_m *databaseMock) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 pgx.Rows
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) pgx.Rows); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Rows)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func TestJsonColumn(t *testing.T) {
	for _, r := range resources {
		_, err := getResourceValues(&r)
		assert.Nil(t, err)
	}
}
