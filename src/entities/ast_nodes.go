package entities

type ExpressionReturn struct{}

type ExpressionVisitor interface {
	VisitBinaryExpression(expression Binary) ExpressionReturn
}

type Expression interface {
	Accept(visitor ExpressionVisitor)
}

type Binary struct {
	left Expression
	operator Token
	right Expression
}

func (e *Binary) Accept(visitor ExpressionVisitor) ExpressionReturn {
	return visitor.VisitBinaryExpression(*e)
}

