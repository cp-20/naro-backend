package main

import (
	"database/sql"
	"testing"
)

func Test_calculatePopulation_empty(t *testing.T) {
	// ここにテストを書いていく
	cities := []City{}
	got := calculatePopulation(cities)
	want := map[string]int{}
	// 長さが0になっているかどうかを確認する
	if len(got) != 0 {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
}

func Test_calculatePopulation_one(t *testing.T) {
	// ここにテストを書いていく
	cities := []City{
		{
			ID:          1,
			Name:        sql.NullString{String: "Kabul", Valid: true},
			CountryCode: sql.NullString{String: "AFG", Valid: true},
			District:    sql.NullString{String: "Kabol", Valid: true},
			Population:  sql.NullInt64{Int64: 1780000, Valid: true},
		},
		{
			ID:          2,
			Name:        sql.NullString{String: "Qandahar", Valid: true},
			CountryCode: sql.NullString{String: "AFG", Valid: true},
			District:    sql.NullString{String: "Qandahar", Valid: true},
			Population:  sql.NullInt64{Int64: 237500, Valid: true},
		},
		{
			ID:          3,
			Name:        sql.NullString{String: "Herat", Valid: true},
			CountryCode: sql.NullString{String: "AFG", Valid: true},
			District:    sql.NullString{String: "Herat", Valid: true},
			Population:  sql.NullInt64{Int64: 186800, Valid: true},
		},
	}
	got := calculatePopulation(cities)
	want := map[string]int{
		"AFG": 1780000 + 237500 + 186800,
	}
	if len(got) != 1 || got["AFG"] != want["AFG"] {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
}

func Test_calculatePopulation_many(t *testing.T) {
	// ここにテストを書いていく
	cities := []City{
		{
			ID:          1,
			Name:        sql.NullString{String: "Kabul", Valid: true},
			CountryCode: sql.NullString{String: "AFG", Valid: true},
			District:    sql.NullString{String: "Kabol", Valid: true},
			Population:  sql.NullInt64{Int64: 1780000, Valid: true},
		},
		{
			ID:          2,
			Name:        sql.NullString{String: "Qandahar", Valid: true},
			CountryCode: sql.NullString{String: "AFG", Valid: true},
			District:    sql.NullString{String: "Qandahar", Valid: true},
			Population:  sql.NullInt64{Int64: 237500, Valid: true},
		},
		{
			ID:          3,
			Name:        sql.NullString{String: "Herat", Valid: true},
			CountryCode: sql.NullString{String: "AFG", Valid: true},
			District:    sql.NullString{String: "Herat", Valid: true},
			Population:  sql.NullInt64{Int64: 186800, Valid: true},
		},
		{
			ID:          11,
			Name:        sql.NullString{String: "Buenos Aires", Valid: true},
			CountryCode: sql.NullString{String: "ARG", Valid: true},
			District:    sql.NullString{String: "Distrito Federal", Valid: true},
			Population:  sql.NullInt64{Int64: 2982146, Valid: true},
		},
		{
			ID:          12,
			Name:        sql.NullString{String: "La Matanza", Valid: true},
			CountryCode: sql.NullString{String: "ARG", Valid: true},
			District:    sql.NullString{String: "Buenos Aires", Valid: true},
			Population:  sql.NullInt64{Int64: 1266461, Valid: true},
		},
		{
			ID:          13,
			Name:        sql.NullString{String: "Córdoba", Valid: true},
			CountryCode: sql.NullString{String: "ARG", Valid: true},
			District:    sql.NullString{String: "Córdoba", Valid: true},
			Population:  sql.NullInt64{Int64: 1157507, Valid: true},
		},
		{
			ID:          21,
			Name:        sql.NullString{String: "Sydney", Valid: true},
			CountryCode: sql.NullString{String: "AUS", Valid: true},
			District:    sql.NullString{String: "New South Wales", Valid: true},
			Population:  sql.NullInt64{Int64: 3276207, Valid: true},
		},
		{
			ID:          22,
			Name:        sql.NullString{String: "Melbourne", Valid: true},
			CountryCode: sql.NullString{String: "AUS", Valid: true},
			District:    sql.NullString{String: "Victoria", Valid: true},
			Population:  sql.NullInt64{Int64: 2865329, Valid: true},
		},
		{
			ID:          23,
			Name:        sql.NullString{String: "Brisbane", Valid: true},
			CountryCode: sql.NullString{String: "AUS", Valid: true},
			District:    sql.NullString{String: "Queensland", Valid: true},
			Population:  sql.NullInt64{Int64: 1291117, Valid: true},
		},
	}
	got := calculatePopulation(cities)
	want := map[string]int{
		"AFG": 1780000 + 237500 + 186800,
		"ARG": 2982146 + 1266461 + 1157507,
		"AUS": 3276207 + 2865329 + 1291117,
	}
	if len(got) != 3 || got["AFG"] != want["AFG"] || got["ARG"] != want["ARG"] || got["AUS"] != want["AUS"] {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
}

func Test_calculatePopulation_not_valid(t *testing.T) {
	// ここにテストを書いていく
	cities := []City{
		{
			ID:          1,
			Name:        sql.NullString{String: "Kabul", Valid: true},
			CountryCode: sql.NullString{String: "AFG", Valid: true},
			District:    sql.NullString{String: "Kabol", Valid: true},
			Population:  sql.NullInt64{Int64: 1780000, Valid: true},
		},
		{
			ID:          2,
			Name:        sql.NullString{String: "Qandahar", Valid: true},
			CountryCode: sql.NullString{String: "AFG", Valid: true},
			District:    sql.NullString{String: "Qandahar", Valid: true},
			Population:  sql.NullInt64{Int64: 237500, Valid: true},
		},
		{
			ID:          3,
			Name:        sql.NullString{String: "Herat", Valid: true},
			CountryCode: sql.NullString{String: "AFG", Valid: true},
			District:    sql.NullString{String: "Herat", Valid: true},
			Population:  sql.NullInt64{Int64: 186800, Valid: true},
		},
		{
			ID:          4,
			Name:        sql.NullString{String: "Buenos Aires", Valid: true},
			CountryCode: sql.NullString{String: "ARG", Valid: false},
			District:    sql.NullString{String: "Distrito Federal", Valid: true},
			Population:  sql.NullInt64{Int64: 2982146, Valid: true},
		},
	}
	got := calculatePopulation(cities)
	want := map[string]int{
		"AFG": 1780000 + 237500 + 186800,
	}
	// 長さが0になっているかどうかを確認する
	if len(got) != 1 || got["AFG"] != want["AFG"] {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
}
