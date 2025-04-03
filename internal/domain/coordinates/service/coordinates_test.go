package service_test

import (
	"testing"

	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/service"
	"github.com/stretchr/testify/assert"
)

// Структура теста
type testCase struct {
	input       string
	wantSuccess bool
	wantErr     error
}

func TestParseCoordinates(t *testing.T) {
	// Таблица тестов
	tests := []testCase{
		{"55.7558, 37.6173", true, nil},   // Обычный случай
		{"-12.3456   98.7654", true, nil}, // Минус, без запятой, несколько пробелов
		{"45.0, -93.0", true, nil},        // Отрицательная долгота
		{"0.0 0.0", true, nil},            // Нулевые координаты
		{"  22.22, 33.33  ", true, nil},   // Лишние пробелы вокруг

		// Негативные тесты
		{"55.7558,37.6173,100", false, service.ErrInvalidCoordinatesFormat}, // Лишнее значение
		{"55.7558-37.6173", false, service.ErrInvalidCoordinatesFormat},     // Без пробела или запятой
		{"abc, def", false, service.ErrInvalidCoordinatesFormat},            // Некорректные данные
		{"", false, service.ErrInvalidCoordinatesFormat},                    // Пустая строка
		{"123.45", false, service.ErrInvalidCoordinatesFormat},              // Только одно число
	}

	service := service.NewCoordinatesService(nil)

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			gotCoordinates, err := service.ParseCoordinates(tc.input)

			if tc.wantSuccess && err != nil {
				t.Errorf("ParseCoordinates(%q) = (%f, %f),  error = %v", tc.input, gotCoordinates.Latitude, gotCoordinates.Longitude, err)
			}

			if !tc.wantSuccess {
				assert.ErrorIs(t, err, tc.wantErr)
			}
		})
	}
}
