package services

import (
	"image"
	"net/http"
	"os"
	"path/filepath"
	"sugoi-api/db"
	"sugoi-api/types"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/google/uuid"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

type ImagesService struct{}

func NewImagesService() *ImagesService {
	return &ImagesService{}
}

func (s *ImagesService) CreateImage(req *http.Request) (types.Image, map[string]string) {
	err := req.ParseMultipartForm(10 << 20) // 10MB max file size
	if err != nil {
		return types.Image{}, map[string]string{"msg": "Could not parse multipart form"}
	}

	source := req.FormValue("source")
	if source == "" {
		return types.Image{}, map[string]string{"msg": "Image URL and Source are required"}
	}

	file, _, err := req.FormFile("image")
	if err != nil {
		return types.Image{}, map[string]string{"msg": "Could not retrieve the file"}
	}
	defer file.Close()

	uploadDir := "uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.Mkdir(uploadDir, 0755)
		if err != nil {
			return types.Image{}, map[string]string{"msg": "Failed to create 'uploads' directory"}
		}
	}

	id := uuid.New()
	fileName := id.String() + ".webp"
	filePath := filepath.Join(uploadDir, fileName)

	img, _, err := image.Decode(file)
	if err != nil {
		return types.Image{}, map[string]string{"msg": "Invalid image format"}
	}

	out, err := os.Create(filePath)
	if err != nil {
		return types.Image{}, map[string]string{"msg": "Could not create file"}
	}
	defer out.Close()

	options, _ := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 80)
	webp.Encode(out, img, options)

	imageWidth := img.Bounds().Dx()
	imageHeight := img.Bounds().Dy()

	imgInfo, err := out.Stat()
	if err != nil {
		return types.Image{}, map[string]string{"msg": "Failed to get file information"}
	}
	imageSize := int(imgInfo.Size())

	image := types.Image{
		ID:          &id,
		ImageURL:    "/" + filePath,
		Source:      source,
		ImageSize:   imageSize,
		ImageWidth:  imageWidth,
		ImageHeight: imageHeight,
	}

	if err = db.DB.Create(&image).Error; err != nil {
		return types.Image{}, map[string]string{"msg": "Could not creaet an image"}
	}

	return image, nil
}

func (s *ImagesService) GetImages() ([]types.Image, map[string]string) {
	var images []types.Image

	result := db.DB.Find(&images)
	if result.Error != nil {
		return nil, map[string]string{"msg": "Images not found"}
	}

	return images, nil
}

func (s *ImagesService) GetImageByID(id string) (types.Image, map[string]string) {
	var image types.Image

	result := db.DB.First(&image, "id = ?", id)
	if result.Error != nil {
		return types.Image{}, map[string]string{"msg": "Image not found"}
	}

	return image, nil
}

func (s *ImagesService) GetRandomImages(limit int) ([]types.Image, map[string]string) {
	var images []types.Image

	result := db.DB.Order("RANDOM()").Limit(limit).Find(&images)
	if result.Error != nil {
		return nil, map[string]string{"msg": "Images not found"}
	}

	return images, nil
}
