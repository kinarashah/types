package mapper

import (
	"github.com/rancher/norman/types"
)

type TestMapper struct {
}

func (s TestMapper) FromInternal(data map[string]interface{}) {
	//test := convert.ToString(values.GetValueN(data, "data", "digitaloceancredentialConfig"))
	//if test != "" {
	//	ans, _ := json.Marshal(data)
	//	logrus.Infof("entered here %v", string(ans))
	//
	//	delete(data, "data")
	//	values.PutValue(data, convert.ToString(test), "digitaloceancredentialConfig", "accessToken")
	//
	//	ans, _ = json.Marshal(data)
	//	logrus.Infof("now from here %v", string(ans))
	//}
}

func (s TestMapper) ToInternal(data map[string]interface{}) error {
	//ans, _ := json.Marshal(data)
	//logrus.Infof("entered here %v", string(ans))
	//
	//stringData := map[string]string{}
	//stringData["digitaloceancredentialConfig"] = convert.ToString(values.GetValueN(data, "digitaloceancredentialConfig", "accessToken"))
	//
	//values.PutValue(data, stringData, "stringData")
	//delete(data, "digitaloceancredentialConfig")
	//
	//ans, _ = json.Marshal(data)
	//logrus.Infof("now here %v", string(ans))
	return  nil
}

func (s TestMapper) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}

