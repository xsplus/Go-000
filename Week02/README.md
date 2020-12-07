## 作业

#### 问题

我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

答：不应该 Wrap，直接告诉调用者有没有查到数据，查到什么数据就好，调用者不需要关心底层实际 error。（调用者不需要也不应该关心 dao 层调用的是 MySQL 还是 MongoDB 等）

#### 代码

```go
package main

import (
    "database/sql"
    "errors"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    actName, err := GetActNameByID(100)
    fmt.Println(actName, err)

    if err != nil {
        if err.Error() == NOT_FOUND {
            fmt.Println("查不到 相关处理逻辑")
            return
        }
        fmt.Println("未知的异常 相关处理逻辑")
        return
    }
    fmt.Println("查到了 相关处理逻辑")
}

// dao

// 定义dao层数据操作相关常量
const (
    NOT_FOUND string = "Not Found"
)

// GetActNameByID 获取活动名称
// @param id 活动ID
// @return actName 活动名称（默认为空字符串）
// @return error 错误
func GetActNameByID(id int) (actName string, err error) {
    db, _ := sql.Open("mysql", "root:123456@tcp(110.110.110.110:3306)/activity?charset=utf8")
    e := db.QueryRow("SELECT act_name FROM t_act_list WHERE Fid = ?", id).Scan(&actName)
    if e != nil {
        if e == sql.ErrNoRows {
            return "", errors.New(NOT_FOUND)
        }
        return "", e
    }
    return actName, nil
}
```

## 笔记

![](https://github.com/xsplus/Go-000/blob/main/Week02/error.png)