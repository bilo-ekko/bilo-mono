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

Both backend servers must be running:
```bash
# Terminal 1 - Start NestJS API (port 3000)
moon run api-nest:dev

# Terminal 2 - Start Go API (port 8080)
moon run api-golang:dev
```

---

## 🧪 Running Tests

### Method 1: Using the Test Script (Recommended) ⭐

**From the repo root:**
```bash
./tests/postman/test.sh
```

**From the tests/postman folder:**
```bash
cd tests/postman
./test.sh
```

The script will:
- ✅ Check if Newman is installed (installs if needed)
- ✅ Verify both APIs are running
- ✅ Run all tests with nice output
- ✅ Show pass/fail summary

**With additional Newman options:**
```bash
./tests/postman/test.sh --reporters cli,html --reporter-html-export report.html
```

---

### Method 2: Using NPM Scripts

**From the repo root:**
```bash
cd tests/postman
npm install  # First time only
npm test
```

**Available scripts:**
```bash
npm test              # Run tests with default output
npm run test:verbose  # Run with JSON and HTML reports
npm run test:html     # Run and auto-open HTML report
```

---

### Method 3: Using Moon Tasks

**From anywhere in the repo:**
```bash
# Run basic tests
moon run postman:test

# Run with verbose output and reports
moon run postman:test-verbose

# Run and generate HTML report
moon run postman:test-html

# Use the shell script
moon run postman:test-script
```

---

### Method 4: Using Newman Directly

**Install Newman globally (one time):**
```bash
npm install -g newman
```

**From repo root:**
```bash
newman run tests/postman/collection.json
```

**From tests/postman folder:**
```bash
cd tests/postman
newman run collection.json
```

**With options:**
```bash
newman run tests/postman/collection.json \
  --reporters cli,html \
  --reporter-html-export report.html
```

---

### Method 5: Using Postman Desktop App

1. Open Postman
2. Click **Import** in the top-left
3. Select the `collection.json` file
4. The collection will appear in your sidebar
5. Click the **Run** button to execute all tests

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

## 📝 Quick Reference

### From Repo Root
```bash
./tests/postman/test.sh                    # Easiest way
cd tests/postman && npm test               # NPM
moon run postman:test                      # Moon
newman run tests/postman/collection.json   # Direct Newman
```

### From tests/postman Folder
```bash
./test.sh                                  # Shell script
npm test                                   # NPM
newman run collection.json                 # Newman
```

---

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
