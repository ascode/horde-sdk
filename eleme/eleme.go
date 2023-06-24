package eleme

type Conf struct {
	Sandbox      bool     `toml:"sandbox"`
	Key          string   `toml:"key"`
	Secret       string   `toml:"secret"`
	CallbackUrl  string   `toml:"callback_url"`
	TestAccounts []string `toml:"test_accounts"`
}

type User struct {
	Id              int64  `json:"userId"`
	Name            string `json:"userName"`
	AuthorizedShops []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"authorizedShops"`
}
