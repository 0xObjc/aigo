package config

var defaultExcludeRules = map[string][]string{
	"Go": {
		"go.sum", "*.test", "*.coverprofile", "vendor/",
	},
	"Python": {
		"__pycache__", "*.pyc", "*.pyo", "*.pyd", ".pytest_cache/", ".coverage",
	},
	"Java": {
		"target/", "*.class", "*.jar",
	},
	"C++": {
		"build/", "*.o", "*.obj", "*.exe",
	},
	"JavaScript": {
		"node_modules/", "*.log", "*.lock", "dist/",
	},
	"Ruby": {
		"vendor/", "*.gem", ".bundle/",
	},
	"PHP": {
		"vendor/", "*.log",
	},
	"C#": {
		"bin/", "obj/", "*.dll", "*.exe",
	},
	"Swift": {
		"DerivedData/", "*.xcodeproj/xcuserdata", "*.xcworkspace/xcuserdata",
	},
	"Kotlin": {
		"build/", "*.kt.bak",
	},
	"Rust": {
		"target/", "Cargo.lock",
	},
	"Scala": {
		"target/", "*.sbt.bak",
	},
	"Haskell": {
		"dist/", "*.hi", "*.o",
	},
	"Clojure": {
		"target/", "*.clj.bak",
	},
	"Erlang": {
		"_build/", "*.beam",
	},
	"Elixir": {
		"_build/", "deps/", "*.exs.bak",
	},
	"Elm": {
		"elm-stuff/", "*.elm.bak",
	},
	"Dart": {
		"build/", "*.dart.bak",
	},
	"Shell": {
		"*.log", "*.bak",
	},
	"Perl": {
		"*.pl.bak",
	},
	"Lua": {
		"*.lua.bak",
	},
	"R": {
		"*.RData", "*.Rhistory",
	},
	"MATLAB": {
		"*.mat", "*.m.bak",
	},
	"Groovy": {
		"build/", "*.groovy.bak",
	},
	"Fortran": {
		"*.mod", "*.f90.bak",
	},
	"Visual Basic": {
		"bin/", "obj/", "*.vb.bak",
	},
	"Assembly": {
		"*.asm.bak",
	},
	"SQL": {
		"*.sql.bak",
	},
	"HTML": {
		"*.html.bak",
	},
	"CSS": {
		"*.css.bak",
	},
	"XML": {
		"*.xml.bak",
	},
	"JSON": {
		"*.json.bak",
	},
	"YAML": {
		"*.yaml.bak",
	},
	"TOML": {
		"*.toml.bak",
	},
	"Markdown": {
		"*.md.bak",
	},
	"Text": {
		"*.txt.bak",
	},
	// 添加其他语言的默认规则
}
