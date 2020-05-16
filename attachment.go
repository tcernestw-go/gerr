package gerr

import "fmt"

type attachments struct {
	mapping map[string]interface{}
}

// it will add attachment into attachments with the specific key
// it will replace origin attachment if key is already existed
func (attms *attachments) add(key string, attachment interface{}) {
	if attms.mapping == nil {
		attms.mapping = map[string]interface{}{}
	}
	attms.mapping[key] = make([]interface{}, 0)
	attms.mapping[key] = attachment
}

func (attms attachments) get(key string) (attachment interface{}) {
	return attms.mapping[key]
}

func (attms attachments) log(padding string) (msg string) {
	for key, attm := range attms.mapping {
		msg += fmt.Sprintf("%v%v: %v\n", padding, key, attm)
	}
	return
}
