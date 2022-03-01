package apiServer

import "strconv"

type Product struct {
	Id          int    `json:"id"`
	OwnerId     int    `json:"owner_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       struct {
		Amount   string `json:"amount"`
		Currency struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		}
		OldAmount string `json:"old_amount"`
		Text      string `json:"text"`
	}
	Dimensions struct {
		Width  int
		Height int
		Length int
	} `json:"dimensions"`
	Weight   int `json:"weight"`
	Category struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Section struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		}
	}
	ThumbPhoto   string `json:"thumb_photo"`
	Date         int    `json:"date"`
	Availability int    `json:"availability"`
	IsFavorite   bool   `json:"is_favorite"`
	Sku          string `json:"sku"`
}

type MMarketGetOkResponse struct {
	Count int       `json:"count"`
	Items []Product `json:"items"`
}

type MMarketGetResponse struct {
	Response MMarketGetOkResponse `json:"response"`
}

func (vkApiServer *VkApiServer) MMarketGet(accessToken string, ownerId int, count int, offset int) ([]Product, error) {
	var resp MMarketGetResponse
	err := vkApiServer.SendMethodRequest(
		"market",
		"get",
		map[string]string{
			"access_token": accessToken,
			"owner_id":     strconv.Itoa(ownerId),
			"count":        strconv.Itoa(count),
			"offset":       strconv.Itoa(offset),
		},
		&resp,
	)
	return resp.Response.Items, err
}

func (vkApiServer *VkApiServer) MMarketGetAll(accessToken string, ownerId int) ([]Product, error) {
	countByRequest := 100
	currentOffset := 0
	var resp []Product

	for {
		currentResp, err := vkApiServer.MMarketGet(accessToken, ownerId, countByRequest, currentOffset)
		if err != nil {
			return nil, err
		}
		resp = append(resp, currentResp...)
		if len(currentResp) < countByRequest {
			break
		}
	}

	return resp, nil
}
