// Code generated from Delphi.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Delphi

import "github.com/antlr4-go/antlr/v4"

// DelphiListener is a complete listener for a parse tree produced by DelphiParser.
type DelphiListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterProgramHead is called when entering the programHead production.
	EnterProgramHead(c *ProgramHeadContext)

	// EnterProgramParmSeq is called when entering the programParmSeq production.
	EnterProgramParmSeq(c *ProgramParmSeqContext)

	// EnterLibrary is called when entering the library production.
	EnterLibrary(c *LibraryContext)

	// EnterLibraryHead is called when entering the libraryHead production.
	EnterLibraryHead(c *LibraryHeadContext)

	// EnterPackageE is called when entering the packageE production.
	EnterPackageE(c *PackageEContext)

	// EnterPackageHead is called when entering the packageHead production.
	EnterPackageHead(c *PackageHeadContext)

	// EnterUnit is called when entering the unit production.
	EnterUnit(c *UnitContext)

	// EnterUnitHead is called when entering the unitHead production.
	EnterUnitHead(c *UnitHeadContext)

	// EnterUnitInterface is called when entering the unitInterface production.
	EnterUnitInterface(c *UnitInterfaceContext)

	// EnterUnitImplementation is called when entering the unitImplementation production.
	EnterUnitImplementation(c *UnitImplementationContext)

	// EnterUnitBlock is called when entering the unitBlock production.
	EnterUnitBlock(c *UnitBlockContext)

	// EnterUnitInitialization is called when entering the unitInitialization production.
	EnterUnitInitialization(c *UnitInitializationContext)

	// EnterUnitFinalization is called when entering the unitFinalization production.
	EnterUnitFinalization(c *UnitFinalizationContext)

	// EnterContainsClause is called when entering the containsClause production.
	EnterContainsClause(c *ContainsClauseContext)

	// EnterRequiresClause is called when entering the requiresClause production.
	EnterRequiresClause(c *RequiresClauseContext)

	// EnterUsesClause is called when entering the usesClause production.
	EnterUsesClause(c *UsesClauseContext)

	// EnterUsesFileClause is called when entering the usesFileClause production.
	EnterUsesFileClause(c *UsesFileClauseContext)

	// EnterNamespaceFileNameList is called when entering the namespaceFileNameList production.
	EnterNamespaceFileNameList(c *NamespaceFileNameListContext)

	// EnterNamespaceFileName is called when entering the namespaceFileName production.
	EnterNamespaceFileName(c *NamespaceFileNameContext)

	// EnterNamespaceNameList is called when entering the namespaceNameList production.
	EnterNamespaceNameList(c *NamespaceNameListContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBlockBody is called when entering the blockBody production.
	EnterBlockBody(c *BlockBodyContext)

	// EnterDeclSection is called when entering the declSection production.
	EnterDeclSection(c *DeclSectionContext)

	// EnterInterfaceDecl is called when entering the interfaceDecl production.
	EnterInterfaceDecl(c *InterfaceDeclContext)

	// EnterLabelDeclSection is called when entering the labelDeclSection production.
	EnterLabelDeclSection(c *LabelDeclSectionContext)

	// EnterConstSection is called when entering the constSection production.
	EnterConstSection(c *ConstSectionContext)

	// EnterConstKey is called when entering the constKey production.
	EnterConstKey(c *ConstKeyContext)

	// EnterConstDeclaration is called when entering the constDeclaration production.
	EnterConstDeclaration(c *ConstDeclarationContext)

	// EnterTypeSection is called when entering the typeSection production.
	EnterTypeSection(c *TypeSectionContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterVarSection is called when entering the varSection production.
	EnterVarSection(c *VarSectionContext)

	// EnterVarKey is called when entering the varKey production.
	EnterVarKey(c *VarKeyContext)

	// EnterVarDeclaration is called when entering the varDeclaration production.
	EnterVarDeclaration(c *VarDeclarationContext)

	// EnterVarValueSpec is called when entering the varValueSpec production.
	EnterVarValueSpec(c *VarValueSpecContext)

	// EnterExportsSection is called when entering the exportsSection production.
	EnterExportsSection(c *ExportsSectionContext)

	// EnterExportItem is called when entering the exportItem production.
	EnterExportItem(c *ExportItemContext)

	// EnterTypeDecl is called when entering the typeDecl production.
	EnterTypeDecl(c *TypeDeclContext)

	// EnterStrucType is called when entering the strucType production.
	EnterStrucType(c *StrucTypeContext)

	// EnterStrucTypePart is called when entering the strucTypePart production.
	EnterStrucTypePart(c *StrucTypePartContext)

	// EnterArrayType is called when entering the arrayType production.
	EnterArrayType(c *ArrayTypeContext)

	// EnterArrayIndex is called when entering the arrayIndex production.
	EnterArrayIndex(c *ArrayIndexContext)

	// EnterArraySubType is called when entering the arraySubType production.
	EnterArraySubType(c *ArraySubTypeContext)

	// EnterSetType is called when entering the setType production.
	EnterSetType(c *SetTypeContext)

	// EnterFileType is called when entering the fileType production.
	EnterFileType(c *FileTypeContext)

	// EnterPointerType is called when entering the pointerType production.
	EnterPointerType(c *PointerTypeContext)

	// EnterStringType is called when entering the stringType production.
	EnterStringType(c *StringTypeContext)

	// EnterCodePageNumber is called when entering the codePageNumber production.
	EnterCodePageNumber(c *CodePageNumberContext)

	// EnterProcedureType is called when entering the procedureType production.
	EnterProcedureType(c *ProcedureTypeContext)

	// EnterMethodType is called when entering the methodType production.
	EnterMethodType(c *MethodTypeContext)

	// EnterSimpleProcedureType is called when entering the simpleProcedureType production.
	EnterSimpleProcedureType(c *SimpleProcedureTypeContext)

	// EnterProcedureReference is called when entering the procedureReference production.
	EnterProcedureReference(c *ProcedureReferenceContext)

	// EnterProcedureTypeHeading is called when entering the procedureTypeHeading production.
	EnterProcedureTypeHeading(c *ProcedureTypeHeadingContext)

	// EnterVariantType is called when entering the variantType production.
	EnterVariantType(c *VariantTypeContext)

	// EnterSimpleType is called when entering the simpleType production.
	EnterSimpleType(c *SimpleTypeContext)

	// EnterSubRangeType is called when entering the subRangeType production.
	EnterSubRangeType(c *SubRangeTypeContext)

	// EnterEnumType is called when entering the enumType production.
	EnterEnumType(c *EnumTypeContext)

	// EnterTypeId is called when entering the typeId production.
	EnterTypeId(c *TypeIdContext)

	// EnterGenericTypeIdent is called when entering the genericTypeIdent production.
	EnterGenericTypeIdent(c *GenericTypeIdentContext)

	// EnterGenericDefinition is called when entering the genericDefinition production.
	EnterGenericDefinition(c *GenericDefinitionContext)

	// EnterSimpleGenericDefinition is called when entering the simpleGenericDefinition production.
	EnterSimpleGenericDefinition(c *SimpleGenericDefinitionContext)

	// EnterConstrainedGenericDefinition is called when entering the constrainedGenericDefinition production.
	EnterConstrainedGenericDefinition(c *ConstrainedGenericDefinitionContext)

	// EnterConstrainedGeneric is called when entering the constrainedGeneric production.
	EnterConstrainedGeneric(c *ConstrainedGenericContext)

	// EnterGenericConstraint is called when entering the genericConstraint production.
	EnterGenericConstraint(c *GenericConstraintContext)

	// EnterGenericPostfix is called when entering the genericPostfix production.
	EnterGenericPostfix(c *GenericPostfixContext)

	// EnterClassDecl is called when entering the classDecl production.
	EnterClassDecl(c *ClassDeclContext)

	// EnterClassTypeTypeDecl is called when entering the classTypeTypeDecl production.
	EnterClassTypeTypeDecl(c *ClassTypeTypeDeclContext)

	// EnterClassTypeDecl is called when entering the classTypeDecl production.
	EnterClassTypeDecl(c *ClassTypeDeclContext)

	// EnterClassState is called when entering the classState production.
	EnterClassState(c *ClassStateContext)

	// EnterClassParent is called when entering the classParent production.
	EnterClassParent(c *ClassParentContext)

	// EnterClassItem is called when entering the classItem production.
	EnterClassItem(c *ClassItemContext)

	// EnterClassHelperDecl is called when entering the classHelperDecl production.
	EnterClassHelperDecl(c *ClassHelperDeclContext)

	// EnterClassHelperItem is called when entering the classHelperItem production.
	EnterClassHelperItem(c *ClassHelperItemContext)

	// EnterInterfaceTypeDecl is called when entering the interfaceTypeDecl production.
	EnterInterfaceTypeDecl(c *InterfaceTypeDeclContext)

	// EnterInterfaceKey is called when entering the interfaceKey production.
	EnterInterfaceKey(c *InterfaceKeyContext)

	// EnterInterfaceGuid is called when entering the interfaceGuid production.
	EnterInterfaceGuid(c *InterfaceGuidContext)

	// EnterInterfaceItem is called when entering the interfaceItem production.
	EnterInterfaceItem(c *InterfaceItemContext)

	// EnterObjectDecl is called when entering the objectDecl production.
	EnterObjectDecl(c *ObjectDeclContext)

	// EnterObjectItem is called when entering the objectItem production.
	EnterObjectItem(c *ObjectItemContext)

	// EnterRecordDecl is called when entering the recordDecl production.
	EnterRecordDecl(c *RecordDeclContext)

	// EnterSimpleRecord is called when entering the simpleRecord production.
	EnterSimpleRecord(c *SimpleRecordContext)

	// EnterVariantRecord is called when entering the variantRecord production.
	EnterVariantRecord(c *VariantRecordContext)

	// EnterRecordItem is called when entering the recordItem production.
	EnterRecordItem(c *RecordItemContext)

	// EnterRecordField is called when entering the recordField production.
	EnterRecordField(c *RecordFieldContext)

	// EnterRecordVariantField is called when entering the recordVariantField production.
	EnterRecordVariantField(c *RecordVariantFieldContext)

	// EnterRecordVariantSection is called when entering the recordVariantSection production.
	EnterRecordVariantSection(c *RecordVariantSectionContext)

	// EnterRecordVariant is called when entering the recordVariant production.
	EnterRecordVariant(c *RecordVariantContext)

	// EnterRecordHelperDecl is called when entering the recordHelperDecl production.
	EnterRecordHelperDecl(c *RecordHelperDeclContext)

	// EnterRecordHelperItem is called when entering the recordHelperItem production.
	EnterRecordHelperItem(c *RecordHelperItemContext)

	// EnterClassMethod is called when entering the classMethod production.
	EnterClassMethod(c *ClassMethodContext)

	// EnterClassField is called when entering the classField production.
	EnterClassField(c *ClassFieldContext)

	// EnterClassProperty is called when entering the classProperty production.
	EnterClassProperty(c *ClassPropertyContext)

	// EnterClassPropertyArray is called when entering the classPropertyArray production.
	EnterClassPropertyArray(c *ClassPropertyArrayContext)

	// EnterClassPropertyIndex is called when entering the classPropertyIndex production.
	EnterClassPropertyIndex(c *ClassPropertyIndexContext)

	// EnterClassPropertySpecifier is called when entering the classPropertySpecifier production.
	EnterClassPropertySpecifier(c *ClassPropertySpecifierContext)

	// EnterClassPropertyEndSpecifier is called when entering the classPropertyEndSpecifier production.
	EnterClassPropertyEndSpecifier(c *ClassPropertyEndSpecifierContext)

	// EnterClassPropertyReadWrite is called when entering the classPropertyReadWrite production.
	EnterClassPropertyReadWrite(c *ClassPropertyReadWriteContext)

	// EnterClassPropertyDispInterface is called when entering the classPropertyDispInterface production.
	EnterClassPropertyDispInterface(c *ClassPropertyDispInterfaceContext)

	// EnterVisibility is called when entering the visibility production.
	EnterVisibility(c *VisibilityContext)

	// EnterExportedProcHeading is called when entering the exportedProcHeading production.
	EnterExportedProcHeading(c *ExportedProcHeadingContext)

	// EnterMethodDecl is called when entering the methodDecl production.
	EnterMethodDecl(c *MethodDeclContext)

	// EnterMethodDeclHeading is called when entering the methodDeclHeading production.
	EnterMethodDeclHeading(c *MethodDeclHeadingContext)

	// EnterMethodKey is called when entering the methodKey production.
	EnterMethodKey(c *MethodKeyContext)

	// EnterMethodName is called when entering the methodName production.
	EnterMethodName(c *MethodNameContext)

	// EnterProcDecl is called when entering the procDecl production.
	EnterProcDecl(c *ProcDeclContext)

	// EnterProcDeclHeading is called when entering the procDeclHeading production.
	EnterProcDeclHeading(c *ProcDeclHeadingContext)

	// EnterFormalParameterSection is called when entering the formalParameterSection production.
	EnterFormalParameterSection(c *FormalParameterSectionContext)

	// EnterFormalParameterList is called when entering the formalParameterList production.
	EnterFormalParameterList(c *FormalParameterListContext)

	// EnterFormalParameter is called when entering the formalParameter production.
	EnterFormalParameter(c *FormalParameterContext)

	// EnterParmType is called when entering the parmType production.
	EnterParmType(c *ParmTypeContext)

	// EnterMethodBody is called when entering the methodBody production.
	EnterMethodBody(c *MethodBodyContext)

	// EnterProcBody is called when entering the procBody production.
	EnterProcBody(c *ProcBodyContext)

	// EnterCustomAttribute is called when entering the customAttribute production.
	EnterCustomAttribute(c *CustomAttributeContext)

	// EnterCustomAttributeList is called when entering the customAttributeList production.
	EnterCustomAttributeList(c *CustomAttributeListContext)

	// EnterCustomAttributeDecl is called when entering the customAttributeDecl production.
	EnterCustomAttributeDecl(c *CustomAttributeDeclContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAnonymousExpression is called when entering the anonymousExpression production.
	EnterAnonymousExpression(c *AnonymousExpressionContext)

	// EnterSimpleExpression is called when entering the simpleExpression production.
	EnterSimpleExpression(c *SimpleExpressionContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterStringFactor is called when entering the stringFactor production.
	EnterStringFactor(c *StringFactorContext)

	// EnterSetSection is called when entering the setSection production.
	EnterSetSection(c *SetSectionContext)

	// EnterDesignator is called when entering the designator production.
	EnterDesignator(c *DesignatorContext)

	// EnterDesignatorItem is called when entering the designatorItem production.
	EnterDesignatorItem(c *DesignatorItemContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterColonConstruct is called when entering the colonConstruct production.
	EnterColonConstruct(c *ColonConstructContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterRelOp is called when entering the relOp production.
	EnterRelOp(c *RelOpContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterCaseStatement is called when entering the caseStatement production.
	EnterCaseStatement(c *CaseStatementContext)

	// EnterCaseItem is called when entering the caseItem production.
	EnterCaseItem(c *CaseItemContext)

	// EnterCaseLabel is called when entering the caseLabel production.
	EnterCaseLabel(c *CaseLabelContext)

	// EnterRepeatStatement is called when entering the repeatStatement production.
	EnterRepeatStatement(c *RepeatStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterWithStatement is called when entering the withStatement production.
	EnterWithStatement(c *WithStatementContext)

	// EnterWithItem is called when entering the withItem production.
	EnterWithItem(c *WithItemContext)

	// EnterCompoundStatement is called when entering the compoundStatement production.
	EnterCompoundStatement(c *CompoundStatementContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterSimpleStatement is called when entering the simpleStatement production.
	EnterSimpleStatement(c *SimpleStatementContext)

	// EnterGotoStatement is called when entering the gotoStatement production.
	EnterGotoStatement(c *GotoStatementContext)

	// EnterConstExpression is called when entering the constExpression production.
	EnterConstExpression(c *ConstExpressionContext)

	// EnterRecordConstExpression is called when entering the recordConstExpression production.
	EnterRecordConstExpression(c *RecordConstExpressionContext)

	// EnterTryStatement is called when entering the tryStatement production.
	EnterTryStatement(c *TryStatementContext)

	// EnterHandlerList is called when entering the handlerList production.
	EnterHandlerList(c *HandlerListContext)

	// EnterHandler is called when entering the handler production.
	EnterHandler(c *HandlerContext)

	// EnterHandlerIdent is called when entering the handlerIdent production.
	EnterHandlerIdent(c *HandlerIdentContext)

	// EnterHandlerStatement is called when entering the handlerStatement production.
	EnterHandlerStatement(c *HandlerStatementContext)

	// EnterRaiseStatement is called when entering the raiseStatement production.
	EnterRaiseStatement(c *RaiseStatementContext)

	// EnterAssemblerStatement is called when entering the assemblerStatement production.
	EnterAssemblerStatement(c *AssemblerStatementContext)

	// EnterMethodDirective is called when entering the methodDirective production.
	EnterMethodDirective(c *MethodDirectiveContext)

	// EnterFunctionDirective is called when entering the functionDirective production.
	EnterFunctionDirective(c *FunctionDirectiveContext)

	// EnterReintroduceDirective is called when entering the reintroduceDirective production.
	EnterReintroduceDirective(c *ReintroduceDirectiveContext)

	// EnterOverloadDirective is called when entering the overloadDirective production.
	EnterOverloadDirective(c *OverloadDirectiveContext)

	// EnterBindingDirective is called when entering the bindingDirective production.
	EnterBindingDirective(c *BindingDirectiveContext)

	// EnterAbstractDirective is called when entering the abstractDirective production.
	EnterAbstractDirective(c *AbstractDirectiveContext)

	// EnterInlineDirective is called when entering the inlineDirective production.
	EnterInlineDirective(c *InlineDirectiveContext)

	// EnterCallConvention is called when entering the callConvention production.
	EnterCallConvention(c *CallConventionContext)

	// EnterCallConventionNoSemi is called when entering the callConventionNoSemi production.
	EnterCallConventionNoSemi(c *CallConventionNoSemiContext)

	// EnterOldCallConventionDirective is called when entering the oldCallConventionDirective production.
	EnterOldCallConventionDirective(c *OldCallConventionDirectiveContext)

	// EnterHintingDirective is called when entering the hintingDirective production.
	EnterHintingDirective(c *HintingDirectiveContext)

	// EnterExternalDirective is called when entering the externalDirective production.
	EnterExternalDirective(c *ExternalDirectiveContext)

	// EnterExternalSpecifier is called when entering the externalSpecifier production.
	EnterExternalSpecifier(c *ExternalSpecifierContext)

	// EnterDispIDDirective is called when entering the dispIDDirective production.
	EnterDispIDDirective(c *DispIDDirectiveContext)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterUsedKeywordsAsNames is called when entering the usedKeywordsAsNames production.
	EnterUsedKeywordsAsNames(c *UsedKeywordsAsNamesContext)

	// EnterIdentList is called when entering the identList production.
	EnterIdentList(c *IdentListContext)

	// EnterIdentListFlat is called when entering the identListFlat production.
	EnterIdentListFlat(c *IdentListFlatContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterIntNum is called when entering the intNum production.
	EnterIntNum(c *IntNumContext)

	// EnterRealNum is called when entering the realNum production.
	EnterRealNum(c *RealNumContext)

	// EnterNamespacedQualifiedIdent is called when entering the namespacedQualifiedIdent production.
	EnterNamespacedQualifiedIdent(c *NamespacedQualifiedIdentContext)

	// EnterNamespaceName is called when entering the namespaceName production.
	EnterNamespaceName(c *NamespaceNameContext)

	// EnterQualifiedIdent is called when entering the qualifiedIdent production.
	EnterQualifiedIdent(c *QualifiedIdentContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitProgramHead is called when exiting the programHead production.
	ExitProgramHead(c *ProgramHeadContext)

	// ExitProgramParmSeq is called when exiting the programParmSeq production.
	ExitProgramParmSeq(c *ProgramParmSeqContext)

	// ExitLibrary is called when exiting the library production.
	ExitLibrary(c *LibraryContext)

	// ExitLibraryHead is called when exiting the libraryHead production.
	ExitLibraryHead(c *LibraryHeadContext)

	// ExitPackageE is called when exiting the packageE production.
	ExitPackageE(c *PackageEContext)

	// ExitPackageHead is called when exiting the packageHead production.
	ExitPackageHead(c *PackageHeadContext)

	// ExitUnit is called when exiting the unit production.
	ExitUnit(c *UnitContext)

	// ExitUnitHead is called when exiting the unitHead production.
	ExitUnitHead(c *UnitHeadContext)

	// ExitUnitInterface is called when exiting the unitInterface production.
	ExitUnitInterface(c *UnitInterfaceContext)

	// ExitUnitImplementation is called when exiting the unitImplementation production.
	ExitUnitImplementation(c *UnitImplementationContext)

	// ExitUnitBlock is called when exiting the unitBlock production.
	ExitUnitBlock(c *UnitBlockContext)

	// ExitUnitInitialization is called when exiting the unitInitialization production.
	ExitUnitInitialization(c *UnitInitializationContext)

	// ExitUnitFinalization is called when exiting the unitFinalization production.
	ExitUnitFinalization(c *UnitFinalizationContext)

	// ExitContainsClause is called when exiting the containsClause production.
	ExitContainsClause(c *ContainsClauseContext)

	// ExitRequiresClause is called when exiting the requiresClause production.
	ExitRequiresClause(c *RequiresClauseContext)

	// ExitUsesClause is called when exiting the usesClause production.
	ExitUsesClause(c *UsesClauseContext)

	// ExitUsesFileClause is called when exiting the usesFileClause production.
	ExitUsesFileClause(c *UsesFileClauseContext)

	// ExitNamespaceFileNameList is called when exiting the namespaceFileNameList production.
	ExitNamespaceFileNameList(c *NamespaceFileNameListContext)

	// ExitNamespaceFileName is called when exiting the namespaceFileName production.
	ExitNamespaceFileName(c *NamespaceFileNameContext)

	// ExitNamespaceNameList is called when exiting the namespaceNameList production.
	ExitNamespaceNameList(c *NamespaceNameListContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBlockBody is called when exiting the blockBody production.
	ExitBlockBody(c *BlockBodyContext)

	// ExitDeclSection is called when exiting the declSection production.
	ExitDeclSection(c *DeclSectionContext)

	// ExitInterfaceDecl is called when exiting the interfaceDecl production.
	ExitInterfaceDecl(c *InterfaceDeclContext)

	// ExitLabelDeclSection is called when exiting the labelDeclSection production.
	ExitLabelDeclSection(c *LabelDeclSectionContext)

	// ExitConstSection is called when exiting the constSection production.
	ExitConstSection(c *ConstSectionContext)

	// ExitConstKey is called when exiting the constKey production.
	ExitConstKey(c *ConstKeyContext)

	// ExitConstDeclaration is called when exiting the constDeclaration production.
	ExitConstDeclaration(c *ConstDeclarationContext)

	// ExitTypeSection is called when exiting the typeSection production.
	ExitTypeSection(c *TypeSectionContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitVarSection is called when exiting the varSection production.
	ExitVarSection(c *VarSectionContext)

	// ExitVarKey is called when exiting the varKey production.
	ExitVarKey(c *VarKeyContext)

	// ExitVarDeclaration is called when exiting the varDeclaration production.
	ExitVarDeclaration(c *VarDeclarationContext)

	// ExitVarValueSpec is called when exiting the varValueSpec production.
	ExitVarValueSpec(c *VarValueSpecContext)

	// ExitExportsSection is called when exiting the exportsSection production.
	ExitExportsSection(c *ExportsSectionContext)

	// ExitExportItem is called when exiting the exportItem production.
	ExitExportItem(c *ExportItemContext)

	// ExitTypeDecl is called when exiting the typeDecl production.
	ExitTypeDecl(c *TypeDeclContext)

	// ExitStrucType is called when exiting the strucType production.
	ExitStrucType(c *StrucTypeContext)

	// ExitStrucTypePart is called when exiting the strucTypePart production.
	ExitStrucTypePart(c *StrucTypePartContext)

	// ExitArrayType is called when exiting the arrayType production.
	ExitArrayType(c *ArrayTypeContext)

	// ExitArrayIndex is called when exiting the arrayIndex production.
	ExitArrayIndex(c *ArrayIndexContext)

	// ExitArraySubType is called when exiting the arraySubType production.
	ExitArraySubType(c *ArraySubTypeContext)

	// ExitSetType is called when exiting the setType production.
	ExitSetType(c *SetTypeContext)

	// ExitFileType is called when exiting the fileType production.
	ExitFileType(c *FileTypeContext)

	// ExitPointerType is called when exiting the pointerType production.
	ExitPointerType(c *PointerTypeContext)

	// ExitStringType is called when exiting the stringType production.
	ExitStringType(c *StringTypeContext)

	// ExitCodePageNumber is called when exiting the codePageNumber production.
	ExitCodePageNumber(c *CodePageNumberContext)

	// ExitProcedureType is called when exiting the procedureType production.
	ExitProcedureType(c *ProcedureTypeContext)

	// ExitMethodType is called when exiting the methodType production.
	ExitMethodType(c *MethodTypeContext)

	// ExitSimpleProcedureType is called when exiting the simpleProcedureType production.
	ExitSimpleProcedureType(c *SimpleProcedureTypeContext)

	// ExitProcedureReference is called when exiting the procedureReference production.
	ExitProcedureReference(c *ProcedureReferenceContext)

	// ExitProcedureTypeHeading is called when exiting the procedureTypeHeading production.
	ExitProcedureTypeHeading(c *ProcedureTypeHeadingContext)

	// ExitVariantType is called when exiting the variantType production.
	ExitVariantType(c *VariantTypeContext)

	// ExitSimpleType is called when exiting the simpleType production.
	ExitSimpleType(c *SimpleTypeContext)

	// ExitSubRangeType is called when exiting the subRangeType production.
	ExitSubRangeType(c *SubRangeTypeContext)

	// ExitEnumType is called when exiting the enumType production.
	ExitEnumType(c *EnumTypeContext)

	// ExitTypeId is called when exiting the typeId production.
	ExitTypeId(c *TypeIdContext)

	// ExitGenericTypeIdent is called when exiting the genericTypeIdent production.
	ExitGenericTypeIdent(c *GenericTypeIdentContext)

	// ExitGenericDefinition is called when exiting the genericDefinition production.
	ExitGenericDefinition(c *GenericDefinitionContext)

	// ExitSimpleGenericDefinition is called when exiting the simpleGenericDefinition production.
	ExitSimpleGenericDefinition(c *SimpleGenericDefinitionContext)

	// ExitConstrainedGenericDefinition is called when exiting the constrainedGenericDefinition production.
	ExitConstrainedGenericDefinition(c *ConstrainedGenericDefinitionContext)

	// ExitConstrainedGeneric is called when exiting the constrainedGeneric production.
	ExitConstrainedGeneric(c *ConstrainedGenericContext)

	// ExitGenericConstraint is called when exiting the genericConstraint production.
	ExitGenericConstraint(c *GenericConstraintContext)

	// ExitGenericPostfix is called when exiting the genericPostfix production.
	ExitGenericPostfix(c *GenericPostfixContext)

	// ExitClassDecl is called when exiting the classDecl production.
	ExitClassDecl(c *ClassDeclContext)

	// ExitClassTypeTypeDecl is called when exiting the classTypeTypeDecl production.
	ExitClassTypeTypeDecl(c *ClassTypeTypeDeclContext)

	// ExitClassTypeDecl is called when exiting the classTypeDecl production.
	ExitClassTypeDecl(c *ClassTypeDeclContext)

	// ExitClassState is called when exiting the classState production.
	ExitClassState(c *ClassStateContext)

	// ExitClassParent is called when exiting the classParent production.
	ExitClassParent(c *ClassParentContext)

	// ExitClassItem is called when exiting the classItem production.
	ExitClassItem(c *ClassItemContext)

	// ExitClassHelperDecl is called when exiting the classHelperDecl production.
	ExitClassHelperDecl(c *ClassHelperDeclContext)

	// ExitClassHelperItem is called when exiting the classHelperItem production.
	ExitClassHelperItem(c *ClassHelperItemContext)

	// ExitInterfaceTypeDecl is called when exiting the interfaceTypeDecl production.
	ExitInterfaceTypeDecl(c *InterfaceTypeDeclContext)

	// ExitInterfaceKey is called when exiting the interfaceKey production.
	ExitInterfaceKey(c *InterfaceKeyContext)

	// ExitInterfaceGuid is called when exiting the interfaceGuid production.
	ExitInterfaceGuid(c *InterfaceGuidContext)

	// ExitInterfaceItem is called when exiting the interfaceItem production.
	ExitInterfaceItem(c *InterfaceItemContext)

	// ExitObjectDecl is called when exiting the objectDecl production.
	ExitObjectDecl(c *ObjectDeclContext)

	// ExitObjectItem is called when exiting the objectItem production.
	ExitObjectItem(c *ObjectItemContext)

	// ExitRecordDecl is called when exiting the recordDecl production.
	ExitRecordDecl(c *RecordDeclContext)

	// ExitSimpleRecord is called when exiting the simpleRecord production.
	ExitSimpleRecord(c *SimpleRecordContext)

	// ExitVariantRecord is called when exiting the variantRecord production.
	ExitVariantRecord(c *VariantRecordContext)

	// ExitRecordItem is called when exiting the recordItem production.
	ExitRecordItem(c *RecordItemContext)

	// ExitRecordField is called when exiting the recordField production.
	ExitRecordField(c *RecordFieldContext)

	// ExitRecordVariantField is called when exiting the recordVariantField production.
	ExitRecordVariantField(c *RecordVariantFieldContext)

	// ExitRecordVariantSection is called when exiting the recordVariantSection production.
	ExitRecordVariantSection(c *RecordVariantSectionContext)

	// ExitRecordVariant is called when exiting the recordVariant production.
	ExitRecordVariant(c *RecordVariantContext)

	// ExitRecordHelperDecl is called when exiting the recordHelperDecl production.
	ExitRecordHelperDecl(c *RecordHelperDeclContext)

	// ExitRecordHelperItem is called when exiting the recordHelperItem production.
	ExitRecordHelperItem(c *RecordHelperItemContext)

	// ExitClassMethod is called when exiting the classMethod production.
	ExitClassMethod(c *ClassMethodContext)

	// ExitClassField is called when exiting the classField production.
	ExitClassField(c *ClassFieldContext)

	// ExitClassProperty is called when exiting the classProperty production.
	ExitClassProperty(c *ClassPropertyContext)

	// ExitClassPropertyArray is called when exiting the classPropertyArray production.
	ExitClassPropertyArray(c *ClassPropertyArrayContext)

	// ExitClassPropertyIndex is called when exiting the classPropertyIndex production.
	ExitClassPropertyIndex(c *ClassPropertyIndexContext)

	// ExitClassPropertySpecifier is called when exiting the classPropertySpecifier production.
	ExitClassPropertySpecifier(c *ClassPropertySpecifierContext)

	// ExitClassPropertyEndSpecifier is called when exiting the classPropertyEndSpecifier production.
	ExitClassPropertyEndSpecifier(c *ClassPropertyEndSpecifierContext)

	// ExitClassPropertyReadWrite is called when exiting the classPropertyReadWrite production.
	ExitClassPropertyReadWrite(c *ClassPropertyReadWriteContext)

	// ExitClassPropertyDispInterface is called when exiting the classPropertyDispInterface production.
	ExitClassPropertyDispInterface(c *ClassPropertyDispInterfaceContext)

	// ExitVisibility is called when exiting the visibility production.
	ExitVisibility(c *VisibilityContext)

	// ExitExportedProcHeading is called when exiting the exportedProcHeading production.
	ExitExportedProcHeading(c *ExportedProcHeadingContext)

	// ExitMethodDecl is called when exiting the methodDecl production.
	ExitMethodDecl(c *MethodDeclContext)

	// ExitMethodDeclHeading is called when exiting the methodDeclHeading production.
	ExitMethodDeclHeading(c *MethodDeclHeadingContext)

	// ExitMethodKey is called when exiting the methodKey production.
	ExitMethodKey(c *MethodKeyContext)

	// ExitMethodName is called when exiting the methodName production.
	ExitMethodName(c *MethodNameContext)

	// ExitProcDecl is called when exiting the procDecl production.
	ExitProcDecl(c *ProcDeclContext)

	// ExitProcDeclHeading is called when exiting the procDeclHeading production.
	ExitProcDeclHeading(c *ProcDeclHeadingContext)

	// ExitFormalParameterSection is called when exiting the formalParameterSection production.
	ExitFormalParameterSection(c *FormalParameterSectionContext)

	// ExitFormalParameterList is called when exiting the formalParameterList production.
	ExitFormalParameterList(c *FormalParameterListContext)

	// ExitFormalParameter is called when exiting the formalParameter production.
	ExitFormalParameter(c *FormalParameterContext)

	// ExitParmType is called when exiting the parmType production.
	ExitParmType(c *ParmTypeContext)

	// ExitMethodBody is called when exiting the methodBody production.
	ExitMethodBody(c *MethodBodyContext)

	// ExitProcBody is called when exiting the procBody production.
	ExitProcBody(c *ProcBodyContext)

	// ExitCustomAttribute is called when exiting the customAttribute production.
	ExitCustomAttribute(c *CustomAttributeContext)

	// ExitCustomAttributeList is called when exiting the customAttributeList production.
	ExitCustomAttributeList(c *CustomAttributeListContext)

	// ExitCustomAttributeDecl is called when exiting the customAttributeDecl production.
	ExitCustomAttributeDecl(c *CustomAttributeDeclContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAnonymousExpression is called when exiting the anonymousExpression production.
	ExitAnonymousExpression(c *AnonymousExpressionContext)

	// ExitSimpleExpression is called when exiting the simpleExpression production.
	ExitSimpleExpression(c *SimpleExpressionContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitStringFactor is called when exiting the stringFactor production.
	ExitStringFactor(c *StringFactorContext)

	// ExitSetSection is called when exiting the setSection production.
	ExitSetSection(c *SetSectionContext)

	// ExitDesignator is called when exiting the designator production.
	ExitDesignator(c *DesignatorContext)

	// ExitDesignatorItem is called when exiting the designatorItem production.
	ExitDesignatorItem(c *DesignatorItemContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitColonConstruct is called when exiting the colonConstruct production.
	ExitColonConstruct(c *ColonConstructContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitRelOp is called when exiting the relOp production.
	ExitRelOp(c *RelOpContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitCaseStatement is called when exiting the caseStatement production.
	ExitCaseStatement(c *CaseStatementContext)

	// ExitCaseItem is called when exiting the caseItem production.
	ExitCaseItem(c *CaseItemContext)

	// ExitCaseLabel is called when exiting the caseLabel production.
	ExitCaseLabel(c *CaseLabelContext)

	// ExitRepeatStatement is called when exiting the repeatStatement production.
	ExitRepeatStatement(c *RepeatStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitWithStatement is called when exiting the withStatement production.
	ExitWithStatement(c *WithStatementContext)

	// ExitWithItem is called when exiting the withItem production.
	ExitWithItem(c *WithItemContext)

	// ExitCompoundStatement is called when exiting the compoundStatement production.
	ExitCompoundStatement(c *CompoundStatementContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitSimpleStatement is called when exiting the simpleStatement production.
	ExitSimpleStatement(c *SimpleStatementContext)

	// ExitGotoStatement is called when exiting the gotoStatement production.
	ExitGotoStatement(c *GotoStatementContext)

	// ExitConstExpression is called when exiting the constExpression production.
	ExitConstExpression(c *ConstExpressionContext)

	// ExitRecordConstExpression is called when exiting the recordConstExpression production.
	ExitRecordConstExpression(c *RecordConstExpressionContext)

	// ExitTryStatement is called when exiting the tryStatement production.
	ExitTryStatement(c *TryStatementContext)

	// ExitHandlerList is called when exiting the handlerList production.
	ExitHandlerList(c *HandlerListContext)

	// ExitHandler is called when exiting the handler production.
	ExitHandler(c *HandlerContext)

	// ExitHandlerIdent is called when exiting the handlerIdent production.
	ExitHandlerIdent(c *HandlerIdentContext)

	// ExitHandlerStatement is called when exiting the handlerStatement production.
	ExitHandlerStatement(c *HandlerStatementContext)

	// ExitRaiseStatement is called when exiting the raiseStatement production.
	ExitRaiseStatement(c *RaiseStatementContext)

	// ExitAssemblerStatement is called when exiting the assemblerStatement production.
	ExitAssemblerStatement(c *AssemblerStatementContext)

	// ExitMethodDirective is called when exiting the methodDirective production.
	ExitMethodDirective(c *MethodDirectiveContext)

	// ExitFunctionDirective is called when exiting the functionDirective production.
	ExitFunctionDirective(c *FunctionDirectiveContext)

	// ExitReintroduceDirective is called when exiting the reintroduceDirective production.
	ExitReintroduceDirective(c *ReintroduceDirectiveContext)

	// ExitOverloadDirective is called when exiting the overloadDirective production.
	ExitOverloadDirective(c *OverloadDirectiveContext)

	// ExitBindingDirective is called when exiting the bindingDirective production.
	ExitBindingDirective(c *BindingDirectiveContext)

	// ExitAbstractDirective is called when exiting the abstractDirective production.
	ExitAbstractDirective(c *AbstractDirectiveContext)

	// ExitInlineDirective is called when exiting the inlineDirective production.
	ExitInlineDirective(c *InlineDirectiveContext)

	// ExitCallConvention is called when exiting the callConvention production.
	ExitCallConvention(c *CallConventionContext)

	// ExitCallConventionNoSemi is called when exiting the callConventionNoSemi production.
	ExitCallConventionNoSemi(c *CallConventionNoSemiContext)

	// ExitOldCallConventionDirective is called when exiting the oldCallConventionDirective production.
	ExitOldCallConventionDirective(c *OldCallConventionDirectiveContext)

	// ExitHintingDirective is called when exiting the hintingDirective production.
	ExitHintingDirective(c *HintingDirectiveContext)

	// ExitExternalDirective is called when exiting the externalDirective production.
	ExitExternalDirective(c *ExternalDirectiveContext)

	// ExitExternalSpecifier is called when exiting the externalSpecifier production.
	ExitExternalSpecifier(c *ExternalSpecifierContext)

	// ExitDispIDDirective is called when exiting the dispIDDirective production.
	ExitDispIDDirective(c *DispIDDirectiveContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitUsedKeywordsAsNames is called when exiting the usedKeywordsAsNames production.
	ExitUsedKeywordsAsNames(c *UsedKeywordsAsNamesContext)

	// ExitIdentList is called when exiting the identList production.
	ExitIdentList(c *IdentListContext)

	// ExitIdentListFlat is called when exiting the identListFlat production.
	ExitIdentListFlat(c *IdentListFlatContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitIntNum is called when exiting the intNum production.
	ExitIntNum(c *IntNumContext)

	// ExitRealNum is called when exiting the realNum production.
	ExitRealNum(c *RealNumContext)

	// ExitNamespacedQualifiedIdent is called when exiting the namespacedQualifiedIdent production.
	ExitNamespacedQualifiedIdent(c *NamespacedQualifiedIdentContext)

	// ExitNamespaceName is called when exiting the namespaceName production.
	ExitNamespaceName(c *NamespaceNameContext)

	// ExitQualifiedIdent is called when exiting the qualifiedIdent production.
	ExitQualifiedIdent(c *QualifiedIdentContext)
}
