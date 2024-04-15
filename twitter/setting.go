package twitter

type Setting struct {
	// oauth 2.0
	ClientId     string `mapstructure:"clientId" yaml:"clientId"`
	ClientSecret string `mapstructure:"clientSecret" yaml:"clientSecret"`
	// oauth 1.0
	ApiKey            string `mapstructure:"apiKey" yaml:"apiKey"`
	ApiSecret         string `mapstructure:"apiSecret" yaml:"apiSecret"`
	AccessToken       string `mapstructure:"accessToken" yaml:"accessToken"`
	AccessTokenSecret string `mapstructure:"accessTokenSecret" yaml:"accessTokenSecret"`
	// app only access token
	BearerToken string `mapstructure:"bearerToken" yaml:"bearerToken"`
}
