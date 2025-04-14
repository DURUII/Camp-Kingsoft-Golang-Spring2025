package pool

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

// 例如数据库连接、网络连接
type ReusableObj struct {
}

type ObjPool struct {
	bufChan chan *ReusableObj
}

func NewObjPool(numOfObj int) *ObjPool {
	objPool := new(ObjPool)
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		// 初始化
		objPool.bufChan <- &ReusableObj{}
	}
	return objPool
}

func (pool *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case obj := <-pool.bufChan:
		return obj, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}

func (pool *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case pool.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
}
