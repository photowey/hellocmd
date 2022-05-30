package configregistry

import (
	"log"
	"sync"
)

var (
	_lock     sync.Mutex
	_registry = newRegistry()

	_ Registry = (*registry)(nil)
)

// Observer 配置对象 Config 的观察者
type Observer interface {
	// Name 获取观察者的名称
	Name() string
	// Update 执行更新
	Update()
}

// Registry 配置对象 Config 的观察者注册中心
type Registry interface {
	// Register 注册 配置 Config对象 的 Observer
	//
	// @param name Observer 名称
	//
	// @param notifier Observer 实例
	//
	// @dangerous maybe trigger the panic ⭐⭐⭐⭐⭐
	Register(name string, observer Observer)
	// Deregister 注销观察者
	//
	// @param name Observer 名称
	Deregister(name string)
	// Has 是否包含某个观察者
	//
	// @param name Observer 名称
	Has(name string) bool
	// NotifyAll 通知所有的观察者 Observer
	NotifyAll()
}

type registry struct {
	ctx map[string]Observer
}

func newRegistry() *registry {
	return &registry{
		ctx: make(map[string]Observer),
	}
}

// ---------------------------------------------------------------- method

func (r registry) Register(name string, observer Observer) {
	_lock.Lock()
	defer _lock.Unlock()
	if observer == nil {
		panic("config.registry: Register observer is nil")
	}

	if _, dup := r.ctx[name]; dup {
		panic("config.registry: Register called twice for observer: " + name)
	}

	r.ctx[name] = observer
	log.Printf("config.registry: register observer: [%s]", name)
}

func (r registry) Deregister(name string) {
	_lock.Lock()
	defer _lock.Unlock()
	if ok := r.Has(name); ok {
		delete(r.ctx, name)
	}
}

func (r registry) Has(name string) bool {
	_, ok := r.ctx[name]

	return ok
}

func (r registry) NotifyAll() {
	for _, observer := range r.ctx {
		observer.Update()
	}
}

// ---------------------------------------------------------------- func

// Register 注册 配置 Config对象 的 Observer
//
// @param name Observer 名称
//
// @param notifier Observer 实例
//
// @dangerous maybe trigger the panic ⭐⭐⭐⭐⭐
func Register(name string, observer Observer) {
	_registry.Register(name, observer)
}

// Deregister 注销观察者
//
// @param name Observer 名称
func Deregister(name string) {
	_registry.Deregister(name)
}

// Has 是否包含某个观察者
//
// @param name Observer 名称
func Has(name string) bool {
	return _registry.Has(name)
}

// Broadcast 广播所有的观察者
func Broadcast() {
	_registry.NotifyAll()
}
