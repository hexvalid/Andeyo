package translator

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

const UnauthorizedMessage = "Meraklı taze, burada ne işin var? Hadi yaylanda boyunu görelim...."

// Service, eski Java sürümündeki çeviri davranışını tek bir yerde tutar.
type Service struct {
	authorizedChats map[int64]struct{}
	latinAlphabet   map[rune]struct{}
	encodeReplacer  *strings.Replacer
	decodeReplacer  *strings.Replacer
}

func NewService(authorizedChatIDs []int64) *Service {
	// Yetkili chat listesini map'e çevirerek O(1) kontrol elde ediyoruz.
	authorizedChats := make(map[int64]struct{}, len(authorizedChatIDs))
	for _, chatID := range authorizedChatIDs {
		authorizedChats[chatID] = struct{}{}
	}

	// Replacer yapıları başlangıçta bir kez hazırlanır; her mesajda tekrar kurmayız.
	return &Service{
		authorizedChats: authorizedChats,
		latinAlphabet:   buildLatinAlphabetSet(),
		encodeReplacer:  strings.NewReplacer(encodePairs()...),
		decodeReplacer:  strings.NewReplacer(decodePairs()...),
	}
}

func (s *Service) BuildReply(chatID int64, input string) string {
	// Önce yetki kontrolü yapıp eski bottaki erişim kuralını birebir koruyoruz.
	if !s.isAuthorized(chatID) {
		return UnauthorizedMessage
	}

	// İlk karakter Latin alfabesindeyse emoji alfabesine çevir, değilse geri çöz.
	if s.shouldEncode(input) {
		return s.encode(input)
	}

	return s.decode(input)
}

func (s *Service) IsAuthorized(chatID int64) bool {
	// Handler tarafında log/izleme gibi ihtiyaçlar için dışarıya güvenli yetki kontrolü açıyoruz.
	return s.isAuthorized(chatID)
}

func (s *Service) isAuthorized(chatID int64) bool {
	// Eski bottaki izinli kullanıcı mantığını birebir sürdürmek için doğrudan ID karşılaştırıyoruz.
	_, ok := s.authorizedChats[chatID]
	return ok
}

func (s *Service) shouldEncode(input string) bool {
	// Karar mekanizması sadece ilk karaktere bakar; bu eski Java davranışının aynısıdır.
	firstRune, ok := firstLowerRune(input)
	if !ok {
		// Boş mesajı kodlamaya yönlendiriyoruz; iki yol da boş döneceği için davranış güvenli kalır.
		return true
	}

	_, exists := s.latinAlphabet[firstRune]
	return exists
}

func (s *Service) encode(input string) string {
	// Java sürümündeki gibi tüm metni önce küçük harfe çekiyoruz.
	return s.encodeReplacer.Replace(strings.ToLower(input))
}

func (s *Service) decode(input string) string {
	// Emoji alfabesinden Latin'e tek geçişte döndürme işlemi yapıyoruz.
	return s.decodeReplacer.Replace(input)
}

func firstLowerRune(value string) (rune, bool) {
	// UTF-8 uyumlu ilk karakter okuması, çok baytlı emoji/harfleri güvenli ele alır.
	firstRune, size := utf8.DecodeRuneInString(value)
	if size == 0 {
		return 0, false
	}

	// Büyük/küçük harf farkını ortadan kaldırıp karar mekanizmasını sadeleştiriyoruz.
	return unicode.ToLower(firstRune), true
}

func buildLatinAlphabetSet() map[rune]struct{} {
	// Eski koddaki harf listesini koruyoruz; karar mekanizması bunun üstünden çalışıyor.
	letters := []rune{'a', 'b', 'c', 'ç', 'd', 'e', 'f', 'g', 'ğ', 'h', 'ı', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'ö', 'p', 'r', 's', 'ş', 't', 'u', 'ü', 'v', 'y', 'z'}
	set := make(map[rune]struct{}, len(letters))
	for _, letter := range letters {
		set[letter] = struct{}{}
	}

	return set
}

func encodePairs() []string {
	// Harf -> emoji eşlemesi eski sürümle aynı bırakıldı.
	return []string{
		"a", "🎄",
		"b", "🔩",
		"c", "🌜",
		"ç", "💭",
		"d", "🌛",
		"e", "📛",
		"f", "🎏",
		"g", "🐉",
		"ğ", "🐌",
		"h", "⛄",
		"ı", "📏",
		"i", "✏",
		"j", "🎷",
		"k", "🎋",
		"l", "🕒",
		"m", "👓",
		"n", "👠",
		"o", "📯",
		"ö", "🌞",
		"p", "🎧",
		"r", "💃",
		"s", "🐍",
		"ş", "🐢",
		"t", "☔",
		"u", "🔧",
		"ü", "🍇",
		"v", "☑",
		"y", "🎌",
		"z", "⚡",
		" ", "     ",
	}
}

func decodePairs() []string {
	// Emoji -> harf eşlemesi ve 5 boşluk -> 1 boşluk geri dönüşü burada tutuluyor.
	return []string{
		"🎄", "a",
		"🔩", "b",
		"🌜", "c",
		"💭", "ç",
		"🌛", "d",
		"📛", "e",
		"🎏", "f",
		"🐉", "g",
		"🐌", "ğ",
		"⛄", "h",
		"📏", "ı",
		"✏", "i",
		"🎷", "j",
		"🎋", "k",
		"🕒", "l",
		"👓", "m",
		"👠", "n",
		"📯", "o",
		"🌞", "ö",
		"🎧", "p",
		"💃", "r",
		"🐍", "s",
		"🐢", "ş",
		"☔", "t",
		"🔧", "u",
		"🍇", "ü",
		"☑", "v",
		"🎌", "y",
		"⚡", "z",
		"     ", " ",
	}
}
