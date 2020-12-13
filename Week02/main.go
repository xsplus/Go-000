package main

import (
    "database/sql"
    "fmt"

    xerrors "github.com/pkg/errors"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    actName, err := GetActNameByID(10)

    if xerrors.Is(err, NotFound) {
        //fmt.Printf("original error: \n %T %v\n", xerrors.Cause(err), xerrors.Cause(err))
        //fmt.Printf("stack trace: \n%+v\n", err)
        fmt.Println(err.Error())
        fmt.Println(xerrors.Cause(err).Error())
        return
    }
    if err != nil {
        fmt.Printf("stack trace: \n%+v\n", err)
        return
    }
    fmt.Println(actName)
}

// dao

var (
    NotFound = xerrors.New("Data Not Found")
)

// GetActNameByID 获取活动名称
// @param id 活动ID
// @return actName 活动名称（默认为空字符串）
// @return error 错误
func GetActNameByID(id int) (actName string, err error) {
    db, _ := sql.Open("mysql", "root:123456@@tcp(110.110.110.110:3306)/study?charset=utf8")
    querySql := fmt.Sprintf("SELECT act_name FROM t_act_list WHERE id = %d", id)
    e := db.QueryRow(querySql).Scan(&actName)
    if e != nil {
        if e == sql.ErrNoRows {
            return "", xerrors.Wrapf(NotFound, fmt.Sprintf("sql: %s error: %v", querySql, e))
        }
        return "", e
    }
    return actName, nil
}
