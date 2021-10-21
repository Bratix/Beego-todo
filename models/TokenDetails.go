package models

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type ExtractedTokenData struct {
	Uuid   string
	UserId int
}
