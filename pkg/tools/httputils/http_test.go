package httputils

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/metadata"

	mockrepository "services/internal/mocks"
	"services/internal/shared_errors"
)

func TestAuthFunc(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mockrepository.NewMockMiddleware(ctrl)

	expiresAt := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &tokenInfo{
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  time.Now().Unix(),
		},
		2,
		[]int{1},
	})

	tkn, err := token.SignedString([]byte("dsfkj98]{32f"))
	if err != nil {
		t.Fatal(err)
	}

	testingError := errors.New("ok")
	mockRepo.EXPECT().CheckToken(tkn).Return(testingError).Times(1)
	mockRepo.EXPECT().CheckToken(tkn).Return(nil).Times(1)

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stdout,
		zap.DebugLevel,
	))
	aFunc := AuthFunc("dsfkj98]{32f", mockRepo, logger)

	for i := 0; i < 6; i++ {
		ctx := context.Background()
		authPrefix := "authorization"
		bearerPrefix := "bearer "
		if i == 3 {
			authPrefix = "wrong"
		}
		if i == 4 {
			bearerPrefix = "wrong"
		}
		if i == 5 {
			tkn += "wrong"
		}
		if i != 2 {
			ctx = metadata.NewIncomingContext(ctx, metadata.New(map[string]string{authPrefix: bearerPrefix + tkn}))
		}
		_, err := aFunc(ctx)
		if err != nil && !errors.As(err, &shared_errors.ErrWrongToken) && !errors.As(err, &testingError) &&
			!errors.As(err, &shared_errors.ErrMetadataFailed) && !errors.As(err, &shared_errors.ErrWrongTokenWithStatus) &&
			!errors.As(err, &shared_errors.ErrWrongParseToken) {
			t.Fatalf("Auth func failed: %v", err)
		}
	}
}
