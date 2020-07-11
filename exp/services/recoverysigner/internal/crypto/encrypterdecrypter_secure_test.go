package crypto

import (
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/google/tink/go/hybrid"
	"github.com/google/tink/go/integration/awskms"
	"github.com/google/tink/go/keyset"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSecureEncrypterDecrypter_localKMS(t *testing.T) {
	// Connect to KMS running locally.
	// This test requires running local-kms on port 8080.
	awsEndpoint := "http://localhost:8080"
	awsSession := session.Must(session.NewSession())
	awsKMS := kms.New(awsSession, aws.NewConfig().WithEndpoint(awsEndpoint).WithRegion("eu-west-2"))
	awsKeyARN := "arn:aws:kms:eu-west-2:111122223333:key/bfb5861b-796a-448d-98da-fa35d091e0f7"

	// Setup Tink KMS client.
	keyURI := "aws-kms://" + awsKeyARN
	client, err := awskms.NewClientWithKMS(keyURI, awsKMS)
	require.NoError(t, err)

	// Create Tink keyset.
	khPriv, err := keyset.NewHandle(hybrid.ECIESHKDFAES128GCMKeyTemplate())
	require.NoError(t, err)

	// Encrypt the keyset.
	aead, err := client.GetAEAD(keyURI)
	require.NoError(t, err)
	keysetEncrypted := strings.Builder{}
	err = khPriv.Write(keyset.NewJSONWriter(&keysetEncrypted), aead)
	require.NoError(t, err)
	t.Log("Keyset encrypted:", keysetEncrypted.String())

	// Create secure encrypter and decrypter using keyset encrypted by AWS KMS.
	enc, dec, err := newSecureEncrypterDecrypter(client, keyURI, keysetEncrypted.String())
	require.NoError(t, err)
	assert.NotNil(t, enc)
	assert.NotNil(t, dec)

	// Encrypt and decrypt some text and check that it comes back the same.
	pt := []byte("abcdefghijklmnopqrstuvwxyz")
	t.Log("Plaintext:", pt)

	ct, err := enc.Encrypt(pt, nil)
	require.NoError(t, err)
	t.Log("Ciphertext:", ct)
	assert.NotEqual(t, pt, ct)

	pt2, err := dec.Decrypt(ct, nil)
	require.NoError(t, err)
	t.Log("Plaintext rounttrip:", pt2)
	assert.Equal(t, pt, pt2)
}

func TestNewSecureEncrypterDecrypter(t *testing.T) {
	ksPriv := generateKeysetEncrypted(t, hybrid.ECIESHKDFAES128GCMKeyTemplate())
	enc, dec, err := newSecureEncrypterDecrypter(mockKMSClient{}, "mock-key-uri", ksPriv)
	require.NoError(t, err)
	assert.NotNil(t, enc)
	assert.NotNil(t, dec)

	enc, dec, err = newSecureEncrypterDecrypter(mockKMSClient{}, "mock-key-uri", "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "reading encrypted keyset")
	assert.Nil(t, enc)
	assert.Nil(t, dec)
}

func TestSecureEncrypterDecrypter_encryptDecrypt(t *testing.T) {
	ksPriv := generateKeysetEncrypted(t, hybrid.ECIESHKDFAES128GCMKeyTemplate())
	enc, dec, err := newSecureEncrypterDecrypter(mockKMSClient{}, "mock-key-uri", ksPriv)
	require.NoError(t, err)

	plaintext := []byte("secure message")
	contextInfo := []byte("context info")
	ciphertext, err := enc.Encrypt(plaintext, contextInfo)
	require.NoError(t, err)

	plaintext2, err := dec.Decrypt(ciphertext, contextInfo)
	require.NoError(t, err)
	assert.Equal(t, plaintext, plaintext2)

	// context info not matching will result in a failed decryption
	_, err = dec.Decrypt(ciphertext, []byte("wrong info"))
	assert.Error(t, err)
}

func TestNewSecureEncrypterDecrypter_rotatedKeyset(t *testing.T) {
	ksPriv1 := generateKeysetEncrypted(t, hybrid.ECIESHKDFAES128GCMKeyTemplate())

	// add an additional ECIESHKDFAES128GCM Key
	ksPriv2 := rotateKeysetEncrypted(t, ksPriv1, hybrid.ECIESHKDFAES128GCMKeyTemplate())
	enc, dec, err := newSecureEncrypterDecrypter(mockKMSClient{}, "mock-key-uri", ksPriv2)
	require.NoError(t, err)
	assert.NotNil(t, enc)
	assert.NotNil(t, dec)

	// add a new ECIESHKDFAES128CTRHMACSHA256 Key on top of the current ECIESHKDFAES128GCM Key
	ksPriv3 := rotateKeysetEncrypted(t, ksPriv1, hybrid.ECIESHKDFAES128CTRHMACSHA256KeyTemplate())
	enc, dec, err = newSecureEncrypterDecrypter(mockKMSClient{}, "mock-key-uri", ksPriv3)
	require.NoError(t, err)
	assert.NotNil(t, enc)
	assert.NotNil(t, dec)
}

func TestSecureEncrypterDecrypter_rotatedKeysetEncryptDecrypt(t *testing.T) {
	ksPriv1 := generateKeysetEncrypted(t, hybrid.ECIESHKDFAES128GCMKeyTemplate())
	enc1, dec1, err := newSecureEncrypterDecrypter(mockKMSClient{}, "mock-key-uri", ksPriv1)
	require.NoError(t, err)

	// add an additional ECIESHKDFAES128GCM Key
	ksPriv2 := rotateKeysetEncrypted(t, ksPriv1, hybrid.ECIESHKDFAES128GCMKeyTemplate())
	enc2, dec2, err := newSecureEncrypterDecrypter(mockKMSClient{}, "mock-key-uri", ksPriv2)
	require.NoError(t, err)

	plaintext := []byte("secure message")
	contextInfo := []byte("context info")

	// verify that the new keyset private is able to decrypt what the new keyset public encrypts
	ciphertext, err := enc2.Encrypt(plaintext, contextInfo)
	require.NoError(t, err)

	plaintext2, err := dec2.Decrypt(ciphertext, contextInfo)
	require.NoError(t, err)
	assert.Equal(t, plaintext, plaintext2)

	// context info not matching will result in a failed decryption
	_, err = dec2.Decrypt(ciphertext, []byte("wrong info"))
	assert.Error(t, err)

	// verify that the new keyset is able to decrypt what the old keyset encrypts
	ciphertext, err = enc1.Encrypt(plaintext, contextInfo)
	require.NoError(t, err)

	plaintext2, err = dec2.Decrypt(ciphertext, contextInfo)
	require.NoError(t, err)
	assert.Equal(t, plaintext, plaintext2)

	// context info not matching will still result in a failed decryption
	_, err = dec2.Decrypt(ciphertext, []byte("wrong info"))
	assert.Error(t, err)

	// verify that the old keyset is not able to decrypt what the new keyset encrypts
	ciphertext, err = enc2.Encrypt(plaintext, contextInfo)
	require.NoError(t, err)

	_, err = dec1.Decrypt(ciphertext, contextInfo)
	assert.Error(t, err)
}

func TestSecureEncrypterDecrypter_rotatedKeysetMixedKeysEncryptDecrypt(t *testing.T) {
	ksPriv1 := generateKeysetEncrypted(t, hybrid.ECIESHKDFAES128GCMKeyTemplate())
	enc1, dec1, err := newSecureEncrypterDecrypter(mockKMSClient{}, "mock-key-uri", ksPriv1)
	require.NoError(t, err)

	// add a new ECIESHKDFAES128CTRHMACSHA256 Key on top of the current ECIESHKDFAES128GCM Key
	ksPriv2 := rotateKeysetEncrypted(t, ksPriv1, hybrid.ECIESHKDFAES128CTRHMACSHA256KeyTemplate())
	enc2, dec2, err := newSecureEncrypterDecrypter(mockKMSClient{}, "mock-key-uri", ksPriv2)
	require.NoError(t, err)

	plaintext := []byte("secure message")
	contextInfo := []byte("context info")

	// verify that the new keyset private is able to decrypt what the new keyset public encrypts
	ciphertext, err := enc2.Encrypt(plaintext, contextInfo)
	require.NoError(t, err)

	plaintext2, err := dec2.Decrypt(ciphertext, contextInfo)
	require.NoError(t, err)
	assert.Equal(t, plaintext, plaintext2)

	// context info not matching will result in a failed decryption
	_, err = dec2.Decrypt(ciphertext, []byte("wrong info"))
	assert.Error(t, err)

	// verify that the new keyset is able to decrypt what the old keyset encrypts
	ciphertext, err = enc1.Encrypt(plaintext, contextInfo)
	require.NoError(t, err)

	plaintext2, err = dec2.Decrypt(ciphertext, contextInfo)
	require.NoError(t, err)
	assert.Equal(t, plaintext, plaintext2)

	// context info not matching will still result in a failed decryption
	_, err = dec2.Decrypt(ciphertext, []byte("wrong info"))
	assert.Error(t, err)

	// verify that the old keyset is not able to decrypt what the new keyset encrypts
	ciphertext, err = enc2.Encrypt(plaintext, contextInfo)
	require.NoError(t, err)

	_, err = dec1.Decrypt(ciphertext, contextInfo)
	assert.Error(t, err)
}
