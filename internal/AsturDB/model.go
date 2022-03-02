package AsturDB

type RandmonData struct {
	customers            int
	employees            int
	products             int
	districts            int
	warehouses           int
	districts_warehouses int
	regions              int
}

// wfg.region
type Region struct {
	region_id   int
	region_name string
}

// wfg.region_country
type RegionCountry struct {
	region_id  int
	country_id int
}

// wfg.country
type Country struct {
	country_id      int
	country_name    string
	country_name_es string
	iso2            string
	iso3            string
	phone_code      string
}

// wfg.product
type Product struct {
	product_id    int
	barcode       string
	product_name  string
	description   string
	standard_cost float64
	list_price    float64
	category_id   int
}

// wfg.customer
type Customer struct {
	customer_id  int
	name         string
	address      string
	website      string
	credit_limit float64
}

// wfg.customer_contact
type CustomerContact struct {
	contact_id  int
	first_name  string
	last_name   string
	email       string
	phone       string
	customer_id int
}

// wfg.employee
type Employee struct {
	employee_id int
	first_name  string
	last_name   string
	email       string
	phone       string
	hire_date   string
	manager_id  int
	job_title   string
}

// wfg.warehouse
type Warehouse struct {
	warehouse_id   int
	warehouse_name string
	city_id        int
	address        string
	postal_code    string
}

// wfg.district
type District struct {
	district_id   int
	district_name string
	address       string
	postal_code   string
	warehouse_id  int
}

//wfg.district
type Inventory struct {
	district_id int
	product_id  int
	quantity    int
}

//wfg.district
type Orders struct {
	order_id    string
	customer_id int
	status_id   int
	salesman_id int
	order_date  string
	invoice_no  string
}

type OrderItems struct {
	order_id    string
	item_id     int
	product_id  int
	quantity    int
	unit_price  float64
	district_id int
}
