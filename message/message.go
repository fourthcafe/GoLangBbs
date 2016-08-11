package message

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type MsgBundleInfo struct {
	FilePath string
	Lang     string
}

type MsgBundle map[string]string
type MsgBundleList map[string]MsgBundle

var cachedMsg = make(MsgBundleList)

func Load(msgBundles []MsgBundleInfo) {
	for _, info := range msgBundles {
		if file, err := ioutil.ReadFile(fmt.Sprintf(info.FilePath)); err == nil {
			msgBunble := MsgBundle{}
			json.Unmarshal(file, &msgBunble)

			cachedMsg[info.Lang] = msgBunble

		} else {
			log.Println(err)
		}
	}
}

func Bundle(lang string) MsgBundle {
	return cachedMsg[lang]
}
