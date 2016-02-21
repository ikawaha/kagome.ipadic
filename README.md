Kagome Japanese Morphological Analyzer (IPADic only)
===

kagome.ipadic is a small version of [kagome](https://github.com/ikawaha/kagome).
This package supports the IPADic only.


# Programming example

Below is a simple go example that demonstrates how a simple text can be segmented.

See aloso https://github.com/ikawaha/kagome 


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

License
---
Kagome is licensed under the Apache License v2.0 and uses the MeCab-IPADIC model. See NOTICE.txt for license details.
