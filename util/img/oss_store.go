package img

import (
	"context"
	"crypto/rand"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var (
	bucketUrl = "https://fuever-1313037799.cos.ap-nanjing.myqcloud.com"
	SecretID  = "AKIDd8QsEVfqwCG6i5E7k7Rr7D4f1Bk25tRm"
	SecretKey = "ECuPgjaBXfaGivtnX3KcnjX5RZimeJo6"
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
	key := strconv.FormatInt(time.Now().UnixNano()+randInt.Int64(), 10) + ".jpg"
	println(key)
	_, err := client.Object.Put(context.Background(), key, file, nil)
	if err != nil {
		panic(err)
	}
	return bucketUrl + "/" + key
}
