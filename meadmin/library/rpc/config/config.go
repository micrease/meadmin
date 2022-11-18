package config

type Config struct {
	Network        string   `yaml:"network"` //tcp,http
	Listen         string   `yaml:"listen"`  //ip:port
	BasePath       string   `yaml:"base_path"`
	ServiceName    string   `yaml:"service_name"`
	DiscoveryAddrs []string `yaml:"discovery_addrs"`
}
