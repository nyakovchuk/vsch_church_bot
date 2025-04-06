package service

import (
	"context"
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/dto"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/repository"
)

const coordPattern = `^(-?\d+\.\d+),?\s+(-?\d+\.\d+)$`

var (
	ErrInvalidCoordinatesFormat         = errors.New("некорректный формат координат")
	ErrInvalidCoordinatesTransformation = errors.New("ошибка преобразования координат")
)

type CoordinatesService interface {
	ParseCoordinates(string) (model.Coordinates, error)
	Save(context.Context, model.Coordinates) (model.Coordinates, error)
	GetCoordinates(context.Context, int) (model.Coordinates, error)
}

type coordinatesService struct {
	repo repository.CoordinatesRepository
}

func NewCoordinatesService(repo repository.CoordinatesRepository) CoordinatesService {
	return &coordinatesService{
		repo: repo,
	}
}

func (c *coordinatesService) Save(ctx context.Context, coords model.Coordinates) (model.Coordinates, error) {
	if err := coords.Validate(); err != nil {
		return model.Coordinates{}, err
	}

	repoDTO := &dto.RepositoryCoordinates{
		Latitude:  coords.Latitude,
		Longitude: coords.Longitude,
		IsOnText:  coords.IsOnText,
	}

	repoCoords, err := c.repo.Save(ctx, repoDTO)
	if err != nil {
		return model.Coordinates{}, err
	}

	return dto.ToModel(repoCoords), nil
}

func (c *coordinatesService) GetCoordinates(ctx context.Context, id int) (model.Coordinates, error) {
	repoCoords, err := c.repo.GetByID(ctx, id)
	if err != nil {
		return model.Coordinates{}, err
	}

	return dto.ToModel(*repoCoords), nil
}

func (c *coordinatesService) ParseCoordinates(text string) (model.Coordinates, error) {
	latStr, lonStr, err := splitCoordinates(text)
	if err != nil {
		return model.Coordinates{}, err
	}

	lat, lon, err := parseFloats(latStr, lonStr)
	if err != nil {
		return model.Coordinates{}, err
	}

	coordinates := dto.CoordinatesToModel(lat, lon)
	if err := coordinates.Validate(); err != nil {
		return model.Coordinates{}, err
	}

	return coordinates, nil
}

func splitCoordinates(input string) (string, string, error) {
	re := regexp.MustCompile(coordPattern)
	matches := re.FindStringSubmatch(strings.TrimSpace(input))

	if len(matches) != 3 {
		return "", "", ErrInvalidCoordinatesFormat
	}

	return matches[1], matches[2], nil
}

func parseFloats(latStr, lonStr string) (float64, float64, error) {
	lat, err1 := strconv.ParseFloat(latStr, 64)
	lon, err2 := strconv.ParseFloat(lonStr, 64)

	if err1 != nil || err2 != nil {
		return 0, 0, ErrInvalidCoordinatesTransformation
	}

	return lat, lon, nil
}
