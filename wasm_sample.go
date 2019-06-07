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
