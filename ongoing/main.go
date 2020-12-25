package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type logicOperator string

type comparisonOperator string

var (
	andOperator logicOperator = "and"
	orOperator  logicOperator = "or"

	equalOperator comparisonOperator = "=="
	gteOperator   comparisonOperator = ">="
	lteOperator   comparisonOperator = "<="
	gtOperator    comparisonOperator = ">"
	ltOperator    comparisonOperator = "<"
)

type segment struct {
	ID         int              `json:"id"`
	Name       string           `json:"name"`
	Conditions []conditionGroup `json:"condition"`
}

type conditionGroup struct {
	NGs   []nGroup      `json:"ngroup"`
	RGs   []rGroup      `json:"rgroup"`
	Logic logicOperator `json:"logic"`
}

func (cg *conditionGroup) ToExpr() string {
	rgLen := len(cg.RGs)
	ngLen := len(cg.NGs)
	if ngLen == 0 && rgLen == 0 {
		return ""
	}
	var rgExpr, ngExpr string
	for _, rg := range cg.RGs {
		rgExpr += rg.ToExpr()
	}
	for _, ng := range cg.NGs {
		ngExpr += ng.ToExpr()
	}
	return fmt.Sprintf("(%v %v) %v ", rgExpr, ngExpr, cg.Logic)
}

type nGroup struct {
	NGs   []nGroup      `json:"ngroup"`
	RGs   []rGroup      `json:"rgroup"`
	Logic logicOperator `json:"logic"`
}

func (ng *nGroup) ToExpr() string {
	rgLen := len(ng.RGs)
	ngLen := len(ng.NGs)
	if rgLen == 0 && ngLen == 0 {
		return ""
	}
	var ngExpr, rgExpr string
	for _, g := range ng.NGs {
		ngExpr += g.ToExpr()
	}
	for _, g := range ng.RGs {
		rgExpr += g.ToExpr()
	}
	return fmt.Sprintf("(%v %v) %v ", rgExpr, ngExpr, ng.Logic)
}

type rGroup struct {
	Attribute string             `json:"attribute"`
	Operator  comparisonOperator `json:"operator"`
	Value     string             `json:"value"`
	Logic     logicOperator      `json:"logic"`
}

func (rg *rGroup) ToExpr() string {
	expr := fmt.Sprintf("(%v %v \"%v\")", rg.Attribute, rg.Operator, rg.Value)
	if rg.Logic != "" {
		expr = expr + fmt.Sprintf(" %v ", rg.Logic)
	}
	return expr
}

func main() {
	var segment segment
	file, _ := ioutil.ReadFile("schema(no_nested).json")
	if err := json.Unmarshal(file, &segment); err != nil {
		panic(err)
	}
	fmt.Println(segment.Conditions[0].ToExpr(), segment.Conditions[1].ToExpr())
}
