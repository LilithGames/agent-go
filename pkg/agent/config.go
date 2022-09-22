package agent

import (
	"sync"

	"github.com/spf13/viper"
)

type variable struct {
	sync.Mutex
	viper *viper.Viper
}

func initVariable(viper *viper.Viper) *variable {
	return &variable{viper: viper}
}

func (v *variable) GetString(key string) string {
	v.Lock()
	defer v.Unlock()
	return v.viper.GetString(key)
}

func (v *variable) Set(key string, val any) {
	v.Lock()
	defer v.Unlock()
	v.viper.Set(key, val)
}
