package memcached

import (
	"errors"
	. "farm.e-pedion.com/repo/cache"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	cacheClient    = new(MemcachedClient)
	memcacheClient *GoMemcacheClientMock
	setted         = false
	key1           = "8b06603b-9b0d-4e8c-8aae-10f988639fe6"
	expires        = 60
	testConfig     *Configuration
)

func init() {
	testConfig = &Configuration{
		Provider: "memcached",
		URL:      "mock://cache",
	}
}

func setup() error {
	setupErr := Setup(testConfig)
	if setupErr != nil {
		setted = true
	}
	return setupErr
}

func before() error {
	if !setted {
		if err := setup(); err != nil {
			return err
		}
	}
	memcacheClient = NewGoMemcacheClientMock()
	cacheClient.client = memcacheClient
	return nil
}

func TestNewClientPanic(t *testing.T) {
	assert.Panics(t,
		func() {
			NewClient()
		},
	)
}

func TestNewClient(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	assert.NotPanics(t,
		func() {
			NewClient()
		},
	)
}

func TestAddItem(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	assert.NotNil(t, cacheClient)
	memcacheClient.On("Add", mock.Anything).Return(nil)
	err := cacheClient.Add(key1, expires, []byte("1234567890"))
	assert.Nil(t, err)
}

func TestGetItem(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	assert.NotNil(t, cacheClient)
	memcacheClient.On("Get", mock.Anything).Return(&memcache.Item{Value: []byte("CacheMock")}, nil)
	item, err := cacheClient.Get(key1)
	assert.Nil(t, err)
	assert.NotNil(t, item)
}

func TestDelItem(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	assert.NotNil(t, cacheClient)
	memcacheClient.On("Delete", mock.Anything).Return(nil)
	err := cacheClient.Delete(key1)
	assert.Nil(t, err)
}

func TestGetEmptyItem(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	assert.NotNil(t, cacheClient)
	memcacheClient.On("Get", mock.Anything).Return(nil, errors.New("ErrMockGet"))
	item, err := cacheClient.Get(key1)
	assert.NotNil(t, err)
	assert.Nil(t, item)
}

func TestSetItem(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	assert.NotNil(t, cacheClient)
	memcacheClient.On("Set", mock.Anything).Return(nil)
	err := cacheClient.Set("cache_test", 120, []byte("golang test"))
	assert.Nil(t, err)
}
