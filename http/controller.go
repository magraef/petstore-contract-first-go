//go:generate oapi-codegen --config ./.oapi-codegen/oapi-codegen-server.cfg.yaml ../docs/openapi.yaml
//go:generate oapi-codegen --config ./.oapi-codegen/oapi-codegen-types.cfg.yaml ../docs/openapi.yaml

package http

import (
	"context"
)

var _ StrictServerInterface = (*PetstoreController)(nil)

type PetstoreController struct {
}

func NewPetstoreController() PetstoreController {
	return PetstoreController{}
}

func (p PetstoreController) AddPet(ctx context.Context, request AddPetRequestObject) (AddPetResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p PetstoreController) UpdatePet(ctx context.Context, request UpdatePetRequestObject) (UpdatePetResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p PetstoreController) FindPetsByStatus(ctx context.Context, request FindPetsByStatusRequestObject) (FindPetsByStatusResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p PetstoreController) FindPetsByTags(ctx context.Context, request FindPetsByTagsRequestObject) (FindPetsByTagsResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p PetstoreController) DeletePet(ctx context.Context, request DeletePetRequestObject) (DeletePetResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p PetstoreController) GetPetById(ctx context.Context, request GetPetByIdRequestObject) (GetPetByIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p PetstoreController) UpdatePetWithForm(ctx context.Context, request UpdatePetWithFormRequestObject) (UpdatePetWithFormResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p PetstoreController) UploadFile(ctx context.Context, request UploadFileRequestObject) (UploadFileResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p PetstoreController) GetInventory(ctx context.Context, request GetInventoryRequestObject) (GetInventoryResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p PetstoreController) PlaceOrder(ctx context.Context, request PlaceOrderRequestObject) (PlaceOrderResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p PetstoreController) DeleteOrder(ctx context.Context, request DeleteOrderRequestObject) (DeleteOrderResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p PetstoreController) GetOrderById(ctx context.Context, request GetOrderByIdRequestObject) (GetOrderByIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p PetstoreController) CreateUser(ctx context.Context, request CreateUserRequestObject) (CreateUserResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p PetstoreController) CreateUsersWithArrayInput(ctx context.Context, request CreateUsersWithArrayInputRequestObject) (CreateUsersWithArrayInputResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p PetstoreController) CreateUsersWithListInput(ctx context.Context, request CreateUsersWithListInputRequestObject) (CreateUsersWithListInputResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p PetstoreController) LoginUser(ctx context.Context, request LoginUserRequestObject) (LoginUserResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p PetstoreController) LogoutUser(ctx context.Context, request LogoutUserRequestObject) (LogoutUserResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p PetstoreController) DeleteUser(ctx context.Context, request DeleteUserRequestObject) (DeleteUserResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p PetstoreController) GetUserByName(ctx context.Context, request GetUserByNameRequestObject) (GetUserByNameResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p PetstoreController) UpdateUser(ctx context.Context, request UpdateUserRequestObject) (UpdateUserResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
