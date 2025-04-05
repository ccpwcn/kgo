package kgo

import (
	"bytes"
	"strings"
	"unicode"
)

type FormatOptions struct {
	UpperCaseKeywords bool   // 是否将关键字转为大写
	Indent            string // 缩进字符串
	LineWidth         int    // 最大行宽
}

// DefaultOptions 返回默认的格式化选项
func DefaultOptions() FormatOptions {
	return FormatOptions{
		UpperCaseKeywords: true,
		Indent:            "    ",
		LineWidth:         80,
	}
}

// SQLFormatOption 是一个修改 FormatOptions 的函数类型
type SQLFormatOption func(*FormatOptions)

// --- SQLFormatOption 设置函数 ---

// WithUpperCaseKeywords 设置是否将关键字大写
func WithUpperCaseKeywords(upper bool) SQLFormatOption {
	return func(opts *FormatOptions) {
		opts.UpperCaseKeywords = upper
	}
}

// WithIndent 设置缩进字符串
func WithIndent(indent string) SQLFormatOption {
	return func(opts *FormatOptions) {
		opts.Indent = indent
	}
}

// WithLineWidth 设置最大行宽
func WithLineWidth(width int) SQLFormatOption {
	return func(opts *FormatOptions) {
		opts.LineWidth = width
	}
}

// 主要关键字列表（用于换行和缩进）
var majorKeywords = map[string]bool{
	"SELECT": true, "FROM": true, "WHERE": true, "SET": true, "VALUES": true,
	"LEFT JOIN": true, "RIGHT JOIN": true, "INNER JOIN": true, "OUTER JOIN": true, "JOIN": true, "ON": true,
	"GROUP BY": true, "ORDER BY": true, "HAVING": true, "LIMIT": true, "OFFSET": true,
	"UNION": true, "UNION ALL": true, "INTERSECT": true, "EXCEPT": true, "WITH": true,
	"UPDATE": true, "INSERT INTO": true, "DELETE FROM": true, // INSERT INTO 和 DELETE FROM 作为复合词处理可能更好
	// 注意: AND/OR 单独处理，因为它们通常需要不同的缩进
}

// 复合关键字列表 (需要合并的)
var compoundKeywords = map[string]bool{
	"order by":    true,
	"group by":    true,
	"union all":   true,
	"left join":   true,
	"right join":  true,
	"inner join":  true,
	"outer join":  true,
	"insert into": true, // 新增
	"delete from": true, // 新增
}

// 需要加空格的运算符
var spacedOperators = map[string]bool{
	"=": true, ">": true, "<": true, ">=": true, "<=": true, "!=": true, "<>": true,
	"AND": true, "OR": true, "LIKE": true, "BETWEEN": true, "IS": true,
	"+": true, "-": true, "*": true, "/": true, "%": true, // 可选：添加算术运算符
}

// FormatSQL 格式化 SQL 语句，接受可选的配置项
func FormatSQL(sql string, optionSetters ...SQLFormatOption) string {
	// 1. 获取默认选项
	opts := DefaultOptions()

	// 2. 应用所有传入的 SQLFormatOption 函数来修改选项
	for _, setter := range optionSetters {
		setter(&opts)
	}

	// --- 接下来的逻辑使用最终的 opts 变量 ---
	sql = cleanSQL(sql)
	tokens := tokenizeSQL(sql)
	tokens = mergeCompoundTokens(tokens) // 合并复合关键字

	var buf bytes.Buffer
	indentLevel := 0
	currentLineLen := 0
	needsIndent := true
	parenLevel := 0

	for i, token := range tokens {
		upperToken := strings.ToUpper(token)

		// --- 换行逻辑 (使用 opts.LineWidth, opts.Indent) ---
		breakBefore := false
		if i > 0 {
			prevToken := tokens[i-1]
			if majorKeywords[upperToken] {
				if prevToken != "(" || upperToken == "SELECT" {
					breakBefore = true
				}
			}
			if upperToken == "AND" || upperToken == "OR" {
				breakBefore = true
			}

			estimatedLen := currentLineLen
			if needsSpace(prevToken, token) {
				estimatedLen++
			}
			estimatedLen += len(token)

			// 使用 opts.LineWidth
			if !breakBefore && opts.LineWidth > 0 && estimatedLen > opts.LineWidth && currentLineLen > indentLevel*len(opts.Indent) && prevToken != "(" && prevToken != "," {
				breakBefore = true
			}
		}

		if breakBefore {
			buf.WriteByte('\n')
			currentLineLen = 0
			// 使用 opts 计算缩进
			indentLevel = calculateIndent(parenLevel, upperToken, tokens, i)
			needsIndent = true
		}

		if needsIndent {
			// 使用 opts.Indent
			indentStr := strings.Repeat(opts.Indent, indentLevel)
			buf.WriteString(indentStr)
			currentLineLen = len(indentStr)
			needsIndent = false
		}

		// --- 处理空格 ---
		spaceNeeded := false
		if i > 0 {
			spaceNeeded = needsSpace(tokens[i-1], token)
		}
		if spaceNeeded {
			buf.WriteByte(' ')
			currentLineLen++
		}

		// --- 格式化 Token (使用 opts.UpperCaseKeywords) ---
		// 注意：formatToken 内部需要接收 opts 参数
		formattedToken := formatToken(token, opts, i, tokens) // 传递 opts

		// --- 写入 Token ---
		buf.WriteString(formattedToken)
		currentLineLen += len(formattedToken)

		// --- 更新括号层级 ---
		switch token {
		case "(":
			parenLevel++
		case ")":
			parenLevel--
			if parenLevel < 0 {
				parenLevel = 0
			}
		}
	}

	return strings.TrimSpace(buf.String())
}

// calculateIndent 计算缩进级别 (基于当前 Token 和括号层级)
func calculateIndent(parenLevel int, upperToken string, tokens []string, pos int) int {
	// 最终缩进基于括号层级和当前行首 token 类型调整
	finalIndent := parenLevel // 基础缩进等于括号层级

	// 如果行首是主要关键字 (FROM, WHERE, JOIN, GROUP BY, ORDER BY...) 或 AND/OR
	// 需要判断这是否真的是“行首” - 在当前实现中，这个函数是在决定换行后调用的，
	// 所以 upperToken 就是新行的第一个 token。
	if majorKeywords[upperToken] || upperToken == "AND" || upperToken == "OR" {
		// 在括号外 (parenLevel == 0)，这些关键字通常需要增加一级缩进（相对于上一行的 SELECT/UPDATE 等）
		if parenLevel == 0 {
			finalIndent = 1 // 主要子句缩进 1 级
			if upperToken == "AND" || upperToken == "OR" {
				finalIndent = 2 // AND/OR 再缩进 1 级 (总共 2 级)
			}
		} else {
			finalIndent = parenLevel + 1 // 子句相对括号缩进 1 级
			if upperToken == "AND" || upperToken == "OR" {
				finalIndent = parenLevel + 2 // AND/OR 再缩进 1 级
			}
			// 子查询的 SELECT 特殊处理?
			if upperToken == "SELECT" {
				finalIndent = parenLevel + 1 // 子查询 SELECT 相对括号缩进 1
			}
		}
	} else if upperToken == "VALUES" {
		// VALUES 关键字通常也需要缩进
		if parenLevel == 0 {
			finalIndent = 1
		} else {
			finalIndent = parenLevel + 1
		}
		// 如果 VALUES 后面紧跟 '('，缩进可能需要调整，但通常另起一行写括号和值
	}

	// 防止负缩进（理论上不应发生）
	if finalIndent < 0 {
		finalIndent = 0
	}

	return finalIndent
}

// needsSpace 判断 prevToken 和 currentToken 之间是否需要空格 (关键改进点)
func needsSpace(prevToken, currentToken string) bool {
	if prevToken == "" {
		return false // 第一个 Token 前没有空格
	}

	// 规则 1: 左括号后、右括号前、逗号前、分号前不需要空格
	if prevToken == "(" || currentToken == ")" || currentToken == "," || currentToken == ";" {
		return false
	}

	// 规则 2: 函数名后紧跟左括号不需要空格 (避免 COUNT (*))
	if currentToken == "(" && !isKeyword(prevToken) && !spacedOperators[strings.ToUpper(prevToken)] && prevToken != ")" && prevToken != "," {
		return false
	}

	// 规则 3: 点号两侧通常不需要空格 (如 table.column)
	if prevToken == "." || currentToken == "." {
		// 假设 tokenizer 能正确处理 table.column 或数字 1.23 为单个 token
		// 如果 tokenizer 分割了 t, ., c，则此规则阻止 t . c
		// 如果 tokenizer 产生 t.c, 此规则不生效
		return false
	}

	// 规则 4: 需要加空格的运算符两侧需要空格
	prevUpper := strings.ToUpper(prevToken)
	currentUpper := strings.ToUpper(currentToken)
	if spacedOperators[prevUpper] || spacedOperators[currentUpper] {
		return true
	}

	// 规则 5: 逗号后需要空格 (除非后面是右括号，由规则1处理)
	if prevToken == "," {
		return true
	}

	// --- 移除规则 6 ---
	// 不再需要对 SELECT * 进行特殊处理，让默认规则决定

	// 默认情况: 其他大部分情况都需要空格 (如关键字之间, 关键字和标识符之间)
	return true
}

// mergeCompoundTokens 合并复合关键字（如 ORDER BY）
func mergeCompoundTokens(tokens []string) []string {
	merged := make([]string, 0, len(tokens))
	i := 0
	for i < len(tokens) {
		// 优先检查三词复合（未来可能需要）

		// 检查两词复合
		if i < len(tokens)-1 {
			compound := strings.ToLower(tokens[i] + " " + tokens[i+1])
			if compoundKeywords[compound] {
				merged = append(merged, tokens[i]+" "+tokens[i+1])
				i += 2 // 跳过两个 Token
				continue
			}
		}
		// 非复合，添加单个 token
		merged = append(merged, tokens[i])
		i++
	}
	return merged
}

// formatToken 格式化 Token (关键字大写)，需要接收 FormatOptions
func formatToken(token string, opts FormatOptions, pos int, tokens []string) string {
	isComp := compoundKeywords[strings.ToLower(token)]
	isKw := isKeyword(token) || isComp

	// 使用传入的 opts
	if isKw && opts.UpperCaseKeywords {
		return strings.ToUpper(token)
	}
	return token
}

// cleanSQL 清理 SQL 语句 (保持不变)
func cleanSQL(sql string) string {
	sql = strings.TrimSpace(sql)
	// 替换多个空白为一个空格，但保留换行符用于可能的注释处理
	sql = strings.Join(strings.FieldsFunc(sql, func(r rune) bool {
		return unicode.IsSpace(r) && r != '\n' // 按非换行的空白分割
	}), " ")

	// 简单处理单行注释（保留）和多行注释（移除，或格式化）
	// 这是一个复杂的问题，简单的移除可能破坏SQL。暂时保持原样。
	/*
		var cleanedLines []string
		lines := strings.Split(sql, "\n")
		inMultiLineComment := false
		for _, line := range lines {
			trimmedLine := strings.TrimSpace(line)
			if strings.HasPrefix(trimmedLine, "--") {
				// 保留单行注释？或者移除？看需求。这里先移除。
				continue
			}
			if strings.Contains(trimmedLine, "/*") {
				// 移除多行注释？
				// 更复杂的处理逻辑...
			}
			cleanedLines = append(cleanedLines, line)
		}
		sql = strings.Join(cleanedLines, "\n") // 如果保留换行，用 \n 连接
		sql = strings.Join(strings.Fields(sql), " ") // 如果不保留，则彻底清理
	*/

	// 再次 Fields 确保移除因注释产生的多余空格
	sql = strings.Join(strings.Fields(sql), " ")

	return sql
}

// tokenizeSQL 分词器 (保持基本逻辑，但需要注意 '.' 的处理)
// 注意：当前 Tokenizer 会将 `table.column` 分割为 `table`, `.`, `column`
// 这需要在 needsSpace 中处理 `.` 两侧不加空格。
// 一个更健壮的 Tokenizer 会识别 `table.column` 为一个标识符 Token。
// 暂时使用现有 Tokenizer 并调整 needsSpace。
func tokenizeSQL(sql string) []string {
	var tokens []string
	var currentToken bytes.Buffer
	inString := false
	stringQuote := byte(0)

	for i := 0; i < len(sql); i++ {
		c := sql[i]

		switch {
		case inString:
			currentToken.WriteByte(c)
			if c == stringQuote {
				// 简单处理转义: '' or "" or ``
				if i+1 < len(sql) && sql[i+1] == stringQuote {
					currentToken.WriteByte(sql[i+1])
					i++ // 跳过下一个引号
				} else {
					inString = false
					tokens = append(tokens, currentToken.String())
					currentToken.Reset()
				}
			}

		case c == '\'' || c == '"' || c == '`':
			flushCurrentToken(&tokens, &currentToken)
			inString = true
			stringQuote = c
			currentToken.WriteByte(c)

		// 修改：将 '.' 视为可构成 Token 的一部分，而不是单独分隔符，除非它是唯一字符
		// case isSQLPunctuation(c) || isComparisonOperatorChar(c):
		case isSeparatorChar(c): // 使用新的判断函数
			// 检查是否为多字符运算符
			isMultiCharOp := false
			if (c == '>' || c == '<' || c == '!' || c == ':') && i+1 < len(sql) && sql[i+1] == '=' {
				flushCurrentToken(&tokens, &currentToken)
				tokens = append(tokens, string(c)+string(sql[i+1]))
				i++
				isMultiCharOp = true
			} else if c == '<' && i+1 < len(sql) && sql[i+1] == '>' { // 处理 <>
				flushCurrentToken(&tokens, &currentToken)
				tokens = append(tokens, "<>")
				i++
				isMultiCharOp = true
			}

			if !isMultiCharOp {
				flushCurrentToken(&tokens, &currentToken)
				tokens = append(tokens, string(c))
			}

		case unicode.IsSpace(rune(c)):
			flushCurrentToken(&tokens, &currentToken) // 空格分隔

		default: // 包括字母、数字、点号(.)、下划线(_)等
			currentToken.WriteByte(c)
		}
	}
	flushCurrentToken(&tokens, &currentToken)
	return tokens
}

// isSeparatorChar 判断是否为需要独立处理的分隔符或单字符运算符
func isSeparatorChar(c byte) bool {
	switch c {
	case '(', ')', ',', ';', '+', '-', '*', '/', '%', '=', '>', '<', '!', ':': // 冒号用于类型转换或特殊语法
		return true
	default:
		return false
	}
}

// flushCurrentToken (保持不变)
func flushCurrentToken(tokens *[]string, buf *bytes.Buffer) {
	if buf.Len() > 0 {
		*tokens = append(*tokens, buf.String())
		buf.Reset()
	}
}

// isComparisonOperatorChar (现在合并到 isSeparatorChar，或用于多字符判断)
func isComparisonOperatorChar(c byte) bool {
	return c == '=' || c == '>' || c == '<' || c == '!'
}

// isComparisonOperator 判断 Token 是否为比较运算符 (现在使用 spacedOperators)
func isComparisonOperator(token string) bool {
	_, exists := spacedOperators[strings.ToUpper(token)]
	// 仅限比较运算符
	switch strings.ToUpper(token) {
	case "=", ">", "<", ">=", "<=", "!=", "<>", "LIKE", "BETWEEN", "IS":
		return exists
	default:
		return false
	}
}

// isSQLPunctuation (现在合并到 isSeparatorChar)
func isSQLPunctuation(c byte) bool {
	return c == '(' || c == ')' || c == ',' || c == ';' // 只包含这些非运算符标点
}

// isPunctuation 判断 token 是否为标点符号 (仅用于单字符判断)
func isPunctuation(token string) bool {
	if len(token) != 1 {
		return false
	}
	return isSQLPunctuation(token[0])
}

// isKeyword 判断 token 是否为 SQL 关键字 (基础列表，可能需要补充)
// 注意：现在也包含了 spacedOperators 中的 AND/OR 等逻辑运算符
func isKeyword(token string) bool {
	lower := strings.ToLower(token)
	// 检查是否为复合关键字
	if compoundKeywords[lower] {
		return true
	}
	// 检查是否为主要关键字
	if majorKeywords[strings.ToUpper(token)] {
		return true
	}
	// 检查是否为需要空格的运算符（也视为广义的关键字）
	if spacedOperators[strings.ToUpper(token)] {
		return true
	}

	// 其他常用关键字
	switch lower {
	case "select", "from", "where", "and", "or", "not", "in", "like", "between",
		"is", "null", "join", "inner", "outer", "left", "right", "full", "on",
		"group", "by", "having", "order", "asc", "desc", "limit", "offset",
		"insert", "into", "values", "update", "set", "delete", "create", "table",
		"alter", "drop", "index", "primary", "key", "foreign", "references",
		"unique", "check", "default", "auto_increment", "as", "distinct",
		"count", "sum", "avg", "min", "max", "case", "when", "then", "else", "end",
		"exists", "any", "all", "union", "intersect", "except", "with":
		return true
	// 添加更多数据库特定关键字...
	default:
		return false
	}
}
