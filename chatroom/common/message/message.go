package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

// Constants for user status
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code    int    `json:"code"` // 500: the user is not registered. 200: the login is successful
	UsersId []int  // Save a slice of user id
	Error   string `json:"error"` //Return error message
}

type RegisterMes struct {
	User User `json:"user"` // The type is the User structure.
}
type RegisterResMes struct {
	Code  int    `json:"code"`  // 400: The user has already exi. 200: registered successfully
	Error string `json:"error"` // Return error message
}

// In order to cooperate with the server to push user status changes
type NotifyUserStatusMes struct {
	UserId int `json:"userId"` // User id
	Status int `json:"status"` // User status
}

// Message sent
type SmsMes struct {
	Content string `json:"content"`
	User
}
