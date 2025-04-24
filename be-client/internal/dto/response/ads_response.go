package response

import "time"

type AdsComicResponse struct {
	A_Id          int64      `json:"ads_id"`
	A_Title       string     `json:"ads_title"`
	A_Image       string     `json:"ads_image"`
	A_Type        string     `json:"ads_type"`
	A_DirectURL   string     `json:"ads_direct_url"`
	A_ActiveFrom  *time.Time `json:"ads_active_from"`
	A_ActiveTo    *time.Time `json:"ads_active_to"`
	C_Id          int64      `json:"comic_id,omitempty"`
	C_Name        string     `json:"comic_name,omitempty"`
	C_Code        string     `json:"comic_code,omitempty"`
	C_Language    string     `json:"comic_language,omitempty"`
	C_Audience    string     `json:"comic_audience,omitempty"`
	C_Description string     `json:"comic_description,omitempty"`
	C_Cover       string     `json:"comic_cover,omitempty"`
	C_Active      bool       `json:"comic_active,omitempty"`
}
