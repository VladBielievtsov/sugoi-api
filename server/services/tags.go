package services

import (
	"encoding/json"
	"net/http"
	"sugoi-api/db"
	"sugoi-api/types"

	"github.com/google/uuid"
)

type TagsService struct{}

func NewTagsService() *TagsService {
	return &TagsService{}
}

func (s *TagsService) CreateTag(r *http.Request) (types.Tag, map[string]string) {
	var req types.CreateTagBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return types.Tag{}, map[string]string{"msg": "Invalid request payload"}
	}

	id := uuid.New()

	tag := types.Tag{
		ID:          &id,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := db.DB.Create(&tag).Error; err != nil {
		return types.Tag{}, map[string]string{"msg": "Could not creaet an tag"}
	}

	return tag, nil
}

func (s *TagsService) GetTagByName(name string) (types.Tag, map[string]string) {
	var tag types.Tag

	result := db.DB.First(&tag, "name = ?", name)
	if result.Error != nil {
		return types.Tag{}, map[string]string{"msg": "Tag not found"}
	}

	return tag, nil
}

func (s *TagsService) GetTags() ([]types.Tag, map[string]string) {
	var tags []types.Tag

	result := db.DB.Find(&tags)
	if result.Error != nil {
		return []types.Tag{}, map[string]string{"msg": "Tags not found"}
	}

	return tags, nil
}

func (s *TagsService) GetOrCreateTags(tagNames []string) ([]types.Tag, map[string]string) {
	var tags []types.Tag
	for _, tagName := range tagNames {
		var tag types.Tag
		if err := db.DB.Where("name = ?", tagName).First(&tag).Error; err != nil {
			id := uuid.New()
			if err := db.DB.Create(&types.Tag{ID: &id, Name: tagName, Description: "NULL"}).Error; err != nil {
				return nil, map[string]string{"msg": "Failed to create tag"}
			}
			db.DB.Where("name = ?", tagName).First(&tag)
		}
		tags = append(tags, tag)
	}
	return tags, nil
}
