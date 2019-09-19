package internal

// ParamsConf describe the config for params
// purpose to change name of params which not is `client_id` or `client_secret`
// other providers not follow the oauth2 standard, such as Zalo Social API
type ParamsConf struct {

	// ClientID name of param client id, default is `client_id`
	ClientID string

	// ClientSecret name of param client secret, default is `client_secret`
	ClientSecret string
}

// InitializeIfEmpty initialize params config if empty
// default:
//    + ClientID     : `client_id`
//    + ClientSecret : `client_secret`
func (p *ParamsConf) InitializeIfEmpty() {
	if p.ClientID == "" {
		p.ClientID = "client_id"
	}
	if p.ClientSecret == "" {
		p.ClientSecret = "client_secret"
	}
}
