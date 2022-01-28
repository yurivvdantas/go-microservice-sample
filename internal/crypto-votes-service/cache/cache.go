package cache

import "go-microservice-sample/internal/crypto-votes-service/model"

var cache map[string]*model.Cryptos

func GetCache(name string) *model.Cryptos {
	Init()
	return cache[name]
}

func Init() {
	if cache != nil {
		return
	}
	cache = make(map[string]*model.Cryptos)
}

func SetCache(name string, crypto *model.Cryptos) {
	Init()
	cache[name] = crypto
}

func Clean(name string) {
	cache[name] = nil
}
