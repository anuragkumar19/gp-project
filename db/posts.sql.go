// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: posts.sql

package database

import (
	"context"
)

const createPost = `-- name: CreatePost :many
INSERT INTO posts (
    title, text, image, video, link, subreddit_id, creator_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING id
`

type CreatePostParams struct {
	Title       string `json:"title"`
	Text        string `json:"text"`
	Image       string `json:"image"`
	Video       string `json:"video"`
	Link        string `json:"link"`
	SubredditID int32  `json:"subreddit_id"`
	CreatorID   int32  `json:"creator_id"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) ([]int32, error) {
	rows, err := q.db.QueryContext(ctx, createPost,
		arg.Title,
		arg.Text,
		arg.Image,
		arg.Video,
		arg.Link,
		arg.SubredditID,
		arg.CreatorID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var id int32
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findPostById = `-- name: FindPostById :many
SELECT id, creator_id, subreddit_id FROM posts
WHERE id = $1
`

type FindPostByIdRow struct {
	ID          int32 `json:"id"`
	CreatorID   int32 `json:"creator_id"`
	SubredditID int32 `json:"subreddit_id"`
}

func (q *Queries) FindPostById(ctx context.Context, id int32) ([]FindPostByIdRow, error) {
	rows, err := q.db.QueryContext(ctx, findPostById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindPostByIdRow
	for rows.Next() {
		var i FindPostByIdRow
		if err := rows.Scan(&i.ID, &i.CreatorID, &i.SubredditID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getVote = `-- name: GetVote :many
SELECT post_id, user_id, down FROM vote_post
WHERE post_id = $1 AND user_id = $2
`

type GetVoteParams struct {
	PostID int32 `json:"post_id"`
	UserID int32 `json:"user_id"`
}

type GetVoteRow struct {
	PostID int32 `json:"post_id"`
	UserID int32 `json:"user_id"`
	Down   bool  `json:"down"`
}

func (q *Queries) GetVote(ctx context.Context, arg GetVoteParams) ([]GetVoteRow, error) {
	rows, err := q.db.QueryContext(ctx, getVote, arg.PostID, arg.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetVoteRow
	for rows.Next() {
		var i GetVoteRow
		if err := rows.Scan(&i.PostID, &i.UserID, &i.Down); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeVote = `-- name: RemoveVote :exec
DELETE FROM vote_post
WHERE vote_post.post_id = $1 AND vote_post.user_id = $2
`

type RemoveVoteParams struct {
	PostID int32 `json:"post_id"`
	UserID int32 `json:"user_id"`
}

func (q *Queries) RemoveVote(ctx context.Context, arg RemoveVoteParams) error {
	_, err := q.db.ExecContext(ctx, removeVote, arg.PostID, arg.UserID)
	return err
}

const votePost = `-- name: VotePost :exec
INSERT INTO vote_post (post_id, user_id, down)
VALUES ($1, $2, $3)
ON CONFLICT
DO UPDATE SET down = $3
WHERE vote_post.post_id = $1 AND vote_post.user_id = $2
`

type VotePostParams struct {
	PostID int32 `json:"post_id"`
	UserID int32 `json:"user_id"`
	Down   bool  `json:"down"`
}

func (q *Queries) VotePost(ctx context.Context, arg VotePostParams) error {
	_, err := q.db.ExecContext(ctx, votePost, arg.PostID, arg.UserID, arg.Down)
	return err
}
