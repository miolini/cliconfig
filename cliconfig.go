package cliconfig

import (
	"reflect"
	"strings"

	"github.com/serenize/snaker"
	"github.com/urfave/cli"
)

func Fill(config interface{}, envPrefix string) []cli.Flag {
	configValue := reflect.Indirect(reflect.ValueOf(config))
	var flags []cli.Flag
	for i := 0; i < configValue.NumField(); i++ {
		fieldValue := configValue.Field(i)
		fieldType := configValue.Type().Field(i)
		name := snaker.CamelToSnake(fieldType.Name)
		flagName := fieldType.Tag.Get("flag")
		if flagName == "" {
			flagName = name
		}
		envName := fieldType.Tag.Get("env")
		if envName == "" {
			envName = strings.ToUpper(flagName)
		}
		envName = envPrefix + envName
		switch fieldType.Type.Kind() {
		case reflect.String:
			flag := cli.StringFlag{
				Name:        flagName,
				EnvVar:      envName,
				Destination: fieldValue.Addr().Interface().(*string),
			}
			flags = append(flags, flag)
		}
	}
	return flags
}
