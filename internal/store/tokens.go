package store

import (
	"database/sql"
	"time"

	"github.com/chunyukuo88/workoutsV2/internal/tokens"
)

type PostgresTokenStore struct {
	db *sql.DB
}

func NewPostgresTokenStore(db *sql.DB) *PostgresTokenStore {
	return &PostgresTokenStore{
		db,
	}
}

type TokenStore interface {
	Insert(token *tokens.Token) error
	CreateNewToken(userID int, ttl time.Duration, scope string) (*tokens.Token, error)
	DeleteAllTokensForUser(userID int, scope string) error
}

func (ts *PostgresTokenStore) CreateNewToken(userID int, ttl time.Duration, scope string) (*tokens.Token, error) {
	token, err := tokens.GenerateToken(userID, ttl, scope)
	if err != nil {
		return nil, err
	}

	err = ts.Insert(token)

	return token, err
}

func (ts *PostgresTokenStore) Insert(token *tokens.Token) error {
	query := `
INSERT INTO tokens (hash, user_id, expiry, scope)
VALUES ($1, $2, $3, $4)
`

	_, err := ts.db.Exec(query, token.Hash, token.UserID, token.Expiry, token.Scope)
	return err
}

func (ts *PostgresTokenStore) DeleteAllTokensForUser(userID int, scope string) error {
	query := `
DELETE FROM tokens
WHERE scope = $1 AND user_id = $2
`

	_, err := ts.db.Exec(query, userID, scope)
	return err
}
