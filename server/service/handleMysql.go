package service

import (
	"fmt"
	"net/http"
	"wiki-server/inits"
	"wiki-server/types"
	"wiki-server/types/response"

	"github.com/gin-gonic/gin"
)

func ProcessGetWikis(ctx *gin.Context) {
	var wikiTitles []types.WikiTitle
	inits.DBEngine.Table("wikis").Select("title", "uuid").Find(&wikiTitles)
	response.Result(http.StatusOK, wikiTitles, "success", ctx)
}

func ProcessGetWiki(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	var wikiEntry types.WikiEntry
	inits.DBEngine.Table("wikis").Where("uuid = ?", uuid).First(&wikiEntry)
	res := types.GetWikiEntryResponse{
		Creator:    wikiEntry.Creator,
		Title:      wikiEntry.Title,
		Content:    wikiEntry.Content,
		CreateTime: wikiEntry.CreateTime,
		UpdateTime: wikiEntry.UpdateTime,
		Uuid:       wikiEntry.Uuid,
	}
	response.Result(http.StatusOK, res, "success", ctx)
}

func ProcessCreateWiki(ctx *gin.Context) {
	var createWikiEntryRequest types.CreateWikiEntryRequest
	if err := ctx.BindJSON(&createWikiEntryRequest); err != nil {
		// 处理错误
		response.Fail(ctx)
		return
	}
	var wikiEntry types.WikiEntry
	inits.DBEngine.Table("wikis")
	fmt.Println(wikiEntry)
	response.Result(http.StatusOK, wikiEntry, "success", ctx)
}

func ProcessUpdateWiki(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	var updateWikiEntryRequest types.UpdateWikiEntryRequest
	if err := ctx.BindJSON(&updateWikiEntryRequest); err != nil {
		response.Fail(ctx)
		return
	}
	var wikiEntry types.WikiEntry
	inits.DBEngine.Table("wikis").Where("uuid = ?", uuid).First(&wikiEntry)
	wikiEntry.Content = updateWikiEntryRequest.WikiContent
	inits.DBEngine.Table("wikis").Where("uuid = ?", uuid).Updates(&wikiEntry)
	response.Result(http.StatusOK, wikiEntry, "success", ctx)
}

func ProcessDeleteWiki(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	var wikiEntry types.WikiEntry
	inits.DBEngine.Table("wikis").Where("uuid = ?", uuid).First(&wikiEntry)
	fmt.Println(wikiEntry)
	response.Result(http.StatusOK, wikiEntry, "success", ctx)
}
