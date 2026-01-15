# Postman API Tests

Comprehensive API testing collection for NestJS and Go backend services.

## 🚀 Quick Start

### 1. Start Backend Servers
```bash
moon run api-nest:dev    # Port 3000
moon run api-golang:dev  # Port 8080
```

### 2. Run Tests (Choose One)

**Easiest (from repo root):**
```bash
./tests/postman/test.sh
```

**Using NPM:**
```bash
cd tests/postman
npm install  # First time only
npm test
```

**Using Moon:**
```bash
moon run postman:test
```

**Using Newman:**
```bash
newman run tests/postman/collection.json
```

## 📊 What Gets Tested

- ✅ 18 API endpoints
- ✅ 45+ test assertions
- ✅ Both NestJS (port 3000) and Go (port 8080) APIs
- ✅ All CRUD operations
- ✅ Query filtering
- ✅ Data validation

## 📚 Full Documentation

See [README.md](./README.md) for complete documentation.
