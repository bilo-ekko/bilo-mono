# @bilo/common

Shared TypeScript utilities for bilo-mono projects.

## Features

### Logger

A simple logging utility with different log levels:

```typescript
import { Logger } from '@bilo/common';

const logger = new Logger('MyApp');
logger.info('Application started');
logger.warn('This is a warning');
logger.error('An error occurred', new Error('Something went wrong'));
```

### CalculateX

A utility function that performs calculations based on input parameters:

```typescript
import { calculateX, calculateXBatch } from '@bilo/common';

// Single calculation
const result = calculateX({ value: 5, multiplier: 3, offset: 10 });
console.log(result.formula); // (5 * 3) + 10 = 25

// Batch calculation
const results = calculateXBatch([1, 2, 3], 2, 5);
```

## Installation

This package is part of the bilo-mono monorepo and can be used internally by other packages.

```json
{
  "dependencies": {
    "@bilo/common": "workspace:*"
  }
}
```

## Development

```bash
# Build the package
npm run build

# Watch mode
npm run dev
```
