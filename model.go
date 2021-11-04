package main

// Basic Info
type BasicInfoItem struct {
	KodeProduk      string
	SkuInduk        string
	NamaProduk      string
	DeskripsiProduk string
}

type BasicInfo struct {
	Items []BasicInfoItem
}

func (basicInfo *BasicInfo) AddItem(basicInfoItem BasicInfoItem) []BasicInfoItem {
	basicInfo.Items = append(basicInfo.Items, basicInfoItem)
	return basicInfo.Items
}

// Sales Info
type SalesInfoItem struct {
	KodeProduk  string
	NamaProduk  string
	KodeVariasi string
	NamaVariasi string
	SkuInduk    string
	Sku         string
	Harga       string
	Stok        string
}

type SalesInfo struct {
	Items []SalesInfoItem
}

func (salesInfo *SalesInfo) AddItem(salesInfoItem SalesInfoItem) []SalesInfoItem {
	salesInfo.Items = append(salesInfo.Items, salesInfoItem)
	return salesInfo.Items
}

// Shipping Info
type ShippingInfoItem struct {
	KodeProduk  string
	SkuInduk    string
	NamaProduk  string
	BeratProduk string
	Panjang     string
	Lebar       string
	Tinggi      string
}

type ShippingInfo struct {
	Items []ShippingInfoItem
}

func (shippingInfo *ShippingInfo) AddItem(shippingInfoItem ShippingInfoItem) []ShippingInfoItem {
	shippingInfo.Items = append(shippingInfo.Items, shippingInfoItem)
	return shippingInfo.Items
}

// Media Info
type MediaInfoItem struct {
	KodeProduk  string
	SkuInduk    string
	NamaProduk  string
	Kategori    string
	FotoSampul  string
	FotoProduk  []string
	NamaVariasi string
	Variasi     []string
	FotoVariasi []string
}

type MediaInfo struct {
	Items []MediaInfoItem
}

func (mediaInfo *MediaInfo) AddItem(mediaInfoItem MediaInfoItem) []MediaInfoItem {
	mediaInfo.Items = append(mediaInfo.Items, mediaInfoItem)
	return mediaInfo.Items
}

// Mass Upload
type MassUploadItem struct {
	Kategori             string
	NamaProduk           string
	DeskripsiProduk      string
	SkuInduk             string
	ProdukBerbahaya      string
	KodeIntegrasiVariasi string
	NamaVariasi1         string
	Varian1              string
	FotoVariasi1         string
	NamaVariasi2         string
	Varian2              string
	FotoVariasi2         string
	Harga                string
	Stok                 string
	KodeVariasi          string
	FotoSampul           string
	FotoProduk1          string
	FotoProduk2          string
	FotoProduk3          string
	FotoProduk4          string
	FotoProduk5          string
	FotoProduk6          string
	FotoProduk7          string
	FotoProduk8          string
	Berat                string
	Panjang              string
	Lebar                string
	Tinggi               string
}

type MassUpload struct {
	Items []MassUploadItem
}

func (massUpload *MassUpload) AddItem(massUploadItem MassUploadItem) []MassUploadItem {
	massUpload.Items = append(massUpload.Items, massUploadItem)
	return massUpload.Items
}
