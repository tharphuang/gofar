package ghttp

import (
	"context"
	"net/http"
)

type Client struct {
	http.Client
	ctx    context.Context
	parent *Client
}
