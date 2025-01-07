package job

import (
	"fmt"
	"xiaozhu/internal/model/pay"
)

type OrderJob struct {
	pay.Order
}

func (j *OrderJob) Run() {

	fmt.Println(6666)
}
