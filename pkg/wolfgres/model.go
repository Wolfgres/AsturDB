package wolfgres

type WfgQueries struct {
	Name   string `yaml:"name"`
	Limit  int    `yaml:"limit"`
	Global bool   `yaml:"global"`
	Query  string `yaml:"query"`
}

type Queries struct {
	Queries []WfgQueries `yaml:"queries"`
}
