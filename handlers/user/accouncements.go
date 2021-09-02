package user

import (
	"backend/global"
	"backend/handlers"
	"backend/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AnnouncementsHandler(ctx *gin.Context) {
	var announcements []types.Announcement
	global.DB.Order("id desc").Limit(10).Find(&announcements)
	var res []types.AnnouncementResponse
	for _, announcement := range announcements {
		res = append(res, types.AnnouncementResponse{
			ID:       int(announcement.ID),
			Title:    announcement.Title,
			Publish:  announcement.CreatedAt,
			Update:   announcement.UpdatedAt,
			Category: announcement.Category,
			Content:  announcement.Content,
		})
	}
	handlers.HandleSuccess(ctx, res, "公告获取成功", http.StatusOK)
}
