package slinguist

// Extracted from github/linguist.
// This one-liner prints the latest version:
// curl https://raw.githubusercontent.com/github/linguist/master/lib/linguist/languages.yml | python -c 'from collections import defaultdict; import os; import json; import yaml; l = yaml.load(os.sys.stdin.read()); exts = reduce(lambda acc, x: acc + x, [v.get("extensions", []) for k, v in l.items()], []); m = {ext: [k for k, v in l.items() if ext in v.get("extensions", [])] for ext in exts}; print "package utils\n\nvar languagesByExtension = map[string][]string{\n"+ json.dumps(m).replace("]", "}").replace("},", "},\n").replace("[", "{")[1:-1] + ",\n}\n"' | gofmt
var languagesByExtension = map[string][]string{
	".obj":                 {"Wavefront Object"},
	".sage":                {"Sage"},
	".tmux":                {"Shell"},
	".kt":                  {"Kotlin"},
	".builder":             {"Ruby"},
	".lfe":                 {"LFE"},
	".arpa":                {"DNS Zone"},
	".plsql":               {"PLSQL"},
	".nsh":                 {"NSIS"},
	".sh-session":          {"ShellSession"},
	".ma":                  {"Mathematica"},
	".raw":                 {"Raw token data"},
	".moon":                {"MoonScript"},
	".rl":                  {"Ragel in Ruby Host"},
	".ebuild":              {"Gentoo Ebuild"},
	".targets":             {"XML"},
	".lid":                 {"Dylan"},
	".rb":                  {"Ruby"},
	".iss":                 {"Inno Setup"},
	".pony":                {"Pony"},
	".thor":                {"Ruby"},
	".xml.dist":            {"XML"},
	".resx":                {"XML"},
	".bmx":                 {"BlitzMax"},
	".asd":                 {"Common Lisp"},
	".geojson":             {"JSON"},
	".bro":                 {"Bro"},
	".geo":                 {"GLSL"},
	".mkdn":                {"Markdown"},
	".irclog":              {"IRC log"},
	".rs.in":               {"Rust"},
	".rkt":                 {"Racket"},
	".bb":                  {"BitBake", "BlitzBasic"},
	".rd":                  {"R"},
	".mms":                 {"Module Management System"},
	".glslv":               {"GLSL"},
	".el":                  {"Emacs Lisp"},
	".nuspec":              {"XML"},
	".csx":                 {"C#"},
	".rss":                 {"XML"},
	".eh":                  {"eC"},
	".c++objdump":          {"Cpp-ObjDump"},
	".sublime-settings":    {"JavaScript"},
	".ec":                  {"eC"},
	".filters":             {"XML"},
	".csl":                 {"XML"},
	".self":                {"Self"},
	".pytb":                {"Python traceback"},
	".ex":                  {"Elixir"},
	".csd":                 {"Csound Document"},
	".xliff":               {"XML"},
	".pm6":                 {"Perl6"},
	".ant":                 {"XML"},
	".pde":                 {"Processing"},
	".pm":                  {"Perl6", "Perl"},
	".pl":                  {"Perl6", "Prolog", "Perl"},
	".rabl":                {"Ruby"},
	".f03":                 {"FORTRAN"},
	".ph":                  {"Perl"},
	".handlebars":          {"Handlebars"},
	".asmx":                {"ASP"},
	".chs":                 {"C2hs Haskell"},
	".pb":                  {"PureBasic"},
	".perl":                {"Perl"},
	".gtpl":                {"Groovy"},
	".py":                  {"Python"},
	".djs":                 {"Dogescript"},
	".xml":                 {"XML"},
	".xmi":                 {"XML"},
	".pp":                  {"Puppet", "Pascal"},
	".ps":                  {"PostScript"},
	".htm":                 {"HTML"},
	".gsx":                 {"Gosu"},
	".fxml":                {"XML"},
	".sagews":              {"Sage"},
	".sig":                 {"Standard ML"},
	".stTheme":             {"XML"},
	".vba":                 {"Visual Basic"},
	".gsp":                 {"Groovy Server Pages"},
	".tst":                 {"Scilab", "GAP"},
	".gst":                 {"Gosu"},
	".sparql":              {"SPARQL"},
	".lagda":               {"Literate Agda"},
	".ily":                 {"LilyPond"},
	".dita":                {"XML"},
	".watchr":              {"Ruby"},
	".vbs":                 {"Visual Basic"},
	".prg":                 {"xBase"},
	".ihlp":                {"Stata"},
	".pri":                 {"QMake"},
	".ecr":                 {"HTML+ECR"},
	".scrbl":               {"Racket"},
	".pro":                 {"INI", "Prolog", "QMake", "IDL"},
	".axd":                 {"ASP"},
	".txt":                 {"Text"},
	".minid":               {"MiniD"},
	".maxpat":              {"Max"},
	".hh":                  {"Hack", "C++"},
	".jinja":               {"HTML+Django"},
	".db2":                 {"SQLPL"},
	".xsjs":                {"JavaScript"},
	".xojo_report":         {"Xojo"},
	".p6":                  {"Perl6"},
	".escript":             {"Erlang"},
	".p6l":                 {"Perl6"},
	".p6m":                 {"Perl6"},
	".muf":                 {"MUF"},
	".viw":                 {"SQL"},
	".exs":                 {"Elixir"},
	".qml":                 {"QML"},
	".flex":                {"JFlex"},
	".sbt":                 {"Scala"},
	".tml":                 {"XML"},
	".nginxconf":           {"Nginx"},
	".xaml":                {"XML"},
	".zcml":                {"XML"},
	".urdf":                {"XML"},
	".hbs":                 {"Handlebars"},
	".nit":                 {"Nit"},
	".clp":                 {"CLIPS"},
	".wiki":                {"MediaWiki"},
	".reb":                 {"Rebol"},
	".ado":                 {"Stata"},
	".red":                 {"Red"},
	".adb":                 {"Ada"},
	".ada":                 {"Ada"},
	".bats":                {"Shell"},
	".xpy":                 {"Python"},
	".hxsl":                {"Haxe"},
	".f90":                 {"FORTRAN"},
	".ox":                  {"Ox"},
	".wsf":                 {"XML"},
	".oz":                  {"Oz"},
	".ads":                 {"Ada"},
	".xpl":                 {"XProc"},
	".adp":                 {"Tcl"},
	".graphql":             {"GraphQL"},
	".rktl":                {"Racket"},
	".3m":                  {"Groff"},
	".owl":                 {"Web Ontology Language"},
	".tmSnippet":           {"XML"},
	".vue":                 {"Vue"},
	".pir":                 {"Parrot Internal Representation"},
	".yaml":                {"YAML"},
	".axs":                 {"NetLinx"},
	".rst.txt":             {"reStructuredText"},
	".wsgi":                {"Python"},
	".hlsl":                {"HLSL"},
	".lmi":                 {"Python"},
	".command":             {"Shell"},
	".vert":                {"GLSL"},
	".objdump":             {"ObjDump"},
	".yml":                 {"YAML"},
	".fr":                  {"Text", "Frege", "Forth"},
	".pig":                 {"PigLatin"},
	".phtml":               {"HTML+PHP"},
	".ml4":                 {"OCaml"},
	".bas":                 {"Visual Basic"},
	".fy":                  {"Fancy"},
	".fx":                  {"HLSL", "FLUX"},
	".wxi":                 {"XML"},
	".wxl":                 {"XML"},
	".gap":                 {"GAP"},
	".aug":                 {"Augeas"},
	".ipynb":               {"Jupyter Notebook"},
	".diff":                {"Diff"},
	".wl":                  {"Mathematica"},
	".cljs.hl":             {"Clojure"},
	".wxs":                 {"XML"},
	".tsx":                 {"XML", "TypeScript"},
	".f08":                 {"FORTRAN"},
	".rst":                 {"reStructuredText"},
	".mll":                 {"OCaml"},
	".mli":                 {"OCaml"},
	".grace":               {"Grace"},
	".cljc":                {"Clojure"},
	".rsx":                 {"R"},
	".txl":                 {"TXL"},
	".cljx":                {"Clojure"},
	".sublime-syntax":      {"YAML"},
	".pmod":                {"Pike"},
	".cobol":               {"COBOL"},
	".mly":                 {"OCaml"},
	".opa":                 {"Opa"},
	".storyboard":          {"XML"},
	".cljs":                {"Clojure"},
	".feature":             {"Cucumber"},
	".rsh":                 {"RenderScript"},
	".arc":                 {"Arc"},
	".vshader":             {"GLSL"},
	".matah":               {"Stata"},
	".ts":                  {"XML", "TypeScript"},
	".xojo_window":         {"Xojo"},
	".tu":                  {"Turing"},
	".bsv":                 {"Bluespec"},
	".tm":                  {"Tcl"},
	".mo":                  {"Modelica"},
	".rhtml":               {"RHTML"},
	".xib":                 {"XML"},
	".tf":                  {"Terraform"},
	".clixml":              {"XML"},
	".smt":                 {"SMT"},
	".applescript":         {"AppleScript"},
	".rviz":                {"YAML"},
	".factor":              {"Factor"},
	".eclxml":              {"ECL"},
	".sma":                 {"SourcePawn"},
	".fun":                 {"Standard ML"},
	".sml":                 {"Standard ML"},
	".axml":                {"XML"},
	".reds":                {"Red"},
	".dae":                 {"COLLADA"},
	".spin":                {"Propeller Spin"},
	".bash":                {"Shell"},
	".tmCommand":           {"XML"},
	".cbx":                 {"TeX"},
	".numpyw":              {"NumPy"},
	".cmake.in":            {"CMake"},
	".intr":                {"Dylan"},
	".vhdl":                {"VHDL"},
	".textile":             {"Textile"},
	".ditaval":             {"XML"},
	".vhost":               {"Nginx", "ApacheConf"},
	".tac":                 {"Python"},
	".tab":                 {"SQL"},
	".pov":                 {"POV-Ray SDL"},
	".rno":                 {"Groff"},
	".sublime-snippet":     {"XML"},
	".haml.deface":         {"Haml"},
	".rest":                {"reStructuredText"},
	".sublime-commands":    {"JavaScript"},
	".cfml":                {"ColdFusion"},
	".cp":                  {"Component Pascal", "C++"},
	"._coffee":             {"CoffeeScript"},
	".rebol":               {"Rebol"},
	".cs":                  {"C#", "Smalltalk"},
	".ct":                  {"XML"},
	".cu":                  {"Cuda"},
	".cw":                  {"Redcode"},
	".ch":                  {"Charity", "xBase"},
	".wsdl":                {"XML"},
	".fan":                 {"Fantom"},
	".ck":                  {"ChucK"},
	".ahk":                 {"AutoHotkey"},
	".dpatch":              {"Darcs Patch"},
	".fcgi":                {"PHP", "Python", "Shell", "Perl", "Lua", "Ruby"},
	".cc":                  {"C++"},
	".yap":                 {"Prolog"},
	".fancypack":           {"Fancy"},
	".1m":                  {"Groff"},
	".sty":                 {"TeX"},
	".cshtml":              {"C#"},
	".bf":                  {"HyPhy", "Brainfuck"},
	".cl2":                 {"Clojure"},
	".mcr":                 {"MAXScript"},
	".rpy":                 {"Ren'Py", "Python"},
	".vark":                {"Gosu"},
	".osm":                 {"XML"},
	".upc":                 {"Unified Parallel C"},
	".xojo_code":           {"Xojo"},
	".anim":                {"Unity3D Asset"},
	".smt2":                {"SMT"},
	".gms":                 {"GAMS"},
	".hic":                 {"Clojure"},
	".raml":                {"RAML"},
	".mawk":                {"Awk"},
	".iced":                {"CoffeeScript"},
	".pkgproj":             {"XML"},
	".dylan":               {"Dylan"},
	".em":                  {"EmberScript"},
	".gml":                 {"XML", "Graph Modeling Language", "Game Maker Language"},
	".dotsettings":         {"XML"},
	".ahkl":                {"AutoHotkey"},
	".ps1":                 {"PowerShell"},
	".mdpolicy":            {"XML"},
	".sublime_metrics":     {"JavaScript"},
	".pwn":                 {"PAWN"},
	".ldml":                {"Lasso"},
	".i7x":                 {"Inform 7"},
	".clj":                 {"Clojure"},
	".robot":               {"RobotFramework"},
	".1":                   {"Groff"},
	".gvy":                 {"Groovy"},
	".xproc":               {"XProc"},
	".cls":                 {"Visual Basic", "Apex", "TeX", "OpenEdge ABL"},
	".unity":               {"Unity3D Asset"},
	".clw":                 {"Clarion"},
	".opencl":              {"OpenCL"},
	".blade":               {"Blade"},
	".csv":                 {"CSV"},
	".fish":                {"fish"},
	".idc":                 {"C"},
	".1x":                  {"Groff"},
	".ltx":                 {"TeX"},
	".pac":                 {"JavaScript"},
	".sch":                 {"Eagle", "KiCad"},
	".php3":                {"PHP"},
	".idr":                 {"Idris"},
	".mkfile":              {"Makefile"},
	".php5":                {"PHP"},
	".php4":                {"PHP"},
	".css":                 {"CSS"},
	".ditamap":             {"XML"},
	".pas":                 {"Pascal"},
	".grxml":               {"XML"},
	".jflex":               {"JFlex"},
	".bison":               {"Bison"},
	".a51":                 {"Assembly"},
	".mtl":                 {"Wavefront Material"},
	".metal":               {"Metal"},
	".opal":                {"Opal"},
	".mkdown":              {"Markdown"},
	".tcc":                 {"C++"},
	".csh":                 {"Tcsh"},
	".nproj":               {"XML"},
	".eps":                 {"PostScript"},
	".xul":                 {"XML"},
	".hb":                  {"Harbour"},
	".markdown":            {"Markdown"},
	".sublime_session":     {"JavaScript"},
	"._ls":                 {"LiveScript"},
	".lgt":                 {"Logtalk"},
	".phps":                {"PHP"},
	".hs":                  {"Haskell"},
	".phpt":                {"PHP"},
	".fxh":                 {"HLSL"},
	".epj":                 {"Ecere Projects"},
	".hy":                  {"Hy"},
	".hx":                  {"Haxe"},
	".JSON-tmLanguage":     {"JSON"},
	".prefab":              {"Unity3D Asset"},
	".tex":                 {"TeX"},
	".rbx":                 {"Ruby"},
	".5":                   {"Groff"},
	".pluginspec":          {"XML", "Ruby"},
	".yy":                  {"Yacc"},
	".rbw":                 {"Ruby"},
	".cl":                  {"OpenCL", "Cool", "Common Lisp"},
	".es":                  {"JavaScript", "Erlang"},
	".rest.txt":            {"reStructuredText"},
	".sublime-keymap":      {"JavaScript"},
	".tea":                 {"Tea"},
	".eq":                  {"EQ"},
	".f77":                 {"FORTRAN"},
	".zmpl":                {"Zimpl"},
	".mediawiki":           {"MediaWiki"},
	".man":                 {"Groff"},
	".4th":                 {"Forth"},
	".pan":                 {"Pan"},
	".d-objdump":           {"D-ObjDump"},
	".scxml":               {"XML"},
	".pd_lua":              {"Lua"},
	".god":                 {"Ruby"},
	".hats":                {"ATS"},
	".lvproj":              {"LabVIEW"},
	".pat":                 {"Max"},
	".purs":                {"PureScript"},
	".aux":                 {"TeX"},
	".cproject":            {"XML"},
	".pxd":                 {"Cython"},
	".creole":              {"Creole"},
	".moo":                 {"Mercury", "Moocode"},
	".hpp":                 {"C++"},
	".mod":                 {"Modula-2", "XML", "Linux Kernel Module", "AMPL"},
	".pxi":                 {"Cython"},
	".auk":                 {"Awk"},
	".befunge":             {"Befunge"},
	".bib":                 {"TeX"},
	".pd":                  {"Pure Data"},
	".jsb":                 {"JavaScript"},
	".mata":                {"Stata"},
	".3qt":                 {"Groff"},
	".xhtml":               {"HTML"},
	".mako":                {"Mako"},
	".mask":                {"Mask"},
	".eex":                 {"HTML+EEX"},
	".duby":                {"Mirah"},
	".coq":                 {"Coq"},
	".scss":                {"SCSS"},
	".emacs":               {"Emacs Lisp"},
	".lex":                 {"Lex"},
	".cob":                 {"COBOL"},
	".toml":                {"TOML"},
	".ll":                  {"LLVM"},
	".adoc":                {"AsciiDoc"},
	".ooc":                 {"ooc"},
	".json":                {"JSON"},
	".tpp":                 {"C++"},
	".au3":                 {"AutoIt"},
	".vcl":                 {"VCL"},
	".grt":                 {"Groovy"},
	".lua":                 {"Lua"},
	".vxml":                {"XML"},
	".6":                   {"Groff"},
	".tpl":                 {"Smarty"},
	".numpy":               {"NumPy"},
	".7":                   {"Groff"},
	".sh.in":               {"Shell"},
	".odd":                 {"XML"},
	".pyde":                {"Python"},
	".emberscript":         {"EmberScript"},
	".plist":               {"XML"},
	".sexp":                {"Common Lisp"},
	".pt":                  {"XML"},
	".toc":                 {"World of Warcraft Addon Data", "TeX"},
	".pod":                 {"Perl", "Pod"},
	".ni":                  {"Inform 7"},
	".mathematica":         {"Mathematica"},
	".ruby":                {"Ruby"},
	".sthlp":               {"Stata"},
	".g4":                  {"ANTLR"},
	".dyl":                 {"Dylan"},
	".pot":                 {"Gettext Catalog"},
	".coffee":              {"CoffeeScript"},
	".nim":                 {"Nimrod"},
	".cirru":               {"Cirru"},
	".kicad_pcb":           {"KiCad"},
	".groovy":              {"Groovy"},
	".nse":                 {"Lua"},
	".ron":                 {"Markdown"},
	".ls":                  {"LiveScript", "LoomScript"},
	".lasso9":              {"Lasso"},
	".lasso8":              {"Lasso"},
	".xqy":                 {"XQuery"},
	".asciidoc":            {"AsciiDoc"},
	".sublime-menu":        {"JavaScript"},
	".ecl":                 {"ECLiPSe", "ECL"},
	".E":                   {"E"},
	".myt":                 {"Myghty"},
	".agc":                 {"Apollo Guidance Computer"},
	".xtend":               {"Xtend"},
	".dpr":                 {"Pascal"},
	".latte":               {"Latte"},
	".xqm":                 {"XQuery"},
	".xql":                 {"XQuery"},
	".ivy":                 {"XML"},
	".gf":                  {"Grammatical Framework"},
	".numsc":               {"NumPy"},
	".shader":              {"GLSL"},
	".hcl":                 {"HCL"},
	".go":                  {"Go"},
	".gi":                  {"GAP"},
	".yang":                {"YANG"},
	".gv":                  {"Graphviz (DOT)"},
	".sfproj":              {"XML"},
	".gp":                  {"Gnuplot"},
	".njs":                 {"JavaScript"},
	".gs":                  {"JavaScript", "Gosu"},
	".lslp":                {"LSL"},
	".com":                 {"DIGITAL Command Language"},
	".cy":                  {"Cycript"},
	".scpt":                {"AppleScript"},
	".rg":                  {"Rouge"},
	".boot":                {"Clojure"},
	".lasso":               {"Lasso"},
	".rs":                  {"RenderScript", "Rust"},
	".flux":                {"FLUX"},
	".slim":                {"Slim"},
	".pike":                {"Pike"},
	".ru":                  {"Ruby"},
	".rbuild":              {"Ruby"},
	".wlt":                 {"Mathematica"},
	".1in":                 {"Groff"},
	".hqf":                 {"SQF"},
	".dats":                {"ATS"},
	".asset":               {"Unity3D Asset"},
	".hlsli":               {"HLSL"},
	".zpl":                 {"Zimpl"},
	".prc":                 {"SQL"},
	".for":                 {"Forth", "Formatted", "FORTRAN"},
	".mkd":                 {"Markdown"},
	".ccp":                 {"COBOL"},
	".cr":                  {"Crystal"},
	".ijs":                 {"J"},
	".emacs.desktop":       {"Emacs Lisp"},
	".xojo_toolbar":        {"Xojo"},
	".x10":                 {"X10"},
	".ttl":                 {"Turtle"},
	".doh":                 {"Stata"},
	".ux":                  {"XML"},
	".gawk":                {"Awk"},
	".twig":                {"Twig"},
	".ur":                  {"UrWeb"},
	".mumps":               {"M"},
	".darcspatch":          {"Darcs Patch"},
	".bbx":                 {"TeX"},
	".ui":                  {"XML"},
	".r3":                  {"Rebol"},
	".r2":                  {"Rebol"},
	".ksh":                 {"Shell"},
	".ccxml":               {"XML"},
	".dot":                 {"Graphviz (DOT)"},
	".xproj":               {"XML"},
	".sls":                 {"Scheme", "SaltStack"},
	".vbproj":              {"XML"},
	".sql":                 {"PLpgSQL", "PLSQL", "SQL", "SQLPL"},
	".hxx":                 {"C++"},
	".thrift":              {"Thrift"},
	".yaml-tmlanguage":     {"YAML"},
	".ktm":                 {"Kotlin"},
	".dfm":                 {"Pascal"},
	".litcoffee":           {"Literate CoffeeScript"},
	".aj":                  {"AspectJ"},
	".kts":                 {"Kotlin"},
	".sld":                 {"Scheme"},
	".icl":                 {"Clean"},
	".topojson":            {"JSON"},
	".cjsx":                {"CoffeeScript"},
	".cdf":                 {"Mathematica"},
	".ik":                  {"Ioke"},
	".xrl":                 {"Erlang"},
	".io":                  {"Io"},
	".f":                   {"Forth", "FORTRAN"},
	".sqf":                 {"SQF"},
	".cmd":                 {"Batchfile"},
	".dyalog":              {"APL"},
	".lisp":                {"NewLisp", "Common Lisp"},
	".sublime-mousemap":    {"JavaScript"},
	".al":                  {"Perl"},
	".bones":               {"JavaScript"},
	".jscad":               {"JavaScript"},
	".vala":                {"Vala"},
	".druby":               {"Mirah"},
	".psc":                 {"Papyrus"},
	".dlm":                 {"IDL"},
	".cxx-objdump":         {"Cpp-ObjDump"},
	".xojo_menu":           {"Xojo"},
	".pkb":                 {"PLSQL"},
	".lookml":              {"LookML"},
	".boo":                 {"Boo"},
	".h++":                 {"C++"},
	".pasm":                {"Parrot Assembly"},
	".dart":                {"Dart"},
	".cxx":                 {"C++"},
	".ceylon":              {"Ceylon"},
	".pks":                 {"PLSQL"},
	".bat":                 {"Batchfile"},
	".desktop.in":          {"desktop"},
	".srt":                 {"SRecode Template", "SubRip Text"},
	".c++-objdump":         {"Cpp-ObjDump"},
	".lol":                 {"LOLCODE"},
	".dtx":                 {"TeX"},
	".gco":                 {"G-code"},
	".es6":                 {"JavaScript"},
	".pbi":                 {"PureBasic"},
	".qbs":                 {"QML"},
	".cats":                {"C"},
	".psm1":                {"PowerShell"},
	".rbuistate":           {"REALbasic"},
	".tmLanguage":          {"XML"},
	".glf":                 {"Glyph"},
	".irbrc":               {"Ruby"},
	".tmPreferences":       {"XML"},
	".tla":                 {"TLA"},
	".vrx":                 {"GLSL"},
	".rake":                {"Ruby"},
	".omgrofl":             {"Omgrofl"},
	".org":                 {"Org"},
	".tmTheme":             {"XML"},
	".cql":                 {"SQL"},
	".xslt":                {"XSLT"},
	".apl":                 {"APL"},
	".lpr":                 {"Pascal"},
	".vapi":                {"Vala"},
	".wlua":                {"Lua"},
	".uc":                  {"UnrealScript"},
	".edn":                 {"edn"},
	".hrl":                 {"Erlang"},
	".maxproj":             {"Max"},
	".builds":              {"XML"},
	".veo":                 {"Verilog"},
	".nbp":                 {"Mathematica"},
	".vh":                  {"SystemVerilog"},
	".udo":                 {"Csound"},
	".ins":                 {"TeX"},
	".pogo":                {"PogoScript"},
	".udf":                 {"SQL"},
	".psd1":                {"PowerShell"},
	".bzl":                 {"Python"},
	".glade":               {"XML"},
	".cppobjdump":          {"Cpp-ObjDump"},
	".inc":                 {"PHP", "HTML", "SourcePawn", "Assembly", "PAWN", "SQL", "Pascal", "POV-Ray SDL", "C++"},
	".ini":                 {"INI"},
	".vhd":                 {"VHDL"},
	".webidl":              {"WebIDL"},
	".sublime-project":     {"JavaScript"},
	".inl":                 {"C++"},
	".ino":                 {"Arduino"},
	".frg":                 {"GLSL"},
	".zsh":                 {"Shell"},
	".decls":               {"BlitzBasic"},
	".xacro":               {"XML"},
	".frm":                 {"Visual Basic"},
	".ddl":                 {"SQL"},
	".cson":                {"CoffeeScript"},
	".cake":                {"CoffeeScript", "C#"},
	".ssjs":                {"JavaScript"},
	".frt":                 {"Forth"},
	".dcl":                 {"Clean"},
	".rbres":               {"REALbasic"},
	".sce":                 {"Scilab"},
	".ascx":                {"ASP"},
	".frx":                 {"Visual Basic"},
	".tcl":                 {"Tcl"},
	".agda":                {"Agda"},
	".yrl":                 {"Erlang"},
	".eclass":              {"Gentoo Eclass"},
	".3in":                 {"Groff"},
	".mustache":            {"HTML+Django"},
	".fpp":                 {"FORTRAN"},
	".geom":                {"GLSL"},
	".fsproj":              {"XML"},
	".json5":               {"JSON5"},
	".scaml":               {"Scaml"},
	".c++":                 {"C++"},
	".jade":                {"Jade"},
	".gradle":              {"Gradle"},
	".mt":                  {"Mathematica"},
	".mu":                  {"mupad"},
	".ms":                  {"MAXScript", "GAS", "Groff"},
	".f95":                 {"FORTRAN"},
	".cpp-objdump":         {"Cpp-ObjDump"},
	".rbbas":               {"REALbasic"},
	".3x":                  {"Groff"},
	".md":                  {"GCC Machine Description", "Markdown"},
	".me":                  {"Groff"},
	".mxt":                 {"Max"},
	".apacheconf":          {"ApacheConf"},
	".ml":                  {"OCaml"},
	".mm":                  {"Objective-C++", "XML"},
	".mk":                  {"Makefile"},
	".nawk":                {"Awk"},
	".axi.erb":             {"NetLinx+ERB"},
	".di":                  {"D"},
	".orc":                 {"Csound"},
	".podspec":             {"Ruby"},
	".dm":                  {"DM"},
	".cmake":               {"CMake"},
	".do":                  {"Stata"},
	".zep":                 {"Zephir"},
	".chpl":                {"Chapel"},
	".svg":                 {"SVG"},
	".props":               {"XML"},
	".podsl":               {"Common Lisp"},
	".roff":                {"Groff"},
	".svh":                 {"SystemVerilog"},
	".ampl":                {"AMPL"},
	".m4":                  {"M4", "M4Sugar"},
	".oxygene":             {"Oxygene"},
	".jbuilder":            {"Ruby"},
	".jsx":                 {"JSX"},
	".cuh":                 {"Cuda"},
	".mat":                 {"Unity3D Asset"},
	".jsp":                 {"Java Server Pages"},
	".jss":                 {"JavaScript"},
	".jsm":                 {"JavaScript"},
	".jsproj":              {"XML"},
	".mak":                 {"Makefile"},
	".parrot":              {"Parrot"},
	".properties":          {"INI"},
	".mao":                 {"Mako"},
	".hlean":               {"Lean"},
	".ML":                  {"Standard ML"},
	".jl":                  {"Julia"},
	".prw":                 {"xBase"},
	".stan":                {"Stan"},
	".gcode":               {"G-code"},
	".plot":                {"Gnuplot"},
	".lsl":                 {"LSL"},
	".matlab":              {"Matlab"},
	".axi":                 {"NetLinx"},
	".pyp":                 {"Python"},
	".zimpl":               {"Zimpl"},
	".pyt":                 {"Python"},
	".pyw":                 {"Python"},
	".lsp":                 {"NewLisp", "Common Lisp"},
	".pyx":                 {"Cython"},
	".psgi":                {"Perl"},
	".dockerfile":          {"Dockerfile"},
	".tcsh":                {"Tcsh"},
	".xojo_script":         {"Xojo"},
	".lean":                {"Lean"},
	".prolog":              {"Prolog"},
	".capnp":               {"Cap'n Proto"},
	".wisp":                {"wisp"},
	".jake":                {"JavaScript"},
	".golo":                {"Golo"},
	".nlogo":               {"NetLogo"},
	".rbmnu":               {"REALbasic"},
	".krl":                 {"KRL"},
	".nsi":                 {"NSIS"},
	".vhs":                 {"VHDL"},
	".vht":                 {"VHDL"},
	".rbfrm":               {"REALbasic"},
	".vhw":                 {"VHDL"},
	".forth":               {"Forth"},
	".sublime-theme":       {"JavaScript"},
	".smali":               {"Smali"},
	".scm":                 {"Scheme"},
	".sco":                 {"Csound Score"},
	".pl6":                 {"Perl6"},
	".sci":                 {"Scilab"},
	".vhf":                 {"VHDL"},
	".apib":                {"API Blueprint"},
	".scd":                 {"SuperCollider"},
	".vhi":                 {"VHDL"},
	".nut":                 {"Squirrel"},
	".ninja":               {"Ninja"},
	".pkl":                 {"Pickle"},
	".vho":                 {"VHDL"},
	".syntax":              {"YAML"},
	".swift":               {"Swift"},
	".rbxs":                {"Lua"},
	".las":                 {"Lasso"},
	".axs.erb":             {"NetLinx+ERB"},
	".js":                  {"JavaScript"},
	".jq":                  {"JSONiq"},
	".click":               {"Click"},
	".sublime-build":       {"JavaScript"},
	".aspx":                {"ASP"},
	".plb":                 {"PLSQL"},
	".mss":                 {"CartoCSS"},
	".sublime-workspace":   {"JavaScript"},
	".ipf":                 {"IGOR Pro"},
	".mspec":               {"Ruby"},
	".gshader":             {"GLSL"},
	".pls":                 {"PLSQL"},
	".http":                {"HTTP"},
	".plt":                 {"Gnuplot"},
	".c-objdump":           {"C-ObjDump"},
	".eliom":               {"OCaml"},
	".lhs":                 {"Literate Haskell"},
	".plx":                 {"Perl"},
	".vb":                  {"Visual Basic"},
	".rktd":                {"Racket"},
	".yacc":                {"Yacc"},
	".pck":                 {"PLSQL"},
	".logtalk":             {"Logtalk"},
	".eliomi":              {"OCaml"},
	".cpp":                 {"C++"},
	".cps":                 {"Component Pascal"},
	".mirah":               {"Mirah"},
	".zone":                {"DNS Zone"},
	".cpy":                 {"COBOL"},
	".erb.deface":          {"HTML+ERB"},
	".liquid":              {"Liquid"},
	".asp":                 {"ASP"},
	".jsonld":              {"JSONLD"},
	".xq":                  {"XQuery"},
	".proto":               {"Protocol Buffer"},
	".mmk":                 {"Module Management System"},
	".ps1xml":              {"XML"},
	".fshader":             {"GLSL"},
	".asc":                 {"Public Key", "AsciiDoc", "AGS Script"},
	".rq":                  {"SPARQL"},
	".xc":                  {"XC"},
	".xm":                  {"Logos"},
	".vsh":                 {"GLSL"},
	".ash":                 {"AGS Script"},
	".xi":                  {"Logos"},
	".csproj":              {"XML"},
	".asm":                 {"Assembly"},
	".sh":                  {"Shell"},
	".sj":                  {"Objective-J"},
	".sl":                  {"Slash"},
	".dll.config":          {"XML"},
	".volt":                {"Volt"},
	".rdoc":                {"RDoc"},
	".sc":                  {"SuperCollider", "Scala"},
	".xlf":                 {"XML"},
	".sp":                  {"SourcePawn"},
	".tool":                {"Shell"},
	".urs":                 {"UrWeb"},
	".ss":                  {"Scheme"},
	".st":                  {"HTML", "Smalltalk"},
	".sv":                  {"SystemVerilog"},
	".cbl":                 {"COBOL"},
	".iml":                 {"XML"},
	".sps":                 {"Scheme"},
	".sats":                {"ATS"},
	".frag":                {"GLSL", "JavaScript"},
	".kit":                 {"Kit"},
	".pug":                 {"Jade"},
	".sjs":                 {"JavaScript"},
	".pub":                 {"Public Key"},
	".html":                {"HTML"},
	".xsp.metadata":        {"XPages"},
	".eam.fs":              {"Formatted"},
	".kid":                 {"Genshi"},
	".ld":                  {"Linker Script"},
	".haml":                {"Haml"},
	".styl":                {"Stylus"},
	".po":                  {"Gettext Catalog"},
	".fsi":                 {"F#"},
	".fsh":                 {"GLSL"},
	".elm":                 {"Elm"},
	".nix":                 {"Nix"},
	".lbx":                 {"TeX"},
	".html.hl":             {"HTML"},
	".gyp":                 {"Python"},
	".hsc":                 {"Haskell"},
	".pascal":              {"Pascal"},
	".fsx":                 {"F#"},
	".shen":                {"Shen"},
	".maxhelp":             {"Max"},
	".nqp":                 {"Perl6"},
	".6pl":                 {"Perl6"},
	".6pm":                 {"Perl6"},
	".sass":                {"Sass"},
	".meta":                {"Unity3D Asset"},
	".brs":                 {"Brightscript"},
	".mkiv":                {"TeX"},
	".vcxproj":             {"XML"},
	".rmd":                 {"RMarkdown"},
	".8":                   {"Groff"},
	".9":                   {"Groff"},
	".scala":               {"Scala"},
	".jsfl":                {"JavaScript"},
	".mkii":                {"TeX"},
	".fs":                  {"GLSL", "F#", "Forth", "Filterscript"},
	".2":                   {"Groff"},
	".3":                   {"Groff"},
	".4":                   {"Groff"},
	".scad":                {"OpenSCAD"},
	".brd":                 {"Eagle", "KiCad"},
	".glsl":                {"GLSL"},
	".no":                  {"Text"},
	".gd":                  {"GAP", "GDScript"},
	".nl":                  {"NewLisp", "NL"},
	".rdf":                 {"XML"},
	".xsl":                 {"XSLT"},
	"._js":                 {"JavaScript"},
	".fp":                  {"GLSL"},
	".nc":                  {"nesC"},
	".nb":                  {"Text", "Mathematica"},
	".gemspec":             {"Ruby"},
	".xsd":                 {"XML"},
	".less":                {"Less"},
	".erl":                 {"Erlang"},
	".ny":                  {"Common Lisp"},
	".cgi":                 {"Python", "Shell", "Perl"},
	".nu":                  {"Nu"},
	".erb":                 {"HTML+ERB"},
	".h":                   {"C", "Objective-C", "C++"},
	".cljscm":              {"Clojure"},
	".j":                   {"Objective-J", "Jasmin"},
	".vssettings":          {"XML"},
	".l":                   {"PicoLisp", "Lex", "Groff", "Common Lisp"},
	".m":                   {"Mercury", "Limbo", "Objective-C", "MUF", "Mathematica", "Matlab", "M"},
	".php":                 {"PHP", "Hack"},
	".gnuplot":             {"Gnuplot"},
	".psc1":                {"XML"},
	".oxh":                 {"Ox"},
	".b":                   {"Limbo", "Brainfuck"},
	".c":                   {"C"},
	".d":                   {"DTrace", "Makefile", "D"},
	".e":                   {"Eiffel"},
	".oxo":                 {"Ox"},
	".g":                   {"G-code", "GAP"},
	".x":                   {"Logos"},
	".y":                   {"Yacc"},
	".aw":                  {"PHP"},
	".mxml":                {"XML"},
	".p":                   {"OpenEdge ABL"},
	".nasm":                {"Assembly"},
	".r":                   {"Rebol", "R"},
	".s":                   {"GAS"},
	".t":                   {"Perl6", "Terra", "Turing", "Perl"},
	".ston":                {"STON"},
	".v":                   {"Verilog", "Coq"},
	".w":                   {"C"},
	".desktop":             {"desktop"},
	".ejs":                 {"EJS"},
	".gnu":                 {"Gnuplot"},
	".awk":                 {"Awk"},
	".x3d":                 {"XML"},
	".vim":                 {"VimL"},
	".sublime-completions": {"JavaScript"},
	".asax":                {"ASP"},
	".ctp":                 {"PHP"},
	".nimrod":              {"Nimrod"},
	".als":                 {"Alloy"},
	".mir":                 {"Mirah"},
	".lidr":                {"Idris"},
	".xquery":              {"XQuery"},
	".prefs":               {"INI"},
	".app.src":             {"Erlang"},
	".monkey":              {"Monkey"},
	".srdf":                {"XML"},
	".jelly":               {"XML"},
	".n":                   {"Nemerle", "Groff"},
	".sublime-macro":       {"JavaScript"},
	".kml":                 {"XML"},
	".thy":                 {"Isabelle"},
	".ncl":                 {"Text", "NCL"},
	".weechatlog":          {"IRC log"},
	".ashx":                {"ASP"},
	".xsjslib":             {"JavaScript"},
	".xht":                 {"HTML"},
	".blade.php":           {"Blade"},
	".uno":                 {"Uno"},
	".ly":                  {"LilyPond"},
	".patch":               {"Diff"},
	".launch":              {"XML"},
	".mkvi":                {"TeX"},
	".xsp-config":          {"XPages"},
	".mtml":                {"MTML"},
	".xs":                  {"XS"},
	".lds":                 {"Linker Script"},
	".reek":                {"YAML"},
	".rbtbar":              {"REALbasic"},
	".abap":                {"ABAP"},
	".cfc":                 {"ColdFusion CFC"},
	".ipp":                 {"C++"},
	".vbhtml":              {"Visual Basic"},
	".cfg":                 {"INI"},
	".sas":                 {"SAS"},
	".ftl":                 {"FreeMarker"},
	".java":                {"Java"},
	".fth":                 {"Forth"},
	".cfm":                 {"ColdFusion"},
	".as":                  {"ActionScript"},
}
