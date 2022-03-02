package wolfgres

type Config struct {
	Deamon struct {
		PidFile string `yaml:"pid_file"`
	}
	Log struct {
		LogLevel  string `yaml:"log_level"`
		LogFile   string `yaml:"log_file"`
		LogFormat string `yaml:"log_format"`
	}
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}
}

type Tanke struct {
	RandmonData struct {
		Regions    int `yaml:"regions"`
		Warehouses int `yaml:"warehouses"`
		Districts  int `yaml:"districts"`
		Products   int `yaml:"products"`
		Employees  int `yaml:"employees"`
		Customers  int `yaml:"customers"`
	}
	Target struct {
		Size            int `yaml:"size"`
		ConcurrentUsers struct {
			Min int `yaml:"min"`
			Max int `yaml:"max"`
		}
		ConcurrentTime struct {
			Min int `yaml:"min"`
			Max int `yaml:"max"`
		}
	}
}
