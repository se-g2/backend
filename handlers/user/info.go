package user

import (
	"backend/global"
	"backend/handlers"
	"backend/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InfoHandler(ctx *gin.Context) {
	var u types.User

	global.DB.First(&u, "id = ?", ctx.GetInt("uid"))

	var vehicles []types.VehicleInfo
	if u.IsDriver {
		global.DB.Find(&vehicles, "owner = ?", u.ID)
	}

	res := types.UserBaseResponse{
		ID: u.ID,
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
	}

	handlers.HandleSuccess(ctx, res, "信息获取成功", http.StatusOK)

}
