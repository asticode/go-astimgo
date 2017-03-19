package astimgo

import (
	"flag"
	"time"
)

// Flags
var (
	Addr     = flag.String("mgo-addr", "", "the mongo addr")
	Password = flag.String("mgo-password", "", "the mongo password")
	Timeout  = flag.Duration("mgo-timeout", 0, "the mongo timeout")
	Username = flag.String("mgo-username", "", "the mongo username")
)

// Configuration represents the AMQP configuration
type Configuration struct {
	Addr     string        `toml:"addr"`
	Password string        `toml:"password"`
	Timeout  time.Duration `toml:"timeout"`
	Username string        `toml:"username"`
}

// FlagConfig returns an AMQP config based on flags
func FlagConfig() Configuration {
	return Configuration{
		Addr:     *Addr,
		Password: *Password,
		Timeout:  *Timeout,
		Username: *Username,
	}
}
