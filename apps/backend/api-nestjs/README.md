# api-nest

A NestJS API service implementing equivalents and quotes management for climate impact tracking.

## 🏗️ Project Structure

**Feature-First Organization** - Each feature is a self-contained module.

```
api-nest/
└── src/
    ├── equivalents/                # Climate Equivalents feature
    │   ├── dto/
    │   │   ├── create-equivalent.dto.ts
    │   │   └── update-equivalent.dto.ts
    │   ├── entities/
    │   │   └── equivalent.entity.ts
    │   ├── equivalents.controller.ts
    │   ├── equivalents.service.ts
    │   └── equivalents.module.ts
    ├── quotes/                     # Quotes feature
    │   ├── dto/
    │   │   ├── create-quote.dto.ts
    │   │   └── update-quote.dto.ts
    │   ├── entities/
    │   │   └── quote.entity.ts
    │   ├── quotes.controller.ts
    │   ├── quotes.service.ts
    │   └── quotes.module.ts
    ├── app.module.ts
    └── main.ts
```

## 📦 Feature Modules

### 1. Equivalents Module
Manages carbon equivalents for different activities.

**Entity:**
```typescript
{
  id: string;
  category: string;      // e.g., "transportation", "energy", "lifestyle"
  value: number;         // CO2 amount
  unit: string;          // e.g., "kg CO2"
  description: string;
  createdAt: Date;
  updatedAt: Date;
}
```

**Endpoints:**
- `GET /equivalents` - Get all equivalents
- `GET /equivalents?category=transportation` - Filter by category
- `GET /equivalents/:id` - Get specific equivalent
- `POST /equivalents` - Create new equivalent
- `PATCH /equivalents/:id` - Update equivalent
- `DELETE /equivalents/:id` - Delete equivalent

### 2. Quotes Module
Manages customer quotes for climate impact products/services.

**Entity:**
```typescript
{
  id: string;
  customerId: string;
  items: QuoteItem[];
  totalAmount: number;
  currency: string;
  status: QuoteStatus;   // draft, pending, accepted, rejected, expired
  validUntil: Date;
  createdAt: Date;
  updatedAt: Date;
}
```

**Endpoints:**
- `GET /quotes` - Get all quotes
- `GET /quotes?customerId=123` - Filter by customer
- `GET /quotes?status=pending` - Filter by status
- `GET /quotes/:id` - Get specific quote
- `POST /quotes` - Create new quote
- `PATCH /quotes/:id` - Update quote (including status)
- `DELETE /quotes/:id` - Delete quote

## 🚀 Getting Started

**Server runs on:** `http://localhost:3000`

### Run Development Server
```bash
moon run api-nest:dev
# or
npm run start:dev
```

### Build
```bash
moon run api-nest:build
# or
npm run build
```

### Run Production
```bash
moon run api-nest:start
# or
npm run start:prod
```

## 📡 API Examples

### Equivalents

**Get all equivalents:**
```bash
curl http://localhost:3000/equivalents
```

**Filter by category:**
```bash
curl http://localhost:3000/equivalents?category=transportation
```

**Create equivalent:**
```bash
curl -X POST http://localhost:3000/equivalents \
  -H "Content-Type: application/json" \
  -d '{
    "category": "energy",
    "value": 0.5,
    "unit": "kg CO2",
    "description": "1 kWh of electricity"
  }'
```

### Quotes

**Get all quotes:**
```bash
curl http://localhost:3000/quotes
```

**Filter by customer:**
```bash
curl http://localhost:3000/quotes?customerId=customer-1
```

**Create quote:**
```bash
curl -X POST http://localhost:3000/quotes \
  -H "Content-Type: application/json" \
  -d '{
    "customerId": "customer-3",
    "items": [
      {
        "productId": "carbon-offset-1",
        "quantity": 100,
        "unitPrice": 15,
        "totalPrice": 1500
      }
    ],
    "currency": "USD"
  }'
```

**Update quote status:**
```bash
curl -X PATCH http://localhost:3000/quotes/1 \
  -H "Content-Type: application/json" \
  -d '{"status": "accepted"}'
```

## 🧪 Testing

```bash
# Unit tests
moon run api-nest:test
# or
npm run test

# E2E tests
moon run api-nest:test-e2e
# or
npm run test:e2e

# Test coverage
moon run api-nest:test-cov
# or
npm run test:cov
```

## 🔧 Other Commands

```bash
# Format code
moon run api-nest:format
# or
npm run format

# Lint code
moon run api-nest:lint
# or
npm run lint
```

## 📝 Sample Data

The application comes pre-seeded with sample data:

**Equivalents:**
- Average car trip of 10 miles (4.6 kg CO2)
- 1 kWh of electricity (0.5 kg CO2)
- One meal with beef (2.5 kg CO2)

**Quotes:**
- Customer 1: Carbon offset + Tree planting (Total: $2,750)
- Customer 2: Solar energy project (Total: $6,000)

## 🏛️ Architecture

This project follows NestJS best practices with feature-first organization:

### Feature Module Structure
Each feature contains:
1. **DTOs** (`dto/`) - Data Transfer Objects for validation
2. **Entities** (`entities/`) - Domain models
3. **Controller** - HTTP request handlers
4. **Service** - Business logic
5. **Module** - Feature module definition

### Benefits of Feature-First
- ✅ **High Cohesion** - Related code stays together
- ✅ **Easy Navigation** - All quote-related code in `quotes/`
- ✅ **Scalability** - Add features without touching others
- ✅ **Clear Boundaries** - Each feature is independent
- ✅ **Team-Friendly** - Parallel development without conflicts

## 🔗 Related Projects

- `api-golang` - Go API with similar feature structure
- `web-dashboard` - Next.js dashboard
- `web-sdks` - SvelteKit UI components
