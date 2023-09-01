package main

import (
	"fmt"
	"mysql/orm"
)

/*
	orm在干什么：《业务人员只需要关注，数据连接参数，以及sql语句的书写》
		1. 链接管理
		2. 管理sql语句之下，封装成session
*/

func main() {
	dbCfg := orm.LoadDBConfig()
	option := dbCfg.ToOpition("charset", "interpolateParams", "allowNativePasswords")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", dbCfg.DbUser, dbCfg.DbPswd, dbCfg.DbHost, dbCfg.DbPort, dbCfg.DbName, option)
	engine, _ := orm.NewEngine("mysql", dataSourceName)
	defer engine.Close()

	s := engine.NewSession()

	createTableSQL := `
		CREATE TABLE IF NOT EXISTS item (
			id INT AUTO_INCREMENT PRIMARY KEY,
			price INT NOT NULL
		)
	`

	// 使用 WriteString 方法执行建表语句
	//_, _ = s.Raw(`DROP TABLE IF EXISTS item`).Exec()  // 如果存在表item 则删除
	_, _ = s.Raw(createTableSQL).Exec()

	// 插入数据
	_, _ = s.Raw("INSERT INTO item(`id`,`price`) VALUES (?, ?)", 1, 3200).Exec()
	_, _ = s.Raw("INSERT INTO item(`id`,`price`) VALUES (?, ?)", 2, 3200).Exec()

}
