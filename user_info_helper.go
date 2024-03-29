package workwx

import (
	"fmt"
	"strconv"
)

func reshapeDeptInfo(
	ids []int64,
	orders []uint32,
	leaderStatuses []int,
) []UserDeptInfo {
	result := make([]UserDeptInfo, len(ids))
	for i := range ids {
		result[i].DeptID = ids[i]

		if len(orders) == len(ids) {
			result[i].Order = orders[i]
		}
		if len(leaderStatuses) == len(ids) {
			result[i].IsLeader = leaderStatuses[i] != 0
		}
	}

	return result
}

func mustFromGenderStr(x string) UserGender {
	if x == "" {
		return UserGenderUnspecified
	}
	n, err := strconv.Atoi(x)
	if err != nil {
		panic(fmt.Sprintf("gender string parse failed: %+v", err))
	}

	return UserGender(n)
}

func (x userDetailResp) intoUserInfo() UserInfo {
	deptInfo := reshapeDeptInfo(x.DeptIDs, x.DeptOrder, x.IsLeaderInDept)

	return UserInfo{
		UserID:         x.UserID,
		Name:           x.Name,
		Position:       x.Position,
		Departments:    deptInfo,
		Mobile:         x.Mobile,
		Gender:         mustFromGenderStr(x.Gender),
		Email:          x.Email,
		AvatarURL:      x.AvatarURL,
		Telephone:      x.Telephone,
		IsEnabled:      x.IsEnabled != 0,
		Alias:          x.Alias,
		Status:         x.Status,
		QRCodeURL:      x.QRCodeURL,
		DeptIDs:        x.DeptIDs,
		DeptOrder:      x.DeptOrder,
		IsLeaderInDept: x.IsLeaderInDept,
	}
}
