const mergeAlternately = (word1: string, word2: string): string => {
  const max = Math.max(word1.length, word2.length);
  let acc = "";
  for (let i = 0; i < max; i++) {
    acc += word1.charAt(i) + word2.charAt(i);
  }
  return acc;
};

export default mergeAlternately;
