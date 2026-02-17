# Andeyo

Andeyo botu, `github.com/go-telegram/bot` kullanılarak Go ile yeniden yazıldı.

## Davranış

- Sadece config dosyasında tanımlı yetkili chat ID'leri için metin dönüşümü yapar.
- Mesajın ilk karakteri Latin alfabesindeyse metni emoji alfabesine çevirir.
- Mesajın ilk karakteri Latin alfabesinde değilse emoji alfabesini Latin metne geri çözer.
- Yetkisiz chat'lerde sabit bir uyarı mesajı döner.

## Config

```bash
cp config.example.json config.json
```

```json
{
  "telegram_bot_token": "YOUR_BOT_TOKEN_HERE",
  "authorized_chat_ids": [81329453, 119941223]
}
```

## Çalıştırma

```bash
go run ./cmd/andeyo
```

## Test

```bash
go test ./...
```

## Notlar

- Token ve yetkili chat listesi `config.json` dosyasından okunur.
- `config.json` ve `.idea/` versiyon kontrolüne dahil edilmez.

## Credits

```text
========================================
   A N D E Y O   P R O J E S I
========================================
```

- Original concept and Java version: [vyagmur92](https://instagram.com/vyagmur92)
- Go rewrite and modernization: Andeyo contributors
