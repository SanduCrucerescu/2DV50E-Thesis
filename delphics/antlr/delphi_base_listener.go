// Code generated from Delphi.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Delphi

import "github.com/antlr4-go/antlr/v4"

// BaseDelphiListener is a complete listener for a parse tree produced by DelphiParser.
type BaseDelphiListener struct{}

var _ DelphiListener = &BaseDelphiListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDelphiListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDelphiListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDelphiListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDelphiListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseDelphiListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseDelphiListener) ExitFile(ctx *FileContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseDelphiListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseDelphiListener) ExitProgram(ctx *ProgramContext) {}

// EnterProgramHead is called when production programHead is entered.
func (s *BaseDelphiListener) EnterProgramHead(ctx *ProgramHeadContext) {}

// ExitProgramHead is called when production programHead is exited.
func (s *BaseDelphiListener) ExitProgramHead(ctx *ProgramHeadContext) {}

// EnterProgramParmSeq is called when production programParmSeq is entered.
func (s *BaseDelphiListener) EnterProgramParmSeq(ctx *ProgramParmSeqContext) {}

// ExitProgramParmSeq is called when production programParmSeq is exited.
func (s *BaseDelphiListener) ExitProgramParmSeq(ctx *ProgramParmSeqContext) {}

// EnterLibrary is called when production library is entered.
func (s *BaseDelphiListener) EnterLibrary(ctx *LibraryContext) {}

// ExitLibrary is called when production library is exited.
func (s *BaseDelphiListener) ExitLibrary(ctx *LibraryContext) {}

// EnterLibraryHead is called when production libraryHead is entered.
func (s *BaseDelphiListener) EnterLibraryHead(ctx *LibraryHeadContext) {}

// ExitLibraryHead is called when production libraryHead is exited.
func (s *BaseDelphiListener) ExitLibraryHead(ctx *LibraryHeadContext) {}

// EnterPackageE is called when production packageE is entered.
func (s *BaseDelphiListener) EnterPackageE(ctx *PackageEContext) {}

// ExitPackageE is called when production packageE is exited.
func (s *BaseDelphiListener) ExitPackageE(ctx *PackageEContext) {}

// EnterPackageHead is called when production packageHead is entered.
func (s *BaseDelphiListener) EnterPackageHead(ctx *PackageHeadContext) {}

// ExitPackageHead is called when production packageHead is exited.
func (s *BaseDelphiListener) ExitPackageHead(ctx *PackageHeadContext) {}

// EnterUnit is called when production unit is entered.
func (s *BaseDelphiListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BaseDelphiListener) ExitUnit(ctx *UnitContext) {}

// EnterUnitHead is called when production unitHead is entered.
func (s *BaseDelphiListener) EnterUnitHead(ctx *UnitHeadContext) {}

// ExitUnitHead is called when production unitHead is exited.
func (s *BaseDelphiListener) ExitUnitHead(ctx *UnitHeadContext) {}

// EnterUnitInterface is called when production unitInterface is entered.
func (s *BaseDelphiListener) EnterUnitInterface(ctx *UnitInterfaceContext) {}

// ExitUnitInterface is called when production unitInterface is exited.
func (s *BaseDelphiListener) ExitUnitInterface(ctx *UnitInterfaceContext) {}

// EnterUnitImplementation is called when production unitImplementation is entered.
func (s *BaseDelphiListener) EnterUnitImplementation(ctx *UnitImplementationContext) {}

// ExitUnitImplementation is called when production unitImplementation is exited.
func (s *BaseDelphiListener) ExitUnitImplementation(ctx *UnitImplementationContext) {}

// EnterUnitBlock is called when production unitBlock is entered.
func (s *BaseDelphiListener) EnterUnitBlock(ctx *UnitBlockContext) {}

// ExitUnitBlock is called when production unitBlock is exited.
func (s *BaseDelphiListener) ExitUnitBlock(ctx *UnitBlockContext) {}

// EnterUnitInitialization is called when production unitInitialization is entered.
func (s *BaseDelphiListener) EnterUnitInitialization(ctx *UnitInitializationContext) {}

// ExitUnitInitialization is called when production unitInitialization is exited.
func (s *BaseDelphiListener) ExitUnitInitialization(ctx *UnitInitializationContext) {}

// EnterUnitFinalization is called when production unitFinalization is entered.
func (s *BaseDelphiListener) EnterUnitFinalization(ctx *UnitFinalizationContext) {}

// ExitUnitFinalization is called when production unitFinalization is exited.
func (s *BaseDelphiListener) ExitUnitFinalization(ctx *UnitFinalizationContext) {}

// EnterContainsClause is called when production containsClause is entered.
func (s *BaseDelphiListener) EnterContainsClause(ctx *ContainsClauseContext) {}

// ExitContainsClause is called when production containsClause is exited.
func (s *BaseDelphiListener) ExitContainsClause(ctx *ContainsClauseContext) {}

// EnterRequiresClause is called when production requiresClause is entered.
func (s *BaseDelphiListener) EnterRequiresClause(ctx *RequiresClauseContext) {}

// ExitRequiresClause is called when production requiresClause is exited.
func (s *BaseDelphiListener) ExitRequiresClause(ctx *RequiresClauseContext) {}

// EnterUsesClause is called when production usesClause is entered.
func (s *BaseDelphiListener) EnterUsesClause(ctx *UsesClauseContext) {}

// ExitUsesClause is called when production usesClause is exited.
func (s *BaseDelphiListener) ExitUsesClause(ctx *UsesClauseContext) {}

// EnterUsesFileClause is called when production usesFileClause is entered.
func (s *BaseDelphiListener) EnterUsesFileClause(ctx *UsesFileClauseContext) {}

// ExitUsesFileClause is called when production usesFileClause is exited.
func (s *BaseDelphiListener) ExitUsesFileClause(ctx *UsesFileClauseContext) {}

// EnterNamespaceFileNameList is called when production namespaceFileNameList is entered.
func (s *BaseDelphiListener) EnterNamespaceFileNameList(ctx *NamespaceFileNameListContext) {}

// ExitNamespaceFileNameList is called when production namespaceFileNameList is exited.
func (s *BaseDelphiListener) ExitNamespaceFileNameList(ctx *NamespaceFileNameListContext) {}

// EnterNamespaceFileName is called when production namespaceFileName is entered.
func (s *BaseDelphiListener) EnterNamespaceFileName(ctx *NamespaceFileNameContext) {}

// ExitNamespaceFileName is called when production namespaceFileName is exited.
func (s *BaseDelphiListener) ExitNamespaceFileName(ctx *NamespaceFileNameContext) {}

// EnterNamespaceNameList is called when production namespaceNameList is entered.
func (s *BaseDelphiListener) EnterNamespaceNameList(ctx *NamespaceNameListContext) {}

// ExitNamespaceNameList is called when production namespaceNameList is exited.
func (s *BaseDelphiListener) ExitNamespaceNameList(ctx *NamespaceNameListContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseDelphiListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseDelphiListener) ExitBlock(ctx *BlockContext) {}

// EnterBlockBody is called when production blockBody is entered.
func (s *BaseDelphiListener) EnterBlockBody(ctx *BlockBodyContext) {}

// ExitBlockBody is called when production blockBody is exited.
func (s *BaseDelphiListener) ExitBlockBody(ctx *BlockBodyContext) {}

// EnterDeclSection is called when production declSection is entered.
func (s *BaseDelphiListener) EnterDeclSection(ctx *DeclSectionContext) {}

// ExitDeclSection is called when production declSection is exited.
func (s *BaseDelphiListener) ExitDeclSection(ctx *DeclSectionContext) {}

// EnterInterfaceDecl is called when production interfaceDecl is entered.
func (s *BaseDelphiListener) EnterInterfaceDecl(ctx *InterfaceDeclContext) {}

// ExitInterfaceDecl is called when production interfaceDecl is exited.
func (s *BaseDelphiListener) ExitInterfaceDecl(ctx *InterfaceDeclContext) {}

// EnterLabelDeclSection is called when production labelDeclSection is entered.
func (s *BaseDelphiListener) EnterLabelDeclSection(ctx *LabelDeclSectionContext) {}

// ExitLabelDeclSection is called when production labelDeclSection is exited.
func (s *BaseDelphiListener) ExitLabelDeclSection(ctx *LabelDeclSectionContext) {}

// EnterConstSection is called when production constSection is entered.
func (s *BaseDelphiListener) EnterConstSection(ctx *ConstSectionContext) {}

// ExitConstSection is called when production constSection is exited.
func (s *BaseDelphiListener) ExitConstSection(ctx *ConstSectionContext) {}

// EnterConstKey is called when production constKey is entered.
func (s *BaseDelphiListener) EnterConstKey(ctx *ConstKeyContext) {}

// ExitConstKey is called when production constKey is exited.
func (s *BaseDelphiListener) ExitConstKey(ctx *ConstKeyContext) {}

// EnterConstDeclaration is called when production constDeclaration is entered.
func (s *BaseDelphiListener) EnterConstDeclaration(ctx *ConstDeclarationContext) {}

// ExitConstDeclaration is called when production constDeclaration is exited.
func (s *BaseDelphiListener) ExitConstDeclaration(ctx *ConstDeclarationContext) {}

// EnterTypeSection is called when production typeSection is entered.
func (s *BaseDelphiListener) EnterTypeSection(ctx *TypeSectionContext) {}

// ExitTypeSection is called when production typeSection is exited.
func (s *BaseDelphiListener) ExitTypeSection(ctx *TypeSectionContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseDelphiListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseDelphiListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterVarSection is called when production varSection is entered.
func (s *BaseDelphiListener) EnterVarSection(ctx *VarSectionContext) {}

// ExitVarSection is called when production varSection is exited.
func (s *BaseDelphiListener) ExitVarSection(ctx *VarSectionContext) {}

// EnterVarKey is called when production varKey is entered.
func (s *BaseDelphiListener) EnterVarKey(ctx *VarKeyContext) {}

// ExitVarKey is called when production varKey is exited.
func (s *BaseDelphiListener) ExitVarKey(ctx *VarKeyContext) {}

// EnterVarDeclaration is called when production varDeclaration is entered.
func (s *BaseDelphiListener) EnterVarDeclaration(ctx *VarDeclarationContext) {}

// ExitVarDeclaration is called when production varDeclaration is exited.
func (s *BaseDelphiListener) ExitVarDeclaration(ctx *VarDeclarationContext) {}

// EnterVarValueSpec is called when production varValueSpec is entered.
func (s *BaseDelphiListener) EnterVarValueSpec(ctx *VarValueSpecContext) {}

// ExitVarValueSpec is called when production varValueSpec is exited.
func (s *BaseDelphiListener) ExitVarValueSpec(ctx *VarValueSpecContext) {}

// EnterExportsSection is called when production exportsSection is entered.
func (s *BaseDelphiListener) EnterExportsSection(ctx *ExportsSectionContext) {}

// ExitExportsSection is called when production exportsSection is exited.
func (s *BaseDelphiListener) ExitExportsSection(ctx *ExportsSectionContext) {}

// EnterExportItem is called when production exportItem is entered.
func (s *BaseDelphiListener) EnterExportItem(ctx *ExportItemContext) {}

// ExitExportItem is called when production exportItem is exited.
func (s *BaseDelphiListener) ExitExportItem(ctx *ExportItemContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *BaseDelphiListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *BaseDelphiListener) ExitTypeDecl(ctx *TypeDeclContext) {}

// EnterStrucType is called when production strucType is entered.
func (s *BaseDelphiListener) EnterStrucType(ctx *StrucTypeContext) {}

// ExitStrucType is called when production strucType is exited.
func (s *BaseDelphiListener) ExitStrucType(ctx *StrucTypeContext) {}

// EnterStrucTypePart is called when production strucTypePart is entered.
func (s *BaseDelphiListener) EnterStrucTypePart(ctx *StrucTypePartContext) {}

// ExitStrucTypePart is called when production strucTypePart is exited.
func (s *BaseDelphiListener) ExitStrucTypePart(ctx *StrucTypePartContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseDelphiListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseDelphiListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterArrayIndex is called when production arrayIndex is entered.
func (s *BaseDelphiListener) EnterArrayIndex(ctx *ArrayIndexContext) {}

// ExitArrayIndex is called when production arrayIndex is exited.
func (s *BaseDelphiListener) ExitArrayIndex(ctx *ArrayIndexContext) {}

// EnterArraySubType is called when production arraySubType is entered.
func (s *BaseDelphiListener) EnterArraySubType(ctx *ArraySubTypeContext) {}

// ExitArraySubType is called when production arraySubType is exited.
func (s *BaseDelphiListener) ExitArraySubType(ctx *ArraySubTypeContext) {}

// EnterSetType is called when production setType is entered.
func (s *BaseDelphiListener) EnterSetType(ctx *SetTypeContext) {}

// ExitSetType is called when production setType is exited.
func (s *BaseDelphiListener) ExitSetType(ctx *SetTypeContext) {}

// EnterFileType is called when production fileType is entered.
func (s *BaseDelphiListener) EnterFileType(ctx *FileTypeContext) {}

// ExitFileType is called when production fileType is exited.
func (s *BaseDelphiListener) ExitFileType(ctx *FileTypeContext) {}

// EnterPointerType is called when production pointerType is entered.
func (s *BaseDelphiListener) EnterPointerType(ctx *PointerTypeContext) {}

// ExitPointerType is called when production pointerType is exited.
func (s *BaseDelphiListener) ExitPointerType(ctx *PointerTypeContext) {}

// EnterStringType is called when production stringType is entered.
func (s *BaseDelphiListener) EnterStringType(ctx *StringTypeContext) {}

// ExitStringType is called when production stringType is exited.
func (s *BaseDelphiListener) ExitStringType(ctx *StringTypeContext) {}

// EnterCodePageNumber is called when production codePageNumber is entered.
func (s *BaseDelphiListener) EnterCodePageNumber(ctx *CodePageNumberContext) {}

// ExitCodePageNumber is called when production codePageNumber is exited.
func (s *BaseDelphiListener) ExitCodePageNumber(ctx *CodePageNumberContext) {}

// EnterProcedureType is called when production procedureType is entered.
func (s *BaseDelphiListener) EnterProcedureType(ctx *ProcedureTypeContext) {}

// ExitProcedureType is called when production procedureType is exited.
func (s *BaseDelphiListener) ExitProcedureType(ctx *ProcedureTypeContext) {}

// EnterMethodType is called when production methodType is entered.
func (s *BaseDelphiListener) EnterMethodType(ctx *MethodTypeContext) {}

// ExitMethodType is called when production methodType is exited.
func (s *BaseDelphiListener) ExitMethodType(ctx *MethodTypeContext) {}

// EnterSimpleProcedureType is called when production simpleProcedureType is entered.
func (s *BaseDelphiListener) EnterSimpleProcedureType(ctx *SimpleProcedureTypeContext) {}

// ExitSimpleProcedureType is called when production simpleProcedureType is exited.
func (s *BaseDelphiListener) ExitSimpleProcedureType(ctx *SimpleProcedureTypeContext) {}

// EnterProcedureReference is called when production procedureReference is entered.
func (s *BaseDelphiListener) EnterProcedureReference(ctx *ProcedureReferenceContext) {}

// ExitProcedureReference is called when production procedureReference is exited.
func (s *BaseDelphiListener) ExitProcedureReference(ctx *ProcedureReferenceContext) {}

// EnterProcedureTypeHeading is called when production procedureTypeHeading is entered.
func (s *BaseDelphiListener) EnterProcedureTypeHeading(ctx *ProcedureTypeHeadingContext) {}

// ExitProcedureTypeHeading is called when production procedureTypeHeading is exited.
func (s *BaseDelphiListener) ExitProcedureTypeHeading(ctx *ProcedureTypeHeadingContext) {}

// EnterVariantType is called when production variantType is entered.
func (s *BaseDelphiListener) EnterVariantType(ctx *VariantTypeContext) {}

// ExitVariantType is called when production variantType is exited.
func (s *BaseDelphiListener) ExitVariantType(ctx *VariantTypeContext) {}

// EnterSimpleType is called when production simpleType is entered.
func (s *BaseDelphiListener) EnterSimpleType(ctx *SimpleTypeContext) {}

// ExitSimpleType is called when production simpleType is exited.
func (s *BaseDelphiListener) ExitSimpleType(ctx *SimpleTypeContext) {}

// EnterSubRangeType is called when production subRangeType is entered.
func (s *BaseDelphiListener) EnterSubRangeType(ctx *SubRangeTypeContext) {}

// ExitSubRangeType is called when production subRangeType is exited.
func (s *BaseDelphiListener) ExitSubRangeType(ctx *SubRangeTypeContext) {}

// EnterEnumType is called when production enumType is entered.
func (s *BaseDelphiListener) EnterEnumType(ctx *EnumTypeContext) {}

// ExitEnumType is called when production enumType is exited.
func (s *BaseDelphiListener) ExitEnumType(ctx *EnumTypeContext) {}

// EnterTypeId is called when production typeId is entered.
func (s *BaseDelphiListener) EnterTypeId(ctx *TypeIdContext) {}

// ExitTypeId is called when production typeId is exited.
func (s *BaseDelphiListener) ExitTypeId(ctx *TypeIdContext) {}

// EnterGenericTypeIdent is called when production genericTypeIdent is entered.
func (s *BaseDelphiListener) EnterGenericTypeIdent(ctx *GenericTypeIdentContext) {}

// ExitGenericTypeIdent is called when production genericTypeIdent is exited.
func (s *BaseDelphiListener) ExitGenericTypeIdent(ctx *GenericTypeIdentContext) {}

// EnterGenericDefinition is called when production genericDefinition is entered.
func (s *BaseDelphiListener) EnterGenericDefinition(ctx *GenericDefinitionContext) {}

// ExitGenericDefinition is called when production genericDefinition is exited.
func (s *BaseDelphiListener) ExitGenericDefinition(ctx *GenericDefinitionContext) {}

// EnterSimpleGenericDefinition is called when production simpleGenericDefinition is entered.
func (s *BaseDelphiListener) EnterSimpleGenericDefinition(ctx *SimpleGenericDefinitionContext) {}

// ExitSimpleGenericDefinition is called when production simpleGenericDefinition is exited.
func (s *BaseDelphiListener) ExitSimpleGenericDefinition(ctx *SimpleGenericDefinitionContext) {}

// EnterConstrainedGenericDefinition is called when production constrainedGenericDefinition is entered.
func (s *BaseDelphiListener) EnterConstrainedGenericDefinition(ctx *ConstrainedGenericDefinitionContext) {
}

// ExitConstrainedGenericDefinition is called when production constrainedGenericDefinition is exited.
func (s *BaseDelphiListener) ExitConstrainedGenericDefinition(ctx *ConstrainedGenericDefinitionContext) {
}

// EnterConstrainedGeneric is called when production constrainedGeneric is entered.
func (s *BaseDelphiListener) EnterConstrainedGeneric(ctx *ConstrainedGenericContext) {}

// ExitConstrainedGeneric is called when production constrainedGeneric is exited.
func (s *BaseDelphiListener) ExitConstrainedGeneric(ctx *ConstrainedGenericContext) {}

// EnterGenericConstraint is called when production genericConstraint is entered.
func (s *BaseDelphiListener) EnterGenericConstraint(ctx *GenericConstraintContext) {}

// ExitGenericConstraint is called when production genericConstraint is exited.
func (s *BaseDelphiListener) ExitGenericConstraint(ctx *GenericConstraintContext) {}

// EnterGenericPostfix is called when production genericPostfix is entered.
func (s *BaseDelphiListener) EnterGenericPostfix(ctx *GenericPostfixContext) {}

// ExitGenericPostfix is called when production genericPostfix is exited.
func (s *BaseDelphiListener) ExitGenericPostfix(ctx *GenericPostfixContext) {}

// EnterClassDecl is called when production classDecl is entered.
func (s *BaseDelphiListener) EnterClassDecl(ctx *ClassDeclContext) {}

// ExitClassDecl is called when production classDecl is exited.
func (s *BaseDelphiListener) ExitClassDecl(ctx *ClassDeclContext) {}

// EnterClassTypeTypeDecl is called when production classTypeTypeDecl is entered.
func (s *BaseDelphiListener) EnterClassTypeTypeDecl(ctx *ClassTypeTypeDeclContext) {}

// ExitClassTypeTypeDecl is called when production classTypeTypeDecl is exited.
func (s *BaseDelphiListener) ExitClassTypeTypeDecl(ctx *ClassTypeTypeDeclContext) {}

// EnterClassTypeDecl is called when production classTypeDecl is entered.
func (s *BaseDelphiListener) EnterClassTypeDecl(ctx *ClassTypeDeclContext) {}

// ExitClassTypeDecl is called when production classTypeDecl is exited.
func (s *BaseDelphiListener) ExitClassTypeDecl(ctx *ClassTypeDeclContext) {}

// EnterClassState is called when production classState is entered.
func (s *BaseDelphiListener) EnterClassState(ctx *ClassStateContext) {}

// ExitClassState is called when production classState is exited.
func (s *BaseDelphiListener) ExitClassState(ctx *ClassStateContext) {}

// EnterClassParent is called when production classParent is entered.
func (s *BaseDelphiListener) EnterClassParent(ctx *ClassParentContext) {}

// ExitClassParent is called when production classParent is exited.
func (s *BaseDelphiListener) ExitClassParent(ctx *ClassParentContext) {}

// EnterClassItem is called when production classItem is entered.
func (s *BaseDelphiListener) EnterClassItem(ctx *ClassItemContext) {}

// ExitClassItem is called when production classItem is exited.
func (s *BaseDelphiListener) ExitClassItem(ctx *ClassItemContext) {}

// EnterClassHelperDecl is called when production classHelperDecl is entered.
func (s *BaseDelphiListener) EnterClassHelperDecl(ctx *ClassHelperDeclContext) {}

// ExitClassHelperDecl is called when production classHelperDecl is exited.
func (s *BaseDelphiListener) ExitClassHelperDecl(ctx *ClassHelperDeclContext) {}

// EnterClassHelperItem is called when production classHelperItem is entered.
func (s *BaseDelphiListener) EnterClassHelperItem(ctx *ClassHelperItemContext) {}

// ExitClassHelperItem is called when production classHelperItem is exited.
func (s *BaseDelphiListener) ExitClassHelperItem(ctx *ClassHelperItemContext) {}

// EnterInterfaceTypeDecl is called when production interfaceTypeDecl is entered.
func (s *BaseDelphiListener) EnterInterfaceTypeDecl(ctx *InterfaceTypeDeclContext) {}

// ExitInterfaceTypeDecl is called when production interfaceTypeDecl is exited.
func (s *BaseDelphiListener) ExitInterfaceTypeDecl(ctx *InterfaceTypeDeclContext) {}

// EnterInterfaceKey is called when production interfaceKey is entered.
func (s *BaseDelphiListener) EnterInterfaceKey(ctx *InterfaceKeyContext) {}

// ExitInterfaceKey is called when production interfaceKey is exited.
func (s *BaseDelphiListener) ExitInterfaceKey(ctx *InterfaceKeyContext) {}

// EnterInterfaceGuid is called when production interfaceGuid is entered.
func (s *BaseDelphiListener) EnterInterfaceGuid(ctx *InterfaceGuidContext) {}

// ExitInterfaceGuid is called when production interfaceGuid is exited.
func (s *BaseDelphiListener) ExitInterfaceGuid(ctx *InterfaceGuidContext) {}

// EnterInterfaceItem is called when production interfaceItem is entered.
func (s *BaseDelphiListener) EnterInterfaceItem(ctx *InterfaceItemContext) {}

// ExitInterfaceItem is called when production interfaceItem is exited.
func (s *BaseDelphiListener) ExitInterfaceItem(ctx *InterfaceItemContext) {}

// EnterObjectDecl is called when production objectDecl is entered.
func (s *BaseDelphiListener) EnterObjectDecl(ctx *ObjectDeclContext) {}

// ExitObjectDecl is called when production objectDecl is exited.
func (s *BaseDelphiListener) ExitObjectDecl(ctx *ObjectDeclContext) {}

// EnterObjectItem is called when production objectItem is entered.
func (s *BaseDelphiListener) EnterObjectItem(ctx *ObjectItemContext) {}

// ExitObjectItem is called when production objectItem is exited.
func (s *BaseDelphiListener) ExitObjectItem(ctx *ObjectItemContext) {}

// EnterRecordDecl is called when production recordDecl is entered.
func (s *BaseDelphiListener) EnterRecordDecl(ctx *RecordDeclContext) {}

// ExitRecordDecl is called when production recordDecl is exited.
func (s *BaseDelphiListener) ExitRecordDecl(ctx *RecordDeclContext) {}

// EnterSimpleRecord is called when production simpleRecord is entered.
func (s *BaseDelphiListener) EnterSimpleRecord(ctx *SimpleRecordContext) {}

// ExitSimpleRecord is called when production simpleRecord is exited.
func (s *BaseDelphiListener) ExitSimpleRecord(ctx *SimpleRecordContext) {}

// EnterVariantRecord is called when production variantRecord is entered.
func (s *BaseDelphiListener) EnterVariantRecord(ctx *VariantRecordContext) {}

// ExitVariantRecord is called when production variantRecord is exited.
func (s *BaseDelphiListener) ExitVariantRecord(ctx *VariantRecordContext) {}

// EnterRecordItem is called when production recordItem is entered.
func (s *BaseDelphiListener) EnterRecordItem(ctx *RecordItemContext) {}

// ExitRecordItem is called when production recordItem is exited.
func (s *BaseDelphiListener) ExitRecordItem(ctx *RecordItemContext) {}

// EnterRecordField is called when production recordField is entered.
func (s *BaseDelphiListener) EnterRecordField(ctx *RecordFieldContext) {}

// ExitRecordField is called when production recordField is exited.
func (s *BaseDelphiListener) ExitRecordField(ctx *RecordFieldContext) {}

// EnterRecordVariantField is called when production recordVariantField is entered.
func (s *BaseDelphiListener) EnterRecordVariantField(ctx *RecordVariantFieldContext) {}

// ExitRecordVariantField is called when production recordVariantField is exited.
func (s *BaseDelphiListener) ExitRecordVariantField(ctx *RecordVariantFieldContext) {}

// EnterRecordVariantSection is called when production recordVariantSection is entered.
func (s *BaseDelphiListener) EnterRecordVariantSection(ctx *RecordVariantSectionContext) {}

// ExitRecordVariantSection is called when production recordVariantSection is exited.
func (s *BaseDelphiListener) ExitRecordVariantSection(ctx *RecordVariantSectionContext) {}

// EnterRecordVariant is called when production recordVariant is entered.
func (s *BaseDelphiListener) EnterRecordVariant(ctx *RecordVariantContext) {}

// ExitRecordVariant is called when production recordVariant is exited.
func (s *BaseDelphiListener) ExitRecordVariant(ctx *RecordVariantContext) {}

// EnterRecordHelperDecl is called when production recordHelperDecl is entered.
func (s *BaseDelphiListener) EnterRecordHelperDecl(ctx *RecordHelperDeclContext) {}

// ExitRecordHelperDecl is called when production recordHelperDecl is exited.
func (s *BaseDelphiListener) ExitRecordHelperDecl(ctx *RecordHelperDeclContext) {}

// EnterRecordHelperItem is called when production recordHelperItem is entered.
func (s *BaseDelphiListener) EnterRecordHelperItem(ctx *RecordHelperItemContext) {}

// ExitRecordHelperItem is called when production recordHelperItem is exited.
func (s *BaseDelphiListener) ExitRecordHelperItem(ctx *RecordHelperItemContext) {}

// EnterClassMethod is called when production classMethod is entered.
func (s *BaseDelphiListener) EnterClassMethod(ctx *ClassMethodContext) {}

// ExitClassMethod is called when production classMethod is exited.
func (s *BaseDelphiListener) ExitClassMethod(ctx *ClassMethodContext) {}

// EnterClassField is called when production classField is entered.
func (s *BaseDelphiListener) EnterClassField(ctx *ClassFieldContext) {}

// ExitClassField is called when production classField is exited.
func (s *BaseDelphiListener) ExitClassField(ctx *ClassFieldContext) {}

// EnterClassProperty is called when production classProperty is entered.
func (s *BaseDelphiListener) EnterClassProperty(ctx *ClassPropertyContext) {}

// ExitClassProperty is called when production classProperty is exited.
func (s *BaseDelphiListener) ExitClassProperty(ctx *ClassPropertyContext) {}

// EnterClassPropertyArray is called when production classPropertyArray is entered.
func (s *BaseDelphiListener) EnterClassPropertyArray(ctx *ClassPropertyArrayContext) {}

// ExitClassPropertyArray is called when production classPropertyArray is exited.
func (s *BaseDelphiListener) ExitClassPropertyArray(ctx *ClassPropertyArrayContext) {}

// EnterClassPropertyIndex is called when production classPropertyIndex is entered.
func (s *BaseDelphiListener) EnterClassPropertyIndex(ctx *ClassPropertyIndexContext) {}

// ExitClassPropertyIndex is called when production classPropertyIndex is exited.
func (s *BaseDelphiListener) ExitClassPropertyIndex(ctx *ClassPropertyIndexContext) {}

// EnterClassPropertySpecifier is called when production classPropertySpecifier is entered.
func (s *BaseDelphiListener) EnterClassPropertySpecifier(ctx *ClassPropertySpecifierContext) {}

// ExitClassPropertySpecifier is called when production classPropertySpecifier is exited.
func (s *BaseDelphiListener) ExitClassPropertySpecifier(ctx *ClassPropertySpecifierContext) {}

// EnterClassPropertyEndSpecifier is called when production classPropertyEndSpecifier is entered.
func (s *BaseDelphiListener) EnterClassPropertyEndSpecifier(ctx *ClassPropertyEndSpecifierContext) {}

// ExitClassPropertyEndSpecifier is called when production classPropertyEndSpecifier is exited.
func (s *BaseDelphiListener) ExitClassPropertyEndSpecifier(ctx *ClassPropertyEndSpecifierContext) {}

// EnterClassPropertyReadWrite is called when production classPropertyReadWrite is entered.
func (s *BaseDelphiListener) EnterClassPropertyReadWrite(ctx *ClassPropertyReadWriteContext) {}

// ExitClassPropertyReadWrite is called when production classPropertyReadWrite is exited.
func (s *BaseDelphiListener) ExitClassPropertyReadWrite(ctx *ClassPropertyReadWriteContext) {}

// EnterClassPropertyDispInterface is called when production classPropertyDispInterface is entered.
func (s *BaseDelphiListener) EnterClassPropertyDispInterface(ctx *ClassPropertyDispInterfaceContext) {
}

// ExitClassPropertyDispInterface is called when production classPropertyDispInterface is exited.
func (s *BaseDelphiListener) ExitClassPropertyDispInterface(ctx *ClassPropertyDispInterfaceContext) {}

// EnterVisibility is called when production visibility is entered.
func (s *BaseDelphiListener) EnterVisibility(ctx *VisibilityContext) {}

// ExitVisibility is called when production visibility is exited.
func (s *BaseDelphiListener) ExitVisibility(ctx *VisibilityContext) {}

// EnterExportedProcHeading is called when production exportedProcHeading is entered.
func (s *BaseDelphiListener) EnterExportedProcHeading(ctx *ExportedProcHeadingContext) {}

// ExitExportedProcHeading is called when production exportedProcHeading is exited.
func (s *BaseDelphiListener) ExitExportedProcHeading(ctx *ExportedProcHeadingContext) {}

// EnterMethodDecl is called when production methodDecl is entered.
func (s *BaseDelphiListener) EnterMethodDecl(ctx *MethodDeclContext) {}

// ExitMethodDecl is called when production methodDecl is exited.
func (s *BaseDelphiListener) ExitMethodDecl(ctx *MethodDeclContext) {}

// EnterMethodDeclHeading is called when production methodDeclHeading is entered.
func (s *BaseDelphiListener) EnterMethodDeclHeading(ctx *MethodDeclHeadingContext) {}

// ExitMethodDeclHeading is called when production methodDeclHeading is exited.
func (s *BaseDelphiListener) ExitMethodDeclHeading(ctx *MethodDeclHeadingContext) {}

// EnterMethodKey is called when production methodKey is entered.
func (s *BaseDelphiListener) EnterMethodKey(ctx *MethodKeyContext) {}

// ExitMethodKey is called when production methodKey is exited.
func (s *BaseDelphiListener) ExitMethodKey(ctx *MethodKeyContext) {}

// EnterMethodName is called when production methodName is entered.
func (s *BaseDelphiListener) EnterMethodName(ctx *MethodNameContext) {}

// ExitMethodName is called when production methodName is exited.
func (s *BaseDelphiListener) ExitMethodName(ctx *MethodNameContext) {}

// EnterProcDecl is called when production procDecl is entered.
func (s *BaseDelphiListener) EnterProcDecl(ctx *ProcDeclContext) {}

// ExitProcDecl is called when production procDecl is exited.
func (s *BaseDelphiListener) ExitProcDecl(ctx *ProcDeclContext) {}

// EnterProcDeclHeading is called when production procDeclHeading is entered.
func (s *BaseDelphiListener) EnterProcDeclHeading(ctx *ProcDeclHeadingContext) {}

// ExitProcDeclHeading is called when production procDeclHeading is exited.
func (s *BaseDelphiListener) ExitProcDeclHeading(ctx *ProcDeclHeadingContext) {}

// EnterFormalParameterSection is called when production formalParameterSection is entered.
func (s *BaseDelphiListener) EnterFormalParameterSection(ctx *FormalParameterSectionContext) {}

// ExitFormalParameterSection is called when production formalParameterSection is exited.
func (s *BaseDelphiListener) ExitFormalParameterSection(ctx *FormalParameterSectionContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BaseDelphiListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BaseDelphiListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterFormalParameter is called when production formalParameter is entered.
func (s *BaseDelphiListener) EnterFormalParameter(ctx *FormalParameterContext) {}

// ExitFormalParameter is called when production formalParameter is exited.
func (s *BaseDelphiListener) ExitFormalParameter(ctx *FormalParameterContext) {}

// EnterParmType is called when production parmType is entered.
func (s *BaseDelphiListener) EnterParmType(ctx *ParmTypeContext) {}

// ExitParmType is called when production parmType is exited.
func (s *BaseDelphiListener) ExitParmType(ctx *ParmTypeContext) {}

// EnterMethodBody is called when production methodBody is entered.
func (s *BaseDelphiListener) EnterMethodBody(ctx *MethodBodyContext) {}

// ExitMethodBody is called when production methodBody is exited.
func (s *BaseDelphiListener) ExitMethodBody(ctx *MethodBodyContext) {}

// EnterProcBody is called when production procBody is entered.
func (s *BaseDelphiListener) EnterProcBody(ctx *ProcBodyContext) {}

// ExitProcBody is called when production procBody is exited.
func (s *BaseDelphiListener) ExitProcBody(ctx *ProcBodyContext) {}

// EnterCustomAttribute is called when production customAttribute is entered.
func (s *BaseDelphiListener) EnterCustomAttribute(ctx *CustomAttributeContext) {}

// ExitCustomAttribute is called when production customAttribute is exited.
func (s *BaseDelphiListener) ExitCustomAttribute(ctx *CustomAttributeContext) {}

// EnterCustomAttributeList is called when production customAttributeList is entered.
func (s *BaseDelphiListener) EnterCustomAttributeList(ctx *CustomAttributeListContext) {}

// ExitCustomAttributeList is called when production customAttributeList is exited.
func (s *BaseDelphiListener) ExitCustomAttributeList(ctx *CustomAttributeListContext) {}

// EnterCustomAttributeDecl is called when production customAttributeDecl is entered.
func (s *BaseDelphiListener) EnterCustomAttributeDecl(ctx *CustomAttributeDeclContext) {}

// ExitCustomAttributeDecl is called when production customAttributeDecl is exited.
func (s *BaseDelphiListener) ExitCustomAttributeDecl(ctx *CustomAttributeDeclContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDelphiListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDelphiListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAnonymousExpression is called when production anonymousExpression is entered.
func (s *BaseDelphiListener) EnterAnonymousExpression(ctx *AnonymousExpressionContext) {}

// ExitAnonymousExpression is called when production anonymousExpression is exited.
func (s *BaseDelphiListener) ExitAnonymousExpression(ctx *AnonymousExpressionContext) {}

// EnterSimpleExpression is called when production simpleExpression is entered.
func (s *BaseDelphiListener) EnterSimpleExpression(ctx *SimpleExpressionContext) {}

// ExitSimpleExpression is called when production simpleExpression is exited.
func (s *BaseDelphiListener) ExitSimpleExpression(ctx *SimpleExpressionContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseDelphiListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseDelphiListener) ExitFactor(ctx *FactorContext) {}

// EnterStringFactor is called when production stringFactor is entered.
func (s *BaseDelphiListener) EnterStringFactor(ctx *StringFactorContext) {}

// ExitStringFactor is called when production stringFactor is exited.
func (s *BaseDelphiListener) ExitStringFactor(ctx *StringFactorContext) {}

// EnterSetSection is called when production setSection is entered.
func (s *BaseDelphiListener) EnterSetSection(ctx *SetSectionContext) {}

// ExitSetSection is called when production setSection is exited.
func (s *BaseDelphiListener) ExitSetSection(ctx *SetSectionContext) {}

// EnterDesignator is called when production designator is entered.
func (s *BaseDelphiListener) EnterDesignator(ctx *DesignatorContext) {}

// ExitDesignator is called when production designator is exited.
func (s *BaseDelphiListener) ExitDesignator(ctx *DesignatorContext) {}

// EnterDesignatorItem is called when production designatorItem is entered.
func (s *BaseDelphiListener) EnterDesignatorItem(ctx *DesignatorItemContext) {}

// ExitDesignatorItem is called when production designatorItem is exited.
func (s *BaseDelphiListener) ExitDesignatorItem(ctx *DesignatorItemContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseDelphiListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseDelphiListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterColonConstruct is called when production colonConstruct is entered.
func (s *BaseDelphiListener) EnterColonConstruct(ctx *ColonConstructContext) {}

// ExitColonConstruct is called when production colonConstruct is exited.
func (s *BaseDelphiListener) ExitColonConstruct(ctx *ColonConstructContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseDelphiListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseDelphiListener) ExitOperator(ctx *OperatorContext) {}

// EnterRelOp is called when production relOp is entered.
func (s *BaseDelphiListener) EnterRelOp(ctx *RelOpContext) {}

// ExitRelOp is called when production relOp is exited.
func (s *BaseDelphiListener) ExitRelOp(ctx *RelOpContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseDelphiListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseDelphiListener) ExitStatement(ctx *StatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseDelphiListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseDelphiListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterCaseStatement is called when production caseStatement is entered.
func (s *BaseDelphiListener) EnterCaseStatement(ctx *CaseStatementContext) {}

// ExitCaseStatement is called when production caseStatement is exited.
func (s *BaseDelphiListener) ExitCaseStatement(ctx *CaseStatementContext) {}

// EnterCaseItem is called when production caseItem is entered.
func (s *BaseDelphiListener) EnterCaseItem(ctx *CaseItemContext) {}

// ExitCaseItem is called when production caseItem is exited.
func (s *BaseDelphiListener) ExitCaseItem(ctx *CaseItemContext) {}

// EnterCaseLabel is called when production caseLabel is entered.
func (s *BaseDelphiListener) EnterCaseLabel(ctx *CaseLabelContext) {}

// ExitCaseLabel is called when production caseLabel is exited.
func (s *BaseDelphiListener) ExitCaseLabel(ctx *CaseLabelContext) {}

// EnterRepeatStatement is called when production repeatStatement is entered.
func (s *BaseDelphiListener) EnterRepeatStatement(ctx *RepeatStatementContext) {}

// ExitRepeatStatement is called when production repeatStatement is exited.
func (s *BaseDelphiListener) ExitRepeatStatement(ctx *RepeatStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseDelphiListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseDelphiListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseDelphiListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseDelphiListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterWithStatement is called when production withStatement is entered.
func (s *BaseDelphiListener) EnterWithStatement(ctx *WithStatementContext) {}

// ExitWithStatement is called when production withStatement is exited.
func (s *BaseDelphiListener) ExitWithStatement(ctx *WithStatementContext) {}

// EnterWithItem is called when production withItem is entered.
func (s *BaseDelphiListener) EnterWithItem(ctx *WithItemContext) {}

// ExitWithItem is called when production withItem is exited.
func (s *BaseDelphiListener) ExitWithItem(ctx *WithItemContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *BaseDelphiListener) EnterCompoundStatement(ctx *CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *BaseDelphiListener) ExitCompoundStatement(ctx *CompoundStatementContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseDelphiListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseDelphiListener) ExitStatementList(ctx *StatementListContext) {}

// EnterSimpleStatement is called when production simpleStatement is entered.
func (s *BaseDelphiListener) EnterSimpleStatement(ctx *SimpleStatementContext) {}

// ExitSimpleStatement is called when production simpleStatement is exited.
func (s *BaseDelphiListener) ExitSimpleStatement(ctx *SimpleStatementContext) {}

// EnterGotoStatement is called when production gotoStatement is entered.
func (s *BaseDelphiListener) EnterGotoStatement(ctx *GotoStatementContext) {}

// ExitGotoStatement is called when production gotoStatement is exited.
func (s *BaseDelphiListener) ExitGotoStatement(ctx *GotoStatementContext) {}

// EnterConstExpression is called when production constExpression is entered.
func (s *BaseDelphiListener) EnterConstExpression(ctx *ConstExpressionContext) {}

// ExitConstExpression is called when production constExpression is exited.
func (s *BaseDelphiListener) ExitConstExpression(ctx *ConstExpressionContext) {}

// EnterRecordConstExpression is called when production recordConstExpression is entered.
func (s *BaseDelphiListener) EnterRecordConstExpression(ctx *RecordConstExpressionContext) {}

// ExitRecordConstExpression is called when production recordConstExpression is exited.
func (s *BaseDelphiListener) ExitRecordConstExpression(ctx *RecordConstExpressionContext) {}

// EnterTryStatement is called when production tryStatement is entered.
func (s *BaseDelphiListener) EnterTryStatement(ctx *TryStatementContext) {}

// ExitTryStatement is called when production tryStatement is exited.
func (s *BaseDelphiListener) ExitTryStatement(ctx *TryStatementContext) {}

// EnterHandlerList is called when production handlerList is entered.
func (s *BaseDelphiListener) EnterHandlerList(ctx *HandlerListContext) {}

// ExitHandlerList is called when production handlerList is exited.
func (s *BaseDelphiListener) ExitHandlerList(ctx *HandlerListContext) {}

// EnterHandler is called when production handler is entered.
func (s *BaseDelphiListener) EnterHandler(ctx *HandlerContext) {}

// ExitHandler is called when production handler is exited.
func (s *BaseDelphiListener) ExitHandler(ctx *HandlerContext) {}

// EnterHandlerIdent is called when production handlerIdent is entered.
func (s *BaseDelphiListener) EnterHandlerIdent(ctx *HandlerIdentContext) {}

// ExitHandlerIdent is called when production handlerIdent is exited.
func (s *BaseDelphiListener) ExitHandlerIdent(ctx *HandlerIdentContext) {}

// EnterHandlerStatement is called when production handlerStatement is entered.
func (s *BaseDelphiListener) EnterHandlerStatement(ctx *HandlerStatementContext) {}

// ExitHandlerStatement is called when production handlerStatement is exited.
func (s *BaseDelphiListener) ExitHandlerStatement(ctx *HandlerStatementContext) {}

// EnterRaiseStatement is called when production raiseStatement is entered.
func (s *BaseDelphiListener) EnterRaiseStatement(ctx *RaiseStatementContext) {}

// ExitRaiseStatement is called when production raiseStatement is exited.
func (s *BaseDelphiListener) ExitRaiseStatement(ctx *RaiseStatementContext) {}

// EnterAssemblerStatement is called when production assemblerStatement is entered.
func (s *BaseDelphiListener) EnterAssemblerStatement(ctx *AssemblerStatementContext) {}

// ExitAssemblerStatement is called when production assemblerStatement is exited.
func (s *BaseDelphiListener) ExitAssemblerStatement(ctx *AssemblerStatementContext) {}

// EnterMethodDirective is called when production methodDirective is entered.
func (s *BaseDelphiListener) EnterMethodDirective(ctx *MethodDirectiveContext) {}

// ExitMethodDirective is called when production methodDirective is exited.
func (s *BaseDelphiListener) ExitMethodDirective(ctx *MethodDirectiveContext) {}

// EnterFunctionDirective is called when production functionDirective is entered.
func (s *BaseDelphiListener) EnterFunctionDirective(ctx *FunctionDirectiveContext) {}

// ExitFunctionDirective is called when production functionDirective is exited.
func (s *BaseDelphiListener) ExitFunctionDirective(ctx *FunctionDirectiveContext) {}

// EnterReintroduceDirective is called when production reintroduceDirective is entered.
func (s *BaseDelphiListener) EnterReintroduceDirective(ctx *ReintroduceDirectiveContext) {}

// ExitReintroduceDirective is called when production reintroduceDirective is exited.
func (s *BaseDelphiListener) ExitReintroduceDirective(ctx *ReintroduceDirectiveContext) {}

// EnterOverloadDirective is called when production overloadDirective is entered.
func (s *BaseDelphiListener) EnterOverloadDirective(ctx *OverloadDirectiveContext) {}

// ExitOverloadDirective is called when production overloadDirective is exited.
func (s *BaseDelphiListener) ExitOverloadDirective(ctx *OverloadDirectiveContext) {}

// EnterBindingDirective is called when production bindingDirective is entered.
func (s *BaseDelphiListener) EnterBindingDirective(ctx *BindingDirectiveContext) {}

// ExitBindingDirective is called when production bindingDirective is exited.
func (s *BaseDelphiListener) ExitBindingDirective(ctx *BindingDirectiveContext) {}

// EnterAbstractDirective is called when production abstractDirective is entered.
func (s *BaseDelphiListener) EnterAbstractDirective(ctx *AbstractDirectiveContext) {}

// ExitAbstractDirective is called when production abstractDirective is exited.
func (s *BaseDelphiListener) ExitAbstractDirective(ctx *AbstractDirectiveContext) {}

// EnterInlineDirective is called when production inlineDirective is entered.
func (s *BaseDelphiListener) EnterInlineDirective(ctx *InlineDirectiveContext) {}

// ExitInlineDirective is called when production inlineDirective is exited.
func (s *BaseDelphiListener) ExitInlineDirective(ctx *InlineDirectiveContext) {}

// EnterCallConvention is called when production callConvention is entered.
func (s *BaseDelphiListener) EnterCallConvention(ctx *CallConventionContext) {}

// ExitCallConvention is called when production callConvention is exited.
func (s *BaseDelphiListener) ExitCallConvention(ctx *CallConventionContext) {}

// EnterCallConventionNoSemi is called when production callConventionNoSemi is entered.
func (s *BaseDelphiListener) EnterCallConventionNoSemi(ctx *CallConventionNoSemiContext) {}

// ExitCallConventionNoSemi is called when production callConventionNoSemi is exited.
func (s *BaseDelphiListener) ExitCallConventionNoSemi(ctx *CallConventionNoSemiContext) {}

// EnterOldCallConventionDirective is called when production oldCallConventionDirective is entered.
func (s *BaseDelphiListener) EnterOldCallConventionDirective(ctx *OldCallConventionDirectiveContext) {
}

// ExitOldCallConventionDirective is called when production oldCallConventionDirective is exited.
func (s *BaseDelphiListener) ExitOldCallConventionDirective(ctx *OldCallConventionDirectiveContext) {}

// EnterHintingDirective is called when production hintingDirective is entered.
func (s *BaseDelphiListener) EnterHintingDirective(ctx *HintingDirectiveContext) {}

// ExitHintingDirective is called when production hintingDirective is exited.
func (s *BaseDelphiListener) ExitHintingDirective(ctx *HintingDirectiveContext) {}

// EnterExternalDirective is called when production externalDirective is entered.
func (s *BaseDelphiListener) EnterExternalDirective(ctx *ExternalDirectiveContext) {}

// ExitExternalDirective is called when production externalDirective is exited.
func (s *BaseDelphiListener) ExitExternalDirective(ctx *ExternalDirectiveContext) {}

// EnterExternalSpecifier is called when production externalSpecifier is entered.
func (s *BaseDelphiListener) EnterExternalSpecifier(ctx *ExternalSpecifierContext) {}

// ExitExternalSpecifier is called when production externalSpecifier is exited.
func (s *BaseDelphiListener) ExitExternalSpecifier(ctx *ExternalSpecifierContext) {}

// EnterDispIDDirective is called when production dispIDDirective is entered.
func (s *BaseDelphiListener) EnterDispIDDirective(ctx *DispIDDirectiveContext) {}

// ExitDispIDDirective is called when production dispIDDirective is exited.
func (s *BaseDelphiListener) ExitDispIDDirective(ctx *DispIDDirectiveContext) {}

// EnterIdent is called when production ident is entered.
func (s *BaseDelphiListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BaseDelphiListener) ExitIdent(ctx *IdentContext) {}

// EnterUsedKeywordsAsNames is called when production usedKeywordsAsNames is entered.
func (s *BaseDelphiListener) EnterUsedKeywordsAsNames(ctx *UsedKeywordsAsNamesContext) {}

// ExitUsedKeywordsAsNames is called when production usedKeywordsAsNames is exited.
func (s *BaseDelphiListener) ExitUsedKeywordsAsNames(ctx *UsedKeywordsAsNamesContext) {}

// EnterIdentList is called when production identList is entered.
func (s *BaseDelphiListener) EnterIdentList(ctx *IdentListContext) {}

// ExitIdentList is called when production identList is exited.
func (s *BaseDelphiListener) ExitIdentList(ctx *IdentListContext) {}

// EnterIdentListFlat is called when production identListFlat is entered.
func (s *BaseDelphiListener) EnterIdentListFlat(ctx *IdentListFlatContext) {}

// ExitIdentListFlat is called when production identListFlat is exited.
func (s *BaseDelphiListener) ExitIdentListFlat(ctx *IdentListFlatContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseDelphiListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseDelphiListener) ExitLabel(ctx *LabelContext) {}

// EnterIntNum is called when production intNum is entered.
func (s *BaseDelphiListener) EnterIntNum(ctx *IntNumContext) {}

// ExitIntNum is called when production intNum is exited.
func (s *BaseDelphiListener) ExitIntNum(ctx *IntNumContext) {}

// EnterRealNum is called when production realNum is entered.
func (s *BaseDelphiListener) EnterRealNum(ctx *RealNumContext) {}

// ExitRealNum is called when production realNum is exited.
func (s *BaseDelphiListener) ExitRealNum(ctx *RealNumContext) {}

// EnterNamespacedQualifiedIdent is called when production namespacedQualifiedIdent is entered.
func (s *BaseDelphiListener) EnterNamespacedQualifiedIdent(ctx *NamespacedQualifiedIdentContext) {}

// ExitNamespacedQualifiedIdent is called when production namespacedQualifiedIdent is exited.
func (s *BaseDelphiListener) ExitNamespacedQualifiedIdent(ctx *NamespacedQualifiedIdentContext) {}

// EnterNamespaceName is called when production namespaceName is entered.
func (s *BaseDelphiListener) EnterNamespaceName(ctx *NamespaceNameContext) {}

// ExitNamespaceName is called when production namespaceName is exited.
func (s *BaseDelphiListener) ExitNamespaceName(ctx *NamespaceNameContext) {}

// EnterQualifiedIdent is called when production qualifiedIdent is entered.
func (s *BaseDelphiListener) EnterQualifiedIdent(ctx *QualifiedIdentContext) {}

// ExitQualifiedIdent is called when production qualifiedIdent is exited.
func (s *BaseDelphiListener) ExitQualifiedIdent(ctx *QualifiedIdentContext) {}
