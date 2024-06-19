package main

import (
	"github.com/he-he-heng/toktok-backend/config"
	"github.com/he-he-heng/toktok-backend/config/database/bolt"
	"github.com/he-he-heng/toktok-backend/config/database/mysql"
	"github.com/he-he-heng/toktok-backend/interval/model"
	"github.com/he-he-heng/toktok-backend/pkg/utils"
)

func main() {
	config := utils.Must(config.NewConfig("./env.yaml"))

	db := utils.Must(mysql.NewDatabse(config))
	bolt := utils.Must(bolt.NewDatabase(config))

	utils.Check(db.AutoMigrate(&model.User{}))

	_ = bolt
}
