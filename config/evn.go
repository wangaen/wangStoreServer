package config

import (
	"fmt"
	"github.com/spf13/viper"
	"reflect"
	"strconv"
	"wangStoreServer/common/models"
)

func GetConfigEnv() *models.Configs {
	viperConfig := viper.New()
	viperConfig.AddConfigPath("config") //设置读取的文件路径
	viperConfig.SetConfigName("config") //设置读取的文件名
	viperConfig.SetConfigType("yaml")   //设置文件的类型
	//尝试进行配置读取
	if err := viperConfig.ReadInConfig(); err != nil {
		fmt.Println("读取配置文件异常：", err.Error())
		return nil
	}
	var envs = models.Configs{}
	setEnvValue(&envs, viperConfig)
	return &envs
}

// 修改结构体属性值
func setEnvValue(envObj *models.Configs, viperConfig *viper.Viper) {
	envObjT := reflect.TypeOf(envObj)
	envObjV := reflect.ValueOf(envObj)
	for i := 0; i < envObjT.Elem().NumField(); i++ {
		// viperConfig 父配置项名
		ParentYamlTag := envObjT.Elem().Field(i).Tag.Get("yaml")
		// Configs 子配置项的所有字段
		childField := envObjT.Elem().Field(i).Type
		for j := 0; j < childField.NumField(); j++ {
			// Configs 子配置项的名
			childFieldName := childField.Field(j).Name
			// Configs 子配置项的值
			childFieldValueName := envObjV.Elem().Field(i).FieldByName(childFieldName)
			// viperConfig子配置项名
			childYamlTag := childField.Field(j).Tag.Get("yaml")
			// viperConfig 子配置项的值
			childConfigValue := viperConfig.Get(ParentYamlTag + "." + childYamlTag)
			// 赋值
			switch childFieldValueName.Type().Name() {
			case "int64":
				val := int64(childConfigValue.(int))
				childFieldValueName.SetInt(val)
			case "string":
				val, ok := childConfigValue.(int)
				if ok {
					// 纯数字的字符串误认为int
					childFieldValueName.SetString(strconv.Itoa(val))
				} else {
					childFieldValueName.SetString(childConfigValue.(string))
				}
			case "bool":
				childFieldValueName.SetBool(childConfigValue.(bool))
			}
		}
	}
}
