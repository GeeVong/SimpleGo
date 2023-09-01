package orm

import "strings"

type DBConfig struct {
	DbName   string //数据库名称
	DBType   string //数据库类型
	DbHost   string //数据库地址
	DbPort   int    //数据库端口
	DbUser   string //数据库用户
	DbPswd   string //数据库密码
	DbKey    string //数据库索引
	DbBackUp int    //数据库备份时间
	Opitions map[string]string
}

func (self *DBConfig) GetOpition(name string) string {
	value, ok := self.Opitions[name]
	if ok {
		return value
	}
	return ""
}

func (self *DBConfig) ToOpition(params ...string) string {
	options := make([]string, 0, len(params))
	for i := 0; i < len(params); i++ {
		value := self.GetOpition(params[i])
		if value != "" {
			options = append(options, params[i]+"="+value)
		}
	}
	return strings.Join(options, "&")
}

func LoadDBConfig() *DBConfig {
	cfg := &DBConfig{
		DbName: "bxch",
		DbUser: "root",
		DbPswd: "root",
		DbHost: "127.0.0.1",
		DbPort: 3306,
	}

	varlueMap := make(map[string]string)
	varlueMap["charset"] = "utf8mb4"
	varlueMap["interpolateParams"] = "true"
	varlueMap["allowNativePasswords"] = "true"
	varlueMap["maxIdle"] = "1"
	varlueMap["maxOpen"] = "16"
	varlueMap["parseTime"] = ""
	varlueMap["loc"] = ""
	cfg.Opitions = varlueMap

	return cfg
}
