package mapper

import (
	"encoding/json"
	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/convert"
	"github.com/rancher/norman/types/values"
	"github.com/sirupsen/logrus"
)

type TestMapper struct {
}

func (s TestMapper) FromInternal(data map[string]interface{}) {
	test := values.GetValueN(data, "data", "digitaloceancredentialConfig")
	if test != nil {
		ans, _ := json.Marshal(data)
		logrus.Infof("entered here %v", string(ans))

		delete(data, "data")
		values.PutValue(data, convert.ToString(test), "digitaloceancredentialConfig", "accessToken")

		//values.PutValue(data, fmt.Sprintf("%s:%s", convert.ToString(values.GetValueN(data, "id")), convert.ToString(values.GetValueN(data, "namespaceId"))), "id")
		
		ans, _ = json.Marshal(data)
		logrus.Infof("now from here %v", string(ans))
	}
}

func (s TestMapper) ToInternal(data map[string]interface{}) error {
	ans, _ := json.Marshal(data)
	logrus.Infof("entered here %v", string(ans))

	stringData := map[string]string{}
	stringData["digitaloceancredentialConfig"] = convert.ToString(values.GetValueN(data, "digitaloceancredentialConfig", "accessToken"))

	values.PutValue(data, stringData, "stringData")
	delete(data, "digitaloceancredentialConfig")

	ans, _ = json.Marshal(data)
	logrus.Infof("now here %v", string(ans))
	return  nil
}

func (s TestMapper) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}

