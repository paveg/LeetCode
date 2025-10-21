package code

/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */

// @lc code=start
func minWindow(s string, t string) string {
	// Time complexity: O(m+n)
	// Space complexity: O(1)

	// m: 検索対象文字列sの長さ、n: 必要な文字数（tの長さ）
	m, n := len(s), len(t)
	// sがtより短い場合は、tのすべての文字を含むことができない
	if m < n {
		return ""
	}

	// 各文字の出現回数を記録する配列（ASCII文字用）
	count := [128]int{}
	// tに含まれる各文字の出現回数をカウント
	for i := 0; i < n; i++ {
		count[t[i]]++
	}

	// left: ウィンドウの左端、right: ウィンドウの右端
	left, right := 0, 0
	// minLeft: 最小ウィンドウの開始位置、minSize: 最小ウィンドウのサイズ
	minLeft, minSize := 0, m+1
	for right < m {
		// 現在の文字がtに含まれる文字の場合、必要な文字数を減らす
		if count[s[right]] > 0 {
			n--
		}
		// 現在の文字をウィンドウに追加（カウントを減らす）
		count[s[right]]--
		// 右ポインタを進めながらウィンドウを拡大
		right++

		// tのすべての文字が含まれている場合（n == 0）
		for n == 0 {
			// より小さいウィンドウが見つかった場合、記録を更新
			if right-left < minSize {
				minLeft = left
				minSize = right - left
			}
			// 左端の文字をウィンドウから除外（カウントを増やす）
			count[s[left]]++
			// 除外した文字がtに必要な文字だった場合、必要な文字数を増やす
			if count[s[left]] > 0 {
				n++
			}
			// 左ポインタを進めてウィンドウを縮小
			left++
		}
	}

	// 有効なウィンドウが見つからなかった場合
	if minSize == m+1 {
		return ""
	}
	// 最小ウィンドウの部分文字列を返す
	return s[minLeft : minLeft+minSize]
}

// @lc code=end
