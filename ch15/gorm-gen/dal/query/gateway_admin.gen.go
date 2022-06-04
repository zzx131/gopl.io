// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gopl.io/ch15/gorm-gen/dal/model"
)

func newGatewayAdmin(db *gorm.DB) gatewayAdmin {
	_gatewayAdmin := gatewayAdmin{}

	_gatewayAdmin.gatewayAdminDo.UseDB(db)
	_gatewayAdmin.gatewayAdminDo.UseModel(&model.GatewayAdmin{})

	tableName := _gatewayAdmin.gatewayAdminDo.TableName()
	_gatewayAdmin.ALL = field.NewField(tableName, "*")
	_gatewayAdmin.ID = field.NewInt64(tableName, "id")
	_gatewayAdmin.UserName = field.NewString(tableName, "user_name")
	_gatewayAdmin.Salt = field.NewString(tableName, "salt")
	_gatewayAdmin.Password = field.NewString(tableName, "password")
	_gatewayAdmin.CreateTime = field.NewInt64(tableName, "create_time")
	_gatewayAdmin.UpdateTime = field.NewInt64(tableName, "update_time")
	_gatewayAdmin.DeletedAt = field.NewField(tableName, "deleted_at")

	_gatewayAdmin.fillFieldMap()

	return _gatewayAdmin
}

type gatewayAdmin struct {
	gatewayAdminDo gatewayAdminDo

	ALL        field.Field
	ID         field.Int64
	UserName   field.String
	Salt       field.String
	Password   field.String
	CreateTime field.Int64
	UpdateTime field.Int64
	DeletedAt  field.Field

	fieldMap map[string]field.Expr
}

func (g gatewayAdmin) Table(newTableName string) *gatewayAdmin {
	g.gatewayAdminDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g gatewayAdmin) As(alias string) *gatewayAdmin {
	g.gatewayAdminDo.DO = *(g.gatewayAdminDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *gatewayAdmin) updateTableName(table string) *gatewayAdmin {
	g.ALL = field.NewField(table, "*")
	g.ID = field.NewInt64(table, "id")
	g.UserName = field.NewString(table, "user_name")
	g.Salt = field.NewString(table, "salt")
	g.Password = field.NewString(table, "password")
	g.CreateTime = field.NewInt64(table, "create_time")
	g.UpdateTime = field.NewInt64(table, "update_time")
	g.DeletedAt = field.NewField(table, "deleted_at")

	g.fillFieldMap()

	return g
}

func (g *gatewayAdmin) WithContext(ctx context.Context) *gatewayAdminDo {
	return g.gatewayAdminDo.WithContext(ctx)
}

func (g gatewayAdmin) TableName() string { return g.gatewayAdminDo.TableName() }

func (g gatewayAdmin) Alias() string { return g.gatewayAdminDo.Alias() }

func (g *gatewayAdmin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *gatewayAdmin) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 7)
	g.fieldMap["id"] = g.ID
	g.fieldMap["user_name"] = g.UserName
	g.fieldMap["salt"] = g.Salt
	g.fieldMap["password"] = g.Password
	g.fieldMap["create_time"] = g.CreateTime
	g.fieldMap["update_time"] = g.UpdateTime
	g.fieldMap["deleted_at"] = g.DeletedAt
}

func (g gatewayAdmin) clone(db *gorm.DB) gatewayAdmin {
	g.gatewayAdminDo.ReplaceDB(db)
	return g
}

type gatewayAdminDo struct{ gen.DO }

func (g gatewayAdminDo) Debug() *gatewayAdminDo {
	return g.withDO(g.DO.Debug())
}

func (g gatewayAdminDo) WithContext(ctx context.Context) *gatewayAdminDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g gatewayAdminDo) Clauses(conds ...clause.Expression) *gatewayAdminDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g gatewayAdminDo) Returning(value interface{}, columns ...string) *gatewayAdminDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g gatewayAdminDo) Not(conds ...gen.Condition) *gatewayAdminDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g gatewayAdminDo) Or(conds ...gen.Condition) *gatewayAdminDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g gatewayAdminDo) Select(conds ...field.Expr) *gatewayAdminDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g gatewayAdminDo) Where(conds ...gen.Condition) *gatewayAdminDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g gatewayAdminDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *gatewayAdminDo {
	return g.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (g gatewayAdminDo) Order(conds ...field.Expr) *gatewayAdminDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g gatewayAdminDo) Distinct(cols ...field.Expr) *gatewayAdminDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g gatewayAdminDo) Omit(cols ...field.Expr) *gatewayAdminDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g gatewayAdminDo) Join(table schema.Tabler, on ...field.Expr) *gatewayAdminDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g gatewayAdminDo) LeftJoin(table schema.Tabler, on ...field.Expr) *gatewayAdminDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g gatewayAdminDo) RightJoin(table schema.Tabler, on ...field.Expr) *gatewayAdminDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g gatewayAdminDo) Group(cols ...field.Expr) *gatewayAdminDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g gatewayAdminDo) Having(conds ...gen.Condition) *gatewayAdminDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g gatewayAdminDo) Limit(limit int) *gatewayAdminDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g gatewayAdminDo) Offset(offset int) *gatewayAdminDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g gatewayAdminDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *gatewayAdminDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g gatewayAdminDo) Unscoped() *gatewayAdminDo {
	return g.withDO(g.DO.Unscoped())
}

func (g gatewayAdminDo) Create(values ...*model.GatewayAdmin) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g gatewayAdminDo) CreateInBatches(values []*model.GatewayAdmin, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g gatewayAdminDo) Save(values ...*model.GatewayAdmin) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g gatewayAdminDo) First() (*model.GatewayAdmin, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GatewayAdmin), nil
	}
}

func (g gatewayAdminDo) Take() (*model.GatewayAdmin, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GatewayAdmin), nil
	}
}

func (g gatewayAdminDo) Last() (*model.GatewayAdmin, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GatewayAdmin), nil
	}
}

func (g gatewayAdminDo) Find() ([]*model.GatewayAdmin, error) {
	result, err := g.DO.Find()
	return result.([]*model.GatewayAdmin), err
}

func (g gatewayAdminDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GatewayAdmin, err error) {
	buf := make([]*model.GatewayAdmin, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g gatewayAdminDo) FindInBatches(result *[]*model.GatewayAdmin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g gatewayAdminDo) Attrs(attrs ...field.AssignExpr) *gatewayAdminDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g gatewayAdminDo) Assign(attrs ...field.AssignExpr) *gatewayAdminDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g gatewayAdminDo) Joins(fields ...field.RelationField) *gatewayAdminDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g gatewayAdminDo) Preload(fields ...field.RelationField) *gatewayAdminDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g gatewayAdminDo) FirstOrInit() (*model.GatewayAdmin, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GatewayAdmin), nil
	}
}

func (g gatewayAdminDo) FirstOrCreate() (*model.GatewayAdmin, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GatewayAdmin), nil
	}
}

func (g gatewayAdminDo) FindByPage(offset int, limit int) (result []*model.GatewayAdmin, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g gatewayAdminDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g *gatewayAdminDo) withDO(do gen.Dao) *gatewayAdminDo {
	g.DO = *do.(*gen.DO)
	return g
}
