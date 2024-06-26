// Package http provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// returns all existing pet matching optional filters
	// (GET /pets)
	GetPets(w http.ResponseWriter, r *http.Request, params GetPetsParams)
	// Add a new pet to the store
	// (POST /pets)
	AddPet(w http.ResponseWriter, r *http.Request)
	// Deletes a pet
	// (DELETE /pets/{petId})
	DeletePet(w http.ResponseWriter, r *http.Request, petId int64)
	// Find pet by ID
	// (GET /pets/{petId})
	GetPetById(w http.ResponseWriter, r *http.Request, petId int64)
	// Updates a pet in the store
	// (PUT /pets/{petId})
	UpdatePet(w http.ResponseWriter, r *http.Request, petId int64)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// returns all existing pet matching optional filters
// (GET /pets)
func (_ Unimplemented) GetPets(w http.ResponseWriter, r *http.Request, params GetPetsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add a new pet to the store
// (POST /pets)
func (_ Unimplemented) AddPet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Deletes a pet
// (DELETE /pets/{petId})
func (_ Unimplemented) DeletePet(w http.ResponseWriter, r *http.Request, petId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Find pet by ID
// (GET /pets/{petId})
func (_ Unimplemented) GetPetById(w http.ResponseWriter, r *http.Request, petId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Updates a pet in the store
// (PUT /pets/{petId})
func (_ Unimplemented) UpdatePet(w http.ResponseWriter, r *http.Request, petId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetPets operation middleware
func (siw *ServerInterfaceWrapper) GetPets(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPetsParams

	// ------------- Optional query parameter "category" -------------

	err = runtime.BindQueryParameter("form", true, false, "category", r.URL.Query(), &params.Category)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "category", Err: err})
		return
	}

	// ------------- Optional query parameter "from" -------------

	err = runtime.BindQueryParameter("form", true, false, "from", r.URL.Query(), &params.From)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "from", Err: err})
		return
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetPets(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddPet operation middleware
func (siw *ServerInterfaceWrapper) AddPet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddPet(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeletePet operation middleware
func (siw *ServerInterfaceWrapper) DeletePet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "petId" -------------
	var petId int64

	err = runtime.BindStyledParameterWithOptions("simple", "petId", chi.URLParam(r, "petId"), &petId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "petId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeletePet(w, r, petId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetPetById operation middleware
func (siw *ServerInterfaceWrapper) GetPetById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "petId" -------------
	var petId int64

	err = runtime.BindStyledParameterWithOptions("simple", "petId", chi.URLParam(r, "petId"), &petId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "petId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetPetById(w, r, petId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdatePet operation middleware
func (siw *ServerInterfaceWrapper) UpdatePet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "petId" -------------
	var petId int64

	err = runtime.BindStyledParameterWithOptions("simple", "petId", chi.URLParam(r, "petId"), &petId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "petId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdatePet(w, r, petId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/pets", wrapper.GetPets)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/pets", wrapper.AddPet)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/pets/{petId}", wrapper.DeletePet)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/pets/{petId}", wrapper.GetPetById)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/pets/{petId}", wrapper.UpdatePet)
	})

	return r
}

type GetPetsRequestObject struct {
	Params GetPetsParams
}

type GetPetsResponseObject interface {
	VisitGetPetsResponse(w http.ResponseWriter) error
}

type GetPets200JSONResponse []Pet

func (response GetPets200JSONResponse) VisitGetPetsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetPets400ApplicationProblemPlusJSONResponse Problem

func (response GetPets400ApplicationProblemPlusJSONResponse) VisitGetPetsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type AddPetRequestObject struct {
	Body *AddPetJSONRequestBody
}

type AddPetResponseObject interface {
	VisitAddPetResponse(w http.ResponseWriter) error
}

type AddPet201JSONResponse Pet

func (response AddPet201JSONResponse) VisitAddPetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type AddPet400ApplicationProblemPlusJSONResponse Problem

func (response AddPet400ApplicationProblemPlusJSONResponse) VisitAddPetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type AddPet404ApplicationProblemPlusJSONResponse Problem

func (response AddPet404ApplicationProblemPlusJSONResponse) VisitAddPetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type AddPet409ApplicationProblemPlusJSONResponse Problem

func (response AddPet409ApplicationProblemPlusJSONResponse) VisitAddPetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type DeletePetRequestObject struct {
	PetId int64 `json:"petId"`
}

type DeletePetResponseObject interface {
	VisitDeletePetResponse(w http.ResponseWriter) error
}

type DeletePet204Response struct {
}

func (response DeletePet204Response) VisitDeletePetResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeletePet400JSONResponse Problem

func (response DeletePet400JSONResponse) VisitDeletePetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type DeletePet404ApplicationProblemPlusJSONResponse Problem

func (response DeletePet404ApplicationProblemPlusJSONResponse) VisitDeletePetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetPetByIdRequestObject struct {
	PetId int64 `json:"petId"`
}

type GetPetByIdResponseObject interface {
	VisitGetPetByIdResponse(w http.ResponseWriter) error
}

type GetPetById200JSONResponse Pet

func (response GetPetById200JSONResponse) VisitGetPetByIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetPetById400ApplicationProblemPlusJSONResponse Problem

func (response GetPetById400ApplicationProblemPlusJSONResponse) VisitGetPetByIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetPetById404ApplicationProblemPlusJSONResponse Problem

func (response GetPetById404ApplicationProblemPlusJSONResponse) VisitGetPetByIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type UpdatePetRequestObject struct {
	PetId int64 `json:"petId"`
	Body  *UpdatePetJSONRequestBody
}

type UpdatePetResponseObject interface {
	VisitUpdatePetResponse(w http.ResponseWriter) error
}

type UpdatePet202Response struct {
}

func (response UpdatePet202Response) VisitUpdatePetResponse(w http.ResponseWriter) error {
	w.WriteHeader(202)
	return nil
}

type UpdatePet400ApplicationProblemPlusJSONResponse Problem

func (response UpdatePet400ApplicationProblemPlusJSONResponse) VisitUpdatePetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type UpdatePet404ApplicationProblemPlusJSONResponse Problem

func (response UpdatePet404ApplicationProblemPlusJSONResponse) VisitUpdatePetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type UpdatePet409ApplicationProblemPlusJSONResponse Problem

func (response UpdatePet409ApplicationProblemPlusJSONResponse) VisitUpdatePetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// returns all existing pet matching optional filters
	// (GET /pets)
	GetPets(ctx context.Context, request GetPetsRequestObject) (GetPetsResponseObject, error)
	// Add a new pet to the store
	// (POST /pets)
	AddPet(ctx context.Context, request AddPetRequestObject) (AddPetResponseObject, error)
	// Deletes a pet
	// (DELETE /pets/{petId})
	DeletePet(ctx context.Context, request DeletePetRequestObject) (DeletePetResponseObject, error)
	// Find pet by ID
	// (GET /pets/{petId})
	GetPetById(ctx context.Context, request GetPetByIdRequestObject) (GetPetByIdResponseObject, error)
	// Updates a pet in the store
	// (PUT /pets/{petId})
	UpdatePet(ctx context.Context, request UpdatePetRequestObject) (UpdatePetResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// GetPets operation middleware
func (sh *strictHandler) GetPets(w http.ResponseWriter, r *http.Request, params GetPetsParams) {
	var request GetPetsRequestObject

	request.Params = params

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetPets(ctx, request.(GetPetsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetPets")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetPetsResponseObject); ok {
		if err := validResponse.VisitGetPetsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// AddPet operation middleware
func (sh *strictHandler) AddPet(w http.ResponseWriter, r *http.Request) {
	var request AddPetRequestObject

	var body AddPetJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.AddPet(ctx, request.(AddPetRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AddPet")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(AddPetResponseObject); ok {
		if err := validResponse.VisitAddPetResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeletePet operation middleware
func (sh *strictHandler) DeletePet(w http.ResponseWriter, r *http.Request, petId int64) {
	var request DeletePetRequestObject

	request.PetId = petId

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.DeletePet(ctx, request.(DeletePetRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeletePet")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DeletePetResponseObject); ok {
		if err := validResponse.VisitDeletePetResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetPetById operation middleware
func (sh *strictHandler) GetPetById(w http.ResponseWriter, r *http.Request, petId int64) {
	var request GetPetByIdRequestObject

	request.PetId = petId

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetPetById(ctx, request.(GetPetByIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetPetById")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetPetByIdResponseObject); ok {
		if err := validResponse.VisitGetPetByIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// UpdatePet operation middleware
func (sh *strictHandler) UpdatePet(w http.ResponseWriter, r *http.Request, petId int64) {
	var request UpdatePetRequestObject

	request.PetId = petId

	var body UpdatePetJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.UpdatePet(ctx, request.(UpdatePetRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdatePet")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(UpdatePetResponseObject); ok {
		if err := validResponse.VisitUpdatePetResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}
