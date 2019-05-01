package iworkanalyzer

// 正则表达式
var regexs = []string{"^[a-zA-Z0-9]+\\(", "^\\)", "^`.*?`", "^[0-9]+", "^\\$[a-zA-Z_0-9](\\.[a-zA-Z0-9])+", "^\\+", ","}

// 正则表达式对应的词语
var regexLexers = []string{"func(", ")", "S", "N", "V", "+", ","}
