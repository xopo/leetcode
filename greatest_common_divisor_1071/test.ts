import { assertEquals } from "jsr:@std/assert";
import greatestComonDivisor from "./index.ts";

const tests = [
  { i: ["ABCABC", "ABC"], o: "ABC" },
  { i: ["ABABAB", "ABAB"], o: "AB" },
  { i: ["LEET", "CODE"], o: "" },
  {
    i: [
      "TAUXXTAUXXTAUXXTAUXXTAUXX",
      "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX",
    ],
    o: "TAUXX",
  },
  {
    i: [
      "CXTXNCXTXNCXTXNCXTXNCXTXN",
      "CXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXN",
    ],
    o: "CXTXN",
  },
];

Deno.test("first test", () => {
  for (let t of tests) {
    const result = greatestComonDivisor(t.i[0], t.i[1]);
    assertEquals(result, t.o);
  }
});
