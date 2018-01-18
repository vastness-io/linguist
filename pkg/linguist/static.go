// DO NOT EDIT ** This file was generated with the bake tool ** DO NOT EDIT //

package linguist

var files = map[string]string{
	"data/languages.yml": `# Defines all Languages known to GitHub.
#
# type              - Either data, programming, markup, prose, or nil
# aliases           - An Array of additional aliases (implicitly
#                     includes name.downcase)
# ace_mode          - A String name of the Ace Mode used for highlighting whenever
#                     a file is edited. This must match one of the filenames in http://git.io/3XO_Cg.
#                     Use "text" if a mode does not exist.
# codemirror_mode   - A String name of the CodeMirror Mode used for highlighting whenever a file is edited.
#                     This must match a mode from https://git.io/vi9Fx
# wrap              - Boolean wrap to enable line wrapping (default: false)
# extensions        - An Array of associated extensions (the first one is
#                     considered the primary extension, the others should be
#                     listed alphabetically)
# interpreters      - An Array of associated interpreters
# searchable        - Boolean flag to enable searching (defaults to true)
# language_id       - Integer used as a language-name-independent indexed field so that we can rename
#                     languages in Linguist without reindexing all the code on GitHub. Must not be
#                     changed for existing languages without the explicit permission of GitHub staff.
# color             - CSS hex color to represent the language. Only used if type is "programming" or "prose".
# tm_scope          - The TextMate scope that represents this programming
#                     language. This should match one of the scopes listed in
#                     the grammars.yml file. Use "none" if there is no grammar
#                     for this language.
# group             - Name of the parent language. Languages in a group are counted
#                     in the statistics as the parent language.
#
# Any additions or modifications (even trivial) should have corresponding
# test changes in ` + "`" + `test/test_blob.rb` + "`" + `.
#
# Please keep this list alphabetized. Capitalization comes before lowercase.

---
1C Enterprise:
  type: programming
  color: "#814CCC"
  extensions:
  - ".bsl"
  - ".os"
  tm_scope: source.bsl
  ace_mode: text
  language_id: 0
ABAP:
  type: programming
  color: "#E8274B"
  extensions:
  - ".abap"
  ace_mode: abap
  language_id: 1
ABNF:
  type: data
  ace_mode: text
  extensions:
  - ".abnf"
  tm_scope: source.abnf
  language_id: 429
AGS Script:
  type: programming
  color: "#B9D9FF"
  aliases:
  - ags
  extensions:
  - ".asc"
  - ".ash"
  tm_scope: source.c++
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-c++src
  language_id: 2
AMPL:
  type: programming
  color: "#E6EFBB"
  extensions:
  - ".ampl"
  - ".mod"
  tm_scope: source.ampl
  ace_mode: text
  language_id: 3
ANTLR:
  type: programming
  color: "#9DC3FF"
  extensions:
  - ".g4"
  ace_mode: text
  language_id: 4
API Blueprint:
  type: markup
  color: "#2ACCA8"
  ace_mode: markdown
  extensions:
  - ".apib"
  tm_scope: text.html.markdown.source.gfm.apib
  language_id: 5
APL:
  type: programming
  color: "#5A8164"
  extensions:
  - ".apl"
  - ".dyalog"
  interpreters:
  - apl
  - aplx
  - dyalog
  tm_scope: source.apl
  ace_mode: text
  codemirror_mode: apl
  codemirror_mime_type: text/apl
  language_id: 6
ASN.1:
  type: data
  extensions:
  - ".asn"
  - ".asn1"
  tm_scope: source.asn
  ace_mode: text
  codemirror_mode: asn.1
  codemirror_mime_type: text/x-ttcn-asn
  language_id: 7
ASP:
  type: programming
  color: "#6a40fd"
  tm_scope: text.html.asp
  aliases:
  - aspx
  - aspx-vb
  extensions:
  - ".asp"
  - ".asax"
  - ".ascx"
  - ".ashx"
  - ".asmx"
  - ".aspx"
  - ".axd"
  ace_mode: text
  codemirror_mode: htmlembedded
  codemirror_mime_type: application/x-aspx
  language_id: 8
ATS:
  type: programming
  color: "#1ac620"
  aliases:
  - ats2
  extensions:
  - ".dats"
  - ".hats"
  - ".sats"
  tm_scope: source.ats
  ace_mode: ocaml
  language_id: 9
ActionScript:
  type: programming
  tm_scope: source.actionscript.3
  color: "#882B0F"
  aliases:
  - actionscript 3
  - actionscript3
  - as3
  extensions:
  - ".as"
  ace_mode: actionscript
  language_id: 10
Ada:
  type: programming
  color: "#02f88c"
  extensions:
  - ".adb"
  - ".ada"
  - ".ads"
  aliases:
  - ada95
  - ada2005
  ace_mode: ada
  language_id: 11
Adobe Font Metrics:
  type: data
  tm_scope: source.afm
  extensions:
  - ".afm"
  aliases:
  - acfm
  - adobe composite font metrics
  - adobe multiple font metrics
  - amfm
  ace_mode: text
  language_id: 147198098
Agda:
  type: programming
  color: "#315665"
  extensions:
  - ".agda"
  ace_mode: text
  language_id: 12
Alloy:
  type: programming
  color: "#64C800"
  extensions:
  - ".als"
  ace_mode: text
  language_id: 13
Alpine Abuild:
  type: programming
  group: Shell
  aliases:
  - abuild
  - apkbuild
  filenames:
  - APKBUILD
  tm_scope: source.shell
  ace_mode: sh
  codemirror_mode: shell
  codemirror_mime_type: text/x-sh
  language_id: 14
AngelScript:
  type: programming
  color: "#C7D7DC"
  extensions:
  - ".as"
  - ".angelscript"
  tm_scope: source.angelscript
  ace_mode: text
  codemirror_mode: clike
  codemirror_mime_type: text/x-c++src
  language_id: 389477596
Ant Build System:
  type: data
  tm_scope: text.xml.ant
  filenames:
  - ant.xml
  - build.xml
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: application/xml
  language_id: 15
ApacheConf:
  type: data
  aliases:
  - aconf
  - apache
  extensions:
  - ".apacheconf"
  - ".vhost"
  tm_scope: source.apache-config
  ace_mode: apache_conf
  language_id: 16
Apex:
  type: programming
  extensions:
  - ".cls"
  tm_scope: source.java
  ace_mode: java
  codemirror_mode: clike
  codemirror_mime_type: text/x-java
  language_id: 17
Apollo Guidance Computer:
  type: programming
  group: Assembly
  extensions:
  - ".agc"
  tm_scope: source.agc
  ace_mode: assembly_x86
  language_id: 18
AppleScript:
  type: programming
  aliases:
  - osascript
  extensions:
  - ".applescript"
  - ".scpt"
  interpreters:
  - osascript
  ace_mode: applescript
  color: "#101F1F"
  language_id: 19
Arc:
  type: programming
  color: "#aa2afe"
  extensions:
  - ".arc"
  tm_scope: none
  ace_mode: text
  language_id: 20
AsciiDoc:
  type: prose
  ace_mode: asciidoc
  wrap: true
  extensions:
  - ".asciidoc"
  - ".adoc"
  - ".asc"
  tm_scope: text.html.asciidoc
  language_id: 22
AspectJ:
  type: programming
  color: "#a957b0"
  extensions:
  - ".aj"
  tm_scope: source.aspectj
  ace_mode: text
  language_id: 23
Assembly:
  type: programming
  color: "#6E4C13"
  aliases:
  - nasm
  extensions:
  - ".asm"
  - ".a51"
  - ".inc"
  - ".nasm"
  tm_scope: source.assembly
  ace_mode: assembly_x86
  language_id: 24
Augeas:
  type: programming
  extensions:
  - ".aug"
  tm_scope: none
  ace_mode: text
  language_id: 25
AutoHotkey:
  type: programming
  color: "#6594b9"
  aliases:
  - ahk
  extensions:
  - ".ahk"
  - ".ahkl"
  tm_scope: source.ahk
  ace_mode: autohotkey
  language_id: 26
AutoIt:
  type: programming
  color: "#1C3552"
  aliases:
  - au3
  - AutoIt3
  - AutoItScript
  extensions:
  - ".au3"
  tm_scope: source.autoit
  ace_mode: autohotkey
  language_id: 27
Awk:
  type: programming
  extensions:
  - ".awk"
  - ".auk"
  - ".gawk"
  - ".mawk"
  - ".nawk"
  interpreters:
  - awk
  - gawk
  - mawk
  - nawk
  ace_mode: text
  language_id: 28
Ballerina:
  type: programming
  extensions:
  - ".bal"
  tm_scope: source.ballerina
  ace_mode: text
  color: "#FF5000"
  language_id: 720859680
Batchfile:
  type: programming
  aliases:
  - bat
  - batch
  - dosbatch
  - winbatch
  extensions:
  - ".bat"
  - ".cmd"
  tm_scope: source.batchfile
  ace_mode: batchfile
  color: "#C1F12E"
  language_id: 29
Befunge:
  type: programming
  extensions:
  - ".befunge"
  ace_mode: text
  language_id: 30
Bison:
  type: programming
  group: Yacc
  tm_scope: source.bison
  extensions:
  - ".bison"
  ace_mode: text
  language_id: 31
BitBake:
  type: programming
  tm_scope: none
  extensions:
  - ".bb"
  ace_mode: text
  language_id: 32
Blade:
  type: markup
  group: HTML
  extensions:
  - ".blade"
  - ".blade.php"
  tm_scope: text.html.php.blade
  ace_mode: text
  language_id: 33
BlitzBasic:
  type: programming
  aliases:
  - b3d
  - blitz3d
  - blitzplus
  - bplus
  extensions:
  - ".bb"
  - ".decls"
  tm_scope: source.blitzmax
  ace_mode: text
  language_id: 34
BlitzMax:
  type: programming
  color: "#cd6400"
  extensions:
  - ".bmx"
  aliases:
  - bmax
  ace_mode: text
  language_id: 35
Bluespec:
  type: programming
  extensions:
  - ".bsv"
  tm_scope: source.bsv
  ace_mode: verilog
  language_id: 36
Boo:
  type: programming
  color: "#d4bec1"
  extensions:
  - ".boo"
  ace_mode: text
  tm_scope: source.boo
  language_id: 37
Brainfuck:
  type: programming
  color: "#2F2530"
  extensions:
  - ".b"
  - ".bf"
  tm_scope: source.bf
  ace_mode: text
  codemirror_mode: brainfuck
  codemirror_mime_type: text/x-brainfuck
  language_id: 38
Brightscript:
  type: programming
  extensions:
  - ".brs"
  tm_scope: source.brightscript
  ace_mode: text
  language_id: 39
Bro:
  type: programming
  extensions:
  - ".bro"
  ace_mode: text
  language_id: 40
C:
  type: programming
  color: "#555555"
  extensions:
  - ".c"
  - ".cats"
  - ".h"
  - ".idc"
  interpreters:
  - tcc
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csrc
  language_id: 41
C#:
  type: programming
  ace_mode: csharp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csharp
  tm_scope: source.cs
  color: "#178600"
  aliases:
  - csharp
  extensions:
  - ".cs"
  - ".cake"
  - ".cshtml"
  - ".csx"
  language_id: 42
C++:
  type: programming
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-c++src
  color: "#f34b7d"
  aliases:
  - cpp
  extensions:
  - ".cpp"
  - ".c++"
  - ".cc"
  - ".cp"
  - ".cxx"
  - ".h"
  - ".h++"
  - ".hh"
  - ".hpp"
  - ".hxx"
  - ".inc"
  - ".inl"
  - ".ino"
  - ".ipp"
  - ".re"
  - ".tcc"
  - ".tpp"
  language_id: 43
C-ObjDump:
  type: data
  extensions:
  - ".c-objdump"
  tm_scope: objdump.x86asm
  ace_mode: assembly_x86
  language_id: 44
C2hs Haskell:
  type: programming
  group: Haskell
  aliases:
  - c2hs
  extensions:
  - ".chs"
  tm_scope: source.haskell
  ace_mode: haskell
  codemirror_mode: haskell
  codemirror_mime_type: text/x-haskell
  language_id: 45
CLIPS:
  type: programming
  extensions:
  - ".clp"
  tm_scope: source.clips
  ace_mode: text
  language_id: 46
CMake:
  type: programming
  extensions:
  - ".cmake"
  - ".cmake.in"
  filenames:
  - CMakeLists.txt
  ace_mode: text
  codemirror_mode: cmake
  codemirror_mime_type: text/x-cmake
  language_id: 47
COBOL:
  type: programming
  extensions:
  - ".cob"
  - ".cbl"
  - ".ccp"
  - ".cobol"
  - ".cpy"
  ace_mode: cobol
  codemirror_mode: cobol
  codemirror_mime_type: text/x-cobol
  language_id: 48
COLLADA:
  type: data
  extensions:
  - ".dae"
  tm_scope: text.xml
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  language_id: 49
CSON:
  type: data
  group: CoffeeScript
  tm_scope: source.coffee
  ace_mode: coffee
  codemirror_mode: coffeescript
  codemirror_mime_type: text/x-coffeescript
  searchable: false
  extensions:
  - ".cson"
  language_id: 424
CSS:
  type: markup
  tm_scope: source.css
  ace_mode: css
  codemirror_mode: css
  codemirror_mime_type: text/css
  color: "#563d7c"
  extensions:
  - ".css"
  language_id: 50
CSV:
  type: data
  ace_mode: text
  tm_scope: none
  extensions:
  - ".csv"
  language_id: 51
CWeb:
  type: programming
  extensions:
  - ".w"
  tm_scope: none
  ace_mode: text
  language_id: 657332628
Cap'n Proto:
  type: programming
  tm_scope: source.capnp
  extensions:
  - ".capnp"
  ace_mode: text
  language_id: 52
CartoCSS:
  type: programming
  aliases:
  - Carto
  extensions:
  - ".mss"
  ace_mode: text
  tm_scope: source.css.mss
  language_id: 53
Ceylon:
  type: programming
  color: "#dfa535"
  extensions:
  - ".ceylon"
  tm_scope: source.ceylon
  ace_mode: text
  language_id: 54
Chapel:
  type: programming
  color: "#8dc63f"
  aliases:
  - chpl
  extensions:
  - ".chpl"
  ace_mode: text
  language_id: 55
Charity:
  type: programming
  extensions:
  - ".ch"
  tm_scope: none
  ace_mode: text
  language_id: 56
ChucK:
  type: programming
  extensions:
  - ".ck"
  tm_scope: source.java
  ace_mode: java
  codemirror_mode: clike
  codemirror_mime_type: text/x-java
  language_id: 57
Cirru:
  type: programming
  color: "#ccccff"
  ace_mode: cirru
  extensions:
  - ".cirru"
  language_id: 58
Clarion:
  type: programming
  color: "#db901e"
  ace_mode: text
  extensions:
  - ".clw"
  tm_scope: source.clarion
  language_id: 59
Clean:
  type: programming
  color: "#3F85AF"
  extensions:
  - ".icl"
  - ".dcl"
  tm_scope: source.clean
  ace_mode: text
  language_id: 60
Click:
  type: programming
  color: "#E4E6F3"
  extensions:
  - ".click"
  tm_scope: source.click
  ace_mode: text
  language_id: 61
Clojure:
  type: programming
  ace_mode: clojure
  codemirror_mode: clojure
  codemirror_mime_type: text/x-clojure
  color: "#db5855"
  extensions:
  - ".clj"
  - ".boot"
  - ".cl2"
  - ".cljc"
  - ".cljs"
  - ".cljs.hl"
  - ".cljscm"
  - ".cljx"
  - ".hic"
  filenames:
  - riemann.config
  language_id: 62
Closure Templates:
  type: markup
  group: HTML
  ace_mode: soy_template
  codemirror_mode: soy
  codemirror_mime_type: text/x-soy
  alias:
  - soy
  extensions:
  - ".soy"
  tm_scope: text.html.soy
  language_id: 357046146
CoffeeScript:
  type: programming
  tm_scope: source.coffee
  ace_mode: coffee
  codemirror_mode: coffeescript
  codemirror_mime_type: text/x-coffeescript
  color: "#244776"
  aliases:
  - coffee
  - coffee-script
  extensions:
  - ".coffee"
  - "._coffee"
  - ".cake"
  - ".cjsx"
  - ".iced"
  filenames:
  - Cakefile
  interpreters:
  - coffee
  language_id: 63
ColdFusion:
  type: programming
  ace_mode: coldfusion
  color: "#ed2cd6"
  aliases:
  - cfm
  - cfml
  - coldfusion html
  extensions:
  - ".cfm"
  - ".cfml"
  tm_scope: text.html.cfm
  language_id: 64
ColdFusion CFC:
  type: programming
  group: ColdFusion
  ace_mode: coldfusion
  aliases:
  - cfc
  extensions:
  - ".cfc"
  tm_scope: source.cfscript
  language_id: 65
Common Lisp:
  type: programming
  tm_scope: source.lisp
  color: "#3fb68b"
  aliases:
  - lisp
  extensions:
  - ".lisp"
  - ".asd"
  - ".cl"
  - ".l"
  - ".lsp"
  - ".ny"
  - ".podsl"
  - ".sexp"
  interpreters:
  - lisp
  - sbcl
  - ccl
  - clisp
  - ecl
  ace_mode: lisp
  codemirror_mode: commonlisp
  codemirror_mime_type: text/x-common-lisp
  language_id: 66
Common Workflow Language:
  alias: cwl
  type: programming
  ace_mode: yaml
  codemirror_mode: yaml
  codemirror_mime_type: text/x-yaml
  extensions:
  - ".cwl"
  interpreters:
  - cwl-runner
  color: "#B5314C"
  tm_scope: source.cwl
  language_id: 988547172
Component Pascal:
  type: programming
  color: "#B0CE4E"
  extensions:
  - ".cp"
  - ".cps"
  tm_scope: source.pascal
  aliases:
  - delphi
  - objectpascal
  ace_mode: pascal
  codemirror_mode: pascal
  codemirror_mime_type: text/x-pascal
  language_id: 67
Cool:
  type: programming
  extensions:
  - ".cl"
  tm_scope: source.cool
  ace_mode: text
  language_id: 68
Coq:
  type: programming
  extensions:
  - ".coq"
  - ".v"
  ace_mode: text
  language_id: 69
Cpp-ObjDump:
  type: data
  extensions:
  - ".cppobjdump"
  - ".c++-objdump"
  - ".c++objdump"
  - ".cpp-objdump"
  - ".cxx-objdump"
  tm_scope: objdump.x86asm
  aliases:
  - c++-objdump
  ace_mode: assembly_x86
  language_id: 70
Creole:
  type: prose
  wrap: true
  extensions:
  - ".creole"
  tm_scope: text.html.creole
  ace_mode: text
  language_id: 71
Crystal:
  type: programming
  color: "#776791"
  extensions:
  - ".cr"
  ace_mode: ruby
  codemirror_mode: crystal
  codemirror_mime_type: text/x-crystal
  tm_scope: source.crystal
  interpreters:
  - crystal
  language_id: 72
Csound:
  type: programming
  aliases:
  - csound-orc
  extensions:
  - ".orc"
  - ".udo"
  tm_scope: source.csound
  ace_mode: csound_orchestra
  language_id: 73
Csound Document:
  type: programming
  aliases:
  - csound-csd
  extensions:
  - ".csd"
  tm_scope: source.csound-document
  ace_mode: csound_document
  language_id: 74
Csound Score:
  type: programming
  aliases:
  - csound-sco
  extensions:
  - ".sco"
  tm_scope: source.csound-score
  ace_mode: csound_score
  language_id: 75
Cuda:
  type: programming
  extensions:
  - ".cu"
  - ".cuh"
  tm_scope: source.cuda-c++
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-c++src
  color: "#3A4E3A"
  language_id: 77
Cycript:
  type: programming
  extensions:
  - ".cy"
  tm_scope: source.js
  ace_mode: javascript
  codemirror_mode: javascript
  codemirror_mime_type: text/javascript
  language_id: 78
Cython:
  type: programming
  group: Python
  extensions:
  - ".pyx"
  - ".pxd"
  - ".pxi"
  aliases:
  - pyrex
  ace_mode: text
  codemirror_mode: python
  codemirror_mime_type: text/x-cython
  language_id: 79
D:
  type: programming
  color: "#ba595e"
  extensions:
  - ".d"
  - ".di"
  ace_mode: d
  codemirror_mode: d
  codemirror_mime_type: text/x-d
  language_id: 80
D-ObjDump:
  type: data
  extensions:
  - ".d-objdump"
  tm_scope: objdump.x86asm
  ace_mode: assembly_x86
  language_id: 81
DIGITAL Command Language:
  type: programming
  aliases:
  - dcl
  extensions:
  - ".com"
  tm_scope: none
  ace_mode: text
  language_id: 82
DM:
  type: programming
  color: "#447265"
  extensions:
  - ".dm"
  aliases:
  - byond
  tm_scope: source.dm
  ace_mode: c_cpp
  language_id: 83
DNS Zone:
  type: data
  extensions:
  - ".zone"
  - ".arpa"
  tm_scope: text.zone_file
  ace_mode: text
  language_id: 84
DTrace:
  type: programming
  aliases:
  - dtrace-script
  extensions:
  - ".d"
  interpreters:
  - dtrace
  tm_scope: source.c
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csrc
  language_id: 85
Darcs Patch:
  type: data
  aliases:
  - dpatch
  extensions:
  - ".darcspatch"
  - ".dpatch"
  tm_scope: none
  ace_mode: text
  language_id: 86
Dart:
  type: programming
  color: "#00B4AB"
  extensions:
  - ".dart"
  interpreters:
  - dart
  ace_mode: dart
  codemirror_mode: dart
  codemirror_mime_type: application/dart
  language_id: 87
DataWeave:
  type: programming
  color: "#003a52"
  extensions:
  - ".dwl"
  ace_mode: text
  tm_scope: source.data-weave
  language_id: 974514097
Diff:
  type: data
  extensions:
  - ".diff"
  - ".patch"
  aliases:
  - udiff
  tm_scope: source.diff
  ace_mode: diff
  codemirror_mode: diff
  codemirror_mime_type: text/x-diff
  language_id: 88
Dockerfile:
  type: data
  tm_scope: source.dockerfile
  extensions:
  - ".dockerfile"
  filenames:
  - Dockerfile
  ace_mode: dockerfile
  codemirror_mode: dockerfile
  codemirror_mime_type: text/x-dockerfile
  language_id: 89
Dogescript:
  type: programming
  color: "#cca760"
  extensions:
  - ".djs"
  tm_scope: none
  ace_mode: text
  language_id: 90
Dylan:
  type: programming
  color: "#6c616e"
  extensions:
  - ".dylan"
  - ".dyl"
  - ".intr"
  - ".lid"
  ace_mode: text
  codemirror_mode: dylan
  codemirror_mime_type: text/x-dylan
  language_id: 91
E:
  type: programming
  color: "#ccce35"
  extensions:
  - ".E"
  interpreters:
  - rune
  tm_scope: none
  ace_mode: text
  language_id: 92
EBNF:
  type: data
  extensions:
  - ".ebnf"
  tm_scope: source.ebnf
  ace_mode: text
  codemirror_mode: ebnf
  codemirror_mime_type: text/x-ebnf
  language_id: 430
ECL:
  type: programming
  color: "#8a1267"
  extensions:
  - ".ecl"
  - ".eclxml"
  tm_scope: none
  ace_mode: text
  codemirror_mode: ecl
  codemirror_mime_type: text/x-ecl
  language_id: 93
ECLiPSe:
  type: programming
  group: prolog
  extensions:
  - ".ecl"
  tm_scope: source.prolog.eclipse
  ace_mode: prolog
  language_id: 94
EJS:
  type: markup
  group: HTML
  extensions:
  - ".ejs"
  tm_scope: text.html.js
  ace_mode: ejs
  language_id: 95
EQ:
  type: programming
  color: "#a78649"
  extensions:
  - ".eq"
  tm_scope: source.cs
  ace_mode: csharp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csharp
  language_id: 96
Eagle:
  type: data
  extensions:
  - ".sch"
  - ".brd"
  tm_scope: text.xml
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  language_id: 97
Easybuild:
  type: data
  group: Python
  ace_mode: python
  codemirror_mode: python
  codemirror_mime_type: text/x-python
  tm_scope: source.python
  extensions:
  - ".eb"
  language_id: 342840477
Ecere Projects:
  type: data
  group: JavaScript
  extensions:
  - ".epj"
  tm_scope: source.json
  ace_mode: json
  codemirror_mode: javascript
  codemirror_mime_type: application/json
  language_id: 98
Edje Data Collection:
  type: data
  extensions:
  - ".edc"
  tm_scope: source.json
  ace_mode: json
  codemirror_mode: javascript
  codemirror_mime_type: application/json
  language_id: 342840478
Eiffel:
  type: programming
  color: "#946d57"
  extensions:
  - ".e"
  ace_mode: eiffel
  codemirror_mode: eiffel
  codemirror_mime_type: text/x-eiffel
  language_id: 99
Elixir:
  type: programming
  color: "#6e4a7e"
  extensions:
  - ".ex"
  - ".exs"
  ace_mode: elixir
  filenames:
  - mix.lock
  interpreters:
  - elixir
  language_id: 100
Elm:
  type: programming
  color: "#60B5CC"
  extensions:
  - ".elm"
  tm_scope: source.elm
  ace_mode: elm
  codemirror_mode: elm
  codemirror_mime_type: text/x-elm
  language_id: 101
Emacs Lisp:
  type: programming
  tm_scope: source.emacs.lisp
  color: "#c065db"
  aliases:
  - elisp
  - emacs
  filenames:
  - ".abbrev_defs"
  - ".emacs"
  - ".emacs.desktop"
  - ".gnus"
  - ".spacemacs"
  - ".viper"
  - Cask
  - Project.ede
  - _emacs
  - abbrev_defs
  extensions:
  - ".el"
  - ".emacs"
  - ".emacs.desktop"
  ace_mode: lisp
  codemirror_mode: commonlisp
  codemirror_mime_type: text/x-common-lisp
  language_id: 102
EmberScript:
  type: programming
  color: "#FFF4F3"
  extensions:
  - ".em"
  - ".emberscript"
  tm_scope: source.coffee
  ace_mode: coffee
  codemirror_mode: coffeescript
  codemirror_mime_type: text/x-coffeescript
  language_id: 103
Erlang:
  type: programming
  color: "#B83998"
  extensions:
  - ".erl"
  - ".app.src"
  - ".es"
  - ".escript"
  - ".hrl"
  - ".xrl"
  - ".yrl"
  filenames:
  - Emakefile
  - rebar.config
  - rebar.config.lock
  - rebar.lock
  ace_mode: erlang
  codemirror_mode: erlang
  codemirror_mime_type: text/x-erlang
  interpreters:
  - escript
  language_id: 104
F#:
  type: programming
  color: "#b845fc"
  aliases:
  - fsharp
  extensions:
  - ".fs"
  - ".fsi"
  - ".fsx"
  tm_scope: source.fsharp
  ace_mode: text
  codemirror_mode: mllike
  codemirror_mime_type: text/x-fsharp
  language_id: 105
FLUX:
  type: programming
  color: "#88ccff"
  extensions:
  - ".fx"
  - ".flux"
  tm_scope: none
  ace_mode: text
  language_id: 106
Factor:
  type: programming
  color: "#636746"
  extensions:
  - ".factor"
  filenames:
  - ".factor-boot-rc"
  - ".factor-rc"
  ace_mode: text
  codemirror_mode: factor
  codemirror_mime_type: text/x-factor
  language_id: 108
Fancy:
  type: programming
  color: "#7b9db4"
  extensions:
  - ".fy"
  - ".fancypack"
  filenames:
  - Fakefile
  ace_mode: text
  language_id: 109
Fantom:
  type: programming
  color: "#14253c"
  extensions:
  - ".fan"
  tm_scope: source.fan
  ace_mode: text
  language_id: 110
Filebench WML:
  type: programming
  extensions:
  - ".f"
  tm_scope: none
  ace_mode: text
  language_id: 111
Filterscript:
  type: programming
  group: RenderScript
  extensions:
  - ".fs"
  tm_scope: none
  ace_mode: text
  language_id: 112
Formatted:
  type: data
  extensions:
  - ".for"
  - ".eam.fs"
  tm_scope: none
  ace_mode: text
  language_id: 113
Forth:
  type: programming
  color: "#341708"
  extensions:
  - ".fth"
  - ".4th"
  - ".f"
  - ".for"
  - ".forth"
  - ".fr"
  - ".frt"
  - ".fs"
  ace_mode: forth
  codemirror_mode: forth
  codemirror_mime_type: text/x-forth
  language_id: 114
Fortran:
  type: programming
  color: "#4d41b1"
  extensions:
  - ".f90"
  - ".f"
  - ".f03"
  - ".f08"
  - ".f77"
  - ".f95"
  - ".for"
  - ".fpp"
  tm_scope: source.fortran.modern
  ace_mode: text
  codemirror_mode: fortran
  codemirror_mime_type: text/x-fortran
  language_id: 107
FreeMarker:
  type: programming
  color: "#0050b2"
  aliases:
  - ftl
  extensions:
  - ".ftl"
  tm_scope: text.html.ftl
  ace_mode: ftl
  language_id: 115
Frege:
  type: programming
  color: "#00cafe"
  extensions:
  - ".fr"
  tm_scope: source.haskell
  ace_mode: haskell
  language_id: 116
G-code:
  type: data
  extensions:
  - ".g"
  - ".gco"
  - ".gcode"
  tm_scope: source.gcode
  ace_mode: gcode
  language_id: 117
GAMS:
  type: programming
  extensions:
  - ".gms"
  tm_scope: none
  ace_mode: text
  language_id: 118
GAP:
  type: programming
  extensions:
  - ".g"
  - ".gap"
  - ".gd"
  - ".gi"
  - ".tst"
  tm_scope: source.gap
  ace_mode: text
  language_id: 119
GCC Machine Description:
  type: programming
  extensions:
  - ".md"
  tm_scope: source.lisp
  ace_mode: lisp
  codemirror_mode: commonlisp
  codemirror_mime_type: text/x-common-lisp
  language_id: 121
GDB:
  type: programming
  extensions:
  - ".gdb"
  - ".gdbinit"
  tm_scope: source.gdb
  ace_mode: text
  language_id: 122
GDScript:
  type: programming
  extensions:
  - ".gd"
  tm_scope: source.gdscript
  ace_mode: text
  language_id: 123
GLSL:
  type: programming
  extensions:
  - ".glsl"
  - ".fp"
  - ".frag"
  - ".frg"
  - ".fs"
  - ".fsh"
  - ".fshader"
  - ".geo"
  - ".geom"
  - ".glslv"
  - ".gshader"
  - ".shader"
  - ".tesc"
  - ".tese"
  - ".vert"
  - ".vrx"
  - ".vsh"
  - ".vshader"
  ace_mode: glsl
  language_id: 124
GN:
  type: data
  extensions:
  - ".gn"
  - ".gni"
  interpreters:
  - gn
  tm_scope: source.gn
  ace_mode: python
  codemirror_mode: python
  codemirror_mime_type: text/x-python
  language_id: 302957008
Game Maker Language:
  type: programming
  color: "#8fb200"
  extensions:
  - ".gml"
  tm_scope: source.c++
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-c++src
  language_id: 125
Genie:
  type: programming
  ace_mode: text
  extensions:
  - ".gs"
  color: "#fb855d"
  tm_scope: none
  language_id: 792408528
Genshi:
  type: programming
  extensions:
  - ".kid"
  tm_scope: text.xml.genshi
  aliases:
  - xml+genshi
  - xml+kid
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  language_id: 126
Gentoo Ebuild:
  type: programming
  group: Shell
  extensions:
  - ".ebuild"
  tm_scope: source.shell
  ace_mode: sh
  codemirror_mode: shell
  codemirror_mime_type: text/x-sh
  language_id: 127
Gentoo Eclass:
  type: programming
  group: Shell
  extensions:
  - ".eclass"
  tm_scope: source.shell
  ace_mode: sh
  codemirror_mode: shell
  codemirror_mime_type: text/x-sh
  language_id: 128
Gerber Image:
  type: data
  aliases:
  - rs-274x
  extensions:
  - ".gbr"
  - ".gbl"
  - ".gbo"
  - ".gbp"
  - ".gbs"
  - ".gko"
  - ".gpb"
  - ".gpt"
  - ".gtl"
  - ".gto"
  - ".gtp"
  - ".gts"
  interpreters:
  - gerbv
  - gerbview
  tm_scope: source.gerber
  ace_mode: text
  language_id: 404627610
Gettext Catalog:
  type: prose
  searchable: false
  aliases:
  - pot
  extensions:
  - ".po"
  - ".pot"
  tm_scope: source.po
  ace_mode: text
  language_id: 129
Gherkin:
  type: programming
  extensions:
  - ".feature"
  tm_scope: text.gherkin.feature
  aliases:
  - cucumber
  ace_mode: text
  color: "#5B2063"
  language_id: 76
Glyph:
  type: programming
  color: "#e4cc98"
  extensions:
  - ".glf"
  tm_scope: source.tcl
  ace_mode: tcl
  codemirror_mode: tcl
  codemirror_mime_type: text/x-tcl
  language_id: 130
Gnuplot:
  type: programming
  color: "#f0a9f0"
  extensions:
  - ".gp"
  - ".gnu"
  - ".gnuplot"
  - ".plot"
  - ".plt"
  interpreters:
  - gnuplot
  ace_mode: text
  language_id: 131
Go:
  type: programming
  color: "#375eab"
  aliases:
  - golang
  extensions:
  - ".go"
  ace_mode: golang
  codemirror_mode: go
  codemirror_mime_type: text/x-go
  language_id: 132
Golo:
  type: programming
  color: "#88562A"
  extensions:
  - ".golo"
  tm_scope: source.golo
  ace_mode: text
  language_id: 133
Gosu:
  type: programming
  color: "#82937f"
  extensions:
  - ".gs"
  - ".gst"
  - ".gsx"
  - ".vark"
  tm_scope: source.gosu.2
  ace_mode: text
  language_id: 134
Grace:
  type: programming
  extensions:
  - ".grace"
  tm_scope: source.grace
  ace_mode: text
  language_id: 135
Gradle:
  type: data
  extensions:
  - ".gradle"
  tm_scope: source.groovy.gradle
  ace_mode: text
  language_id: 136
Grammatical Framework:
  type: programming
  aliases:
  - gf
  wrap: false
  extensions:
  - ".gf"
  searchable: true
  color: "#79aa7a"
  tm_scope: source.haskell
  ace_mode: haskell
  codemirror_mode: haskell
  codemirror_mime_type: text/x-haskell
  language_id: 137
Graph Modeling Language:
  type: data
  extensions:
  - ".gml"
  tm_scope: none
  ace_mode: text
  language_id: 138
GraphQL:
  type: data
  extensions:
  - ".graphql"
  - ".gql"
  tm_scope: source.graphql
  ace_mode: text
  language_id: 139
Graphviz (DOT):
  type: data
  tm_scope: source.dot
  extensions:
  - ".dot"
  - ".gv"
  ace_mode: text
  language_id: 140
Groovy:
  type: programming
  ace_mode: groovy
  codemirror_mode: groovy
  codemirror_mime_type: text/x-groovy
  color: "#e69f56"
  extensions:
  - ".groovy"
  - ".grt"
  - ".gtpl"
  - ".gvy"
  interpreters:
  - groovy
  filenames:
  - Jenkinsfile
  language_id: 142
Groovy Server Pages:
  type: programming
  group: Groovy
  aliases:
  - gsp
  - java server page
  extensions:
  - ".gsp"
  tm_scope: text.html.jsp
  ace_mode: jsp
  codemirror_mode: htmlembedded
  codemirror_mime_type: application/x-jsp
  language_id: 143
HCL:
  type: programming
  extensions:
  - ".hcl"
  - ".tf"
  - ".tfvars"
  ace_mode: ruby
  codemirror_mode: ruby
  codemirror_mime_type: text/x-ruby
  tm_scope: source.terraform
  language_id: 144
HLSL:
  type: programming
  extensions:
  - ".hlsl"
  - ".cginc"
  - ".fx"
  - ".fxh"
  - ".hlsli"
  ace_mode: text
  tm_scope: source.hlsl
  language_id: 145
HTML:
  type: markup
  tm_scope: text.html.basic
  ace_mode: html
  codemirror_mode: htmlmixed
  codemirror_mime_type: text/html
  color: "#e34c26"
  aliases:
  - xhtml
  extensions:
  - ".html"
  - ".htm"
  - ".html.hl"
  - ".inc"
  - ".st"
  - ".xht"
  - ".xhtml"
  language_id: 146
HTML+Django:
  type: markup
  tm_scope: text.html.django
  group: HTML
  extensions:
  - ".jinja"
  - ".mustache"
  - ".njk"
  aliases:
  - django
  - html+django/jinja
  - html+jinja
  - htmldjango
  - njk
  - nunjucks
  ace_mode: django
  codemirror_mode: django
  codemirror_mime_type: text/x-django
  language_id: 147
HTML+ECR:
  type: markup
  tm_scope: text.html.ecr
  group: HTML
  aliases:
  - ecr
  extensions:
  - ".ecr"
  ace_mode: text
  codemirror_mode: htmlmixed
  codemirror_mime_type: text/html
  language_id: 148
HTML+EEX:
  type: markup
  tm_scope: text.html.elixir
  group: HTML
  aliases:
  - eex
  extensions:
  - ".eex"
  ace_mode: text
  codemirror_mode: htmlmixed
  codemirror_mime_type: text/html
  language_id: 149
HTML+ERB:
  type: markup
  tm_scope: text.html.erb
  group: HTML
  aliases:
  - erb
  extensions:
  - ".erb"
  - ".erb.deface"
  ace_mode: text
  codemirror_mode: htmlembedded
  codemirror_mime_type: application/x-erb
  language_id: 150
HTML+PHP:
  type: markup
  tm_scope: text.html.php
  group: HTML
  extensions:
  - ".phtml"
  ace_mode: php
  codemirror_mode: php
  codemirror_mime_type: application/x-httpd-php
  language_id: 151
HTTP:
  type: data
  extensions:
  - ".http"
  tm_scope: source.httpspec
  ace_mode: text
  codemirror_mode: http
  codemirror_mime_type: message/http
  language_id: 152
Hack:
  type: programming
  ace_mode: php
  codemirror_mode: php
  codemirror_mime_type: application/x-httpd-php
  extensions:
  - ".hh"
  - ".php"
  tm_scope: text.html.php
  color: "#878787"
  language_id: 153
Haml:
  group: HTML
  type: markup
  extensions:
  - ".haml"
  - ".haml.deface"
  ace_mode: haml
  codemirror_mode: haml
  codemirror_mime_type: text/x-haml
  language_id: 154
Handlebars:
  type: markup
  group: HTML
  aliases:
  - hbs
  - htmlbars
  extensions:
  - ".handlebars"
  - ".hbs"
  tm_scope: text.html.handlebars
  ace_mode: handlebars
  language_id: 155
Harbour:
  type: programming
  color: "#0e60e3"
  extensions:
  - ".hb"
  tm_scope: source.harbour
  ace_mode: text
  language_id: 156
Haskell:
  type: programming
  color: "#5e5086"
  extensions:
  - ".hs"
  - ".hsc"
  interpreters:
  - runhaskell
  ace_mode: haskell
  codemirror_mode: haskell
  codemirror_mime_type: text/x-haskell
  language_id: 157
Haxe:
  type: programming
  ace_mode: haxe
  codemirror_mode: haxe
  codemirror_mime_type: text/x-haxe
  color: "#df7900"
  extensions:
  - ".hx"
  - ".hxsl"
  tm_scope: source.haxe.2
  language_id: 158
Hy:
  type: programming
  ace_mode: text
  color: "#7790B2"
  extensions:
  - ".hy"
  aliases:
  - hylang
  tm_scope: none
  language_id: 159
HyPhy:
  type: programming
  ace_mode: text
  extensions:
  - ".bf"
  tm_scope: none
  language_id: 160
IDL:
  type: programming
  color: "#a3522f"
  extensions:
  - ".pro"
  - ".dlm"
  ace_mode: text
  codemirror_mode: idl
  codemirror_mime_type: text/x-idl
  language_id: 161
IGOR Pro:
  type: programming
  extensions:
  - ".ipf"
  aliases:
  - igor
  - igorpro
  tm_scope: none
  ace_mode: text
  language_id: 162
INI:
  type: data
  extensions:
  - ".ini"
  - ".cfg"
  - ".prefs"
  - ".pro"
  - ".properties"
  filenames:
  - buildozer.spec
  tm_scope: source.ini
  aliases:
  - dosini
  ace_mode: ini
  codemirror_mode: properties
  codemirror_mime_type: text/x-properties
  language_id: 163
IRC log:
  type: data
  aliases:
  - irc
  - irc logs
  extensions:
  - ".irclog"
  - ".weechatlog"
  tm_scope: none
  ace_mode: text
  codemirror_mode: mirc
  codemirror_mime_type: text/mirc
  language_id: 164
Idris:
  type: programming
  color: "#b30000"
  extensions:
  - ".idr"
  - ".lidr"
  ace_mode: text
  tm_scope: source.idris
  language_id: 165
Inform 7:
  type: programming
  wrap: true
  extensions:
  - ".ni"
  - ".i7x"
  tm_scope: source.inform7
  aliases:
  - i7
  - inform7
  ace_mode: text
  language_id: 166
Inno Setup:
  type: programming
  extensions:
  - ".iss"
  tm_scope: none
  ace_mode: text
  language_id: 167
Io:
  type: programming
  color: "#a9188d"
  extensions:
  - ".io"
  interpreters:
  - io
  ace_mode: io
  language_id: 168
Ioke:
  type: programming
  color: "#078193"
  extensions:
  - ".ik"
  interpreters:
  - ioke
  ace_mode: text
  language_id: 169
Isabelle:
  type: programming
  color: "#FEFE00"
  extensions:
  - ".thy"
  tm_scope: source.isabelle.theory
  ace_mode: text
  language_id: 170
Isabelle ROOT:
  type: programming
  group: Isabelle
  filenames:
  - ROOT
  tm_scope: source.isabelle.root
  ace_mode: text
  language_id: 171
J:
  type: programming
  color: "#9EEDFF"
  extensions:
  - ".ijs"
  interpreters:
  - jconsole
  tm_scope: source.j
  ace_mode: text
  language_id: 172
JFlex:
  type: programming
  group: Lex
  extensions:
  - ".flex"
  - ".jflex"
  tm_scope: source.jflex
  ace_mode: text
  language_id: 173
JSON:
  type: data
  tm_scope: source.json
  group: JavaScript
  ace_mode: json
  codemirror_mode: javascript
  codemirror_mime_type: application/json
  searchable: false
  extensions:
  - ".json"
  - ".geojson"
  - ".JSON-tmLanguage"
  - ".topojson"
  filenames:
  - ".arcconfig"
  - ".jshintrc"
  - composer.lock
  - mcmod.info
  language_id: 174
JSON5:
  type: data
  extensions:
  - ".json5"
  filenames:
  - ".babelrc"
  tm_scope: source.js
  ace_mode: javascript
  codemirror_mode: javascript
  codemirror_mime_type: application/json
  language_id: 175
JSONLD:
  type: data
  group: JavaScript
  ace_mode: javascript
  extensions:
  - ".jsonld"
  tm_scope: source.js
  language_id: 176
JSONiq:
  color: "#40d47e"
  type: programming
  ace_mode: jsoniq
  codemirror_mode: javascript
  codemirror_mime_type: application/json
  extensions:
  - ".jq"
  tm_scope: source.jq
  language_id: 177
JSX:
  type: programming
  group: JavaScript
  extensions:
  - ".jsx"
  tm_scope: source.js.jsx
  ace_mode: javascript
  codemirror_mode: jsx
  codemirror_mime_type: text/jsx
  language_id: 178
Jasmin:
  type: programming
  ace_mode: java
  extensions:
  - ".j"
  tm_scope: source.jasmin
  language_id: 180
Java:
  type: programming
  ace_mode: java
  codemirror_mode: clike
  codemirror_mime_type: text/x-java
  color: "#b07219"
  extensions:
  - ".java"
  language_id: 181
Java Server Pages:
  type: programming
  group: Java
  aliases:
  - jsp
  extensions:
  - ".jsp"
  tm_scope: text.html.jsp
  ace_mode: jsp
  codemirror_mode: htmlembedded
  codemirror_mime_type: application/x-jsp
  language_id: 182
JavaScript:
  type: programming
  tm_scope: source.js
  ace_mode: javascript
  codemirror_mode: javascript
  codemirror_mime_type: text/javascript
  color: "#f1e05a"
  aliases:
  - js
  - node
  extensions:
  - ".js"
  - "._js"
  - ".bones"
  - ".es"
  - ".es6"
  - ".frag"
  - ".gs"
  - ".jake"
  - ".jsb"
  - ".jscad"
  - ".jsfl"
  - ".jsm"
  - ".jss"
  - ".mjs"
  - ".njs"
  - ".pac"
  - ".sjs"
  - ".ssjs"
  - ".xsjs"
  - ".xsjslib"
  filenames:
  - Jakefile
  interpreters:
  - node
  language_id: 183
Jison:
  type: programming
  group: Yacc
  extensions:
  - ".jison"
  tm_scope: source.jison
  ace_mode: text
  language_id: 284531423
Jison Lex:
  type: programming
  group: Lex
  extensions:
  - ".jisonlex"
  tm_scope: source.jisonlex
  ace_mode: text
  language_id: 406395330
Jolie:
  type: programming
  extensions:
  - ".ol"
  - ".iol"
  interpreters:
  - jolie
  color: "#843179"
  ace_mode: text
  tm_scope: source.jolie
  language_id: 998078858
Julia:
  type: programming
  extensions:
  - ".jl"
  interpreters:
  - julia
  color: "#a270ba"
  ace_mode: julia
  codemirror_mode: julia
  codemirror_mime_type: text/x-julia
  language_id: 184
Jupyter Notebook:
  type: markup
  ace_mode: json
  codemirror_mode: javascript
  codemirror_mime_type: application/json
  tm_scope: source.json
  color: "#DA5B0B"
  extensions:
  - ".ipynb"
  filenames:
  - Notebook
  aliases:
  - IPython Notebook
  language_id: 185
KRL:
  type: programming
  color: "#28431f"
  extensions:
  - ".krl"
  tm_scope: none
  ace_mode: text
  language_id: 186
KiCad Layout:
  type: data
  aliases:
  - pcbnew
  extensions:
  - ".kicad_pcb"
  - ".kicad_mod"
  - ".kicad_wks"
  filenames:
  - fp-lib-table
  tm_scope: source.pcb.sexp
  ace_mode: lisp
  codemirror_mode: commonlisp
  codemirror_mime_type: text/x-common-lisp
  language_id: 187
KiCad Legacy Layout:
  type: data
  extensions:
  - ".brd"
  tm_scope: source.pcb.board
  ace_mode: text
  language_id: 140848857
KiCad Schematic:
  type: data
  aliases:
  - eeschema schematic
  extensions:
  - ".sch"
  tm_scope: source.pcb.schematic
  ace_mode: text
  language_id: 622447435
Kit:
  type: markup
  ace_mode: html
  codemirror_mode: htmlmixed
  codemirror_mime_type: text/html
  extensions:
  - ".kit"
  tm_scope: text.html.basic
  language_id: 188
Kotlin:
  type: programming
  color: "#F18E33"
  extensions:
  - ".kt"
  - ".ktm"
  - ".kts"
  tm_scope: source.Kotlin
  ace_mode: text
  codemirror_mode: clike
  codemirror_mime_type: text/x-kotlin
  language_id: 189
LFE:
  type: programming
  color: "#4C3023"
  extensions:
  - ".lfe"
  tm_scope: source.lisp
  ace_mode: lisp
  codemirror_mode: commonlisp
  codemirror_mime_type: text/x-common-lisp
  language_id: 190
LLVM:
  type: programming
  extensions:
  - ".ll"
  ace_mode: text
  color: "#185619"
  language_id: 191
LOLCODE:
  type: programming
  extensions:
  - ".lol"
  color: "#cc9900"
  tm_scope: none
  ace_mode: text
  language_id: 192
LSL:
  type: programming
  ace_mode: lsl
  extensions:
  - ".lsl"
  - ".lslp"
  interpreters:
  - lsl
  color: "#3d9970"
  language_id: 193
LabVIEW:
  type: programming
  extensions:
  - ".lvproj"
  tm_scope: text.xml
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  language_id: 194
Lasso:
  type: programming
  color: "#999999"
  extensions:
  - ".lasso"
  - ".las"
  - ".lasso8"
  - ".lasso9"
  - ".ldml"
  tm_scope: file.lasso
  aliases:
  - lassoscript
  ace_mode: text
  language_id: 195
Latte:
  type: markup
  group: HTML
  extensions:
  - ".latte"
  tm_scope: text.html.smarty
  ace_mode: smarty
  codemirror_mode: smarty
  codemirror_mime_type: text/x-smarty
  language_id: 196
Lean:
  type: programming
  extensions:
  - ".lean"
  - ".hlean"
  ace_mode: text
  language_id: 197
Less:
  type: markup
  group: CSS
  extensions:
  - ".less"
  tm_scope: source.css.less
  ace_mode: less
  codemirror_mode: css
  codemirror_mime_type: text/css
  language_id: 198
Lex:
  type: programming
  color: "#DBCA00"
  aliases:
  - flex
  extensions:
  - ".l"
  - ".lex"
  tm_scope: none
  ace_mode: text
  language_id: 199
LilyPond:
  type: programming
  extensions:
  - ".ly"
  - ".ily"
  ace_mode: text
  language_id: 200
Limbo:
  type: programming
  extensions:
  - ".b"
  - ".m"
  tm_scope: none
  ace_mode: text
  language_id: 201
Linker Script:
  type: data
  extensions:
  - ".ld"
  - ".lds"
  filenames:
  - ld.script
  tm_scope: none
  ace_mode: text
  language_id: 202
Linux Kernel Module:
  type: data
  extensions:
  - ".mod"
  tm_scope: none
  ace_mode: text
  language_id: 203
Liquid:
  type: markup
  extensions:
  - ".liquid"
  tm_scope: text.html.liquid
  ace_mode: liquid
  language_id: 204
Literate Agda:
  type: programming
  group: Agda
  extensions:
  - ".lagda"
  tm_scope: none
  ace_mode: text
  language_id: 205
Literate CoffeeScript:
  type: programming
  tm_scope: source.litcoffee
  group: CoffeeScript
  ace_mode: text
  wrap: true
  aliases:
  - litcoffee
  extensions:
  - ".litcoffee"
  language_id: 206
Literate Haskell:
  type: programming
  group: Haskell
  aliases:
  - lhaskell
  - lhs
  extensions:
  - ".lhs"
  tm_scope: text.tex.latex.haskell
  ace_mode: text
  codemirror_mode: haskell-literate
  codemirror_mime_type: text/x-literate-haskell
  language_id: 207
LiveScript:
  type: programming
  color: "#499886"
  aliases:
  - live-script
  - ls
  extensions:
  - ".ls"
  - "._ls"
  filenames:
  - Slakefile
  ace_mode: livescript
  codemirror_mode: livescript
  codemirror_mime_type: text/x-livescript
  language_id: 208
Logos:
  type: programming
  extensions:
  - ".xm"
  - ".x"
  - ".xi"
  ace_mode: text
  tm_scope: source.logos
  language_id: 209
Logtalk:
  type: programming
  extensions:
  - ".lgt"
  - ".logtalk"
  ace_mode: text
  language_id: 210
LookML:
  type: programming
  ace_mode: yaml
  codemirror_mode: yaml
  codemirror_mime_type: text/x-yaml
  color: "#652B81"
  extensions:
  - ".lookml"
  - ".model.lkml"
  - ".view.lkml"
  tm_scope: source.yaml
  language_id: 211
LoomScript:
  type: programming
  extensions:
  - ".ls"
  tm_scope: source.loomscript
  ace_mode: text
  language_id: 212
Lua:
  type: programming
  ace_mode: lua
  codemirror_mode: lua
  codemirror_mime_type: text/x-lua
  color: "#000080"
  extensions:
  - ".lua"
  - ".fcgi"
  - ".nse"
  - ".pd_lua"
  - ".rbxs"
  - ".wlua"
  interpreters:
  - lua
  language_id: 213
M:
  type: programming
  aliases:
  - mumps
  extensions:
  - ".mumps"
  - ".m"
  ace_mode: text
  codemirror_mode: mumps
  codemirror_mime_type: text/x-mumps
  language_id: 214
  tm_scope: none
M4:
  type: programming
  extensions:
  - ".m4"
  tm_scope: none
  ace_mode: text
  language_id: 215
M4Sugar:
  type: programming
  group: M4
  aliases:
  - autoconf
  extensions:
  - ".m4"
  filenames:
  - configure.ac
  tm_scope: none
  ace_mode: text
  language_id: 216
MAXScript:
  type: programming
  color: "#00a6a6"
  extensions:
  - ".ms"
  - ".mcr"
  tm_scope: source.maxscript
  ace_mode: text
  language_id: 217
MQL4:
  type: programming
  color: "#62A8D6"
  extensions:
  - ".mq4"
  - ".mqh"
  tm_scope: source.mql5
  ace_mode: c_cpp
  language_id: 426
MQL5:
  type: programming
  color: "#4A76B8"
  extensions:
  - ".mq5"
  - ".mqh"
  tm_scope: source.mql5
  ace_mode: c_cpp
  language_id: 427
MTML:
  type: markup
  color: "#b7e1f4"
  extensions:
  - ".mtml"
  tm_scope: text.html.basic
  ace_mode: html
  codemirror_mode: htmlmixed
  codemirror_mime_type: text/html
  language_id: 218
MUF:
  type: programming
  group: Forth
  extensions:
  - ".muf"
  - ".m"
  tm_scope: none
  ace_mode: forth
  codemirror_mode: forth
  codemirror_mime_type: text/x-forth
  language_id: 219
Makefile:
  type: programming
  color: "#427819"
  aliases:
  - bsdmake
  - make
  - mf
  extensions:
  - ".mak"
  - ".d"
  - ".make"
  - ".mk"
  - ".mkfile"
  filenames:
  - BSDmakefile
  - GNUmakefile
  - Kbuild
  - Makefile
  - Makefile.am
  - Makefile.boot
  - Makefile.frag
  - Makefile.in
  - Makefile.inc
  - Makefile.wat
  - makefile
  - makefile.sco
  - mkfile
  interpreters:
  - make
  ace_mode: makefile
  codemirror_mode: cmake
  codemirror_mime_type: text/x-cmake
  language_id: 220
Mako:
  type: programming
  extensions:
  - ".mako"
  - ".mao"
  tm_scope: text.html.mako
  ace_mode: text
  language_id: 221
Markdown:
  type: prose
  aliases:
  - pandoc
  ace_mode: markdown
  codemirror_mode: gfm
  codemirror_mime_type: text/x-gfm
  wrap: true
  extensions:
  - ".md"
  - ".markdown"
  - ".mdown"
  - ".mdwn"
  - ".mkd"
  - ".mkdn"
  - ".mkdown"
  - ".ron"
  - ".workbook"
  tm_scope: source.gfm
  language_id: 222
Marko:
  group: HTML
  type: markup
  tm_scope: text.marko
  extensions:
  - ".marko"
  aliases:
  - markojs
  ace_mode: text
  codemirror_mode: htmlmixed
  codemirror_mime_type: text/html
  language_id: 932782397
Mask:
  type: markup
  color: "#f97732"
  ace_mode: mask
  extensions:
  - ".mask"
  tm_scope: source.mask
  language_id: 223
Mathematica:
  type: programming
  extensions:
  - ".mathematica"
  - ".cdf"
  - ".m"
  - ".ma"
  - ".mt"
  - ".nb"
  - ".nbp"
  - ".wl"
  - ".wlt"
  aliases:
  - mma
  ace_mode: text
  codemirror_mode: mathematica
  codemirror_mime_type: text/x-mathematica
  language_id: 224
Matlab:
  type: programming
  color: "#e16737"
  aliases:
  - octave
  extensions:
  - ".matlab"
  - ".m"
  ace_mode: matlab
  codemirror_mode: octave
  codemirror_mime_type: text/x-octave
  language_id: 225
Maven POM:
  type: data
  tm_scope: text.xml.pom
  filenames:
  - pom.xml
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  language_id: 226
Max:
  type: programming
  color: "#c4a79c"
  aliases:
  - max/msp
  - maxmsp
  extensions:
  - ".maxpat"
  - ".maxhelp"
  - ".maxproj"
  - ".mxt"
  - ".pat"
  tm_scope: source.json
  ace_mode: json
  codemirror_mode: javascript
  codemirror_mime_type: application/json
  language_id: 227
MediaWiki:
  type: prose
  wrap: true
  extensions:
  - ".mediawiki"
  - ".wiki"
  tm_scope: text.html.mediawiki
  ace_mode: text
  language_id: 228
Mercury:
  type: programming
  color: "#ff2b2b"
  ace_mode: prolog
  interpreters:
  - mmi
  extensions:
  - ".m"
  - ".moo"
  tm_scope: source.mercury
  language_id: 229
Meson:
  type: programming
  color: "#007800"
  filenames:
  - meson.build
  - meson_options.txt
  tm_scope: source.meson
  ace_mode: text
  language_id: 799141244
Metal:
  type: programming
  color: "#8f14e9"
  extensions:
  - ".metal"
  tm_scope: source.c++
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-c++src
  language_id: 230
MiniD:
  type: programming
  searchable: false
  extensions:
  - ".minid"
  tm_scope: none
  ace_mode: text
  language_id: 231
Mirah:
  type: programming
  color: "#c7a938"
  extensions:
  - ".druby"
  - ".duby"
  - ".mir"
  - ".mirah"
  tm_scope: source.ruby
  ace_mode: ruby
  codemirror_mode: ruby
  codemirror_mime_type: text/x-ruby
  language_id: 232
Modelica:
  type: programming
  extensions:
  - ".mo"
  tm_scope: source.modelica
  ace_mode: text
  codemirror_mode: modelica
  codemirror_mime_type: text/x-modelica
  language_id: 233
Modula-2:
  type: programming
  extensions:
  - ".mod"
  tm_scope: source.modula2
  ace_mode: text
  language_id: 234
Module Management System:
  type: programming
  extensions:
  - ".mms"
  - ".mmk"
  filenames:
  - descrip.mmk
  - descrip.mms
  tm_scope: none
  ace_mode: text
  language_id: 235
Monkey:
  type: programming
  extensions:
  - ".monkey"
  - ".monkey2"
  ace_mode: text
  tm_scope: source.monkey
  language_id: 236
Moocode:
  type: programming
  extensions:
  - ".moo"
  tm_scope: none
  ace_mode: text
  language_id: 237
MoonScript:
  type: programming
  extensions:
  - ".moon"
  interpreters:
  - moon
  ace_mode: text
  language_id: 238
Myghty:
  type: programming
  extensions:
  - ".myt"
  tm_scope: none
  ace_mode: text
  language_id: 239
NCL:
  type: programming
  color: "#28431f"
  extensions:
  - ".ncl"
  tm_scope: source.ncl
  ace_mode: text
  language_id: 240
NL:
  type: data
  extensions:
  - ".nl"
  tm_scope: none
  ace_mode: text
  language_id: 241
NSIS:
  type: programming
  extensions:
  - ".nsi"
  - ".nsh"
  ace_mode: text
  codemirror_mode: nsis
  codemirror_mime_type: text/x-nsis
  language_id: 242
Nearley:
  type: programming
  ace_mode: text
  color: "#990000"
  extensions:
  - ".ne"
  - ".nearley"
  tm_scope: source.ne
  language_id: 521429430
Nemerle:
  type: programming
  color: "#3d3c6e"
  extensions:
  - ".n"
  ace_mode: text
  language_id: 243
NetLinx:
  type: programming
  color: "#0aa0ff"
  extensions:
  - ".axs"
  - ".axi"
  tm_scope: source.netlinx
  ace_mode: text
  language_id: 244
NetLinx+ERB:
  type: programming
  color: "#747faa"
  extensions:
  - ".axs.erb"
  - ".axi.erb"
  tm_scope: source.netlinx.erb
  ace_mode: text
  language_id: 245
NetLogo:
  type: programming
  color: "#ff6375"
  extensions:
  - ".nlogo"
  tm_scope: source.lisp
  ace_mode: lisp
  codemirror_mode: commonlisp
  codemirror_mime_type: text/x-common-lisp
  language_id: 246
NewLisp:
  type: programming
  lexer: NewLisp
  color: "#87AED7"
  extensions:
  - ".nl"
  - ".lisp"
  - ".lsp"
  interpreters:
  - newlisp
  tm_scope: source.lisp
  ace_mode: lisp
  codemirror_mode: commonlisp
  codemirror_mime_type: text/x-common-lisp
  language_id: 247
Nextflow:
  type: programming
  ace_mode: groovy
  tm_scope: source.nextflow
  color: "#3ac486"
  extensions:
  - ".nf"
  filenames:
  - "nextflow.config"  
  interpreters:
  - nextflow
  language_id: 506780613
Nginx:
  type: data
  extensions:
  - ".nginxconf"
  - ".vhost"
  filenames:
  - nginx.conf
  tm_scope: source.nginx
  aliases:
  - nginx configuration file
  ace_mode: text
  codemirror_mode: nginx
  codemirror_mime_type: text/x-nginx-conf
  language_id: 248
Nim:
  type: programming
  color: "#37775b"
  extensions:
  - ".nim"
  - ".nimrod"
  ace_mode: text
  tm_scope: source.nim
  language_id: 249
Ninja:
  type: data
  tm_scope: source.ninja
  extensions:
  - ".ninja"
  ace_mode: text
  language_id: 250
Nit:
  type: programming
  color: "#009917"
  extensions:
  - ".nit"
  tm_scope: source.nit
  ace_mode: text
  language_id: 251
Nix:
  type: programming
  color: "#7e7eff"
  extensions:
  - ".nix"
  aliases:
  - nixos
  tm_scope: source.nix
  ace_mode: nix
  language_id: 252
Nu:
  type: programming
  color: "#c9df40"
  aliases:
  - nush
  extensions:
  - ".nu"
  filenames:
  - Nukefile
  tm_scope: source.nu
  ace_mode: scheme
  codemirror_mode: scheme
  codemirror_mime_type: text/x-scheme
  interpreters:
  - nush
  language_id: 253
NumPy:
  type: programming
  group: Python
  extensions:
  - ".numpy"
  - ".numpyw"
  - ".numsc"
  tm_scope: none
  ace_mode: text
  codemirror_mode: python
  codemirror_mime_type: text/x-python
  language_id: 254
OCaml:
  type: programming
  ace_mode: ocaml
  codemirror_mode: mllike
  codemirror_mime_type: text/x-ocaml
  color: "#3be133"
  extensions:
  - ".ml"
  - ".eliom"
  - ".eliomi"
  - ".ml4"
  - ".mli"
  - ".mll"
  - ".mly"
  interpreters:
  - ocaml
  - ocamlrun
  - ocamlscript
  tm_scope: source.ocaml
  language_id: 255
ObjDump:
  type: data
  extensions:
  - ".objdump"
  tm_scope: objdump.x86asm
  ace_mode: assembly_x86
  language_id: 256
Objective-C:
  type: programming
  tm_scope: source.objc
  color: "#438eff"
  aliases:
  - obj-c
  - objc
  - objectivec
  extensions:
  - ".m"
  - ".h"
  ace_mode: objectivec
  codemirror_mode: clike
  codemirror_mime_type: text/x-objectivec
  language_id: 257
Objective-C++:
  type: programming
  tm_scope: source.objc++
  color: "#6866fb"
  aliases:
  - obj-c++
  - objc++
  - objectivec++
  extensions:
  - ".mm"
  ace_mode: objectivec
  codemirror_mode: clike
  codemirror_mime_type: text/x-objectivec
  language_id: 258
Objective-J:
  type: programming
  color: "#ff0c5a"
  aliases:
  - obj-j
  - objectivej
  - objj
  extensions:
  - ".j"
  - ".sj"
  tm_scope: source.js.objj
  ace_mode: text
  language_id: 259
Omgrofl:
  type: programming
  extensions:
  - ".omgrofl"
  color: "#cabbff"
  tm_scope: none
  ace_mode: text
  language_id: 260
Opa:
  type: programming
  extensions:
  - ".opa"
  ace_mode: text
  language_id: 261
Opal:
  type: programming
  color: "#f7ede0"
  extensions:
  - ".opal"
  tm_scope: source.opal
  ace_mode: text
  language_id: 262
OpenCL:
  type: programming
  group: C
  extensions:
  - ".cl"
  - ".opencl"
  tm_scope: source.c
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csrc
  language_id: 263
OpenEdge ABL:
  type: programming
  aliases:
  - progress
  - openedge
  - abl
  extensions:
  - ".p"
  - ".cls"
  - ".w"
  tm_scope: source.abl
  ace_mode: text
  language_id: 264
OpenRC runscript:
  type: programming
  group: Shell
  aliases:
  - openrc
  interpreters:
  - openrc-run
  tm_scope: source.shell
  ace_mode: sh
  codemirror_mode: shell
  codemirror_mime_type: text/x-sh
  language_id: 265
OpenSCAD:
  type: programming
  extensions:
  - ".scad"
  tm_scope: source.scad
  ace_mode: scad
  language_id: 266
OpenType Feature File:
  type: data
  aliases:
  - AFDKO
  extensions:
  - ".fea"
  tm_scope: source.opentype
  ace_mode: text
  language_id: 374317347
Org:
  type: prose
  wrap: true
  extensions:
  - ".org"
  tm_scope: none
  ace_mode: text
  language_id: 267
Ox:
  type: programming
  extensions:
  - ".ox"
  - ".oxh"
  - ".oxo"
  tm_scope: source.ox
  ace_mode: text
  language_id: 268
Oxygene:
  type: programming
  color: "#cdd0e3"
  extensions:
  - ".oxygene"
  tm_scope: none
  ace_mode: text
  language_id: 269
Oz:
  type: programming
  color: "#fab738"
  extensions:
  - ".oz"
  tm_scope: source.oz
  ace_mode: text
  codemirror_mode: oz
  codemirror_mime_type: text/x-oz
  language_id: 270
P4:
  type: programming
  color: "#7055b5"
  extensions:
  - ".p4"
  tm_scope: source.p4
  ace_mode: text
  language_id: 348895984
PAWN:
  type: programming
  color: "#dbb284"
  extensions:
  - ".pwn"
  - ".inc"
  tm_scope: source.pawn
  ace_mode: text
  language_id: 271
PHP:
  type: programming
  tm_scope: text.html.php
  ace_mode: php
  codemirror_mode: php
  codemirror_mime_type: application/x-httpd-php
  color: "#4F5D95"
  extensions:
  - ".php"
  - ".aw"
  - ".ctp"
  - ".fcgi"
  - ".inc"
  - ".php3"
  - ".php4"
  - ".php5"
  - ".phps"
  - ".phpt"
  filenames:
  - ".php_cs"
  - ".php_cs.dist"
  - Phakefile
  interpreters:
  - php
  aliases:
  - inc
  language_id: 272
PLSQL:
  type: programming
  ace_mode: sql
  codemirror_mode: sql
  codemirror_mime_type: text/x-plsql
  tm_scope: none
  color: "#dad8d8"
  extensions:
  - ".pls"
  - ".bdy"
  - ".ddl"
  - ".fnc"
  - ".pck"
  - ".pkb"
  - ".pks"
  - ".plb"
  - ".plsql"
  - ".prc"
  - ".spc"
  - ".sql"
  - ".tpb"
  - ".tps"
  - ".trg"
  - ".vw"
  language_id: 273
PLpgSQL:
  type: programming
  ace_mode: pgsql
  codemirror_mode: sql
  codemirror_mime_type: text/x-sql
  tm_scope: source.sql
  extensions:
  - ".sql"
  language_id: 274
POV-Ray SDL:
  type: programming
  aliases:
  - pov-ray
  - povray
  extensions:
  - ".pov"
  - ".inc"
  ace_mode: text
  language_id: 275
Pan:
  type: programming
  color: "#cc0000"
  extensions:
  - ".pan"
  tm_scope: source.pan
  ace_mode: text
  language_id: 276
Papyrus:
  type: programming
  color: "#6600cc"
  extensions:
  - ".psc"
  tm_scope: source.papyrus.skyrim
  ace_mode: text
  language_id: 277
Parrot:
  type: programming
  color: "#f3ca0a"
  extensions:
  - ".parrot"
  tm_scope: none
  ace_mode: text
  language_id: 278
Parrot Assembly:
  group: Parrot
  type: programming
  aliases:
  - pasm
  extensions:
  - ".pasm"
  interpreters:
  - parrot
  tm_scope: none
  ace_mode: text
  language_id: 279
Parrot Internal Representation:
  group: Parrot
  tm_scope: source.parrot.pir
  type: programming
  aliases:
  - pir
  extensions:
  - ".pir"
  interpreters:
  - parrot
  ace_mode: text
  language_id: 280
Pascal:
  type: programming
  color: "#E3F171"
  extensions:
  - ".pas"
  - ".dfm"
  - ".dpr"
  - ".inc"
  - ".lpr"
  - ".pascal"
  - ".pp"
  interpreters:
  - instantfpc
  ace_mode: pascal
  codemirror_mode: pascal
  codemirror_mime_type: text/x-pascal
  language_id: 281
Pep8:
  type: programming
  color: "#C76F5B"
  extensions:
  - ".pep"
  ace_mode: text
  tm_scope: source.pep8
  language_id: 840372442
Perl:
  type: programming
  tm_scope: source.perl
  ace_mode: perl
  codemirror_mode: perl
  codemirror_mime_type: text/x-perl
  color: "#0298c3"
  extensions:
  - ".pl"
  - ".al"
  - ".cgi"
  - ".fcgi"
  - ".perl"
  - ".ph"
  - ".plx"
  - ".pm"
  - ".psgi"
  - ".t"
  filenames:
  - cpanfile
  interpreters:
  - perl
  language_id: 282
Perl 6:
  type: programming
  color: "#0000fb"
  extensions:
  - ".6pl"
  - ".6pm"
  - ".nqp"
  - ".p6"
  - ".p6l"
  - ".p6m"
  - ".pl"
  - ".pl6"
  - ".pm"
  - ".pm6"
  - ".t"
  filenames:
  - Rexfile
  interpreters:
  - perl6
  aliases:
  - perl6
  tm_scope: source.perl6fe
  ace_mode: perl
  codemirror_mode: perl
  codemirror_mime_type: text/x-perl
  language_id: 283
Pic:
  type: markup
  group: Roff
  tm_scope: source.pic
  extensions:
  - ".pic"
  - ".chem"
  ace_mode: text
  codemirror_mode: troff
  codemirror_mime_type: text/troff
  language_id: 425
Pickle:
  type: data
  extensions:
  - ".pkl"
  tm_scope: none
  ace_mode: text
  language_id: 284
PicoLisp:
  type: programming
  extensions:
  - ".l"
  interpreters:
  - picolisp
  - pil
  tm_scope: source.lisp
  ace_mode: lisp
  language_id: 285
PigLatin:
  type: programming
  color: "#fcd7de"
  extensions:
  - ".pig"
  tm_scope: source.pig_latin
  ace_mode: text
  language_id: 286
Pike:
  type: programming
  color: "#005390"
  extensions:
  - ".pike"
  - ".pmod"
  interpreters:
  - pike
  ace_mode: text
  language_id: 287
Pod:
  type: prose
  ace_mode: perl
  codemirror_mode: perl
  codemirror_mime_type: text/x-perl
  wrap: true
  extensions:
  - ".pod"
  interpreters:
  - perl
  tm_scope: none
  language_id: 288
PogoScript:
  type: programming
  color: "#d80074"
  extensions:
  - ".pogo"
  tm_scope: source.pogoscript
  ace_mode: text
  language_id: 289
Pony:
  type: programming
  extensions:
  - ".pony"
  tm_scope: source.pony
  ace_mode: text
  language_id: 290
PostCSS:
  type: markup
  tm_scope: source.postcss
  group: CSS
  extensions:
  - ".pcss"
  ace_mode: text
  language_id: 262764437
PostScript:
  type: markup
  color: "#da291c"
  extensions:
  - ".ps"
  - ".eps"
  - ".pfa"
  tm_scope: source.postscript
  aliases:
  - postscr
  ace_mode: text
  language_id: 291
PowerBuilder:
  type: programming
  color: "#8f0f8d"
  extensions:
  - ".pbt"
  - ".sra"
  - ".sru"
  - ".srw"
  tm_scope: none
  ace_mode: text
  language_id: 292
PowerShell:
  type: programming
  color: "#012456"
  ace_mode: powershell
  codemirror_mode: powershell
  codemirror_mime_type: application/x-powershell
  aliases:
  - posh
  extensions:
  - ".ps1"
  - ".psd1"
  - ".psm1"
  language_id: 293
Processing:
  type: programming
  color: "#0096D8"
  extensions:
  - ".pde"
  ace_mode: text
  language_id: 294
Prolog:
  type: programming
  color: "#74283c"
  extensions:
  - ".pl"
  - ".pro"
  - ".prolog"
  - ".yap"
  interpreters:
  - swipl
  - yap
  tm_scope: source.prolog
  ace_mode: prolog
  language_id: 295
Propeller Spin:
  type: programming
  color: "#7fa2a7"
  extensions:
  - ".spin"
  tm_scope: source.spin
  ace_mode: text
  language_id: 296
Protocol Buffer:
  type: data
  aliases:
  - protobuf
  - Protocol Buffers
  extensions:
  - ".proto"
  tm_scope: source.protobuf
  ace_mode: protobuf
  codemirror_mode: protobuf
  codemirror_mime_type: text/x-protobuf
  language_id: 297
Public Key:
  type: data
  extensions:
  - ".asc"
  - ".pub"
  tm_scope: none
  ace_mode: text
  codemirror_mode: asciiarmor
  codemirror_mime_type: application/pgp
  language_id: 298
Pug:
  group: HTML
  type: markup
  extensions:
  - ".jade"
  - ".pug"
  tm_scope: text.jade
  ace_mode: jade
  codemirror_mode: pug
  codemirror_mime_type: text/x-pug
  language_id: 179
Puppet:
  type: programming
  color: "#302B6D"
  extensions:
  - ".pp"
  filenames:
  - Modulefile
  ace_mode: text
  codemirror_mode: puppet
  codemirror_mime_type: text/x-puppet
  tm_scope: source.puppet
  language_id: 299
Pure Data:
  type: data
  extensions:
  - ".pd"
  tm_scope: none
  ace_mode: text
  language_id: 300
PureBasic:
  type: programming
  color: "#5a6986"
  extensions:
  - ".pb"
  - ".pbi"
  tm_scope: none
  ace_mode: text
  language_id: 301
PureScript:
  type: programming
  color: "#1D222D"
  extensions:
  - ".purs"
  tm_scope: source.purescript
  ace_mode: haskell
  codemirror_mode: haskell
  codemirror_mime_type: text/x-haskell
  language_id: 302
Python:
  type: programming
  ace_mode: python
  codemirror_mode: python
  codemirror_mime_type: text/x-python
  color: "#3572A5"
  extensions:
  - ".py"
  - ".bzl"
  - ".cgi"
  - ".fcgi"
  - ".gyp"
  - ".gypi"
  - ".lmi"
  - ".py3"
  - ".pyde"
  - ".pyi"
  - ".pyp"
  - ".pyt"
  - ".pyw"
  - ".rpy"
  - ".spec"
  - ".tac"
  - ".wsgi"
  - ".xpy"
  filenames:
  - ".gclient"
  - BUCK
  - BUILD
  - BUILD.bazel
  - SConscript
  - SConstruct
  - Snakefile
  - WORKSPACE
  - wscript
  interpreters:
  - python
  - python2
  - python3
  aliases:
  - rusthon
  language_id: 303
Python console:
  type: programming
  group: Python
  searchable: false
  aliases:
  - pycon
  tm_scope: text.python.console
  ace_mode: text
  language_id: 428
Python traceback:
  type: data
  group: Python
  searchable: false
  extensions:
  - ".pytb"
  tm_scope: text.python.traceback
  ace_mode: text
  language_id: 304
QML:
  type: programming
  color: "#44a51c"
  extensions:
  - ".qml"
  - ".qbs"
  tm_scope: source.qml
  ace_mode: text
  language_id: 305
QMake:
  type: programming
  extensions:
  - ".pro"
  - ".pri"
  interpreters:
  - qmake
  ace_mode: text
  language_id: 306
R:
  type: programming
  color: "#198CE7"
  aliases:
  - R
  - Rscript
  - splus
  extensions:
  - ".r"
  - ".rd"
  - ".rsx"
  filenames:
  - ".Rprofile"
  interpreters:
  - Rscript
  ace_mode: r
  codemirror_mode: r
  codemirror_mime_type: text/x-rsrc
  language_id: 307
RAML:
  type: markup
  ace_mode: yaml
  codemirror_mode: yaml
  codemirror_mime_type: text/x-yaml
  tm_scope: source.yaml
  color: "#77d9fb"
  extensions:
  - ".raml"
  language_id: 308
RDoc:
  type: prose
  ace_mode: rdoc
  wrap: true
  extensions:
  - ".rdoc"
  tm_scope: text.rdoc
  language_id: 309
REALbasic:
  type: programming
  extensions:
  - ".rbbas"
  - ".rbfrm"
  - ".rbmnu"
  - ".rbres"
  - ".rbtbar"
  - ".rbuistate"
  tm_scope: source.vbnet
  ace_mode: text
  language_id: 310
REXX:
  type: programming
  aliases:
  - arexx
  extensions:
  - ".rexx"
  - ".pprx"
  - ".rex"
  interpreters:
  - regina
  - rexx
  tm_scope: source.rexx
  ace_mode: text
  language_id: 311
RHTML:
  type: markup
  group: HTML
  extensions:
  - ".rhtml"
  tm_scope: text.html.erb
  aliases:
  - html+ruby
  ace_mode: rhtml
  codemirror_mode: htmlembedded
  codemirror_mime_type: application/x-erb
  language_id: 312
RMarkdown:
  type: prose
  wrap: true
  ace_mode: markdown
  codemirror_mode: gfm
  codemirror_mime_type: text/x-gfm
  extensions:
  - ".rmd"
  tm_scope: source.gfm
  language_id: 313
RPM Spec:
  type: data
  tm_scope: source.rpm-spec
  extensions:
  - ".spec"
  aliases:
  - specfile
  ace_mode: text
  codemirror_mode: rpm
  codemirror_mime_type: text/x-rpm-spec
  language_id: 314
RUNOFF:
  type: markup
  color: "#665a4e"
  extensions:
  - ".rnh"
  - ".rno"
  tm_scope: text.runoff
  ace_mode: text
  language_id: 315
Racket:
  type: programming
  color: "#22228f"
  extensions:
  - ".rkt"
  - ".rktd"
  - ".rktl"
  - ".scrbl"
  interpreters:
  - racket
  tm_scope: source.racket
  ace_mode: lisp
  language_id: 316
Ragel:
  type: programming
  color: "#9d5200"
  extensions:
  - ".rl"
  aliases:
  - ragel-rb
  - ragel-ruby
  tm_scope: none
  ace_mode: text
  language_id: 317
Rascal:
  type: programming
  color: "#fffaa0"
  extensions:
  - ".rsc"
  tm_scope: source.rascal
  ace_mode: text
  language_id: 173616037
Raw token data:
  type: data
  aliases:
  - raw
  extensions:
  - ".raw"
  tm_scope: none
  ace_mode: text
  language_id: 318
Reason:
  type: programming
  group: OCaml
  ace_mode: rust
  codemirror_mode: rust
  codemirror_mime_type: text/x-rustsrc
  extensions:
  - ".re"
  - ".rei"
  interpreters:
  - ocaml
  tm_scope: source.reason
  language_id: 869538413
Rebol:
  type: programming
  color: "#358a5b"
  extensions:
  - ".reb"
  - ".r"
  - ".r2"
  - ".r3"
  - ".rebol"
  ace_mode: text
  tm_scope: source.rebol
  language_id: 319
Red:
  type: programming
  color: "#f50000"
  extensions:
  - ".red"
  - ".reds"
  aliases:
  - red/system
  tm_scope: source.red
  ace_mode: text
  language_id: 320
Redcode:
  type: programming
  extensions:
  - ".cw"
  tm_scope: none
  ace_mode: text
  language_id: 321
Regular Expression:
  type: data
  extensions:
  - ".regexp"
  - ".regex"
  aliases:
  - regexp
  - regex
  ace_mode: text
  tm_scope: source.regexp
  language_id: 363378884
Ren'Py:
  type: programming
  aliases:
  - renpy
  color: "#ff7f7f"
  extensions:
  - ".rpy"
  tm_scope: source.renpy
  ace_mode: python
  language_id: 322
RenderScript:
  type: programming
  extensions:
  - ".rs"
  - ".rsh"
  tm_scope: none
  ace_mode: text
  language_id: 323
Ring:
  type: programming
  color: "#0e60e3"
  extensions:
  - ".ring"
  tm_scope: source.ring
  ace_mode: text
  language_id: 431
RobotFramework:
  type: programming
  extensions:
  - ".robot"
  tm_scope: text.robot
  ace_mode: text
  language_id: 324
Roff:
  type: markup
  color: "#ecdebe"
  extensions:
  - ".man"
  - ".1"
  - ".1in"
  - ".1m"
  - ".1x"
  - ".2"
  - ".3"
  - ".3in"
  - ".3m"
  - ".3qt"
  - ".3x"
  - ".4"
  - ".5"
  - ".6"
  - ".7"
  - ".8"
  - ".9"
  - ".l"
  - ".me"
  - ".ms"
  - ".n"
  - ".nr"
  - ".rno"
  - ".roff"
  - ".tmac"
  filenames:
  - mmn
  - mmt
  tm_scope: text.roff
  aliases:
  - nroff
  ace_mode: text
  codemirror_mode: troff
  codemirror_mime_type: text/troff
  language_id: 141
Rouge:
  type: programming
  ace_mode: clojure
  codemirror_mode: clojure
  codemirror_mime_type: text/x-clojure
  color: "#cc0088"
  extensions:
  - ".rg"
  tm_scope: source.clojure
  language_id: 325
Ruby:
  type: programming
  ace_mode: ruby
  codemirror_mode: ruby
  codemirror_mime_type: text/x-ruby
  color: "#701516"
  aliases:
  - jruby
  - macruby
  - rake
  - rb
  - rbx
  extensions:
  - ".rb"
  - ".builder"
  - ".eye"
  - ".fcgi"
  - ".gemspec"
  - ".god"
  - ".jbuilder"
  - ".mspec"
  - ".pluginspec"
  - ".podspec"
  - ".rabl"
  - ".rake"
  - ".rbuild"
  - ".rbw"
  - ".rbx"
  - ".ru"
  - ".ruby"
  - ".spec"
  - ".thor"
  - ".watchr"
  interpreters:
  - ruby
  - macruby
  - rake
  - jruby
  - rbx
  filenames:
  - ".irbrc"
  - ".pryrc"
  - Appraisals
  - Berksfile
  - Brewfile
  - Buildfile
  - Dangerfile
  - Deliverfile
  - Fastfile
  - Gemfile
  - Gemfile.lock
  - Guardfile
  - Jarfile
  - Mavenfile
  - Podfile
  - Puppetfile
  - Rakefile
  - Snapfile
  - Thorfile
  - Vagrantfile
  - buildfile
  language_id: 326
Rust:
  type: programming
  color: "#dea584"
  extensions:
  - ".rs"
  - ".rs.in"
  ace_mode: rust
  codemirror_mode: rust
  codemirror_mime_type: text/x-rustsrc
  language_id: 327
SAS:
  type: programming
  color: "#B34936"
  extensions:
  - ".sas"
  tm_scope: source.sas
  ace_mode: text
  codemirror_mode: sas
  codemirror_mime_type: text/x-sas
  language_id: 328
SCSS:
  type: markup
  tm_scope: source.scss
  group: CSS
  ace_mode: scss
  codemirror_mode: css
  codemirror_mime_type: text/x-scss
  extensions:
  - ".scss"
  language_id: 329
SMT:
  type: programming
  extensions:
  - ".smt2"
  - ".smt"
  interpreters:
  - boolector
  - cvc4
  - mathsat5
  - opensmt
  - smtinterpol
  - smt-rat
  - stp
  - verit
  - yices2
  - z3
  tm_scope: source.smt
  ace_mode: text
  language_id: 330
SPARQL:
  type: data
  tm_scope: source.sparql
  ace_mode: text
  codemirror_mode: sparql
  codemirror_mime_type: application/sparql-query
  extensions:
  - ".sparql"
  - ".rq"
  language_id: 331
SQF:
  type: programming
  color: "#3F3F3F"
  extensions:
  - ".sqf"
  - ".hqf"
  tm_scope: source.sqf
  ace_mode: text
  language_id: 332
SQL:
  type: data
  tm_scope: source.sql
  ace_mode: sql
  codemirror_mode: sql
  codemirror_mime_type: text/x-sql
  extensions:
  - ".sql"
  - ".cql"
  - ".ddl"
  - ".inc"
  - ".mysql"
  - ".prc"
  - ".tab"
  - ".udf"
  - ".viw"
  language_id: 333
SQLPL:
  type: programming
  ace_mode: sql
  codemirror_mode: sql
  codemirror_mime_type: text/x-sql
  tm_scope: source.sql
  extensions:
  - ".sql"
  - ".db2"
  language_id: 334
SRecode Template:
  type: markup
  color: "#348a34"
  tm_scope: source.lisp
  ace_mode: lisp
  codemirror_mode: commonlisp
  codemirror_mime_type: text/x-common-lisp
  extensions:
  - ".srt"
  language_id: 335
STON:
  type: data
  group: Smalltalk
  extensions:
  - ".ston"
  tm_scope: source.smalltalk
  ace_mode: text
  language_id: 336
SVG:
  type: data
  extensions:
  - ".svg"
  tm_scope: text.xml
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  language_id: 337
Sage:
  type: programming
  group: Python
  extensions:
  - ".sage"
  - ".sagews"
  tm_scope: source.python
  ace_mode: python
  codemirror_mode: python
  codemirror_mime_type: text/x-python
  language_id: 338
SaltStack:
  type: programming
  color: "#646464"
  aliases:
  - saltstate
  - salt
  extensions:
  - ".sls"
  tm_scope: source.yaml.salt
  ace_mode: yaml
  codemirror_mode: yaml
  codemirror_mime_type: text/x-yaml
  language_id: 339
Sass:
  type: markup
  tm_scope: source.sass
  group: CSS
  extensions:
  - ".sass"
  ace_mode: sass
  codemirror_mode: sass
  codemirror_mime_type: text/x-sass
  language_id: 340
Scala:
  type: programming
  ace_mode: scala
  codemirror_mode: clike
  codemirror_mime_type: text/x-scala
  color: "#c22d40"
  extensions:
  - ".scala"
  - ".kojo"  
  - ".sbt"
  - ".sc"
  interpreters:
  - scala
  language_id: 341
Scaml:
  group: HTML
  type: markup
  extensions:
  - ".scaml"
  tm_scope: source.scaml
  ace_mode: text
  language_id: 342
Scheme:
  type: programming
  color: "#1e4aec"
  extensions:
  - ".scm"
  - ".sch"
  - ".sld"
  - ".sls"
  - ".sps"
  - ".ss"
  interpreters:
  - guile
  - bigloo
  - chicken
  - csi
  - gosh
  - r6rs
  ace_mode: scheme
  codemirror_mode: scheme
  codemirror_mime_type: text/x-scheme
  language_id: 343
Scilab:
  type: programming
  extensions:
  - ".sci"
  - ".sce"
  - ".tst"
  ace_mode: text
  language_id: 344
Self:
  type: programming
  color: "#0579aa"
  extensions:
  - ".self"
  tm_scope: none
  ace_mode: text
  language_id: 345
ShaderLab:
  type: programming
  extensions:
  - ".shader"
  ace_mode: text
  tm_scope: source.shaderlab
  language_id: 664257356
Shell:
  type: programming
  color: "#89e051"
  aliases:
  - sh
  - shell-script
  - bash
  - zsh
  extensions:
  - ".sh"
  - ".bash"
  - ".bats"
  - ".cgi"
  - ".command"
  - ".fcgi"
  - ".ksh"
  - ".sh.in"
  - ".tmux"
  - ".tool"
  - ".zsh"
  filenames:
  - ".bash_history"
  - ".bash_logout"
  - ".bash_profile"
  - ".bashrc"
  - PKGBUILD
  - gradlew
  interpreters:
  - ash
  - bash
  - dash
  - ksh
  - mksh
  - pdksh
  - rc
  - sh
  - zsh
  ace_mode: sh
  codemirror_mode: shell
  codemirror_mime_type: text/x-sh
  language_id: 346
ShellSession:
  type: programming
  extensions:
  - ".sh-session"
  aliases:
  - bash session
  - console
  tm_scope: text.shell-session
  ace_mode: sh
  codemirror_mode: shell
  codemirror_mime_type: text/x-sh
  language_id: 347
Shen:
  type: programming
  color: "#120F14"
  extensions:
  - ".shen"
  tm_scope: source.shen
  ace_mode: text
  language_id: 348
Slash:
  type: programming
  color: "#007eff"
  extensions:
  - ".sl"
  tm_scope: text.html.slash
  ace_mode: text
  language_id: 349
Slim:
  group: HTML
  type: markup
  extensions:
  - ".slim"
  tm_scope: text.slim
  ace_mode: text
  codemirror_mode: slim
  codemirror_mime_type: text/x-slim
  language_id: 350
Smali:
  type: programming
  extensions:
  - ".smali"
  ace_mode: text
  tm_scope: source.smali
  language_id: 351
Smalltalk:
  type: programming
  color: "#596706"
  extensions:
  - ".st"
  - ".cs"
  aliases:
  - squeak
  ace_mode: text
  codemirror_mode: smalltalk
  codemirror_mime_type: text/x-stsrc
  language_id: 352
Smarty:
  type: programming
  extensions:
  - ".tpl"
  ace_mode: smarty
  codemirror_mode: smarty
  codemirror_mime_type: text/x-smarty
  tm_scope: text.html.smarty
  language_id: 353
Solidity:
  type: programming
  color: "#AA6746"
  ace_mode: text
  tm_scope: source.solidity
  language_id: 237469032
SourcePawn:
  type: programming
  color: "#5c7611"
  aliases:
  - sourcemod
  extensions:
  - ".sp"
  - ".inc"
  - ".sma"
  tm_scope: source.sp
  ace_mode: text
  language_id: 354
Spline Font Database:
  type: data
  extensions:
  - ".sfd"
  tm_scope: text.sfd
  ace_mode: yaml
  language_id: 767169629
Squirrel:
  type: programming
  color: "#800000"
  extensions:
  - ".nut"
  tm_scope: source.c++
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-c++src
  language_id: 355
Stan:
  type: programming
  color: "#b2011d"
  extensions:
  - ".stan"
  ace_mode: text
  tm_scope: source.stan
  language_id: 356
Standard ML:
  type: programming
  color: "#dc566d"
  aliases:
  - sml
  extensions:
  - ".ML"
  - ".fun"
  - ".sig"
  - ".sml"
  tm_scope: source.ml
  ace_mode: text
  codemirror_mode: mllike
  codemirror_mime_type: text/x-ocaml
  language_id: 357
Stata:
  type: programming
  extensions:
  - ".do"
  - ".ado"
  - ".doh"
  - ".ihlp"
  - ".mata"
  - ".matah"
  - ".sthlp"
  ace_mode: text
  language_id: 358
Stylus:
  type: markup
  group: CSS
  extensions:
  - ".styl"
  tm_scope: source.stylus
  ace_mode: stylus
  language_id: 359
SubRip Text:
  type: data
  extensions:
  - ".srt"
  ace_mode: text
  tm_scope: text.srt
  language_id: 360
Sublime Text Config:
  type: data
  group: JSON
  tm_scope: source.js
  ace_mode: javascript
  codemirror_mode: javascript
  codemirror_mime_type: text/javascript
  extensions:
  - ".sublime-build"
  - ".sublime-commands"
  - ".sublime-completions"
  - ".sublime-keymap"
  - ".sublime-macro"
  - ".sublime-menu"
  - ".sublime-mousemap"
  - ".sublime-project"
  - ".sublime-settings"
  - ".sublime-theme"
  - ".sublime-workspace"
  - ".sublime_metrics"
  - ".sublime_session"
  language_id: 423
SugarSS:
  type: markup
  tm_scope: source.css.postcss.sugarss
  group: CSS
  extensions:
  - ".sss"
  ace_mode: text
  language_id: 826404698
SuperCollider:
  type: programming
  color: "#46390b"
  extensions:
  - ".sc"
  - ".scd"
  interpreters:
  - sclang
  - scsynth
  tm_scope: source.supercollider
  ace_mode: text
  language_id: 361
Swift:
  type: programming
  color: "#ffac45"
  extensions:
  - ".swift"
  ace_mode: text
  codemirror_mode: swift
  codemirror_mime_type: text/x-swift
  language_id: 362
SystemVerilog:
  type: programming
  color: "#DAE1C2"
  extensions:
  - ".sv"
  - ".svh"
  - ".vh"
  ace_mode: verilog
  codemirror_mode: verilog
  codemirror_mime_type: text/x-systemverilog
  language_id: 363
TI Program:
  type: programming
  ace_mode: text
  color: "#A0AA87"
  extensions:
  - ".8xp"
  - ".8xk"
  - ".8xk.txt"
  - ".8xp.txt"
  language_id: 422
  tm_scope: none
TLA:
  type: programming
  extensions:
  - ".tla"
  tm_scope: source.tla
  ace_mode: text
  language_id: 364
TOML:
  type: data
  extensions:
  - ".toml"
  tm_scope: source.toml
  ace_mode: toml
  codemirror_mode: toml
  codemirror_mime_type: text/x-toml
  language_id: 365
TXL:
  type: programming
  extensions:
  - ".txl"
  tm_scope: source.txl
  ace_mode: text
  language_id: 366
Tcl:
  type: programming
  color: "#e4cc98"
  extensions:
  - ".tcl"
  - ".adp"
  - ".tm"
  interpreters:
  - tclsh
  - wish
  ace_mode: tcl
  codemirror_mode: tcl
  codemirror_mime_type: text/x-tcl
  language_id: 367
Tcsh:
  type: programming
  group: Shell
  extensions:
  - ".tcsh"
  - ".csh"
  tm_scope: source.shell
  ace_mode: sh
  codemirror_mode: shell
  codemirror_mime_type: text/x-sh
  language_id: 368
TeX:
  type: markup
  color: "#3D6117"
  ace_mode: tex
  codemirror_mode: stex
  codemirror_mime_type: text/x-stex
  wrap: true
  aliases:
  - latex
  extensions:
  - ".tex"
  - ".aux"
  - ".bbx"
  - ".bib"
  - ".cbx"
  - ".cls"
  - ".dtx"
  - ".ins"
  - ".lbx"
  - ".ltx"
  - ".mkii"
  - ".mkiv"
  - ".mkvi"
  - ".sty"
  - ".toc"
  language_id: 369
Tea:
  type: markup
  extensions:
  - ".tea"
  tm_scope: source.tea
  ace_mode: text
  language_id: 370
Terra:
  type: programming
  extensions:
  - ".t"
  color: "#00004c"
  ace_mode: lua
  codemirror_mode: lua
  codemirror_mime_type: text/x-lua
  interpreters:
  - lua
  language_id: 371
Text:
  type: prose
  wrap: true
  aliases:
  - fundamental
  extensions:
  - ".txt"
  - ".fr"
  - ".nb"
  - ".ncl"
  - ".no"
  filenames:
  - COPYING
  - COPYRIGHT.regex
  - FONTLOG
  - INSTALL
  - INSTALL.mysql
  - LICENSE
  - LICENSE.mysql
  - NEWS
  - README.1ST
  - README.me
  - README.mysql
  - click.me
  - delete.me
  - keep.me
  - read.me
  - test.me
  tm_scope: none
  ace_mode: text
  language_id: 372
Textile:
  type: prose
  ace_mode: textile
  codemirror_mode: textile
  codemirror_mime_type: text/x-textile
  wrap: true
  extensions:
  - ".textile"
  tm_scope: none
  language_id: 373
Thrift:
  type: programming
  tm_scope: source.thrift
  extensions:
  - ".thrift"
  ace_mode: text
  language_id: 374
Turing:
  type: programming
  color: "#cf142b"
  extensions:
  - ".t"
  - ".tu"
  tm_scope: source.turing
  ace_mode: text
  language_id: 375
Turtle:
  type: data
  extensions:
  - ".ttl"
  tm_scope: source.turtle
  ace_mode: text
  codemirror_mode: turtle
  codemirror_mime_type: text/turtle
  language_id: 376
Twig:
  type: markup
  group: HTML
  extensions:
  - ".twig"
  tm_scope: text.html.twig
  ace_mode: twig
  codemirror_mode: twig
  codemirror_mime_type: text/x-twig
  language_id: 377
Type Language:
  type: data
  aliases:
  - tl
  extensions:
  - ".tl"
  tm_scope: source.tl
  ace_mode: text
  language_id: 632765617
TypeScript:
  type: programming
  color: "#2b7489"
  aliases:
  - ts
  extensions:
  - ".ts"
  - ".tsx"
  tm_scope: source.ts
  ace_mode: typescript
  codemirror_mode: javascript
  codemirror_mime_type: application/typescript
  language_id: 378
Unified Parallel C:
  type: programming
  group: C
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csrc
  extensions:
  - ".upc"
  tm_scope: source.c
  language_id: 379
Unity3D Asset:
  type: data
  ace_mode: yaml
  codemirror_mode: yaml
  codemirror_mime_type: text/x-yaml
  extensions:
  - ".anim"
  - ".asset"
  - ".mat"
  - ".meta"
  - ".prefab"
  - ".unity"
  tm_scope: source.yaml
  language_id: 380
Unix Assembly:
  type: programming
  group: Assembly
  extensions:
  - ".s"
  - ".ms"
  tm_scope: source.assembly
  ace_mode: assembly_x86
  language_id: 120
Uno:
  type: programming
  extensions:
  - ".uno"
  ace_mode: csharp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csharp
  tm_scope: source.cs
  language_id: 381
UnrealScript:
  type: programming
  color: "#a54c4d"
  extensions:
  - ".uc"
  tm_scope: source.java
  ace_mode: java
  codemirror_mode: clike
  codemirror_mime_type: text/x-java
  language_id: 382
UrWeb:
  type: programming
  aliases:
  - Ur/Web
  - Ur
  extensions:
  - ".ur"
  - ".urs"
  tm_scope: source.ur
  ace_mode: text
  language_id: 383
VCL:
  type: programming
  color: "#0298c3"
  extensions:
  - ".vcl"
  tm_scope: source.varnish.vcl
  ace_mode: text
  language_id: 384
VHDL:
  type: programming
  color: "#adb2cb"
  extensions:
  - ".vhdl"
  - ".vhd"
  - ".vhf"
  - ".vhi"
  - ".vho"
  - ".vhs"
  - ".vht"
  - ".vhw"
  ace_mode: vhdl
  codemirror_mode: vhdl
  codemirror_mime_type: text/x-vhdl
  language_id: 385
Vala:
  type: programming
  color: "#fbe5cd"
  extensions:
  - ".vala"
  - ".vapi"
  ace_mode: vala
  language_id: 386
Verilog:
  type: programming
  color: "#b2b7f8"
  extensions:
  - ".v"
  - ".veo"
  ace_mode: verilog
  codemirror_mode: verilog
  codemirror_mime_type: text/x-verilog
  language_id: 387
Vim script:
  type: programming
  color: "#199f4b"
  tm_scope: source.viml
  aliases:
  - vim
  - viml
  - nvim
  extensions:
  - ".vim"
  filenames:
  - ".nvimrc"
  - ".vimrc"
  - _vimrc
  - gvimrc
  - nvimrc
  - vimrc
  ace_mode: text
  language_id: 388
Visual Basic:
  type: programming
  color: "#945db7"
  extensions:
  - ".vb"
  - ".bas"
  - ".cls"
  - ".frm"
  - ".frx"
  - ".vba"
  - ".vbhtml"
  - ".vbs"
  tm_scope: source.vbnet
  aliases:
  - vb.net
  - vbnet
  ace_mode: text
  codemirror_mode: vb
  codemirror_mime_type: text/x-vb
  language_id: 389
Volt:
  type: programming
  color: "#1F1F1F"
  extensions:
  - ".volt"
  tm_scope: source.d
  ace_mode: d
  codemirror_mode: d
  codemirror_mime_type: text/x-d
  language_id: 390
Vue:
  type: markup
  color: "#2c3e50"
  extensions:
  - ".vue"
  tm_scope: text.html.vue
  ace_mode: html
  language_id: 391
Wavefront Material:
  type: data
  extensions:
  - ".mtl"
  tm_scope: source.wavefront.mtl
  ace_mode: text
  language_id: 392
Wavefront Object:
  type: data
  extensions:
  - ".obj"
  tm_scope: source.wavefront.obj
  ace_mode: text
  language_id: 393
Web Ontology Language:
  type: data
  extensions:
  - ".owl"
  tm_scope: text.xml
  ace_mode: xml
  language_id: 394
WebAssembly:
  type: programming
  color: "#04133b"
  extensions:
  - ".wast"
  - ".wat"
  aliases:
  - wast
  - wasm
  tm_scope: source.webassembly
  ace_mode: lisp
  codemirror_mode: commonlisp
  codemirror_mime_type: text/x-common-lisp
  language_id: 956556503
WebIDL:
  type: programming
  extensions:
  - ".webidl"
  tm_scope: source.webidl
  ace_mode: text
  codemirror_mode: webidl
  codemirror_mime_type: text/x-webidl
  language_id: 395
World of Warcraft Addon Data:
  type: data
  extensions:
  - ".toc"
  tm_scope: source.toc
  ace_mode: text
  language_id: 396
X10:
  type: programming
  aliases:
  - xten
  ace_mode: text
  extensions:
  - ".x10"
  color: "#4B6BEF"
  tm_scope: source.x10
  language_id: 397
XC:
  type: programming
  color: "#99DA07"
  extensions:
  - ".xc"
  tm_scope: source.xc
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csrc
  language_id: 398
XCompose:
  type: data
  filenames:
  - ".XCompose"
  - XCompose
  - xcompose
  tm_scope: config.xcompose
  ace_mode: text
  language_id: 225167241
XML:
  type: data
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  aliases:
  - rss
  - xsd
  - wsdl
  extensions:
  - ".xml"
  - ".adml"
  - ".admx"
  - ".ant"
  - ".axml"
  - ".builds"
  - ".ccproj"
  - ".ccxml"
  - ".clixml"
  - ".cproject"
  - ".cscfg"
  - ".csdef"
  - ".csl"
  - ".csproj"
  - ".ct"
  - ".depproj"
  - ".dita"
  - ".ditamap"
  - ".ditaval"
  - ".dll.config"
  - ".dotsettings"
  - ".filters"
  - ".fsproj"
  - ".fxml"
  - ".glade"
  - ".gml"
  - ".grxml"
  - ".iml"
  - ".ivy"
  - ".jelly"
  - ".jsproj"
  - ".kml"
  - ".launch"
  - ".mdpolicy"
  - ".mjml"
  - ".mm"
  - ".mod"
  - ".mxml"
  - ".natvis"
  - ".ndproj"
  - ".nproj"
  - ".nuspec"
  - ".odd"
  - ".osm"
  - ".pkgproj"
  - ".plist"
  - ".pluginspec"
  - ".proj"
  - ".props"
  - ".ps1xml"
  - ".psc1"
  - ".pt"
  - ".rdf"
  - ".resx"
  - ".rss"
  - ".sch"
  - ".scxml"
  - ".sfproj"
  - ".shproj"
  - ".srdf"
  - ".storyboard"
  - ".stTheme"
  - ".sublime-snippet"
  - ".targets"
  - ".tmCommand"
  - ".tml"
  - ".tmLanguage"
  - ".tmPreferences"
  - ".tmSnippet"
  - ".tmTheme"
  - ".ts"
  - ".tsx"
  - ".ui"
  - ".urdf"
  - ".ux"
  - ".vbproj"
  - ".vcxproj"
  - ".vsixmanifest"
  - ".vssettings"
  - ".vstemplate"
  - ".vxml"
  - ".wixproj"
  - ".wsdl"
  - ".wsf"
  - ".wxi"
  - ".wxl"
  - ".wxs"
  - ".x3d"
  - ".xacro"
  - ".xaml"
  - ".xib"
  - ".xlf"
  - ".xliff"
  - ".xmi"
  - ".xml.dist"
  - ".xproj"
  - ".xsd"
  - ".xspec"
  - ".xul"
  - ".zcml"
  filenames:
  - ".classpath"
  - ".project"
  - App.config
  - NuGet.config
  - Settings.StyleCop
  - Web.Debug.config
  - Web.Release.config
  - Web.config
  - packages.config
  language_id: 399
XPM:
  type: data
  extensions:
  - ".xpm"
  - ".pm"
  ace_mode: c_cpp
  tm_scope: source.c
  language_id: 781846279
XPages:
  type: data
  extensions:
  - ".xsp-config"
  - ".xsp.metadata"
  tm_scope: text.xml
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  language_id: 400
XProc:
  type: programming
  extensions:
  - ".xpl"
  - ".xproc"
  tm_scope: text.xml
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  language_id: 401
XQuery:
  type: programming
  color: "#5232e7"
  extensions:
  - ".xquery"
  - ".xq"
  - ".xql"
  - ".xqm"
  - ".xqy"
  ace_mode: xquery
  codemirror_mode: xquery
  codemirror_mime_type: application/xquery
  tm_scope: source.xq
  language_id: 402
XS:
  type: programming
  extensions:
  - ".xs"
  tm_scope: source.c
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csrc
  language_id: 403
XSLT:
  type: programming
  aliases:
  - xsl
  extensions:
  - ".xslt"
  - ".xsl"
  tm_scope: text.xml.xsl
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  color: "#EB8CEB"
  language_id: 404
Xojo:
  type: programming
  extensions:
  - ".xojo_code"
  - ".xojo_menu"
  - ".xojo_report"
  - ".xojo_script"
  - ".xojo_toolbar"
  - ".xojo_window"
  tm_scope: source.vbnet
  ace_mode: text
  language_id: 405
Xtend:
  type: programming
  extensions:
  - ".xtend"
  ace_mode: text
  language_id: 406
YAML:
  type: data
  tm_scope: source.yaml
  aliases:
  - yml
  extensions:
  - ".yml"
  - ".reek"
  - ".rviz"
  - ".sublime-syntax"
  - ".syntax"
  - ".yaml"
  - ".yaml-tmlanguage"
  - ".yml.mysql"
  filenames:
  - ".clang-format"
  - ".clang-tidy"
  ace_mode: yaml
  codemirror_mode: yaml
  codemirror_mime_type: text/x-yaml
  language_id: 407
YANG:
  type: data
  extensions:
  - ".yang"
  tm_scope: source.yang
  ace_mode: text
  language_id: 408
YARA:
  type: data
  ace_mode: text
  extensions:
  - ".yar"
  - ".yara"
  tm_scope: source.yara
  language_id: 805122868
Yacc:
  type: programming
  extensions:
  - ".y"
  - ".yacc"
  - ".yy"
  tm_scope: source.bison
  ace_mode: text
  color: "#4B6C4B"
  language_id: 409
Zephir:
  type: programming
  color: "#118f9e"
  extensions:
  - ".zep"
  tm_scope: source.php.zephir
  ace_mode: php
  language_id: 410
Zimpl:
  type: programming
  extensions:
  - ".zimpl"
  - ".zmpl"
  - ".zpl"
  tm_scope: none
  ace_mode: text
  language_id: 411
desktop:
  type: data
  extensions:
  - ".desktop"
  - ".desktop.in"
  tm_scope: source.desktop
  ace_mode: text
  language_id: 412
eC:
  type: programming
  color: "#913960"
  extensions:
  - ".ec"
  - ".eh"
  tm_scope: source.c.ec
  ace_mode: text
  language_id: 413
edn:
  type: data
  ace_mode: clojure
  codemirror_mode: clojure
  codemirror_mime_type: text/x-clojure
  extensions:
  - ".edn"
  tm_scope: source.clojure
  language_id: 414
fish:
  type: programming
  group: Shell
  interpreters:
  - fish
  extensions:
  - ".fish"
  tm_scope: source.fish
  ace_mode: text
  language_id: 415
mupad:
  type: programming
  extensions:
  - ".mu"
  ace_mode: text
  language_id: 416
nesC:
  type: programming
  color: "#94B0C7"
  extensions:
  - ".nc"
  ace_mode: text
  tm_scope: source.nesc
  language_id: 417
ooc:
  type: programming
  color: "#b0b77e"
  extensions:
  - ".ooc"
  ace_mode: text
  language_id: 418
reStructuredText:
  type: prose
  wrap: true
  aliases:
  - rst
  extensions:
  - ".rst"
  - ".rest"
  - ".rest.txt"
  - ".rst.txt"
  ace_mode: text
  codemirror_mode: rst
  codemirror_mime_type: text/x-rst
  language_id: 419
wdl:
  type: programming
  color: "#42f1f4"
  extensions:
  - ".wdl"
  tm_scope: source.wdl
  ace_mode: text
  language_id: 374521672
wisp:
  type: programming
  ace_mode: clojure
  codemirror_mode: clojure
  codemirror_mime_type: text/x-clojure
  color: "#7582D1"
  extensions:
  - ".wisp"
  tm_scope: source.clojure
  language_id: 420
xBase:
  type: programming
  color: "#403a40"
  aliases:
  - advpl
  - clipper
  - foxpro
  extensions:
  - ".prg"
  - ".ch"
  - ".prw"
  tm_scope: source.harbour
  ace_mode: text
  language_id: 421
`,

	"data/vendor.yml": `# Vendored files and directories are excluded from language
# statistics.
#
# Lines in this file are Regexps that are matched against the file
# pathname.
#
# Please add additional test coverage to
# ` + "`" + `test/test_blob.rb#test_vendored` + "`" + ` if you make any changes.

## Vendor Conventions ##

# Caches
- (^|/)cache/

# Dependencies
- ^[Dd]ependencies/

# Distributions
- (^|/)dist/

# C deps
- ^deps/
- (^|/)configure$
- (^|/)config.guess$
- (^|/)config.sub$

# stuff autogenerated by autoconf - still C deps
- (^|/)aclocal.m4
- (^|/)libtool.m4
- (^|/)ltoptions.m4
- (^|/)ltsugar.m4
- (^|/)ltversion.m4
- (^|/)lt~obsolete.m4

# Linters
- cpplint.py

# Node dependencies
- node_modules/

# Bower Components
- bower_components/

# Erlang bundles
- ^rebar$
- erlang.mk

# Go dependencies
- Godeps/_workspace/

# GNU indent profiles
- .indent.pro

# Minified JavaScript and CSS
- (\.|-)min\.(js|css)$

# Stylesheets imported from packages
- ([^\s]*)import\.(css|less|scss|styl)$

# Bootstrap css and js
- (^|/)bootstrap([^.]*)\.(js|css|less|scss|styl)$
- (^|/)custom\.bootstrap([^\s]*)(js|css|less|scss|styl)$

# Font Awesome
- (^|/)font-awesome\.(css|less|scss|styl)$
- (^|/)font-awesome/.*\.(css|less|scss|styl)$

# Foundation css
- (^|/)foundation\.(css|less|scss|styl)$

# Normalize.css
- (^|/)normalize\.(css|less|scss|styl)$

# Skeleton.css
- (^|/)skeleton\.(css|less|scss|styl)$

# Bourbon css
- (^|/)[Bb]ourbon/.*\.(css|less|scss|styl)$

# Animate.css
- (^|/)animate\.(css|less|scss|styl)$

# Materialize.css
- (^|/)materialize\.(css|less|scss|styl|js)$

# Select2
- (^|/)select2/.*\.(css|scss|js)$

# Vendored dependencies
- third[-_]?party/
- 3rd[-_]?party/
- vendors?/
- extern(al)?/
- (^|/)[Vv]+endor/

# Debian packaging
- ^debian/

# Haxelib projects often contain a neko bytecode file named run.n
- run.n$

# Bootstrap Datepicker
- bootstrap-datepicker/

## Commonly Bundled JavaScript frameworks ##

# jQuery
- (^|/)jquery([^.]*)\.js$
- (^|/)jquery\-\d\.\d+(\.\d+)?\.js$

# jQuery UI
- (^|/)jquery\-ui(\-\d\.\d+(\.\d+)?)?(\.\w+)?\.(js|css)$
- (^|/)jquery\.(ui|effects)\.([^.]*)\.(js|css)$

# jQuery Gantt
- jquery.fn.gantt.js

# jQuery fancyBox
- jquery.fancybox.(js|css)

# Fuel UX
- fuelux.js

# jQuery File Upload
- (^|/)jquery\.fileupload(-\w+)?\.js$

# jQuery dataTables
- jquery.dataTables.js

# bootboxjs
- bootbox.js

# pdf-worker
- pdf.worker.js

# Slick
- (^|/)slick\.\w+.js$

# Leaflet plugins
- (^|/)Leaflet\.Coordinates-\d+\.\d+\.\d+\.src\.js$
- leaflet.draw-src.js
- leaflet.draw.css
- Control.FullScreen.css
- Control.FullScreen.js
- leaflet.spin.js
- wicket-leaflet.js

# Sublime Text workspace files
- .sublime-project
- .sublime-workspace

# VS Code workspace files
- .vscode

# Prototype
- (^|/)prototype(.*)\.js$
- (^|/)effects\.js$
- (^|/)controls\.js$
- (^|/)dragdrop\.js$

# Typescript definition files
- (.*?)\.d\.ts$

# MooTools
- (^|/)mootools([^.]*)\d+\.\d+.\d+([^.]*)\.js$

# Dojo
- (^|/)dojo\.js$

# MochiKit
- (^|/)MochiKit\.js$

# YUI
- (^|/)yahoo-([^.]*)\.js$
- (^|/)yui([^.]*)\.js$

# WYS editors
- (^|/)ckeditor\.js$
- (^|/)tiny_mce([^.]*)\.js$
- (^|/)tiny_mce/(langs|plugins|themes|utils)

# Ace Editor
- (^|/)ace-builds/

# Fontello CSS files
- (^|/)fontello(.*?)\.css$

# MathJax
- (^|/)MathJax/

# Chart.js
- (^|/)Chart\.js$

# CodeMirror
- (^|/)[Cc]ode[Mm]irror/(\d+\.\d+/)?(lib|mode|theme|addon|keymap|demo)

# SyntaxHighlighter - http://alexgorbatchev.com/
- (^|/)shBrush([^.]*)\.js$
- (^|/)shCore\.js$
- (^|/)shLegacy\.js$

# AngularJS
- (^|/)angular([^.]*)\.js$

# D3.js
- (^|\/)d3(\.v\d+)?([^.]*)\.js$

# React
- (^|/)react(-[^.]*)?\.js$

# flow-typed
- (^|/)flow-typed/.*\.js$

# Modernizr
- (^|/)modernizr\-\d\.\d+(\.\d+)?\.js$
- (^|/)modernizr\.custom\.\d+\.js$

# Knockout
- (^|/)knockout-(\d+\.){3}(debug\.)?js$

## Python ##

# Sphinx
- (^|/)docs?/_?(build|themes?|templates?|static)/

# django
- (^|/)admin_media/
- (^|/)env/

# Fabric
- ^fabfile\.py$

# WAF
- ^waf$

# .osx
- ^.osx$

## Obj-C ##

# Xcode

- \.xctemplate/
- \.imageset/

# Carthage
- (^|/)Carthage/

# Sparkle
- (^|/)Sparkle/

# Crashlytics
- Crashlytics.framework/

# Fabric
- Fabric.framework/

# BuddyBuild
- BuddyBuildSDK.framework/

# Realm
- Realm.framework

# RealmSwift
- RealmSwift.framework

# git config files
- gitattributes$
- gitignore$
- gitmodules$

## Groovy ##

# Gradle
- (^|/)gradlew$
- (^|/)gradlew\.bat$
- (^|/)gradle/wrapper/

## .NET ##

# Visual Studio IntelliSense
- -vsdoc\.js$
- \.intellisense\.js$

# jQuery validation plugin (MS bundles this with asp.net mvc)
- (^|/)jquery([^.]*)\.validate(\.unobtrusive)?\.js$
- (^|/)jquery([^.]*)\.unobtrusive\-ajax\.js$

# Microsoft Ajax
- (^|/)[Mm]icrosoft([Mm]vc)?([Aa]jax|[Vv]alidation)(\.debug)?\.js$

# NuGet
- ^[Pp]ackages\/.+\.\d+\/

# ExtJS
- (^|/)extjs/.*?\.js$
- (^|/)extjs/.*?\.xml$
- (^|/)extjs/.*?\.txt$
- (^|/)extjs/.*?\.html$
- (^|/)extjs/.*?\.properties$
- (^|/)extjs/.sencha/
- (^|/)extjs/docs/
- (^|/)extjs/builds/
- (^|/)extjs/cmd/
- (^|/)extjs/examples/
- (^|/)extjs/locale/
- (^|/)extjs/packages/
- (^|/)extjs/plugins/
- (^|/)extjs/resources/
- (^|/)extjs/src/
- (^|/)extjs/welcome/

# Html5shiv
- (^|/)html5shiv\.js$

# Test fixtures
- ^[Tt]ests?/fixtures/
- ^[Ss]pecs?/fixtures/

# PhoneGap/Cordova
- (^|/)cordova([^.]*)\.js$
- (^|/)cordova\-\d\.\d(\.\d)?\.js$

# Foundation js
- foundation(\..*)?\.js$

# Vagrant
- ^Vagrantfile$

# .DS_Stores
- .[Dd][Ss]_[Ss]tore$

# R packages
- ^vignettes/
- ^inst/extdata/

# Octicons
- octicons.css
- sprockets-octicons.scss

# Typesafe Activator
- (^|/)activator$
- (^|/)activator\.bat$

# ProGuard
- proguard.pro
- proguard-rules.pro

# PuPHPet
- ^puphpet/

# Android Google APIs
- (^|/)\.google_apis/

# Jenkins Pipeline
- ^Jenkinsfile$
`,

	"data/documentation.yml": `# Documentation files and directories are excluded from language
# statistics.
#
# Lines in this file are Regexps that are matched against the file
# pathname.
#
# Please add additional test coverage to
# ` + "`" + `test/test_blob.rb#test_documentation` + "`" + ` if you make any changes.

## Documentation directories ##

- ^[Dd]ocs?/
- (^|/)[Dd]ocumentation/
- (^|/)[Jj]avadoc/
- ^[Mm]an/
- ^[Ee]xamples/
- ^[Dd]emos?/

## Documentation files ##

- (^|/)CHANGE(S|LOG)?(\.|$)
- (^|/)CONTRIBUTING(\.|$)
- (^|/)COPYING(\.|$)
- (^|/)INSTALL(\.|$)
- (^|/)LICEN[CS]E(\.|$)
- (^|/)[Ll]icen[cs]e(\.|$)
- (^|/)README(\.|$)
- (^|/)[Rr]eadme(\.|$)

# Samples folders
- ^[Ss]amples?/
`,
}
