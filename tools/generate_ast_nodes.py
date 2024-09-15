import io

def define_ast_nodes(types: list[str]):
    builder = io.StringIO()
    builder.write("package entities\n\n")

    define_expression_return(builder)
    define_expression_visitor(builder)
    define_expression(builder)
    define_types(types, builder)

    with open('../src/entities/ast_nodes.go', 'w') as f:
        f.write(builder.getvalue())

def define_expression_return(builder):
    builder.write("type ExpressionReturn struct{}\n\n")

def define_expression_visitor(builder):
    builder.write("type ExpressionVisitor interface {\n")
    builder.write("\tVisitBinaryExpression(expression Binary) ExpressionReturn\n")
    builder.write("}\n\n")

def define_expression(builder):
    builder.write("type Expression interface {\n")
    builder.write("\tAccept(visitor ExpressionVisitor)\n")
    builder.write("}\n\n")

def define_types(types, builder):
    for t in types:
        split = t.split(":")

        name = split[0].strip()
        builder.write(f"type {name} struct {{\n")

        fields = [field.strip() for field in split[1:]]
        builder.write("\t")
        builder.write("\n\t".join(fields))
        builder.write("\n}\n\n")

        # Add Accept method for each struct
        builder.write(f"func (e *{name}) Accept(visitor ExpressionVisitor) ExpressionReturn {{\n")
        builder.write("\treturn visitor.VisitBinaryExpression(*e)\n")
        builder.write("}\n\n")

define_ast_nodes([
    "Binary : left Expression : operator Token : right Expression",
])