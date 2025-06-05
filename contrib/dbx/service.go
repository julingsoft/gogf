package dbx

import (
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
	Service[T, D any] struct {
		model *gdb.Model
	}
)

// IService 是一个通用的泛型接口，用于获取单个实体
// 使用 [T any] 来定义一个泛型类型参数 T
type IService[T, D any] interface {
	Model() *gdb.Model
	Create(do *D) error
	CreateBatch(dos []*D) error
	InsertAndGetId(do *D) (int64, error)
	DeleteById(id any) error
	DeleteByIds(ids []any) error
	UpdateById(do *D, id any) error
	GetById(id any) (*T, error)
	GetOne(where interface{}, args ...interface{}) (*T, error)
	ListByIds(ids []any) ([]*T, error)
	List(where interface{}, args ...interface{}) ([]*T, error)
	Count(where interface{}, args ...interface{}) (int, error)
	Page(req PageRequest, where interface{}, args ...interface{}) (*PageResult[T], error)
}

// New 创建并返回一个 Service 的泛型实例
func New[T, D any](model *gdb.Model) *Service[T, D] {
	return &Service[T, D]{
		model: model,
	}
}

func (s *Service[T, D]) Model() *gdb.Model {
	return s.model
}

func (s *Service[T, D]) Create(do *D) error {
	if _, err := s.Model().Insert(do); err != nil {
		return err
	}
	return nil
}

func (s *Service[T, D]) CreateBatch(dos []*D) error {
	if len(dos) == 0 {
		return nil
	}
	if _, err := s.Model().Insert(dos); err != nil {
		return err
	}
	return nil
}

func (s *Service[T, D]) InsertAndGetId(do *D) (int64, error) {
	return s.Model().InsertAndGetId(do)
}

func (s *Service[T, D]) DeleteById(id any) error {
	if _, err := s.Model().WherePri(id).Delete(); err != nil {
		return err
	}
	return nil
}

func (s *Service[T, D]) DeleteByIds(ids []any) error {
	if len(ids) == 0 {
		return nil
	}
	if _, err := s.Model().WherePri(ids).Delete(); err != nil {
		return err
	}
	return nil
}

func (s *Service[T, D]) UpdateById(do *D, id any) error {
	if _, err := s.Model().WherePri(id).Update(do); err != nil {
		return err
	}
	return nil
}

func (s *Service[T, D]) GetById(id any) (entity *T, err error) {
	err = s.Model().WherePri(id).Scan(&entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (s *Service[T, D]) GetOne(where interface{}, args ...interface{}) (entity *T, err error) {
	err = s.Model().Where(where, args...).Scan(&entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (s *Service[T, D]) ListByIds(ids []any) (entities []*T, err error) {
	if len(ids) == 0 {
		return entities, nil
	}
	err = s.Model().WherePri(ids).Scan(&entities)
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (s *Service[T, D]) List(where interface{}, args ...interface{}) (entities []*T, err error) {
	err = s.Model().Where(where, args...).Scan(&entities)
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (s *Service[T, D]) Count(where interface{}, args ...interface{}) (int, error) {
	return s.Model().Where(where, args...).Count()
}

func (s *Service[T, D]) Page(req PageRequest, where interface{}, args ...interface{}) (result *PageResult[T], err error) {
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

	total, err = s.Count(where, args...)
	if err != nil {
		return nil, err
	}

	if total == 0 {
		return &PageResult[T]{Records: records, PageNum: req.PageNum, PageSize: req.PageSize, Total: 0}, nil
	}

	offset := (req.PageNum - 1) * req.PageSize
	err = s.Model().Where(where, args...).Offset(offset).Limit(req.PageSize).Scan(&records)
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
