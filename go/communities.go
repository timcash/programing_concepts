// A struct to hold the data for a community
package main

type Community struct {
	ID          int64
	Name        string
	Description string
	UserID      int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Language: go
// Path: go/communities.go
// A struct to hold the data for a community user
type CommunityUser struct {
	ID          int64
	CommunityID int64
	UserID      int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Language: go
// Path: go/communities.go
// A struct to hold the data for a community post
type CommunityPost struct {
	ID          int64
	CommunityID int64
	UserID      int64
	Title       string
	Body        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Language: go
// Path: go/communities.go
// A struct to hold the data for a community post comment
type CommunityPostComment struct {
	ID          int64
	CommunityPostID int64
	UserID      int64
	Body        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Language: go
// Path: go/communities.go
// A struct to hold the data for a community post like
type CommunityPostLike struct {
	ID          int64
	CommunityPostID int64
	UserID      int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Language: go
// Path: go/communities.go
// A struct to hold the data for a community post comment like
type CommunityPostCommentLike struct {
	ID          int64
	CommunityPostCommentID int64
	UserID      int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Language: go
// Path: go/communities.go
// A struct to hold the data for a community post comment reply
type CommunityPostCommentReply struct {
	ID          int64
	CommunityPostCommentID int64
	UserID      int64
	Body        string
	CreatedAt   time.Time