package service

import (
	"context"
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/nyakovchuk/vsch_church_bot/internal/apperrors"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/dto"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/repository"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/external"
)

const coordPattern = `^(-?\d+\.\d+),?\s+(-?\d+\.\d+)$`

var (
	ErrInvalidCoordinatesFormat         = errors.New("некорректный формат координат")
	ErrInvalidCoordinatesTransformation = errors.New("ошибка преобразования координат")
)

type CoordinatesService interface {
	ParseCoordinates(string) (model.Coordinates, error)
	Save(context.Context, model.Coordinates) (model.Coordinates, error)
	GetCoordinates(context.Context, external.External) (model.Coordinates, error)
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
		PlatformID: coords.PlatformID,
		ExternalID: coords.ExternalID,
		Latitude:   coords.Latitude,
		Longitude:  coords.Longitude,
		IsOnText:   coords.IsOnText,
	}

	repoCoords, err := c.repo.Save(ctx, repoDTO)
	if err != nil {
		return model.Coordinates{}, apperrors.Wrap(apperrors.ErrSaveCoordinates, err)
	}

	return repoCoords.ToModel(), nil
}

func (c *coordinatesService) GetCoordinates(ctx context.Context, external external.External) (model.Coordinates, error) {

	repoExternal := external.ToRepository()
	repoCoords, err := c.repo.GetCoordinatesByExternal(ctx, repoExternal)
	if err != nil {
		return model.Coordinates{}, err
	}

	return repoCoords.ToModel(), nil
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

	coordinates := model.GeoToModel(lat, lon)
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
