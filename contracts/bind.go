package contracts

import "context"

type Bind interface {
	StoreMedia(ctx context.Context)
}
