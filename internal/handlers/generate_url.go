package handlers

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/laysaalves/url-miuda/internal/entity"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func GerarUrlAleatoria() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	length := 5 + r.Intn(6)

	result := make([]byte, length)
	for i := range result {
		switch r.Intn(3) {
		case 0:
			result[i] = byte(r.Intn(10) + '0')
		case 1:
			result[i] = byte(r.Intn(26) + 'A')
		case 2:
			result[i] = byte(r.Intn(26) + 'a')
		}
	}

	return string(result)
}

func EncurtarUrl(db *gorm.DB, UrlOriginal string) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	Link := entity.Link{
		ID:        uuid.New(),
		LinkLong:  UrlOriginal,
		LinkMiudo: GerarUrlAleatoria(),
		CreatedAt: time.Now(),
	}

	if err := db.Create(&Link).Error; err != nil {
		logger.Error("Erro ao salvar link", zap.Error(err))
	} else {
		logger.Info("Link salvo com sucesso", zap.String("LinkMiudo", Link.LinkMiudo))
	}
}
