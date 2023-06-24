package crypto

import "testing"

const (
	PodUidCipherKey    = "ioyaThaipei5Zosh8eephugae5eex9ei"
	ZoUserUidCipherKey = "sdcBnjNnnsdsNnjasnjasDSjnxsxjsxs"
)

func TestAESEncrypt(t *testing.T) {
	s, err := AESEncrypt([]byte(PodUidCipherKey), "1234567")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}

func TestAESDecrypt(t *testing.T) {
	s, err := AESDecrypt([]byte(PodUidCipherKey), "VMYZiyBP34xb8QfVHpCtuIw0jS+EouT4DKZCf5sCXNQ")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}
