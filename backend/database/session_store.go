/**
** @see https://velog.io/@byron1st/GoFiber-%EC%97%90-%EC%84%B8%EC%85%98-%EA%B8%B0%EB%B0%98-%EC%9D%B8%EC%A6%9D-%EC%B6%94%EA%B0%80%ED%95%98%EA%B8%B0
**/
package database

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/bbolt"
)

var store *session.Store
var _once sync.Once

func GetSessionStore() *session.Store {
	if store == nil {
		_once.Do(func() {
			storage := bbolt.New(bbolt.Config{
				Database: "session.db",
				Bucket:   "BoltSessionBucket",
				Reset:    false,
			})

			STAGE := os.Getenv("STAGE")
			cookieSecure := false
			if STAGE != "prod" {
				cookieSecure = true
			}

			store = session.New(session.Config{
				Storage:        storage,
				Expiration:     time.Minute * 10,
				CookieSecure:   cookieSecure, // Indicates if cookie is secure.
				CookieHTTPOnly: true,         // client에서 document.cookie로 접근 불가
				CookieSameSite: "None",       // For cross-origin
			})
			fmt.Println("Session create success")
		})
	}

	return store
}

func SetSession(accountId string, c *fiber.Ctx) error {
	sess, err := GetSessionStore().Get(c)

	if err != nil {
		return err
	}

	sess.Set("accountId", accountId)

	// set-cookie 헤더에 자동으로 세션키를 포함한다
	if err := sess.Save(); err != nil {
		return err
	}

	return nil
}

func GetAccountIdFromSession(accountId string, c *fiber.Ctx) (string, error) {
	sess, err := GetSessionStore().Get(c)

	if err != nil {
		return "", err
	}

	raw := sess.Get("accountId")

	if raw == nil {
		return "", errors.New("user not logged in")
	}

	id, ok := raw.(string)

	if !ok {
		return "", errors.New("malformed session")
	}

	return id, nil
}

func DeleteAccountIdFromSession(accountId string, c *fiber.Ctx) error {
	sess, err := GetSessionStore().Get(c)

	if err != nil {
		return err
	}

	sess.Delete(accountId)

	return nil
}
