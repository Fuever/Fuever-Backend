package service

import "Fuever/model"

func GetBlocks(limit int, offset int) ([]*model.BlockInfo, error) {
	blockInfo, err := model.GetBlockAndAuthorNameWithLimitOffset(limit, offset)
	if err != nil {
		return nil, err
	}
	return blockInfo, nil
}
