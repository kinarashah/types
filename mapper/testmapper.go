package mapper

import (
	"encoding/json"
	"github.com/rancher/norman/types"
	"github.com/sirupsen/logrus"
)

type TestMapper struct {
}

func (s TestMapper) FromInternal(data map[string]interface{}) {
	ans, _ := json.Marshal(data)
	logrus.Infof("entered here %v", string(ans))
}

func (s TestMapper) ToInternal(data map[string]interface{}) error {
	ans, _ := json.Marshal(data)
	logrus.Infof("entered here %v", string(ans))
	return  nil
}

func (s TestMapper) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}

