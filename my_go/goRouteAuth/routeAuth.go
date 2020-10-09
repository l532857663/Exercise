package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Accounts map[string]string
type HandlerFunc func(*Context)

type Context struct {
}

const AuthUserKey = "user"

type authPair struct {
	value string
	user  string
}

type authPairs []authPair

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	BasicAuthForRealm(Accounts{
		"user1": "love",
		"user2": "god",
		"user3": "sex",
	}, "")

	authorized := r.Group("/")
	authorized.GET("/secret", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"secret": "The secret ingredient to the BBQ sauce is stiring it in an old whiskey barrel.",
		})
	})
	r.Run("127.0.0.1:4396") // listen and serve on 0.0.0.0:8080
}

func (a authPairs) searchCredential(authValue string) (string, bool) {
	if authValue == "" {
		return "", false
	}
	for _, pair := range a {
		if pair.value == authValue {
			return pair.user, true
		}
	}
	return "", false
}

func BasicAuthForRealm(accounts Accounts, realm string) {
	if realm == "" {
		realm = "Authorization Required"
	}
	realm = "Basic realm=" + strconv.Quote(realm)
	pairs := processAccounts(accounts)
	fmt.Printf("%+v\n", pairs)
	/*
		return func(c *Context) {
			// Search user in the slice of allowed credentials
			user, found := pairs.searchCredential(c.requestHeader("Authorization"))
			if !found {
				// Credentials doesn't match, we return 401 and abort handlers chain.
				c.Header("WWW-Authenticate", realm)
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			// The user credentials was found, set user's id to key AuthUserKey in this context, the user's id can be read later using
			// c.MustGet(gin.AuthUserKey).
			c.Set(AuthUserKey, user)
		}
	*/
}

func processAccounts(accounts Accounts) authPairs {
	pairs := make(authPairs, 0, len(accounts))
	for user, password := range accounts {
		value := authorizationHeader(user, password)
		pairs = append(pairs, authPair{
			value: value,
			user:  user,
		})
	}
	return pairs
}

func authorizationHeader(user, password string) string {
	base := user + ":" + password
	// return "Basic " + base64.StdEncoding.EncodeToString(bytesconv.StringToBytes(base))
	return base
}
