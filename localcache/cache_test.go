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

	result, ok := s.localcache.Get(key)
	assert.True(s.T(), ok)
	assert.Equal(s.T(), value, result)
}

func (s *LocalCacheTestSuite) TestOverwrite() {
	key := "myKey"
	value1 := "myValue1"
	value2 := "myValue2"

	s.localcache.Set(key, value1)
	s.localcache.Set(key, value2)

	result, ok := s.localcache.Get(key)
	assert.True(s.T(), ok)
	assert.Equal(s.T(), value2, result)
}

func (s *LocalCacheTestSuite) TestExpiration() {
	key := "myKey"
	value := "myValue"

	s.localcache.Set(key, value)

	time.Sleep((EXPIRATION_TTL + 1) * time.Second)

	_, ok := s.localcache.Get(key)
	assert.False(s.T(), ok)
}

func TestLocalCache(t *testing.T) {
	suite.Run(t, new(LocalCacheTestSuite))
}
