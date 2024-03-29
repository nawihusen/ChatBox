package handler

// import (
// 	"alpha-test/domain"
// 	"alpha-test/helper"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/gofiber/fiber/v2"
// 	log "github.com/sirupsen/logrus"

// 	// "github.com/spf13/viper"
// 	"github.com/valyala/fasthttp"
// )

// // Handler is REST API handler for Service System
// type ArticleHandler struct {
// 	ArticleUsecase domain.ArticleUsecase
// }

// func (sh *ArticleHandler) PostArticleHandler(c *fiber.Ctx) error {
// 	var input domain.Article
// 	err := c.BodyParser(&input)
// 	if err != nil {
// 		log.Error(err)
// 		return helper.HTTPSimpleResponse(c, fasthttp.StatusBadRequest)
// 	}

// 	validate := validator.New()
// 	err = validate.Struct(input)
// 	if err != nil {
// 		log.Error(err)
// 		return helper.HTTPSimpleResponse(c, fasthttp.StatusBadRequest)
// 	}

// 	err = sh.ArticleUsecase.PostArticle(c.Context(), input)
// 	if err != nil {
// 		log.Error(err)
// 		return helper.HTTPSimpleResponse(c, fasthttp.StatusInternalServerError)
// 	}

// 	return c.Status(fasthttp.StatusOK).JSON("Success")
// }

// func (sh *ArticleHandler) GetArticlesHandler(c *fiber.Ctx) (err error) {
// 	author := c.Query("author")

// 	title := c.Query("title")

// 	body := c.Query("body")

// 	response, err := sh.ArticleUsecase.GetArticles(c.Context(), author, title, body)
// 	if err != nil {
// 		log.Error(err)
// 		return c.Status(500).SendString(fasthttp.StatusMessage(fasthttp.StatusInternalServerError))
// 	} else if len(response) == 0 {
// 		return helper.HTTPSimpleResponse(c, fasthttp.StatusNotFound)
// 	}

// 	err = c.Status(fasthttp.StatusOK).JSON(response)
// 	return
// }
