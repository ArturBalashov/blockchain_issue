package httputils

import (
	"context"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"services/internal/repository"
	"services/internal/shared_errors"
)

func AuthFunc(signingKey string, repo repository.Middleware, logger *zap.Logger) auth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		method := strings.Split(grpc.ServerTransportStreamFromContext(ctx).Method(), "/")
		if len(method) < 3 {
			return nil, shared_errors.ErrMetadataFailed
		}
		if method[2] == "Health" {
			return ctx, nil
		}

		meta, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return ctx, shared_errors.ErrMetadataFailed
		}

		authHeader := meta.Get("authorization")
		if len(authHeader) == 0 {
			return ctx, shared_errors.ErrWrongTokenWithStatus
		}
		header := strings.Split(authHeader[0], " ")
		if strings.ToLower(header[0]) != "bearer" {
			return ctx, shared_errors.ErrWrongTokenWithStatus
		}

		token, err := parseToken(header[1], signingKey)
		if err != nil {
			logger.Error("can't parse token", zap.Error(err))
			return ctx, shared_errors.ErrWrongParseToken
		}

		if err := repo.CheckToken(header[1]); err != nil {
			logger.Error("can't authorization user", zap.Error(err))
			return ctx, shared_errors.ErrWrongTokenWithStatus
		}

		if err := repo.CheckPermission(token.UserID, method[2]); err != nil {
			logger.Error("permission denied", zap.Error(err))
			return ctx, shared_errors.ErrPermissionDenied
		}
		ctx = context.WithValue(ctx, "user_info", UserContext{
			token:  header[1],
			userID: token.UserID,
		})

		return ctx, nil
	}
}

func parseToken(t, signingKey string) (*tokenInfo, error) {
	userToken := &tokenInfo{}

	token, err := jwt.ParseWithClaims(t, userToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, shared_errors.ErrWrongToken
	}

	return userToken, nil
}
