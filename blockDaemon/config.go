package blockDaemon

const baseUrl = "https://svc.blockdaemon.com/universal/v1"

type ConfigBlockDaemon struct {
	Protocol string
	Network  string
	Token    string
}

var (
	MainNet = ConfigBlockDaemon{
		Protocol: "bitcoin",
		Network:  "mainnet",
		Token:    "im4YrpAa9tjvFcwlZDci22aQGzp4JtAqnQtdzcMXAIdj-Aoi",
	}
	TestNet = ConfigBlockDaemon{
		Protocol: "bitcoin",
		Network:  "testnet",
		Token:    "zpka_b9d83e436ee54ec8b9089cf33760b37a_53d49665",
	}
)
