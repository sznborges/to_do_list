package shutdown

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sznborges/to_do_list/infra/logger"
	"golang.org/x/sync/errgroup"
)

func Subscribe(start, stop func(ctx context.Context) error) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return start(ctx)
	})
	g.Go(func() error {
		<-gCtx.Done()
		return stop(ctx)
	})

	if err := g.Wait(); err != nil {
		logger.Logger.Errorf("exit reason: %s \n", err)
	}
}