package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfigSuccess(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")
	content := `{
  "telegram_bot_token": "  token-value  ",
  "authorized_chat_ids": [1, 2, 2, 3]
}`

	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("config dosyası yazılamadı: %v", err)
	}

	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("config yüklenemedi: %v", err)
	}

	if cfg.TelegramBotToken != "token-value" {
		t.Fatalf("token normalize hatası: %q", cfg.TelegramBotToken)
	}

	if len(cfg.AuthorizedChatIDs) != 3 {
		t.Fatalf("id normalize hatası, beklenen 3 kayıt, gelen %d", len(cfg.AuthorizedChatIDs))
	}
}

func TestLoadConfigValidationErrors(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")
	content := `{
  "telegram_bot_token": "",
  "authorized_chat_ids": []
}`

	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("config dosyası yazılamadı: %v", err)
	}

	if _, err := Load(path); err == nil {
		t.Fatalf("hatalı config için doğrulama hatası bekleniyordu")
	}
}
