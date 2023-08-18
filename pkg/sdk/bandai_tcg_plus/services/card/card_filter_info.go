package card

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/DesistDaydream/dtcg/pkg/sdk/bandai_tcg_plus/models"
)

func GetCardFilterInfo(r *models.CardFilterInfoReq) (*models.CardMetadataGetResp, error) {
	url := "https://api.bandai-tcg-plus.com/api/user/card"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("game_title_id", r.GameTitleID)
	q.Add("language_code", r.LanguageCode)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var CardFilterInfo *models.CardMetadataGetResp
	err = json.Unmarshal(body, &CardFilterInfo)
	if err != nil {
		return nil, err
	}

	return CardFilterInfo, nil

}
