package message

import "fmt"

func Main() {
	Load([]MsgBundleInfo{
		{"./message/message_ko.json", "ko"},
		{"./message/message_en.json", "en"},
	})

	lang := "ko"
	msg := Bundle(lang)

	fmt.Printf("%s", msg["global_input_box"])
}
