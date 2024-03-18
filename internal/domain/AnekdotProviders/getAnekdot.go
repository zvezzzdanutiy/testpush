package AnekdotProviders

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (d *Domain) GetAnekdot(ctx context.Context, category string) (string, error) {
	response, err := d.client.Get(category)
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("wrong response got, code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), err
}
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
func (d *Domain) GenerateURL() string {
	const apiKey = "b7f30f7550d0120aaea92d9a9c64f9037bd441474c86717ca96f1e81ccafb5ec"
	x := fmt.Sprintf("%d", time.Now().Unix())
	str := "pid=805v15r485hdqkshtvse&method=getRandItem&uts=" + x // Установите значение num равным желаемому количеству анекдотов
	hash := GetMD5Hash(str + apiKey)
	return "http://anecdotica.ru/api?" + str + "&hash=" + hash
}
