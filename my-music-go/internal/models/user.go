package models

import (
	"time"
)

/*
db:"name" 是提供给sqlx来看的
在执行 db.Get() 或 db.Select() 时，自动、准确地将查询结果映射到结构体中，而无需关心字段顺序。

json:"name", 是给 Go 的 encoding/json 包（以及 Gin 框架）看的, 控制了这个字段在序列化成 JSON 字符串时的行为
omitempty, 意思是 "omit if empty" (如果为空则忽略)
json:"-", 意思是“永远不要将此字段序列化到 JSON 中”
*/
// User 对应数据库中的 `user` 表
type User struct {
	ID          int64      `db:"id" json:"id"`
	Username    string     `db:"username" json:"username"`
	Password    string     `db:"password" json:"password"` // 这里可能要返回比对, 而且反正是密文
	Name        *string    `db:"name" json:"name,omitempty"`
	Phone       *string    `db:"phone" json:"phone,omitempty"`
	Email       *string    `db:"email" json:"email,omitempty"`
	Avatar      string     `db:"avatar" json:"avatar"`
	Bio         *string    `db:"bio" json:"bio,omitempty"`
	Gender      int8       `db:"gender" json:"gender"`
	Birthday    *time.Time `db:"birthday" json:"birthday,omitempty"`
	Location    *string    `db:"location" json:"location,omitempty"`
	Status      int8       `db:"status" json:"status"`
	CreatedTime time.Time  `db:"created_time" json:"created_time"`
	UpdatedTime time.Time  `db:"updated_time" json:"updated_time"`
}

// Manager 对应数据库中的 `manager` 表
type Manager struct {
	ID       int64  `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"-"`
}

// Follow 对应数据库中的 `follow` 关注关系表
type Follow struct {
	ID          int64     `db:"id" json:"id"`
	FollowerID  int64     `db:"follower_id" json:"follower_id"`   // 关注者
	FollowingID int64     `db:"following_id" json:"following_id"` // 被关注者
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

// ---- DTOs for User APIs ----

// RegisterRequest 用于绑定用户注册的JSON请求 (DTO)
type RegisterRequest struct {
	Username string     `json:"username" binding:"required"`
	Password string     `json:"password" binding:"required"`
	Name     *string    `json:"name" binding:"required"`
	Email    *string    `json:"email,omitempty"`
	Phone    *string    `json:"phone,omitempty"`
	Avatar   *string    `json:"avatar,omitempty"`
	Bio      *string    `json:"bio,omitempty"`
	Gender   *int8      `json:"gender,omitempty"`
	Birthday *time.Time `db:"birthday" json:"birthday,omitempty"`
	Location *string    `db:"location" json:"location,omitempty"`
}

// LoginRequest 用于绑定用户登录的JSON请求 (DTO)
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserProfile struct {
	User
	Followings int `db:"followings" json:"followings"`
	Followers  int `db:"followers" json:"followers"`
}

// UpdateUserRequest 用于绑定更新用户信息的JSON请求 (DTO)
type UpdateUserRequest struct {
	Name     *string    `json:"name,omitempty"`
	Avatar   *string    `json:"avatar,omitempty"`
	Email    *string    `json:"email,omitempty"`
	Phone    *string    `json:"phone,omitempty"`
	Gender   *int8      `json:"gender,omitempty"`
	Birthday *time.Time `json:"birthday,omitempty"`
	Location *string    `json:"location,omitempty"`
	Bio      *string    `json:"bio,omitempty"`
}

// UsernameResponse 用于专门返回用户名的API (DTO)
type UsernameResponse struct {
	Username string  `json:"username"`
	Name     *string `json:"name,omitempty"` // 可以顺便带上昵称
}
