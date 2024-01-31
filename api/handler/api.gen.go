// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package handler

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for GameCondition.
const (
	GameConditionFair GameCondition = "fair"
	GameConditionGood GameCondition = "good"
	GameConditionMint GameCondition = "mint"
	GameConditionPoor GameCondition = "poor"
)

// Defines values for NewGameCondition.
const (
	NewGameConditionFair NewGameCondition = "fair"
	NewGameConditionGood NewGameCondition = "good"
	NewGameConditionMint NewGameCondition = "mint"
	NewGameConditionPoor NewGameCondition = "poor"
)

// Defines values for OfferStatus.
const (
	Accepted OfferStatus = "accepted"
	Pending  OfferStatus = "pending"
	Rejected OfferStatus = "rejected"
)

// Game defines model for Game.
type Game struct {
	Condition *GameCondition      `json:"condition,omitempty"`
	Id        *openapi_types.UUID `json:"id,omitempty"`
	Name      *string             `json:"name,omitempty"`
	Publisher *string             `json:"publisher,omitempty"`
}

// GameCondition defines model for Game.Condition.
type GameCondition string

// NewGame defines model for NewGame.
type NewGame struct {
	Condition NewGameCondition `json:"condition"`
	Name      string           `json:"name"`
	Publisher string           `json:"publisher"`
}

// NewGameCondition defines model for NewGame.Condition.
type NewGameCondition string

// NewOffer defines model for NewOffer.
type NewOffer struct {
	GamesOffered   []openapi_types.UUID `json:"gamesOffered"`
	GamesRequested []openapi_types.UUID `json:"gamesRequested"`
	Offeree        openapi_types.UUID   `json:"offeree"`
	Offeror        openapi_types.UUID   `json:"offeror"`
}

// NewUser defines model for NewUser.
type NewUser struct {
	Address string `json:"address"`
	Email   string `json:"email"`
	Name    string `json:"name"`
}

// Offer defines model for Offer.
type Offer struct {
	GamesOffered   *[]openapi_types.UUID `json:"gamesOffered,omitempty"`
	GamesRequested *[]openapi_types.UUID `json:"gamesRequested,omitempty"`
	Id             *openapi_types.UUID   `json:"id,omitempty"`
	Offeree        *openapi_types.UUID   `json:"offeree,omitempty"`
	Offeror        *openapi_types.UUID   `json:"offeror,omitempty"`
	Status         *OfferStatus          `json:"status,omitempty"`
	Timestamp      *string               `json:"timestamp,omitempty"`
}

// OfferStatus defines model for Offer.Status.
type OfferStatus string

// User defines model for User.
type User struct {
	Address *string             `json:"address,omitempty"`
	Email   *string             `json:"email,omitempty"`
	Id      *openapi_types.UUID `json:"id,omitempty"`
	Name    *string             `json:"name,omitempty"`
}

// PostGamesJSONRequestBody defines body for PostGames for application/json ContentType.
type PostGamesJSONRequestBody = NewGame

// PutGamesIdJSONRequestBody defines body for PutGamesId for application/json ContentType.
type PutGamesIdJSONRequestBody = NewGame

// PostOffersJSONRequestBody defines body for PostOffers for application/json ContentType.
type PostOffersJSONRequestBody = NewOffer

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody = NewUser

// PutUsersIdJSONRequestBody defines body for PutUsersId for application/json ContentType.
type PutUsersIdJSONRequestBody = NewUser

// PutUsersIdChangePasswordJSONRequestBody defines body for PutUsersIdChangePassword for application/json ContentType.
type PutUsersIdChangePasswordJSONRequestBody = User

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create a new video game
	// (POST /games)
	PostGames(c *gin.Context)
	// Delete a video game
	// (DELETE /games/{id})
	DeleteGamesId(c *gin.Context, id string)
	// Get a specific video game
	// (GET /games/{id})
	GetGamesId(c *gin.Context, id string)
	// Update a video game
	// (PUT /games/{id})
	PutGamesId(c *gin.Context, id string)
	// Create a new trade offer
	// (POST /offers)
	PostOffers(c *gin.Context)
	// Delete a trade offer
	// (DELETE /offers/{id})
	DeleteOffersId(c *gin.Context, id string)
	// Get a specific trade offer
	// (GET /offers/{id})
	GetOffersId(c *gin.Context, id string)
	// Register a new user
	// (POST /users)
	PostUsers(c *gin.Context)
	// Get a specific user
	// (GET /users/{id})
	GetUser(c *gin.Context, id string)
	// Update user information
	// (PUT /users/{id})
	PutUsersId(c *gin.Context, id string)
	// Change user password
	// (PUT /users/{id}/change-password)
	PutUsersIdChangePassword(c *gin.Context, id int)
	// Reset user password
	// (POST /users/{id}/reset-password)
	PostUsersIdResetPassword(c *gin.Context, id int)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// PostGames operation middleware
func (siw *ServerInterfaceWrapper) PostGames(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostGames(c)
}

// DeleteGamesId operation middleware
func (siw *ServerInterfaceWrapper) DeleteGamesId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteGamesId(c, id)
}

// GetGamesId operation middleware
func (siw *ServerInterfaceWrapper) GetGamesId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetGamesId(c, id)
}

// PutGamesId operation middleware
func (siw *ServerInterfaceWrapper) PutGamesId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutGamesId(c, id)
}

// PostOffers operation middleware
func (siw *ServerInterfaceWrapper) PostOffers(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostOffers(c)
}

// DeleteOffersId operation middleware
func (siw *ServerInterfaceWrapper) DeleteOffersId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteOffersId(c, id)
}

// GetOffersId operation middleware
func (siw *ServerInterfaceWrapper) GetOffersId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetOffersId(c, id)
}

// PostUsers operation middleware
func (siw *ServerInterfaceWrapper) PostUsers(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostUsers(c)
}

// GetUser operation middleware
func (siw *ServerInterfaceWrapper) GetUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUser(c, id)
}

// PutUsersId operation middleware
func (siw *ServerInterfaceWrapper) PutUsersId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutUsersId(c, id)
}

// PutUsersIdChangePassword operation middleware
func (siw *ServerInterfaceWrapper) PutUsersIdChangePassword(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutUsersIdChangePassword(c, id)
}

// PostUsersIdResetPassword operation middleware
func (siw *ServerInterfaceWrapper) PostUsersIdResetPassword(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostUsersIdResetPassword(c, id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/games", wrapper.PostGames)
	router.DELETE(options.BaseURL+"/games/:id", wrapper.DeleteGamesId)
	router.GET(options.BaseURL+"/games/:id", wrapper.GetGamesId)
	router.PUT(options.BaseURL+"/games/:id", wrapper.PutGamesId)
	router.POST(options.BaseURL+"/offers", wrapper.PostOffers)
	router.DELETE(options.BaseURL+"/offers/:id", wrapper.DeleteOffersId)
	router.GET(options.BaseURL+"/offers/:id", wrapper.GetOffersId)
	router.POST(options.BaseURL+"/users", wrapper.PostUsers)
	router.GET(options.BaseURL+"/users/:id", wrapper.GetUser)
	router.PUT(options.BaseURL+"/users/:id", wrapper.PutUsersId)
	router.PUT(options.BaseURL+"/users/:id/change-password", wrapper.PutUsersIdChangePassword)
	router.POST(options.BaseURL+"/users/:id/reset-password", wrapper.PostUsersIdResetPassword)
}

type PostGamesRequestObject struct {
	Body *PostGamesJSONRequestBody
}

type PostGamesResponseObject interface {
	VisitPostGamesResponse(w http.ResponseWriter) error
}

type PostGames201JSONResponse Game

func (response PostGames201JSONResponse) VisitPostGamesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type DeleteGamesIdRequestObject struct {
	Id string `json:"id"`
}

type DeleteGamesIdResponseObject interface {
	VisitDeleteGamesIdResponse(w http.ResponseWriter) error
}

type DeleteGamesId200Response struct {
}

func (response DeleteGamesId200Response) VisitDeleteGamesIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type DeleteGamesId404Response struct {
}

func (response DeleteGamesId404Response) VisitDeleteGamesIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetGamesIdRequestObject struct {
	Id string `json:"id"`
}

type GetGamesIdResponseObject interface {
	VisitGetGamesIdResponse(w http.ResponseWriter) error
}

type GetGamesId200JSONResponse Game

func (response GetGamesId200JSONResponse) VisitGetGamesIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetGamesId400Response struct {
}

func (response GetGamesId400Response) VisitGetGamesIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type GetGamesId404Response struct {
}

func (response GetGamesId404Response) VisitGetGamesIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PutGamesIdRequestObject struct {
	Id   string `json:"id"`
	Body *PutGamesIdJSONRequestBody
}

type PutGamesIdResponseObject interface {
	VisitPutGamesIdResponse(w http.ResponseWriter) error
}

type PutGamesId200JSONResponse Game

func (response PutGamesId200JSONResponse) VisitPutGamesIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PutGamesId400Response struct {
}

func (response PutGamesId400Response) VisitPutGamesIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type PutGamesId404Response struct {
}

func (response PutGamesId404Response) VisitPutGamesIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PostOffersRequestObject struct {
	Body *PostOffersJSONRequestBody
}

type PostOffersResponseObject interface {
	VisitPostOffersResponse(w http.ResponseWriter) error
}

type PostOffers201JSONResponse Offer

func (response PostOffers201JSONResponse) VisitPostOffersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostOffers400Response struct {
}

func (response PostOffers400Response) VisitPostOffersResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type DeleteOffersIdRequestObject struct {
	Id string `json:"id"`
}

type DeleteOffersIdResponseObject interface {
	VisitDeleteOffersIdResponse(w http.ResponseWriter) error
}

type DeleteOffersId200Response struct {
}

func (response DeleteOffersId200Response) VisitDeleteOffersIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type DeleteOffersId404Response struct {
}

func (response DeleteOffersId404Response) VisitDeleteOffersIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetOffersIdRequestObject struct {
	Id string `json:"id"`
}

type GetOffersIdResponseObject interface {
	VisitGetOffersIdResponse(w http.ResponseWriter) error
}

type GetOffersId200JSONResponse Offer

func (response GetOffersId200JSONResponse) VisitGetOffersIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetOffersId404Response struct {
}

func (response GetOffersId404Response) VisitGetOffersIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PostUsersRequestObject struct {
	Body *PostUsersJSONRequestBody
}

type PostUsersResponseObject interface {
	VisitPostUsersResponse(w http.ResponseWriter) error
}

type PostUsers201JSONResponse User

func (response PostUsers201JSONResponse) VisitPostUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostUsers400Response struct {
}

func (response PostUsers400Response) VisitPostUsersResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type GetUserRequestObject struct {
	Id string `json:"id"`
}

type GetUserResponseObject interface {
	VisitGetUserResponse(w http.ResponseWriter) error
}

type GetUser200JSONResponse User

func (response GetUser200JSONResponse) VisitGetUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetUser404Response struct {
}

func (response GetUser404Response) VisitGetUserResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PutUsersIdRequestObject struct {
	Id   string `json:"id"`
	Body *PutUsersIdJSONRequestBody
}

type PutUsersIdResponseObject interface {
	VisitPutUsersIdResponse(w http.ResponseWriter) error
}

type PutUsersId200JSONResponse User

func (response PutUsersId200JSONResponse) VisitPutUsersIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PutUsersId400Response struct {
}

func (response PutUsersId400Response) VisitPutUsersIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type PutUsersId404Response struct {
}

func (response PutUsersId404Response) VisitPutUsersIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PutUsersIdChangePasswordRequestObject struct {
	Id   int `json:"id"`
	Body *PutUsersIdChangePasswordJSONRequestBody
}

type PutUsersIdChangePasswordResponseObject interface {
	VisitPutUsersIdChangePasswordResponse(w http.ResponseWriter) error
}

type PutUsersIdChangePassword200Response struct {
}

func (response PutUsersIdChangePassword200Response) VisitPutUsersIdChangePasswordResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PutUsersIdChangePassword400Response struct {
}

func (response PutUsersIdChangePassword400Response) VisitPutUsersIdChangePasswordResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type PutUsersIdChangePassword404Response struct {
}

func (response PutUsersIdChangePassword404Response) VisitPutUsersIdChangePasswordResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PostUsersIdResetPasswordRequestObject struct {
	Id int `json:"id"`
}

type PostUsersIdResetPasswordResponseObject interface {
	VisitPostUsersIdResetPasswordResponse(w http.ResponseWriter) error
}

type PostUsersIdResetPassword200Response struct {
}

func (response PostUsersIdResetPassword200Response) VisitPostUsersIdResetPasswordResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PostUsersIdResetPassword400Response struct {
}

func (response PostUsersIdResetPassword400Response) VisitPostUsersIdResetPasswordResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type PostUsersIdResetPassword404Response struct {
}

func (response PostUsersIdResetPassword404Response) VisitPostUsersIdResetPasswordResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Create a new video game
	// (POST /games)
	PostGames(ctx context.Context, request PostGamesRequestObject) (PostGamesResponseObject, error)
	// Delete a video game
	// (DELETE /games/{id})
	DeleteGamesId(ctx context.Context, request DeleteGamesIdRequestObject) (DeleteGamesIdResponseObject, error)
	// Get a specific video game
	// (GET /games/{id})
	GetGamesId(ctx context.Context, request GetGamesIdRequestObject) (GetGamesIdResponseObject, error)
	// Update a video game
	// (PUT /games/{id})
	PutGamesId(ctx context.Context, request PutGamesIdRequestObject) (PutGamesIdResponseObject, error)
	// Create a new trade offer
	// (POST /offers)
	PostOffers(ctx context.Context, request PostOffersRequestObject) (PostOffersResponseObject, error)
	// Delete a trade offer
	// (DELETE /offers/{id})
	DeleteOffersId(ctx context.Context, request DeleteOffersIdRequestObject) (DeleteOffersIdResponseObject, error)
	// Get a specific trade offer
	// (GET /offers/{id})
	GetOffersId(ctx context.Context, request GetOffersIdRequestObject) (GetOffersIdResponseObject, error)
	// Register a new user
	// (POST /users)
	PostUsers(ctx context.Context, request PostUsersRequestObject) (PostUsersResponseObject, error)
	// Get a specific user
	// (GET /users/{id})
	GetUser(ctx context.Context, request GetUserRequestObject) (GetUserResponseObject, error)
	// Update user information
	// (PUT /users/{id})
	PutUsersId(ctx context.Context, request PutUsersIdRequestObject) (PutUsersIdResponseObject, error)
	// Change user password
	// (PUT /users/{id}/change-password)
	PutUsersIdChangePassword(ctx context.Context, request PutUsersIdChangePasswordRequestObject) (PutUsersIdChangePasswordResponseObject, error)
	// Reset user password
	// (POST /users/{id}/reset-password)
	PostUsersIdResetPassword(ctx context.Context, request PostUsersIdResetPasswordRequestObject) (PostUsersIdResetPasswordResponseObject, error)
}

type StrictHandlerFunc = strictgin.StrictGinHandlerFunc
type StrictMiddlewareFunc = strictgin.StrictGinMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// PostGames operation middleware
func (sh *strictHandler) PostGames(ctx *gin.Context) {
	var request PostGamesRequestObject

	var body PostGamesJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostGames(ctx, request.(PostGamesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostGames")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PostGamesResponseObject); ok {
		if err := validResponse.VisitPostGamesResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteGamesId operation middleware
func (sh *strictHandler) DeleteGamesId(ctx *gin.Context, id string) {
	var request DeleteGamesIdRequestObject

	request.Id = id

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteGamesId(ctx, request.(DeleteGamesIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteGamesId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(DeleteGamesIdResponseObject); ok {
		if err := validResponse.VisitDeleteGamesIdResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetGamesId operation middleware
func (sh *strictHandler) GetGamesId(ctx *gin.Context, id string) {
	var request GetGamesIdRequestObject

	request.Id = id

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetGamesId(ctx, request.(GetGamesIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetGamesId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetGamesIdResponseObject); ok {
		if err := validResponse.VisitGetGamesIdResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// PutGamesId operation middleware
func (sh *strictHandler) PutGamesId(ctx *gin.Context, id string) {
	var request PutGamesIdRequestObject

	request.Id = id

	var body PutGamesIdJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PutGamesId(ctx, request.(PutGamesIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PutGamesId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PutGamesIdResponseObject); ok {
		if err := validResponse.VisitPutGamesIdResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostOffers operation middleware
func (sh *strictHandler) PostOffers(ctx *gin.Context) {
	var request PostOffersRequestObject

	var body PostOffersJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostOffers(ctx, request.(PostOffersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostOffers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PostOffersResponseObject); ok {
		if err := validResponse.VisitPostOffersResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteOffersId operation middleware
func (sh *strictHandler) DeleteOffersId(ctx *gin.Context, id string) {
	var request DeleteOffersIdRequestObject

	request.Id = id

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteOffersId(ctx, request.(DeleteOffersIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteOffersId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(DeleteOffersIdResponseObject); ok {
		if err := validResponse.VisitDeleteOffersIdResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetOffersId operation middleware
func (sh *strictHandler) GetOffersId(ctx *gin.Context, id string) {
	var request GetOffersIdRequestObject

	request.Id = id

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetOffersId(ctx, request.(GetOffersIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetOffersId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetOffersIdResponseObject); ok {
		if err := validResponse.VisitGetOffersIdResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostUsers operation middleware
func (sh *strictHandler) PostUsers(ctx *gin.Context) {
	var request PostUsersRequestObject

	var body PostUsersJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostUsers(ctx, request.(PostUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PostUsersResponseObject); ok {
		if err := validResponse.VisitPostUsersResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetUser operation middleware
func (sh *strictHandler) GetUser(ctx *gin.Context, id string) {
	var request GetUserRequestObject

	request.Id = id

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUser(ctx, request.(GetUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUser")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetUserResponseObject); ok {
		if err := validResponse.VisitGetUserResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// PutUsersId operation middleware
func (sh *strictHandler) PutUsersId(ctx *gin.Context, id string) {
	var request PutUsersIdRequestObject

	request.Id = id

	var body PutUsersIdJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PutUsersId(ctx, request.(PutUsersIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PutUsersId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PutUsersIdResponseObject); ok {
		if err := validResponse.VisitPutUsersIdResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// PutUsersIdChangePassword operation middleware
func (sh *strictHandler) PutUsersIdChangePassword(ctx *gin.Context, id int) {
	var request PutUsersIdChangePasswordRequestObject

	request.Id = id

	var body PutUsersIdChangePasswordJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PutUsersIdChangePassword(ctx, request.(PutUsersIdChangePasswordRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PutUsersIdChangePassword")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PutUsersIdChangePasswordResponseObject); ok {
		if err := validResponse.VisitPutUsersIdChangePasswordResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostUsersIdResetPassword operation middleware
func (sh *strictHandler) PostUsersIdResetPassword(ctx *gin.Context, id int) {
	var request PostUsersIdResetPasswordRequestObject

	request.Id = id

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostUsersIdResetPassword(ctx, request.(PostUsersIdResetPasswordRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostUsersIdResetPassword")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PostUsersIdResetPasswordResponseObject); ok {
		if err := validResponse.VisitPostUsersIdResetPasswordResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RYTW/jNhD9K8S0R+3a+Siw0Knd3SJIgTZBtumlyIErjmUGEsklqaRB4P9ecCjZkiNZ",
	"ahoL7ckWzRnOvPdmhvIzZLo0WqHyDtJncNkaS05fL3iJ4dNYbdB6ibSaaSWkl1qFB1RVCemfUErlIYFc",
	"awEJrLi0kIDR2sJdAvgXL02BkDbb/JMJT85bqXLYJCBFcLbStuQeUqgqKfq2qTqgncMvMle8kK5vt6m+",
	"FtKt0XZNbrTDd6hyqfCl1Wa7or/eY+aDn9/wcT4kZknR4rdKWhQhXjqw7Slp5XXXD8fVahWP7OKR8xId",
	"/YbEp/RYuknE1gvcWv4UnsnTDX6r0Pl/60tTQDjJlvZqO2HvHoqN4e64pAvHi5wGoL11fchyISw612X5",
	"5PSM/cqlYl+8RfQJ+7kw0iXs5PTs/Ie+7LDksuj6uNdrJTT+WK+8z3Q5TZW/6LVin/VkfcWzk20mfen/",
	"r2Q1sWUdR30JOM995dqNx6AS4ccEeJah8SQ6iwHcjtxaKckSneelCW4mNML/pjZfPTsOSngv+bAk1UoT",
	"UtKT/R9SoGZhNLDfLRfIfrq+hAQe0DqaCXDyfvl+ScQaVNxISOGMlhIw3K8JsgVpklDVzofPgC0P3fdS",
	"QArX2vkL2hLLCp3/qMVTPX48KrLhxhQyI6vFvYsDKQ7y8O17iytI4bvFbtIv6jG/aIbbZhML1xmtXAzo",
	"dHnyZsc0ZyQg0GVWmjg2awwDBiyzyD0K5qosQ+dWVVE8ERWuKktunyCFT7SFcabwkT1sTQOBPHehCCKa",
	"d8EsIrt4lmITohNYoMeXAH+mdYL4MlZMB4JltG3HTIxHd/vRJnC+PB+wUNqzla6U2EsqBsD44YQSyLFH",
	"HhfoD4c+F3sWvZX40ItID4YfuWC1nF+H2gV6xpkzmMmVzMawM9zyEj3asP4MMrgPJQhNXwApoD23vK0w",
	"aaGz3yDuwo2pr1yrDh8z1+usjFdGcD8b37d02liVhLKnKTrSUa/inqNRFG8yR+6p20P2SYrTiHAY6KpT",
	"WBruvH7nv8VBfG5zMLH3RjImNt92bq7aptT040Fxte1GO3F/grR6sBWPZDIf64e68SvB2Wu4hzVwjI4b",
	"dFW50dK+dUetbLoDH7mwmzP2GQ7rb1XQN5hL59HWJR2AbfEYcY61TN+3pTwk/dvo4GiyH8LkqgmEBaN/",
	"cCcjMCeKfQCdee8VJOxj3iuGpb2cR9pvfqM4yHF9owhksvBqF94fg9lYFSyyNVc5vjPcuUdt6e3zrXQg",
	"lcecEBoXwicK47qJ4jiyGNFEF+4mFhYRmovICEQk0uzgGGHRokM/A4mHB9WluAlxdFicDDOlMBPIFOY4",
	"xsEG7UODYWULSGHtvUkXi0JnvFhr59MPyw9LCOjU9s9DTd2xlbbx1JIrnmOJ9Ld5zUc8dpOM2NPrSp99",
	"fG0YtQ93nfjPWm0Xb4Kbu83fAQAA//+a6o0bPRkAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
