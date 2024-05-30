// Code generated from /Users/sid/Desktop/filter1/BooleanExpression.g4 by ANTLR 4.13.1. DO NOT EDIT.

package lib // BooleanExpression

import "github.com/antlr4-go/antlr/v4"

// BaseBooleanExpressionListener is a complete listener for a parse tree produced by BooleanExpressionParser.
type BaseBooleanExpressionListener struct{}

var _ BooleanExpressionListener = &BaseBooleanExpressionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBooleanExpressionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBooleanExpressionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBooleanExpressionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBooleanExpressionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseBooleanExpressionListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseBooleanExpressionListener) ExitParse(ctx *ParseContext) {}

// EnterUnaryArithmeticExpression is called when production unaryArithmeticExpression is entered.
func (s *BaseBooleanExpressionListener) EnterUnaryArithmeticExpression(ctx *UnaryArithmeticExpressionContext) {
}

// ExitUnaryArithmeticExpression is called when production unaryArithmeticExpression is exited.
func (s *BaseBooleanExpressionListener) ExitUnaryArithmeticExpression(ctx *UnaryArithmeticExpressionContext) {
}

// EnterBinaryExpression is called when production binaryExpression is entered.
func (s *BaseBooleanExpressionListener) EnterBinaryExpression(ctx *BinaryExpressionContext) {}

// ExitBinaryExpression is called when production binaryExpression is exited.
func (s *BaseBooleanExpressionListener) ExitBinaryExpression(ctx *BinaryExpressionContext) {}

// EnterTypesExpression is called when production typesExpression is entered.
func (s *BaseBooleanExpressionListener) EnterTypesExpression(ctx *TypesExpressionContext) {}

// ExitTypesExpression is called when production typesExpression is exited.
func (s *BaseBooleanExpressionListener) ExitTypesExpression(ctx *TypesExpressionContext) {}

// EnterInExpression is called when production inExpression is entered.
func (s *BaseBooleanExpressionListener) EnterInExpression(ctx *InExpressionContext) {}

// ExitInExpression is called when production inExpression is exited.
func (s *BaseBooleanExpressionListener) ExitInExpression(ctx *InExpressionContext) {}

// EnterArrayExpression is called when production arrayExpression is entered.
func (s *BaseBooleanExpressionListener) EnterArrayExpression(ctx *ArrayExpressionContext) {}

// ExitArrayExpression is called when production arrayExpression is exited.
func (s *BaseBooleanExpressionListener) ExitArrayExpression(ctx *ArrayExpressionContext) {}

// EnterToExpression is called when production toExpression is entered.
func (s *BaseBooleanExpressionListener) EnterToExpression(ctx *ToExpressionContext) {}

// ExitToExpression is called when production toExpression is exited.
func (s *BaseBooleanExpressionListener) ExitToExpression(ctx *ToExpressionContext) {}

// EnterArithmeticFunctionExpression is called when production arithmeticFunctionExpression is entered.
func (s *BaseBooleanExpressionListener) EnterArithmeticFunctionExpression(ctx *ArithmeticFunctionExpressionContext) {
}

// ExitArithmeticFunctionExpression is called when production arithmeticFunctionExpression is exited.
func (s *BaseBooleanExpressionListener) ExitArithmeticFunctionExpression(ctx *ArithmeticFunctionExpressionContext) {
}

// EnterNotExpression is called when production notExpression is entered.
func (s *BaseBooleanExpressionListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production notExpression is exited.
func (s *BaseBooleanExpressionListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterArithmeticExpression is called when production arithmeticExpression is entered.
func (s *BaseBooleanExpressionListener) EnterArithmeticExpression(ctx *ArithmeticExpressionContext) {}

// ExitArithmeticExpression is called when production arithmeticExpression is exited.
func (s *BaseBooleanExpressionListener) ExitArithmeticExpression(ctx *ArithmeticExpressionContext) {}

// EnterComparatorExpression is called when production comparatorExpression is entered.
func (s *BaseBooleanExpressionListener) EnterComparatorExpression(ctx *ComparatorExpressionContext) {}

// ExitComparatorExpression is called when production comparatorExpression is exited.
func (s *BaseBooleanExpressionListener) ExitComparatorExpression(ctx *ComparatorExpressionContext) {}

// EnterParentExpression is called when production parentExpression is entered.
func (s *BaseBooleanExpressionListener) EnterParentExpression(ctx *ParentExpressionContext) {}

// ExitParentExpression is called when production parentExpression is exited.
func (s *BaseBooleanExpressionListener) ExitParentExpression(ctx *ParentExpressionContext) {}

// EnterComparator is called when production comparator is entered.
func (s *BaseBooleanExpressionListener) EnterComparator(ctx *ComparatorContext) {}

// ExitComparator is called when production comparator is exited.
func (s *BaseBooleanExpressionListener) ExitComparator(ctx *ComparatorContext) {}

// EnterArithmeticOperator is called when production arithmeticOperator is entered.
func (s *BaseBooleanExpressionListener) EnterArithmeticOperator(ctx *ArithmeticOperatorContext) {}

// ExitArithmeticOperator is called when production arithmeticOperator is exited.
func (s *BaseBooleanExpressionListener) ExitArithmeticOperator(ctx *ArithmeticOperatorContext) {}

// EnterArithmeticFunction is called when production arithmeticFunction is entered.
func (s *BaseBooleanExpressionListener) EnterArithmeticFunction(ctx *ArithmeticFunctionContext) {}

// ExitArithmeticFunction is called when production arithmeticFunction is exited.
func (s *BaseBooleanExpressionListener) ExitArithmeticFunction(ctx *ArithmeticFunctionContext) {}

// EnterWordlist is called when production wordlist is entered.
func (s *BaseBooleanExpressionListener) EnterWordlist(ctx *WordlistContext) {}

// ExitWordlist is called when production wordlist is exited.
func (s *BaseBooleanExpressionListener) ExitWordlist(ctx *WordlistContext) {}

// EnterArrayOperators is called when production arrayOperators is entered.
func (s *BaseBooleanExpressionListener) EnterArrayOperators(ctx *ArrayOperatorsContext) {}

// ExitArrayOperators is called when production arrayOperators is exited.
func (s *BaseBooleanExpressionListener) ExitArrayOperators(ctx *ArrayOperatorsContext) {}

// EnterNumericTypes is called when production numericTypes is entered.
func (s *BaseBooleanExpressionListener) EnterNumericTypes(ctx *NumericTypesContext) {}

// ExitNumericTypes is called when production numericTypes is exited.
func (s *BaseBooleanExpressionListener) ExitNumericTypes(ctx *NumericTypesContext) {}

// EnterTypes is called when production types is entered.
func (s *BaseBooleanExpressionListener) EnterTypes(ctx *TypesContext) {}

// ExitTypes is called when production types is exited.
func (s *BaseBooleanExpressionListener) ExitTypes(ctx *TypesContext) {}

// EnterBinary is called when production binary is entered.
func (s *BaseBooleanExpressionListener) EnterBinary(ctx *BinaryContext) {}

// ExitBinary is called when production binary is exited.
func (s *BaseBooleanExpressionListener) ExitBinary(ctx *BinaryContext) {}

// EnterBool is called when production bool is entered.
func (s *BaseBooleanExpressionListener) EnterBool(ctx *BoolContext) {}

// ExitBool is called when production bool is exited.
func (s *BaseBooleanExpressionListener) ExitBool(ctx *BoolContext) {}
