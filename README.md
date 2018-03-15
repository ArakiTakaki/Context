
# ContextTool

 ServletのApplicationScopeのようなもの。
自分用

## 使用例

　重たいファイルなどを読み込む時に１度だけに済ませる用
基本的なメタデータは永久に保持しておきたいのでこの形を取った。

```

// GetMeta サイトの情報を返却する。
func GetMeta() Site {
	data, flag := context.Get("meta")
	if !flag {
		file, err := ioutil.ReadFile(`./conf/conf.xml`)
		if err != nil {
			panic(err)
		}
		var site Site
		xml.Unmarshal(file, &site)
		context.Put("meta", site)
		return site
	}
	return data.(Site)
}

```