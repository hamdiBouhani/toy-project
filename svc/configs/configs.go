package configs

type Config struct {
	HostPort          string
	GRPCAddress       string
	GqlAddress        string
	Dsn               string
	CORSHosts         string
	GqlDebugUrlPrefix string
}
