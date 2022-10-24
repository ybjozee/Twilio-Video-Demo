package helper

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/twilio/twilio-go/client/jwt"
	"os"
)

func GenerateToken(roomName string) string {

	params := jwt.AccessTokenParams{
		AccountSid:    os.Getenv("TWILIO_ACCOUNT_SID"),
		SigningKeySid: os.Getenv("TWILIO_API_KEY_SID"),
		Secret:        os.Getenv("TWILIO_API_SECRET_KEY"),
		Identity:      identity(),
	}

	jwtToken := jwt.CreateAccessToken(params)

	videoGrant := &jwt.VideoGrant{
		Room: roomName,
	}

	jwtToken.AddGrant(videoGrant)
	token, err := jwtToken.ToJwt()

	if err != nil {
		fmt.Println(err)
	}

	return token
}

func identity() string {

	input, _ := rand.Prime(rand.Reader, 128)
	hash := md5.Sum([]byte(input.String()))
	return hex.EncodeToString(hash[:])
}
