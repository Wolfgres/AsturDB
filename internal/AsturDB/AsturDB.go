package AsturDB

import (
	"AsturDB/pkg/wolfgres"
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var datname string

// Run Stress Test with name
// NameStressTest
func RunStressTest(dbTest string) {
	datname = dbTest

	randomData := RandmonData{
		customers:            viper.GetInt("random_data.customers"),
		employees:            viper.GetInt("random_data.employees"),
		products:             viper.GetInt("random_data.products"),
		districts:            viper.GetInt("random_data.districts"),
		warehouses:           viper.GetInt("random_data.warehouses"),
		districts_warehouses: viper.GetInt("random_data.warehouses") * viper.GetInt("random_data.districts"),
		regions:              viper.GetInt("random_data.regions"),
	}

	if validateTestSchema() {
		log.Info("The user exist's, Do you want resumen the Test?")
	} else {
		createDBUser()
		loadCatalogs(randomData)
	}

	log.Info("Init Orders")
	orders := viper.GetInt("test.orders_by_test")
	var order4Cycles, updates4Clycles, deletes4Clycles int

	// TODO: Ay que verificar como hacer esta validacion

	if viper.GetInt("test.orders_by_loop") > 50 {
		order4Cycles = 50
	} else {
		order4Cycles = viper.GetInt("test.orders_by_loop")
		updates4Clycles = viper.GetInt("test.updates_by_loop")
		deletes4Clycles = viper.GetInt("test.deletes_by_loop")
	}

	min_workers := viper.GetInt("test.min_workers")
	max_workers := viper.GetInt("test.max_workers")
	time_sleep := viper.GetInt("test.time_sleep")
	test_type := viper.GetString("test.type")
	max_conn_sleep := viper.GetInt("test.max_conn_sleep")

	var type_worker int

	for {

	Resume:
		log.Debug("Validate if the max_connections is not reached")
		if maxConnectionsReached() {
			time.Sleep(time.Duration(max_conn_sleep) * time.Second)
			goto Resume
		}

		workers := min_workers + rand.Intn(max_workers-min_workers)

		log.Debug("#########################")
		log.Debug("Workers:: -> ", workers)
		log.Debug("#########################")
		orders = orders - order4Cycles

		for i := 0; i < workers; i++ {

			if test_type != "load" {
				type_worker = 1 + rand.Intn(5-1)
			} else {
				type_worker = 1
			}

			switch type_worker {
			case 1:
				go insertOrder(order4Cycles, randomData)
			case 2:
				go updateQuery(updates4Clycles)
			case 3:
				go deleteQuery(deletes4Clycles)
			case 4:
				go query(order4Cycles)
			}
		}

		if !nextLoopTest() {
			log.Info("Stop test")
			break
		}

		log.Debug("Sleep:: ", time_sleep)
		time.Sleep(time.Duration(time_sleep) * time.Second)
	}

}

// Check a target type and call function to validate this.
func nextLoopTest() bool {

	var r bool

	if viper.GetString("test.target_type") == "size" {
		if targetSizeReached() {
			log.Info("The target size reached! ")
			log.Info("Stop test")
			r = false
		} else {
			r = true
		}
	}

	if viper.GetString("test.target_type") == "orders" {
		if targetOrdersReached() {
			log.Info("The target orders reached! ")
			r = false
		} else {
			r = true
		}
	}

	return r
}

// Generate Orders Inserts and insert in single transaction
func insertOrder(orders4Cycle int, randmonData RandmonData) {
	log.Debug("Generate ", orders4Cycle, " Inserts ")
	var sql string
	sql = GenerateOrders(orders4Cycle, randmonData.customers, randmonData.employees, randmonData.products, randmonData.districts_warehouses)
	conn, ctx := wolfgres.PgxConnDB(datname)
	defer conn.Close(ctx)

	tx, err := conn.Begin(ctx)
	_, err = tx.Exec(ctx, sql)
	if err != nil {
		log.Error(err)
		tx.Rollback(ctx)
	}
	tx.Commit(ctx)
}

func updateQuery(n int) {
	log.Debug("Generate ", n, " Updates ")
	var sql string
	sql = GenerateUpdateOrder(n, datname)
	conn, ctx := wolfgres.PgxConnDB(datname)
	conn.Close(ctx)
	tx, err := conn.Begin(ctx)
	_, err = tx.Exec(ctx, sql)
	if err != nil {
		log.Error(err)
		tx.Rollback(ctx)
	}
	tx.Commit(ctx)
}

func deleteQuery(n int) {
	log.Debug("Generate ", n, " Deletes")
	// TODO: Generate Deletes Worker
	var sql string
	sql = GenerateDeleteOrder(n, datname)
	conn, ctx := wolfgres.PgxConnDB(datname)
	defer conn.Close(ctx)
	tx, err := conn.Begin(ctx)
	_, err = tx.Exec(ctx, sql)
	if err != nil {
		log.Error(err)
		tx.Rollback(ctx)
	}
	tx.Commit(ctx)

}

func query(n int) {
	log.Debug("Generate Query ")
	// TODO: Generate Query Worker
	var sql string
	sql = GenerateSelectsQueries(n, datname)
	conn, ctx := wolfgres.PgxConnDB(datname)
	defer conn.Close(ctx)
	tx, err := conn.Begin(ctx)
	_, err = tx.Exec(ctx, sql)
	if err != nil {
		log.Error(err)
	}
}

// This function check if the target size reached
func targetSizeReached() bool {
	target_size := viper.GetInt("test.target_size")
	log.Debug("Check Database Size :: ", datname)
	var size int
	conn, ctx := wolfgres.PgxConn()
	defer conn.Close(ctx)
	sql := fmt.Sprintf("SELECT pg_database_size('%s') / 1024 / 1024;", datname)
	log.Debug(sql)
	row := conn.QueryRow(ctx, sql)
	err := row.Scan(&size)

	if err != nil {
		log.Error(err)
	}

	log.Debug("Size DB Test :: ", size)
	log.Debug("Target Size :: ", target_size)

	if size <= target_size {
		return false
	} else {
		return true
	}
}

func targetOrdersReached() bool {
	targetOrders := viper.GetInt("test.target_orders")
	var noOrders int
	conn, ctx := wolfgres.PgxConn()
	defer conn.Close(ctx)
	sql := "SELECT count(order_id) FROM wfg.orders;"
	log.Debug(sql)
	row := conn.QueryRow(ctx, sql)
	err := row.Scan(&noOrders)

	if err != nil {
		log.Error(err)
	}

	log.Debug("Orders Test :: ", noOrders)
	log.Debug("Target Orders :: ", targetOrders)

	if noOrders <= targetOrders {
		return false
	} else {
		return true
	}
}

// When
func maxConnectionsReached() bool {
	max_conn_percent := viper.GetInt("test.max_conn_percent")
	var availableConnPercent int
	conn, ctx := wolfgres.PgxConn()
	defer conn.Close(ctx)
	sql := "SELECT  100 - ((t.conn_active * 100)::integer/m.max_connection::integer)  FROM (SELECT 1 AS id, setting AS max_connection FROM pg_settings WHERE name = 'max_connections') m JOIN (SELECT 1 AS id, COALESCE(COUNT(pid),0) AS conn_active FROM pg_stat_activity WHERE state = 'active' AND pid <> pg_backend_pid()) t ON t.id = m.id"
	row := conn.QueryRow(ctx, sql)
	err := row.Scan(&availableConnPercent)

	if err != nil {
		log.Error(err)
	}

	log.Debug(" max_conn_percent:: ", max_conn_percent, " > availableConnPercent:: ", availableConnPercent)
	log.Debug("result:: ", max_conn_percent > availableConnPercent)

	return max_conn_percent > availableConnPercent
}

/*
	Create Test Database set in configfile and Create Test User
*/
func createDBUser() (bool, error) {
	conn, ctx := wolfgres.PgxConn()
	defer conn.Close(ctx)
	test_user := viper.GetString("database.test_user")
	test_pass := viper.GetString("database.test_pass")
	log.Info("Create User:: ", test_user)
	sql := fmt.Sprintf("CREATE USER %s WITH PASSWORD '%s';\n", test_user, test_pass)
	log.Debug(sql)

	_, err := conn.Exec(ctx, sql)
	if err != nil {
		log.Error(err)
		//return false, err
	} else {
		log.Info("User test:: %s Created ", test_user)
	}

	log.Info("Create database:: ", datname)
	sql = fmt.Sprintf("CREATE DATABASE %s OWNER %s ;\n", datname, test_user)
	log.Debug("SQL Query to create Database -> ", sql)
	_, err = conn.Exec(ctx, sql)
	if err != nil {
		log.Error(err)
		return false, err
	} else {
		log.Info("Database %s Created ", datname)
	}

	return true, err
}

func loadCatalogs(randmonData RandmonData) (bool, error) {
	conn, ctx := wolfgres.PgxConnDB(datname)
	defer conn.Close(ctx)
	if loadScript(conn, ctx, "schema") {
		log.Info("Database Schema Created")
	}

	log.Info("ALTER DATABASE to set search_path")
	sql := fmt.Sprintf("ALTER DATABASE %s SET search_path TO wfg;\n", datname)
	log.Debug("SQL Query to alter database -> ", sql)

	_, err := conn.Exec(ctx, sql)
	if err != nil {
		log.Error(err)
		return false, err
	} else {
		log.Info("Alter Database set wfg ")
	}

	log.Info("Load catalogs")
	log.Info("*************")
	log.Info("Load category category")

	loadScript(conn, ctx, "category")
	log.Info("Load order status catalog")
	loadScript(conn, ctx, "order_status")

	log.Info("Load country country catalog")
	loadScript(conn, ctx, "country")

	log.Info("Load city city catalog")
	loadScript(conn, ctx, "city")

	log.Info("Create Region Catalog")

	regions, err := strconv.Atoi(fmt.Sprintf("%v", randmonData.regions))
	sql = GenerateRegion(regions)
	_, err = conn.Exec(ctx, sql)
	if err != nil {
		log.Error(err)
		return false, err
	}

	log.Info("Create Region Warehouses and Districts")
	sql = GenerateWarehouse(datname, 0, randmonData.warehouses, randmonData.districts)
	_, err = conn.Exec(ctx, sql)
	if err != nil {
		log.Error(err)
		return false, err
	}

	ds := randmonData.warehouses * randmonData.districts

	log.Info("Create Products Catalogo")
	sql = GenerateProducts(0, randmonData.products)
	_, err = conn.Exec(ctx, sql)
	if err != nil {
		log.Error(err)
		return false, err
	}

	log.Info("Generate Inventory per District")
	for i := 1; i <= ds; i++ {
		sql = GenerateInventory(i, randmonData.products)
		_, err = conn.Exec(ctx, sql)
		if err != nil {
			log.Error(err)
			return false, err
		}
	}

	log.Info("Create Employees Catalog")
	sql = GenerateEmployee(0, randmonData.employees)
	_, err = conn.Exec(ctx, sql)
	if err != nil {
		log.Error(err)
		return false, err
	}

	log.Info("Create Customers & Contact Customers Catalog")
	sql = GenerateCustomers(0, randmonData.customers)
	_, err = conn.Exec(ctx, sql)
	if err != nil {
		log.Error(err)
		return false, err
	}

	return true, err
}

// Validate database user, database and all tables in the schema before start test
func validateTestSchema() bool {
	var r bool
	conn, ctx := wolfgres.PgxConn()

	v := viper.GetString("database.test_user")
	s := "SELECT CASE WHEN EXISTS (SELECT usename FROM pg_user WHERE usename = '%s') THEN CAST(True AS BOOL) ELSE CAST(False AS BOOL) END;"
	sql := fmt.Sprintf(s, v)
	row, err := conn.Query(ctx, sql)
	if err != nil {
		log.Error(err)
	}
	for row.Next() {
		err = row.Scan(&r)
		log.Debug("The user exists -> ", r)
	}

	s = "SELECT CASE WHEN EXISTS (SELECT usename FROM pg_user WHERE usename = '%s') THEN CAST(True AS BOOL) ELSE CAST(False AS BOOL) END;"
	sql = fmt.Sprintf(s, datname)
	row, err = conn.Query(ctx, sql)
	if err != nil {
		log.Error(err)
	}
	for row.Next() {
		err = row.Scan(&r)
		log.Debug("The database exists -> ", r)
	}

	conn.Close(ctx)
	conn, ctx = wolfgres.PgxConnDB(datname)
	defer conn.Close(ctx)

	s = "SELECT CASE WHEN EXISTS ( SELECT datname FROM pg_database WHERE datname = '%s') THEN CAST(True AS BOOL) ELSE CAST(False AS BOOL) END;"
	sql = fmt.Sprintf(s, datname)
	row, err = conn.Query(ctx, sql)
	if err != nil {
		log.Error(err)
	}
	for row.Next() {
		err = row.Scan(&r)
		log.Debug("The database exists -> ", r)
	}

	return r
}

func loadScript(conn *pgx.Conn, ctx context.Context, scriptName string) bool {
	f := fmt.Sprintf("./scripts/%s.sql", scriptName)
	log.Debug("---> ", f)
	absPath, err := filepath.Abs(f)

	c, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Error(err)
		return false
	}

	sql := string(c)

	_, err = conn.Exec(ctx, sql)
	if err != nil {
		log.Error(err)
		return false
	}

	return true
}

// Load randmon products
func LoadProducts(max int) {
	var sql string
	log.Debug("Start load products")
	log.Debug("Load ", max, " products")
	conn, ctx := wolfgres.PgxConn()
	defer conn.Close(ctx)

	log.Debug("Get lastId")

	lastId, err := getLastId(ctx, conn, "employee_id", "wfg", "employee")
	if err != nil {
		log.Error(err)
	}

	log.Debug("lastId --> ", lastId)
	tx, err := conn.Begin(ctx)

	defer tx.Rollback(ctx)

	sql = GenerateProducts(lastId, 10)

	if err != nil {
		log.Error(err)
	}

	_, err = tx.Exec(ctx, sql)

	if err != nil {
		log.Error(err)
	}

	err = tx.Commit(ctx)

	if err != nil {
		log.Error(err)
	}

}

func LoadEmployees(max int) {
	var sql string
	log.Debug("Start load products")
	log.Debug("Load ", max, " products")
	conn, ctx := wolfgres.PgxConn()
	defer conn.Close(ctx)

	log.Debug("Get lastId")
	lastId, err := getLastId(ctx, conn, "employee_id", "wfg", "employee")
	if err != nil {
		log.Error(err)
	}

	tx, err := conn.Begin(ctx)
	defer tx.Rollback(ctx)

	sql = GenerateEmployee(lastId, max)
	if err != nil {
		log.Error(err)
	}
	_, err = tx.Exec(ctx, sql)

	if err != nil {
		log.Error(err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		log.Error(err)
	}
}

// Save a PID number in a file
func SavePID(pid int) {

	file, err := os.Create(viper.GetString("deamon.pid_file"))

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
