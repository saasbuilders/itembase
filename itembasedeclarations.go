package itembase

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

// TODO: Some entities/models don't have the full set of fields from the API.
// Some of the implementation detail structs (Contacts, Billing, pagination
// containers, etc.) could perhaps be unexported.

type ProfileID string

func (profileID ProfileID) String() string {
	return string(profileID)
}

// A Profile represents a user profile entity from the itembase API.
//
// See http://sandbox.api.itembase.io/swagger-ui/
type Profile struct {
	Active    bool   `json:"active,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	Contact   struct {
		Contact []Contact `json:"contact,omitempty"`
	} `json:"contact,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	Currency          string     `json:"currency,omitempty"`
	DisplayName       string     `json:"display_name,omitempty"`
	ID                ProfileID  `json:"id"`
	Language          string     `json:"language,omitempty"`
	Locale            string     `json:"locale,omitempty"`
	OriginalReference string     `json:"original_reference,omitempty"`
	PlatformID        string     `json:"platform_id,omitempty"`
	PlatformName      string     `json:"platform_name,omitempty"`
	SourceID          string     `json:"source_id,omitempty"`
	Status            string     `json:"status,omitempty"`
	Type              string     `json:"type,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	URL               string     `json:"url,omitempty"`
}

// An Address represents a mailing address model from the itembase API.
type Address struct {
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	Line1   string `json:"line_1,omitempty"`
	Name    string `json:"name,omitempty"`
	Zip     string `json:"zip,omitempty"`
}

// A Contact represents a container of contact information from itembase API
// models.
type Contact struct {
	Addresses []Address `json:"addresses,omitempty"`
	Emails    []struct {
		Value string `json:"value,omitempty"`
	} `json:"emails,omitempty"`
	Phones []interface{} `json:"phones,omitempty"`
}

// GetName returns a string with a combined FirstName and
// LastName of a Buyer Profile
func (buyer *Buyer) GetName() string {
	return buyer.FirstName + " " + buyer.LastName
}

type BuyerID string

func (buyerID BuyerID) String() string {
	return string(buyerID)
}

// A Buyer represents a buyer entity from the itembase API.
//
// See http://sandbox.api.itembase.io/swagger-ui/
type Buyer struct {
	Active            bool       `json:"active,omitempty"`
	Contact           Contact    `json:"contact,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	Currency          string     `json:"currency,omitempty"`
	DateOfBirth       string     `json:"date_of_birth,omitempty"`
	FirstName         string     `json:"first_name,omitempty"`
	ID                BuyerID    `json:"id"`
	Language          string     `json:"language,omitempty"`
	LastName          string     `json:"last_name,omitempty"`
	Locale            string     `json:"locale,omitempty"`
	Note              string     `json:"note,omitempty"`
	OptOut            bool       `json:"opt_out,omitempty"`
	OriginalReference string     `json:"original_reference,omitempty"`
	SourceID          string     `json:"source_id,omitempty"`
	Status            string     `json:"status,omitempty"`
	Type              string     `json:"type,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	URL               string     `json:"url,omitempty"`
}

// GetEmail returns an Email for a Profile
func (buyer *Buyer) GetEmail() string {
	if len(buyer.Contact.Emails) > 0 {
		for _, EmailValue := range buyer.Contact.Emails {
			return EmailValue.Value
		}
	}

	return ""
}

// GetEmails returns all Emails for a Profile
func (buyer *Buyer) GetEmails() (emails []string) {
	if len(buyer.Contact.Emails) > 0 {
		for _, EmailValue := range buyer.Contact.Emails {
			emails = append(emails, EmailValue.Value)
		}
	}

	return
}

// A Category represents a product category model from the itembase API.
type Category struct {
	CategoryID string `json:"category_id,omitempty"`
	Language   string `json:"language,omitempty"`
	Value      string `json:"value,omitempty"`
}

// A ProductDescription represents a product description model from the itembase
// API, which may be in a specified language.
type ProductDescription struct {
	Language string `json:"language,omitempty"`
	Value    string `json:"value,omitempty"`
}

// A Brand represents a product brand model from the itembase API.
type Brand struct {
	Name struct {
		Language string `json:"language,omitempty"`
		Value    string `json:"value,omitempty"`
	} `json:"name,omitempty"`
}

type Identifier struct {
	ID string `json:"id,omitempty"`
}

type StockInformation struct {
	InStock        bool    `json:"in_stock,omitempty"`
	InventoryLevel float64 `json:"inventory_level,omitempty"`
	InventoryUnit  string  `json:"inventory_unit,omitempty"`
}

type ProductID string

func (productID ProductID) String() string {
	return string(productID)
}

// A Product represents a product entity from the itembase API.
//
// See http://sandbox.api.itembase.io/swagger-ui/
type Product struct {
	Active      bool                 `json:"active,omitempty"`
	Brand       Brand                `json:"brand,omitempty"`
	Categories  []Category           `json:"categories,omitempty"`
	Condition   string               `json:"condition,omitempty"`
	CreatedAt   *time.Time           `json:"created_at,omitempty"`
	Currency    string               `json:"currency,omitempty"`
	Description []ProductDescription `json:"description,omitempty"`
	ID          ProductID            `json:"id"`
	Identifier  Identifier           `json:"identifier,omitempty"`
	Name        []struct {
		Language string `json:"language,omitempty"`
		Value    string `json:"value,omitempty"`
	} `json:"name,omitempty"`
	OriginalReference string `json:"original_reference,omitempty"`
	PictureUrls       []struct {
		URLOriginal string `json:"url_original,omitempty"`
	} `json:"picture_urls,omitempty"`
	PricePerUnit float64 `json:"price_per_unit,omitempty"`
	Shipping     []struct {
		Price           float64 `json:"price,omitempty"`
		ShippingService string  `json:"shipping_service,omitempty"`
	} `json:"shipping,omitempty"`
	SourceID         string           `json:"source_id,omitempty"`
	StockInformation StockInformation `json:"stock_information,omitempty"`
	Tax              float64          `json:"tax,omitempty"`
	TaxRate          float64          `json:"tax_rate,omitempty"`
	UpdatedAt        *time.Time       `json:"updated_at,omitempty"`
	URL              string           `json:"url,omitempty"`
	Variants         []interface{}    `json:"variants,omitempty"`
}

func (product *Product) InStock() bool {
	return product.StockInformation.InStock
}

// Returns name for specified preferred language if present
func (product *Product) GetName(preferredLanguage string) (name string, ok bool) {

	for _, productName := range product.Name {
		if preferredLanguage == productName.Language {
			return cleanItembaseUnicode(productName.Value), true
		}
	}

	// if []struct{} is empty, return empty string
	return "", false
}

// Returns any name for Product
func (product *Product) GetDefaultName() (name string, ok bool) {

	for _, productName := range product.Name {
		return cleanItembaseUnicode(productName.Value), true
	}

	return "", false

}

func cleanItembaseUnicode(str string) string {
	str = strings.Replace(str, "\u00a0", " ", -1)
	str = strings.Replace(str, "\ufeff", "", -1)
	return str
}

// Billing represents a model from the itembase API containing the billing
// address of a Transaction.
type Billing struct {
	Address Address `json:"address,omitempty"`
}

type Shipping struct {
	Address Address `json:"address,omitempty"`
}

// Status describes a transactions' status
type Status struct {
	Global   string `json:"global,omitempty"`
	Payment  string `json:"payment,omitempty"`
	Shipping string `json:"shipping,omitempty"`
}

type TransactionID string

func (transactionID TransactionID) String() string {
	return string(transactionID)
}

// A Transaction represents a transaction entity from the itembase API.
//
// See http://sandbox.api.itembase.io/swagger-ui/
type Transaction struct {
	Billing           Billing       `json:"billing,omitempty"`
	Buyer             Buyer         `json:"buyer,omitempty"`
	CreatedAt         *time.Time    `json:"created_at,omitempty"`
	Currency          string        `json:"currency,omitempty"`
	ID                TransactionID `json:"id"`
	OriginalReference string        `json:"original_reference,omitempty"`
	Products          []Product     `json:"products,omitempty"`
	Shipping          Shipping      `json:"shipping,omitempty"`
	SourceID          string        `json:"source_id,omitempty"`
	Status            Status        `json:"status,omitempty"`
	TotalPrice        float64       `json:"total_price,omitempty"`
	TotalPriceNet     float64       `json:"total_price_net,omitempty"`
	TotalTax          float64       `json:"total_tax,omitempty"`
	UpdatedAt         *time.Time    `json:"updated_at,omitempty"`
}

func (t *Transaction) Completed() bool {
	if t.Status.Global == "completed" {
		return true
	}
	return false
}

// ItembaseResponse is a container for any Itembase response.
// It returns the resultset, Number of found documents and Number of documents returned
type ItembaseResponse struct {
	Documents            []interface{} `json:"documents"`
	NumDocumentsFound    int           `json:"num_documents_found"`
	NumDocumentsReturned int           `json:"num_documents_returned"`
}

// Transactions is a container for pagination of Transaction entities.
type Transactions struct {
	Transactions []Transaction `json:"documents"`
}

func (transactions *Transactions) Add(transaction interface{}) error {

	var newTransaction Transaction
	err := ConvertTo(transaction, &newTransaction)
	if err != nil {
		log.Println("Error when converting", "error", err)
		return err
	}
	if !transactions.Exists(newTransaction) {
		transactions.Transactions = append(transactions.Transactions, newTransaction)
	}
	return nil

}

func (transactions *Transactions) Exists(searchTransaction Transaction) bool {

	for _, transaction := range transactions.Transactions {
		if searchTransaction.ID == transaction.ID {
			log.Println("Transaction already exists in collection")
			return true
		}
	}
	return false
}

func (transactions *Transactions) Count() int {
	return len(transactions.Transactions)
}

// Return date of heighest Created At transaction
func (transactions *Transactions) MaxCreatedAt() time.Time {
	var maxCreatedAt time.Time

	for _, transaction := range transactions.Transactions {
		if transaction.CreatedAt.After(maxCreatedAt) {
			maxCreatedAt = *transaction.CreatedAt
		}
	}

	return maxCreatedAt
}

// Return date of heighest Updated At transaction
func (transactions *Transactions) MaxUpdatedAt() time.Time {
	var maxUpdatedAt time.Time

	for _, transaction := range transactions.Transactions {
		if transaction.UpdatedAt != nil && transaction.UpdatedAt.After(maxUpdatedAt) {
			maxUpdatedAt = *transaction.UpdatedAt
		}
	}

	return maxUpdatedAt
}

// Return only completed transactions
func (transactions *Transactions) Completed() (filteredTransactions Transactions) {
	for _, transaction := range transactions.Transactions {
		if transaction.Completed() {
			filteredTransactions.Add(transaction)
		}
	}

	return
}

// Profiles is a container for pagination of Profile entities.
type Profiles struct {
	Profiles []Profile `json:"documents"`
}

func (profiles *Profiles) Add(profile interface{}) error {

	var newProfile Profile
	err := ConvertTo(profile, &newProfile)
	if err != nil {
		log.Println("Error when converting", "error", err)
		return err
	}
	if !profiles.Exists(newProfile) {
		profiles.Profiles = append(profiles.Profiles, newProfile)
	}
	return nil

}

func (profiles *Profiles) Exists(searchProfile Profile) bool {

	for _, profile := range profiles.Profiles {
		if searchProfile.ID == profile.ID {
			log.Println("Profile already exists in collection")
			return true
		}
	}
	return false
}

func (profiles *Profiles) Count() int {
	return len(profiles.Profiles)
}

// Return date of heighest Created At profile
func (profiles *Profiles) MaxCreatedAt() time.Time {
	var maxCreatedAt time.Time

	for _, profile := range profiles.Profiles {
		if profile.CreatedAt.After(maxCreatedAt) {
			maxCreatedAt = *profile.CreatedAt
		}
	}

	return maxCreatedAt
}

// Return date of heighest Updated At profile
func (profiles *Profiles) MaxUpdatedAt() time.Time {
	var maxUpdatedAt time.Time

	for _, profile := range profiles.Profiles {
		if profile.UpdatedAt != nil && profile.UpdatedAt.After(maxUpdatedAt) {
			maxUpdatedAt = *profile.UpdatedAt
		}
	}

	return maxUpdatedAt
}

// Products is a container for pagination of Product entities.
type Products struct {
	Products []Product `json:"documents"`
}

func (products *Products) Add(product interface{}) error {

	var newProduct Product
	err := ConvertTo(product, &newProduct)
	if err != nil {
		log.Println("Error when converting", "error", err)
		return err
	}
	if !products.Exists(newProduct) {
		products.Products = append(products.Products, newProduct)
	}
	return nil

}

func (products *Products) Exists(searchProduct Product) bool {

	for _, product := range products.Products {
		if searchProduct.ID == product.ID {
			log.Println("Product already exists in collection", searchProduct.ID, product.ID)
			return true
		}
	}
	return false
}

func (products *Products) Count() int {
	return len(products.Products)
}

// Return date of heighest Created At product
func (products *Products) MaxCreatedAt() time.Time {
	var maxCreatedAt time.Time

	for _, product := range products.Products {
		if product.CreatedAt.After(maxCreatedAt) {
			maxCreatedAt = *product.CreatedAt
		}
	}

	return maxCreatedAt
}

// Return date of heighest Updated At product
func (products *Products) MaxUpdatedAt() time.Time {
	var maxUpdatedAt time.Time

	for _, product := range products.Products {
		if product.UpdatedAt != nil && product.UpdatedAt.After(maxUpdatedAt) {
			maxUpdatedAt = *product.UpdatedAt
		}
	}

	return maxUpdatedAt
}

func (products *Products) InStock() (filteredProducts Products) {
	for _, product := range products.Products {
		if product.InStock() {
			filteredProducts.Add(product)
		}
	}
	return
}

// Get Products based on shopID
func (products *Products) ByShop(shopID string) (filteredProducts Products) {
	for _, product := range products.Products {
		if product.SourceID == shopID {
			filteredProducts.Add(product)
		}
	}
	return
}

// Buyers is a container for pagination of Buyer entities.
type Buyers struct {
	Buyers []Buyer `json:"documents"`
}

func (buyers *Buyers) Add(buyer interface{}) error {

	var newBuyer Buyer
	err := ConvertTo(buyer, &newBuyer)
	if err != nil {
		log.Println("Error when converting", "error", err)
		return err
	}
	if !buyers.Exists(newBuyer) {
		buyers.Buyers = append(buyers.Buyers, newBuyer)
	}
	return nil

}

func (buyers *Buyers) Exists(searchBuyer Buyer) bool {

	for _, buyer := range buyers.Buyers {
		if searchBuyer.ID == buyer.ID {
			log.Println("Buyer already exists in collection")
			return true
		}
	}
	return false
}

func (buyers *Buyers) Count() int {
	return len(buyers.Buyers)
}

// Return date of heighest Created At buyer
func (buyers *Buyers) MaxCreatedAt() time.Time {
	var maxCreatedAt time.Time

	for _, buyer := range buyers.Buyers {
		if buyer.CreatedAt.After(maxCreatedAt) {
			maxCreatedAt = *buyer.CreatedAt
		}
	}

	return maxCreatedAt
}

// Return date of heighest Updated At buyer
func (buyers *Buyers) MaxUpdatedAt() time.Time {
	var maxUpdatedAt time.Time

	for _, buyer := range buyers.Buyers {
		if buyer.UpdatedAt != nil && buyer.UpdatedAt.After(maxUpdatedAt) {
			maxUpdatedAt = *buyer.UpdatedAt
		}
	}

	return maxUpdatedAt
}

func (buyers *Buyers) ByShop(shopID string) (filteredBuyers Buyers) {

	for _, buyer := range buyers.Buyers {
		if buyer.SourceID == shopID {
			filteredBuyers.Add(buyer)
		}
	}
	return
}

// A User represents a user entity from the itembase API, such as returned from
// the "me" endpoint.
type User struct {
	UUID              string `json:"uuid"`
	Username          string `json:"username,omitempty"`
	FirstName         string `json:"first_name,omitempty"`
	LastName          string `json:"last_name,omitempty"`
	MiddleName        string `json:"middle_name,omitempty"`
	NameFormat        string `json:"name_format,omitempty"`
	Locale            string `json:"locale,omitempty"`
	Email             string `json:"email,omitempty"`
	PreferredCurrency string `json:"preferred_currency,omitempty"`
}

func ConvertTo(inputInterface, outputType interface{}) error {

	jsonBLOB, err := json.Marshal(inputInterface)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = json.Unmarshal(jsonBLOB, &outputType)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil

}
