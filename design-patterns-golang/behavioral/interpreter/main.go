package interpreter

import "fmt"

func TestInterpreterDemo() {
	// male、women表达式
	maleExpression := TerminalExpression{data: "male"}
	womenExpression := TerminalExpression{data: "women"}
	// people is male or women
	peopleExpression := OrExpression{
		exp1: maleExpression,
		exp2: womenExpression,
	}
	// 调用示例
	context := "Jon is male."
	fmt.Printf("Jon is people? %v \n", peopleExpression.interpret(context))

	marriedExpression := TerminalExpression{data: "married"}
	// married women expression
	marriedWomenExpression := AndExpression{
		exp1: marriedExpression,
		exp2: womenExpression,
	}
	// 调用示例
	context1 := "Mary(women) is married."
	fmt.Printf("Mary is married women? %v \n", marriedWomenExpression.interpret(context1))

}
