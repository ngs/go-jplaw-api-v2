package lawapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Client provides access to the Japan Law API
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new API client
func NewClient() *Client {
	return &Client{
		baseURL:    DefaultBaseURL,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

// SetHTTPClient sets a custom HTTP client
func (c *Client) SetHTTPClient(client *http.Client) {
	c.httpClient = client
}

// GetAttachmentParams contains query parameters for GetAttachment
type GetAttachmentParams struct {
	// Src represents 法令XML中のFig要素のsrc属性 > jpgの例：`./pict/H11HO127-001.jpg` > pdfの例：`./pict/2FH00000007000.pdf`
	Src *string
}

// GetAttachment field from the API response
func (c *Client) GetAttachment(lawRevisionId string, params *GetAttachmentParams) (*string, error) {
	urlPath := c.baseURL + "/attachment" + "/" + lawRevisionId
	if params != nil {
		queryParams := url.Values{}
		if params.Src != nil {
			queryParams.Set("src", fmt.Sprintf("%v", *params.Src))
		}
		if len(queryParams) > 0 {
			urlPath += "?" + queryParams.Encode()
		}
	}
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	result := string(body)
	return &result, nil
}

// GetKeywordParams contains query parameters for GetKeyword
type GetKeywordParams struct {
	// Keyword represents field from the API response
	Keyword string
	// LawNum represents field from the API response
	LawNum *string
	// LawNumEra represents law numberの元号 > 例： `Heisei`
	LawNumEra *LawNumEra
	// LawNumNum represents law numberの号数 > 例： `006`
	LawNumNum *string
	// LawNumType represents law numberの法令type 種類の定義はSchemasの"#model-law_num_type">`law_num_type`を参照してください。 > 例： `Rule`
	LawNumType *LawNumType
	// LawNumYear represents law numberの年 > 例： `28`
	LawNumYear *int
	// LawType represents 法令type（複数指定可） > 例： `Act,Rule`
	LawType *[]LawType
	// Asof represents 法令の時点。指定時点以前で最新のamendmenthistoryを、各法令の `revision_info` に格納します。省略した場合、現時点でsearchします。 > 例： `2024-05-27`
	Asof *Date
	// CategoryCd represents 事項別分類コード（複数指定可） コードの定義はSchemasの"#model-category_cd">`category_cd`を参照してください。 > 例： `011,021`
	CategoryCd *[]CategoryCd
	// PromulgationDateFrom represents promulgation date（開始） > 例： `2016-12-15`
	PromulgationDateFrom *Date
	// PromulgationDateTo represents promulgation date（終了） > 例： `2016-12-15`
	PromulgationDateTo *Date
	// Limit represents レスポンスの`sentences`の`position`数の総和の上限。 > 例：`50` > 既定値： `100` > 上限値： `1000`
	Limit *int32
	// Offset represents field from the API response
	Offset *int32
	// Order represents field from the API response
	Order *string
	// ResponseFormat represents レスポンスformat（`json` 又は `xml`）。指定なしの場合はAcceptヘッダから判断、判断できない場合は `json` とする。 > 例： `json` > 既定値： 指定なし
	ResponseFormat *ResponseFormat
	// SentencesLimit represents field from the API response
	SentencesLimit *int32
	// SentenceTextSize represents レスポンス：`items`->`sentences`->`text` の表示文字数（`highlight_tag`で指定したHTMLタグを含む） > 例：`20` > 既定値： `100`
	SentenceTextSize *int32
	// HighlightTag represents `keyword`で指定された文言のヒット箇所を囲むHTMLタグ名。 > 例： `em` > 規定値： `span`
	HighlightTag *string
}

// GetKeyword field from the API response
func (c *Client) GetKeyword(params *GetKeywordParams) (*KeywordResponse, error) {
	urlPath := c.baseURL + "/keyword"
	if params != nil {
		queryParams := url.Values{}
		queryParams.Set("keyword", fmt.Sprintf("%v", params.Keyword))
		if params.LawNum != nil {
			queryParams.Set("law_num", fmt.Sprintf("%v", *params.LawNum))
		}
		if params.LawNumEra != nil {
			queryParams.Set("law_num_era", fmt.Sprintf("%v", *params.LawNumEra))
		}
		if params.LawNumNum != nil {
			queryParams.Set("law_num_num", fmt.Sprintf("%v", *params.LawNumNum))
		}
		if params.LawNumType != nil {
			queryParams.Set("law_num_type", fmt.Sprintf("%v", *params.LawNumType))
		}
		if params.LawNumYear != nil {
			queryParams.Set("law_num_year", fmt.Sprintf("%v", *params.LawNumYear))
		}
		if params.LawType != nil {
			for _, v := range *params.LawType {
				queryParams.Add("law_type", fmt.Sprintf("%v", v))
			}
		}
		if params.Asof != nil {
			queryParams.Set("asof", fmt.Sprintf("%v", *params.Asof))
		}
		if params.CategoryCd != nil {
			for _, v := range *params.CategoryCd {
				queryParams.Add("category_cd", fmt.Sprintf("%v", v))
			}
		}
		if params.PromulgationDateFrom != nil {
			queryParams.Set("promulgation_date_from", fmt.Sprintf("%v", *params.PromulgationDateFrom))
		}
		if params.PromulgationDateTo != nil {
			queryParams.Set("promulgation_date_to", fmt.Sprintf("%v", *params.PromulgationDateTo))
		}
		if params.Limit != nil {
			queryParams.Set("limit", fmt.Sprintf("%v", *params.Limit))
		}
		if params.Offset != nil {
			queryParams.Set("offset", fmt.Sprintf("%v", *params.Offset))
		}
		if params.Order != nil {
			queryParams.Set("order", fmt.Sprintf("%v", *params.Order))
		}
		if params.ResponseFormat != nil {
			queryParams.Set("response_format", fmt.Sprintf("%v", *params.ResponseFormat))
		}
		if params.SentencesLimit != nil {
			queryParams.Set("sentences_limit", fmt.Sprintf("%v", *params.SentencesLimit))
		}
		if params.SentenceTextSize != nil {
			queryParams.Set("sentence_text_size", fmt.Sprintf("%v", *params.SentenceTextSize))
		}
		if params.HighlightTag != nil {
			queryParams.Set("highlight_tag", fmt.Sprintf("%v", *params.HighlightTag))
		}
		if len(queryParams) > 0 {
			urlPath += "?" + queryParams.Encode()
		}
	}
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	var result KeywordResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// GetLawDataParams contains query parameters for GetLawData
type GetLawDataParams struct {
	// LawFullTextFormat represents 法令text contentのformat（`json` 又は `xml`）。指定なしの場合は`response_format`により判断されるformatに合わせる。 > 例： `json` > 既定値： 指定なし
	LawFullTextFormat *ResponseFormat
	// Asof represents field from the API response
	Asof *Date
	// Elm represents field from the API response
	Elm *Elm
	// OmitAmendmentSupplProvision represents `true`の場合、レスポンスの`law_full_text`の中にamendment法令の附則を含めない > 例： `true` > 既定値： `false`
	OmitAmendmentSupplProvision *bool
	// IncludeAttachedFileContent represents `true`の場合、レスポンスの`attached_files_info`の`image_data`を返却します。 > 例： `true` > 既定値： `false`
	IncludeAttachedFileContent *bool
	// ResponseFormat represents レスポンスformat（`json` 又は `xml`）。指定なしの場合はAcceptヘッダから判断、判断できない場合は `json` とする。 > 例： `json` > 既定値： 指定なし
	ResponseFormat *ResponseFormat
}

// GetLawData field from the API response
func (c *Client) GetLawData(lawIdOrNumOrRevisionId string, params *GetLawDataParams) (*LawDataResponse, error) {
	urlPath := c.baseURL + "/law_data" + "/" + lawIdOrNumOrRevisionId
	if params != nil {
		queryParams := url.Values{}
		if params.LawFullTextFormat != nil {
			queryParams.Set("law_full_text_format", fmt.Sprintf("%v", *params.LawFullTextFormat))
		}
		if params.Asof != nil {
			queryParams.Set("asof", fmt.Sprintf("%v", *params.Asof))
		}
		if params.Elm != nil {
			queryParams.Set("elm", fmt.Sprintf("%v", *params.Elm))
		}
		if params.OmitAmendmentSupplProvision != nil {
			queryParams.Set("omit_amendment_suppl_provision", fmt.Sprintf("%v", *params.OmitAmendmentSupplProvision))
		}
		if params.IncludeAttachedFileContent != nil {
			queryParams.Set("include_attached_file_content", fmt.Sprintf("%v", *params.IncludeAttachedFileContent))
		}
		if params.ResponseFormat != nil {
			queryParams.Set("response_format", fmt.Sprintf("%v", *params.ResponseFormat))
		}
		if len(queryParams) > 0 {
			urlPath += "?" + queryParams.Encode()
		}
	}
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	var result LawDataResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// GetLawFileParams contains query parameters for GetLawFile
type GetLawFileParams struct {
	// Asof represents field from the API response
	Asof *Date
}

// GetLawFile field from the API response
func (c *Client) GetLawFile(lawIdOrNumOrRevisionId string, fileType string, params *GetLawFileParams) (*string, error) {
	urlPath := c.baseURL + "/law_file" + "/" + fileType + "/" + lawIdOrNumOrRevisionId
	if params != nil {
		queryParams := url.Values{}
		if params.Asof != nil {
			queryParams.Set("asof", fmt.Sprintf("%v", *params.Asof))
		}
		if len(queryParams) > 0 {
			urlPath += "?" + queryParams.Encode()
		}
	}
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	result := string(body)
	return &result, nil
}

// GetRevisionsParams contains query parameters for GetRevisions
type GetRevisionsParams struct {
	// LawTitle represents field from the API response
	LawTitle *string
	// LawTitleKana represents field from the API response
	LawTitleKana *string
	// AmendmentDateFrom represents amendment法令施行期日（指定値を含む、それ以後） > 例： `2024-06-07`
	AmendmentDateFrom *Date
	// AmendmentDateTo represents amendment法令施行期日（指定値を含む、それ以前） > 例： `2024-06-07`
	AmendmentDateTo *Date
	// AmendmentLawId represents amendment法令のlaw ID（部分一致） > 例： `506AC0000000046`
	AmendmentLawId *string
	// AmendmentLawNum represents field from the API response
	AmendmentLawNum *string
	// AmendmentLawTitle represents field from the API response
	AmendmentLawTitle *string
	// AmendmentLawTitleKana represents field from the API response
	AmendmentLawTitleKana *string
	// AmendmentPromulgateDateFrom represents amendment法令promulgation date（指定値を含む、それ以後） > 例： `2024-06-07`
	AmendmentPromulgateDateFrom *Date
	// AmendmentPromulgateDateTo represents amendment法令promulgation date（指定値を含む、それ以前） > 例： `2024-06-07`
	AmendmentPromulgateDateTo *Date
	// AmendmentType represents amendmenttype（複数指定可） amendmenttypeの定義はSchemasの"#model-amendment_type">`amendment_type`を参照してください。 > 例： `1,3`
	AmendmentType *[]AmendmentType
	// CategoryCd represents 事項別分類コード（複数指定可） コードの定義はSchemasの"#model-category_cd">`category_cd`を参照してください。 > 例： `011,021`
	CategoryCd *[]CategoryCd
	// CurrentRevisionStatus represents field from the API response
	CurrentRevisionStatus *[]CurrentRevisionStatus
	// Mission represents 新規制定又は被amendment法令（`New`）・一部amendment法令（`Partial`）を指定（複数指定可） > 例： `New,Partial`
	Mission *[]Mission
	// RemainInForce represents repeal後の効力（`true`:repeal後でも効力を有するもの / `false`:repeal後に効力を有しないもの） > 例： `false`
	RemainInForce *bool
	// RepealDateFrom represents repeal日（指定値を含む、それ以後） > 例： `2024-04-01`
	RepealDateFrom *Date
	// RepealDateTo represents repeal日（指定値を含む、それ以前） > 例： `2024-04-01`
	RepealDateTo *Date
	// RepealStatus represents field from the API response
	RepealStatus *[]RepealStatus
	// UpdatedFrom represents dataの更新日（指定値を含む、それ以後） > 例： `2024-06-07`
	UpdatedFrom *Date
	// UpdatedTo represents dataの更新日（指定値を含む、それ以前） > 例： `2024-06-07`
	UpdatedTo *Date
	// ResponseFormat represents レスポンスformat（`json` 又は `xml`）。指定なしの場合はAcceptヘッダから判断、判断できない場合は `json` とする。 > 例： `json` > 既定値： 指定なし
	ResponseFormat *ResponseFormat
}

// GetRevisions field from the API response
func (c *Client) GetRevisions(lawIdOrNum string, params *GetRevisionsParams) (*LawRevisionsResponse, error) {
	urlPath := c.baseURL + "/law_revisions" + "/" + lawIdOrNum
	if params != nil {
		queryParams := url.Values{}
		if params.LawTitle != nil {
			queryParams.Set("law_title", fmt.Sprintf("%v", *params.LawTitle))
		}
		if params.LawTitleKana != nil {
			queryParams.Set("law_title_kana", fmt.Sprintf("%v", *params.LawTitleKana))
		}
		if params.AmendmentDateFrom != nil {
			queryParams.Set("amendment_date_from", fmt.Sprintf("%v", *params.AmendmentDateFrom))
		}
		if params.AmendmentDateTo != nil {
			queryParams.Set("amendment_date_to", fmt.Sprintf("%v", *params.AmendmentDateTo))
		}
		if params.AmendmentLawId != nil {
			queryParams.Set("amendment_law_id", fmt.Sprintf("%v", *params.AmendmentLawId))
		}
		if params.AmendmentLawNum != nil {
			queryParams.Set("amendment_law_num", fmt.Sprintf("%v", *params.AmendmentLawNum))
		}
		if params.AmendmentLawTitle != nil {
			queryParams.Set("amendment_law_title", fmt.Sprintf("%v", *params.AmendmentLawTitle))
		}
		if params.AmendmentLawTitleKana != nil {
			queryParams.Set("amendment_law_title_kana", fmt.Sprintf("%v", *params.AmendmentLawTitleKana))
		}
		if params.AmendmentPromulgateDateFrom != nil {
			queryParams.Set("amendment_promulgate_date_from", fmt.Sprintf("%v", *params.AmendmentPromulgateDateFrom))
		}
		if params.AmendmentPromulgateDateTo != nil {
			queryParams.Set("amendment_promulgate_date_to", fmt.Sprintf("%v", *params.AmendmentPromulgateDateTo))
		}
		if params.AmendmentType != nil {
			for _, v := range *params.AmendmentType {
				queryParams.Add("amendment_type", fmt.Sprintf("%v", v))
			}
		}
		if params.CategoryCd != nil {
			for _, v := range *params.CategoryCd {
				queryParams.Add("category_cd", fmt.Sprintf("%v", v))
			}
		}
		if params.CurrentRevisionStatus != nil {
			for _, v := range *params.CurrentRevisionStatus {
				queryParams.Add("current_revision_status", fmt.Sprintf("%v", v))
			}
		}
		if params.Mission != nil {
			for _, v := range *params.Mission {
				queryParams.Add("mission", fmt.Sprintf("%v", v))
			}
		}
		if params.RemainInForce != nil {
			queryParams.Set("remain_in_force", fmt.Sprintf("%v", *params.RemainInForce))
		}
		if params.RepealDateFrom != nil {
			queryParams.Set("repeal_date_from", fmt.Sprintf("%v", *params.RepealDateFrom))
		}
		if params.RepealDateTo != nil {
			queryParams.Set("repeal_date_to", fmt.Sprintf("%v", *params.RepealDateTo))
		}
		if params.RepealStatus != nil {
			for _, v := range *params.RepealStatus {
				queryParams.Add("repeal_status", fmt.Sprintf("%v", v))
			}
		}
		if params.UpdatedFrom != nil {
			queryParams.Set("updated_from", fmt.Sprintf("%v", *params.UpdatedFrom))
		}
		if params.UpdatedTo != nil {
			queryParams.Set("updated_to", fmt.Sprintf("%v", *params.UpdatedTo))
		}
		if params.ResponseFormat != nil {
			queryParams.Set("response_format", fmt.Sprintf("%v", *params.ResponseFormat))
		}
		if len(queryParams) > 0 {
			urlPath += "?" + queryParams.Encode()
		}
	}
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	var result LawRevisionsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// GetLawsParams contains query parameters for GetLaws
type GetLawsParams struct {
	// LawId represents law ID（部分一致） > 例： `322CO0000000016`
	LawId *string
	// LawNum represents field from the API response
	LawNum *string
	// LawNumEra represents law numberの元号 > 例： `Showa`
	LawNumEra *LawNumEra
	// LawNumNum represents law numberの号数 > 例： `88`
	LawNumNum *string
	// LawNumType represents law numberの法令type 種類の定義はSchemasの"#model-law_num_type">`law_num_type`を参照してください。 > 例： `Act`
	LawNumType *LawNumType
	// LawNumYear represents law numberの年 > 例： `60`
	LawNumYear *int
	// LawTitle represents field from the API response
	LawTitle *string
	// LawTitleKana represents field from the API response
	LawTitleKana *string
	// LawType represents 法令type（複数指定可） > 例： `Act,Rule`
	LawType *[]LawType
	// AmendmentLawId represents amendment法令のlaw ID（部分一致） > 注意：本パラメータを指定した場合、パラメータ：法令の時点（`asof`）を無視します。 > 例： `429AC0000000054`
	AmendmentLawId *string
	// Asof represents 法令の時点。指定時点以前で最新のamendmenthistoryを、各法令の `revision_info` に格納します。省略した場合、現時点でsearchします。 > 例： `2023-07-01`
	Asof *Date
	// CategoryCd represents 事項別分類コード（複数指定可） コードの定義はSchemasの"#model-category_cd">`category_cd`を参照してください。 > 例： `001,002`
	CategoryCd *[]CategoryCd
	// Mission represents 新規制定又は被amendment法令（`New`）・一部amendment法令（`Partial`）を指定（複数指定可） > 例： `New,Partial`
	Mission *[]Mission
	// OmitCurrentRevisionInfo represents `true`の場合、法令の時点（`asof`）に依存しない現在以前の最新の版のinformation（`current_revision_info`）をレスポンスに含めない > 例： `true` > 既定値： `false`
	OmitCurrentRevisionInfo *bool
	// PromulgationDateFrom represents promulgation date（指定値を含む、それ以後） > 例： `2023-07-01`
	PromulgationDateFrom *Date
	// PromulgationDateTo represents promulgation date（指定値を含む、それ以前） > 例： `2023-07-01`
	PromulgationDateTo *Date
	// RepealStatus represents field from the API response
	RepealStatus *[]RepealStatus
	// Limit represents レスポンスの `laws` のretrieve件数の上限。 > 例：`50` > 既定値：`100`
	Limit *int32
	// Offset represents field from the API response
	Offset *int32
	// Order represents field from the API response
	Order *string
	// ResponseFormat represents レスポンスformat（`json` 又は `xml`）。指定なしの場合はAcceptヘッダから判断、判断できない場合は `json` とする。 > 例： `json` > 既定値： 指定なし
	ResponseFormat *ResponseFormat
}

// GetLaws field from the API response
func (c *Client) GetLaws(params *GetLawsParams) (*LawsResponse, error) {
	urlPath := c.baseURL + "/laws"
	if params != nil {
		queryParams := url.Values{}
		if params.LawId != nil {
			queryParams.Set("law_id", fmt.Sprintf("%v", *params.LawId))
		}
		if params.LawNum != nil {
			queryParams.Set("law_num", fmt.Sprintf("%v", *params.LawNum))
		}
		if params.LawNumEra != nil {
			queryParams.Set("law_num_era", fmt.Sprintf("%v", *params.LawNumEra))
		}
		if params.LawNumNum != nil {
			queryParams.Set("law_num_num", fmt.Sprintf("%v", *params.LawNumNum))
		}
		if params.LawNumType != nil {
			queryParams.Set("law_num_type", fmt.Sprintf("%v", *params.LawNumType))
		}
		if params.LawNumYear != nil {
			queryParams.Set("law_num_year", fmt.Sprintf("%v", *params.LawNumYear))
		}
		if params.LawTitle != nil {
			queryParams.Set("law_title", fmt.Sprintf("%v", *params.LawTitle))
		}
		if params.LawTitleKana != nil {
			queryParams.Set("law_title_kana", fmt.Sprintf("%v", *params.LawTitleKana))
		}
		if params.LawType != nil {
			for _, v := range *params.LawType {
				queryParams.Add("law_type", fmt.Sprintf("%v", v))
			}
		}
		if params.AmendmentLawId != nil {
			queryParams.Set("amendment_law_id", fmt.Sprintf("%v", *params.AmendmentLawId))
		}
		if params.Asof != nil {
			queryParams.Set("asof", fmt.Sprintf("%v", *params.Asof))
		}
		if params.CategoryCd != nil {
			for _, v := range *params.CategoryCd {
				queryParams.Add("category_cd", fmt.Sprintf("%v", v))
			}
		}
		if params.Mission != nil {
			for _, v := range *params.Mission {
				queryParams.Add("mission", fmt.Sprintf("%v", v))
			}
		}
		if params.OmitCurrentRevisionInfo != nil {
			queryParams.Set("omit_current_revision_info", fmt.Sprintf("%v", *params.OmitCurrentRevisionInfo))
		}
		if params.PromulgationDateFrom != nil {
			queryParams.Set("promulgation_date_from", fmt.Sprintf("%v", *params.PromulgationDateFrom))
		}
		if params.PromulgationDateTo != nil {
			queryParams.Set("promulgation_date_to", fmt.Sprintf("%v", *params.PromulgationDateTo))
		}
		if params.RepealStatus != nil {
			for _, v := range *params.RepealStatus {
				queryParams.Add("repeal_status", fmt.Sprintf("%v", v))
			}
		}
		if params.Limit != nil {
			queryParams.Set("limit", fmt.Sprintf("%v", *params.Limit))
		}
		if params.Offset != nil {
			queryParams.Set("offset", fmt.Sprintf("%v", *params.Offset))
		}
		if params.Order != nil {
			queryParams.Set("order", fmt.Sprintf("%v", *params.Order))
		}
		if params.ResponseFormat != nil {
			queryParams.Set("response_format", fmt.Sprintf("%v", *params.ResponseFormat))
		}
		if len(queryParams) > 0 {
			urlPath += "?" + queryParams.Encode()
		}
	}
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	var result LawsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// Helper functions for creating pointer values

// StringPtr returns a pointer to the string value
func StringPtr(v string) *string {
	return &v
}

// IntPtr returns a pointer to the int value
func IntPtr(v int) *int {
	return &v
}

// Int32Ptr returns a pointer to the int32 value
func Int32Ptr(v int32) *int32 {
	return &v
}

// Int64Ptr returns a pointer to the int64 value
func Int64Ptr(v int64) *int64 {
	return &v
}

// BoolPtr returns a pointer to the bool value
func BoolPtr(v bool) *bool {
	return &v
}

// Float32Ptr returns a pointer to the float32 value
func Float32Ptr(v float32) *float32 {
	return &v
}

// Float64Ptr returns a pointer to the float64 value
func Float64Ptr(v float64) *float64 {
	return &v
}

