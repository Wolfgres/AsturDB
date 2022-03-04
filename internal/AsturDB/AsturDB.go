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
var test_user string

// Run Stress Test with name
// NameStressTest
func RunStressTest(dbTest string) {
	datname = dbTest
	test_user = viper.GetString("database.test_user")

	randomData := RandmonData{
		customers:            viper.GetInt("random_data.customers"),
		employees:            viper.GetInt("random_data.employees"),
		products:             viper.GetInt("random_data.products"),
		districts:            viper.GetInt("random_data.districts"),
		warehouses:           viper.GetInt("random_data.warehouses"),
		districts_warehouses: viper.GetInt("random_data.warehouses") * viper.GetInt("random_data.districts"),
		regions:              viper.GetInt("random_data.regions"),
	}

	createDBUser()
	loadCatalogs(randomData)

	log.Info("Init Orders")
	orders := viper.GetInt("test.orders_by_test")
	var orderXciclos int

	if viper.GetInt("test.orders_by_loop") > 50 {
		orderXciclos = 50
	} else {
		orderXciclos = viper.GetInt("test.orders_by_loop")
	}

	min_workers := viper.GetInt("test.min_workers")
	max_workers := viper.GetInt("test.max_workers")
	time_sleep := viper.GetInt("test.time_sleep")
	stress_type := viper.GetString("test.stress_type")

	var type_worker int

	for {

		workers := min_workers + rand.Intn(max_workers-min_workers)

		log.Debug("#########################")
		log.Debug("Workers:: -> ", workers)
		log.Debug("#########################")
		orders = orders - orderXciclos

		for i := 0; i < workers; i++ {

			if stress_type != "load" {
				type_worker = 1 + rand.Intn(5-1)
			} else {
				type_worker = 1
			}

			switch type_worker {
			case 1:
				go insertOrder(orderXciclos, randomData)
			case 2:
				go updateQuery(orderXciclos)
			case 3:
				go deleteQuery(orderXciclos)
			case 4:
				go query(orderXciclos)
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

func insertOrder(orderXciclos int, randmonData RandmonData) {
	log.Debug("Generate ", orderXciclos, " Inserts ")
	var sql string
	sql = GenerateOrders(orderXciclos, randmonData.customers, randmonData.employees, randmonData.products, randmonData.districts_warehouses)
	conn, ctx := wolfgres.PgxConnDB(datname, test_user)
	tx, err := conn.Begin(ctx)
	_, err = tx.Exec(ctx, sql)
	if err != nil {
		log.Error(err)
		tx.Rollback(ctx)
	}
	tx.Commit(ctx)
	conn.Close(ctx)
}

func updateQuery(n int) {
	log.Debug("Generate Updates ")
	// TODO: Generate Updates Worker
}

func deleteQuery(n int) {
	log.Debug("Generate Deletes ")
	// TODO: Generate Deletes Worker
}

func query(n int) {
	log.Debug("Generate Query ")
	// TODO: Generate Query Worker

}

// This function check if the target size reached
func targetSizeReached() bool {
	datname := "wolfgres_db" //tem
	target_size := viper.GetInt("test.target_size")
	log.Debug("Check Database Size :: ", datname)
	var size int
	conn, ctx := wolfgres.PgxConn()
	sql := fmt.Sprintf("SELECT pg_database_size('%s') / 1024 / 1024;", datname)
	log.Debug(sql)
	row := conn.QueryRow(ctx, sql)
	err := row.Scan(&size)

	if err != nil {
		log.Error(err)
	}

	conn.Close(ctx)

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
	sql := "SELECT count(order_id) FROM wfg.orders;"
	log.Debug(sql)
	row := conn.QueryRow(ctx, sql)
	err := row.Scan(&noOrders)

	if err != nil {
		log.Error(err)
	}

	conn.Close(ctx)

	log.Debug("Orders Test :: ", noOrders)
	log.Debug("Target Orders :: ", targetOrders)

	if noOrders <= targetOrders {
		return false
	} else {
		return true
	}
}

/*
	Create Test Database set in configfile and Create Test User
*/
func createDBUser() (bool, error) {
	conn, ctx := wolfgres.PgxConn()
	test_user := viper.GetString("database.test_user")
	log.Info("Create User:: ", test_user)
	sql := fmt.Sprintf("CREATE USER %s WITH PASSWORD '%s';\n", test_user, test_user)
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

	conn.Close(ctx)
	return true, err
}

func loadCatalogs(randmonData RandmonData) (bool, error) {
	conn, ctx := wolfgres.PgxConnDB(datname, test_user)
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
	conn.Close(ctx)

	return true, err
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
	config, err := wolfgres.GetConfig()

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
