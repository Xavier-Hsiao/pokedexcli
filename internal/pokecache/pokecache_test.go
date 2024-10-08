package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	const interval = 5 * time.Second
	testCache := CreateNewCache(interval)
	if testCache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := CreateNewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 10 * time.Millisecond
	const waitTime = baseTime + 10*time.Millisecond

	testCache := CreateNewCache(baseTime)
	testCache.Add("https://example.com", []byte("testdata"))

	_, ok := testCache.Get("https://example.com")
	if !ok {
		t.Errorf("failed to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = testCache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

func TestFailedReapLoop(t *testing.T) {
	const baseTime = 10 * time.Millisecond
	const waitTime = baseTime + 10*time.Millisecond

	testCache := CreateNewCache(baseTime)
	key := "https://example.com"
	testCache.Add(key, []byte("testdata"))

	_, ok := testCache.Get(key)
	if !ok {
		t.Errorf("failed to find key")
		return
	}

	time.Sleep(baseTime / 2)

	_, ok = testCache.Get(key)
	if !ok {
		t.Errorf("%s should not have been purged", key)
		return
	}
}
