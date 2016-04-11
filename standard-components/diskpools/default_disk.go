package diskpools

import "github.com/xchapter7x/enaml"

func NewDiskPool(name string, size int) enaml.DiskPool {
	return enaml.DiskPool{
		Name:     name,
		DiskSize: size,
	}
}
