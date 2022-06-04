package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"strings"
)

// generate code
func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/query",
		Mode:    gen.WithDefaultQuery,
	})
	dataMap := map[string]func(detailType string) (dataType string){
		"int": func(detailType string) (dataType string) { return "int64" },
		// bool mapping
		"tinyint": func(detailType string) (dataType string) {
			if strings.HasPrefix(detailType, "tinyint(1)") {
				return "bool"
			}
			return "int8"
		},
	}
	g.WithDataTypeMap(dataMap)
	db, _ := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/go-gateway?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db)
	// g.ApplyBasic(g.GenerateAllTable()...)
	g.ApplyBasic(g.GenerateModel("gateway_admin"))
	g.Execute()

}
