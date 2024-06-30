package services

import (
	"encoding/json"
	"net/http"
	"strings"
	"sugoi-api/db"
	"sugoi-api/types"

	"github.com/google/uuid"
)

type CharactersService struct{}

func NewCharactersService() *CharactersService {
	return &CharactersService{}
}

func (s *CharactersService) CreateCharacter(r *http.Request) (types.Character, map[string]string) {
	var req types.CreateCharacterBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return types.Character{}, map[string]string{"msg": "Invalid request payload"}
	}

	id := uuid.New()

	character := types.Character{
		ID:          &id,
		Name:        req.Name,
		Description: req.Description,
		Gender:      req.Gender,
		Species:     req.Species,
	}

	if err := db.DB.Create(&character).Error; err != nil {
		return types.Character{}, map[string]string{"msg": "Could not creaet character"}
	}

	return character, nil
}

func (s *CharactersService) GetCharacters(name, gender, species string) ([]types.Character, map[string]string) {
	var characters []types.Character
	query := db.DB

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	if gender != "" {
		query = query.Where("LOWER(gender) = ?", strings.ToLower(gender))
	}

	if species != "" {
		query = query.Where("LOWER(species) = ?", strings.ToLower(species))
	}

	result := query.Find(&characters)

	if result.Error != nil {
		return nil, map[string]string{"msg": "Characters not found"}
	}

	return characters, nil
}

func (s *CharactersService) GetCharactersByNames(names []string) ([]types.Character, map[string]string) {
	var characters []types.Character
	for _, name := range names {
		var character types.Character
		if err := db.DB.Where("name = ?", name).First(&character).Error; err != nil {
			return nil, map[string]string{"msg": "Failed to find character"}
		}
		characters = append(characters, character)
	}
	return characters, nil
}
