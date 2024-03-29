package user

type InfoResponse struct {
	NameFirst       string `json:"name_first"`
	NameLast        string `json:"name_last"`
	NameFirstKana   string `json:"name_first_kana"`
	NameLastKana    string `json:"name_last_kana"`
	Type            int    `json:"type"`
	ProfileImageURL string `json:"profile_icon_url"`
	Email           string `json:"email"`
	Gender          int    `json:"gender"`
	StudentID       string `json:"student_id"`
	Status          int    `json:"status"`
}
