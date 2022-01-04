package models

type LoginFormData struct {
	AppName     string `json:"appName"`
	AppEntrance string `json:"appEntrance"`
	CsrfToken   string `json:"_csrf_token"`
	UmIdToken   string `json:"umidToken"`
	IsMobile    bool   `json:"isMobile"`
	Lang        string `json:"lang"`
	ReturnUrl   string `json:"returnUrl"`
	Hsiz        string `json:"hsiz"`
	FromSite    int    `json:"fromSite"`
	BizParams   string `json:"bizParams"`
}

type LoginConfig struct {
	AppEntrance        string         `json:"appEntrance"`
	AppName            string         `json:"appName"`
	CurrentTime        string         `json:"currentTime"`
	NoCaptchaAppKey    string         `json:"nocaptchaAppKey"`
	UmIdEncryptAppName string         `json:"umidEncryptAppName"`
	UmIdToken          string         `json:"umidToken"`
	LoginForm          *LoginFormData `json:"loginFormData"`
}
