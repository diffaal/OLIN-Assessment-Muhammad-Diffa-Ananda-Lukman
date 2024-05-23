package main

import (
	"fmt"
	"olin-assessment-muhammad-diffa/services"
)

func main() {
	assessmentService := services.NewAssessmentService()
	resultCodeOne := assessmentService.TwoSumArrayTarget([]int{2, 7, 11, 15}, 9)
	resultCodeTwo := assessmentService.ThreeSumArrayZero([]int{-1, 0, 1, 2, -1, -4})

	fmt.Printf("Result Code Question One: %v\n", resultCodeOne)
	fmt.Printf("Result Code Question Two: %v", resultCodeTwo)
}
