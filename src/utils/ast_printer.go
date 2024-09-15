package utils

import "github.com/SanteriSuomi/gtiny/src/entities"

type AstPrinter struct{}

func (ap *AstPrinter) VisitBinaryExpression(expression entities.Binary) entities.ExpressionReturn {
	return expression.Accept(ap)
}
