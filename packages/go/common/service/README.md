# Service Package

This package provides a generic base service struct to reduce boilerplate when implementing service layers across domains.

## Usage

Instead of writing:

```go
type DefaultService struct {
    repo Repository
}

func NewService(repo Repository) *DefaultService {
    return &DefaultService{repo: repo}
}
```

You can now write:

```go
import "github.com/bilo-mono/packages/common/service"

type DefaultService struct {
    service.BaseService[Repository]
}

func NewService(repo Repository) *DefaultService {
    return &DefaultService{
        BaseService: service.NewBaseService(repo),
    }
}
```

Then access the repository via `s.Repo` instead of `s.repo`.

## Benefits

1. **Reduces boilerplate**: No need to define the struct field and constructor manually
2. **Type-safe**: Uses Go generics to ensure type safety
3. **Consistent**: All services follow the same pattern
4. **Flexible**: Works with any repository interface type

## Examples

### Single Repository

```go
type Repository interface {
    GetByID(ctx context.Context, id string) (*Entity, error)
}

type DefaultService struct {
    service.BaseService[Repository]
}

func NewService(repo Repository) *DefaultService {
    return &DefaultService{
        BaseService: service.NewBaseService(repo),
    }
}

func (s *DefaultService) GetEntity(ctx context.Context, id string) (*Entity, error) {
    return s.Repo.GetByID(ctx, id)
}
```

### Multiple Repositories

For services with multiple repositories, you can still use BaseService for one and add others manually:

```go
type FactorRepository interface {
    GetFactor(ctx context.Context, mcc string, countryID string) (*Factor, error)
}

type FootprintRepository interface {
    Create(ctx context.Context, footprint *Footprint) error
}

type DefaultService struct {
    service.BaseService[FactorRepository]
    footprintRepo FootprintRepository
}

func NewService(factorRepo FactorRepository, footprintRepo FootprintRepository) *DefaultService {
    return &DefaultService{
        BaseService: service.NewBaseService(factorRepo),
        footprintRepo: footprintRepo,
    }
}
```

## Migration Guide

To migrate existing services:

1. Add import: `"github.com/bilo-mono/packages/common/service"`
2. Change struct definition from `repo Repository` to `service.BaseService[Repository]`
3. Update constructor to use `service.NewBaseService(repo)`
4. Change all `s.repo` references to `s.Repo` (capital R)
