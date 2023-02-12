// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package lib // BooleanExpression

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BooleanExpressionListener is a complete listener for a parse tree produced by BooleanExpressionParser.
type BooleanExpressionListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterBinaryExpression is called when entering the binaryExpression production.
	EnterBinaryExpression(c *BinaryExpressionContext)

	// EnterTypesExpression is called when entering the typesExpression production.
	EnterTypesExpression(c *TypesExpressionContext)

	// EnterInExpression is called when entering the inExpression production.
	EnterInExpression(c *InExpressionContext)

	// EnterToExpression is called when entering the toExpression production.
	EnterToExpression(c *ToExpressionContext)

	// EnterNotExpression is called when entering the notExpression production.
	EnterNotExpression(c *NotExpressionContext)

	// EnterComparatorExpression is called when entering the comparatorExpression production.
	EnterComparatorExpression(c *ComparatorExpressionContext)

	// EnterParentExpression is called when entering the parentExpression production.
	EnterParentExpression(c *ParentExpressionContext)

	// EnterComparator is called when entering the comparator production.
	EnterComparator(c *ComparatorContext)

	// EnterWordlist is called when entering the wordlist production.
	EnterWordlist(c *WordlistContext)

	// EnterNumericTypes is called when entering the numericTypes production.
	EnterNumericTypes(c *NumericTypesContext)

	// EnterTypes is called when entering the types production.
	EnterTypes(c *TypesContext)

	// EnterBinary is called when entering the binary production.
	EnterBinary(c *BinaryContext)

	// EnterBool is called when entering the bool production.
	EnterBool(c *BoolContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitBinaryExpression is called when exiting the binaryExpression production.
	ExitBinaryExpression(c *BinaryExpressionContext)

	// ExitTypesExpression is called when exiting the typesExpression production.
	ExitTypesExpression(c *TypesExpressionContext)

	// ExitInExpression is called when exiting the inExpression production.
	ExitInExpression(c *InExpressionContext)

	// ExitToExpression is called when exiting the toExpression production.
	ExitToExpression(c *ToExpressionContext)

	// ExitNotExpression is called when exiting the notExpression production.
	ExitNotExpression(c *NotExpressionContext)

	// ExitComparatorExpression is called when exiting the comparatorExpression production.
	ExitComparatorExpression(c *ComparatorExpressionContext)

	// ExitParentExpression is called when exiting the parentExpression production.
	ExitParentExpression(c *ParentExpressionContext)

	// ExitComparator is called when exiting the comparator production.
	ExitComparator(c *ComparatorContext)

	// ExitWordlist is called when exiting the wordlist production.
	ExitWordlist(c *WordlistContext)

	// ExitNumericTypes is called when exiting the numericTypes production.
	ExitNumericTypes(c *NumericTypesContext)

	// ExitTypes is called when exiting the types production.
	ExitTypes(c *TypesContext)

	// ExitBinary is called when exiting the binary production.
	ExitBinary(c *BinaryContext)

	// ExitBool is called when exiting the bool production.
	ExitBool(c *BoolContext)
}
