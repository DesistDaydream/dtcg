package subset

// 判断一个数组是否是另一个数组的子集
func IsSubset(subset, superset []string) bool {
	// 超集的 Map。
	supersetMap := make(map[string]int)
	// 遍历超集，统计每个元素的个数
	for _, supersetElement := range superset {
		supersetMap[supersetElement] += 1
	}

	// 遍历子集
	for _, subsetElement := range subset {
		// 将子集中元素作为 超集Map 的键，**逐一**判断子集中每个元素的个数是否与超集中同一元素的个数是否相等
		// 任何一个元素没有找到都返回假
		// 任何一个元素的个数小于1都返回假
		if count, isFound := supersetMap[subsetElement]; !isFound {
			return false
		} else if count < 1 {
			return false
		} else {
			// 若子集中 X 元素的个数大于 1，则每次判断，都要将对应超集 X 元素的数量减 1
			// 这样可以防止，当数组中具有相同元素的时候，会导致的误判。
			// 如果不进行该操作，当子集中 X 元素的个数大于 1 的时候，则会出现如下错误：
			// 命题: [a a] 是 [z d a c b c] 的子集.
			// 结论: true
			// 因为在两个数组中第一次判断 a 之后，长数组中 a 的数量没有减 1，那么下一次判断段数组中第二个 a 时，就会认为也存在在长数组中。
			supersetMap[subsetElement] = count - 1
		}
	}

	return true
}
