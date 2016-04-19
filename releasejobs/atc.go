package releasejobs
/*
* File Generated by enaml generator
*/
type Atc struct {

	OldResourceGracePeriod string `yaml:"old_resource_grace_period,omitempty"`

	GithubAuth_apiUrl interface{} `yaml:"github_auth.api_url,omitempty"`

	Retention_container_successDuration string `yaml:"retention.container.success_duration,omitempty"`

	PeerUrl interface{} `yaml:"peer_url,omitempty"`

	BasicAuthUsername string `yaml:"basic_auth_username,omitempty"`

	BasicAuthPassword string `yaml:"basic_auth_password,omitempty"`

	GithubAuth_tokenUrl interface{} `yaml:"github_auth.token_url,omitempty"`

	Yeller_apiKey string `yaml:"yeller.api_key,omitempty"`

	Riemann_port int `yaml:"riemann.port,omitempty"`

	ExternalUrl interface{} `yaml:"external_url,omitempty"`

	GithubAuth_clientId string `yaml:"github_auth.client_id,omitempty"`

	ResourceCacheCleanupInterval string `yaml:"resource_cache_cleanup_interval,omitempty"`

	BindIp string `yaml:"bind_ip,omitempty"`

	GithubAuth_authorize []interface {} `yaml:"github_auth.authorize,omitempty"`

	PostgresqlDatabase interface{} `yaml:"postgresql_database,omitempty"`

	Yeller_environmentName string `yaml:"yeller.environment_name,omitempty"`

	Riemann_host string `yaml:"riemann.host,omitempty"`

	DefaultCheckInterval string `yaml:"default_check_interval,omitempty"`

	BindPort int `yaml:"bind_port,omitempty"`

	Postgresql_role_name string `yaml:"postgresql.role.name,omitempty"`

	Postgresql_role_password interface{} `yaml:"postgresql.role.password,omitempty"`

	PubliclyViewable bool `yaml:"publicly_viewable,omitempty"`

	DevelopmentMode bool `yaml:"development_mode,omitempty"`

	GithubAuth_authUrl interface{} `yaml:"github_auth.auth_url,omitempty"`

	Postgresql_address interface{} `yaml:"postgresql.address,omitempty"`

	GithubAuth_clientSecret string `yaml:"github_auth.client_secret,omitempty"`

	Retention_container_failureDuration string `yaml:"retention.container.failure_duration,omitempty"`

	Postgresql_database string `yaml:"postgresql.database,omitempty"`

}