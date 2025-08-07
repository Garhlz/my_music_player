package services

import "errors"

// 定义已知的业务逻辑错误
// 注意, 如果是"来自第三方库或标准库的非预期错误", 不用设置哨兵错误, 而是使用 fmt.Errorf("简短描述: %w", err) 来包装它，然后返回
var (
	// 注册
	ErrUserNotFound = errors.New("user not found")

	// 登录
	ErrInvalidCredentials = errors.New("invalid credentials") // 凭证无效（通常指密码错误）
	ErrUserAlreadyExists  = errors.New("username already exists")

	// 歌曲
	ErrSongNotFound = errors.New("song not found")

	// 歌手
	ErrArtistNotFound = errors.New("artist not found")

	ErrAlbumNotFound = errors.New("album not found")

	ErrSongAlreadyLiked = errors.New("song already liked")
	ErrLikeNotFound     = errors.New("like record not found")
)
