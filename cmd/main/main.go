package main

import (
	"fmt"
	"time"

	"github.com/he-he-heng/toktok-backend/config"
	mybolt "github.com/he-he-heng/toktok-backend/config/database/bolt"
	mymysql "github.com/he-he-heng/toktok-backend/config/database/mysql"
	"github.com/he-he-heng/toktok-backend/model"
	"github.com/he-he-heng/toktok-backend/pnumstate"
	"github.com/he-he-heng/toktok-backend/pnumstate/repo"

	"github.com/he-he-heng/toktok-backend/pkg/utils"
)

func main() {
	config := utils.Must(config.NewConfig(".env.yaml"))

	db := utils.Must(mymysql.NewDatabse(config))
	bolt := utils.Must(mybolt.NewDatabase(config))
	utils.Check(db.AutoMigrate(&model.User{}))

	var pnumState pnumstate.PnumStateRepo = repo.NewPnumStateRepository(bolt)

	_, err := pnumState.Create("01081315423", &model.PnumStatus{
		Cnt:       5,
		Timestamp: time.Now(),
		Break:     true,
		AuthCode:  192931,
	})

	if err != nil {
		panic(err)
	}

	res, err := pnumState.ReadByPnum("01081315423")
	if err != nil {
		panic(err)
	}

	fmt.Println("res: ", res)

}
