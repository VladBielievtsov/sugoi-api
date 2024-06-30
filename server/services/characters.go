package services

import (
	"encoding/json"
	"net/http"
	"sugoi-api/db"
	"sugoi-api/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (s *CharactersService) GetCharacters(name string) ([]types.Character, map[string]string) {
	var characters []types.Character
	var result *gorm.DB

	if name != "" {
		result = db.DB.Where("name ILIKE ?", "%"+name+"%").Find(&characters)
	} else {
		result = db.DB.Find(&characters)
	}

	if result.Error != nil {
		return nil, map[string]string{"msg": "Characters not found"}
	}

	return characters, nil
}
