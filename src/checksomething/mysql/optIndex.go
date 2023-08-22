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

func main() {
	// 连接到数据库
	dbCfg := loadDBConfig()
	option := dbCfg.ToOpition("charset", "interpolateParams", "allowNativePasswords")

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

	_, err = fmt.Println("数据库注册success")

	// 插入数据
	{
		start := time.Now().Unix()
		var batchSize = 10000
		var control = 10

		// 重置index： 每次插入的时候需要做修改
		incr = 1010086

		fmt.Printf("插入%d条数据，总共用时：%d s\n", batchSize*control, insertDataV5(mysql, batchSize, control)-start)
	}

	// todo 查询接口有问题
	/*
		创建索引1
			ALTER TABLE `pt` ADD INDEX `idx_query` (`next_at`, `deleted`, `invalid`)
	*/
	//createIndex(mysql, "ALTER TABLE pt ADD INDEX idx_query (next_at, deleted, invalid)")

	// 数据查找
	findData(mysql)

	/* 索引优化
	创建索引2:
		ALTER TABLE `pt` ADD INDEX `idx_query` (`deleted`, `invalid`, `next_at`)
	*/
	// createIndex(mysql, "ALTER TABLE pt ADD INDEX idx_query1 (next_at, deleted, invalid)")
	// findData(mysql,"EXPLAIN SELECT id FROM pt WHERE next_at <= 1514822400 AND deleted=0 AND invalid=0")

}

var chIn chan bool
var incr int

func init() {
	chIn = make(chan bool, 10)
	incr = 10086
}

func insertDataV5(db *sql.DB, batchSize int, control int) int64 {
	var totalCount int

	insertQuery := "INSERT INTO pt (data, next_at, update_at, created_at, deleted, invalid) VALUES (?, ?, ?, ?, ?, ?)"

	// 准备插入语句
	stmt, err := db.Prepare(insertQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		fmt.Println(stmt.Close())
	}()

	var wg sync.WaitGroup
	for k := 0; k < control; k++ {
		// 生成数据

		batchData := genDataV2(batchSize)
		fmt.Println("batchData info lent:", len(batchData))

		// 执行批量插入操作
		if <-chIn {
			wg.Add(1)
			go func(data []interface{}) {
				defer wg.Done()

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
		}
		wg.Wait()
	}

	fmt.Println("插入数据完成，总计插入", totalCount, "条数据")
	return time.Now().Unix()
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
		incr++
		//tasnextAt := time.Now().Unix() + int64(j)
		data := fmt.Sprintf("data index %d", j)
		nextAt := int64(incr)
		updateAt := time.Now().Unix()
		createdAt := time.Now().Unix()
		rowData := []interface{}{data, nextAt, updateAt, createdAt, deleted, invalid}
		batchData = append(batchData, rowData)
	}

	chIn <- true
	return batchData
}

// 创建索引："ALTER TABLE pt ADD INDEX idx_query (next_at, deleted, invalid)"
func createIndex(db *sql.DB, sql_str string) {
	// 执行ALTER TABLE命令
	_, err := db.Exec(sql_str)
	if err != nil {
		fmt.Println("ALTER TABLE命令执行失败：", err.Error())
		return
	}

	fmt.Println("ALTER TABLE命令执行成功")
}

//  查询

func findData1(db *sql.DB) {

	// 执行查询并解释执行计划
	rows, err := db.Query("EXPLAIN SELECT id FROM pt WHERE next_at >= 1692721116 AND deleted=1 AND invalid=0")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 打印解释执行计划结果
	var (
		id int
	)
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("id: %d\n", id)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func findData(db *sql.DB) {
	query := "SELECT * FROM bxch"

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		// process rows
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
	}

	// Explain query execution plan
	explainQuery := "EXPLAIN FORMAT=json " + query
	explainRows, err := db.Query(explainQuery)
	if err != nil {
		fmt.Println("Error getting query execution plan:", err)
		return
	}
	defer explainRows.Close()

	for explainRows.Next() {
		// process explain rows
	}

	if err := explainRows.Err(); err != nil {
		fmt.Println("Error iterating explain rows:", err)
	}

}

/*

表数据：1110086条
+-------+------------+-------------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+---------+------------+
| Table | Non_unique | Key_name    | Seq_in_index | Column_name | Collation | Cardinality | Sub_part | Packed | Null | Index_type | Comment | Index_comment | Visible | Expression |
+-------+------------+-------------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+---------+------------+
| pt    |          0 | PRIMARY     |            1 | id          | A         |     4285914 |     NULL |   NULL |      | BTREE      |         |               | YES     | NULL       |
| pt    |          1 | idx_next_at |            1 | next_at     | A         |     2347438 |     NULL |   NULL | YES  | BTREE      |         |               | YES     | NULL       |
| pt    |          1 | Uindex_1    |            1 | next_at     | A         |     2407392 |     NULL |   NULL | YES  | BTREE      |         |               | YES     | NULL       |
| pt    |          1 | Uindex_1    |            2 | deleted     | A         |     2928384 |     NULL |   NULL |      | BTREE      |         |               | YES     | NULL       |
| pt    |          1 | Uindex_1    |            3 | invalid     | A         |     2820960 |     NULL |   NULL |      | BTREE      |         |               | YES     | NULL       |
+-------+------------+-------------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+---------+------------+
5 rows in set (0.01 sec)

索引数据：
+------------+--------------+------+-----+---------+----------------+
| Field      | Type         | Null | Key | Default | Extra          |
+------------+--------------+------+-----+---------+----------------+
| id         | int unsigned | NO   | PRI | NULL    | auto_increment |
| data       | longtext     | YES  |     | NULL    |                |
| next_at    | bigint       | YES  | MUL | NULL    |                |
| update_at  | int          | NO   |     | NULL    |                |
| created_at | int          | NO   |     | NULL    |                |
| deleted    | tinyint(1)   | NO   |     | 0       |                |
| invalid    | tinyint(1)   | NO   |     | 0       |                |
+------------+--------------+------+-----+---------+----------------+

*/

/*
rowData info : [data index 10000 1110086 1692718483 1692718483 true false]

利用联合索引 Uindex_1
	ALTER TABLE `pt` ADD INDEX `Uindex_1` (`next_at`, `deleted`, `invalid`);

	mysql> EXPLAIN SELECT id FROM pt WHERE next_at <= 555043 AND deleted = 1 AND invalid = 0;
	+----+-------------+-------+------------+-------+---------------+----------+---------+------+--------+----------+--------------------------+
	| id | select_type | table | partitions | type  | possible_keys | key      | key_len | ref  | rows   | filtered | Extra                    |
	+----+-------------+-------+------------+-------+---------------+----------+---------+------+--------+----------+--------------------------+
	|  1 | SIMPLE      | pt    | NULL       | range | Uindex_1      | Uindex_1 | 11      | NULL | 524635 |     1.00 | Using where; Using index |
	+----+-------------+-------+------------+-------+---------------+----------+---------+------+--------+----------+--------------------------+
	1 row in set, 1 warning (0.00 sec)


对整个表进行对半查询，555043，查询了524635条
	mysql> EXPLAIN SELECT id FROM pt WHERE next_at <= 555043 AND deleted = 1 AND invalid = 0;
	+----+-------------+-------+------------+-------+-------------------+----------+---------+------+--------+----------+--------------------------+
	| id | select_type | table | partitions | type  | possible_keys     | key      | key_len | ref  | rows   | filtered | Extra                    |
	+----+-------------+-------+------------+-------+-------------------+----------+---------+------+--------+----------+--------------------------+
	|  1 | SIMPLE      | pt    | NULL       | range | Uindex_1,Uindex_2 | Uindex_1 | 11      | NULL | 524635 |    10.00 | Using where; Using index |
	+----+-------------+-------+------------+-------+-------------------+----------+---------+------+--------+----------+--------------------------+
	1 row in set, 1 warning (0.01 sec)

	mysql> EXPLAIN SELECT id FROM pt WHERE next_at <= 110087 AND deleted = 1 AND invalid = 0;
	+----+-------------+-------+------------+-------+---------------+----------+---------+------+--------+----------+--------------------------+
	| id | select_type | table | partitions | type  | possible_keys | key      | key_len | ref  | rows   | filtered | Extra                    |
	+----+-------------+-------+------------+-------+---------------+----------+---------+------+--------+----------+--------------------------+
	|  1 | SIMPLE      | pt    | NULL       | range | Uindex_1      | Uindex_1 | 11      | NULL | 236296 |     1.00 | Using where; Using index |
	+----+-------------+-------+------------+-------+---------------+----------+---------+------+--------+----------+--------------------------+

查询语句对效率的影响
	EXPLAIN SELECT id FROM pt WHERE deleted = 1 AND invalid = 0 AND next_at <= 555043;  ===== 555043;
	EXPLAIN SELECT id FROM pt WHERE next_at <= 110087 AND deleted = 1 AND invalid = 0;  ===== 236296
	EXPLAIN SELECT id FROM pt WHERE next_at <= 110087 AND invalid = 0 AND deleted = 1;  ===== 236296
	EXPLAIN SELECT id FROM pt WHERE deleted = 1 AND next_at <= 555043 AND invalid = 0;  ===== 524635


*/

/*
	index 优化

	----ALTER TABLE `pt` ADD INDEX `Uindex_2` ( `deleted`, `invalid`，`next_at`);
	----空格报错	// ERROR 1064 (42000): You have an error in your SQL syntax; check the manual that corresponds to
	----your MySQL server version for the right syntax to use near '，`next_at`)' at line 1

创建新的联合索引
	ALTER TABLE `pt` ADD INDEX `Uindex_2` (`deleted`, `invalid`, `next_at`);


	mysql> EXPLAIN SELECT id FROM pt WHERE next_at = 1110086 AND deleted = 1 AND invalid = 0;
	+----+-------------+-------+------------+------+---------------+----------+---------+-------------------+------+----------+-------------+
	| id | select_type | table | partitions | type | possible_keys | key      | key_len | ref               | rows | filtered | Extra       |
	+----+-------------+-------+------------+------+---------------+----------+---------+-------------------+------+----------+-------------+
	|  1 | SIMPLE      | pt    | NULL       | ref  | Uindex_2      | Uindex_2 | 11      | const,const,const |    1 |   100.00 | Using index |
	+----+-------------+-------+------------+------+---------------+----------+---------+-------------------+------+----------+-------------+


	mysql> EXPLAIN SELECT id FROM pt WHERE next_at <= 110087 AND deleted = 1 AND invalid = 0;
	+----+-------------+-------+------------+-------+---------------+----------+---------+------+--------+----------+--------------------------+
	| id | select_type | table | partitions | type  | possible_keys | key      | key_len | ref  | rows   | filtered | Extra                    |
	+----+-------------+-------+------------+-------+---------------+----------+---------+------+--------+----------+--------------------------+
	|  1 | SIMPLE      | pt    | NULL       | range | Uindex_2      | Uindex_2 | 11      | NULL | 103210 |   100.00 | Using where; Using index |
	+----+-------------+-------+------------+-------+---------------+----------+---------+------+--------+----------+--------------------------+


*/

/*
	对比：


	mysql> EXPLAIN SELECT id FROM pt WHERE next_at <= 110087 AND deleted = 1 AND invalid = 0;
	+----+-------------+-------+------------+-------+---------------+----------+---------+------+--------+----------+--------------------------+
	| id | select_type | table | partitions | type  | possible_keys | key      | key_len | ref  | rows   | filtered | Extra                    |
	+----+-------------+-------+------------+-------+---------------+----------+---------+------+--------+----------+--------------------------+
	|  1 | SIMPLE      | pt    | NULL       | range | Uindex_2      | Uindex_2 | 11      | NULL | 103210 |   100.00 | Using where; Using index |
	+----+-------------+-------+------------+-------+---------------+----------+---------+------+--------+----------+--------------------------+



	mysql> EXPLAIN SELECT id FROM pt WHERE next_at <= 110087 AND deleted = 1 AND invalid = 0;
	+----+-------------+-------+------------+-------+---------------+----------+---------+------+--------+----------+--------------------------+
	| id | select_type | table | partitions | type  | possible_keys | key      | key_len | ref  | rows   | filtered | Extra                    |
	+----+-------------+-------+------------+-------+---------------+----------+---------+------+--------+----------+--------------------------+
	|  1 | SIMPLE      | pt    | NULL       | range | Uindex_1      | Uindex_1 | 11      | NULL | 236296 |     1.00 | Using where; Using index |
	+----+-------------+-------+------------+-------+---------------+----------+---------+------+--------+----------+--------------------------+


	Uindex_1：236296
	Uindex_2：103210
	优化后的Uindex_2 查询效率确实是比Uindex_1效率要高



*/
