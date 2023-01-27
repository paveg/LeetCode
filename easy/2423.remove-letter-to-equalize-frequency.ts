/*
 * @lc app=leetcode id=2423 lang=typescript
 *
 * [2423] Remove Letter To Equalize Frequency
 */

// @lc code=start
function equalFrequency(word: string): boolean {
  const cnt: { [K: string]: number } = {}
  for (let w of word) { cnt[w] = (cnt[w] || 0) + 1; }
  for (let k in cnt) {
    const n = Object.assign({}, cnt);
    if (--n[k] === 0) delete n[k];
    if (new Set([...Object.values(n)]).size === 1) return true;
  };

  return false;
};
// @lc code=end

