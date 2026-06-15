package http

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/axizkhan/go_postgresSQL/db/sqlc"
	"github.com/axizkhan/go_postgresSQL/internal/models"
	userService "github.com/axizkhan/go_postgresSQL/internal/service/user"
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
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"success":false,
				"error":"invalid request body",
			},
		)
	}

	if err:=validatorPkg.Validate.Struct(req); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"success":false,
				"error":"invalid request body",
			},
		)
	}

	if req.DOB.After(time.Now()){
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"success":false,
				"error":"invalid request body",
			},
		)
	}

	user,err:=h.service.CreateUser(c.Context(),sqlc.CreateUserParams{Name: req.Name, Dob: pgtype.Date{
	Time:  req.DOB,
	Valid: true,
},},)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"error":   err.Error(),
			},
		)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success":true,
		"data":user,
	})

}

func (h *UserHandler) GetUserById(c *fiber.Ctx)error{
	idParam:=c.Params("id")

	id,err:=strconv.ParseInt(idParam,10,32)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":false,
			"error":"invalid user id",
		},)
	}

	user,err:=h.service.GetUserById(c.Context(),int32(id))

	if err != nil{
		return c.Status(fiber.StatusNotFound).JSON(
			fiber.Map{
				"success": false,
				"error":   err.Error(),
			},
		)
	}


	return c.JSON(
		fiber.Map{
			"success": true,
			"data":    user,
		},
	)
}

func (h *UserHandler) ListUser(c *fiber.Ctx,) error{
	users, err := h.service.ListUsers(
		c.Context(),10,0,
	)

	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"error":   err.Error(),
			},
		)
	}

	return c.JSON(
		fiber.Map{
			"success": true,
			"data":    users,
		},
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

		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"success": false,
				"error":   "invalid user id",
			},
		)
	}

	var req models.UpdateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"success": false,
				"error":   "invalid request body",
			},
		)
	}


	if err := validatorPkg.Validate.Struct(req); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"success": false,
				"error":   err.Error(),
			},
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

		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"error":   err.Error(),
			},
		)
	}

	return c.JSON(
		fiber.Map{
			"success":true,
			"data":user,
		},
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

		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"success": false,
				"error":   "invalid user id",
			},
		)
	}

	err = h.service.DeleteUser(
		c.Context(),
		int32(id),
	)

	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"error":   err.Error(),
			},
		)
	}

	return c.JSON(
		fiber.Map{
			"success": true,
			"message": "user deleted successfully",
		},
	)
}