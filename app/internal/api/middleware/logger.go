package middleware

// func LoggerMiddleware(l *logger.Logger) fiber.Handler {
// 	return func(c fiber.Ctx) error {
// 		requestID := requestid.FromContext(c)

// 		key, val := logger.AddToCtx(l.With(
// 			zap.String("request_id", requestID),
// 		))

// 		c.Context().SetUserValue(key, val)
// 		c.Locals(key, val)
// 		c.SetUserContext(context.WithValue(c.UserContext(), key, val))

// 		return c.Next()
// 	}
// }

// func LoggerRequestMiddleware(l *logger.Logger) fiber.Handler {
// 	return fiberlogger.New(fiberlogger.Config{
// 		LoggerFunc: func(c fiber.Ctx, data *fiberlogger.Data, cfg fiberlogger.Config) error {
// 			statusCode := c.Response().StatusCode()
// 			body := c.Body()
// 			responseBody := c.Response().Body()

// 			fields := []zap.Field{
// 				zap.String("module", "fiber"),
// 				zap.String("method", c.Method()),
// 				zap.String("path", c.Path()),
// 				zap.String("ip", c.IP()),
// 				zap.Int("status_code", statusCode),
// 				zap.Duration("latency", data.Stop.Sub(data.Start)),
// 				zap.String("request_id", requestid.FromContext(c)),
// 				zap.String("body", string(body)),
// 				zap.String("response_body", string(responseBody)),
// 			}

// 			if data.ChainErr != nil {
// 				fields = append(fields, zap.Error(data.ChainErr))
// 			}

// 			if wallet := c.Params("wallet"); wallet != "" {
// 				fields = append(fields, zap.String("wallet", wallet))
// 			}

// 			if orderID := c.Params("order_id"); orderID != "" {
// 				fields = append(fields, zap.String("order_id", orderID))
// 			}

// 			if actionID := c.Params("action_id"); actionID != "" {
// 				fields = append(fields, zap.String("action_id", actionID))
// 			}

// 			switch {
// 			case statusCode >= 500:
// 				l.Error("handle request", fields...)
// 			case statusCode >= 400:
// 				l.Warn("handle request", fields...)
// 			default:
// 				l.Info("handle request", fields...)
// 			}

// 			return nil
// 		},
// 	})
// }
