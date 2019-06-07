[![Build Status](https://travis-ci.org/ikawaha/kagome.ipadic.svg?branch=master)](https://travis-ci.org/ikawaha/kagome.ipadic) [![BuildStatus(Windows)](https://ci.appveyor.com/api/projects/status/1y9pvcrdyarpmsy1/branch/master?svg=true)](https://ci.appveyor.com/project/ikawaha/kagome-ipadic) [![Coverage Status](https://coveralls.io/repos/ikawaha/kagome.ipadic/badge.svg?branch=master&service=github)](https://coveralls.io/github/ikawaha/kagome.ipadic?branch=master)  [![GoDoc](https://godoc.org/github.com/ikawaha/kagome.ipadic?status.svg)](https://godoc.org/github.com/ikawaha/kagome.ipadic)

Kagome Japanese Morphological Analyzer (IPADic only)
===

kagome.ipadic is a small version of [kagome](https://github.com/ikawaha/kagome).
This package supports the IPADic only.


# Programming example

Below is a simple go example that demonstrates how a simple text can be segmented.

See also https://github.com/ikawaha/kagome 


sample code:

```go:example
package main

import (
	"fmt"
	"strings"

	"github.com/ikawaha/kagome.ipadic/tokenizer"
)

func main() {
	t := tokenizer.New()
	tokens := t.Tokenize("寿司が食べたい。") // t.Analyze("寿司が食べたい。", tokenizer.Normal)
	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			// BOS: Begin Of Sentence, EOS: End Of Sentence.
			fmt.Printf("%s\n", token.Surface)
			continue
		}
		features := strings.Join(token.Features(), ",")
		fmt.Printf("%s\t%v\n", token.Surface, features)
	}
}
```

output:

```text:outputs
BOS
寿司    名詞,一般,*,*,*,*,寿司,スシ,スシ
が      助詞,格助詞,一般,*,*,*,が,ガ,ガ
食べ    動詞,自立,*,*,一段,連用形,食べる,タベ,タベ
たい    助動詞,*,*,*,特殊・タイ,基本形,たい,タイ,タイ
。      記号,句点,*,*,*,*,。,。,。
EOS
```

# Working with GAE/Go

Using fully kagome.ipadic on GAE/Go needs at least B4 instance (>512MB memory).
If you use a simple dictionary without contents other than part of speech, it will be able to run on B1 instance.
Even in that case, the analysis result does not change, the output contents (`活用型`, `活用形`, `基本形`, `読み`, `発音`) are omitted.

|Instance Class|Memory Limit|CPU Limit|
|:---|:---|:---|
|B1|128 MB|600 Mhz|
|B2|256 MB|1.2 Ghz|
|B4|512 MB|2.4 Ghz|
|B4_1G|1024 MB|2.4 Ghz|
|B8|1024 MB|4.8 Ghz|
|F1|128 MB|600 Mhz|
|F2|256 MB|1.2 Ghz|
|F4|512 MB|2.4 Ghz|
|F4_1G|1024 MB|2.4 Ghz|

## Usage

```
command:
    use `-sysdic=simple` option. ex, kagome -sysdic=simple
lib:
    use `dic := tokenizer.SysDicIPASimple()` instead of `dic := tokenizer.SysDic()`
```

### Full Dict.

```text:outputs
BOS
寿司    名詞,一般,*,*,*,*,寿司,スシ,スシ
が      助詞,格助詞,一般,*,*,*,が,ガ,ガ
食べ    動詞,自立,*,*,一段,連用形,食べる,タベ,タベ
たい    助動詞,*,*,*,特殊・タイ,基本形,たい,タイ,タイ
。      記号,句点,*,*,*,*,。,。,。
EOS
```

### Simple Dict.

```text:outputs
BOS
寿司    名詞,一般,*,*,*,*
が      助詞,格助詞,一般,*,*,*
食べ    動詞,自立,*,*,一段,連用形
たい    助動詞,*,*,*,特殊・タイ,基本形
。      記号,句点,*,*,*,*
EOS
```

# WebAssembly

You can see how kagome wasm works in [demo site.](http://ikawaha.github.io/kagome.ipadic/demo/)

## Sample main.go

```Go
package main

import (
	"syscall/js"

	"github.com/ikawaha/kagome.ipadic/tokenizer"
)

func tokenize(_ js.Value, args []js.Value) interface{} {
	t := tokenizer.New()
	if len(args) == 0 {
		return nil
	}
	ret := []interface{}{}
	tokens := t.Tokenize(args[0].String())
	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			//fmt.Printf("%s\n", token.Surface)
			continue
		}
		features := token.Features()
		for i := 9 - len(features); i > 0; i-- {
			features = append(features, "*")
		}
		//fmt.Printf("%s\t%v\n", token.Surface, strings.Join(features, ","))
		ret = append(ret, map[string]interface{}{
			"word_id":         token.ID,
			"word_type":       token.Class.String(),
			"word_position":   token.Start,
			"surface_form":    token.Surface,
			"pos":             features[0],
			"pos_detail_1":    features[1],
			"pos_detail_2":    features[2],
			"pos_detail_3":    features[3],
			"conjugated_type": features[4],
			"conjugated_form": features[5],
			"basic_form":      features[6],
			"reading":         features[7],
			"pronunciation":   features[8],
		})
	}
	return ret
}

var global = js.Global()

func main() {
	_ = tokenizer.New()
	c := make(chan struct{}, 0)
	println("Go Web Assembly Ready")

	global.Set("kagome", js.FuncOf(tokenize))
	<-c
}
```

## Build wasm

```shellsession
$ GOOS=js GOARCH=wasm go build -o kagome.wasm ./main.go
```

License
---
Kagome is licensed under the Apache License v2.0 and uses the MeCab-IPADIC model. See NOTICE.txt for license details.
