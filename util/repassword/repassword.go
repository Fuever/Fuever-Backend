package repassword

import "crypto/md5"

func MD5(text string) string {
	key := md5.Sum([]byte(text))
	return string(key[:])
}

func SaltHash(text, salt string) string {
	return MD5(MD5(text) + salt)
}

func ProofingHashes(saltHash, password, salt string) bool {
	return saltHash == SaltHash(password, salt)
}
