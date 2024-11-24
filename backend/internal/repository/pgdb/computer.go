package pgdb

import (
	"access-platform/internal/entity"
	"access-platform/internal/lib/ip-generator"
	"access-platform/pkg/postgres"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var (
	ErrComputerAlreadyTaken = errors.New("computer already taken")
)

type ComputerRepository struct {
	*postgres.Postgres
}

func NewComputerRepository(pg *postgres.Postgres) *ComputerRepository {
	return &ComputerRepository{pg}
}

func (r *ComputerRepository) AddComputer(ctx context.Context, os string, cpu string, ram int) (uuid.UUID, string, error) {
	const op = "repository.computer.AddComputer"

	var computerID uuid.UUID

	ssh := "root@" + ipgen.RandomIPv4()

	query := `INSERT INTO computers_schema.computer(os, cpu, ram, status, ssh) VALUES(@computerOS, @computerCPU, @computerRAM, @computerStatus, @computerSSH) RETURNING id`
	args := pgx.NamedArgs{
		"computerOS":     os,
		"computerCPU":    cpu,
		"computerRAM":    ram,
		"computerStatus": true,
		"computerSSH":    ssh,
	}

	err := r.DB.QueryRow(ctx, query, args).Scan(&computerID)
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("%s: %w", op, err)
	}

	return computerID, ssh, nil
}

func (r *ComputerRepository) GetComputer(ctx context.Context, id uuid.UUID) (entity.Computer, error) {
	const op = "repository.computer.GetComputer"

	var computerOS string
	var computerCPU string
	var computerRAM int
	var computerStatus bool
	var computerSSH string

	query := `SELECT os, cpu, ram, status, ssh FROM computers_schema.computer WHERE id = @computerID`
	args := pgx.NamedArgs{
		"computerID": id,
	}

	err := r.DB.QueryRow(ctx, query, args).Scan(&computerOS, &computerCPU, &computerRAM, &computerStatus, &computerSSH)
	if err != nil {
		return entity.Computer{}, fmt.Errorf("%s: %w", op, err)
	}

	computer := entity.Computer{
		ID:     id,
		OS:     computerOS,
		CPU:    computerCPU,
		RAM:    computerRAM,
		Status: computerStatus,
		SSH:    computerSSH,
	}

	return computer, nil
}

func (r *ComputerRepository) ReserveComputer(ctx context.Context, id uuid.UUID) error {
	const op = "repository.computer.ReserveComputer"

	var computerStatus bool

	query := `SELECT status FROM computers_schema.computer WHERE id = @computerID`
	args := pgx.NamedArgs{
		"computerID": id,
	}

	err := r.DB.QueryRow(ctx, query, args).Scan(&computerStatus)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if computerStatus != true {
		return fmt.Errorf("%s: %w", op, ErrComputerAlreadyTaken)
	}

	query = `UPDATE computers_schema.computer SET status = @computerStatus WHERE id = @computerID`
	args = pgx.NamedArgs{
		"computerStatus": false,
		"computerID":     id,
	}

	_, err = r.DB.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *ComputerRepository) RelieveComputer(ctx context.Context, id uuid.UUID) error {
	const op = "repository.computer.RelieveComputer"

	query := `UPDATE computers_schema.computer SET status = @computerStatus WHERE id = @computerID`
	args := pgx.NamedArgs{
		"computerStatus": true,
		"computerID":     id,
	}

	_, err := r.DB.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *ComputerRepository) GetAllComputers(ctx context.Context) ([]entity.Computer, error) {
	const op = "repository.computer.GetAllComputers"

	// SQL-запрос для получения всех записей из таблицы
	query := `SELECT id, os, cpu, ram, status, ssh FROM computers_schema.computer`

	// Выполнение запроса
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to execute query: %w", op, err)
	}
	defer rows.Close()

	// Срез для хранения результатов
	var computers []entity.Computer

	// Итерация по результатам
	for rows.Next() {
		var computer entity.Computer

		err := rows.Scan(&computer.ID, &computer.OS, &computer.CPU, &computer.RAM, &computer.Status, &computer.SSH)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to scan row: %w", op, err)
		}

		computers = append(computers, computer)
	}

	// Проверка ошибок, возникших во время итерации
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: rows iteration error: %w", op, err)
	}

	return computers, nil
}
