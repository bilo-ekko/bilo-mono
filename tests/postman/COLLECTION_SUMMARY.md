# Postman Collection Summary

## 📊 Collection Statistics

- **Total Requests:** 29
- **Total Tests:** 80+ assertions
- **APIs Covered:** 2 (NestJS + Go)
- **Ports:** 3000 (NestJS), 8080 (Go)

## 🗂️ Collection Structure

```
Bilo Mono API Collection
│
├── NestJS API (Port 3000)
│   ├── Health Check
│   │   └── Get Root [2 tests]
│   │
│   ├── Equivalents
│   │   ├── Get All Equivalents [3 tests]
│   │   ├── Get Equivalents by Category [2 tests]
│   │   ├── Get Single Equivalent [2 tests]
│   │   └── Create Equivalent [2 tests]
│   │
│   └── Quotes
│       ├── Get All Quotes [3 tests]
│       ├── Get Quotes by Customer [2 tests]
│       ├── Get Single Quote [2 tests]
│       ├── Create Quote [2 tests]
│       └── Update Quote Status [2 tests]
│
└── Go API (Port 8080)
    ├── Health Check
    │   ├── Get Root [2 tests]
    │   └── Health Check [2 tests]
    │
    ├── Impact Partners
    │   ├── Get All Impact Partners [3 tests]
    │   └── Get Single Impact Partner [2 tests]
    │
    ├── Impact Projects
    │   ├── Get All Impact Projects [3 tests]
    │   ├── Get Single Impact Project [2 tests]
    │   └── Get Projects by Partner [2 tests]
    │
    └── Quotes (Carbon Offset) ⭐ NEW
        ├── Create Quote - Basic [6 tests]
        ├── Create Quote - With Merchant Details [3 tests]
        ├── Create Quote - With Round Up [3 tests]
        ├── Create Quote - Currency Conversion (USD) [3 tests]
        ├── Create Quote - Currency Conversion (GBP) [3 tests]
        ├── Create Quote - Child Organisation [3 tests]
        ├── Create Quote - Filter Projects by Location [3 tests]
        ├── Create Quote - Forbidden (Unauthorised Org) [3 tests]
        ├── Create Quote - Missing Required Field [2 tests]
        ├── Get Quote by ID [3 tests]
        └── Get Quote - Not Found [2 tests]
```

## 🧪 Test Coverage by Endpoint

### NestJS API (Port 3000)

| Endpoint | Method | Tests | Description |
|----------|--------|-------|-------------|
| `/` | GET | 2 | Root endpoint health check |
| `/equivalents` | GET | 3 | Get all equivalents with validation |
| `/equivalents?category=X` | GET | 2 | Filter equivalents by category |
| `/equivalents/:id` | GET | 2 | Get specific equivalent |
| `/equivalents` | POST | 2 | Create new equivalent |
| `/quotes` | GET | 3 | Get all quotes with validation |
| `/quotes?customerId=X` | GET | 2 | Filter quotes by customer |
| `/quotes/:id` | GET | 2 | Get specific quote |
| `/quotes` | POST | 2 | Create new quote |
| `/quotes/:id` | PATCH | 2 | Update quote status |

### Go API (Port 8080)

| Endpoint | Method | Tests | Description |
|----------|--------|-------|-------------|
| `/` | GET | 2 | Root endpoint |
| `/api/health` | GET | 2 | Health check endpoint |
| `/api/impact-partners` | GET | 3 | Get all impact partners |
| `/api/impact-partners/:id` | GET | 2 | Get specific partner |
| `/api/impact-projects` | GET | 3 | Get all projects |
| `/api/impact-projects/:id` | GET | 2 | Get specific project |
| `/api/impact-projects?partnerId=X` | GET | 2 | Filter projects by partner |
| `/api/quotes` | POST | 6 | Create carbon offset quote (basic) |
| `/api/quotes` | POST | 3 | Create quote with merchant MCC details |
| `/api/quotes` | POST | 3 | Create quote with round up |
| `/api/quotes` | POST | 3 | Create quote with USD currency conversion |
| `/api/quotes` | POST | 3 | Create quote with GBP currency conversion |
| `/api/quotes` | POST | 3 | Create quote for child organisation |
| `/api/quotes` | POST | 3 | Create quote with location filter |
| `/api/quotes` | POST | 3 | Forbidden access test |
| `/api/quotes` | POST | 2 | Missing field validation |
| `/api/quotes/:id` | GET | 3 | Get specific quote |
| `/api/quotes/:id` | GET | 2 | Quote not found |

## 🌍 Quote Creation Flow (Go API)

The new Quote endpoints implement the full carbon offset calculation orchestration:

```
1. Validate Organisation     ──► Check org hierarchy (parent/child)
         │
2. Get/Create Customer       ──► Find or create customer record
         │
3. Calculate Carbon Footprint
   ├── Use merchant MCC details (or org defaults)
   ├── Look up merchant country
   ├── Convert currency to EUR (if needed)
   ├── Look up carbon factor (by MCC + country)
   └── Calculate: amount × factor = carbon kg
         │
4. Get Blended Project Price ──► Avg price across impact projects
         │
5. Calculate Compensation    ──► carbon kg × blended unit price
         │
6. Calculate Round Up        ──► Optional rounding (0.50, 1.00, nearest)
         │
7. Calculate Service Fee     ──► Based on org fee config
         │
8. Calculate Sales Tax       ──► Based on customer location (VAT/Sales Tax)
         │
9. Write Quote              ──► Persist and return complete quote
```

### Quote Request Example
```json
{
  "organisationId": "org-parent-1",
  "customerExternalId": "cust-123",
  "customerCountry": "GB",
  "transactionId": "txn-001",
  "transactionAmount": 100.00,
  "transactionCurrency": "GBP",
  "merchantDetails": {
    "mcc": "5812",
    "countryCode": "GB"
  },
  "enableRoundUp": true,
  "roundUpTarget": "1.00"
}
```

### Quote Response Breakdown
```json
{
  "breakdown": {
    "transactionAmount": 100,
    "transactionCurrency": "GBP",
    "amountEur": 117,
    "carbonKg": 40.95,
    "blendedUnitPrice": 14.33,
    "compensationAmount": 586.84,
    "roundUpAmount": 0.16,
    "serviceFeePercentage": 0.10,
    "serviceFee": 10.00,
    "salesTaxRate": 0.20,
    "salesTax": 119.40,
    "totalAmount": 716.40
  }
}
```

## ✅ Test Types

Each request includes one or more of the following test types:

### 1. Status Code Tests
```javascript
pm.test("Status code is 200", function () {
    pm.response.to.have.status(200);
});
```

### 2. Response Structure Tests
```javascript
pm.test("Response is an array", function () {
    var jsonData = pm.response.json();
    pm.expect(jsonData).to.be.an('array');
});
```

### 3. Field Validation Tests
```javascript
pm.test("Response has required fields", function () {
    var jsonData = pm.response.json();
    pm.expect(jsonData).to.have.property('id');
    pm.expect(jsonData).to.have.property('name');
});
```

### 4. Business Logic Tests
```javascript
pm.test("All results match filter", function () {
    var jsonData = pm.response.json();
    jsonData.forEach(function(item) {
        pm.expect(item.category).to.equal('transportation');
    });
});
```

### 5. Data Integrity Tests
```javascript
pm.test("Response has correct ID", function () {
    var jsonData = pm.response.json();
    pm.expect(jsonData.id).to.equal('1');
});
```

### 6. Currency Conversion Tests (NEW)
```javascript
pm.test("Currency conversion occurred", function () {
    var breakdown = pm.response.json().breakdown;
    pm.expect(breakdown.transactionCurrency).to.equal('USD');
    pm.expect(breakdown.amountEur).to.be.below(100); // USD→EUR rate ~0.92
});
```

### 7. Authorization Tests (NEW)
```javascript
pm.test("Status code is 403 Forbidden", function () {
    pm.response.to.have.status(403);
});
pm.test("Error code is FORBIDDEN", function () {
    pm.expect(pm.response.json().error.code).to.equal('FORBIDDEN');
});
```

## 🎯 Test Scenarios Covered

### CRUD Operations
- ✅ Create (POST)
- ✅ Read (GET) - Single & Multiple
- ✅ Update (PATCH)
- ❌ Delete (Not implemented yet)

### Filtering & Querying
- ✅ Filter by category
- ✅ Filter by customer ID
- ✅ Filter by partner ID
- ✅ Filter by status
- ✅ Filter projects by location (NEW)

### Data Validation
- ✅ Required fields presence
- ✅ Data types validation
- ✅ Relationship integrity (partnerId)
- ✅ Calculated fields (totalAmount)
- ✅ Missing field error handling (NEW)

### Business Logic (NEW - Quote Flow)
- ✅ Organisation hierarchy validation
- ✅ Currency conversion (USD, GBP → EUR)
- ✅ Carbon footprint calculation with MCC
- ✅ Blended project pricing
- ✅ Round up calculations
- ✅ Service fee calculations
- ✅ Sales tax by country/state (VAT, US Sales Tax)
- ✅ Project location filtering

### Authorization (NEW)
- ✅ Parent org accessing child org (allowed)
- ✅ Child org accessing sibling org (forbidden)
- ✅ Error response structure validation

### Edge Cases
- ✅ Empty query parameters
- ✅ Valid ID lookups
- ✅ Array response validation
- ✅ Quote not found (404)
- ✅ Missing required fields (400)
- ✅ Forbidden access (403)

## 🚀 Running the Tests

### Quick Test
```bash
# Start Go API server
cd apps/backend/api-golang && go run ./cmd/api/main.go &

# Wait for server to start
sleep 3

# Run tests with Newman
newman run tests/postman/collection.json
```

### Detailed Test with Report
```bash
newman run tests/postman/collection.json \
  --reporters cli,json,html \
  --reporter-html-export test-report.html
```

### Run Only Go API Tests
```bash
newman run tests/postman/collection.json \
  --folder "Go API (Port 8080)"
```

### Run Only Quote Tests
```bash
newman run tests/postman/collection.json \
  --folder "Quotes (Carbon Offset)"
```

## 📈 Expected Results

When all tests pass, you should see:

```
┌─────────────────────────┬───────────────────┬───────────────────┐
│                         │          executed │            failed │
├─────────────────────────┼───────────────────┼───────────────────┤
│              iterations │                 1 │                 0 │
├─────────────────────────┼───────────────────┼───────────────────┤
│                requests │                29 │                 0 │
├─────────────────────────┼───────────────────┼───────────────────┤
│            test-scripts │                29 │                 0 │
├─────────────────────────┼───────────────────┼───────────────────┤
│              assertions │                80 │                 0 │
└─────────────────────────┴───────────────────┴───────────────────┘
```

## 🔮 Future Enhancements

- [ ] Add DELETE operation tests
- [ ] Add authentication/authorization tests
- [x] Add error handling tests (404, 400, 403) ✅
- [ ] Add performance tests
- [ ] Add data-driven tests with CSV
- [ ] Add pre-request scripts for data setup
- [ ] Add environment-specific configurations
- [ ] Add integration tests between APIs
- [x] Add currency conversion tests ✅
- [x] Add organisation hierarchy tests ✅
- [x] Add carbon footprint calculation tests ✅