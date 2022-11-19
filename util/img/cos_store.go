package img

import (
	"context"
	"crypto/rand"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var (
	bucketUrl = os.Getenv("COS_BUCKET_URL")
	SecretID  = os.Getenv("COS_SECRET_ID")
	SecretKey = os.Getenv("COS_SECRET_KEY")
)

func SaveImage(file io.Reader) string {
	u, _ := url.Parse(bucketUrl)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  SecretID,
			SecretKey: SecretKey,
		},
	})
	randInt, _ := rand.Int(rand.Reader, big.NewInt(100))
	key := strconv.FormatInt(time.Now().UnixNano()+randInt.Int64()%100, 10) + ".jpg"
	println(key)
	_, err := client.Object.Put(context.Background(), key, file, nil)
	if err != nil {
		panic(err)
	}
	return bucketUrl + "/" + key
}
