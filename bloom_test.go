package bloom

import (
	"testing"

	"github.com/go-redis/redis/v8"

	"github.com/stretchr/testify/assert"
)

var REDIS_URL = "redis://:@localhost:6479/0"

func TestRedisBitSet_New_Set_Test(t *testing.T) {
	// store, clean, err := redistest.CreateRedis()
	opt, err := redis.ParseURL(REDIS_URL)

	if err != nil {
		panic(err)
	}

	store := redis.NewClient(opt)

	bitSet := newRedisBitSet(store, "test_key", 1024)
	isSetBefore, err := bitSet.check([]uint{0})
	if err != nil {
		t.Fatal(err)
	}
	if isSetBefore {
		t.Fatal("Bit should not be set")
	}
	err = bitSet.set([]uint{512})
	if err != nil {
		t.Fatal(err)
	}
	isSetAfter, err := bitSet.check([]uint{512})
	if err != nil {
		t.Fatal(err)
	}
	if !isSetAfter {
		t.Fatal("Bit should be set")
	}
	err = bitSet.expire(3600)
	if err != nil {
		t.Fatal(err)
	}
	err = bitSet.del()
	if err != nil {
		t.Fatal(err)
	}
}

func TestRedisBitSet_Add(t *testing.T) {
	opt, err := redis.ParseURL(REDIS_URL)

	if err != nil {
		panic(err)
	}

	store := redis.NewClient(opt)
	filter := New(store, "test_key", 64)
	assert.Nil(t, filter.Add([]byte("hello")))
	assert.Nil(t, filter.Add([]byte("world")))
	ok, err := filter.Exists([]byte("hello"))
	assert.Nil(t, err)
	assert.True(t, ok)
}
