package context

var instance = make(map[string]interface{})

// Put (key string, value mapping) 対応した構造体を返却する。
func Put(key string, value interface{}) {
	instance[key] = value
}

// Get 引数Key返り値mapping存在チェック
func Get(key string) (interface{}, bool) {
	work, swt := instance[key]
	return work, swt
}

// Len マップの長さを返却する
func Len() int {
	return len(instance)
}

// Remove 値の削除
func Remove(key string) bool {
	_, swt := instance[key]
	if swt {
		delete(instance, key)
	} else {
		return false
	}
	return true
}
