package service

//
//func encodeSegment(seg []byte) string {
//	return s.TrimRight(base64.URLEncoding.EncodeToString(seg), "=")
//}
//func decodeSegment(seg string) string {
//	str, _ := base64.URLEncoding.DecodeString(seg + "==")
//	return string(str)
//}
//func VerifyJWT(JWT string) (bool, string) {
//	JWTarr := s.Split(JWT, ".")
//	header := JWTarr[0]
//	payload := JWTarr[1]
//	signature := JWTarr[2]
//	key := "key"
//	mac := hmac.New(sha256.New, []byte(key))
//	mac.Write([]byte(header + "." + payload))
//	expectedSignature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
//	if signature == expectedSignature {
//		payloadJSON := decodeSegment(payload)
//		P := model.Payload{}
//		json.Unmarshal([]byte(payloadJSON), &P)
//		if a := time.Now().Unix() - P.Iat; a > 57600 {
//			fmt.Println(a)
//			fmt.Println(P)
//			fmt.Println(payloadJSON)
//			fmt.Println(time.Now().Unix())
//			return false, "time out"
//		} else {
//			return true, ""
//		}
//	}
//	return false, "invalid JWT"
//}
//
//func GenerateJWT(header string, payload string) string {
//	H := encodeSegment([]byte(header))
//	P := encodeSegment([]byte(payload))
//	key := "key"
//	mac := hmac.New(sha256.New, []byte(key))
//	mac.Write([]byte(H + "." + P))
//	S := base64.StdEncoding.EncodeToString(mac.Sum(nil))
//	return H + "." + P + "." + S
//}
//func CheckUser(user model.User) string {
//	return repository.CheckUser(user)
//}
//func GetOrderService(id string) (*model.Order, error) {
//	return repository.GetOrder(id)
//}
//func CheckPriceOfOrder(order model.Order) bool {
//	itemList := order.ItemList
//	leng := len(itemList)
//	price := 0
//	for i := 0; i < leng; i++ {
//		price += itemList[i].Amount * repository.GetItem(itemList[i].Item_id).Price
//	}
//	return price-order.Discount == order.Price
//}
//func InsertOrderService(order model.Order) error {
//	return repository.InsertOrder(order)
//}
//func DeleteOrderService(id string) error {
//	return repository.DeleteOrder(id)
//}
//func UpdateOrderStatusService(id string, status int) error {
//	return repository.UpdateOrderStatus(id, status)
//}
//func DeleteItemService(id string) error {
//	return repository.DeleteItem(id)
//}
//func FilterBuyerService(buyer string) []string { // " quang    khoi   "  => []{'quang';'khoi'}
//	buyer = s.TrimSpace(buyer)
//	ar := s.Split(buyer, " ")
//	l := len(ar)
//	var b []string
//	for i := 0; i < l; i++ {
//		if ar[i] != "" {
//			b = append(b, UpperFirstLetter(ar[i]))
//		}
//	}
//	return b
//}
//func FormatBuyerService(buyer []string) string { //[]{'quang';'khoi'} to "quang khoi"
//	return s.Join(buyer, " ")
//}
//func UpperFirstLetter(str string) string { // "van" to "Van"
//	runes := []rune(str)
//	if unicode.IsLetter(runes[0]) {
//		// Check if the first letter is not already uppercase
//		if !unicode.IsUpper(runes[0]) {
//			// Convert the first letter to uppercase
//			runes[0] = unicode.ToUpper(runes[0])
//		}
//	}
//	return string(runes)
//}
//func UpdateItemService(item model.Item) error {
//	return repository.UpdateItem(item)
//}
