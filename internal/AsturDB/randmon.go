package AsturDB

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func RandTimestamp() time.Time {
	randomTime := rand.Int63n(time.Now().Unix()-94608000) + 94608000
	randomNow := time.Unix(randomTime, 0)

	return randomNow
}

func RandSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n) // rune type is an alias of int32
	for i := range b {
		b[i] = letters[rand.Intn(len(letters)-1)+1]
	}
	return string(b)
}

/*
func RandProduct() string {
	rand.Seed(time.Now().UnixNano())
	s := material[rand.Intn(len(material))] + " " + products[rand.Intn(len(products))] + " " + colors[rand.Intn(len(colors))]
	return s
}*/

func RandEmail() string {
	rand.Seed(time.Now().UnixNano())
	s := strings.ToLower(RandSeq(10)) + "@" + strings.ToLower(RandSeq(10)) + ".com"
	return s
}

func RandBarcode(min int) string {
	rand.Seed(time.Now().UnixNano())
	var r string
	for i := 0; i < min; i++ {
		r += strconv.Itoa(rand.Intn(10))
	}
	return r
}

func RandFloat(min int, max int) float64 {
	rand.Seed(time.Now().UnixNano())
	r := (rand.Float64()*(float64(max)-float64(min)) + float64(min))
	return r
}

func RandInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	r := min + rand.Intn(max-min)
	return r
}

func RandCompanyWebsite(str string) string {
	rand.Seed(time.Now().UnixNano())
	str = strings.ToLower(strings.Join(strings.Fields(str), ""))
	return str + "." + domains[rand.Intn(len(domains))]
}

/*
func RandName() (string, string, string) {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(2)
	var g, fn, ln string
	g = gender[i]

	if i == 0 {
		fn = girlsNames[rand.Intn(len(girlsNames))]
	} else {
		fn = boysNames[rand.Intn(len(boysNames))]
	}

	ln = lastNames[rand.Intn((len(lastNames)))]
	log.Debug("Name -> ", g, " - ", fn, " - ", ln)

	return g, fn, ln
}

func RandCompany() string {
	rand.Seed(time.Now().UnixNano())
	return companies[rand.Intn(len(companies))] + " " + companiesTypes[rand.Intn(len(companiesTypes))]
}

func RandAddress() string {
	rand.Seed(time.Now().UnixNano())
	return streets[rand.Intn(len(streets))]
}
*/

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//var streets = []string{"S Bomar Ave", "College Dr", "W Penn Corner Rd", "Sable Way Trce", "N Pine St", "Cross Dr", "Bellevue Rd", "19th St N", "Gratiot Ave", "Ocean Blvd", "E 300th S", "Chippendale Ct", "Meadow Lake Dr", "Po Box", "W Pillar Rd", "Wales Dr", "Allensville Rd", "King St #2", "Preston Ave", "SW 57th Pl", "Morgan Ave", "Howard Hill Rd", "Bay Shore Ave", "101st St", "Peoga Rd", "Grey Eagle Ct #13", "Sherman St", "Copeland Dr", "Par Ave", "Valley Grove Rd", "Goodwin St", "Keen St #B", "E Jewett Pl", "Parker Ave", "State Ave", "Preston St", "NE 21st Way", "N 2nd St", "Cedar Rd E", "E Fort Rd #I1", "Lowell St", "183rd Ave", "State St #APT B-3", "North Ave", "Duke Dr", "Macarthur Ln", "Greers Ferry Rd", "County 775 Rd", "E Fm 273", "S Main St", "Twisted Hill Rd", "Willow Way #APT 27", "Post Oak Dr", "N Chestnut Dr", "556th Rd", "Maywood Dr", "Linden Ct #A", "Crestwood Dr", "Laura Ln", "11th St SE #205", "Sheeley St", "Tranquility Ct", "Smith Rd", "Hillcrest Ln", "Rowena Ave", "644th Hwy", "Willow Rd #DD", "Pent Rd", "Colony Rd", "Erica Ln", "S Norton Ave", "Ash St", "Melissa Dr", "Crest Hill Rd", "N State 130 Rte", "Augusta Dr", "Old Mill Dr", "E Bayshore Rd", "SW 8th Ct", "S 2050th W", "E Washington Ave", "Sweetwater Ln", "Eddy St", "Woodchip Rd", "Crisafulli Dr #50", "Nixon St", "Smith St", "County 1 Hwy", "Greenside Ter", "101st St", "Arbor View Ln", "A Ave", "S Cobb Dr", "Beaumont Ln", "Summerlee Ave", "Sumac Ave SW", "Alec Ln", "Bay Blvd #A", "W Pearl St", "Aiken Ave", "Bielby Rd", "48th Ave S", "Athens Rd", "Wheelock Rd", "Portola Rd", "Rose Cir", "W State 158 Hwy", "N Hillside Ave", "Ga 117 Hwy", "NW 166th St", "Lafayette Ln", "Fort Meade Rd #419", "Rr 1", "N Colmar Dr", "Dolly Rd", "W Vine St", "220 Speights St", "Wood St", "E Church St", "W Otmo Dr", "Little Bend Ct", "Lafayette Dr", "Warren St", "Creek Rd", "Wildwood Ave", "Twisting Trl", "Whisper Ridge Dr", "Azel Scott Ln", "Lexington Ct", "Byers St", "Chestnut St #11", "State 68 Hwy", "Longfield Rd", "Mesquite Ln", "N 15th Rd", "Fm 777", "49th Hwy E", "State Rd", "Private 1740 Rd", "Massachusetts St", "Jill Blvd"}
var domains = []string{"com", "net", "org", "io", "biz", "info", "co", "xyz", "online", "site", "me", "app", "solutions", "support", "global"}

//var girlsNames = []string{"Olivia", "Emma", "Ava", "Sophia", "Isabella", "Charlotte", "Amelia", "Mia", "Harper", "Evelyn", "Abigail", "Emily", "Ella", "Elizabeth", "Camila", "Luna", "Sofia", "Avery", "Mila", "Aria", "Scarlett", "Penelope", "Layla", "Chloe", "Victoria", "Madison", "Eleanor", "Grace", "Nora", "Riley", "Zoey", "Hannah", "Hazel", "Lily", "Ellie", "Violet", "Lillian", "Zoe", "Stella", "Aurora", "Natalie", "Emilia", "Everly", "Leah", "Aubrey", "Willow", "Addison", "Lucy", "Audrey", "Bella", "Nova", "Brooklyn", "Paisley", "Savannah", "Claire", "Skylar", "Isla", "Genesis", "Naomi", "Elena", "Caroline", "Eliana", "Anna", "Maya", "Valentina", "Ruby", "Kennedy", "Ivy", "Ariana", "Aaliyah", "Cora", "Madelyn", "Alice", "Kinsley", "Hailey", "Gabriella", "Allison", "Gianna", "Serenity", "Samantha", "Sarah", "Autumn", "Quinn", "Eva", "Piper", "Sophie", "Sadie", "Delilah", "Josephine", "Nevaeh", "Adeline", "Arya", "Emery", "Lydia", "Clara", "Vivian", "Madeline", "Peyton", "Julia", "Rylee", "Brielle", "Reagan", "Natalia", "Jade", "Athena", "Maria", "Leilani", "Everleigh", "Liliana", "Melanie", "Mackenzie", "Hadley", "Raelynn", "Kaylee", "Rose", "Arianna", "Isabelle", "Melody", "Eliza", "Lyla", "Katherine", "Aubree", "Adalynn", "Kylie", "Faith", "Mary", "Margaret", "Ximena", "Iris", "Alexandra", "Jasmine", "Charlie", "Amaya", "Taylor", "Isabel", "Ashley", "Khloe", "Ryleigh", "Alexa", "Amara", "Valeria", "Andrea", "Parker", "Norah", "Eden", "Elliana", "Brianna", "Emersyn", "Valerie", "Anastasia", "Eloise", "Emerson", "Cecilia", "Remi", "Josie", "Alina", "Reese", "Bailey", "Lucia", "Adalyn", "Molly", "Ayla", "Sara", "Daisy", "London", "Jordyn", "Esther", "Genevieve", "Harmony", "Annabelle", "Alyssa", "Ariel", "Aliyah", "Londyn", "Juliana", "Morgan", "Summer", "Juliette", "Trinity", "Callie", "Sienna", "Blakely", "Alaia", "Kayla", "Teagan", "Alaina", "Brynlee", "Finley", "Catalina", "Sloane", "Rachel", "Lilly", "Ember", "Kimberly", "Juniper", "Sydney", "Arabella", "Gemma", "Jocelyn", "Freya"}
//var boysNames = []string{"Liam", "Noah", "Oliver", "William", "Elijah", "James", "Benjamin", "Lucas", "Mason", "Ethan", "Alexander", "Henry", "Jacob", "Michael", "Daniel", "Logan", "Jackson", "Sebastian", "Jack", "Aiden", "Owen", "Samuel", "Matthew", "Joseph", "Levi", "Mateo", "David", "John", "Wyatt", "Carter", "Julian", "Luke", "Grayson", "Isaac", "Jayden", "Theodore", "Gabriel", "Anthony", "Dylan", "Leo", "Lincoln", "Jaxon", "Asher", "Christopher", "Josiah", "Andrew", "Thomas", "Joshua", "Ezra", "Hudson", "Charles", "Caleb", "Isaiah", "Ryan", "Nathan", "Adrian", "Christian", "Maverick", "Colton", "Elias", "Aaron", "Eli", "Landon", "Jonathan", "Nolan", "Hunter", "Cameron", "Connor", "Santiago", "Jeremiah", "Ezekiel", "Angel", "Roman", "Easton", "Miles", "Robert", "Jameson", "Nicholas", "Greyson", "Cooper", "Ian", "Carson", "Axel", "Jaxson", "Dominic", "Leonardo", "Luca", "Austin", "Jordan", "Adam", "Xavier", "Jose", "Jace", "Everett", "Declan", "Evan", "Kayden", "Parker", "Wesley", "Kai", "Brayden", "Bryson", "Weston", "Jason", "Emmett", "Sawyer", "Silas", "Bennett", "Brooks", "Micah", "Damian", "Harrison", "Waylon", "Ayden", "Vincent", "Ryder", "Kingston", "Rowan", "George", "Luis", "Chase", "Cole", "Nathaniel", "Zachary", "Ashton", "Braxton", "Gavin", "Tyler", "Diego", "Bentley", "Amir", "Beau", "Gael", "Carlos", "Ryker", "Jasper", "Max", "Juan", "Ivan", "Brandon", "Jonah", "Giovanni", "Kaiden", "Myles", "Calvin", "Lorenzo", "Maxwell", "Jayce", "Kevin", "Legend", "Tristan", "Jesus", "Jude", "Zion", "Justin", "Maddox", "Abel", "King", "Camden", "Elliott", "Malachi", "Milo", "Emmanuel", "Karter", "Rhett", "Alex", "August", "River", "Xander", "Antonio", "Brody", "Finn", "Elliot", "Dean", "Emiliano", "Eric", "Miguel", "Arthur", "Matteo", "Graham", "Alan", "Nicolas", "Blake", "Thiago", "Adriel", "Victor", "Joel", "Timothy", "Hayden", "Judah", "Abraham", "Edward", "Messiah", "Zayden", "Theo", "Tucker", "Grant", "Richard", "Alejandro", "Steven"}
//var lastNames = []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez", "Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson", "Thomas", "Taylor", "Moore", "Jackson", "Martin", "Lee", "Perez", "Thompson", "White", "Harris", "Sanchez", "Clark", "Ramirez", "Lewis", "Robinson", "Walker", "Young", "Allen", "King", "Wright", "Scott", "Torres", "Nguyen", "Hill", "Flores", "Green", "Adams", "Nelson", "Baker", "Hall", "Rivera", "Campbell", "Mitchell", "Carter", "Roberts", "Gomez", "Phillips", "Evans", "Turner", "Diaz", "Parker", "Cruz", "Edwards", "Collins", "Reyes", "Stewart", "Morris", "Morales", "Murphy", "Cook", "Rogers", "Gutierrez", "Ortiz", "Morgan", "Cooper", "Peterson", "Bailey", "Reed", "Kelly", "Howard", "Ramos", "Kim", "Cox", "Ward", "Richardson", "Watson", "Brooks", "Chavez", "Wood", "James", "Bennett", "Gray", "Mendoza", "Ruiz", "Hughes", "Price", "Alvarez", "Castillo", "Sanders", "Patel", "Myers", "Long", "Ross", "Foster", "Jimenez", "Powell", "Jenkins", "Perry", "Russell", "Sullivan", "Bell", "Coleman", "Butler", "Henderson", "Barnes", "Gonzales", "Fisher", "Vasquez", "Simmons", "Romero", "Jordan", "Patterson", "Alexander", "Hamilton", "Graham", "Reynolds", "Griffin", "Wallace", "Moreno", "West", "Cole", "Hayes", "Bryant", "Herrera", "Gibson", "Ellis", "Tran", "Medina", "Aguilar", "Stevens", "Murray", "Ford", "Castro", "Marshall", "Owens", "Harrison", "Fernandez", "Mcdonald", "Woods", "Washington", "Kennedy", "Wells", "Vargas", "Henry", "Chen", "Freeman", "Webb", "Tucker", "Guzman", "Burns", "Crawford", "Olson", "Simpson", "Porter", "Hunter", "Gordon", "Mendez", "Silva", "Shaw", "Snyder", "Mason", "Dixon", "Munoz", "Hunt", "Hicks", "Holmes", "Palmer", "Wagner", "Black", "Robertson", "Boyd", "Rose", "Stone", "Salazar", "Fox", "Warren", "Mills", "Meyer", "Rice", "Schmidt", "Garza", "Daniels", "Ferguson", "Nichols", "Stephens", "Soto", "Weaver", "Ryan", "Gardner", "Payne", "Grant", "Dunn", "Kelley", "Spencer", "Hawkins", "Arnold", "Pierce", "Vazquez", "Hansen", "Peters", "Santos", "Hart", "Bradley", "Knight", "Elliott", "Cunningham", "Duncan", "Armstrong", "Hudson", "Carroll", "Lane", "Riley", "Andrews", "Alvarado", "Ray", "Delgado", "Berry", "Perkins", "Hoffman", "Johnston", "Matthews", "Pena", "Richards", "Contreras", "Willis", "Carpenter", "Lawrence", "Sandoval", "Guerrero", "George", "Chapman", "Rios", "Estrada", "Ortega", "Watkins", "Greene", "Nunez", "Wheeler", "Valdez", "Harper", "Burke", "Larson", "Santiago", "Maldonado", "Morrison", "Franklin", "Carlson", "Austin", "Dominguez", "Carr", "Lawson", "Jacobs", "Obrien", "Lynch", "Singh", "Vega", "Bishop", "Montgomery", "Oliver", "Jensen", "Harvey", "Williamson", "Gilbert", "Dean", "Sims", "Espinoza", "Howell", "Li", "Wong", "Reid", "Hanson", "Le", "Mccoy", "Garrett", "Burton", "Fuller", "Wang", "Weber", "Welch", "Rojas", "Lucas", "Marquez", "Fields", "Park", "Yang", "Little", "Banks", "Padilla", "Day", "Walsh", "Bowman", "Schultz", "Luna", "Fowler", "Mejia", "Davidson", "Acosta", "Brewer", "May", "Holland", "Juarez", "Newman", "Pearson", "Curtis", "Cortez", "Douglas", "Schneider", "Joseph", "Barrett", "Navarro", "Figueroa", "Keller", "Avila", "Wade", "Molina", "Stanley", "Hopkins", "Campos", "Barnett", "Bates", "Chambers", "Caldwell", "Beck", "Lambert", "Miranda", "Byrd", "Craig", "Ayala", "Lowe", "Frazier", "Powers", "Neal", "Leonard", "Gregory", "Carrillo", "Sutton", "Fleming", "Rhodes", "Shelton", "Schwartz", "Norris", "Jennings", "Watts", "Duran", "Walters", "Cohen", "Mcdaniel", "Moran", "Parks", "Steele", "Vaughn", "Becker", "Holt", "Deleon", "Barker", "Terry", "Hale", "Leon", "Hail", "Benson", "Haynes", "Horton", "Miles", "Lyons", "Pham", "Graves", "Bush", "Thornton", "Wolfe", "Warner", "Cabrera", "Mckinney", "Mann", "Zimmerman", "Dawson", "Lara", "Fletcher", "Page", "Mccarthy", "Love", "Robles", "Cervantes", "Solis", "Erickson", "Reeves", "Chang", "Klein", "Salinas", "Fuentes", "Baldwin", "Daniel", "Simon", "Velasquez", "Hardy", "Higgins", "Aguirre", "Lin", "Cummings", "Chandler", "Sharp", "Barber", "Bowen", "Ochoa", "Dennis", "Robbins", "Liu", "Ramsey", "Francis", "Griffith", "Paul", "Blair", "Oconnor", "Cardenas", "Pacheco", "Cross", "Calderon", "Quinn", "Moss", "Swanson", "Chan", "Rivas", "Khan", "Rodgers", "Serrano", "Fitzgerald", "Rosales", "Stevenson", "Christensen", "Manning", "Gill", "Curry", "Mclaughlin", "Harmon", "Mcgee", "Gross", "Doyle", "Garner", "Newton", "Burgess", "Reese", "Walton", "Blake", "Trujillo", "Adkins", "Brady", "Goodman", "Roman", "Webster", "Goodwin", "Fischer", "Huang", "Potter", "Delacruz", "Montoya", "Todd", "Wu", "Hines", "Mullins", "Castaneda", "Malone", "Cannon", "Tate", "Mack", "Sherman", "Hubbard", "Hodges", "Zhang", "Guerra", "Wolf", "Valencia", "Franco", "Saunders", "Rowe", "Gallagher", "Farmer", "Hammond", "Hampton", "Townsend", "Ingram", "Wise", "Gallegos", "Clarke", "Barton", "Schroeder", "Maxwell", "Waters", "Logan", "Camacho", "Strickland", "Norman", "Person", "Colon", "Parsons", "Frank", "Harrington"}
//var products = []string{"chair", "stool", "table", "mug", "cup", "desk lamp", "floor lamp", "desk", "shelf", "sofa", "tea cup", "tea pot", "cutlery", "chess set", "lounge", "alarm clock", "phone dock", "keyboard", "side table", "wallet", "vase", "dog bed", "bird house", "wine holder", "skateboard", "calculator", "coathanger", "salt & pepper shaker", "coasters", "piggy bank", "headphones", "sculpture", "telephone", "flashlight", "mail sorter", "playing cards", "fan", "jewelry box", "mouse", "lantern", "walking cane", "sword", "wall clock", "mirror", "bed", "crib", "hammock", "plate", "bowl", "coffee mug", "espresso cup", "glasses", "fork", "spoon", "knife", "serving tray", "toy train", "action figure", "lamp shade", "cutting board", "dresser", "shoe rack", "rocking chair", "usb key", "8 ball", "frying pan", "drawer handle", "doorknob", "cable organizer", "planter pot", "coat hanger", "bottle opener", "can opener", "coasters", "pocket knife", "surfboard", "shoes", "book", "calendar", "house numbers", "spice rack", "suitcase", "button", "ring", "baking tray", "tape dispenser", "flower pot", "canoe", "basket", "pillow", "rug", "wall tile", "road bike", "bike seat", "handlebars"}
var material = []string{"sound-absorbing", "heavy", "weightless", "liquid", "reflective", "corrosive", "flexible", "malleable", "stiff", "flammable", "fabric", "styrofoam", "vinyl", "felt", "wool", "ceramic", "plastic", "aluminum", "steel", "cast iron", "digital", "paper", "cardboard", "maple", "walnut", "ceramic", "glass", "gold", "silver", "bronze", "brass", "copper", "carbon fiber", "cotton", "foam", "rubber", "silicone", "canvas", "concrete", "stone", "rope", "wire", "plywood", "hair", "brick", "leather", "suede", "3d Printed", "cork", "foam", "putty", "hemp", "bamboo", "dirt", "twig", "sheet metal", "stone", "cast iron", "wrought iron", "nickel", "tin", "zinc", "balsa wood", "oak", "teak", "bone", "ivory", "rattan", "diamond", "charcoal", "plaster", "clay", "feather", "floor", "ceiling", "wall", "corner", "papier mâché", "nylon", "mercury", "marble", "corian", "glowing", "floating", "magnetic", "wood fiber", "foil", "porcelain", "wax", "underwater", "outerspace", "outdoor"}
var colors = []string{"red", "blue", "green", "yellow", "purple", "pink", "orange", "brown"}
var gender = []string{"F", "M"}
var companiesTypes = []string{"Company", "Inc", "Group"}
