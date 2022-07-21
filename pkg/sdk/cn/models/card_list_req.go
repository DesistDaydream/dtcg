package models

type FilterConditionReq struct {
	Page             string `json:"page"`
	Limit            string `json:"limit"`
	Name             string `json:"name"`
	State            string `json:"state"`
	CardGroup        string `json:"cardGroup"`
	RareDegree       string `json:"rareDegree"`
	BelongsType      string `json:"belongsType"`
	CardLevel        string `json:"cardLevel"`
	Form             string `json:"form"`
	Attribute        string `json:"attribute"`
	Type             string `json:"type"`
	Color            string `json:"color"`
	EnvolutionEffect string `json:"envolutionEffect"`
	SafeEffect       string `json:"safeEffect"`
	ParallCard       string `json:"parallCard"`
	KeyEffect        string `json:"keyEffect"`
}
