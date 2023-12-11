package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

{{if .Easy}}
const {{.Type}}CollectionName = "{{.snakeType}}"{{else}}const {{.Type}}CollectionName = "{{.snakeType}}"
{{end}}

var _ {{.Type}}Model = (*custom{{.Type}}Model)(nil)

type (
    // {{.Type}}Model is an interface to be customized, add more methods here,
    // and implement the added methods in custom{{.Type}}Model.
    {{.Type}}Model interface {
        {{.lowerType}}Model
        FindAll(ctx context.Context, filter map[string]interface{}, opts ...*options.FindOptions) ([]{{.Type}}, error)
        Count(ctx context.Context, filter map[string]interface{}, opts ...*options.CountOptions) (int64, error)
        FindOne(ctx context.Context, filter map[string]interface{}, opts ...*options.FindOneOptions) (*{{.Type}}, error)
    }

    custom{{.Type}}Model struct {
        *default{{.Type}}Model
    }

)


// New{{.Type}}Model returns a model for the mongo.
{{if .Easy}}func New{{.Type}}Model(cfg MonConfig, opts ...mon.Option) {{.Type}}Model {
    conn := mon.MustNewModel(cfg.GetURL(), cfg.GetDBName(), {{.Type}}CollectionName, opts...)
    return &custom{{.Type}}Model{
        default{{.Type}}Model: newDefault{{.Type}}Model(conn),
    }
}{{else}}func New{{.Type}}Model(cfg MonConfig, opts ...mon.Option) {{.Type}}Model {
    conn := mon.MustNewModel(cfg.GetURL(), cfg.GetDBName(), {{.Type}}CollectionName, opts...)
    return &custom{{.Type}}Model{
        default{{.Type}}Model: newDefault{{.Type}}Model(conn),
    }
}{{end}}

func (m *custom{{.Type}}Model) parseFilter(filter map[string]interface{}, flag ...string) bson.D {
	var flagType string
	if len(flag) > 0 {
		flagType = flag[0]
	}else {
		flagType = "default"
	}

	query := bson.D{}
	switch flagType {
	default:
		for k, v := range filter {
			query = append(query, bson.E{Key: k, Value: v})
		}
	}
	return query
}

func (m *custom{{.Type}}Model) Count(ctx context.Context, filter map[string]interface{}, opts ...*options.CountOptions) (int64, error) {
	query := m.parseFilter(filter)
	count, err := m.conn.CountDocuments(ctx, query, opts...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (m *custom{{.Type}}Model) FindAll(ctx context.Context, filter map[string]interface{}, opts ...*options.FindOptions) ([]{{.Type}}, error) {
	result := make([]{{.Type}}, 0)
	query := m.parseFilter(filter)
	err := m.conn.Find(ctx, &result, query, opts...)

	switch err {
	case nil:
		return result, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *custom{{.Type}}Model) FindOne(ctx context.Context, filter map[string]interface{}, opts ...*options.FindOneOptions) (*{{.Type}}, error) {
	var data {{.Type}}
	query := m.parseFilter(filter)
	err := m.conn.FindOne(ctx, &data, query, opts...)
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
