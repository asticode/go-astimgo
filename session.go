package astimgo

import (
	"gopkg.in/mgo.v2"
)

// NewSession returns a new session
func NewSession(c Configuration) (*mgo.Session, error) {
	return mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{c.Addr},
		Timeout:  c.Timeout,
		Username: c.Username,
		Password: c.Password,
	})
}
