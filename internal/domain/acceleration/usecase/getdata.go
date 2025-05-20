package usecase

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"time"

	acceleration "github.com/aleksander/Go_API/internal/domain/acceleration/models"
	"github.com/google/uuid"
)

//Return parsed example_acc_data/example_acc_reader.csv

type GetAccelerationDataUseCase struct {
	repo acceleration.Repository
}

func NewGetAccelerationDataUseCase(repo acceleration.Repository) *GetAccelerationDataUseCase {
	return &GetAccelerationDataUseCase{repo: repo}
}

type GetAccelerationDataInput struct {
	ID uuid.UUID
}
type GetAccelerationDataOutput struct {
	Acceleration []*acceleration.Acceleration
	Err          error
}

func readCsvFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil

}

func (uc *GetAccelerationDataUseCase) Execute(input GetAccelerationDataInput) GetAccelerationDataOutput {
	records, err := readCsvFile("example_acc_data/example_acc_reader.csv")

	if err != nil {
		return GetAccelerationDataOutput{Err: err}
	}

	accelerations := make([]*acceleration.Acceleration, 0, len(records)-1)

	// Skip header row, process all data rows
	for i := 1; i < len(records); i++ {
		x, err := strconv.ParseFloat(strings.Replace(records[i][1], ",", ".", 1), 64)
		if err != nil {
			return GetAccelerationDataOutput{Err: err}
		}
		y, err := strconv.ParseFloat(strings.Replace(records[i][2], ",", ".", 1), 64)
		if err != nil {
			return GetAccelerationDataOutput{Err: err}
		}
		z, err := strconv.ParseFloat(strings.Replace(records[i][3], ",", ".", 1), 64)
		if err != nil {
			return GetAccelerationDataOutput{Err: err}
		}

		accelerations = append(accelerations, &acceleration.Acceleration{
			ID:        uuid.New(),
			DeviceID:  input.ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Timestamp: time.Now(),
			X:         x,
			Y:         y,
			Z:         z,
		})
	}

	return GetAccelerationDataOutput{
		Acceleration: accelerations,
		Err:          nil,
	}
}

func (uc *GetAccelerationDataUseCase) GetAllAccelerationData() ([]acceleration.Acceleration, error) {
	accelerations, err := uc.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return accelerations, nil
}
