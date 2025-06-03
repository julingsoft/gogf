package dbx

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
)

type (
	PageRequest struct {
		PageNum  int
		PageSize int
	}

	PageResult[T any] struct {
		Records  []*T `json:"records"`
		PageNum  int  `json:"pageNum"`
		PageSize int  `json:"pageSize"`
		Total    int  `json:"total"`
	}

	// Service 是 IService 接口的一个泛型实现
	Service[T any] struct {
		model *gdb.Model
	}
)

// IService 是一个通用的泛型接口，用于获取单个实体
// 使用 [T any] 来定义一个泛型类型参数 T
type IService[T any] interface {
	Model(ctx context.Context) *gdb.Model
	Create(ctx context.Context, entity *T) error
	CreateBatch(ctx context.Context, entities []*T) error
	InsertAndGetId(ctx context.Context, entity *T) (int64, error)
	DeleteById(ctx context.Context, id any) error
	DeleteByIds(ctx context.Context, ids []any) error
	UpdateById(ctx context.Context, entity *T, id any) error
	GetById(ctx context.Context, id any) (*T, error)
	GetOne(ctx context.Context, where interface{}, args ...interface{}) (*T, error)
	ListByIds(ctx context.Context, ids []any) ([]*T, error)
	List(ctx context.Context, where interface{}, args ...interface{}) ([]*T, error)
	Count(ctx context.Context, where interface{}, args ...interface{}) (int, error)
	Page(ctx context.Context, req PageRequest, where interface{}, args ...interface{}) (*PageResult[T], error)
}

// New 创建并返回一个 Service 的泛型实例
func New[T any](m *gdb.Model) *Service[T] {
	return &Service[T]{
		model: m,
	}
}

func (s *Service[T]) Model(ctx context.Context) *gdb.Model {
	return s.model.Ctx(ctx)
}

func (s *Service[T]) Create(ctx context.Context, entity *T) error {
	if _, err := s.Model(ctx).Insert(entity); err != nil {
		return err
	}
	return nil
}

func (s *Service[T]) CreateBatch(ctx context.Context, entities []*T) error {
	if len(entities) == 0 {
		return nil
	}
	if _, err := s.Model(ctx).Insert(entities); err != nil {
		return err
	}
	return nil
}

func (s *Service[T]) InsertAndGetId(ctx context.Context, entity *T) (int64, error) {
	return s.Model(ctx).InsertAndGetId(entity)
}

func (s *Service[T]) DeleteById(ctx context.Context, id any) error {
	if _, err := s.Model(ctx).WherePri(id).Delete(); err != nil {
		return err
	}
	return nil
}

func (s *Service[T]) DeleteByIds(ctx context.Context, ids []any) error {
	if len(ids) == 0 {
		return nil
	}
	if _, err := s.Model(ctx).WherePri(ids).Delete(); err != nil {
		return err
	}
	return nil
}

func (s *Service[T]) UpdateById(ctx context.Context, entity *T, id any) error {
	if _, err := s.Model(ctx).WherePri(id).Update(entity); err != nil {
		return err
	}
	return nil
}

func (s *Service[T]) GetById(ctx context.Context, id any) (entity *T, err error) {
	err = s.Model(ctx).WherePri(id).Scan(&entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (s *Service[T]) GetOne(ctx context.Context, where interface{}, args ...interface{}) (entity *T, err error) {
	err = s.Model(ctx).Where(where, args).Scan(&entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (s *Service[T]) ListByIds(ctx context.Context, ids []any) (entities []*T, err error) {
	if len(ids) == 0 {
		return entities, nil
	}
	err = s.Model(ctx).WherePri(ids).Scan(&entities)
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (s *Service[T]) List(ctx context.Context, where interface{}, args ...interface{}) (entities []*T, err error) {
	err = s.Model(ctx).Where(where, args).Scan(&entities)
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (s *Service[T]) Count(ctx context.Context, where interface{}, args ...interface{}) (int, error) {
	return s.Model(ctx).Count(where, args)
}

func (s *Service[T]) Page(ctx context.Context, req PageRequest, where interface{}, args ...interface{}) (result *PageResult[T], err error) {
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	if req.PageSize > 200 {
		req.PageSize = 200
	}

	var (
		total   int
		records []*T
	)

	total, err = s.Count(ctx, where, args)
	if err != nil {
		return nil, err
	}

	if total == 0 {
		return &PageResult[T]{Records: records, PageNum: req.PageNum, PageSize: req.PageSize, Total: 0}, nil
	}

	offset := (req.PageNum - 1) * req.PageSize
	err = s.Model(ctx).Where(where, args).Offset(offset).Limit(req.PageSize).Scan(&records)
	if err != nil {
		return nil, err
	}

	return &PageResult[T]{
		Records:  records,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Total:    total,
	}, nil
}
