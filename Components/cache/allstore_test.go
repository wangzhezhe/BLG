package cache

import (
	"log"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/wangzhezhe/BLG/Components/cache/util/sets"
)

func doFIFOSCacheStore(t *testing.T, store *FIFO) {
	mkObj := func(id string, val string) testStoreObject {
		return testStoreObject{id: id, val: val}
	}

	Popfunc := func(item interface{}) error {
		log.Printf("do populated function for item %+v\n", item)
		return nil
	}
	const amount = 10
	wait := sync.WaitGroup{}
	wait.Add(10)
	go func() {
		for i := 0; i < amount; i++ {
			store.Add(mkObj(string('a'+i), strconv.Itoa(i+1)))
			time.Sleep(time.Second)

		}
	}()

	go func() {
		for i := 0; i < amount; i++ {
			store.Pop(Popfunc)
			wait.Done()
			time.Sleep(time.Second)
		}

	}()

	time.Sleep(time.Second * 2)
	wait.Wait()

}

func doSimpleCacheStore(t *testing.T, store SimpleStore) {
	mkObj := func(id string, val string) testStoreObject {
		return testStoreObject{id: id, val: val}
	}

	//test add
	o := mkObj("add", "testadd")
	err := store.Add(o)
	if err != nil {
		t.Error(err)
	}
	value, ok := store.Get("add")
	if !ok {
		t.Error("failed to get key")
	}

	if value.(testStoreObject).val != "testadd" {
		t.Error("get error value")
	}

	//test delete
	store.Delete(mkObj("foo", ""))
	if _, ok := store.Get("foo"); ok {
		t.Errorf("found deleted item??")
	}
	//test update
	store.Update(mkObj("foo", "baz"))
	if item, ok := store.Get("foo"); !ok {
		t.Errorf("didn't find inserted item")
	} else {
		if e, a := "baz", item.(testStoreObject).val; e != a {
			t.Errorf("expected %v, got %v", e, a)
		}
	}

	//test list
	err = store.Clean()
	if err != nil {
		t.Error(err)
	}
	store.Update(mkObj("foo", "baz"))
	store.Update(mkObj("bar", "cdf"))
	found := sets.String{}
	for _, item := range store.ListKeys() {
		found.Insert(item)
	}
	if !found.HasAll("bar", "foo") {
		t.Errorf("missing items, found: %v", found)
	}
	if len(found) != 2 {
		t.Errorf("extra items", found)
	}

}

type testStoreObject struct {
	id  string
	val string
}

func testStoreKeyFunc(obj interface{}) (string, error) {
	return obj.(testStoreObject).id, nil
}

func TestSimpleCacheStore(t *testing.T) {
	doSimpleCacheStore(t, NewSimpleStoreMap(testStoreKeyFunc))
}

func TestLRUStore(t *testing.T) {
	//TODO add test here
}

func TestFIFOCacheStore(t *testing.T) {

	doFIFOSCacheStore(t, NewFIFO(testStoreKeyFunc))
}
