package mapping

import (
	"xiaozhu/backend/internal/model/assets"
)

func Company() (map[int]string, error) {
	company := assets.Company{}
	companies, err := company.GetAll()
	if err != nil {
		return nil, err
	}

	companys := make(map[int]string)
	for _, v := range companies {
		companys[v.Id] = v.Name
	}

	return companys, nil
}
