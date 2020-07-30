package dto

type UserDTO struct {
	UserId   *string `json:"userId"`
	UserName *string `json:"userName"`
	Age      *int    `json:"age"`
	Status   *bool   `json:"status"`
}

func (userDTO *UserDTO) Validate() bool {
	return userDTO.UserId != nil && userDTO.UserName != nil && userDTO.Age != nil && userDTO.Status != nil
}

func (userDTO *UserDTO) SetUserId(userId string) {
	userDTO.UserId = &userId
}
