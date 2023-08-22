package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var MysqlTimeParse = "Local"

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

func loadDBConfig() *DBConfig {
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

func insertDataV3(db *sql.DB, ch chan int) {
	// 插入的数据量
	totalRows := 8000000 - 1699176

	// 构建插入语句
	insertQuery := "INSERT INTO pt (data, next_at, update_at, created_at, deleted, invalid) VALUES (?, ?, ?, ?, ?, ?)"

	// 准备插入语句
	stmt, err := db.Prepare(insertQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// 开始事务
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	// 批量插入数据
	batchSize := 100000 // 每批插入的数量

	for i := 1; i <= totalRows; i += batchSize {
		wg := sync.WaitGroup{}
		batchData := make([]interface{}, 0, batchSize*6) // 预先分配空间
		for j := i; j < i+batchSize && j <= totalRows; j++ {
			wg.Add(1)

			go func(j int) {
				defer wg.Done()
				var del, inva bool
				if j%2 == 0 {
					del = true
					inva = false
				} else {
					del = false
					inva = true
				}

				tasnextAt := time.Now().Unix() + int64(j)
				data := fmt.Sprintf("data index %d", j) // 根据实际需求生成数据
				nextAt := tasnextAt                     // 根据实际需求生成数据
				updateAt := time.Now().Unix()           // 根据实际需求生成数据
				createdAt := time.Now().Unix()          // 根据实际需求生成数据
				deleted := del                          // 根据实际需求生成数据
				invalid := inva                         // 根据实际需求生成数据

				// 将每组数据作为一个整体添加到 batchData 中
				rowData := []interface{}{data, nextAt, updateAt, createdAt, deleted, invalid}
				batchData = append(batchData, rowData)

			}(j)
		}

		// 执行插入语句
		for _, rowData := range batchData {
			rowData, ok := rowData.([]interface{})
			if !ok {
				log.Fatal("rowData 类型断言失败")
			}
			_, err = stmt.Exec(rowData...)
			if err != nil {
				log.Fatal(err)
			}
		}
		wg.Wait()
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("插入数据完成")
}

var totalRows = 8000000

var chMaxCon chan int
var chBatchData chan []interface{}

func init() {
	chMaxCon = make(chan int, 256)
	chBatchData = make(chan []interface{}, 256)
}

// totalRows = 8000000 - 4145125
func insertDataV4(db *sql.DB, ch chan int) {
	var count int
	// 构建插入语句
	insertQuery := "INSERT INTO pt (data, next_at, update_at, created_at, deleted, invalid) VALUES (?, ?, ?, ?, ?, ?)"

	// 准备插入语句
	stmt, err := db.Prepare(insertQuery)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		stmt.Close()
	}()

	for k := 0; k < 250; k++ {
		/*  生成数据 */
		chMaxCon <- 100000 // 每批插入的数量 4197769 4226082 4235453-4247749
		go genData()

		//  插入数据
		go func() {
			batchData := <-chBatchData
			if batchData != nil {
				for _, rowData := range batchData {
					rowData, ok := rowData.([]interface{})
					if !ok {
						log.Fatal("rowData 类型断言失败")
					}
					_, err = stmt.Exec(rowData...)
					if err != nil {
						log.Fatal(err)
					}
					count++
				}
			}
		}()
		fmt.Println("count", count)
	}

	fmt.Println("插入数据完成")
}

// batchSize := 100000 // 每批插入的数量
func genData() {
	var batchSize int
	batchSize = <-chMaxCon
	batchData := make([]interface{}, 0, batchSize*6) // 预先分配空间

	for j := 1; j < batchSize; j++ {
		var del, inva bool
		if j%2 == 0 {
			del = true
			inva = false
		} else {
			del = false
			inva = true
		}

		tasnextAt := time.Now().Unix() + int64(j)
		data := fmt.Sprintf("data index %d", j) // 根据实际需求生成数据
		nextAt := tasnextAt                     // 根据实际需求生成数据
		updateAt := time.Now().Unix()           // 根据实际需求生成数据
		createdAt := time.Now().Unix()          // 根据实际需求生成数据
		deleted := del                          // 根据实际需求生成数据
		invalid := inva                         // 根据实际需求生成数据

		rowData := []interface{}{data, nextAt, updateAt, createdAt, deleted, invalid}
		batchData = append(batchData, rowData)
	}

	fmt.Println("生成数据成功<-chMaxCon:", <-chMaxCon)
	chBatchData <- batchData
}

func main() {
	// 连接到数据库
	dbCfg := loadDBConfig()
	option := dbCfg.ToOpition("charset", "interpolateParams", "allowNativePasswords")
	if MysqlTimeParse != "" {
		option = fmt.Sprintf("%s&parseTime=true&loc=%s", option, MysqlTimeParse)
	}
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", dbCfg.DbUser, dbCfg.DbPswd, dbCfg.DbHost, dbCfg.DbPort, dbCfg.DbName, option)
	mysql, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatal("数据库注册失败")
	}

	err = mysql.Ping()
	if err != nil {
		log.Fatal("Ping数据库注册失败", err)
	}
	defer mysql.Close()

	mysql.SetMaxOpenConns(256)
	mysql.SetMaxIdleConns(100)
	mysql.SetConnMaxLifetime(time.Second * 30)

	insertDataV5(mysql, 10000)
	//log.Fatal("数据库注册success")

	for {
		select {

		default:
			fmt.Println("111")
		}
	}

}

var chIn chan bool

func init() {
	chIn = make(chan bool, 10)
}

func insertDataV5(db *sql.DB, batchSize int) {
	var totalCount int

	insertQuery := "INSERT INTO pt (data, next_at, update_at, created_at, deleted, invalid) VALUES (?, ?, ?, ?, ?, ?)"

	// 准备插入语句
	stmt, err := db.Prepare(insertQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for k := 0; k < 250; k++ {
		// 生成数据

		batchData := genDataV2(batchSize)

		fmt.Println("batchData info lent:", len(batchData))

		// 执行批量插入操作
		if <-chIn {
			go func(data []interface{}) {
				for _, rowData := range data {
					rowData, ok := rowData.([]interface{})
					fmt.Println("rowData info :", rowData)
					if !ok {
						log.Fatal("rowData 类型断言失败")
					}
					_, err := stmt.Exec(rowData...)
					if err != nil {
						log.Fatal(err)
					}
					count := len(data)
					totalCount += count
				}
			}(batchData)
		} else {
			fmt.Println("no data")
		}

	}

	fmt.Println("插入数据完成，总计插入", totalCount, "条数据")
}

func genDataV2(batchSize int) []interface{} {
	batchData := make([]interface{}, 0, batchSize*6) // 预先分配空间

	for j := 1; j <= batchSize; j++ {
		var deleted, invalid bool
		if j%2 == 0 {
			deleted = true
			invalid = false
		} else {
			deleted = false
			invalid = true
		}

		tasnextAt := time.Now().Unix() + int64(j)
		data := fmt.Sprintf("data index %d", j)
		nextAt := tasnextAt
		updateAt := time.Now().Unix()
		createdAt := time.Now().Unix()
		rowData := []interface{}{data, nextAt, updateAt, createdAt, deleted, invalid}
		batchData = append(batchData, rowData)
	}

	chIn <- true
	return batchData
}
