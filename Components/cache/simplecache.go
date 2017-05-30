/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cache

import (
	"errors"
	"fmt"
	"sync"
)

type SimpleStore interface {
	Add(obj interface{}) error
	Update(obj interface{}) error
	Delete(obj interface{}) error
	Get(key string) (interface{}, bool)
	GetByKey(key string) (interface{}, bool)
	ListKeys() []string
	Clean() error
}

type KeyFunc func(obj interface{}) (string, error)

// simplestireMap implements SimpleStore
type simplestoreMap struct {
	lock    sync.RWMutex
	items   map[string]interface{}
	keyFunc KeyFunc
}

func NewSimpleStoreMap(keyFunc KeyFunc) *simplestoreMap {
	return &simplestoreMap{
		items:   map[string]interface{}{},
		keyFunc: keyFunc,
	}
}

func (c *simplestoreMap) Add(obj interface{}) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	key, err := c.keyFunc(obj)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create key from obj %+v", obj))
	}
	c.items[key] = obj
	return nil
}

// Update sets an item in the cache to its updated state.
func (c *simplestoreMap) Update(obj interface{}) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	key, err := c.keyFunc(obj)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create key from obj %+v", obj))
	}
	c.items[key] = obj
	return nil
}

func (c *simplestoreMap) Get(key string) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	item, exists := c.items[key]
	return item, exists
}

func (c *simplestoreMap) Delete(obj interface{}) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	key, err := c.keyFunc(obj)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create key from obj %+v", obj))
	}
	delete(c.items, key)
	return nil
}

func (c *simplestoreMap) GetByKey(key string) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	item, exists := c.items[key]
	return item, exists
}

func (c *simplestoreMap) ListKeys() []string {
	c.lock.Lock()
	defer c.lock.Unlock()
	list := make([]string, 0, len(c.items))
	for key := range c.items {
		list = append(list, key)
	}
	return list
}

func (c *simplestoreMap) Clean() error {
	c.lock.Lock()
	defer c.lock.Unlock()
	for key := range c.items {
		delete(c.items, key)
	}
	return nil
}
