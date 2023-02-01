package platform_train

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func procedure() {
	now := time.Now()
	fmt.Printf("Long procedure started at %+v \n", now)

	rf := rand.Intn(1)
	fmt.Printf("Long procedure will continue %+v seconds \n", rf)

	for i := 0; i <= rf; i++ {
		time.Sleep(time.Second)
	}

	fmt.Println(rand.Int31())
}

func contexter() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

	go func(ctx context.Context, cancel context.CancelFunc) {
		select {
		case <-ctx.Done():
			cancel()
			panic(ctx.Err())
		}
	}(ctx, cancel)
	procedure()

}
