package github

import (
    "crypto/hmac"
    "crypto/sha1"
    "encoding/hex"
    "github.com/gin-gonic/gin"
    "io/ioutil"
)

func GithubHandler(c *gin.Context) {
    payload, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
        panic(err)
    }
    signature := c.GetHeader("X-Hub-Signature")
    hash := hmac.New(sha1.New, []byte("yili"))
    hash.Write(payload)
    hash.Sum(nil)
    encoding := hex.EncodeToString(hash.Sum(nil))
    if !hmac.Equal([]byte(signature[5:], ), []byte(encoding)) {
        c.JSON(502, "signature error")
        return
    }
    c.JSON(200, "success")
}
