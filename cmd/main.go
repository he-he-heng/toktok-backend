package main

// import (
// 	"context"
// 	"log"

// 	"toktok-backend/ent"

// 	_ "toktok-backend/ent/runtime"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	client, err := ent.Open("mysql", "root:Qwe123!@tcp(localhost:3306)/toktok?parseTime=True")
// 	if err != nil {
// 		log.Fatalf("failed opening connection to mysql: %v", err)
// 	}
// 	defer client.Close()
// 	// Run the auto migration tool.
// 	if err := client.Schema.Create(context.Background()); err != nil {
// 		log.Fatalf("failed creating schema resources: %v", err)
// 	}

// 	// client.User.CreateBulk(
// 	// 	client.User.Create().SetUID("2hanjunbum6").SetPassword("123456"),
// 	// 	client.User.Create().SetUID("1ckwoal").SetPassword("3829182u37"),
// 	// ).ExecX(context.Background())

// 	// client.Avatar.Create().
// 	// 	SetOwnerID(1).
// 	// 	SetSex(avatar.SexFemale).
// 	// 	SetBirthday("20240607").
// 	// 	SetNickname("장충동왕족발보쌈").
// 	// 	SetIsOnline(true).
// 	// 	SetPicture("test").
// 	// 	SaveX(context.Background())

// 	// client.Avatar.Create().
// 	// 	SetOwnerID(2).
// 	// 	SetSex(avatar.SexFemale).
// 	// 	SetBirthday("20240607").
// 	// 	SetNickname("abcdedf").
// 	// 	SetIsOnline(true).
// 	// 	SetPicture("test").
// 	// 	SaveX(context.Background())

// 	// client.Relation.Create().
// 	// 	SetState(relation.StateNonResponse).
// 	// 	SetAlterState(relation.AlterStateAllow).
// 	// 	AddAvatarIDs(1).
// 	// 	AddAvatarIDs(2).
// 	// 	SaveX(context.Background())

// 	// client.Avatar.CreateBulk(
// 	// 	client.Avatar.Create().SetOwnerID(2).SetSex(avatar.SexFemale).SetBirthday("20240607").SetNickname("kevin").SetIsOnline(true).SetPicture("test").AddrelationIDs(1),
// 	// ).ExecX(context.Background())

// }
