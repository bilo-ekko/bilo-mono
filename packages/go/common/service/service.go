package service

// BaseService provides a generic base struct for services that depend on a repository.
// This reduces boilerplate when implementing service structs across domains.
//
// Example usage:
//
//	type Repository interface {
//	    GetByID(ctx context.Context, id string) (*Entity, error)
//	}
//
//	type DefaultService struct {
//	    service.BaseService[Repository]
//	}
//
//	func NewService(repo Repository) *DefaultService {
//	    return &DefaultService{
//	        BaseService: service.NewBaseService(repo),
//	    }
//	}
//
//	func (s *DefaultService) GetEntity(ctx context.Context, id string) (*Entity, error) {
//	    return s.Repo.GetByID(ctx, id)
//	}
type BaseService[R any] struct {
	Repo R
}

// NewBaseService creates a new BaseService with the given repository.
func NewBaseService[R any](repo R) BaseService[R] {
	return BaseService[R]{
		Repo: repo,
	}
}
