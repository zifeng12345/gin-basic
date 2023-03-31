package tables

type Waiting struct {
	Bases
	StoreId     string `json:"storeId"`
	UserName    string `json:"userName"`
	Desc        string `json:"desc"`
	PhoneNumber string `json:"phoneNumber"`
	Status      int32  `json:"status"`
	Number      int32  `json:"number"`
	Day         string `json:"day"`
	Pax         int32  `json:"pax"`
}
