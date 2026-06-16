package http

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/axizkhan/go_postgresSQL/db/sqlc"
	"github.com/axizkhan/go_postgresSQL/internal/models"
	userService "github.com/axizkhan/go_postgresSQL/internal/service/user"
	"github.com/axizkhan/go_postgresSQL/pkg/utils"
	validatorPkg "github.com/axizkhan/go_postgresSQL/pkg/validator"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserHandler struct{
	service *userService.Service
}

func NewUserHandler(service *userService.Service) *UserHandler{
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error{
	var req models.CreateUserRequest

	if err:= c.BodyParser(&req); err!=nil{
		return utils.Error(
			c,
			fiber.StatusBadRequest,
			"Invalid request body",
		)
	}

	if err:=validatorPkg.Validate.Struct(req); err!=nil{
		return utils.Error(
			c,
			fiber.StatusBadRequest,
			"Invalid request body",
		)
	}

	if req.DOB.After(time.Now()){
		return utils.Error(
			c,
			fiber.StatusBadRequest,
			"Invalid request body",
		)
	}

	user,err:=h.service.CreateUser(c.Context(),sqlc.CreateUserParams{Name: req.Name, Dob: pgtype.Date{
	Time:  req.DOB,
	Valid: true,
},},)

	if err != nil{
		return utils.Error(
			c,
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}

	return utils.Success(
		c,
		fiber.StatusOK,
		user,
	)

}

func (h *UserHandler) GetUserById(c *fiber.Ctx)error{
	idParam:=c.Params("id")

	id,err:=strconv.ParseInt(idParam,10,32)

	if err != nil {
		return utils.Error(
			c,
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	user,err:=h.service.GetUserById(c.Context(),int32(id))

	if err != nil{
		return utils.Error(
			c,
			fiber.StatusNotFound,
			err.Error(),
		)
	}


	return utils.Success(
		c,
		fiber.StatusOK,
		user,
	)
}

func (h *UserHandler) ListUser(c *fiber.Ctx,) error{

	page,_ := strconv.Atoi(c.Query("page","1"))
	limit,_ := strconv.Atoi(c.Query("limit","10"))

	parsedLimit, offset := utils.GetPagination(page,limit)

	users, err := h.service.ListUsers(
		c.Context(),parsedLimit,offset,
	)

	if err != nil {

		return utils.Error(
			c,
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}

	return utils.Success(
		c,
		fiber.StatusOK,
		users,
	)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx,) error {
	idParam := c.Params("id")

	id, err := strconv.ParseInt(
		idParam,
		10,
		32,
	)


	if err != nil {

		return utils.Error(
			c,
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	var req models.UpdateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.Error(
			c,
			fiber.StatusBadRequest,
			err.Error(),
		)
	}


	if err := validatorPkg.Validate.Struct(req); err != nil {

		return utils.Error(
			c,
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	user, err := h.service.UpdateUser(
		c.Context(),
		sqlc.UpdateUserParams{
			ID: int32(id),
			Name: req.Name,
			Dob: pgtype.Date{
	Time:  req.DOB,
	Valid: true,
},
		},
	)

	
	if err != nil {
return utils.Error(
			c,
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}

	return utils.Success(
		c,
		fiber.StatusOK,
		user,
	)

}

func (h *UserHandler) DeleteUser(
	c *fiber.Ctx,
) error {

	idParam := c.Params("id")

	id, err := strconv.ParseInt(
		idParam,
		10,
		32,
	)

	if err != nil {

		return utils.Error(
			c,
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	err = h.service.DeleteUser(
		c.Context(),
		int32(id),
	)

	if err != nil {

		return utils.Error(
			c,
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}

	return utils.Success(
		c,
		fiber.StatusOK,
		"User deleted",
	)
}