package device

import (
	"sync"
)

type ProductInfo struct {
	sync.RWMutex
	Products map[int]*Product
}

func NewProductInfo() (*ProductInfo, error) {
	products, err := FetchProducts()
	if err != nil {
		return nil, err
	}
	return &ProductInfo{
		Products: products,
	}, nil
}

func (pi *ProductInfo) GetProduct(pid int) *Product {
	pi.Lock()
	defer pi.Unlock()

	return pi.Products[pid]
}
