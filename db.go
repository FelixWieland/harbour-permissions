package harbourpermissions

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql" //sql driver
)

type sqlQuery string

type loginCredentials struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

var errCantConnectToDB = errors.New("Cant connect to DB")

func (sqlQuery sqlQuery) prep(vals ...string) string {
	if len(vals) > strings.Count(string(sqlQuery), "?") {
		panic("Too many arguments in prep call")
	}

	buffer := string(sqlQuery)
	for _, elm := range vals {
		buffer = strings.Replace(string(buffer), "?", "'"+elm+"'", 1)
	}
	return buffer
}

func loadCredentials(path string) loginCredentials {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully Opened " + path)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	data := loginCredentials{}
	json.Unmarshal(byteValue, &data)
	return data
}

func (loginCredentials *loginCredentials) toString() string {
	//username:password@tcp(127.0.0.1:3306)/test
	return loginCredentials.User + ":" + loginCredentials.Password + "@tcp(" + loginCredentials.Host + ":" + loginCredentials.Port + ")/" + loginCredentials.Database + "?charset=utf8mb4"
}

func connectToDB(connString string) (*sql.DB, error) {
	l, err := sql.Open("mysql", connString) //"astaxie:astaxie@/test?charset=utf8"
	if err == nil && nil == l.Ping() {
		return l, nil
	}
	return &sql.DB{}, errCantConnectToDB
}
