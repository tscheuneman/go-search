package implementor

import (
	"fmt"
	"time"

	lru "github.com/hashicorp/golang-lru"
)

type CacheClient struct {
	Client *lru.Cache
	Ttl    int64
}

type CacheItem struct {
	Ttl   int64
	Value interface{}
}

type CallbackFunction func() (interface{}, error)

func (x CacheClient) Resolve(key string, cbFunc CallbackFunction) (interface{}, error) {

	value, found := x.GetCacheItem(key)
	if found {
		fmt.Println("Fetching data from cache")
		return value, nil
	}

	fmt.Println("Fetching data from callback")
	value, err := cbFunc()

	if err != nil {
		return nil, err
	}

	x.UpdateSetCacheItem(key, value)

	return value, nil
}

func (x CacheClient) GetCacheItem(key string) (returnVal interface{}, found bool) {
	value, found := x.Client.Get(key)

	if !found {
		return nil, false
	}

	cacheItem := value.(*CacheItem)

	var currentTime int64 = time.Now().Unix()

	if (currentTime + x.Ttl) > cacheItem.Ttl {
		return cacheItem.Value, true
	}

	_ = x.Client.Remove(key)

	return nil, false
}

func (x CacheClient) UpdateSetCacheItem(key string, value interface{}) {
	x.Client.Add(key, &CacheItem{
		Value: value,
		Ttl:   time.Now().Unix() + x.Ttl,
	})
}

func (x CacheClient) RemoveCacheItem(key string) {
	x.Client.Remove(key)
}
