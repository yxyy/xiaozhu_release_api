package mapping

import (
	"xiaozhu/internal/model/assets"
)

func AppType() (map[int]string, error) {
	appType := assets.AppType{}
	companies, err := appType.GetAll()
	if err != nil {
		return nil, err
	}

	appTypes := make(map[int]string)
	for _, v := range companies {
		appTypes[int(v.Id)] = v.Name
	}

	return appTypes, nil
}
