package bkmkitap

const (
	loginEndpoint        = "https://www.bkmkitap.com/srv/customer/signin/email/"
	favoritesEndpoint    = "https://www.bkmkitap.com/srv/service/profile/get-shopping-list?link=uye-alisveris-listesi"
	getCountriesEndpoint = "https://www.bkmkitap.com/srv/service/region/get-countries"
	getBasketEndpoint    = "https://www.bkmkitap.com/srv/service/customer/get-basket"
	getOffersEndpoint    = "https://www.bkmkitap.com/srv/service/showcase/interactive"
)

type Favorites struct {
	Categories []Category `json:"CATEGORIES"`
	IsShare    int        `json:"IS_SHARE"`
}

type Image struct {
	ID     string `json:"ID"`
	Small  string `json:"SMALL"`
	Medium string `json:"MEDIUM"`
	Big    string `json:"BIG"`
	Title  string `json:"TITLE"`
}

type Product struct {
	ListID                 string `json:"LIST_ID"`
	DateAdded              string `json:"DATE_ADDED"`
	ID                     string `json:"ID"`
	URL                    string `json:"URL"`
	Image                  Image  `json:"IMAGE"`
	Brand                  string `json:"BRAND"`
	BrandURL               string `json:"BRAND_URL"`
	Title                  string `json:"TITLE"`
	PriceSell              string `json:"PRICE_SELL"`
	PriceBuy               string `json:"PRICE_BUY"`
	PriceNotDiscounted     int    `json:"PRICE_NOT_DISCOUNTED"`
	Vat                    int    `json:"VAT"`
	TargetCurrency         string `json:"TARGET_CURRENCY"`
	TargetCurrencyOriginal string `json:"TARGET_CURRENCY_ORIGINAL"`
	Currency               string `json:"CURRENCY"`
	HasVariant             int    `json:"HAS_VARIANT"`
}

type Category struct {
	ID       string    `json:"ID"`
	Name     string    `json:"NAME"`
	Products []Product `json:"PRODUCTS"`
}

type Region struct {
	Countries []Country `json:"countries"`
	Selected  string    `json:"selected"`
}

type Country struct {
	Title     string `json:"title"`
	Code      string `json:"code"`
	PhoneCode string `json:"phone_code"`
	Pattern   string `json:"pattern"`
	HasState  string `json:"has_state"`
}

type Basket struct {
	Currency     string          `json:"currency"`
	Mobile       int             `json:"mobile"`
	Price        float64         `json:"price"`
	PriceVat     int             `json:"price_vat"`
	PriceCargo   int             `json:"price_cargo"`
	MailHash     string          `json:"mail_hash"`
	CustomerCode string          `json:"customer_code"`
	Products     []BasketProduct `json:"products"`
}

type BasketProduct struct {
	Identifier   int     `json:"identifier"`
	Amount       float64 `json:"amount"`
	PriceSell    float64 `json:"price_sell"`
	Currency     string  `json:"currency"`
	Quantity     int     `json:"quantity"`
	Category     string  `json:"category"`
	CategoryID   int     `json:"category_id"`
	Code         int64   `json:"code"`
	Name         string  `json:"name"`
	Brand        string  `json:"brand"`
	SupplierCode string  `json:"supplier_code"`
	Note         string  `json:"note"`
}

type Bundle struct {
	Rows []Offer `json:"rows"`
}

type Offer struct {
	H string `json:"H"`
}
