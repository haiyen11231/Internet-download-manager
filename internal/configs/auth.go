package configs

import "time"

type Hash struct {
	HashCost int `yaml:"hash_cost"`
}

type Token struct {
	ExpireIn                   string `yaml:"expire_in"`
	RegenerateTokenBeforeExpiry string `yaml:"regenerate_token_before_expiry"`
}

type Auth struct {
	Hash  Hash  
	Token Token 
}

func (t Token) GetExpireInDuration() (time.Duration, error) {
	return time.ParseDuration(t.ExpireIn)
}

func (t Token) GetRegenerateTokenBeforeExpiryDuration() (time.Duration, error) {
	return time.ParseDuration(t.RegenerateTokenBeforeExpiry)
}