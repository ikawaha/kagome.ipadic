# build dictionary

install: ipa_internal bindata ipa
	cp ./bindata*.go ../../internal/dic/data
	cp ipa.dic ../../_sample

.PHONY: ipa_internal
ipa_internal:
	go run main.go ipa -mecab ./mecab-ipadic-2.7.0-20070801 -output ./dic/ipa/ -z=false

.PHONY: ipa
ipa:
	go run main.go ipa -mecab ./mecab-ipadic-2.7.0-20070801

.PHONY: bindata
bindata: 
	go-bindata -o bindata.go -nomemcopy -separate -pkg=data dic/...

.PHONY: clean
clean:
	rm ./bindata*.go
	rm ipa.dic

