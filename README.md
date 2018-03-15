
# ContextTool

 ServletのApplicationScopeのようなもの。
自分用

## 使用例

　重たいファイルなどを読み込む時に１度だけに済ませる用
基本的なメタデータは永久に保持しておきたいのでこの形を取った。

```
type Hoge struct {
	HogeTitle string
}

// GetMeta サイトの情報を返却する。
func GetMeta() Hoge {
	data, flag := context.Get("BUZ")
	if !flag {
		file, err := ioutil.ReadFile(`bar.xml`)
		if err != nil {
			panic(err)
		}
		var fuga Hoge
		xml.Unmarshal(file, &fuga)
		context.Put("BUZ", fuga)
		return Hoge
	}
	return data.(Hoge)
}

```