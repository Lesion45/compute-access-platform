package handlers

import (
	"access-platform/internal/controller/http/response"
	"access-platform/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
)

type AddComputerRequest struct {
	OS    string `json:"os"`
	CPU   string `json:"cpu"`
	RAM   int    `json:"ram"`
	Token string `json:"token"`
}

type AddComputerResponse struct {
	ID     uuid.UUID `json:"id"`
	OS     string    `json:"os"`
	CPU    string    `json:"cpu"`
	RAM    int       `json:"ram"`
	Status bool      `json:"status"`
	SSH    string    `json:"ssh"`
}

func AddComputer(log *zap.Logger, services *service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const op = "handlers.v1.AddComputer"

		var req AddComputerRequest

		err := ctx.BindJSON(&req)
		if err != nil {
			log.Error("failed to decode request body", zap.Error(err), zap.String("op", op))
			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("failed to decode request"))

			return
		}

		log.Info("request body decoded", zap.Any("request", req), zap.String("op", op))

		computer, err := services.ComputeService.AddComputer(ctx, req.OS, req.CPU, req.RAM)
		if err != nil {
			log.Info("failed to add computer", zap.String("op", op))

			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("internal server error"))

			return
		}

		log.Info("computer added", zap.String("op", op))

		ctx.IndentedJSON(http.StatusOK, AddComputerResponse{
			ID:     computer.ID,
			OS:     computer.OS,
			CPU:    computer.CPU,
			RAM:    computer.RAM,
			Status: computer.Status,
			SSH:    computer.SSH,
		})

		return
	}
}

type GetComputerRequest struct {
	ID    uuid.UUID `json:"id"`
	Token string    `json:"token"`
}

type GetComputerResponse struct {
	ID     uuid.UUID `json:"id"`
	OS     string    `json:"os"`
	CPU    string    `json:"cpu"`
	RAM    int       `json:"ram"`
	Status bool      `json:"status"`
	SSH    string    `json:"ssh"`
}

func GetComputer(log *zap.Logger, services *service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const op = "handlers.v1.GetComputer"

		var req GetComputerRequest

		err := ctx.BindJSON(&req)
		if err != nil {
			log.Error("failed to decode request body", zap.Error(err), zap.String("op", op))
			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("failed to decode request"))

			return
		}

		log.Info("request body decoded", zap.Any("request", req), zap.String("op", op))

		computer, err := services.ComputeService.GetComputer(ctx, req.ID)
		if err != nil {
			log.Info("failed to get computer", zap.String("op", op))

			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("internal server error"))

			return
		}

		log.Info("computer found", zap.String("op", op))

		ctx.IndentedJSON(http.StatusOK, GetComputerResponse{
			ID:     computer.ID,
			OS:     computer.OS,
			CPU:    computer.CPU,
			RAM:    computer.RAM,
			Status: computer.Status,
			SSH:    computer.SSH,
		})

		return
	}
}

type ReserveComputerRequest struct {
	ID    uuid.UUID `json:"id"`
	Token string    `json:"token"`
}

type ReserveComputerResponse struct {
	ID       uuid.UUID `json:"id"`
	Reserved bool      `json:"reserved"`
}

func ReserveComputer(log *zap.Logger, services *service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const op = "handlers.v1.ReserveComputer"

		var req ReserveComputerRequest

		err := ctx.BindJSON(&req)
		if err != nil {
			log.Error("failed to decode request body", zap.Error(err), zap.String("op", op))
			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("failed to decode request"))

			return
		}

		log.Info("request body decoded", zap.Any("request", req), zap.String("op", op))

		err = services.ComputeService.ReserveComputer(ctx, req.ID)
		if err != nil {
			log.Info("failed to reserve computer", zap.String("op", op))

			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("internal server error"))

			return
		}

		log.Info("computer found", zap.String("op", op))

		ctx.IndentedJSON(http.StatusOK, ReserveComputerResponse{
			ID:       req.ID,
			Reserved: true,
		})

		return
	}
}

type RelieveComputerRequest struct {
	ID    uuid.UUID `json:"id"`
	Token string    `json:"token"`
}

type RelieveComputerResponse struct {
	ID       uuid.UUID `json:"id"`
	Reserved bool      `json:"reserved"`
}

func RelieveComputer(log *zap.Logger, services *service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const op = "handlers.v1.RelieveComputer"

		var req RelieveComputerRequest

		err := ctx.BindJSON(&req)
		if err != nil {
			log.Error("failed to decode request body", zap.Error(err), zap.String("op", op))
			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("failed to decode request"))

			return
		}

		log.Info("request body decoded", zap.Any("request", req), zap.String("op", op))

		err = services.ComputeService.RelieveComputer(ctx, req.ID)
		if err != nil {
			log.Info("failed to reserve computer", zap.String("op", op))

			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("internal server error"))

			return
		}

		log.Info("computer found", zap.String("op", op))

		ctx.IndentedJSON(http.StatusOK, ReserveComputerResponse{
			ID:       req.ID,
			Reserved: false,
		})

		return
	}
}

type GetAllComputersResponse struct {
	Computers []GetComputerResponse `json:"computers"`
}

func GetAllComputers(log *zap.Logger, services *service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const op = "handlers.v1.GetAllComputers"

		log.Info("attempting to fetch all computers", zap.String("op", op))

		// Вызов метода сервиса для получения всех компьютеров
		computers, err := services.ComputeService.GetAllComputers(ctx)
		if err != nil {
			log.Error("failed to fetch all computers", zap.Error(err), zap.String("op", op))
			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("internal server error"))

			return
		}

		log.Info("successfully fetched all computers", zap.Int("count", len(computers)), zap.String("op", op))

		// Подготовка ответа
		var responseComputers []GetComputerResponse
		for _, computer := range computers {
			responseComputers = append(responseComputers, GetComputerResponse{
				ID:     computer.ID,
				OS:     computer.OS,
				CPU:    computer.CPU,
				RAM:    computer.RAM,
				Status: computer.Status,
				SSH:    computer.SSH,
			})
		}

		// Возвращаем JSON-ответ
		ctx.IndentedJSON(http.StatusOK, GetAllComputersResponse{
			Computers: responseComputers,
		})

		return
	}
}
