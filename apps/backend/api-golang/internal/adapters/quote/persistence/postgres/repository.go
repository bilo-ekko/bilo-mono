package postgres

import (
	"context"
	"database/sql"

	"api-golang/internal/quote"
	"api-golang/internal/shared/errors"
)

// PostgresRepository implements quote.Repository interface
// This is a CONCRETE EXAMPLE of a driven adapter (outbound).
// This adapter replaces InMemoryRepository when using PostgreSQL.
type PostgresRepository struct {
	db *sql.DB
}

// NewPostgresRepository creates a new PostgreSQL repository adapter
func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

// Create stores a new quote in PostgreSQL
func (r *PostgresRepository) Create(ctx context.Context, quote *quote.Entity) error {
	query := `
		INSERT INTO quotes (
			id, quote_reference, calculation_reference, organisation_id, 
			customer_id, currency, carbon_credit_total, status, 
			expires_at, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	_, err := r.db.ExecContext(ctx, query,
		quote.ID,
		quote.QuoteReference,
		quote.CalculationReference,
		quote.OrganisationID,
		quote.CustomerID,
		quote.Currency,
		quote.CarbonCreditTotal,
		quote.Status,
		quote.ExpiresAt,
		quote.CreatedAt,
		quote.UpdatedAt,
	)

	return err
}

// GetByID retrieves a quote by ID from PostgreSQL
func (r *PostgresRepository) GetByID(ctx context.Context, id string) (*quote.Entity, error) {
	query := `
		SELECT id, quote_reference, calculation_reference, organisation_id,
		       customer_id, currency, carbon_credit_total, status,
		       expires_at, created_at, updated_at
		FROM quotes
		WHERE id = $1
	`

	var q quote.Entity
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&q.ID,
		&q.QuoteReference,
		&q.CalculationReference,
		&q.OrganisationID,
		&q.CustomerID,
		&q.Currency,
		&q.CarbonCreditTotal,
		&q.Status,
		&q.ExpiresAt,
		&q.CreatedAt,
		&q.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.NewNotFoundError("quote", "quote not found")
	}
	if err != nil {
		return nil, err
	}

	return &q, nil
}

// Update updates an existing quote in PostgreSQL
func (r *PostgresRepository) Update(ctx context.Context, quote *quote.Entity) error {
	query := `
		UPDATE quotes
		SET quote_reference = $2, calculation_reference = $3, organisation_id = $4,
		    customer_id = $5, currency = $6, carbon_credit_total = $7, status = $8,
		    expires_at = $9, updated_at = $10
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query,
		quote.ID,
		quote.QuoteReference,
		quote.CalculationReference,
		quote.OrganisationID,
		quote.CustomerID,
		quote.Currency,
		quote.CarbonCreditTotal,
		quote.Status,
		quote.ExpiresAt,
		quote.UpdatedAt,
	)

	return err
}
