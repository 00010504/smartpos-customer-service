package config

const (
	DateTimeFormat = "2006-01-02 15:04:05"
	DateFormat     = "2006-01-02"

	ConsumerGroupID = "catalog_service_v2"

	PendingShipmentStatusID   = "66034c5b-48a5-4b41-9959-b69da8e02d33"
	CancelledShipmentStatusID = "acc2a42b-d0b8-45d1-9509-72c4a196319b"
	FinishedShipmentStatusID  = "cc6574f9-9ed5-409a-9081-2b1adcf0a85f"

	NewProductConditionID    = "b107ff89-4cfc-4c6b-a838-d97ac1333307"
	UsedProductConditionID   = "ca1ca924-e464-4849-a0ad-565cdf9d193c"
	DefectProductConditionID = "96ef24a8-efd7-4378-85ce-76458e2aeeac"

	ProductProductTypeID = "69e939aa-9b8f-46a9-b605-8b2675475b7b"
	ServiceProductTypeID = "5a0e556a-15f8-47ac-ae07-46972f3c6ab4"
	SetProductTypeID     = "864c77c7-5407-45dc-8289-3162b71dc653"

	ShipmentAdjustmentTypeID = "50bc6ff6-a185-4f4d-a6f9-fd289f5c68d2"
	ShipmentResidueTypeID    = "fd152773-2e12-4c1a-8fb5-a7d5c9955750"
	ShipmentGoodsTypeID      = "a230b02b-46f8-42f4-885e-d81813c297d6"

	ShipmentLoadedStatusID   = "7cbb2295-559e-4b72-a3c1-11ac24dffc6b"
	ShipmentCheckingStatusID = "c80f46d1-da4f-407b-8ad0-51feb6780312"
	ShipmentFinishedStatusID = "31cd30a7-46ae-460c-9530-7c2df1356b62"

	RepricingTypeCurrencyChange     = "currency_change"
	RepricingTypeProfitMarginChange = "profit_margin_change"
	RepricingTypePriceChange        = "price_change"

	RepricingNewStatus       = "c171e477-beb7-4146-8c64-1c5ecd862cbc"
	RepricingOnProcessStatus = "6d3dc7b6-a85b-4882-9cd6-353f0d82bf03"
	RepricingFinishedStatus  = "aed69f3b-9f76-4324-8836-93f5a0cca063"

	// Supplier order retail price change type
	ApplyOldRetailPrice = "apply_old_retail_price"
	ApplyDefaultMargin  = "apply_default_margin"

	// Stocktaking type
	StocktakingTypeFull     = "full"
	StocktakingTypePartial  = "partial"
	StocktakingTypeImport   = "import"
	StocktakingTypeTransfer = "transfer"

	// Table types
	ProductsTableType = "products_table"

	DeleteProductJobStatusNotStarted = -1
	UpdateProductJobStatusNotStarted = 0
	UpdateProductJobStatusInProgress = 1
	UpdateProductJobStatusFinished   = 2

	SupplierOrderEmptyID = "00000000-0000-0000-0000-000000000000"
)

var (
	ErrorMessages = map[string]map[string]string{
		"NAME":                    {"en": "The value of field NAME is required", "uz": "NAME maydonining qiymati talab qilinadi", "ru": "Значение поля Наименование является обязательным"},
		"SCANNED":                 {"en": "The value of field SCANNED is required", "uz": "SCANNED maydonining qiymati talab qilinadi", "ru": "Значение поля отсканировано является обязательным"},
		"SKU":                     {"en": "The value of field SKU is required", "uz": "SKU maydonining qiymati talab qilinadi", "ru": "Значение поля Артикул является обязательным"},
		"BARCODE":                 {"en": "The value of field barcode is required", "uz": "BARCODE maydonining qiymati talab qilinadi", "ru": "Значение поля Баркод является обязательным"},
		"SUPPLY_PRICE_REQUIRED":   {"en": "This value of field SUPPLY_PRICE is required", "uz": "SUPPLY_PRICE maydonining qiymati talab qilinadi", "ru": "Значение поля Цена поставки является обязательным"},
		"SUPPLY_PRICE_IS_NUMERIC": {"en": "The value of field SUPPLY_PRICE should be numeric", "uz": "SUPPLY_PRICE maydonining qiymati raqamli bo'lishi kerak", "ru": "Значение поля Цена поставки должно быть числовым"},
		"RETAIL_PRICE_REQUIRED":   {"en": "The value of field RETAIL_PRICE is required", "uz": "RETAIL_PRICE maydonining qiymati talab qilinadi", "ru": "Значение поля Цена продажи является обязательным"},
		"RETAIL_PRICE_IS_NUMERIC": {"en": "The value of field RETAIL_PRICE should be numeric", "uz": "RETAIL_PRICE maydonining qiymati raqamli bo'lishi kerak", "ru": "Значение поля Цена продажи должно быть числовым"},
		"MEASUREMENT_UNIT":        {"en": "The value of field MEASUREMENT_UNIT does not exists", "uz": "O'lchov birligi maydonining qiymati mavjud emas", "ru": "Значение поля ЕДИНИЦА ИЗМЕРЕНИЯ не существует"},
		"QUANTITY_REQUIRED":       {"en": "The value of field QUANTITY is required", "uz": "QUANTITY maydonining qiymati talab qilinadi", "ru": "Значение поля КОЛИЧЕСТВО является обязательным"},
		"QUANTITY_IS_NUMERIC":     {"en": "The value of field QUANTITY should be numberic", "uz": "QUANTITY raqam bo'lish kere", "ru": "Значение поля КОЛИЧЕСТВО должно быть числовым"},
		"PROPERTY":                {"en": "Required property %v is not provided", "uz": "Kerakli mulk %v ko'rsatilmagan", "ru": "Требуемое свойство %v не указано"},
		"PRICE_REQUIRED":          {"en": "The value of field PRICE is required", "uz": "PRICE maydonining qiymati talab qilinadi", "ru": "Значение поля ЦЕНА является обязательным"},
		"PRICE_IS_NUMERIC":        {"en": "The value of field PRICE should be numberic", "uz": "PRICE raqam bo'lish kere", "ru": "Значение поля ЦЕНА должно быть числовым"},
		"PRODUCT":                 {"en": "Product does not exist: Barcode:%v  SKU:%v", "uz": "Mahsulot mavjud emas : Barcode:%v  SKU:%v ", "ru": "Товара не существует : Баркод:%v  Артикул:%v"},
		"PRODUCT_IS_AVAILABLE":    {"en": "Product is not available in this shop", "uz": "Bu do'konda mahsulot yo'q", "ru": "Товар недоступен в этом магазине"},
		"DUPLICATE_BARCODE_SKU":   {"en": "Duplicate barcode or sku: %v - %v", "uz": "Barcode yoki sku kiritilgan: %v - %v", "ru": "Дублированный штрихкод или артикул: %v - %v"},
		"WRONG_PRODUCT_TYPE_ID":   {"en": "Only products are available for orders. You cannot order a service or package", "uz": "Buyurtma uchun faqat tovarlar mavjud. Xizmat yoki komplektni buyurtma berish mumkin emas.", "ru": "Для заказа доступны только товары. Нельзя заказать услугу или комплект."},
	}

	Properties = map[string]map[string]string{
		"BARCODE":          {"en": "Barcode", "uz": "Barcode", "ru": "Штрихкод"},
		"SKU":              {"en": "SKU", "uz": "SKU", "ru": "Артикул"},
		"NAME":             {"en": "Name", "uz": "Nomi", "ru": "Наименование"},
		"DESCRIPTION":      {"en": "Description", "uz": "Ta'rif", "ru": "Описание"},
		"MEASUREMENT_UNIT": {"en": "Measurement unit", "uz": "O'lchov birligi", "ru": "Единица измерения"},
		"QUANTITY":         {"en": "Quantity", "uz": "Soni", "ru": "Количество"},
		"PRICE":            {"en": "Price", "uz": "Qiymat", "ru": "Цена"},
		"RETAIL_PRICE":     {"en": "Retail price", "uz": "Retail qiymat", "ru": "Розничная цена"},
		"PRODUCT_GROUP":    {"en": "Product group", "uz": "Mahsulot turi", "ru": "Тип товара"},
		"SUPPLY_PRICE":     {"en": "Supply price", "uz": "Tovar qiymati", "ru": "Цена поставщика"},
	}

	ProductFields = map[string]map[string]string{
		"BARCODE":          {"en": "Barcode", "uz": "Barcode", "ru": "Штрихкод"},
		"SKU":              {"en": "SKU", "uz": "SKU", "ru": "Артикул"},
		"NAME":             {"en": "Name", "uz": "Nomi", "ru": "Наименование"},
		"DESCRIPTION":      {"en": "Description", "uz": "Ta'rif", "ru": "Описание"},
		"MEASUREMENT_UNIT": {"en": "Measurement unit", "uz": "O'lchov birligi", "ru": "Единица измерения"},
		"QUANTITY":         {"en": "Quantity", "uz": "Soni", "ru": "Количество"},
		"PRICE":            {"en": "Price", "uz": "Qiymat", "ru": "Цена"},
		"RETAIL_PRICE":     {"en": "Retail price", "uz": "Retail qiymat", "ru": "Розничная цена"},
		"SUPPLY_PRICE":     {"en": "Supply price", "uz": "Tovar qiymati", "ru": "Цена поставщика"},
	}
)
