// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: subreddit.sql

package database

import (
	"context"
)

const createSubreddit = `-- name: CreateSubreddit :many
INSERT INTO subreddit (
    title, name, creator_id
) VALUES (
    $1, $2, $3
) RETURNING id, name
`

type CreateSubredditParams struct {
	Title     string `json:"title"`
	Name      string `json:"name"`
	CreatorID int32  `json:"creator_id"`
}

type CreateSubredditRow struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) CreateSubreddit(ctx context.Context, arg CreateSubredditParams) ([]CreateSubredditRow, error) {
	rows, err := q.db.QueryContext(ctx, createSubreddit, arg.Title, arg.Name, arg.CreatorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CreateSubredditRow
	for rows.Next() {
		var i CreateSubredditRow
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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

const findSubreddit = `-- name: FindSubreddit :many
SELECT id FROM subreddit
WHERE name = $1
`

func (q *Queries) FindSubreddit(ctx context.Context, name string) ([]int32, error) {
	rows, err := q.db.QueryContext(ctx, findSubreddit, name)
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

const findSubredditById = `-- name: FindSubredditById :many
SELECT id, creator_id FROM subreddit
WHERE id = $1
`

type FindSubredditByIdRow struct {
	ID        int32 `json:"id"`
	CreatorID int32 `json:"creator_id"`
}

func (q *Queries) FindSubredditById(ctx context.Context, id int32) ([]FindSubredditByIdRow, error) {
	rows, err := q.db.QueryContext(ctx, findSubredditById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindSubredditByIdRow
	for rows.Next() {
		var i FindSubredditByIdRow
		if err := rows.Scan(&i.ID, &i.CreatorID); err != nil {
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

const updateSubredditAvatar = `-- name: UpdateSubredditAvatar :exec
UPDATE subreddit 
SET avatar = $2
WHERE id = $1
`

type UpdateSubredditAvatarParams struct {
	ID     int32  `json:"id"`
	Avatar string `json:"avatar"`
}

func (q *Queries) UpdateSubredditAvatar(ctx context.Context, arg UpdateSubredditAvatarParams) error {
	_, err := q.db.ExecContext(ctx, updateSubredditAvatar, arg.ID, arg.Avatar)
	return err
}

const updateSubredditCover = `-- name: UpdateSubredditCover :exec
UPDATE subreddit 
SET cover = $2
WHERE id = $1
`

type UpdateSubredditCoverParams struct {
	ID    int32  `json:"id"`
	Cover string `json:"cover"`
}

func (q *Queries) UpdateSubredditCover(ctx context.Context, arg UpdateSubredditCoverParams) error {
	_, err := q.db.ExecContext(ctx, updateSubredditCover, arg.ID, arg.Cover)
	return err
}

const updateSubredditName = `-- name: UpdateSubredditName :exec
UPDATE subreddit 
SET name = $2
WHERE id = $1
`

type UpdateSubredditNameParams struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateSubredditName(ctx context.Context, arg UpdateSubredditNameParams) error {
	_, err := q.db.ExecContext(ctx, updateSubredditName, arg.ID, arg.Name)
	return err
}

const updateSubredditTitle = `-- name: UpdateSubredditTitle :exec
UPDATE subreddit 
SET title = $2
WHERE id = $1
`

type UpdateSubredditTitleParams struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
}

func (q *Queries) UpdateSubredditTitle(ctx context.Context, arg UpdateSubredditTitleParams) error {
	_, err := q.db.ExecContext(ctx, updateSubredditTitle, arg.ID, arg.Title)
	return err
}
