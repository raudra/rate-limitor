package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/raudra/rate-limitor/rule-manager/config"
	"github.com/raudra/rate-limitor/rule-manager/src/models"
)

func GetRules() ([]models.Rule, error) {
	ctx := context.TODO()
	redisClient := config.RedisClient()

	data, err := redisClient.SMembers(ctx, "Rules").
		Result()

	fmt.Println(data)
	rules := []models.Rule{}

	for _, r := range data {
		var rule models.Rule
		_ = json.Unmarshal([]byte(r), &rule)
		rules = append(rules, rule)
	}

	return rules, err
}
