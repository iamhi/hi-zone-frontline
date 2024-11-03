package userhandler

import (
	"time"

	"github.com/google/uuid"
	"github.com/iamhi/frontline/internal/errors"
)

type tokenEntry struct {
	Token      string
	Username   string
	ExpireTime time.Time
}

var active_token_entries = []tokenEntry{}

func generateToken(username string) string {
	token := uuid.New().String()

	active_token_entries = append(active_token_entries, tokenEntry{
		Token:      token,
		ExpireTime: time.Now().Add(time.Hour * 48),
		Username:   username,
	})

	return token
}

func getUserName(token string) (string, errors.UserHandlerError) {
	for idx, token_entry := range active_token_entries {
		if token_entry.Token == token {
			if token_entry.ExpireTime.Before(time.Now()) {
				active_token_entries = append(active_token_entries[:idx], active_token_entries[idx+1:]...)

				return "", errors.UserInvalidTokenError{}
			}

			return token_entry.Username, nil
		}
	}

	return "", errors.UserInvalidTokenError{}
}

func replaceToken(old_token string) (string, errors.UserHandlerError) {
	for idx, token_entry := range active_token_entries {
		if token_entry.Token == old_token {
			active_token_entries = append(active_token_entries[:idx], active_token_entries[idx+1:]...)

			return generateToken(token_entry.Username), nil
		}
	}

	return "", errors.UserInvalidTokenError{}
}

func invalidateToken(token string) {
	for idx, token_entry := range active_token_entries {
		if token_entry.Token == token {
			active_token_entries = append(active_token_entries[:idx], active_token_entries[idx+1:]...)
			return
		}
	}

	return
}
