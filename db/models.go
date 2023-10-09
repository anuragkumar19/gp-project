// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package database

import (
	"database/sql"
	"time"
)

type Post struct {
	ID          int32     `json:"id"`
	Title       string    `json:"title"`
	Text        string    `json:"text"`
	Image       string    `json:"image"`
	Video       string    `json:"video"`
	Link        string    `json:"link"`
	SubredditID int32     `json:"subreddit_id"`
	CreatorID   int32     `json:"creator_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Reply struct {
	ID            int32         `json:"id"`
	UserID        int32         `json:"user_id"`
	PostID        int32         `json:"post_id"`
	ParentReplyID sql.NullInt32 `json:"parent_reply_id"`
	Content       string        `json:"content"`
	CreatedAt     sql.NullTime  `json:"created_at"`
	UpdatedAt     sql.NullTime  `json:"updated_at"`
}

type Subreddit struct {
	ID         int32     `json:"id"`
	Title      string    `json:"title"`
	IsVerified bool      `json:"is_verified"`
	Name       string    `json:"name"`
	About      string    `json:"about"`
	Avatar     string    `json:"avatar"`
	Cover      string    `json:"cover"`
	CreatorID  int32     `json:"creator_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type User struct {
	ID              int32         `json:"id"`
	Name            string        `json:"name"`
	Email           string        `json:"email"`
	Avatar          string        `json:"avatar"`
	IsEmailVerified bool          `json:"is_email_verified"`
	Otp             sql.NullInt32 `json:"otp"`
	OtpExpiry       sql.NullTime  `json:"otp_expiry"`
	Password        string        `json:"password"`
	Username        string        `json:"username"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}

type UserSubredditJoin struct {
	UserID      int32 `json:"user_id"`
	SubredditID int32 `json:"subreddit_id"`
}

type VotePost struct {
	UserID int32 `json:"user_id"`
	PostID int32 `json:"post_id"`
	Down   bool  `json:"down"`
}
