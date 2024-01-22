package controller

//func VerifyCookie(ctx *gin.Context) bool {
//	//return true, ""
//	JWT, err := ctx.Cookie("JWT")
//	if err != nil {
//		fmt.Println("hello")
//		ctx.JSON(400, gin.H{
//			"error": err.Error(),
//		})
//		return false
//	}
//	if JWT == "" {
//		ctx.JSON(400, gin.H{
//			"error": "no JWT",
//		})
//		return false
//	}
//	if verified, message := service.VerifyJWT(JWT); verified {
//		return true
//	} else {
//		ctx.JSON(400, gin.H{
//			"error": message,
//		})
//		return false
//	}
//}
//func HandleLogin(ctx *gin.Context) {
//	var user model.User
//	ctx.ShouldBind(&user)
//	if id := service.CheckUser(user); id != "" {
//		header := `{"alg":"HS256","typ":"JWT"}`
//		payload := fmt.Sprintf(`{"id":%v,"iat":%d}`, `"`+id+`"`, time.Now().Unix())
//		JWT := service.GenerateJWT(header, payload)
//		ctx.SetCookie("JWT", JWT, 57600, "/", "localhost", true, true)
//		ctx.JSON(200, gin.H{
//			"message": JWT,
//		})
//	} else {
//		ctx.JSON(400, gin.H{
//			"message": "wrong password",
//		})
//	}
//}
//func HandleLogout(ctx *gin.Context) {
//	fmt.Println(ctx.Cookie("JWT"))
//	ctx.SetCookie("JWT", "", 100, "/", "localhost", true, true)
//	ctx.JSON(200, gin.H{
//		"message": "log out",
//	})
//}
//func HandleGetOrderId(ctx *gin.Context) {
//	if check := VerifyCookie(ctx); !check {
//		return
//	}
//	id := ctx.Param("id")
//	if order, err := service.GetOrderService(id); err == nil {
//		ctx.JSON(200, gin.H{
//			"data": order,
//		})
//	} else {
//		ctx.JSON(404, gin.H{
//			"error": err.Error(),
//		})
//	}
//	return
//}
//
//func HandleCreateOrder(ctx *gin.Context) {
//	if check := VerifyCookie(ctx); !check {
//		return
//	}
//	var order model.Order
//	ctx.ShouldBind(&order)
//	if !service.CheckPriceOfOrder(order) {
//		ctx.JSON(400, gin.H{
//			"message": "price is wrong",
//		})
//		return
//	}
//	order.Buyer = service.FormatBuyerService(service.FilterBuyerService(order.Buyer))
//	fmt.Println(order)
//	order.Date = time.Now()
//	fmt.Println(order)
//	if err := service.InsertOrderService(order); err != nil {
//		ctx.JSON(404, gin.H{
//			"error": err.Error(),
//		})
//	} else {
//		ctx.JSON(200, gin.H{
//			"message": "successfully create new order",
//		})
//	}
//}
//func HandleDeleteOrder(ctx *gin.Context) {
//	if check := VerifyCookie(ctx); !check {
//		return
//	}
//	id := ctx.Param("id")
//	if err := service.DeleteOrderService(id); err != nil {
//		ctx.JSON(404, gin.H{
//			"error": err.Error(),
//		})
//	} else {
//		ctx.JSON(200, gin.H{
//			"message": " successfully delete order with id: " + id,
//		})
//	}
//}
//func HandleUpdateOrderStatus(ctx *gin.Context) {
//	if check := VerifyCookie(ctx); !check {
//		return
//	}
//	id := ctx.Param("id")
//	status, _ := strconv.Atoi(ctx.Param("status"))
//	if er := service.UpdateOrderStatusService(id, status); er != nil {
//		ctx.JSON(404, gin.H{
//			"error": er.Error(),
//		})
//		return
//	}
//	ctx.JSON(200, gin.H{
//		"message": "successfully updated status",
//	})
//
//}
//func HandleDeleteItem(ctx *gin.Context) {
//	if check := VerifyCookie(ctx); !check {
//		return
//	}
//	id := ctx.Param("id")
//	if err := service.DeleteItemService(id); err != nil {
//		ctx.JSON(404, gin.H{
//			"error": err.Error(),
//		})
//		return
//	}
//	ctx.JSON(200, gin.H{
//		"message": "Item deleted successfully",
//	})
//}
//func HandleUpdateItem(ctx *gin.Context) {
//	if check := VerifyCookie(ctx); check {
//		return
//	}
//	id := ctx.Param("id")
//	var item model.Item
//	ctx.ShouldBind(&item)
//	item.ID = id
//	if err := service.UpdateItemService(item); err != nil {
//		ctx.JSON(404, gin.H{
//			"error": err.Error(),
//		})
//		return
//	}
//	ctx.JSON(200, gin.H{
//		"message": "Successfully updated item",
//	})
//
//}
//func HandleGetOrder1(ctx *gin.Context) {
//
//}
