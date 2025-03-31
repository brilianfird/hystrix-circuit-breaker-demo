package cb

import "github.com/afex/hystrix-go/hystrix"

// BreakerConfig holds the configuration for circuit breaker
type BreakerConfig struct {
	Timeout                int
	MaxConcurrentRequests  int
	RequestVolumeThreshold int
	SleepWindow            int
	ErrorPercentThreshold  int
}

// DefaultBreakerConfig returns default configuration
func DefaultBreakerConfig() BreakerConfig {
	return BreakerConfig{
		Timeout:                1000,
		MaxConcurrentRequests:  100,
		RequestVolumeThreshold: 5,
		SleepWindow:            5000,
		ErrorPercentThreshold:  50,
	}
}

// ConfigureBreaker configures the circuit breaker with the given configuration
func ConfigureBreaker(command string, cfg BreakerConfig) {
	hystrix.ConfigureCommand(command, hystrix.CommandConfig{
		Timeout:                cfg.Timeout,
		MaxConcurrentRequests:  cfg.MaxConcurrentRequests,
		RequestVolumeThreshold: cfg.RequestVolumeThreshold,
		SleepWindow:            cfg.SleepWindow,
		ErrorPercentThreshold:  cfg.ErrorPercentThreshold,
	})
}
