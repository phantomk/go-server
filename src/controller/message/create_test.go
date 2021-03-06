package message_test

import (
	"encoding/json"
	"github.com/axetroy/go-server/src/controller"
	"github.com/axetroy/go-server/src/controller/admin"
	"github.com/axetroy/go-server/src/controller/auth"
	"github.com/axetroy/go-server/src/controller/message"
	"github.com/axetroy/go-server/src/exception"
	"github.com/axetroy/go-server/src/model"
	"github.com/axetroy/go-server/src/schema"
	"github.com/axetroy/go-server/src/util"
	"github.com/axetroy/go-server/tester"
	"github.com/axetroy/mocker"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"testing"
)

func init() {
	// 确保超级管理员存在
	admin.CreateAdmin(admin.CreateAdminParams{
		Account:  "admin",
		Password: "admin",
		Name:     "admin",
	}, true)
}

func TestCreate(t *testing.T) {
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

	// 创建一个公告
	{
		var (
			title   = "test"
			content = "test"
		)

		r := message.Create(controller.Context{
			Uid: adminUid,
		}, message.CreateMessageParams{
			Title:   title,
			Content: content,
		})

		assert.Equal(t, schema.StatusSuccess, r.Status)
		assert.Equal(t, "", r.Message)

		n := model.Message{}

		assert.Nil(t, tester.Decode(r.Data, &n))

		defer message.DeleteMessageById(n.Id)

		assert.Equal(t, title, n.Tittle)
		assert.Equal(t, content, n.Content)
	}

	// 非管理员的uid去创建，应该报错
	{
		// 创建一个普通用户
		var (
			username = "tester-normal"
			uid      string
		)

		{
			rand.Seed(10331)
			password := "123123"

			r := auth.SignUp(auth.SignUpParams{
				Username: &username,
				Password: password,
			})

			profile := schema.Profile{}

			assert.Nil(t, tester.Decode(r.Data, &profile))

			defer func() {
				auth.DeleteUserByUserName(username)
			}()

			uid = profile.Id
		}

		var (
			title   = "test"
			content = "test"
		)

		r := message.Create(controller.Context{
			Uid: uid,
		}, message.CreateMessageParams{
			Title:   title,
			Content: content,
		})

		assert.Equal(t, schema.StatusFail, r.Status)
		assert.Equal(t, exception.AdminNotExist.Error(), r.Message)
	}
}

func TestCreateRouter(t *testing.T) {
	var (
		adminToken string
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

		if _, er := util.ParseToken(util.TokenPrefix+" "+adminInfo.Token, true); er != nil {
			t.Error(er)
		} else {
			adminToken = adminInfo.Token
		}
	}

	// 登陆正确的管理员账号
	{
		var (
			title   = "test"
			content = "test"
		)

		header := mocker.Header{
			"Authorization": util.TokenPrefix + " " + adminToken,
		}

		body, _ := json.Marshal(&message.CreateMessageParams{
			Title:   title,
			Content: content,
		})

		r := tester.Http.Post("/v1/admin/message/create", body, &header)
		res := schema.Response{}

		assert.Equal(t, http.StatusOK, r.Code)
		assert.Nil(t, json.Unmarshal([]byte(r.Body.String()), &res))

		n := schema.Message{}

		assert.Nil(t, tester.Decode(res.Data, &n))

		defer message.DeleteMessageById(n.Id)

		assert.Equal(t, title, n.Tittle)
		assert.Equal(t, content, n.Content)
	}
}
