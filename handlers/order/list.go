package order

import (
	"backend/global"
	"backend/handlers"
	"backend/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getOrder(orderId uint) types.OrderResponse {
	var order types.Order
	global.DB.First(&order, orderId)
	var u types.User
	var passengersResponse []types.UserBaseResponse

	for _, uid := range order.BookingPassengerIDs {
		global.DB.First(&u, uid)

		var vehicles []types.VehicleInfo
		if u.IsDriver {
			global.DB.Find(&vehicles, "owner = ?", u.ID)
		}

		passengersResponse = append(passengersResponse, types.UserBaseResponse{
			ID: uid,
			Profile: types.UserProfile{
				Avatar: u.Avatar,
				Bio:    u.Bio,
			},
			SchoolID:     u.SchoolID,
			Name:         u.Name,
			Tele:         u.Tele,
			Score:        u.Score,
			FinishOrders: u.FinishOrders,
			Like:         u.Like,
			Money:        u.Money,
			IsDriver:     u.IsDriver,
			DriverProfile: types.DriverProfileResponse{
				IdentityCode:  u.IdentityCode,
				DriverLicense: u.DriverLicense,
				Vehicles:      vehicles,
				AvailableTime: u.AvailableTime,
			},
			IsAdmin: u.IsAdmin,
		})
	}

	var vehicle types.VehicleInfo
	global.DB.Find(&vehicle, order.VehicleInfoID)

	var driver types.User
	global.DB.First(&driver, order.DriverID)

	return types.OrderResponse{
		ID:               int(order.ID),
		CreateTime:       order.CreatedAt,
		Creator:          int(order.Creator),
		Departure:        order.Departure,
		StartTime:        order.StartTime,
		EndTime:          order.EndTime,
		Distance:         order.Distance,
		Price:            order.Price,
		BookingPassenger: passengersResponse,
		TargetPassenger:  order.TargetPassenger,
		Vehicle:          vehicle,
		Driver: types.DriverProfileResponse{
			IdentityCode:  driver.IdentityCode,
			DriverLicense: driver.DriverLicense,
			Vehicles:      nil,
			AvailableTime: driver.AvailableTime,
		},
		Status:   order.Status,
		Complete: order.IsComplete,
	}
}

func ListHandler(ctx *gin.Context) {

	var res []types.OrderResponse

	identity := ctx.Param("identity")

	if identity == "passenger" {
		var txs []types.Transaction
		global.DB.Find(&txs, "user_id = ?", ctx.GetInt("uid"))
		for _, tx := range txs {
			res = append(res, getOrder(tx.OrderID))
		}
	} else if identity == "driver" {
		var orders []types.Order
		global.DB.Find(&orders, "driver_id = ?", ctx.GetInt("uid"))
		for _, order := range orders {
			res = append(res, getOrder(order.ID))
		}
	}
	handlers.HandleSuccess(ctx, res, "订单获取成功", http.StatusOK)
}
