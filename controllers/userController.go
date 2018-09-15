package controllers

import (
	"encoding/json"
	// "fmt"
	"github.com/ChallenAi/iTools-service/models"
	"github.com/valyala/fasthttp"
)

func GetAllUsers(ctx *fasthttp.RequestCtx) {
	user := &models.User{
		UserId:   1,
		UserName: "Challen",
	}
	resp, _ := json.Marshal(user)

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	ctx.Write(resp)
}

func Login(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Login")

}

func Register(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Register")

}
