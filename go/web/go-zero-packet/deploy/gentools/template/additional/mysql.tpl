package query

import (
	"context"
	"errors"
    "{{.ModelPackPath}}"
	"gorm.io/gen"
)

type {{.ModelStructName}} struct {
	{{.QueryStructName}}
}

func New{{.ModelStructName}}(p {{.QueryStructName}}) {{.ModelStructName}} {
	return {{.ModelStructName}}{
	    {{.QueryStructName}}: p,
	}
}

func (p *{{.QueryStructName}}) ParseWhere(where map[string]interface{}) []gen.Condition {
	eqConds := make([]gen.Condition, 0)
	for key, value := range where {
        switch key {
            {{.Switch}}
        }
	}

	return eqConds
}

func (p *{{.QueryStructName}}) Find(ctx context.Context, limit, offset int, where map[string]interface{}) ([]*model.{{.ModelStructName}}, error) {
	conds := p.ParseWhere(where)
	_db := p.WithContext(ctx)
	if limit != 0 {
		_db = _db.Limit(limit)
	}
	if offset != 0 {
		_db = _db.Offset(offset)
	}
	if where != nil {
		_db = _db.Where(conds...)
	}

	return _db.Find()
}

func (p *{{.QueryStructName}}) FindOne(ctx context.Context, where map[string]interface{}) (*model.{{.ModelStructName}}, error) {
	conds := p.ParseWhere(where)
	_db := p.WithContext(ctx)

	if where != nil {
		_db = _db.Where(conds...)
	}

	return _db.Take()
}

func (p *{{.QueryStructName}}) UpdateMap(ctx context.Context, where, update map[string]interface{}) (info gen.ResultInfo, err error) {
	conds := p.ParseWhere(where)
	if len(conds) == 0 {
		return gen.ResultInfo{}, errors.New("The update condition cannot be empty")
	}
	info, err = p.WithContext(ctx).Where(conds...).Updates(update)
	return info, err
}
