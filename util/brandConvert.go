package util

func BrandConvert(brandNo int) string {
	switch brandNo {
	case 1:
		return "UberEats"
	case 2:
		return "Rakuten"
	}
	return ""
}

func BrandConvertNo(brand string) int {
	switch brand {
	case "UberEats":
		return 1
	case "Rakuten":
		return 2
	}
	return 0
}

func GetAllBrand() []string {
	return []string{
		"UberEats",
		"Rakuten",
		"100%",
		"CoCo壱番屋",
		"Mitsui",
		"なか卯",
		"ウイッシュオリジナル",
	}
}
