package context

// // Accesser test
// type Accesser interface {
// 	New()
// 	Put(key string, value interface{})
// 	Get(key string) (interface{}, bool)
// 	Len() int
// 	Remove(key string) bool
// }

type mapping struct {
	instance map[string]interface{}
}

var context = new()

// new 明示的にコンテキストの作成処理を行う
func new() *mapping {
	var work = mapping{}
	work.instance = make(map[string]interface{})
	return &work
}

// Put (key string, value mapping) 対応した構造体を返却する。
func Put(key string, value interface{}) {
	context.instance[key] = value
}

// Get 引数Key返り値mapping存在チェック
func Get(key string) (interface{}, bool) {
	work, swt := context.instance[key]
	return work, swt
}

// Len マップの長さを返却する
func Len() int {
	return len(context.instance)
}

// Remove 値の削除
func Remove(key string) bool {
	_, swt := context.instance[key]
	if swt {
		delete(context.instance, key)
	} else {
		return false
	}
	return true
}
