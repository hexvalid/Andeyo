package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const DefaultPath = "config.json"

// AppConfig, botun çalışma ayarlarını tek dosyada toplar.
type AppConfig struct {
	TelegramBotToken  string  `json:"telegram_bot_token"`
	AuthorizedChatIDs []int64 `json:"authorized_chat_ids"`
}

func Load(path string) (*AppConfig, error) {
	// Dosyayı tek seferde okuyup parse ediyoruz; yapı küçük olduğu için yeterli.
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("config dosyası okunamadı (%s): %w", path, err)
	}

	var cfg AppConfig
	if err := json.Unmarshal(content, &cfg); err != nil {
		return nil, fmt.Errorf("config json parse edilemedi (%s): %w", path, err)
	}

	cfg.normalize()
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *AppConfig) normalize() {
	// Token etrafındaki istemsiz boşlukları temizliyoruz.
	c.TelegramBotToken = strings.TrimSpace(c.TelegramBotToken)

	// Yetkili ID listesini tekrar eden değerlerden arındırıyoruz.
	seen := make(map[int64]struct{}, len(c.AuthorizedChatIDs))
	uniqueIDs := make([]int64, 0, len(c.AuthorizedChatIDs))
	for _, id := range c.AuthorizedChatIDs {
		if _, exists := seen[id]; exists {
			continue
		}
		seen[id] = struct{}{}
		uniqueIDs = append(uniqueIDs, id)
	}
	c.AuthorizedChatIDs = uniqueIDs
}

func (c *AppConfig) validate() error {
	if c.TelegramBotToken == "" {
		return errors.New("telegram_bot_token boş olamaz")
	}
	if len(c.AuthorizedChatIDs) == 0 {
		return errors.New("authorized_chat_ids boş olamaz")
	}

	return nil
}
