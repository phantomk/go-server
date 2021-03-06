package notification_test

import (
	"github.com/axetroy/go-server/src/controller"
	"github.com/axetroy/go-server/src/controller/admin"
	"github.com/axetroy/go-server/src/controller/notification"
	"github.com/axetroy/go-server/src/schema"
	"github.com/axetroy/go-server/src/util"
	"github.com/axetroy/go-server/tester"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	var (
		adminUid string
	)
	// 先登陆获取管理员的Token
	{
		// 登陆超级管理员-成功

		r := admin.Login(admin.SignInParams{
			Username: "admin",
			Password: "admin",
		})

		assert.Equal(t, schema.StatusSuccess, r.Status)
		assert.Equal(t, "", r.Message)

		adminInfo := schema.AdminProfileWithToken{}

		if err := tester.Decode(r.Data, &adminInfo); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, "admin", adminInfo.Username)
		assert.True(t, len(adminInfo.Token) > 0)

		if c, er := util.ParseToken(util.TokenPrefix+" "+adminInfo.Token, true); er != nil {
			t.Error(er)
		} else {
			adminUid = c.Uid
		}
	}

	context := controller.Context{
		Uid: adminUid,
	}

	var testNotification schema.Notification

	// 创建一篇系统通知
	{
		var (
			title   = "TestGet"
			content = "TestGet"
		)

		r := notification.Create(context, notification.CreateParams{
			Tittle:  title,
			Content: content,
		})

		assert.Equal(t, schema.StatusSuccess, r.Status)
		assert.Equal(t, "", r.Message)

		testNotification = schema.Notification{}

		assert.Nil(t, tester.Decode(r.Data, &testNotification))

		defer notification.DeleteNotificationById(testNotification.Id)

		assert.Equal(t, title, testNotification.Tittle)
		assert.Equal(t, content, testNotification.Content)
	}

	// 获取详情
	{
		r := notification.Get(context, testNotification.Id)

		assert.Equal(t, schema.StatusSuccess, r.Status)
		assert.Equal(t, "", r.Message)
		n := schema.Notification{}
		assert.Nil(t, tester.Decode(r.Data, &n))

		assert.Equal(t, n.Id, testNotification.Id)
		assert.Equal(t, n.Tittle, testNotification.Tittle)
		assert.Equal(t, n.Content, testNotification.Content)
		assert.Equal(t, false, testNotification.Read)
	}
}
