package types

// DBOptions dummy
type DBOptions struct {
	Host        string
	Port        int
	SSLMode     string
	MaxIdleConn int
	Database    string
	User        string
	Password    string
}
