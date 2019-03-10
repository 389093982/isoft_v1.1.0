package block

import (
	"isoft/isoft_iaas_web/models/iwork"
	"sort"
)

type BlockStep struct {
	Step            *iwork.WorkStep // 步骤
	HasChildren     bool            // 是否有子步骤
	ChildBlockSteps []*BlockStep    // 子步骤列表
}

// 将 steps 转换为 BlockStep,最终执行的是 BlockStep
func ParseToBlockStep(steps []iwork.WorkStep) []*BlockStep {
	bSteps := make([]*BlockStep, 0)
	minIndentIndexs := getMinIndentIndex(steps)
	for index, indentIndex := range minIndentIndexs {
		bStep := &BlockStep{
			Step: &steps[indentIndex],
		}
		childs := getChildBlockSteps(index, minIndentIndexs, steps)
		if len(childs) > 0 {
			bStep.HasChildren = true
			bStep.ChildBlockSteps = childs
		}
		bSteps = append(bSteps, bStep)
	}
	return bSteps
}

// index 当前最小缩进索引
// minIndentIndexs 所有最小缩进位置
func getChildBlockSteps(index int, minIndentIndexs []int, steps []iwork.WorkStep) []*BlockStep {
	var max, min int
	min = minIndentIndexs[index]
	if len(minIndentIndexs)-1 == index { // 最后一个最小缩进索引
		max = len(steps)
	} else {
		max = minIndentIndexs[index+1] // 非最后一个最小缩进索引

	}

	bSteps := make([]*BlockStep, 0)

	if max-min > 1 {
		// 获取所有的 childSteps
		childSteps := make([]iwork.WorkStep, 0)
		for i := min + 1; i < max; i++ {
			childSteps = append(childSteps, steps[i])
		}
		// 转换为 BlockStep
		bSteps = append(bSteps, ParseToBlockStep(childSteps)...)
	}
	return bSteps
}

// 获取同批最小缩进值索引
func getMinIndentIndex(steps []iwork.WorkStep) []int {
	indentMap := make(map[int][]int, 0)
	for index, step := range steps {
		if _, ok := indentMap[step.WorkStepIndent]; !ok {
			indentMap[step.WorkStepIndent] = make([]int, 0)
		}
		indentMap[step.WorkStepIndent] = append(indentMap[step.WorkStepIndent], index)
	}
	var indents []int
	for k, _ := range indentMap {
		indents = append(indents, k)
	}
	sort.Ints(indents)
	return indentMap[indents[0]]
}
