package structs

type EnviromentModel struct {
	Database   database
	Cloudflare cloudflare
}

type database struct {
	Client      string
	MaxIdle     uint
	MaxLifeTime string
	MaxOpenConn uint
	User        string
	Password    string
	Host        string
	Db          string
	Params      string
	Port        string
}

type cloudflare struct {
	APIKey       string
	Email        string
	AccountID    string
	APIDomain    string
	APIVersion   string
	Pem          string
	UtilDomain   string
	KeyID        string
	StreamDomain string
}
