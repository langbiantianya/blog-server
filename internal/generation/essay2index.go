package generation

import (
	"blog-server/internal/entity"
	"blog-server/internal/entity/vo"
	"blog-server/public/utils"
	"encoding/json"
)

func Index(essays []entity.Essay) (string, error) {
	res := utils.Map(essays, func(index int, essay entity.Essay) (vo.IndexVo, error) {
		return vo.IndexVo{
			// 提取出 title
			// 提取出 post
			// 提取出 tag
			// 提取出 id
			Id:    essay.ID,
			Title: essay.Title,
			Post:  essay.Post,
			Tags:  utils.Map(essay.Tags, func(index int, tag entity.Tag) (string, error) { return tag.Name, nil }),
		}, nil
	})

	// 整合序列化为json
	jsonBytes, err := json.Marshal(res)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
