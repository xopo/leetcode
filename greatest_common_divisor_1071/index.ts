function gcdOfStrings(word1: string, word2: string): string {
  // words should have same content but different size
  if (word1 + word2 !== word2 + word1) {
    return "";
  }

  const limit = gcd(word1.length, word2.length);
  return word1.substring(0, limit);
}

function gcd(l1: number, l2: number) {
  if (l2 === 0) {
    return l1;
  }
  return gcd(l2, l1 % l2);
}

// bruteforce
// function greatestComonDivisor(word1: string, target: string): string {
//   for (let gcd of getWordComonDivisor(target)) {
//     if (isADivider(word1, gcd)) {
//       return gcd;
//     }
//   }
//   return "";
// }
//
// function* getWordComonDivisor(word: string): Generator<string> {
//   let left = 0;
//   let width = word.length;
//   while (width > 0) {
//     left = 0;
//     width--;
//     const right = left + width;
//     while (left + right <= word.length) {
//       const substr = word.substring(left, left + right + 1);
//       if (isADivider(word, substr)) {
//         yield substr;
//       }
//       left++;
//     }
//     yield "";
//   }
// }
//
// function isADivider(word: string, target: string) {
//   return word.split(target).join("").length === 0;
// }
//
// export default greatestComonDivisor;
