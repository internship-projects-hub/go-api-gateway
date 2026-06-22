package services

type Service struct {
	Name string
	URL  string
}

type Registry struct {
	services map[string]Service
}

func NewRegistry() *Registry {
	return &Registry{
		services: map[string]Service{
			"user-service": {
				Name: "user-service",
				URL:  "http://localhost:8001",
			},
			"payment-service": {
				Name: "payment-service",
				URL:  "http://localhost:8002",
			},
		},
	}
}

func (r *Registry) GetService(name string) (Service, bool) {
	service, exists := r.services[name]
	return service, exists
}
