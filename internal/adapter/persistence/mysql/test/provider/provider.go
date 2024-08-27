package provider

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"toktok-backend/internal/adapter/persistence/mysql"
	"toktok-backend/internal/adapter/persistence/mysql/ent"

	_ "github.com/go-sql-driver/mysql"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var p *Provider
var once sync.Once
var MyContext = context.Background()

func Get() *Provider {
	if p == nil {

		once.Do(func() {
			p = &Provider{}

			ctx := context.Background()

			conatiner, err := setupContainer()
			if err != nil {
				log.Fatalf("failed to setup conatiner: %v", err)
			}
			p.container = conatiner

			dsn, err := p.getDatabaseDsn(ctx)
			if err != nil {
				log.Fatalf("failed to get database dsn: %v", err)
			}

			client, err := newClientForProvider(dsn)
			if err != nil {
				log.Fatalf("failed to new client for provider: %v", err)
			}

			err = mysql.AutoMigration(ctx, client)
			if err != nil {
				log.Fatalf("failed to auto migration: %s", err)
			}

			p.Client = client
		})
	}

	return p
}

type Provider struct {
	*mysql.Client
	container testcontainers.Container
}

func (p *Provider) TerminateContainer(ctx context.Context) error {
	err := p.container.Terminate(ctx)
	if err != nil {
		return err
	}

	return nil
}

func setupContainer() (testcontainers.Container, error) {
	ctx := context.Background()

	// MySQL 컨테이너 설정
	req := testcontainers.ContainerRequest{
		Image:        "mysql:8.0.36",
		ExposedPorts: []string{"3306/tcp", "33060/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "1234",
			"MYSQL_DATABASE":      "test",
		},
		WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),
	}

	// MySQL 컨테이너 생성 및 시작
	mysqlContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	return mysqlContainer, nil
}

func (p *Provider) getDatabaseDsn(ctx context.Context) (string, error) {
	if p.container == nil {
		return "", errors.New("did not create a container for p")
	}

	// 컨테이너 호스트 주소 얻기
	host, err := p.container.Host(ctx)
	if err != nil {
		return "", err
	}

	// 매핑된 포트 얻기
	port, err := p.container.MappedPort(ctx, "3306/tcp")
	if err != nil {
		return "", err
	}

	// 데이터베이스 연결 문자열 생성
	dsn := fmt.Sprintf("root:1234@tcp(%s:%s)/test?parseTime=True", host, port.Port())

	return dsn, nil
}

func newClientForProvider(dsn string) (*mysql.Client, error) {

	ist, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	client := mysql.Client{
		Client: ist,
	}

	return &client, nil
}
