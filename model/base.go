package model

import(
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-gorp/gorp"
)

var (
    Db *sql.DB
    Dbmap *gorp.DbMap
)

func InitDb()error{
    var err error
    Db, err = sql.Open("mysql", "root:123@tcp(localhost:3306)/goBlog?charset=utf8")
    if err!=nil{
        goto end
    }
    Db.SetMaxOpenConns(20)
    Db.SetMaxIdleConns(10)

    Dbmap = &gorp.DbMap{Db: Db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
    
    Dbmap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")
	Dbmap.AddTableWithName(Category{}, "category").SetKeys(true, "Id")
	Dbmap.AddTableWithName(Tags{}, "tags").SetKeys(true, "Id")
    
    err = Dbmap.CreateTablesIfNotExists()
    
    if err!=nil{
        goto end
    }

end:
    if err!=nil{
        if Db!=nil{
            Db.Close()
        }
    }
    return err
}
