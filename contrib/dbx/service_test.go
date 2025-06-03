package dbx

import (
	"business-trip-switch/internal/model/entity"
	"context"
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func Test_Service(t *testing.T) {
	var o = NewOrder()
	order, err := o.GetById(context.TODO(), 1)
	if err != nil {
		_ = fmt.Errorf("query err : %v", err)
	}
	fmt.Println(order)
}

func NewOrder() IService[entity.Order] {
	return New[entity.Order](g.Model("order"))
}
