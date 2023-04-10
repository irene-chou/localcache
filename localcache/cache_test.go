package localcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LocalCacheTestSuite struct {
	suite.Suite
	localcache *localCache
}

func (s *LocalCacheTestSuite) SetupTest() {
	s.localcache = New().(*localCache)
}

func (s *LocalCacheTestSuite) TestSetAndGet() {
	key := "myKey"
	value := "myValue"
	s.localcache.Set(key, value)

	s.assertCacheValue(key, value)
}

func (s *LocalCacheTestSuite) TestOverwrite() {
	key := "myKey"
	value1 := "myValue1"
	value2 := "myValue2"

	s.localcache.Set(key, value1)
	s.localcache.Set(key, value2)

	s.assertCacheValue(key, value2)
}

func (s *LocalCacheTestSuite) TestExpiration() {
	key := "myKey"
	value := "myValue"

	s.localcache.Set(key, value)

	time.Sleep((EXPIRATION_TTL + 1) * time.Second)

	s.assertCacheValueNotExists(key)
}

func (s *LocalCacheTestSuite) TestOverwriteAndExpiration() {
	key := "myKey"
	value1 := "myValue1"
	value2 := "myValue2"

	s.localcache.Set(key, value1)

	time.Sleep((EXPIRATION_TTL / 2) * time.Second)

	s.localcache.Set(key, value2)

	time.Sleep(((EXPIRATION_TTL / 2) + 1) * time.Second)

	s.assertCacheValue(key, value2)

	time.Sleep((EXPIRATION_TTL / 2) * time.Second)

	s.assertCacheValueNotExists(key)
}

func TestLocalCache(t *testing.T) {
	suite.Run(t, new(LocalCacheTestSuite))
}

func (s *LocalCacheTestSuite) assertCacheValue(key, expectedValue string) {
	value, ok := s.localcache.Get(key)
	assert.True(s.T(), ok)
	assert.Equal(s.T(), expectedValue, value)
}

func (s *LocalCacheTestSuite) assertCacheValueNotExists(key string) {
	_, ok := s.localcache.Get(key)
	assert.False(s.T(), ok)
}
