/**
 * CalculateX takes an arbitrary input and performs a computation
 * This is a simple example that demonstrates shared utility usage
 */

export interface CalculateXInput {
  value: number;
  multiplier?: number;
  offset?: number;
}

export interface CalculateXOutput {
  result: number;
  formula: string;
  timestamp: Date;
}

/**
 * Performs a calculation based on the input parameters
 * Formula: (value * multiplier) + offset
 * 
 * @param input - The input parameters for calculation
 * @returns The calculated result with metadata
 */
export function calculateX(input: CalculateXInput): CalculateXOutput {
  const multiplier = input.multiplier ?? 2;
  const offset = input.offset ?? 10;
  
  const result = (input.value * multiplier) + offset;
  const formula = `(${input.value} * ${multiplier}) + ${offset} = ${result}`;
  
  return {
    result,
    formula,
    timestamp: new Date(),
  };
}

/**
 * Batch calculation for multiple values
 */
export function calculateXBatch(values: number[], multiplier?: number, offset?: number): CalculateXOutput[] {
  return values.map(value => calculateX({ value, multiplier, offset }));
}
