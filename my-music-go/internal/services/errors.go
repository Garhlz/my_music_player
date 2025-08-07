package services

import "errors"

// 定义已知的业务逻辑错误
// 注意, 如果是"来自第三方库或标准库的非预期错误",而非业务逻辑错误, 不用设置哨兵错误, 而是使用 fmt.Errorf("简短描述: %w", err) 来包装它，然后返回
var (
	// User related
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserAlreadyExists  = errors.New("username already exists")

	// Song related
	ErrSongNotFound     = errors.New("song not found")
	ErrSongAlreadyLiked = errors.New("song already liked")
	ErrLikeNotFound     = errors.New("like record not found")

	// Artist/Album related
	ErrArtistNotFound = errors.New("artist not found")
	ErrAlbumNotFound  = errors.New("album not found")

	// Playlist related (新增)
	ErrPlaylistNotFound      = errors.New("playlist not found")
	ErrSongAlreadyInPlaylist = errors.New("song already in playlist")
	ErrSongNotInPlaylist     = errors.New("song not in playlist")

	// General
	ErrForbidden = errors.New("operation not permitted")
)
