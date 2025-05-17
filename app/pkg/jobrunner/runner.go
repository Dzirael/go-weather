package jobrunner

import (
	"time"

	"go.uber.org/zap"
)

type Job struct {
	Name     string
	DoFunc   func()
	Interval time.Duration
	Logger   *zap.Logger

	stop    chan struct{}
	stopped chan struct{}
}

func New(logger *zap.Logger, name string, interval time.Duration, doFunc func()) *Job {
	return &Job{
		Name:     name,
		DoFunc:   doFunc,
		Interval: interval,
		Logger:   logger,

		stop:    make(chan struct{}),
		stopped: make(chan struct{}),
	}
}

func (j *Job) Start() {
	j.Logger.Info("job started", zap.String("job", j.Name), zap.Duration("run_interval", j.Interval))
	go j.run()
}

func (j *Job) Stop() {
	j.Logger.Warn("waiting for job to stop", zap.String("job", j.Name))
	close(j.stop)
	<-j.stopped
}

func (j *Job) run() {
	t := time.NewTicker(j.Interval)

	for {
		select {
		case <-j.stop:
			j.Logger.Info("job stopped", zap.String("job", j.Name))
			close(j.stopped)
			return

		case <-t.C:
			now := time.Now()

			defer func() {
				recovered := recover()

				if recovered != nil {
					j.Logger.Error("job panicked", zap.String("job", j.Name), zap.Any("err", recovered), zap.Stack("stack"))
				}
			}()
			j.DoFunc()
			j.Logger.Debug("job ran", zap.String("job", j.Name), zap.Duration("elapsed", time.Since(now)))
		}
	}
}
