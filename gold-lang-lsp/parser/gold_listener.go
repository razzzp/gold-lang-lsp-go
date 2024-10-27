// Code generated from antlr4/Gold.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Gold
import "github.com/antlr4-go/antlr/v4"

// GoldListener is a complete listener for a parse tree produced by GoldParser.
type GoldListener interface {
	antlr.ParseTreeListener

	// EnterGoldProg is called when entering the goldProg production.
	EnterGoldProg(c *GoldProgContext)

	// EnterGoldDefinittion is called when entering the goldDefinittion production.
	EnterGoldDefinittion(c *GoldDefinittionContext)

	// EnterGoldProcedureDefinition is called when entering the goldProcedureDefinition production.
	EnterGoldProcedureDefinition(c *GoldProcedureDefinitionContext)

	// EnterGoldFunctionDefinition is called when entering the goldFunctionDefinition production.
	EnterGoldFunctionDefinition(c *GoldFunctionDefinitionContext)

	// EnterParameterDefinition is called when entering the parameterDefinition production.
	EnterParameterDefinition(c *ParameterDefinitionContext)

	// EnterGoldClassDefinition is called when entering the goldClassDefinition production.
	EnterGoldClassDefinition(c *GoldClassDefinitionContext)

	// EnterGoldModuleDefinition is called when entering the goldModuleDefinition production.
	EnterGoldModuleDefinition(c *GoldModuleDefinitionContext)

	// EnterVarType is called when entering the varType production.
	EnterVarType(c *VarTypeContext)

	// ExitGoldProg is called when exiting the goldProg production.
	ExitGoldProg(c *GoldProgContext)

	// ExitGoldDefinittion is called when exiting the goldDefinittion production.
	ExitGoldDefinittion(c *GoldDefinittionContext)

	// ExitGoldProcedureDefinition is called when exiting the goldProcedureDefinition production.
	ExitGoldProcedureDefinition(c *GoldProcedureDefinitionContext)

	// ExitGoldFunctionDefinition is called when exiting the goldFunctionDefinition production.
	ExitGoldFunctionDefinition(c *GoldFunctionDefinitionContext)

	// ExitParameterDefinition is called when exiting the parameterDefinition production.
	ExitParameterDefinition(c *ParameterDefinitionContext)

	// ExitGoldClassDefinition is called when exiting the goldClassDefinition production.
	ExitGoldClassDefinition(c *GoldClassDefinitionContext)

	// ExitGoldModuleDefinition is called when exiting the goldModuleDefinition production.
	ExitGoldModuleDefinition(c *GoldModuleDefinitionContext)

	// ExitVarType is called when exiting the varType production.
	ExitVarType(c *VarTypeContext)
}
