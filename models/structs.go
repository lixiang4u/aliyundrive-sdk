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

type QueryFileListParams struct {
	All                   bool   `json:"all"`
	DriveId               string `json:"drive_id"`
	Fields                string `json:"fields"`
	ImageThumbnailProcess string `json:"image_thumbnail_process"`
	ImageUrlProcess       string `json:"image_url_process"`
	Limit                 int    `json:"limit"`
	OrderBy               string `json:"order_by"`
	OrderDirection        string `json:"order_direction"`
	ParentFileId          string `json:"parent_file_id"`
	UrlExpireSec          int    `json:"url_expire_sec"`
	VideoThumbnailProcess string `json:"video_thumbnail_process"`
}

type QueryFileSearchParams struct {
	DriveId               string `json:"drive_id"`
	ImageThumbnailProcess string `json:"image_thumbnail_process"`
	ImageUrlProcess       string `json:"image_url_process"`
	Limit                 int    `json:"limit"`
	OrderBy               string `json:"order_by"`
	Query                 string `json:"query"`
	VideoThumbnailProcess string `json:"video_thumbnail_process"`
}

type QueryFileDownloadParams struct {
	DriveId string `json:"drive_id"`
	FileId  string `json:"file_id"`
}
