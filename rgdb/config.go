package rgdb

import "strconv"

type Config struct {
	ConnString string
	Host       string
	User       string
	Password   string
	Port       string
	Database   string
	SSLMode    bool
	MaxConns   int
}

func (c *Config) GetConnectionString() string {
	if c.ConnString != "" {
		return c.ConnString
	}

	str := `postgres://` + c.User

	if c.Password != "" {
		str = str + `:` + c.Password
	}

	str = str + `@` + c.Host + `:`

	if c.Port != "" {
		str = str + c.Port
	} else {
		str = str + `5432`
	}

	str = str + `/` + c.Database + `?`

	if !c.SSLMode {
		str = str + `sslmode=disable&`
	}

	if c.MaxConns != 0 {
		str = str + `max_conns=` + strconv.Itoa(c.MaxConns)
	}

	return str
}
