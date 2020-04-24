package firebaseclient

import (
	"context"
	"encoding/json"

	"logbook/internal/config"
	"logbook/pkg/errors"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var (
	shadredClient = &firestore.Client{}
	credentials   = map[string]string{
		"type":                        "service_account",
		"project_id":                  "neogenesis-2a947",
		"private_key_id":              "382d5d9682869f29ed27fac6e7e91691b8471118",
		"private_key":                 "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDHYojkp08jiEL8\n77rVMj/3u/sqMnCAZy8jk+J4wAPjvkpHKvdfjJgO2ESCqyME5ca3e1uSUqNqymBK\nG+SN7umQpLkkJmsPY+haZOzKH8Cma04Cb4f/ksUKECf3fO6pjmvzCuvEqeIzO+KZ\nhXC+G59LlpZ+4eR95t6stqKbzN2LNBcjv3LNw+eewl/PXIh+E2+l8klMfyQkZ1M/\nxccqGQczjMForeVBucF3woGe2t/Uln4Dy2/gRQ6ZyeaPsxpeIVO8ozjAUL6qtOOK\nD7ulE8N8/TCH8240RPCb+N7pIuCSstLu84U6xNwq104ZTxXMp0vsbAlsK1UADC2o\neQ57j1n3AgMBAAECggEAPGGLwWslnUYbu0vCUeQ42QZVps7FoS0eanDTqevEgcjT\n57+MUKRcPEqkXMoE/eu0BUbXUXWzRZjqFidiTNVoaERqSRdVXsDL0ew0hXWeOwfO\nBRDPp8dD4qtH3zw3bqPR7zWEdvdXqFpfYky1+uyUjiiZhO1V7AbpsnZQwkwFX60q\n51/RSu9j02I/0WB9MLuu7pyHYH0V88J2sHmRZlOYjKfQ6g0sMy23DXejkLm+Me/S\nMmI77tuxdwKQq7lhdyYknbagD3dFhtOgFvpAru7IsmOwUkV3Zd7+Z76YHmtf5i6v\nKS9a1m4SWKqX5KTI1K31zD40r33g9BxqZanWt5O+aQKBgQDkqRmAp+MAo75Bgvek\n8slL9gdthg4C8Xj8+szQ9rQQf+FrFzzEfhqqO5aSMSEM80UoAMKZUKJ5QGHQj+XQ\nvhTaKwX4dwjI4jtVkg3dfuqvvjyK0WVR81gXkeCVyLD76Azd4jihmsBV1OIev8Xw\nv+kGZFaNulGa+sXJACZ9Y+EbUwKBgQDfOVvVmzwZVwIrBXXX4X46g5d0sFCRV47R\nBuNLJx1qlI2PR8eY8xBZ3gPpBrM/F1zq0ju4JRIxA6fjC7TraAw4TakyvdD7mWE6\njuQ+b26KkMOHBxWTlnofmTtBdvnKIHhbIau7oQmv5hAByd/bPNG+ESvF5E2Gc6uD\nItTC1lsWTQKBgEGcd/HfupYzZeM8ZjzRYYtHVEhL//c+PJ4yhNStBWv6lb750vj6\nykHwiTr3l+k0YLKizTPa3FP4cli7AZGNCUS0taje1wl1SNs75niY81RsPMkEXpNO\nRQe2VwADry0JH7AHOoK8edXnGN8D3hM/reikn+VZAhN7Q4dyYJj4MN57AoGAW3he\nmEY3N07/FW/0XRUKq6GWo3tIWHZfkGnZI8eQd0cORuMr+BHJ5nlics+LDdRd/99M\nF2YhcTJFTnVSmpAVSdEQ/zjm9OJiBCPfhXIWAvPVG3Scs9CalG4u6OF+9AOgLf+7\nJnSTyp9keccmUuk2LUNOA5ZM6kZlxqO0OW6q/pUCgYBZdikB97p2kY48WGAseb5s\n4CYGXKOepYNf1ifn8aDjFShuEyGrpgkxDwPPa7BTZWZ2FFICW2G6LaLKqoo6nxgn\no6lJDxXH1k6Iernv67W7MRx4SZyIr507Pupke9+AzyIR8DAZ0KGwFkCv5njz/VND\ndLIwl9YcQyXUij6srYKOAw==\n-----END PRIVATE KEY-----\n",
		"client_email":                "firebase-adminsdk-7zlk2@neogenesis-2a947.iam.gserviceaccount.com",
		"client_id":                   "110912029148578276583",
		"auth_uri":                    "https://accounts.google.com/o/oauth2/auth",
		"token_uri":                   "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url":        "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-7zlk2%40neogenesis-2a947.iam.gserviceaccount.com",
	}
)

// Client ...
type Client struct {
	Client *firestore.Client
}

// NewClient ...
func NewClient(cfg *config.Config) (*Client, error) {
	var c Client
	cb, err := json.Marshal(credentials)
	if err != nil {
		return &c, errors.Wrap(err, "[FIREBASE] Failed to marshal credentials!")
	}

	option := option.WithCredentialsJSON(cb)
	c.Client, err = firestore.NewClient(context.Background(), cfg.Firebase.ProjectID, option)
	if err != nil {
		return &c, errors.Wrap(err, "[FIREBASE] Failed to initiate firebase client!")

	}
	return &c, err
}
