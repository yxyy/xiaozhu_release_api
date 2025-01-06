package queue

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"
	"time"
	"xiaozhu/internal/model/assets"
	"xiaozhu/internal/model/common"
	"xiaozhu/internal/model/key"
	"xiaozhu/internal/model/pay"
	"xiaozhu/utils"
	"xiaozhu/utils/queue"
)

type OrderQueue struct {
	common.TopicExtra
	Message pay.Order
}

func NewOrderQueue() *queue.Queue {
	return queue.NewQueue(key.OrderQueue, &OrderQueue{})
}

func (l *OrderQueue) Run(q *queue.Queue, topic string) error {
	if err := json.Unmarshal([]byte(topic), &l); err != nil {
		return fmt.Errorf("序列化失败：%w", err)
	}

	if l.Message.OrderNum == "" {
		return fmt.Errorf("无效的订单")
	}

	order := &pay.Order{OrderNum: l.Message.OrderNum}
	if err := order.GetOrderInfo(q.Ctx); err != nil {
		return fmt.Errorf("获取订单信息失败，订单号: %s，错误: %w", l.Message.OrderNum, err)
	}

	game, err := assets.GetAppGameInfo(q.Ctx, order.GameId)
	if err != nil {
		return fmt.Errorf("订单号: %s，无效的游戏", l.Message.OrderNum)
	}
	if game.CpCallbackUrl == "" {
		return fmt.Errorf("订单号: %s，该游戏支付通知地址未配置", l.Message.OrderNum)
	}

	requestBody, err := GetRequestBody(game, order)
	if err != nil {
		return fmt.Errorf("订单号: %s，获取请求 body 错误: %w", l.Message.OrderNum, err)
	}

	resp, err := utils.Request(q.Ctx, "POST", game.CpCallbackUrl, requestBody)
	if err != nil {
		return fmt.Errorf("订单号: %s，发送错误: %w", l.Message.OrderNum, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	response := new(CpOrderResponse)
	if err = json.Unmarshal(body, response); err != nil {
		return fmt.Errorf("读取body序列化失败：%w", err)
	}

	if response.Code == 0 {
		order.CpOrderAt = time.Now().Unix()
		order.CpOrderStatus = 1
	}

	return order.Save(q.Ctx)
}

type CpOrderRequest struct {
	OrderType     int8   `json:"order_type"`
	ZoneId        int    `json:"zone_id"`
	RoleId        int    `json:"role_id"`
	CpOrderNum    string `json:"cp_order_num"`
	GoodsId       string `json:"goods_id"`
	Amount        int    `json:"amount"`
	ActualAmount  int    `json:"actual_amount"`
	UserId        int    `json:"user_id"`
	TradeOrderNum string `json:"trade_order_num"`
	Ts            int64  `json:"ts"`
	ExtData       string `json:"ext_data"`
	Sandbox       int8   `json:"sandbox"`
	Currency      string `json:"currency"`
	Sign          string `json:"sign"`
}

type CpOrderResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func GetRequestBody(game *assets.AppGame, order *pay.Order) (io.Reader, error) {
	if game.AppKey == "" {
		return nil, fmt.Errorf("该游戏支付 key 未配置")
	}

	request := CpOrderRequest{
		OrderType:     order.OrderType,
		ZoneId:        order.ZoneId,
		RoleId:        order.RoleId,
		CpOrderNum:    order.CpOrderNum,
		GoodsId:       order.GoodsId,
		Amount:        order.OrderPrice,
		ActualAmount:  order.PayMoney,
		UserId:        order.UserId,
		TradeOrderNum: order.OrderNum,
		Ts:            order.PayAt,
		ExtData:       order.ExtData,
		Sandbox:       order.SandBox,
		Currency:      order.PayCurrency,
		Sign:          "",
	}

	request.Sign = GetSign(request, game.AppKey)

	marshal, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(marshal), nil
}

func GetSign(request CpOrderRequest, key string) string {
	sortString := SortStruct(request)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(sortString+key)))
}

// SortStruct 按结构体字段名 ASCII 字典顺序拼接 key=value
func SortStruct(input interface{}) string {
	v := reflect.ValueOf(input)
	t := reflect.TypeOf(input)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	// 确保是结构体
	if v.Kind() != reflect.Struct {
		return ""
	}

	// 保存字段名和对应值的键值对
	var kvPairs []string

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Tag.Get("json")
		if fieldName == "" {
			fieldName = field.Name
		}

		if fieldName == "sign" {
			continue
		}

		// 跳过空字段
		fieldValue := fmt.Sprintf("%v", v.Field(i).Interface())
		// if fieldValue == "" {
		// 	continue
		// }

		// 特殊字符（如 & 或 =） 添加 URL 编码
		fieldValue = strings.ReplaceAll(fieldValue, "&", "%26")
		fieldValue = strings.ReplaceAll(fieldValue, "=", "%3D")

		kvPairs = append(kvPairs, fmt.Sprintf("%s=%s", fieldName, fieldValue))
	}

	// 按字段名 ASCII 排序
	sort.Strings(kvPairs)

	// 拼接成字符串
	return strings.Join(kvPairs, "&")
}
