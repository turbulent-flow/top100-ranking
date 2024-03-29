package api_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/LiamYabou/top100-ranking/preference"
	"github.com/LiamYabou/top100-ranking/test"
	"github.com/LiamYabou/top100-ranking/api"
)

func (a *actionSuite) TestFindProducts() {
	assert := assert.New(a.T())
	page := 1
	categoryId := 2
	opts := preference.LoadOptions(preference.WithDB(test.DBpool))
	// # Find products
	// ## Standard procedure
	expected := test.CannedRankings
	actual := api.FindRankings(categoryId, page, opts)
	failedMsg := fmt.Sprintf("Failed, expected the result: %v, got the result: %v", expected, actual)
	assert.Equal(expected, actual, failedMsg)
	// ## Empty result
	categoryId = 8
	expected = `{"status":"success","data":null}`
	actual = api.FindRankings(categoryId, page, opts)
	failedMsg = fmt.Sprintf("Failed, expected the result: %v, got the result: %v", expected, actual)
	assert.Equal(expected, actual, failedMsg)
	// ## Empty product set
	categoryId = 1
	expected = `{"status":"success","data":{"categories":[{"id":2,"name":"Amazon Devices \u0026 Accessories"}],"products":[],"root_category":null,"selected_category_name":"Any Department"}}`
	actual = api.FindRankings(categoryId, page, opts)
	failedMsg = fmt.Sprintf("Failed, expected the result: %v, got the result: %v", expected, actual)
	assert.Equal(expected, actual, failedMsg)
	// ## category = 0
	categoryId = 0
	expected = `{"status":"fail","data":{"category_id":"The category id is invaild, it should be greater than zero."}}`
	actual = api.FindRankings(categoryId, page, opts)
	failedMsg = fmt.Sprintf("Failed, expected the result: %v, got the result: %v", expected, actual)
	assert.Equal(expected, actual, failedMsg)
	// ## page = 0
	categoryId = 2
	page = 0
	expected = `{"status":"fail","data":{"page":"The page number is invaild, it should either be 1 or 2."}}`
	actual = api.FindRankings(categoryId, page, opts)
	failedMsg = fmt.Sprintf("Failed, expected the result: %v, got the result: %v", expected, actual)
	assert.Equal(expected, actual, failedMsg)
	// ## page > 2
	page = 3
	expected = `{"status":"fail","data":{"page":"The page number is invaild, it should either be 1 or 2."}}`
	actual = api.FindRankings(categoryId, page, opts)
	failedMsg = fmt.Sprintf("Failed, expected the result: %v, got the result: %v", expected, actual)
	assert.Equal(expected, actual, failedMsg)
}
