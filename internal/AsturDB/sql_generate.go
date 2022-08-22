package AsturDB

import (
	"AsturDB/pkg/wolfgres"
	"context"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/icrowley/fake"

	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

func GenerateOrders(x int, customers int, emps int, products int, district int) string {
	var sql string
	var order Orders
	var orderItems OrderItems

	for i := 0; i < x; i++ {
		order_id, _ := uuid.NewV4()

		order = Orders{
			order_id:    order_id.String(),
			customer_id: RandInt(1, customers),
			status_id:   RandInt(1, 5),
			salesman_id: RandInt(1, emps),
			order_date:  RandTimestamp().Format(time.RFC3339Nano),
			invoice_no:  fake.DigitsN(10),
		}

		sql += wolfgres.CreateQueryInsert("wfg", order) + "\n"

		items := rand.Intn(100)

		for j := 1; j <= items; j++ {
			q := rand.Intn(300)
			orderItems = OrderItems{
				order_id:    order_id.String(),
				item_id:     j,
				product_id:  RandInt(1, products),
				quantity:    q,
				unit_price:  RandFloat(15, 2000),
				district_id: RandInt(1, district),
			}

			sql += wolfgres.CreateQueryInsert("wfg", orderItems) + "\n"
		}
	}

	return sql

}

// Generate script with UPDATE senteces in order table
func GenerateUpdateOrder(x int, datname string) string {
	var err error
	query, err := wolfgres.GetQuery("get_randmon_order_ids")
	var updateScript string
	if err != nil {
		log.Error(err)
	}

	conn, ctx := wolfgres.PgxConnDB(datname)
	defer conn.Close(ctx)
	sqlQuery := fmt.Sprintf("%s LIMIT %s", query.Query, strconv.Itoa(x))
	rows, err := conn.Query(ctx, sqlQuery)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	updateScript = ""

	for rows.Next() {
		var order_id string

		err = rows.Scan(&order_id)

		if err != nil {
			log.Error("err")
			rows.Close()
		}

		statusId := strconv.Itoa(RandInt(1, 5))

		updateScript += fmt.Sprintf("UPDATE wfg.orders SET status_id = %s WHERE order_id = '%s';\n", statusId, order_id)
	}

	return updateScript
}

// Generate script with UPDATE senteces in order table
func GenerateDeleteOrder(x int, datname string) string {
	var err error
	query, err := wolfgres.GetQuery("get_randmon_order_ids")
	var updateScript string
	if err != nil {
		log.Error(err)
	}

	conn, ctx := wolfgres.PgxConnDB(datname)
	defer conn.Close(ctx)

	sqlQuery := fmt.Sprintf("%s LIMIT %s", query.Query, strconv.Itoa(x))
	rows, err := conn.Query(ctx, sqlQuery)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	updateScript = ""

	for rows.Next() {
		var order_id string

		err = rows.Scan(&order_id)

		if err != nil {
			log.Error("err")
			rows.Close()
		}

		updateScript += fmt.Sprintf("DELETE FROM wfg.order_items WHERE order_id = '%s';\n", order_id)
		updateScript += fmt.Sprintf("DELETE FROM wfg.orders WHERE order_id = '%s';\n", order_id)
	}

	return updateScript
}

//Generate queries
func GenerateSelectsQueries(x int, datname string) string {
	log.Debug(" x -> ", x)
	return ""
}

func GenerateRegion(regionNumber int) string {
	var sql, regionName string
	var region Region
	var regionCountry RegionCountry
	maxCountries := 248 / regionNumber

	var countryId []int

	for i := 1; i <= regionNumber; i++ {
		regionName = fmt.Sprintf("Region %v", i)
		region = Region{
			region_id:   i,
			region_name: regionName,
		}
		sql += wolfgres.CreateQueryInsert("wfg", region) + "\n"

		numCountries := rand.Intn(maxCountries)
		for j := 0; j < numCountries; j++ {
			r := rand.Intn(248)

			for wolfgres.FindInt(countryId, r) {
				r = rand.Intn(248)
			}
			countryId = append(countryId, r)

			regionCountry = RegionCountry{
				region_id:  i,
				country_id: r,
			}
			sql += wolfgres.CreateQueryInsert("wfg", regionCountry) + "\n"

		}
	}
	return sql
}

// Generate a Warehouses and districst queries
func GenerateWarehouse(datname string, lastId int, ws int, ds int) string {
	var sql string
	sql = ""
	citysIds := getCityIdsOnRegion(datname)
	//log.Debug("citysIds --> ", citysIds)

	var warehouse Warehouse
	var district District
	var district_id int

	district_id = 0
	for i := lastId; i < ws; i++ {
		idCity := citysIds[rand.Intn(len(citysIds))]
		id := (i + 1)
		warehouse = Warehouse{
			warehouse_id:   id,
			warehouse_name: fake.Brand(),
			city_id:        idCity,
			postal_code:    fake.Zip(),
			address:        fake.StreetAddress(),
		}

		sql += wolfgres.CreateQueryInsert("wfg", warehouse) + "\n"

		for j := 1; j <= ds; j++ {
			district_id += 1
			district = District{
				district_id:   district_id,
				district_name: fake.Brand(),
				address:       fake.StreetAddress(),
				postal_code:   fake.Zip(),
				warehouse_id:  id,
			}

			sql += wolfgres.CreateQueryInsert("wfg", district) + "\n"
		}
	}

	return sql
}

/*
func GenerateDistrict(lastId int, x int) string {
	var sql string
	sql = ""

	citysIds := getCityIdsOnRegion()

	var warehouse Warehouse

	for i := lastId; i <= x; i++ {
		idCity := citysIds[rand.Intn(len(citysIds))]
		warehouse = Warehouse{
			warehouse_id:   i,
			warehouse_name: fake.Brand(),
			city_id:        idCity,
			postal_code:    fake.Zip(),
			address:        fake.StreetAddress(),
		}

		sql += wolfgres.CreateQueryInsert("wfg", warehouse) + "\n"
	}

	return sql
}*/

func GenerateEmployee(lastId int, x int) string {
	var sql string

	var emp Employee
	for i := lastId; i <= x; i++ {
		hireDate := RandTimestamp().Format("2006-01-01 15:04:05")
		managerId := rand.Intn(x)
		id := (i + 1)
		emp = Employee{
			employee_id: id,
			first_name:  fake.FirstName(),
			last_name:   fake.LastName(),
			email:       fake.EmailAddress(),
			phone:       fake.Phone(),
			hire_date:   hireDate,
			manager_id:  managerId,
			job_title:   fake.JobTitle(),
		}

		sql += wolfgres.CreateQueryInsert("wfg", emp) + "\n"
	}

	return sql
}

func GenerateCustomers(lastId int, x int) string {
	var sql string
	sql = ""
	var customer Customer
	var customerContact CustomerContact

	for i := lastId; i < (lastId + x); i++ {
		companyName := fake.Company()
		id := (i + 1)
		customer = Customer{
			customer_id:  id,
			name:         companyName,
			address:      fake.StreetAddress(),
			website:      RandCompanyWebsite(companyName),
			credit_limit: RandFloat(1, 1000000),
		}

		sql += wolfgres.CreateQueryInsert("wfg", customer) + "\n"

		numContacts := rand.Intn(10)

		for j := 1; j <= numContacts; j++ {
			customerContact = CustomerContact{
				contact_id:  j,
				first_name:  fake.FirstName(),
				last_name:   fake.LastName(),
				email:       strings.ToLower(fake.EmailAddress()),
				phone:       fake.Phone(),
				customer_id: id,
			}

			sql += wolfgres.CreateQueryInsert("wfg", customerContact) + "\n"
		}
	}

	return sql
}

func GenerateProducts(lastId int, x int) string {
	var sql string
	sql = ""
	for i := lastId; i <= (lastId + x); i++ {
		id := (i + 1)
		min := RandInt(1, 100)
		max := RandInt(100, 1000)
		f := RandFloat(min, max)
		f = math.Round(f*100) / 100
		standard_cost := f
		list_price := standard_cost * float64(0.20)
		list_price = (math.Round(list_price*100) / 100)
		category_id := RandInt(1, 7)
		fake.Phone()
		product := Product{id, RandBarcode(10), fake.Product(), fake.CharactersN(255), standard_cost, list_price, category_id}
		sql += wolfgres.CreateQueryInsert("wfg", product) + "\n"
	}

	//log.Info(sql)

	return sql
}

func GenerateInventory(district_id int, products int) string {
	var sql string
	sql = ""
	for j := 1; j <= products; j++ {
		inventory := Inventory{
			district_id: district_id,
			product_id:  j,
			quantity:    RandInt(1, 100000),
		}

		sql += wolfgres.CreateQueryInsert("wfg", inventory) + "\n"
	}
	return sql
}

//TODO: Checar esta parte
func getCityIdsOnRegion(datname string) []int {
	conn, ctx := wolfgres.PgxConnDB(datname)
	defer conn.Close(ctx)

	sql := `SELECT city_id FROM city c
				JOIN country co ON co.country_id = c.country_id
 			WHERE c.country_id IN (SELECT country_id FROM region_country ORDER BY country_id);`
	rows, _ := conn.Query(ctx, sql)
	defer rows.Close()
	var citysIds []int

	var idCity int
	for rows.Next() {
		err := rows.Scan(&idCity)
		if err != nil {
			log.Error(err)
		}
		citysIds = append(citysIds, idCity)
	}

	return citysIds
}

// Get lastId from table in id
func getLastId(ctx context.Context, conn *pgx.Conn, id string, schema string, table string) (int, error) {
	var sql string

	log.Debug("Get lastId")

	sql = fmt.Sprintf("SELECT COALESCE(MAX(%s),0) FROM %s.%s", id, schema, table)

	var lastId int
	err := conn.QueryRow(ctx, sql).Scan(&lastId)
	if err != nil {
		log.Error(err)
	}

	return lastId, err
}
