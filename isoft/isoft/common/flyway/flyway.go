package flyway

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"isoft/isoft/common/hashutil"
	"strings"
)

type FlyWay struct {
	Dsn string						// dsn 连接串
	MigrationLocation string		// 迁移文件路径
	db *sql.DB
}

func checkError(err error)  {
	if err != nil{
		panic(err)
	}
}

// 执行 sql 迁移的方法
func (this *FlyWay) Migrate() {
	this.InitDataSourceConn()
	this.InitFlyWayVersionTable()
	this.CheckAndCompare()
	defer this.db.Close()
}

func (this *FlyWay) InitDataSourceConn()  {
	if this.Dsn == ""{
		panic("empty dsn error...")
	}
	// 建立连接
	db, err := sql.Open("mysql", this.Dsn)
	checkError(err)
	this.db = db
}

// 建立迁移文件版本管理表
func (this *FlyWay) InitFlyWayVersionTable()  {
	versionTable := `CREATE TABLE IF NOT EXISTS flyway_version (id INT(20) PRIMARY KEY AUTO_INCREMENT,hash CHAR(200),sql_detail TEXT,created_time datetime);`
	this.ExecSQL(versionTable)
}

func (this *FlyWay) ExecSQL(sql string)  {
	stmt, err := this.db.Prepare(sql)
	checkError(err)
	_, err = stmt.Exec()
	checkError(err)
}

// 检查迁移文件格式是否正确,并且和 DB 里面的执行记录进行对比
func (this *FlyWay) CheckAndCompare()  {
	migrationsFile := this.ReadAllFileMigrations(this.MigrationLocation)
	migrationsBD := this.ReadAllDBMigrations()
	writeMigrations := this.Compare(migrationsFile,migrationsBD)
	this.Execute(writeMigrations)
}

// 执行sql迁移记录
func (this *FlyWay) Execute(migrations []*Migration)  {
	for _, migration := range migrations{
		this.ExecSQL(migration.sql)
		writeVersionRecord := fmt.Sprintf(
			`INSERT INTO flyway_version(HASH,SQL_DETAIL,CREATED_TIME) VALUES ('%s','%s',NOW());`,migration.hash,migration.sql)
		this.ExecSQL(writeVersionRecord)
	}
}

// 文件中的迁移和DB中的迁移进行对比,返回待执行的迁移
func (this *FlyWay) Compare(migrationsFile,migrationsBD []*Migration) []*Migration {
	if len(migrationsFile) < len(migrationsBD){
		panic("redundant migration record found...")
	}
	for index,migration := range migrationsBD{
		if migrationsFile[index].hash != migration.hash{
			panic(fmt.Sprintf("hash matching error for sql: [file sql:%s],[DB sql:%s], please check it...",
				getShortSql(migrationsFile[index].sql),getShortSql(migration.sql)))
		}
	}
	return migrationsFile[len(migrationsBD):len(migrationsFile)]
}

func getShortSql(sql string) string {
	if len(sql) > 80{
		return sql[0:80] + "..."
	}
	return sql
}

type Migration struct {
	sql string
	hash string
}

func (this *FlyWay) ReadAllDBMigrations() []*Migration {
	sql := `select hash,sql_detail from flyway_version`
	rows, err := this.db.Query(sql)
	checkError(err)
	defer rows.Close()
	migrations := make([]*Migration,0,10)
	for rows.Next(){
		migration := new(Migration)
		err = rows.Scan(&migration.hash, &migration.sql)
		checkError(err)
		migrations = append(migrations, migration)
	}
	return migrations
}

func (this *FlyWay) ReadAllFileMigrations(migrationLocation string) []*Migration {
	if migrationLocation == ""{
		panic("empty migrationLocation file...")
	}
	bytes, err := ioutil.ReadFile(migrationLocation)
	checkError(err)
	migrationStrArr := strings.Split(string(bytes), "\n")
	var migrations []*Migration
	for _,migrationStr := range migrationStrArr{
		migrationStr = strings.TrimSpace(migrationStr)
		// 空行或者注释行不算在内
		if migrationStr == "" || strings.HasPrefix(migrationStr, "--") || strings.HasPrefix(migrationStr, "/*"){
			continue
		}
		migrations = append(migrations, &Migration{sql:migrationStr,hash:hashutil.CalculateHashWithString(migrationStr)})
	}
	return migrations
}
