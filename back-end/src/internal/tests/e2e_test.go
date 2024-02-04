package tests

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path"
	"src/internal/config"
	"src/internal/models"
	"src/internal/services"
	"src/pkg/utils"
	"strings"
	"testing"
	"time"
)

var (
	workingDir, _ = os.Getwd()
	parts         = strings.Split(workingDir, "/")
	globalConfig  = config.LoadAppConfigFrom(
		"/" + path.Join(
			path.Join(parts[:len(parts)-2]...),
			"configs/app_config.json",
		),
	)
)

type MockBody struct {
	data      []byte
	lastIndex int
}

func (mb *MockBody) Close() error { return nil }
func (mb *MockBody) Read(data []byte) (int, error) {
	n := copy(data, mb.data[mb.lastIndex:])
	mb.lastIndex += n
	return n, nil
}

func NewMockBodyStr(data string) *MockBody {
	return &MockBody{data: []byte(data), lastIndex: 0}
}

func NewMockBodyUser(user models.UserSecret) *MockBody {
	data, err := json.Marshal(&user)
	if err != nil {
		panic("ERROR_JSON_MARSHAL" + err.Error())
	}
	return NewMockBodyStr(string(data))
}

func TestRegister(t *testing.T) {
	var err error
	//var response *http.Response
	var secret models.UserSecret

	gin.SetMode(gin.TestMode)

	secret.Username = fmt.Sprintf("test_user_%d", rand.Intn(100000))
	secret.Password = fmt.Sprintf("p@$$w0rd")

	server := services.NewVanillaServerWithConfig(&globalConfig)

	go func() {
		server.Initialize()
		server.Run()
	}()
	time.Sleep(2 * time.Second)

	// register
	client := http.DefaultClient
	address, err := url.Parse("http://localhost:8099/auth/register")
	response, err := client.Do(&http.Request{
		Method: "POST",
		URL:    address,
		Body:   NewMockBodyUser(secret),
	})
	utils.Logger.Info("asdfasfasdf2134123")
	if err != nil {
		utils.Logger.Info(
			"checking user",
			zap.Int("exists_position", strings.Index(err.Error(), "exists")))
	}

	// login
	header := response.Header.Get("Authorization")
	address, err = url.Parse("http://localhost:8099/auth/login")
	response, err = client.Do(&http.Request{
		Method: "POST",
		URL:    address,
		Body:   NewMockBodyUser(secret),
	})
	if err != nil {

	}
	utils.InitLogging()
	utils.Logger.Info("adsfasfd123123123123123123123")
	fmt.Fprintf(os.Stdout, "HEADER: %s", header)
	// run for 2 min
	time.Sleep(time.Minute * 2)
}
