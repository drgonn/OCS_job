package public

import (
	"ocs-app/model"
	"reflect"
	"strings"
)

func parseModel2Table[MODEL AllM](base model.BaseDAO, model MODEL, name, mean, index string) *Table[MODEL] {
	var args []Arg
	var sort bool
	var postTags, putTags, listTags []string

	mtype := reflect.TypeOf(model)
	for j := 0; j < mtype.NumField()-1; j++ {
		field := mtype.Field(j)
		postTag := field.Tag.Get("post")
		if len(postTag) > 0 {
			postTags = strings.Split(postTag, ",")
		} else {
			postTags = []string{}
		}
		putTag := field.Tag.Get("put")
		if len(putTag) > 0 {
			putTags = strings.Split(putTag, ",")
		} else {
			putTags = []string{}
		}
		listTag := field.Tag.Get("list")
		if len(listTag) > 0 {
			listTags = strings.Split(listTag, ",")
		} else {
			listTags = []string{}
		}

		arg := Arg{
			NotNull:  true,
			Post:     postTags,
			Put:      putTags,
			List:     listTags,
			Sort:     sort,
			Name:     field.Name,
			JsonName: strings.Split(field.Tag.Get("json"), ",")[0],
			Type:     field.Type.Name(),
			Mean:     "",
		}
		args = append(args, arg)
	}
	r := Table[MODEL]{
		Name:   name,
		Mean:   mean,
		Index:  index,
		Args:   args,
		Model:  model,
		Struct: base,
	}
	return &r
}
