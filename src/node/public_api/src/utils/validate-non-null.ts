export function validateNonNull<T>(
  value: T,
  name: string,
): asserts value is NonNullable<T> {
  if (value === null || value === undefined) {
    throw new Error(`Expected non-null ${name} `);
  }
}
