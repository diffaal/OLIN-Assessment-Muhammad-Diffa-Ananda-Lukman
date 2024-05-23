package services

import (
	"olin-assessment-muhammad-diffa/package/helper"
	"sort"
)

type AssessmentService struct{}

type IAssessmentService interface {
	TwoSumArrayTarget(nums []int, target int) []int
	ThreeSumArrayZero(nums []int) [][]int
}

func NewAssessmentService() *AssessmentService {
	assessmentService := AssessmentService{}
	return &assessmentService
}

func (as *AssessmentService) TwoSumArrayTarget(nums []int, target int) []int {
	result := []int{}
	numsTable := map[int]int{}

	for i := 0; i < len(nums); i++ {
		numsTable[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		remainder := target - nums[i]
		index, ok := numsTable[remainder]
		if ok && index != i {
			result = append(result, index)
			result = append(result, i)
			break
		}
		numsTable[nums[i]] = i
	}

	return result
}

func (as *AssessmentService) ThreeSumArrayZero(nums []int) [][]int {
	result := [][]int{}
	resultChecker := map[string]int{}
	numsTable := map[int]int{}

	for i := 0; i < len(nums); i++ {
		numsTable[nums[i]] = i
	}

	for a := 0; a < len(nums); a++ {
		for b := a + 1; b < len(nums); b++ {
			c := (nums[a] + nums[b]) * -1

			index, ok := numsTable[c]
			if ok && index != a && index != b {
				triple := []int{c, nums[a], nums[b]}
				sort.Ints(triple)

				mapKey := helper.MapKeyMaker(triple)
				_, ok := resultChecker[mapKey]
				if !ok {
					result = append(result, triple)
					resultChecker[mapKey] = 1
				}
			}
		}
	}
	return result
}
