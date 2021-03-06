// DO NOT EDIT: This file is autogenerated via the builtin command.

package bigtable

import (
	ast "github.com/influxdata/flux/ast"
	runtime "github.com/influxdata/flux/runtime"
)

func init() {
	runtime.RegisterPackage(pkgAST)
}

var pkgAST = &ast.Package{
	BaseNode: ast.BaseNode{
		Errors: nil,
		Loc:    nil,
	},
	Files: []*ast.File{&ast.File{
		BaseNode: ast.BaseNode{
			Errors: nil,
			Loc: &ast.SourceLocation{
				End: ast.Position{
					Column: 13,
					Line:   3,
				},
				File:   "bigtable.flux",
				Source: "package bigtable\n\nbuiltin from",
				Start: ast.Position{
					Column: 1,
					Line:   1,
				},
			},
		},
		Body: []ast.Statement{&ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 13,
						Line:   3,
					},
					File:   "bigtable.flux",
					Source: "builtin from",
					Start: ast.Position{
						Column: 1,
						Line:   3,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 13,
							Line:   3,
						},
						File:   "bigtable.flux",
						Source: "from",
						Start: ast.Position{
							Column: 9,
							Line:   3,
						},
					},
				},
				Name: "from",
			},
			Ty: ast.TypeExpression{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 104,
							Line:   3,
						},
						File:   "bigtable.flux",
						Source: "(token: string, project: string, instance: string, table: string) => [T] where T: Record",
						Start: ast.Position{
							Column: 16,
							Line:   3,
						},
					},
				},
				Constraints: []*ast.TypeConstraint{&ast.TypeConstraint{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 104,
								Line:   3,
							},
							File:   "bigtable.flux",
							Source: "T: Record",
							Start: ast.Position{
								Column: 95,
								Line:   3,
							},
						},
					},
					Kinds: []*ast.Identifier{&ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 104,
									Line:   3,
								},
								File:   "bigtable.flux",
								Source: "Record",
								Start: ast.Position{
									Column: 98,
									Line:   3,
								},
							},
						},
						Name: "Record",
					}},
					Tvar: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 96,
									Line:   3,
								},
								File:   "bigtable.flux",
								Source: "T",
								Start: ast.Position{
									Column: 95,
									Line:   3,
								},
							},
						},
						Name: "T",
					},
				}},
				Ty: &ast.FunctionType{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 88,
								Line:   3,
							},
							File:   "bigtable.flux",
							Source: "(token: string, project: string, instance: string, table: string) => [T]",
							Start: ast.Position{
								Column: 16,
								Line:   3,
							},
						},
					},
					Parameters: []*ast.ParameterType{&ast.ParameterType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 30,
									Line:   3,
								},
								File:   "bigtable.flux",
								Source: "token: string",
								Start: ast.Position{
									Column: 17,
									Line:   3,
								},
							},
						},
						Kind: "Required",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 22,
										Line:   3,
									},
									File:   "bigtable.flux",
									Source: "token",
									Start: ast.Position{
										Column: 17,
										Line:   3,
									},
								},
							},
							Name: "token",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 30,
										Line:   3,
									},
									File:   "bigtable.flux",
									Source: "string",
									Start: ast.Position{
										Column: 24,
										Line:   3,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 30,
											Line:   3,
										},
										File:   "bigtable.flux",
										Source: "string",
										Start: ast.Position{
											Column: 24,
											Line:   3,
										},
									},
								},
								Name: "string",
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 47,
									Line:   3,
								},
								File:   "bigtable.flux",
								Source: "project: string",
								Start: ast.Position{
									Column: 32,
									Line:   3,
								},
							},
						},
						Kind: "Required",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 39,
										Line:   3,
									},
									File:   "bigtable.flux",
									Source: "project",
									Start: ast.Position{
										Column: 32,
										Line:   3,
									},
								},
							},
							Name: "project",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 47,
										Line:   3,
									},
									File:   "bigtable.flux",
									Source: "string",
									Start: ast.Position{
										Column: 41,
										Line:   3,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 47,
											Line:   3,
										},
										File:   "bigtable.flux",
										Source: "string",
										Start: ast.Position{
											Column: 41,
											Line:   3,
										},
									},
								},
								Name: "string",
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 65,
									Line:   3,
								},
								File:   "bigtable.flux",
								Source: "instance: string",
								Start: ast.Position{
									Column: 49,
									Line:   3,
								},
							},
						},
						Kind: "Required",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 57,
										Line:   3,
									},
									File:   "bigtable.flux",
									Source: "instance",
									Start: ast.Position{
										Column: 49,
										Line:   3,
									},
								},
							},
							Name: "instance",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 65,
										Line:   3,
									},
									File:   "bigtable.flux",
									Source: "string",
									Start: ast.Position{
										Column: 59,
										Line:   3,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 65,
											Line:   3,
										},
										File:   "bigtable.flux",
										Source: "string",
										Start: ast.Position{
											Column: 59,
											Line:   3,
										},
									},
								},
								Name: "string",
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 80,
									Line:   3,
								},
								File:   "bigtable.flux",
								Source: "table: string",
								Start: ast.Position{
									Column: 67,
									Line:   3,
								},
							},
						},
						Kind: "Required",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 72,
										Line:   3,
									},
									File:   "bigtable.flux",
									Source: "table",
									Start: ast.Position{
										Column: 67,
										Line:   3,
									},
								},
							},
							Name: "table",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 80,
										Line:   3,
									},
									File:   "bigtable.flux",
									Source: "string",
									Start: ast.Position{
										Column: 74,
										Line:   3,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 80,
											Line:   3,
										},
										File:   "bigtable.flux",
										Source: "string",
										Start: ast.Position{
											Column: 74,
											Line:   3,
										},
									},
								},
								Name: "string",
							},
						},
					}},
					Return: &ast.ArrayType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 88,
									Line:   3,
								},
								File:   "bigtable.flux",
								Source: "[T]",
								Start: ast.Position{
									Column: 85,
									Line:   3,
								},
							},
						},
						ElementType: &ast.TvarType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 87,
										Line:   3,
									},
									File:   "bigtable.flux",
									Source: "T",
									Start: ast.Position{
										Column: 86,
										Line:   3,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 87,
											Line:   3,
										},
										File:   "bigtable.flux",
										Source: "T",
										Start: ast.Position{
											Column: 86,
											Line:   3,
										},
									},
								},
								Name: "T",
							},
						},
					},
				},
			},
		}},
		Imports:  nil,
		Metadata: "parser-type=rust",
		Name:     "bigtable.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 17,
						Line:   1,
					},
					File:   "bigtable.flux",
					Source: "package bigtable",
					Start: ast.Position{
						Column: 1,
						Line:   1,
					},
				},
			},
			Name: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 17,
							Line:   1,
						},
						File:   "bigtable.flux",
						Source: "bigtable",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "bigtable",
			},
		},
	}},
	Package: "bigtable",
	Path:    "experimental/bigtable",
}
