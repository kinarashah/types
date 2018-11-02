package schema

import (
	"encoding/base64"

	"github.com/sirupsen/logrus"

	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/convert"
)

type CloudCredentialMapper struct {
}

func (e CloudCredentialMapper) FromInternal(data map[string]interface{}) {
	logrus.Infof("ENTER ENTER")
}

func (e CloudCredentialMapper) ToInternal(data map[string]interface{}) error {
	logrus.Infof("ENTER ENTER")
	if data == nil {
		return nil
	}

	auth := convert.ToString(data["auth"])
	username := convert.ToString(data["username"])
	password := convert.ToString(data["password"])

	if auth == "" && username != "" && password != "" {
		data["auth"] = base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	}

	return nil
}

func (e CloudCredentialMapper) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
