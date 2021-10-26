package message

// Define a user structure

type User struct {

	// Ensure that the key of the json string of the user information is consistent with the tag name corresponding to the field of the structure
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"`
	Sex        string `json:"sex"`
}
