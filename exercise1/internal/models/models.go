package models

type Product struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BrandName string `json:"brandName"`
	Details   string `json:"details"`
	ImageURL  string `json:"imageURL"`
}

type Variant struct {
	ID             string `json:"id"`
	ProductID      string `json:"productID"`
	VariantName    string `json:"variantName"`
	VariantDetails string `json:"variantDetails"`
}

type Products struct{
	Product Product
	Variant []Variant
}

