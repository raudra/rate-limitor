package config

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/raudra/rate-limitor/rule-manager/src/models"
	"github.com/rs/zerolog/log"
)

type conf struct {
	Rules []models.Rule
}

func LoadRules() {
	path, _ := filepath.Abs("config/rules/test.yml")

	ymlParser := &YAMLParser{
		filePath: path,
	}

	err := ymlParser.Parse()

	if err != nil {
		log.Error().
			Msg(fmt.Sprintf("%s", err))
	}

	conf := &conf{}

	jsonData, _ := json.Marshal(ymlParser.dataStruct)

	_ = json.Unmarshal(jsonData, conf)

	log.Info().
		Interface("ParseData", conf).
		Msg("Parsing done successfully")

	updateCache(conf)

}

func updateCache(cnf *conf) {
	ctx := context.TODO()

	redisClient := RedisClient()

	for _, rule := range cnf.Rules {
		data, _ := json.Marshal(rule)
		_, err := redisClient.SAdd(ctx, "Rules", data).
			Result()
		if err != nil {
			log.Error().
				Msg(fmt.Sprintf("%s", err))
		}
	}

}
