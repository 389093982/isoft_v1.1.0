package block

import (
	"isoft/isoft_iaas_web/models/iwork"
	"sort"
)

type BlockStep struct {
	Step            *iwork.WorkStep // 步骤
	HasChildren     bool            // 是否有子步骤
	ChildBlockSteps []*BlockStep    // 子步骤列表
	ParentBlockStep *BlockStep      // 父级 BlockStep
}

// 将 steps 转换为 BlockStep,最终执行的是 BlockStep
func ParseToBlockStep(steps []iwork.WorkStep) []*BlockStep {
	_, blockSteps := ParseAndGetCurrentBlockStep(nil, steps)
	return blockSteps
}

func ParseAndGetCurrentBlockStep(currentStep *iwork.WorkStep, steps []iwork.WorkStep) (currentBlockStep *BlockStep, blockSteps []*BlockStep) {
	blockSteps = make([]*BlockStep, 0)
	minIndentIndexs := getMinIndentIndex(steps)
	for index, indentIndex := range minIndentIndexs {
		bStep := &BlockStep{
			Step: &steps[indentIndex],
		}
		if currentStep != nil && steps[indentIndex].WorkStepId == currentStep.WorkStepId {
			currentBlockStep = bStep
		}
		_currentBlockStep, childs := getChildBlockSteps(currentStep, index, minIndentIndexs, steps)
		if _currentBlockStep != nil {
			currentBlockStep = _currentBlockStep
		}
		if len(childs) > 0 {
			bStep.HasChildren = true
			bStep.ChildBlockSteps = childs
			for _, child := range childs {
				// 设置 parent 属性
				child.ParentBlockStep = bStep
			}
		}
		blockSteps = append(blockSteps, bStep)
	}
	return
}

// index 当前最小缩进索引
// minIndentIndexs 所有最小缩进位置
func getChildBlockSteps(currentStep *iwork.WorkStep, index int, minIndentIndexs []int,
	steps []iwork.WorkStep) (currentBlockStep *BlockStep, blockSteps []*BlockStep) {
	var max, min int
	min = minIndentIndexs[index]
	if len(minIndentIndexs)-1 == index { // 最后一个最小缩进索引
		max = len(steps)
	} else {
		max = minIndentIndexs[index+1] // 非最后一个最小缩进索引

	}

	blockSteps = make([]*BlockStep, 0)
	if max-min > 1 {
		// 获取所有的 childSteps
		childSteps := make([]iwork.WorkStep, 0)
		for i := min + 1; i < max; i++ {
			childSteps = append(childSteps, steps[i])
		}
		// 转换为 BlockStep
		_currentBlockStep, _blockSteps := ParseAndGetCurrentBlockStep(currentStep, childSteps)
		if _currentBlockStep != nil {
			currentBlockStep = _currentBlockStep
		}
		blockSteps = append(blockSteps, _blockSteps...)
	}
	return
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

// 判断前置 step 在块范围内是否是可访问的
func CheckBlockAccessble(allBlockSteps []*BlockStep, currentBlockStep *BlockStep, checkStepId int64) bool {
	for {
		// 获取父级别 blockStep
		parentBlockStep := currentBlockStep.ParentBlockStep
		if parentBlockStep == nil { // 最外层 block
			for _, blockStep := range allBlockSteps {
				if blockStep.Step.WorkStepId == checkStepId {
					return true
				}
			}
			return false
		}
		for _, cBlockStep := range parentBlockStep.ChildBlockSteps {
			if cBlockStep.Step.WorkStepId == checkStepId {
				return true
			}
		}
		currentBlockStep = parentBlockStep
	}
}
