package identity

import (
	"fmt"
	"strconv"
	"ynufes-mypage-backend/pkg/snowflake"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

type (
	ID snowflake.Snowflake
)

func IssueID() util.ID {
	return ID(snowflake.NewSnowflake())
}

func ImportID(id string) (util.ID, error) {
	result, err := strconv.ParseInt(id, 36, 64)
	if err != nil {
		return ID(0), fmt.Errorf("failed to convert string to ID, req:%s", id)
	}
	return ID(result), nil
}

func NewID(id int64) ID {
	return ID(id)
}

func (i ID) ExportID() string {
	return strconv.FormatInt(int64(i), 36)
}

func (i ID) HasValue() bool {
	return i != 0
}

func (i ID) GetValue() int64 {
	return int64(i)
}
