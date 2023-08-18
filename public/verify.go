package public

import (
	"errors"
	"fmt"
	"ocs-app/model"
	"ocs-app/util"
	"reflect"
	"strings"
)

func (t *Table[MODEL]) VerifyPostArg(entity MODEL) error {
	// 检验必填参数是否有提交
	for _, arg := range t.Args {
		if len(arg.Post) > 0 {
			fmt.Println("ar  ", arg)
			err := t.distributeBindArg(arg.Post, entity, arg, "0")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *Table[MODEL]) VerifyPatchArg(entity MODEL, id string) error {
	for _, arg := range t.Args {
		if len(arg.Put) > 0 {
			fmt.Println("patch  ", arg.Put)
			err := t.distributeBindArg(arg.Put, entity, arg, id)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *Table[MODEL]) distributeBindArg(binds []string, entity MODEL, arg Arg, id string) error {
	for _, bind := range binds {
		switch {
		case bind == "required":
			err := t.handle_required(entity, arg)
			if err != nil {
				return err
			}
		case bind == "ban":
			err := t.handle_ban(entity, arg)
			if err != nil {
				return err
			}
		case strings.HasPrefix(bind, "unique="):
			// 处理 bind 以 "unique=" 开头的情况
			err := t.handle_unique(entity, bind[7:], arg, id)
			if err != nil {
				return err
			}
		default:
			return errors.New("unknown bind: " + bind)
		}
	}
	return nil
}

func (t *Table[MODEL]) handle_required(entity MODEL, arg Arg) error {
	fieldValue := reflect.ValueOf(entity).FieldByName(arg.Name)
	fmt.Println("fieldValue", fieldValue)
	if fieldValue.IsValid() {
		fv := fieldValue.Interface()
		if fv == "" || fv == nil || fv == 0 {
			return errors.New(arg.JsonName + " required")
		}
		return nil
	} else {
		return errors.New(arg.JsonName + " required")
	}
}

func (t *Table[MODEL]) handle_unique(entity MODEL, restrain string, arg Arg, id string) error {
	if restrain == "" {
		return nil
	}
	var count int64
	keys := map[string]interface{}{}
	keyOpts := map[string]interface{}{}
	value := reflect.ValueOf(entity).FieldByName(arg.Name)
	cons_name := util.UnderscoreToCamelCase(restrain)
	cons_value := reflect.ValueOf(entity).FieldByName(cons_name)

	dao := model.DAOOption{
		Where: map[string]interface{}{
			restrain:     cons_value.Interface(),
			arg.JsonName: value.Interface(),
		},
	}
	if id != "0" {
		keys["id"] = id
		keyOpts["id"] = "<>"
	}
	err := t.Struct.CountWithKeys(entity, &count,
		keys,
		keyOpts,
		dao,
	)
	if err != nil {
		return err
	}
	fmt.Println("count", count)
	if count > 0 {
		return errors.New("unique restrain failed")
	}

	return nil
}

func (t *Table[MODEL]) handle_ban(entity MODEL, arg Arg) error {
	value := reflect.ValueOf(entity).FieldByName(arg.Name)
	if util.IsZeroValue(value) {
		return nil
	}
	return errors.New(arg.JsonName + " not allowed")
}
