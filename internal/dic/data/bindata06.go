package data

import(
	"os"
	"time"
)

var _dicIpaUnkDic = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x93\xcf\x6b\x13\x41\x14\xc7\xdf\x4b\xb6\x93\x38\x6e\x43\x58\x42\x09\x39\x78\x10\x85\xd2\x93\x14\xf4\x54\x7b\x28\xf5\x50\x69\x7b\x08\x8a\x3f\x8e\x29\xad\xb1\x88\x96\xb6\x50\xf1\x20\x22\x31\x3f\x88\xd9\x06\xb1\x52\xb5\x36\x20\x0a\x06\x7a\xe8\xc1\xfa\x63\xc1\x7f\x42\xd0\x7f\x40\x27\x1b\x3c\x7b\x10\x0f\x33\x32\x4c\xa4\x8f\xba\x85\x8d\x46\x94\xf0\x20\xdf\xdd\xf7\x3e\xef\x7d\x67\xdf\xf4\xab\x5a\x04\x23\xca\x05\x54\x77\x01\x8e\xab\x6a\x14\xb1\x6f\xea\xda\xe2\x42\x5e\x3f\xc0\x28\xb2\xc9\xd9\xb9\xe5\x89\x71\xb4\x00\x63\xd9\xcb\x97\xf2\x9d\xff\xec\xdc\xac\x16\x68\x01\x80\xc4\x6f\xca\x85\x41\x94\x3c\x63\xe2\xde\x11\x40\xc9\x93\x26\x6e\x2c\x6a\xe1\x98\x70\xd3\x5a\xa4\x4d\x5c\x19\xd0\x82\x9b\xc8\x6f\x01\x32\x64\x28\x57\x18\x29\x9e\x4a\x91\xe2\x9c\x4b\x8a\xc7\xce\xc3\x6e\xc3\xdc\x02\x21\x8d\x8e\x00\x72\xe4\x28\x8f\x8e\x91\x94\x19\x87\x16\x7f\x27\xd8\x99\x61\xd2\x70\xfa\x22\x21\x9d\x5c\xd7\x82\x99\x78\x22\x09\xad\xfa\x92\xd0\x96\xb2\x84\x56\x4b\x13\xc0\x5c\xa6\x63\xaa\xe0\x91\x1e\x4b\x45\x82\x2d\x51\x87\xa5\xeb\x04\x9b\x7b\x4f\x6a\x56\x6e\x12\xec\x85\x71\x32\x4a\xe5\x81\x16\x87\x4c\x7c\x7c\x4a\x68\xf3\x3f\xbb\x5f\x5d\x26\xa4\xec\x30\xe9\x71\xf6\x0c\x21\xcd\x7b\xa4\xc7\xc4\x17\x82\xfd\xf0\x15\xd0\x41\x07\xe5\x29\x97\x8c\x5e\xc9\x41\x42\x95\x2c\x8c\xa8\x32\xa0\xa5\x57\x21\xa3\xca\x70\x10\x06\x9c\xb4\x7d\xd8\x3a\x16\x3f\xcd\xa6\x39\xa4\xec\xe4\x48\x62\x34\x32\x99\x54\x65\x60\xb6\x6d\xf1\x04\xe7\x3c\x15\x77\xe2\xfd\x6a\x55\x6f\x5e\x1d\x50\x15\x01\x6c\x75\x47\xab\x22\xa0\x0d\x20\xfb\x4e\xa8\x3a\x0c\xc6\x98\xa8\xd7\xda\xcd\x4d\x5b\x3c\xf6\x5a\x1b\x15\x23\x98\xd8\xd8\x16\x8d\x06\xfb\xfc\xe6\x56\xbb\xbc\x85\x43\xfa\x17\x9c\x48\x33\xf6\xcd\xf1\x3c\x51\xaf\x85\x80\xf9\x3b\x05\xff\xed\xfd\xbd\xb0\x3d\x3d\x70\x28\x76\xa0\x55\x68\x88\xea\x5a\xbb\xb9\xb9\xfb\xf0\x1f\x0e\x18\xfa\xe8\x7e\x75\xc2\xda\x2f\xd6\xc5\xea\xeb\x80\x17\x7f\xf4\x55\xc2\xcc\x19\xda\x74\xf0\x09\xee\x3f\xf2\xa7\xdb\x3b\xe2\x59\xa5\xe5\x3e\xf7\x5f\x3d\xfa\x5f\x1c\xfd\x95\x1d\xea\xc6\x68\x4f\x57\xac\xdb\x2f\x12\xf6\xc8\xa3\xad\xb5\xed\xdf\x31\xd2\xeb\xfb\xd8\xfb\x6b\x16\xec\xaf\x73\xf9\xfc\xa6\xe7\x3f\x7c\x17\x7a\x93\x7f\x04\x00\x00\xff\xff\x43\x34\x59\xde\xce\x07\x00\x00"

func dicIpaUnkDicBytes() ([]byte, error) {
	return bindataRead(
		_dicIpaUnkDic,
		"dic/ipa/unk.dic",
	)
}

func dicIpaUnkDic() (*asset, error) {
	bytes, err := dicIpaUnkDicBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dic/ipa/unk.dic", size: 1998, mode: os.FileMode(420), modTime: time.Unix(1559270655, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

