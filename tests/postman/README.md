# Postman API Testing Collection

This directory contains Postman collections for testing all backend API endpoints in the bilo-mono project.

## 📋 Collection Overview

The `collection.json` file includes comprehensive tests for:

### NestJS API (Port 3000)
- **Health Check** - Root endpoint
- **Equivalents Module**
  - Get all equivalents
  - Filter by category
  - Get single equivalent
  - Create equivalent
- **Quotes Module**
  - Get all quotes
  - Filter by customer
  - Get single quote
  - Create quote
  - Update quote status

### Go API (Port 8080)
- **Health Check** - Root and health endpoints
- **Impact Partners Module**
  - Get all partners
  - Get single partner
- **Impact Projects Module**
  - Get all projects
  - Get single project
  - Filter by partner

## 🧪 Test Coverage

Each request includes at least one unit test:
- ✅ Status code validation
- ✅ Response structure validation
- ✅ Data type checks
- ✅ Field presence validation
- ✅ Filter functionality tests
- ✅ Business logic validation

## 🚀 Getting Started

### Prerequisites

1. **Postman Desktop App** or **Postman CLI** installed
2. Both backend servers running:
   ```bash
   # Terminal 1 - Start NestJS API (port 3000)
   moon run api-nest:dev
   
   # Terminal 2 - Start Go API (port 8080)
   moon run api-golang:dev
   ```

### Option 1: Using Postman Desktop App

1. Open Postman
2. Click **Import** in the top-left
3. Select the `collection.json` file
4. The collection will appear in your sidebar
5. Click the **Run** button to execute all tests

### Option 2: Using Postman CLI (Newman)

Install Newman (Postman CLI):
```bash
npm install -g newman
```

Run the collection:
```bash
# Run all tests
newman run tests/postman/collection.json

# Run with detailed output
newman run tests/postman/collection.json --reporters cli,json

# Run and generate HTML report
newman run tests/postman/collection.json --reporters html --reporter-html-export report.html
```

### Option 3: Using VS Code REST Client

If you prefer VS Code, you can use the REST Client extension with the examples in the collection.

## 📊 Environment Variables

The collection uses the following variables:

| Variable | Default Value | Description |
|----------|---------------|-------------|
| `nest_base_url` | `http://localhost:3000` | NestJS API base URL |
| `go_base_url` | `http://localhost:8080` | Go API base URL |

To customize these:
1. In Postman, go to **Environments**
2. Create a new environment or edit existing
3. Add these variables with your custom values

## 🧪 Test Structure

Each request follows this pattern:

```javascript
pm.test("Status code is 200", function () {
    pm.response.to.have.status(200);
});

pm.test("Response has required fields", function () {
    var jsonData = pm.response.json();
    pm.expect(jsonData).to.have.property('id');
});
```

## 📝 Example Test Run Output

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
│      prerequest-scripts │                 0 │                 0 │
├─────────────────────────┼───────────────────┼───────────────────┤
│              assertions │                45 │                 0 │
└─────────────────────────┴───────────────────┴───────────────────┘
```

## 🔧 Troubleshooting

### Connection Refused Errors

Make sure both APIs are running:
```bash
# Check if servers are running
curl http://localhost:3000/
curl http://localhost:8080/
```

### Port Already in Use

If ports are occupied, update the port configuration in each API:
- NestJS: Update port in `apps/backend/api-nest/src/main.ts`
- Go: Update port in `apps/backend/api-golang/cmd/api/main.go`

Then update the environment variables in the Postman collection.

### Test Failures

1. Check server logs for errors
2. Verify sample data is seeded correctly
3. Ensure no previous test data is interfering
4. Try restarting both servers

## 📚 Additional Resources

- [Postman Documentation](https://learning.postman.com/)
- [Newman CLI Documentation](https://github.com/postmanlabs/newman)
- [Postman Test Scripts](https://learning.postman.com/docs/writing-scripts/test-scripts/)

## 🎯 CI/CD Integration

To integrate with CI/CD pipelines:

```yaml
# GitHub Actions example
- name: Run API Tests
  run: |
    # Start servers
    moon run api-nest:dev &
    moon run api-golang:dev &
    sleep 5
    # Run tests
    newman run tests/postman/collection.json
```

## 📦 Collection Contents

- **18 Requests** across 2 APIs
- **45+ Test Assertions**
- **100% Endpoint Coverage** for implemented features
- **Feature-First Organization** matching codebase structure
