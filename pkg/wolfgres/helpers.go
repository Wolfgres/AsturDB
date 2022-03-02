package wolfgres

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var cfgFile string

const (
	PathSeparator = string(os.PathSeparator)
)

// Generate Conn in PostgreSQL with pgx driver from a parameters in config file
func PgxConn() (*pgx.Conn, context.Context) {
	var connStr string = fmt.Sprintf("postgresql://%s:%s@%s:%v/%s", viper.Get("database.admin_user"), viper.Get("database.password"),
		viper.Get("database.host"), viper.GetInt("database.port"), viper.Get("database.database"))
	//log.Debug(connStr)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)

	if err != nil {
		log.Fatal(os.Stderr, "Unable to connecto database: ", err)
		os.Exit(1)
	}

	return conn, ctx
}

// Generate Conn in PostgreSQL with pgx driver to another user
func PgxConnDB(database string, user string) (*pgx.Conn, context.Context) {
	//var connStr string = fmt.Sprintf("postgresql://%s:%s@%s:%v/%s", user, viper.Get("database.password"),
	//	viper.Get("database.host"), viper.GetInt("database.port"), database)
	password := user
	var connStr string = fmt.Sprintf("postgresql://%s:%s@%s:%v/%s", user, password,
		viper.Get("database.host"), viper.GetInt("database.port"), database)
	//log.Debug(connStr)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)

	if err != nil {
		log.Fatal(os.Stderr, " Unable to connecto database: ", err)
		os.Exit(1)
	}

	return conn, ctx
}

/*
// Connection SQLX
func SqlxConn() *sqlx.DB {
	config, err := GetConfig("")

	if err != nil {
		log.Error("Error config file: ", err)
	}

	var connStr string = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", config.Database.User, config.Database.Password,
		config.Database.Host, config.Database.Port, config.Database.Database)

	conn, err := sqlx.Connect("pgx", connStr)

	if err != nil {
		log.Error(os.Stderr, "Unable to connecto database: %v \n", err)
		os.Exit(1)
	}

	return conn
}*/

/*
// Generate SQLite connection with database name and store in data file
// TODO: Verificar si el estado lo vamos a manejar a través del sqlite igual y si
func SqliteConn(database string) *sql.DB {
	exePath := GetExecPath()
	filename := "data" + pathSeparator + database
	f := exePath + pathSeparator + filename
	db, err := sql.Open("sqlite3", f)
	if err != nil {
		log.Error(err)
	}

	return db
}*/

/*
// Connection SQLX GORM
func GormSqliteConn(database string) *gorm.DB {
	exePath := GetExecPath()
	filename := "data" + pathSeparator + database
	f := exePath + pathSeparator + filename
	db, err := gorm.Open(sqlite.Open(f), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}
	return db
}*/

// ReadConfig read a config file to functionality for AsturDB
// Checar *****************************
func GetConfig() (Config, error) {
	filename := ""
	log.Debug(filename)

	var c Config

	if len(filename) == 0 {
		path := "configs"
		file := "config.yaml"
		filename = path + PathSeparator + file
	}

	exePath := GetExecPath()
	var f string
	f = exePath + PathSeparator + filename

	yamlFile, err := ioutil.ReadFile(f)
	if err != nil {
		log.Error("ReadFile -> ", err)
		return c, errors.New("The file " + filename + " don't exist")
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Error("Unmarshall ->", err)
		return c, errors.New("The file failed in Unmarshal")
	}

	return c, nil
}

/*
// Get Queries
func GetQueries(filename string) (Queries, error) {
	var queries Queries

	if len(filename) == 0 {
		filename = "scripts" + pathSeparator + "queries.yaml"
	}

	exePath := GetExecPath()
	f := exePath + pathSeparator + filename
	yamlFile, err := ioutil.ReadFile(f)

	if err != nil {
		log.Error("YamlFile read: ", err)
		return queries, errors.New("The file " + filename + " don't exist")
	}

	err = yaml.Unmarshal(yamlFile, &queries)
	if err != nil {
		log.Error("Unmarshall -> ", err)
		return queries, errors.New("The file failed in Unmarshal")
	}

	return queries, nil
}
*/

// Create Query Insert based in struct
func CreateQueryInsert(schema string, q interface{}) string {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := strings.ToLower(CamelCaseToSnakeCase(reflect.TypeOf(q).Name()))
		query := fmt.Sprintf("INSERT INTO %s.%s VALUES(", schema, t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s'%s'", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s,'%s'", query, v.Field(i).String())
				}
			case reflect.Float64:
				if i == 0 {
					query = fmt.Sprintf("%s%.4f", query, v.Field(i).Float())
				} else {
					query = fmt.Sprintf("%s, %.4f", query, v.Field(i).Float())
				}

			//TODO: Check to convert Value to Struct exported time.Time
			// esto no sirve un jaja
			case reflect.TypeOf(time.Time{}).Kind():
				if i == 0 {
					val := time.Now().Format(time.RFC3339Nano)
					query = fmt.Sprintf("%s'%s'", query, val)
				} else {
					val := v.Field(i).Interface().(time.Time).Format(time.RFC3339Nano)
					query = fmt.Sprintf("%s,'%s'", query, val)
				}
			default:
				fmt.Println("Unsupported type")
				return query
			}
		}
		query = fmt.Sprintf("%s);", query)
		//log.Debug(query)
		return query

	} else {
		return ""
	}
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// Convert camelCase string to snake_case
func CamelCaseToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func FindInt(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

//  Get a exec path in string
func GetExecPath() string {
	exe, err := os.Executable()

	if err != nil {
		log.Error(err)
	}
	s := filepath.Dir(exe)

	if s[0:4] == "/tmp" {
		s = "." + PathSeparator
	}

	return s
}

// Get log.Level from string from config file
func GetLogLevel(level string) log.Level {

	var logLevel log.Level

	switch strings.ToLower(level) {

	case "debug":
		logLevel = log.DebugLevel
	case "info":
		logLevel = log.InfoLevel
	case "warning":
		logLevel = log.WarnLevel
	case "error":
		logLevel = log.ErrorLevel
	case "fatal":
		logLevel = log.FatalLevel
	case "panic":
		logLevel = log.PanicLevel

	}

	return logLevel
}

// Make pgx.Rows to JSON byte[]
func PgSqlRowsToJson(rows pgx.Rows) []byte {
	fieldDescriptions := rows.FieldDescriptions()
	var columns []string
	for _, col := range fieldDescriptions {
		columns = append(columns, string(col.Name))
	}

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)

	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		values, _ := rows.Values()
		for i, v := range values {
			log.Info("v --> ", v)
			valuePtrs[i] = reflect.New(reflect.TypeOf(v)).Interface() // allocate pointer to type
		}
		break
	}

	for rows.Next() {
		rows.Scan(valuePtrs...)

		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := reflect.ValueOf(valuePtrs[i]).Elem().Interface() // dereference pointer
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, _ := json.Marshal(tableData)

	return jsonData
}

// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Info("\r- Ctrl+C pressed in Terminal \n")
		os.Exit(0)
	}()
}

// Save a PID number in a file
func SavePID(pid int) {
	config, err := GetConfig()

	if err != nil {
		log.Error(err)
	}

	file, err := os.Create(config.Deamon.PidFile)

	if err != nil {
		log.Error(err)
	}
	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(pid))

	if err != nil {
		log.Error(err)
	}

	file.Sync() // flush to disk

}
