package handler

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type GameApi struct {
	Users  map[uuid.UUID]User
	Games  map[uuid.UUID]Game
	Offers map[uuid.UUID]Offer
}

var _ StrictServerInterface = (*GameApi)(nil)

func NewGameApi() *GameApi {
	return &GameApi{
		Users:  make(map[uuid.UUID]User),
		Games:  make(map[uuid.UUID]Game),
		Offers: make(map[uuid.UUID]Offer),
	}
}

// DeleteGamesId implements StrictServerInterface.
func (a *GameApi) DeleteGamesId(ctx context.Context, request DeleteGamesIdRequestObject) (DeleteGamesIdResponseObject, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return DeleteGamesId404Response{}, err
	}
	delete(a.Games, id)
	return DeleteGamesId200Response{}, nil
}

// DeleteOffersId implements StrictServerInterface.
func (a *GameApi) DeleteOffersId(ctx context.Context, request DeleteOffersIdRequestObject) (DeleteOffersIdResponseObject, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return DeleteOffersId404Response{}, err
	}
	delete(a.Offers, id)
	return DeleteOffersId200Response{}, nil
}

// GetGamesId implements StrictServerInterface.
func (a *GameApi) GetGamesId(ctx context.Context, request GetGamesIdRequestObject) (GetGamesIdResponseObject, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return GetGamesId404Response{}, err
	}
	result := a.Games[id]
	return GetGamesId200JSONResponse(result), nil
}

// GetOffersId implements StrictServerInterface.
func (a *GameApi) GetOffersId(ctx context.Context, request GetOffersIdRequestObject) (GetOffersIdResponseObject, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return GetOffersId404Response{}, err
	}
	result := a.Offers[id]
	return GetOffersId200JSONResponse(result), nil
}

// GetUser implements StrictServerInterface.
func (a *GameApi) GetUser(ctx context.Context, request GetUserRequestObject) (GetUserResponseObject, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return GetUser404Response{}, err
	}
	result := a.Users[id]
	return GetUser200JSONResponse(result), nil
}

// PostGames implements StrictServerInterface.
func (a *GameApi) PostGames(ctx context.Context, request PostGamesRequestObject) (PostGamesResponseObject, error) {
	id := uuid.New()
	newGame := Game{
		Id:        &id,
		Name:      &request.Body.Name,
		Condition: (*GameCondition)(&request.Body.Condition),
		Publisher: &request.Body.Publisher,
	}
	a.Games[id] = newGame
	return PostGames201JSONResponse(newGame), nil
}

// PostOffers implements StrictServerInterface.
func (a *GameApi) PostOffers(ctx context.Context, request PostOffersRequestObject) (PostOffersResponseObject, error) {
	id := uuid.New()
	timeStamp := time.Now().Format(time.Stamp)
	status := Pending
	newOffer := Offer{
		Id:             &id,
		GamesOffered:   &request.Body.GamesOffered,
		GamesRequested: &request.Body.GamesRequested,
		Offeror:        &request.Body.Offeror,
		Offeree:        &request.Body.Offeree,
		Timestamp:      &timeStamp,
		Status:         &status,
	}
	a.Offers[id] = newOffer
	return PostOffers201JSONResponse(newOffer), nil
}

// PostUsers implements StrictServerInterface.
func (a *GameApi) PostUsers(ctx context.Context, request PostUsersRequestObject) (PostUsersResponseObject, error) {
	id := uuid.New()
	newUser := User{
		Id:      &id,
		Name:    &request.Body.Name,
		Email:   &request.Body.Email,
		Address: &request.Body.Address,
	}
	a.Users[id] = newUser
	return PostUsers201JSONResponse(newUser), nil
}

// PostUsersIdResetPassword implements StrictServerInterface.
func (*GameApi) PostUsersIdResetPassword(ctx context.Context, request PostUsersIdResetPasswordRequestObject) (PostUsersIdResetPasswordResponseObject, error) {
	panic("unimplemented")
}

// PutGamesId implements StrictServerInterface.
func (a *GameApi) PutGamesId(ctx context.Context, request PutGamesIdRequestObject) (PutGamesIdResponseObject, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return PutGamesId404Response{}, err
	}
	updatedGame := Game{
		Id:        &id,
		Name:      &request.Body.Name,
		Condition: (*GameCondition)(&request.Body.Condition),
		Publisher: &request.Body.Publisher,
	}
	a.Games[id] = updatedGame
	return PutGamesId200JSONResponse(updatedGame), nil
}

// PutUsersId implements StrictServerInterface.
func (a *GameApi) PutUsersId(ctx context.Context, request PutUsersIdRequestObject) (PutUsersIdResponseObject, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return PutUsersId404Response{}, err
	}
	updatedUser := User{
		Id:      &id,
		Name:    &request.Body.Name,
		Email:   &request.Body.Email,
		Address: &request.Body.Address,
	}
	a.Users[id] = updatedUser
	return PutUsersId200JSONResponse(updatedUser), nil
}

// PutUsersIdChangePassword implements StrictServerInterface.
func (*GameApi) PutUsersIdChangePassword(ctx context.Context, request PutUsersIdChangePasswordRequestObject) (PutUsersIdChangePasswordResponseObject, error) {
	panic("unimplemented")
}
