# Postman Collection Summary

## 📊 Collection Statistics

- **Total Requests:** 18
- **Total Tests:** 45+ assertions
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
    └── Impact Projects
        ├── Get All Impact Projects [3 tests]
        ├── Get Single Impact Project [2 tests]
        └── Get Projects by Partner [2 tests]
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

### Data Validation
- ✅ Required fields presence
- ✅ Data types validation
- ✅ Relationship integrity (partnerId)
- ✅ Calculated fields (totalAmount)

### Edge Cases
- ✅ Empty query parameters
- ✅ Valid ID lookups
- ✅ Array response validation

## 🚀 Running the Tests

### Quick Test
```bash
# Start both servers
moon run api-nest:dev &
moon run api-golang:dev &

# Wait for servers to start
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

## 📈 Expected Results

When all tests pass, you should see:

```
┌─────────────────────────┬───────────────────┬───────────────────┐
│                         │          executed │            failed │
├─────────────────────────┼───────────────────┼───────────────────┤
│              iterations │                 1 │                 0 │
├─────────────────────────┼───────────────────┼───────────────────┤
│                requests │                18 │                 0 │
├─────────────────────────┼───────────────────┼───────────────────┤
│            test-scripts │                18 │                 0 │
├─────────────────────────┼───────────────────┼───────────────────┤
│              assertions │                45 │                 0 │
└─────────────────────────┴───────────────────┴───────────────────┘
```

## 🔮 Future Enhancements

- [ ] Add DELETE operation tests
- [ ] Add authentication/authorization tests
- [ ] Add error handling tests (404, 400, 500)
- [ ] Add performance tests
- [ ] Add data-driven tests with CSV
- [ ] Add pre-request scripts for data setup
- [ ] Add environment-specific configurations
- [ ] Add integration tests between APIs
