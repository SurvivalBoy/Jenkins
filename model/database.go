package model

import (
	"database/sql"
	"errors"
	"runtime"
	"sync"

	"jenkins_demo/setting"

	_ "github.com/go-sql-driver/mysql"
	"github.com/phjt-go/logger"
)

var err error
var DB *sql.DB
var once sync.Once

// init
func init() {
	once.Do(func() {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("MySQL connection error, ", err)
			}
		}()

		//这行代码的作用就是开启MySQL的链接
		DB, err = sql.Open("mysql", setting.GetString("mysql.source_name"))
		if nil != err {
			panic(err)
		}

		//设置最大空闲连接数
		DB.SetMaxIdleConns(setting.GetInt("mysql.max_idle_conn"))

		err = DB.Ping()
		if nil != err {
			panic(err)
		}

		logger.Info("MySQL connection successful！")
	})
}

// Close 关闭数据库连接
func Close() {
	DB.Close()
}

func printCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

// Insert 插入操作
func Insert(sql string, args ...interface{}) (int64, error) {
	stmt, err := DB.Prepare(sql)
	printCallerName := printCallerName()
	err = CheckErr(err, printCallerName, "SQL语句设置失败", sql, args)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(args...)
	if err = CheckErr(err, printCallerName, "参数添加失败", sql, args); err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err = CheckErr(err, printCallerName, "插入失败", sql, args); err != nil {
		return 0, err
	}
	logger.Debug("插入成功，ID", id, sql)
	return id, nil
}

// Delete 删除操作
func Delete(sql string, args ...interface{}) error {
	stmt, err := DB.Prepare(sql)
	printCallerName := printCallerName()
	if err = CheckErr(err, printCallerName, "SQL语句设置失败", sql, args); err != nil {
		return err
	}

	result, err := stmt.Exec(args...)
	if err = CheckErr(err, printCallerName, "参数添加失败", sql, args); err != nil {
		return err
	}
	num, err := result.RowsAffected()
	if err = CheckErr(err, printCallerName, "删除失败", sql, args); err != nil {
		return err
	}
	if int(num) == 0 {
		logger.Error(printCallerName+"方法MYSQL错误", "删除失败", sql, args)
		return errors.New("Delete failed. ")
	}
	logger.Debug("删除成功：", "影响条数", num, sql, args)
	return nil
}

// Update 修改操作
func Update(sql string, args ...interface{}) error {
	stmt, err := DB.Prepare(sql)
	printCallerName := printCallerName()
	if err = CheckErr(err, printCallerName, "SQL语句设置失败", sql, args); err != nil {
		return err
	}
	result, err := stmt.Exec(args...)

	if err = CheckErr(err, printCallerName, "参数添加失败", sql, args); err != nil {
		return err
	}
	num, err := result.RowsAffected()
	if err = CheckErr(err, printCallerName, "修改失败", sql, args); err != nil {
		return err
	}
	if int(num) == 0 {
		logger.Error(printCallerName+"方法MYSQL错误", "修改失败", sql, args)
		return errors.New("Update failed. ")
	}
	logger.Debug("修改成功：", "影响条数", num, sql, args)
	return nil
}

func getRows(sql string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := DB.Prepare(sql)
	if err != nil {
		logger.Error("Prepare data error: %v\n", err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	if err != nil {
		logger.Error("Query data error: %v\n", err)
		return nil, err
	}
	return rows, err
}

// CheckErr 检查error
func CheckErr(err error, printCallerName, msg, sql string, args ...interface{}) error {
	if err != nil {
		logger.Error(printCallerName+"方法MYSQL错误", msg, err.Error(), sql, args)
		return err
	}
	return nil
}
