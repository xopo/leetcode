import { assertEquals } from "jsr:@std/assert";
import mergeAlternately from "./index.ts";

Deno.test("first test", () => {
  const result = mergeAlternately("abc", "pqr");
  assertEquals(result, "apbqcr");
});

Deno.test("second test", () => {
  const result = mergeAlternately("ab", "pqrs");
  assertEquals(result, "apbqrs");
});

Deno.test("third test", () => {
  const result = mergeAlternately("abcd", "pq");
  assertEquals(result, "apbqcd");
});
